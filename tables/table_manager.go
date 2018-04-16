package tables

import (
	"sync"
	log "github.com/cihub/seelog"
	"path"
	"os"
	"github.com/seefan/tablekv/common"
	"time"
)

const (
	TimeFormat  = "2006-01-02 15:04:05"
	NewTable    = 0
	DeleteTable = 1
)

type TableManager struct {
	tableMap   map[string]*Table
	lock       sync.RWMutex
	path       string
	conf       *common.Config
	TableEvent func(name string, eventType byte)
	timer      *time.Ticker
	now        time.Time
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
		if table, err = LoadTable(t.path, common.HashString(name)); err != nil {
			return nil, err
		} else {
			t.tableMap[name] = table
			//synchronize to cdb
			if t.TableEvent != nil {
				t.TableEvent(name, NewTable) //new table is 0
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
	log.Debugf("delete table %s # %s", name, table.createTime.Format(TimeFormat))

	if ok {
		err = table.Close()
		delete(t.tableMap, name)
	}
	if err == nil && !common.FileIsNotExist(table.path) {
		err = os.RemoveAll(table.path)
	}
	if err == nil && t.TableEvent != nil {
		t.TableEvent(name, DeleteTable)
	}
	return
}

//create new table manager
func NewTableManager(cfg *common.Config, tables []*TableInfo) (t *TableManager) {
	t = &TableManager{
		tableMap: make(map[string]*Table),
		conf:     cfg,
		path:     path.Join(cfg.VarPath, "tables"),
		timer:    time.NewTicker(time.Minute),
		now:      time.Now(),
	}
	if common.FileIsNotExist(t.path) {
		if err := os.MkdirAll(t.path, 0764); err != nil {
			log.Error("MkdirAll error", err)
		}
	}
	go t.timeProcessor()
	if tables == nil {
		return
	}
	for _, tb := range tables {
		if _, ok := t.tableMap[tb.Name]; !ok {
			if table, err := LoadTable(t.path, common.HashString(tb.Name)); err == nil {
				table.createTime = tb.CreateTime
				t.tableMap[tb.Name] = table
				log.Debugf("load table %s # %s", tb.Name, tb.CreateTime.Format(TimeFormat))
			} else {
				log.Errorf("load table %s error", tb.Name, err)
			}
		}
	}
	return
}
