package thrift_protocol

import (
	"github.com/seefan/goerr"
	"git.apache.org/thrift.git/lib/go/thrift"
	"context"
	"fmt"
	"strings"

	"github.com/seefan/tablekv/common"
	"sync"
)



type MultiplexedProcessor struct {
	thrift.TMultiplexedProcessor
	serviceProcessorMap map[string]thrift.TProcessor
	processorManager    common.GetProcessor
	lock                sync.Mutex
}

func (t *MultiplexedProcessor) Process(ctx context.Context, in, out thrift.TProtocol) (bool, thrift.TException) {
	name, typeId, seqid, err := in.ReadMessageBegin()
	if err != nil {
		return false, err
	}
	if typeId != thrift.CALL && typeId != thrift.ONEWAY {
		return false, fmt.Errorf("Unexpected message type %v", typeId)
	}
	//extract the service name
	var tableName, methodName string
	v := strings.SplitN(name, thrift.MULTIPLEXED_SEPARATOR, 2)
	if len(v) == 2 {
		tableName = v[0]
		methodName = v[1]
	} else {
		tableName = "default"
		methodName = name
	}
	actualProcessor, ok := t.serviceProcessorMap[tableName]
	if !ok {
		t.lock.Lock()
		tp, err := t.processorManager.GetProcessor(tableName)
		if err != nil {
			return false, goerr.New("Table name not found")
		}
		actualProcessor = NewTableKVProcessor(&ThriftProcessor{
			p: tp,
		})
		t.serviceProcessorMap[tableName] = actualProcessor
		t.lock.Unlock()
	}

	smb := thrift.NewStoredMessageProtocol(in, methodName, typeId, seqid)
	return actualProcessor.Process(ctx, smb, out)
}
