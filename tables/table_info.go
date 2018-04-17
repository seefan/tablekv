package tables

import (
	"time"
	"sync"
	"fmt"
	"strings"
	"strconv"
	"github.com/seefan/goerr"
)

const (
	MinTimeFormat = "20060102150405"
)

type TableInfo struct {
	Name       string
	CreateTime time.Time
	Host       string
	//only for queue,max id
	queueId uint32
	lock    sync.Mutex
}

func (t *TableInfo) ToByte() []byte {
	return []byte(fmt.Sprintf("%s:%s:%s:%x", t.Name, t.Host, t.CreateTime.Format(MinTimeFormat), t.queueId))
}
func (t *TableInfo) FromByte(bs []byte) error {
	ss := strings.Split(string(bs), ":")
	if len(ss) == 4 {
		t.Name = ss[0]
		t.Host = ss[1]
		if tt, err := time.Parse(MinTimeFormat, ss[2]); err != nil {
			return err
		} else {
			t.CreateTime = tt
		}
		if tt, err := strconv.ParseInt(ss[3], 16, 10); err != nil {
			return err
		} else {
			t.queueId = uint32(tt)
		}
		return nil
	} else {
		return goerr.New("error cdb data:%s", bs)
	}
}
func (t *TableInfo) GetQueueId() uint32 {
	t.lock.Lock()
	defer t.lock.Unlock()
	t.queueId += 1
	return t.queueId
}
