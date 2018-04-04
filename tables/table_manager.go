package tables

import (
	"sync"
	log "github.com/cihub/seelog"
	"path"
	"os"
	"github.com/seefan/tablekv/common"
	"time"
)

type TableManager struct {
	tableMap   map[string]*Table
	lock       sync.RWMutex
	path       string
	conf       *common.Config
	TableEvent func(name string, eventType byte)
	timer      *time.Ticker
	timeout    float64
}

//close all table
func (t *TableManager) Close() (err error) {
	for _, tb := range t.tableMap {
		if err = tb.Close(); err != nil {
			log.Error("close table error", err)
		}
	}
	if t.timer != nil {
		t.timer.Stop()
	}
	return
}

//get a table
func (t *TableManager) GetTable(name string) (table *Table, err error) {
	t.lock.Lock()
	defer t.lock.Unlock()
	table, ok := t.tableMap[name]
	if !ok {
		if table, err = LoadTable(t.path, name); err != nil {
			return nil, err
		} else {
			t.tableMap[name] = table
			//synchronize to cdb
			if t.TableEvent != nil {
				t.TableEvent(name, 0) //new table is 0
			}
		}
	}
	return
}

//delete a table
func (t *TableManager) DeleteTable(name string) (err error) {
	t.lock.Lock()
	defer t.lock.Unlock()
	table, ok := t.tableMap[name]
	//delete dir
	log.Debugf("delete table %s", name)

	if ok {
		err = table.Close()
		delete(t.tableMap, name)
	}
	dir := path.Join(t.path, name)
	if err == nil && !common.FileIsNotExist(dir) {
		err = os.RemoveAll(dir)
	}
	if err == nil && t.TableEvent != nil {
		t.TableEvent(name, 1) //delete table is 1
	}
	return
}

//create new table manager
func NewTableManager(cfg *common.Config, tables []string) (t *TableManager) {
	t = &TableManager{
		tableMap: make(map[string]*Table),
		conf:     cfg,
		path:     path.Join(cfg.VarPath, "tables"),
		timer:    time.NewTicker(time.Hour),
	}
	if common.FileIsNotExist(t.path){
		os.MkdirAll(t.path,0764)
	}
	go t.timeProcessor()
	if tables == nil {
		return
	}
	for _, name := range tables {
		log.Debugf("load table %s", name)
		if _, ok := t.tableMap[name]; !ok {
			if table, err := LoadTable(t.path, name); err == nil {
				t.tableMap[name] = table
			} else {
				log.Errorf("load table %s error", name, err)
			}
		}
	}
	switch cfg.TimeoutType {
	case 0:
		t.timeout = float64(cfg.Timeout)
	case 1:
		t.timeout = float64(cfg.Timeout * 24)
	default:
		t.timeout = 1
	}
	if t.timeout < 1 {
		t.timeout = 1
	}
	return
}
