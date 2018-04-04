package tables

import (
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/seefan/goerr"
	"path"
	"time"
	//"github.com/syndtr/goleveldb/leveldb/util"
	log "github.com/cihub/seelog"
	"github.com/syndtr/goleveldb/leveldb/util"
	"github.com/syndtr/goleveldb/leveldb/opt"
)

type Table struct {
	db         *leveldb.DB
	name       string
	isOpen     bool
	lastUpdate time.Time
}
type TableValue struct {
	Key   string
	Value string
}

func LoadTable(p, name string) (t *Table, err error) {
	t = &Table{
		isOpen: false,
	}
	if t.db, err = leveldb.OpenFile(path.Join(p, name), &opt.Options{
		WriteBuffer: 8 * opt.MiB, //8mb
	}); err == nil {
		t.isOpen = true
		t.name = name
	}
	t.lastUpdate = time.Now()
	return
}

func (t *Table) Close() error {
	log.Debugf("%s is close", t.name)
	if t.db != nil {
		t.isOpen = false
		return t.db.Close()
	}
	return nil
}
func (t *Table) Get(key []byte) ([]byte, error) {
	if !t.isOpen {
		return nil, goerr.New("db is not open")
	}
	return t.db.Get([]byte(key), nil)
}
func (t *Table) Set(key, value []byte, ) (error) {
	if !t.isOpen {
		return goerr.New("db is not open")
	}
	return t.db.Put(key, value, nil)
}
func (t *Table) Delete(key []byte) (error) {
	if !t.isOpen {
		return goerr.New("db is not open")
	}
	return t.db.Delete(key, nil)
}

func (t *Table) Exists(key []byte) (bool, error) {
	if !t.isOpen {
		return false, goerr.New("db is not open")
	}
	return t.db.Has(key, nil)
}

func (t *Table) Scan(start, end []byte) (re []TableValue, err error) {
	if !t.isOpen {
		return nil, goerr.New("db is not open")
	}

	iter := t.db.NewIterator(&util.Range{Start: start, Limit: end}, nil)
	for iter.Next() {
		re = append(re, TableValue{Key: string(iter.Key()), Value: string(iter.Value())})
	}
	iter.Release()
	err = iter.Error()
	return
}

func (t *Table) BatchSet(keys [][]byte, values [][]byte) error {
	bat := new(leveldb.Batch)
	if len(keys) != len(values) {
		return goerr.New("The length of keys and values is different.")
	}
	for i, k := range keys {
		bat.Put(k, values[i])
	}
	return t.db.Write(bat, nil)
}
