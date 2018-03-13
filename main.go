package main

import (
	"github.com/go-ini/ini"
	"flag"
	"github.com/seefan/tablekv/common"
	"github.com/cihub/seelog"
	"path"
	"os"
	"github.com/seefan/tablekv/tables"
	"github.com/seefan/tablekv/cdb"
	//"github.com/seefan/tablekv/processor"
	"github.com/seefan/tablekv/processor"
	"github.com/seefan/tablekv/processor/thrift_protocol"
	"syscall"
	"os/signal"
)

func main() {
	defer common.PrintErr()
	//load config file
	confPath := flag.String("config", "./conf.ini", "conf.ini path")
	cfg := new(common.Config)
	if file, err := ini.Load(confPath); err == nil {
		cfg.Load(file)
	} else {
		cfg.Load(nil)
	}
	//create log  directory
	if common.FileIsNotExist(cfg.LogPath) {
		os.MkdirAll(cfg.LogPath, 0764)
	}
	//create data  directory
	if common.FileIsNotExist(cfg.VarPath) {
		os.MkdirAll(cfg.VarPath, 0764)
	}
	//init log config and log file
	common.InitLog(path.Join(cfg.LogPath, "log.xml"), path.Join(cfg.LogPath, "tk.log"))
	defer seelog.Flush()
	seelog.Info(cfg.ToString())
	//start center db
	db := new(cdb.ClusterDB)
	if err := db.Start(cfg.VarPath); err != nil {
		seelog.Error("clusterDB start error", err)
		panic(err)
	}
	//start table manager and load table
	tm := tables.NewTableManager(path.Join(cfg.VarPath, "tables"), nil)
	tm.NewTableEvent = func(name string) {
		if err := db.SetTable(name); err != nil {
			if err = db.SetTable(name); err != nil {
				seelog.Error("write cdb error", err)
			}
		}
	}
	//create processor
	var pd processor.ProcessStart
	pd = &thrift_protocol.Thrift{}

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		if err := pd.Start(tm, cfg.Host, cfg.Port); err != nil {
			seelog.Error("start processor error", err)
			sig <- syscall.SIGABRT
		}
	}()

	seelog.Info("TableKV started")
	<-sig
	if err := pd.Stop(); err != nil {
		seelog.Error("stop processor error", err)
	}
	if err := tm.Close(); err != nil {
		seelog.Error("stop table manager error", err)
	}
	if err := db.Close(); err != nil {
		seelog.Error("stop cdb error", err)
	}
	seelog.Info("TableKV stoped")

}
