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
	common.BlockBuffer = cfg.BlockBuffer
	switch cfg.ExpiredType {
	case 1:
		common.Expired = float64(cfg.ExpiredNum)
	case 2:
		common.Expired = float64(cfg.ExpiredNum * 24)
	default:
		common.Expired = 1
	}
	if common.Expired < 1 {
		common.Expired = 1
	}

	//create log  directory
	if common.FileIsNotExist(cfg.LogPath) {
		if err := os.MkdirAll(cfg.LogPath, 0764); err != nil {
			println("make log dir has error", err)
		}
	}
	//create data  directory
	if common.FileIsNotExist(cfg.VarPath) {
		if err := os.MkdirAll(cfg.VarPath, 0764); err != nil {
			panic("make var dir has error:" + err.Error())
		}
	}
	b.cfg = cfg
	return cfg
}
