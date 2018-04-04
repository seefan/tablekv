package process

import (
	"github.com/seefan/tablekv/tables"
	"github.com/syndtr/goleveldb/leveldb"
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

// Parameters:
//  - Keys
//  - Values
func (b *LevelDbProcess) BatchSet(keys [][]byte, values [][]byte) (err error) {
	return b.table.BatchSet(keys,values)
}
