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
	//connect timeout,0 is none
	Timeout int
	//The num of expiration time
	ExpiredNum int
	// The expiration type
	//0 none 1 hour 2 day
	ExpiredType int
	//Write buffer in Mbs,default is 64mb
	WriteBuffer int
	//Block buffer in Mbs,default is 64mb
	BlockBuffer int
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
	c.Timeout = f.Section("main").Key("timeout").MustInt(0)
	c.ExpiredNum = f.Section("main").Key("expired_num").MustInt(0)
	c.ExpiredType = f.Section("main").Key("expired_type").MustInt(0)
	c.WriteBuffer = f.Section("main").Key("write_buffer").MustInt(64)
	c.BlockBuffer = f.Section("main").Key("block_buffer").MustInt(64)
}

//config value to string
func (c *Config) ToString() string {
	if bs, err := json.Marshal(c); err == nil {
		return string(bs)
	} else {
		return err.Error()
	}
}
