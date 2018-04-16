package tables

import (
	"time"
	"encoding/json"
	"sync"
)

type TableInfo struct {
	Name       string
	CreateTime time.Time
	Host       string
	//only for queue,max id
	QueueId int64
	lock    sync.Mutex
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
		t.QueueId = tmp.QueueId
		return nil
	}
}
func (t *TableInfo) GetQueueId() int64 {
	t.lock.Lock()
	defer t.lock.Unlock()
	t.QueueId += 1
	return t.QueueId
}
