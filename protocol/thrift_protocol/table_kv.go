package thrift_protocol

import (
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
	"context"
	"github.com/seefan/tablekv/common"
)

type Thrift struct {
	server *thrift.TSimpleServer
}

func (t *Thrift) Start(pm common.GetProcessor, host string, port int) error {
	socket, err := thrift.NewTServerSocket(fmt.Sprintf("%s:%d", host, port))
	if err != nil {
		return err
	}
	p1 := &MultiplexedProcessor{
		processorManager:    pm,
		serviceProcessorMap: make(map[string]thrift.TProcessor),
	}

	transportFactory := thrift.NewTFramedTransportFactory(thrift.NewTTransportFactory())
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()
	t.server = thrift.NewTSimpleServer4(p1, socket, transportFactory, protocolFactory)
	return t.server.Serve()
}
func (t *Thrift) Stop() error {
	if t.server != nil {
		return t.server.Stop()
	} else {
		return nil
	}
}

type ThriftProcessor struct {
	p common.Processor
}

// Parameters:
//  - Table
//  - Key
func (t *ThriftProcessor) Get(ctx context.Context, key []byte) (r []byte, err error) {
	return t.p.Get(key)
}

// Parameters:
//  - Table
//  - Key
//  - Value
func (t *ThriftProcessor) Set(ctx context.Context, key []byte, value []byte) (err error) {
	return t.p.Set(key,value)
}

// Parameters:
//  - Tablee
//  - Key
func (t *ThriftProcessor) Exists(ctx context.Context, key []byte) (r bool, err error) {
	return t.p.Exists(key)
}

// Parameters:
//  - Table
//  - Key
func (t *ThriftProcessor) Delete(ctx context.Context, key []byte) (err error) {
	return t.p.Delete(key)
}

// Parameters:
//  - Keys
//  - Values
func (t *ThriftProcessor) BatchSet(ctx context.Context, keys [][]byte, values [][]byte) (err error) {
	return t.p.BatchSet(keys,values)
}
