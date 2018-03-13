package cdb

import (
	"io/ioutil"
	"encoding/json"
	"github.com/seefan/tablekv/common"
	"time"
	"sync"
	log "github.com/cihub/seelog"
)

//Maintain node data information
type ClusterDB struct {
	node  map[string]string
	table map[string]string
	file  string
	clock *time.Timer
	lock  sync.Mutex
}

func (c *ClusterDB) SetTable(name string) {
	c.table[name] = "localhost"
}
func (c *ClusterDB) GetTables() map[string]string {
	return c.table
}

//start
func (c *ClusterDB) Start(path string) (err error) {
	c.file = path
	c.table = make(map[string]string)
	c.node = make(map[string]string)
	if err = c.load(); err == nil {
		go c.timer()
		log.Info("ClusterDB is started")
	}
	c.table["test"]="localhost"
	return
}
func (c *ClusterDB) Close() {
	if err := c.flush(); err != nil {
		log.Error("ClusterDB close error", err)
	}
	c.clock.Stop()
	log.Info("ClusterDB is stoped")
}

//Loading data
func (c *ClusterDB) load() error {
	if common.FileIsNotExist(c.file) {
		c.flush()
	} else {
		bs, err := ioutil.ReadFile(c.file)
		if err != nil {
			return err
		}
		err = json.Unmarshal(bs, c)
		if err != nil {
			return err
		}
	}
	return nil
}

//Saving data
func (c *ClusterDB) flush() error {
	c.lock.Lock()
	defer c.lock.Unlock()
	if bs, err := json.Marshal(c); err == nil {
		ioutil.WriteFile(c.file, bs, 0764)
		return nil
	} else {
		return err
	}
}

//Timer
func (c *ClusterDB) timer() {
	c.clock = time.NewTimer(time.Minute)
	for range c.clock.C {
		if err := c.flush(); err != nil {
			log.Error("ClusterDB flush data error", err)
		}
	}
}
