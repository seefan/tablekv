package main

import (
	"github.com/seefan/tablekv/common"
	"os"

	"syscall"
	"os/signal"
	"github.com/seefan/tablekv/boot"
	"flag"
	"path"
	"github.com/cihub/seelog"
)

func main() {
	defer common.PrintErr()
	confPath := flag.String("config", "./conf.ini", "conf.ini path")
	b := &boot.Boot{}
	cfg := b.LoadConfig(*confPath)
	//init log config and log file
	common.InitLog(path.Join(cfg.LogPath, "log.xml"), path.Join(cfg.LogPath, "tk.log"))
	defer seelog.Flush()
	seelog.Debug("config loaded", cfg.ToString())
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT, os.Interrupt)
	seelog.Info("TableKV is starting")
	go func() {
		if err := b.Start(); err != nil {
			seelog.Info("TableKV startup error", err)
			sig <- syscall.SIGABRT
		}
	}()

	s := <-sig

	seelog.Info("received signal ", s)
	b.Close()
	seelog.Info("TableKV is closed")
}
