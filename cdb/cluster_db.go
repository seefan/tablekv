package cdb

import (
	log "github.com/cihub/seelog"
	"github.com/seefan/tablekv/tables"
	"time"
)

const (
	TimeFormat = "20060102150405"
)

//Maintain node data information
type ClusterDB struct {
	data *tables.Table
}

//set table info
func (c *ClusterDB) SetTable(name string) (error) {
	if exists, err := c.data.Exists([]byte(name)); err == nil && exists {
		return nil
	}
	t := tables.TableInfo{
		Name:       name,
		Host:       "localhost",
		LastUpdate: time.Now(),
	}
	return c.data.Set([]byte(name), t.ToByte())
}

//get all local tables
func (c *ClusterDB) GetLocalTables() (re []*tables.TableInfo, err error) {
	ts, err := c.data.Scan(nil, nil)
	if err != nil {
		return nil, err
	}
	for _, v := range ts {
		ti := new(tables.TableInfo)
		if err = ti.FromByte([]byte(v.Value)); err == nil && ti.Host == "localhost" {
			re = append(re, ti)
		} else {
			log.Debug(err)
		}
	}
	return
}

//get table info
func (c *ClusterDB) GetTable(name string) (*tables.TableInfo, error) {
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
	return c.data.Delete([]byte(name))
}

//start
func (c *ClusterDB) Start(path string) (err error) {
	c.data, err = tables.LoadTable(path, "cdb")
	if err != nil {
		return err
	}
	log.Info("ClusterDB is start")
	return
}

//stop
func (c *ClusterDB) Close() (err error) {
	if c.data != nil {
		err = c.data.Close()
	}
	log.Info("ClusterDB is stop")
	return
}
