package tables

import (
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/seefan/goerr"
	"path"
	"time"
	log "github.com/cihub/seelog"
	"github.com/syndtr/goleveldb/leveldb/util"
	"github.com/syndtr/goleveldb/leveldb/opt"
	"github.com/seefan/tablekv/common"
	"github.com/syndtr/goleveldb/leveldb/errors"
)

type Table struct {
	TableInfo
	db     *leveldb.DB
	isOpen bool
	//last update time ,save memory
	lastTime time.Time
	//file save path
	path string
}

//load a new TableKV
func LoadTable(p string, info TableInfo) (t *Table, err error) {
	t = &Table{
		isOpen:    false,
		path:      path.Join(p, common.HashPath(info.Name)),
		TableInfo: info,
	}
	if t.db, err = leveldb.OpenFile(t.path, &opt.Options{
		WriteBuffer:            common.WriteBuffer * opt.MiB, //write buffer is important.
		BlockCacheCapacity:     common.BlockBuffer * opt.MiB,
		BlockSize:              2 * opt.DefaultBlockSize,
		CompactionTableSize:    16 * opt.DefaultCompactionTableSize,
		CompactionTotalSize:    16 * opt.DefaultCompactionTotalSize,
		WriteL0PauseTrigger:    64,
		WriteL0SlowdownTrigger: 32,
	}); err == nil {
		t.isOpen = true
		t.TableInfo = info
	} else {
		log.Error("load table error", err)
	}

	if t.CreateTime.Sub(time.Time{}).Seconds() < 1 {
		t.CreateTime = time.Now()
	}
	return
}

//close TableKV
func (t *Table) Close() error {
	log.Debugf("%s is close", t.Name)
	if t.db != nil {
		t.isOpen = false
		return t.db.Close()
	}
	return nil
}

//get value from key
func (t *Table) Get(key []byte) ([]byte, error) {
	if !t.isOpen {
		return nil, goerr.New("db is not open")
	}
	bs, err := t.db.Get(key, nil)
	if err != nil && err == errors.ErrNotFound {
		err = nil
	}
	return bs, err
}

//set key and value
func (t *Table) Set(key, value []byte, ) (error) {
	if !t.isOpen {
		return goerr.New("db is not open")
	}
	return t.db.Put(key, value, nil)
}

//delete key and value
func (t *Table) Delete(key []byte) (error) {
	if !t.isOpen {
		return goerr.New("db is not open")
	}
	return t.db.Delete(key, nil)
}

//check key exists or not
func (t *Table) Exists(key []byte) (bool, error) {
	if !t.isOpen {
		return false, goerr.New("db is not open")
	}
	return t.db.Has(key, nil)
}

//scan keys from start to end
func (t *Table) Scan(start, end []byte, limit int) (key [][]byte, value [][]byte, err error) {
	if !t.isOpen {
		return nil, nil, goerr.New("db is not open")
	}
	index := 0
	its := t.db.NewIterator(&util.Range{Start: start, Limit: end}, nil)
	for its.Next() {
		index += 1
		if limit > 0 && index > limit {
			break
		}
		var k []byte
		k = append(k, its.Key()...)
		var v []byte
		v = append(v, its.Value()...)
		key = append(key, k)
		value = append(value, v)
	}
	its.Release()
	err = its.Error()

	return
}

//batch set key and value
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

//batch delete key
func (t *Table) BatchDelete(keys [][]byte) error {
	bat := new(leveldb.Batch)
	for _, k := range keys {
		bat.Delete(k)
	}
	return t.db.Write(bat, nil)
}

//get table info
//leveldb.num-files-at-level{n}
//Returns the number of files at level 'n'.
//leveldb.stats
//Returns statistics of the underlying DB.
//leveldb.iostats
//Returns statistics of effective disk read and write.
//leveldb.writedelay
//Returns cumulative write delay caused by compaction.
//leveldb.sstables
//Returns sstables list for each level.
//leveldb.blockpool
//Returns block pool stats.
//leveldb.cachedblock
//Returns size of cached block.
//leveldb.openedtables
//Returns number of opened tables.
//leveldb.alivesnaps
//Returns number of alive snapshots.
//leveldb.aliveiters
//Returns number of alive iterators.
func (t *Table) Info() map[string]string {
	re := make(map[string]string)
	if t.isOpen {
		re["leveldb.stats"], _ = t.db.GetProperty("leveldb.stats")
		re["leveldb.iostats"], _ = t.db.GetProperty("leveldb.iostats")
		//re["leveldb.writedelay"], _ = t.db.GetProperty("leveldb.writedelay")
		//re["leveldb.sstables"], _ = t.db.GetProperty("leveldb.sstables")
		//re["leveldb.blockpool"], _ = t.db.GetProperty("leveldb.blockpool")
		re["leveldb.cachedblock"], _ = t.db.GetProperty("leveldb.cachedblock")
		//re["leveldb.openedtables"], _ = t.db.GetProperty("leveldb.openedtables")
		//re["leveldb.alivesnaps"], _ = t.db.GetProperty("leveldb.alivesnaps")
		//re["leveldb.aliveiters"], _ = t.db.GetProperty("leveldb.aliveiters")
	} else {
		re["leveldb.stats"] = "not open"
	}
	return re
}
