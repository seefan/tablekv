package tables

import (
	"time"
	"github.com/seefan/tablekv/common"
	log "github.com/cihub/seelog"
	"bytes"
)

//Regularly scan the table to remove expired tables

func (t *TableManager) timeProcessor() {
	for range t.timer.C {
		var ts []string
		for k, tab := range t.tableMap {
			if time.Since(tab.CreateTime).Hours() > common.Timeout {
				ts = append(ts, k)
			}
		}
		if len(ts) > 0 {
			for _, k := range ts {
				if err := t.DeleteTable(k); err != nil {
					log.Error("delete table has error", err)
				}
			}
		}
		log.Debugf("online table count is %d,running time is %s.", len(t.tableMap), time.Since(t.now).String())
		for name, tb := range t.tableMap {

			var bs bytes.Buffer
			for k, v := range tb.Info() {
				bs.WriteString(k)
				bs.WriteString(":\t\t\t")
				bs.WriteString(v)
				bs.WriteRune('\n')
			}
			log.Debugf("-----------------------------\n%s info %s", name,bs.String())
		}
	}
}
