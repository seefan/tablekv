package cdb

import (
	log "github.com/cihub/seelog"
	"github.com/seefan/tablekv/tables"
	"time"
	"sync"
	"github.com/seefan/goerr"
)



//Maintain node data information
type ClusterDB struct {
	data *tables.Table
	lock sync.Mutex
}

//set table info
func (c *ClusterDB) SetTable(name string, info []byte) (error) {
	c.lock.Lock()
	defer c.lock.Unlock()
	if c.data == nil {
		return goerr.New("cdb not init")
	}
	return c.data.Set([]byte(name), info)
}

//get all local tables
func (c *ClusterDB) GetLocalTables() (re []*tables.TableInfo, err error) {
	if c.data == nil {
		return nil, goerr.New("cdb not init")
	}
	ks,vs, err := c.data.Scan(nil, nil, 0)
	if err != nil {
		return nil, err
	}
	for i := range ks {
		ti := new(tables.TableInfo)
		if err := ti.FromByte([]byte(vs[i])); err == nil && ti.Host == "localhost" {
			re = append(re, ti)
		} else {
			log.Error(err)
		}
	}
	return
}

//get table info
func (c *ClusterDB) GetTable(name string) (*tables.TableInfo, error) {
	if c.data == nil {
		return nil, goerr.New("cdb not init")
	}
	c.lock.Lock()
	defer c.lock.Unlock()
	if bs, err := c.data.Get([]byte(name)); err == nil {
		tb := new(tables.TableInfo)
		tb.FromByte(bs)
		return tb, nil
	} else {
		return nil, err
	}
}

//remove table info
func (c *ClusterDB) RemoveTable(name string) (err error) {
	if c.data == nil {
		return goerr.New("cdb not init")
	}
	c.lock.Lock()
	defer c.lock.Unlock()
	return c.data.Delete([]byte(name))
}

//start
func (c *ClusterDB) Start(path string) (err error) {
	c.lock.Lock()
	defer c.lock.Unlock()
	c.data, err = tables.LoadTable(path, tables.TableInfo{
		Name:       "cdb",
		Host:       "localhost",
		CreateTime: time.Now(),
	})
	if err != nil {
		return err
	}
	log.Info("ClusterDB is start")
	return
}

//stop
func (c *ClusterDB) Close() (err error) {
	c.lock.Lock()
	defer c.lock.Unlock()
	if c.data != nil {
		err = c.data.Close()
	}
	log.Info("ClusterDB is stop")
	return
}
