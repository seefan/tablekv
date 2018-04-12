package main

import (
	"github.com/seefan/tablekv/common"
	"github.com/cihub/seelog"
	"path"
	"os"
	"github.com/seefan/tablekv/cdb"
	"github.com/seefan/tablekv/tables"
	"github.com/seefan/tablekv/process"
	"github.com/seefan/tablekv/protocol/thrift_protocol"
	"syscall"
	"os/signal"
	"github.com/seefan/tablekv/boot"
)

func main() {
	defer common.PrintErr()
	cfg := boot.LoadConfig()
	//init log config and log file
	common.InitLog(path.Join(cfg.LogPath, "log.xml"), path.Join(cfg.LogPath, "tk.log"))
	defer seelog.Flush()
	seelog.Debug("config loaded", cfg.ToString())
	//start center db
	db := new(cdb.ClusterDB)
	if err := db.Start(cfg.VarPath); err != nil {
		seelog.Error("clusterDB start error", err)
		panic(err)
	}
	seelog.Debug("clusterDB loaded")
	//start table manager and load table

	tbs, err := db.GetLocalTables()
	if err != nil {
		seelog.Errorf("load local table error")
	}
	tm := tables.NewTableManager(cfg, tbs)
	tm.TableEvent = func(name string, eventType byte) {
		if eventType == 0 {
			if err := db.SetTable(name); err != nil {
				if err = db.SetTable(name); err != nil {
					seelog.Error("write cdb error", err)
				}
			}
		} else {
			if err := db.RemoveTable(name); err != nil {
				if err = db.RemoveTable(name); err != nil {
					seelog.Error("remove table from cdb is error", err)
				}
			}
		}
	}
	seelog.Debug("Table Manager loaded")
	//create processor
	pm := process.NewProcessorManager(tm)

	var pd common.NetLayout
	pd = &thrift_protocol.Thrift{}
	seelog.Debug("Process Manager loaded")
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		if err := pd.Start(pm, cfg.Host, cfg.Port); err != nil {
			seelog.Error("start processor error", err)
			sig <- syscall.SIGABRT
		}
	}()
	defer func() {
		if err := pd.Stop(); err != nil {
			seelog.Error("stop processor error", err)
		}

		if err := tm.Close(); err != nil {
			seelog.Error("stop table manager error", err)
		}
		if err := db.Close(); err != nil {
			seelog.Error("stop cdb error", err)
		}
		seelog.Info("TableKV is stop")
	}()
	seelog.Info("TableKV started")
	<-sig
}
