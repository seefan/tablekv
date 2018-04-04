package tables

import (
	"time"
	"encoding/json"
)

type TableInfo struct {
	Name       string
	LastUpdate time.Time
	Host       string
}

func (t *TableInfo) ToByte() []byte {
	if bs, err := json.Marshal(t); err == nil {
		return bs
	} else {
		return nil
	}
}
func (t *TableInfo) FromByte(bs []byte) error {
	tmp := new(TableInfo)
	if err := json.Unmarshal(bs, tmp); err != nil {
		return err
	} else {
		t.Host = tmp.Host
		t.Name = tmp.Name
		t.LastUpdate = tmp.LastUpdate
		return nil
	}
}
