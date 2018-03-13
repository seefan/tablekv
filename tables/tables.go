package tables

import (
	"sync"
	log "github.com/cihub/seelog"
)

type TableManager struct {
	tableManager  map[string]*Table
	lock          sync.RWMutex
	path          string
	NewTableEvent func(name string)
}

func (t *TableManager) GetTable(name string) (table *Table, err error) {
	t.lock.Lock()
	defer t.lock.Unlock()
	table, ok := t.tableManager[name]
	if !ok {
		if table, err = loadTable(t.path, name); err != nil {
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
func NewTableManager(path string, tables map[string]string) *TableManager {
	t := &TableManager{
		tableManager: make(map[string]*Table),
		path:         path,
	}
	for name := range tables {
		if table, err := loadTable(path, name); err == nil {
			t.tableManager[name] = table
		} else {
			log.Errorf("load table %s error", name, err)
		}
	}
	return t
}
