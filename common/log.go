package common

import (
	"fmt"
	log "github.com/cihub/seelog"
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

var (
	logText = `<seelog type="asynctimer" asyncinterval="5000000" minlevel="debug">
	<outputs formatid="main">
		<console/>
		<rollingfile type="size" filename="__log_url__" maxsize="1024000" maxrolls="10" />
	</outputs>
	<formats>
		<format id="main" format="%Date(2006-01-02 15:04:05) [%Level] %RelFile line:%Line %Msg%n"/>
	</formats>
</seelog>`
)

func InitLog(logConfig, logFile string) {
	if _, err := os.Stat(logConfig); os.IsNotExist(err) {
		ioutil.WriteFile(logConfig, []byte(strings.Replace(logText, "__log_url__", logFile, 1)), 0764)
	}
	if logger, err := log.LoggerFromConfigAsFile(logConfig); err == nil {
		log.ReplaceLogger(logger)
	}
}
func PrintErr() {
	if err := recover(); err != nil {
		path, fe := filepath.Abs(os.Args[0])
		if fe != nil {
			path = os.Args[0]
		}
		path = filepath.Dir(path)
		path += string(os.PathSeparator) + "fault.txt"
		str := fmt.Sprintf("%v\n", err)
		for i := 1; i < 10; i++ {
			funcName, file, line, ok := runtime.Caller(i)
			if ok {
				str += fmt.Sprintf("frame %v:[func:%v,file:%v,line:%v]\n", i, runtime.FuncForPC(funcName).Name(), file, line)
			}
		}
		logFile, err := os.OpenFile(path, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0660)
		if err != nil {
			println(err.Error())
			println(str)
			return
		}
		defer logFile.Close()
		println(str)
		logFile.WriteString(str)
	}
}
