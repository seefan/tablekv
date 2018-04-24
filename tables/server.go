package tables

import (
	"net/http"
	"fmt"
	"bufio"
	log "github.com/cihub/seelog"
	"strconv"
	"github.com/syndtr/goleveldb/leveldb/util"
)

const BadCommand = "Incorrect command format"

type TableServer struct {
	tableName  string
	tablePos   []byte
	tableStart []byte
	tableEnd   []byte
	tm         *TableManager
}

func (t *TableServer) Start(host string, port int) error {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseForm(); err == nil {
			if cmd, ok := r.Form["cmd"]; ok {
				log.Debug("cmd is ", cmd)
				if len(cmd) == 0 {
					w.Write([]byte(BadCommand))
					return
				}
				bw := bufio.NewWriter(w)
				switch cmd[0] {
				case "show":
					t.runShow(cmd, bw)
				case "from":
					t.runFrom(cmd, bw)
				case "next":
					t.runNext(cmd, bw)
				case "get":
					t.runGet(cmd, bw)
				default:
					w.Write([]byte(BadCommand))
				}
				bw.Flush()
			} else {
				w.Write([]byte("cmd not found"))
			}
		} else {
			w.Write([]byte(err.Error()))
		}
	})
	log.Infof("TableKV http manager is starting on %s:%d", host, port)
	return http.ListenAndServe(fmt.Sprintf("%s:%d", host, port), nil)
}
func (t *TableServer) runFrom(cmd []string, w *bufio.Writer) {
	if len(cmd) < 2 {
		badCommand(w)
	} else {
		if tb, ok := t.tm.tableMap[cmd[1]]; ok {
			ks, vs, err := tb.Scan(nil, nil, 10)
			if err != nil {
				w.WriteString(err.Error())
				return
			}
			t.tableName = cmd[1]

			t.tableStart = nil
			t.tableEnd = nil
			for i, k := range ks {
				w.Write(k)
				w.WriteRune('：')
				w.Write(vs[i])
				w.WriteRune('\n')
				t.tablePos = k
			}
		} else {
			w.WriteString("table not found")
		}
	}
}
func (t *TableServer) runNext(cmd []string, w *bufio.Writer) {
	if len(cmd) != 2 {
		badCommand(w)
	} else {
		if size, err := strconv.Atoi(cmd[1]); err != nil {
			w.Write([]byte(err.Error()))
		} else {
			if tb, ok := t.tm.tableMap[t.tableName]; ok {
				r := tb.db.NewIterator(&util.Range{
					Start: t.tableStart,
					Limit: t.tableEnd,
				}, nil)
				pos := 0
				r.Seek(t.tablePos)
				for r.Next() {
					w.Write(r.Key())
					w.WriteRune('：')
					w.Write(r.Value())
					w.WriteRune('\n')
					pos += 1
					t.tablePos = r.Key()
					if pos >= size {
						break
					}
				}
				r.Release()
				if err := r.Error(); err != nil {
					w.Write([]byte(err.Error()))
				}
			} else {
				w.WriteString("table not found")
			}
		}
	}
}
func (t *TableServer) runShow(cmd []string, w *bufio.Writer) {
	if len(cmd) != 2 {
		badCommand(w)
	} else {
		switch cmd[1] {
		case "tables":
			t.runShowTables(w)
		default:
			badCommand(w)
		}
	}
}
func badCommand(w *bufio.Writer) {
	w.WriteString(BadCommand)
}
func (t *TableServer) runShowTables(w *bufio.Writer) {
	for k := range t.tm.tableMap {
		w.WriteString(k)
		w.WriteRune('\n')
	}
}
func (t *TableServer) runGet(cmd []string, w *bufio.Writer) {
	if len(cmd) != 3 {
		badCommand(w)
	} else {
		if tb, err := t.tm.GetTable(cmd[1]); err != nil {
			w.WriteString(err.Error())
		} else {
			if v, err := tb.Get([]byte(cmd[2])); err != nil {
				w.WriteString(err.Error())
			} else {
				w.Write(v)
			}
		}
	}
}
