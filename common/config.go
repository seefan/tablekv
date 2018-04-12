package common

import (
	"github.com/go-ini/ini"
	"encoding/json"
)

//config
type Config struct {
	//ip
	Host string
	Port int
	//data path
	VarPath string
	//log path
	LogPath string
	//is master db
	IsMaster bool
	//The expiration time of the table in days
	Timeout int
	// The expiration type
	//0 按小时过期 1 按天过期
	TimeoutType int
	//Write buffer in Mbs,default is 256mb
	WriteBuffer int
}

//load config and set default value
func (c *Config) Load(f *ini.File) {
	if f == nil {
		f = ini.Empty()
	}
	c.Host = f.Section("main").Key("host").MustString("127.0.0.1")
	c.Port = f.Section("main").Key("port").MustInt(12321)
	c.VarPath = f.Section("main").Key("var").MustString("./var")
	c.LogPath = f.Section("main").Key("log").MustString("./log")
	c.Timeout = f.Section("main").Key("timeout").MustInt(1)
	c.TimeoutType = f.Section("main").Key("timeout_type").MustInt(1)
	c.WriteBuffer = f.Section("main").Key("write_buffer").MustInt(256)
}

//config value to string
func (c *Config) ToString() string {
	if bs, err := json.Marshal(c); err == nil {
		return string(bs)
	} else {
		return err.Error()
	}
}
