package tables

import (
	"sync"
	log "github.com/cihub/seelog"
	"path"
	"os"
	"github.com/seefan/tablekv/common"
)

type TableManager struct {
	tableManager  map[string]*Table
	lock          sync.RWMutex
	path          string
	NewTableEvent func(name string)
}
//close all table
func (t *TableManager) Close() (err error) {
	for _, tb := range t.tableManager {
		if err = tb.Close(); err != nil {
			log.Error("close table error", err)
		}
	}
	return
}
//get a table
func (t *TableManager) GetTable(name string) (table *Table, err error) {
	t.lock.Lock()
	defer t.lock.Unlock()
	table, ok := t.tableManager[name]
	if !ok {
		if table, err = LoadTable(t.path, name); err != nil {
			return nil, err
		} else {
			t.tableManager[name] = table
			//synchronize to cdb
			if t.NewTableEvent != nil {
				t.NewTableEvent(name)
			}
		}
	}

	return
}
//delete a table
func (t *TableManager) DeleteTable(name string) (err error) {
	t.lock.Lock()
	defer t.lock.Unlock()
	table, ok := t.tableManager[name]
	//delete dir

	if ok {
		err = table.Close()
		delete(t.tableManager, name)
	}
	dir := path.Join(t.path, name)
	if err == nil && !common.FileIsNotExist(dir) {
		err = os.RemoveAll(dir)
	}
	return
}
//create new table manager
func NewTableManager(path string, tables []string) (t *TableManager) {
	t = &TableManager{
		tableManager: make(map[string]*Table),
		path:         path,
	}
	if tables == nil {
		return
	}
	for _, name := range tables {
		if table, err := LoadTable(path, name); err == nil {
			log.Debugf("load table %s", name)
			t.tableManager[name] = table
		} else {
			log.Errorf("load table %s error", name, err)
		}
	}
	return
}
