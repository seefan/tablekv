package client_go

import (
	"github.com/seefan/gopool"
	"fmt"
	"net"
	"git.apache.org/thrift.git/lib/go/thrift"
	"github.com/seefan/tablekv/protocol/thrift_protocol"
	"context"
)

var (
	pool = gopool.NewPool()
	Host = "127.0.0.1"
	Port = "12321"
)

func init() {
	pool.NewClient = func() gopool.IClient {
		return &TableKVClient{}
	}
}
func Start() error {
	return pool.Start()
}
func Close() {
	if pool != nil {
		pool.Close()
	}
}

type TableKVClient struct {
	trans           thrift.TTransport
	protocolFactory thrift.TProtocolFactory
	client          *thrift_protocol.TableKVClient
}

//打开连接
//
// 返回，error。如果连接到服务器时出错，就返回错误信息，否则返回nil
func (t *TableKVClient) Start() (err error) {
	trans, err := thrift.NewTSocket(net.JoinHostPort(Host, Port))
	if err != nil {
		return fmt.Errorf("error resolving address:", err)
	}
	t.trans = thrift.NewTFramedTransport(trans)
	err = t.trans.Open()
	if err != nil {
		return
	}
	t.protocolFactory = thrift.NewTBinaryProtocolFactoryDefault()
	iprot := t.protocolFactory.GetProtocol(t.trans)
	oprot := t.protocolFactory.GetProtocol(t.trans)
	t.client = thrift_protocol.NewTableKVClient(thrift.NewTStandardClient(iprot, oprot))
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
