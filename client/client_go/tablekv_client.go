package client_go

import (
	"github.com/seefan/gopool"
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
	"github.com/seefan/tablekv/protocol/thrift_protocol"
	"context"
)

type TablePool struct {
	pool      *gopool.Pool
	host      string
	port      int
	tableName string
}

func NewTablePool(name, host string, port int) *TablePool {
	return &TablePool{
		pool:      gopool.NewPool(),
		host:      host,
		port:      port,
		tableName: name,
	}
}
func (t *TablePool) Start() error {
	t.pool.NewClient = func() gopool.IClient {
		return &TableKVClient{
			host: t.host,
			port: t.port,
			name: t.tableName,
		}
	}
	return t.Start()
}
func (t *TablePool) Close() {
	t.pool.Close()
}

type TableKVClient struct {
	trans  thrift.TTransport
	client *thrift_protocol.TableKVClient
	host   string
	port   int
	name   string
}

//打开连接
//
// 返回，error。如果连接到服务器时出错，就返回错误信息，否则返回nil
func (t *TableKVClient) Start() (err error) {
	trans, err := thrift.NewTSocket(fmt.Sprintf("%s:%d", t.host, t.port))
	if err != nil {
		return fmt.Errorf("error resolving address %s:%d %v", t.host, t.port, err)
	}
	t.trans = thrift.NewTFramedTransport(trans)
	err = t.trans.Open()
	if err != nil {
		return
	}
	bpt := thrift.NewTBinaryProtocolTransport(t.trans)
	mp := thrift.NewTMultiplexedProtocol(bpt, t.name)
	t.client = thrift_protocol.NewTableKVClient(thrift.NewTStandardClient(mp, mp))
	return
}

//关闭连接
//
// 返回，error。如果关闭连接时出错，就返回错误信息，否则返回nil
func (t *TableKVClient) Close() error {
	if t.IsOpen() {
		return t.trans.Close()
	}
	return nil
}

//是否打开
//
// 返回，bool。如果已连接到服务器，就返回true。
func (t *TableKVClient) IsOpen() bool {
	return t.trans != nil && t.trans.IsOpen()
}

//检查连接状态
//
// 返回，bool。如果无法访问服务器，就返回false。
func (t *TableKVClient) Ping() bool {
	if t.IsOpen() {
		if _, err := t.client.Ping(context.Background()); err == nil {
			return true
		}
	}
	return false
}
