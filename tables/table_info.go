package tables

import (
	"time"
	"encoding/json"
)

type TableInfo struct {
	Name       string
	CreateTime time.Time
	Host       string
	//only for queue,max id
	QueueId int64
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
		t.CreateTime = tmp.CreateTime
		return nil
	}
}
