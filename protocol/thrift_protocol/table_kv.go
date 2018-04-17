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
	return t.p.Set(key, value)
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
	return t.p.BatchSet(keys, values)
}

func (t *ThriftProcessor) ScanMap(ctx context.Context, key_start []byte, key_end []byte, limit int32) (r map[string][]byte, err error) {
	if ks, vs, err := t.p.Scan(key_start, key_end, int(limit)); err == nil {
		r = make(map[string][]byte)
		for i, k := range ks {
			r[string(k)] = []byte(vs[i])
		}
	}
	return
}

// Parameters:
//  - KeyStart
//  - KeyEnd
//  - Limit
func (t *ThriftProcessor)Scan(ctx context.Context,key_start []byte, key_end []byte, limit int32) (r [][][]byte, err error) {
	if ks, vs, err := t.p.Scan(key_start, key_end, int(limit)); err == nil {
		for i, k := range ks {
			r = append(r, [][]byte{k, vs[i]})
		}
	}
	return
}

func (t *ThriftProcessor) QGet(ctx context.Context) (r []byte, err error) {
	return t.p.QGet()
}

// Parameters:
//  - Value
func (t *ThriftProcessor) QSet(ctx context.Context, value []byte) (err error) {
	return t.p.QSet(value)
}

// Parameters:
//  - Value
func (t *ThriftProcessor) BatchQSet(ctx context.Context, value [][]byte) (err error) {
	return t.p.BatchQSet(value)
}

// Parameters:
//  - Size
func (t *ThriftProcessor) BatchQGet(ctx context.Context, size int32) (r [][]byte, err error) {
	return t.p.BatchQGet(int(size))
}
