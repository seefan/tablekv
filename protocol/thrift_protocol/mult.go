package thrift_protocol

import (
	"github.com/seefan/goerr"
	"git.apache.org/thrift.git/lib/go/thrift"
	"context"
	"fmt"
	"strings"

	"github.com/seefan/tablekv/common"
)

const (
	INVALID_TMESSAGE_TYPE  = 0
	CALL                   = 1
	REPLY                  = 2
	EXCEPTION              = 3
	ONEWAY                 = 4
	MULTIPLEXED_SEPARATOR=":"
)

type MultiplexedProcessor struct {
	thrift.TMultiplexedProcessor
	serviceProcessorMap map[string]thrift.TProcessor
	processorManager common.GetProcessor
}

func (t *MultiplexedProcessor) Process(ctx context.Context, in, out thrift.TProtocol) (bool, thrift.TException) {
	name, typeId, seqid, err := in.ReadMessageBegin()
	if err != nil {
		return false, err
	}
	if typeId != CALL && typeId != ONEWAY {
		return false, fmt.Errorf("Unexpected message type %v", typeId)
	}
	//extract the service name
	v := strings.SplitN(name, MULTIPLEXED_SEPARATOR, 2)
	if len(v) != 2 {
		return false, fmt.Errorf("Table name not found in message name: %s.  Did you forget to use a TMultiplexProtocol in your client?", name)
	}
	actualProcessor, ok := t.serviceProcessorMap[v[0]]
	if !ok {
		tp,err:=t.processorManager.GetProcessor(v[0])
		if err!=nil{
			return false,goerr.New("Table name not found")
		}
		actualProcessor = NewTableKVProcessor(&ThriftProcessor{
			p:tp,
		})
		t.serviceProcessorMap[v[0]]=actualProcessor
	}

	smb := thrift.NewStoredMessageProtocol(in, v[1], typeId, seqid)
	return actualProcessor.Process(ctx,smb, out)
}
