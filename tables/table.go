package tables

import (
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/seefan/goerr"
	"path"
)

type Table struct {
	db     *leveldb.DB
	name   string
	isOpen bool
}

func loadTable(p, name string) (t *Table, err error) {
	t = &Table{
		isOpen: false,
	}
	if t.db, err = leveldb.OpenFile(path.Join(p, name), nil); err == nil {
		t.isOpen = true
		t.name = name
	}
	return
}

func (t *Table) Close() error {
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
