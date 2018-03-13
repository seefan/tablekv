package common

import (
	"github.com/go-ini/ini"
	"encoding/json"
)

type Config struct {
	Host    string
	Port    int
	VarPath string
	LogPath string
	IsMaster bool
}

func (c *Config) Load(f *ini.File) {
	if f == nil {
		f = ini.Empty()
	}
	c.Host = f.Section("main").Key("host").MustString("127.0.0.1")
	c.Port = f.Section("main").Key("port").MustInt(7788)
	c.VarPath = f.Section("main").Key("var").MustString("./var")
	c.LogPath = f.Section("main").Key("log").MustString("./log")
}

func (c *Config) ToString() string {
	if bs, err := json.Marshal(c); err == nil {
		return string(bs)
	} else {
		return err.Error()
	}
}
