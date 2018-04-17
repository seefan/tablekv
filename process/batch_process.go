package process

import (
	"github.com/seefan/tablekv/tables"
	"github.com/syndtr/goleveldb/leveldb"
	"strconv"
	"github.com/seefan/goerr"
)

type LevelDbProcess struct {
	name  string
	table *tables.Table
	bat   *leveldb.Batch
}

func NewLevelDbProcess(name string, tb *tables.Table) *LevelDbProcess {
	return &LevelDbProcess{
		name:  name,
		table: tb,
	}
}

// Parameters:
//  - Key
func (b *LevelDbProcess) Get(key []byte) ([]byte, error) {
	return b.table.Get(key)
}

// Parameters:
//  - Key
//  - Value
func (b *LevelDbProcess) Set(key []byte, value []byte) (err error) {
	return b.table.Set(key, value)
}

// Parameters:
//  - Key
func (b *LevelDbProcess) Exists(key []byte) (r bool, err error) {
	return b.table.Exists(key)
}

// Parameters:
//  - Key
func (b *LevelDbProcess) Delete(key []byte) (err error) {
	return b.table.Delete(key)
}
func (b *LevelDbProcess) Scan(key_start []byte, key_end []byte, limit int) (key [][]byte, value [][]byte, err error) {
	return b.table.Scan(key_start, key_end, int(limit))
}


// Parameters:
//  - Keys
//  - Values
func (b *LevelDbProcess) BatchSet(keys [][]byte, values [][]byte) (err error) {
	return b.table.BatchSet(keys, values)
}
func (b *LevelDbProcess) QGet() (r []byte, err error) {
	if ks, vs, err := b.table.Scan(nil, nil, 1); err != nil {
		return nil, err
	} else {
		if len(ks) > 0 {
			for i, k := range ks {
				if err := b.table.Delete([]byte(k)); err != nil {
					return []byte(vs[i]), err
				}
				return []byte(vs[i]), nil
			}
		} else {
			return nil, goerr.New("empty")
		}
	}
	return nil, nil //never reach
}

// Parameters:
//  - Value
func (b *LevelDbProcess) QSet(value []byte) (err error) {
	id := b.table.GetQueueId()
	key := []byte("q:00000000")
	k := strconv.AppendInt(nil, int64(id), 16)
	return b.table.Set(append(key[:10-len(k)], k...), value)
}

// Parameters:
//  - Value
func (b *LevelDbProcess) BatchQSet(value [][]byte) (err error) {
	var ks [][]byte
	for range value {
		id := b.table.GetQueueId()
		ks = append(ks, strconv.AppendInt(nil, int64(id), 10))
	}
	return b.table.BatchSet(ks, value)
}

// Parameters:
//  - Size
func (b *LevelDbProcess) BatchQGet(size int) (r [][]byte, err error) {
	if ks, vs, err := b.table.Scan(nil, nil, int(size)); err != nil {
		return nil, err
	} else {
		if len(ks) > 0 {
			for i := range ks {
				r = append(r, []byte(vs[i]))
			}
			return r, nil
		} else {
			return nil, goerr.New("empty")
		}
	}
}
