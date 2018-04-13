package boot

import (
	log "github.com/cihub/seelog"
	"github.com/seefan/tablekv/protocol/thrift_protocol"
	"github.com/seefan/tablekv/cdb"
	"github.com/seefan/tablekv/tables"
	"github.com/seefan/tablekv/common"
	"github.com/seefan/tablekv/process"
)

type Boot struct {
	db  *cdb.ClusterDB
	tm  *tables.TableManager
	cnl common.NetLayout
	cfg *common.Config
}

func (b *Boot) Start() error {
	//start center db
	b.db = new(cdb.ClusterDB)
	if err := b.db.Start(b.cfg.VarPath); err != nil {
		log.Error("clusterDB start error", err)
	}
	log.Debug("clusterDB loaded")
	//start table manager and load table

	tbs, err := b.db.GetLocalTables()
	if err != nil {
		log.Errorf("load local table error")
	}
	b.tm = tables.NewTableManager(b.cfg, tbs)
	b.tm.TableEvent = func(name string, eventType byte) {
		if eventType == 0 {
			if err := b.db.SetTable(name); err != nil {
				if err = b.db.SetTable(name); err != nil {
					log.Error("write cdb error", err)
				}
			}
		} else {
			if err := b.db.RemoveTable(name); err != nil {
				if err = b.db.RemoveTable(name); err != nil {
					log.Error("remove table from cdb is error", err)
				}
			}
		}
	}
	log.Debug("Table Manager loaded")
	//create processor
	pm := process.NewProcessorManager(b.tm)

	b.cnl = &thrift_protocol.Thrift{}
	log.Debug("Process Manager loaded")

	if err := b.cnl.Start(pm, b.cfg.Host, b.cfg.Port); err != nil {
		return err
	}
	return nil
}
func (b *Boot) Close() {
	if err := b.cnl.Stop(); err != nil {
		log.Error("stop processor error", err)
	}

	if err := b.tm.Close(); err != nil {
		log.Error("stop table manager error", err)
	}
	if err := b.db.Close(); err != nil {
		log.Error("stop cdb error", err)
	}

}