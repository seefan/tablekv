package tables

import (
	"time"
)
//Regularly scan the table to remove expired tables

func (t *TableManager) timeProcessor() {
	for range t.timer.C {
		var ts []string
		for k, tab := range t.tableMap {
			if time.Since(tab.lastUpdate).Minutes() > t.timeout {
				ts = append(ts, k)
			}
		}
		if len(ts) > 0 {
			for _, k := range ts {
				t.DeleteTable(k)
			}
		}
	}
}

