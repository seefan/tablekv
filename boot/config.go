package boot

import (
	"github.com/go-ini/ini"
	"os"
	"github.com/seefan/tablekv/common"
)

func (b *Boot) LoadConfig(confPath string) *common.Config {
	//load config file
	//confPath := flag.String("config", "./conf.ini", "conf.ini path")
	cfg := new(common.Config)
	if file, err := ini.Load(confPath); err == nil {
		cfg.Load(file)
	} else {
		cfg.Load(nil)
	}
	common.WriteBuffer = cfg.WriteBuffer
	switch cfg.TimeoutType {
	case 0:
		common.Timeout = float64(cfg.Timeout)
	case 1:
		common.Timeout = float64(cfg.Timeout * 24)
	default:
		common.Timeout = 1
	}
	if common.Timeout < 1 {
		common.Timeout = 1
	}

	//create log  directory
	if common.FileIsNotExist(cfg.LogPath) {
		os.MkdirAll(cfg.LogPath, 0764)
	}
	//create data  directory
	if common.FileIsNotExist(cfg.VarPath) {
		os.MkdirAll(cfg.VarPath, 0764)
	}
	b.cfg = cfg
	return cfg
}
