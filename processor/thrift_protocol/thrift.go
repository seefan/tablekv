package thrift_protocol

import (
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"

	"github.com/seefan/tablekv/tables"
	"time"
)

type Thrift struct {
	server *thrift.TSimpleServer
}

func (t *Thrift) Start(tm *tables.TableManager, host string, port int) error {
	socket, err := thrift.NewTServerSocketTimeout(fmt.Sprintf("%s:%d", host, port), time.Duration(30)*time.Second)
	if err != nil {
		return err
	}
	processor := NewTableKVProcessor(&ThriftProcessor{
		tm: tm,
	})
	transportFactory := thrift.NewTFramedTransportFactory(thrift.NewTTransportFactory())
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()
	t.server = thrift.NewTSimpleServer4(processor, socket, transportFactory, protocolFactory)
	return t.server.Serve()
}

type ThriftProcessor struct {
	tm *tables.TableManager
}

// Parameters:
//  - Table
//  - Key
func (t *ThriftProcessor) Get(table []byte, key []byte) (r []byte, err error) {
	if tb, err := t.tm.GetTable(string(table)); err == nil {
		return tb.Get(key)
	} else {
		return nil, err
	}
}

// Parameters:
//  - Table
//  - Key
//  - Value
func (t *ThriftProcessor) Set(table []byte, key []byte, value []byte) (err error) {
	if tb, err := t.tm.GetTable(string(table)); err == nil {
		return tb.Set(key, value)
	} else {
		return err
	}
}

// Parameters:
//  - Table
//  - Key
func (t *ThriftProcessor) Exists(table []byte, key []byte) (r bool, err error) {
	if tb, err := t.tm.GetTable(string(table)); err == nil {
		return tb.Exists(key)
	} else {
		return false, err
	}
}

// Parameters:
//  - Table
//  - Key
func (t *ThriftProcessor) Delete(table []byte, key []byte) (err error) {
	if tb, err := t.tm.GetTable(string(table)); err == nil {
		return tb.Delete(key)
	} else {
		return err
	}
}
