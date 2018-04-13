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
	"os/exec"
	"time"
)

func main() {
	defer common.PrintErr()
	confPath := flag.String("config", "./conf.ini", "conf.ini path")
	b := &boot.Boot{}
	cfg := b.LoadConfig(*confPath)
	//init log config and log file
	common.InitLog(path.Join(cfg.LogPath, "log.xml"), path.Join(cfg.LogPath, "tk.log"))
	defer seelog.Flush()
	//
	cmd := "start"
	if len(os.Args) > 1 {
		cmd = os.Args[1]
	}
	switch cmd {
	case "help", "h":
		println("usage: tablekv [help,h,stop] [config=/var/conf.ini]")
	case "stop":
		stop(cfg)
	default:
		start(cfg, b)
	}

}
func stop(cfg *common.Config) {
	if pid, err := common.GetPid(path.Join(cfg.VarPath, "run.pid")); err == nil {
		checkCmd := exec.Command("kill", "-s", "0", pid)
		killCmd := exec.Command("kill", "-s", "USR1", pid)
		now := time.Now()
		if err := killCmd.Run(); err == nil {
			for {
				if err := checkCmd.Run(); err != nil {
					break
				}
				time.Sleep(time.Millisecond * 100)
				if time.Since(now).Seconds() > 30 {
					break
				}
			}
		} else {
			seelog.Error("TableKV stop error", err)
		}
	}
}
func start(cfg *common.Config, b *boot.Boot) {
	seelog.Debug("config loaded", cfg.ToString())
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM, syscall.SIGUSR1)
	seelog.Info("TableKV is starting")
	go func() {
		if err := b.Start(); err != nil {
			seelog.Info("TableKV startup error", err)
			sig <- syscall.SIGABRT
		}
	}()
	common.SavePid(path.Join(cfg.VarPath, "run.pid"))
	s := <-sig

	seelog.Info("received signal ", s)
	b.Close()
	seelog.Info("TableKV is closed")
}
