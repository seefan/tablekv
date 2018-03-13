package cdb

import (
	log "github.com/cihub/seelog"
	"github.com/seefan/tablekv/tables"
)

//Maintain node data information
type ClusterDB struct {
	data *tables.Table
}

func (c *ClusterDB) SetTable(name string) error {
	return c.data.Set([]byte(name), []byte("localhost"))
}
func (c *ClusterDB) GetTable(name string) (string, error) {
	if bs, err := c.data.Get([]byte(name)); err == nil {
		return string(bs), nil
	} else {
		return "", err
	}
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
func (c *ClusterDB) Close() (err error) {
	if c.data != nil {
		err = c.data.Close()
	}
	log.Info("ClusterDB is stop")
	return
}
