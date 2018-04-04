// Autogenerated by Thrift Compiler (0.11.0)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

package thrift_protocol

import (
	"bytes"
	"reflect"
	"context"
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
)

// (needed to ensure safety because of naive import list construction.)
var _ = thrift.ZERO
var _ = fmt.Printf
var _ = context.Background
var _ = reflect.DeepEqual
var _ = bytes.Equal

type TableKV interface {
  // Parameters:
  //  - Key
  Get(ctx context.Context, key []byte) (r []byte, err error)
  // Parameters:
  //  - Key
  //  - Value
  Set(ctx context.Context, key []byte, value []byte) (err error)
  // Parameters:
  //  - Key
  Exists(ctx context.Context, key []byte) (r bool, err error)
  // Parameters:
  //  - Key
  Delete(ctx context.Context, key []byte) (err error)
  // Parameters:
  //  - Keys
  //  - Values
  BatchSet(ctx context.Context, keys [][]byte, values [][]byte) (err error)
}

type TableKVClient struct {
  c thrift.TClient
}

// Deprecated: Use NewTableKV instead
func NewTableKVClientFactory(t thrift.TTransport, f thrift.TProtocolFactory) *TableKVClient {
  return &TableKVClient{
    c: thrift.NewTStandardClient(f.GetProtocol(t), f.GetProtocol(t)),
  }
}

// Deprecated: Use NewTableKV instead
func NewTableKVClientProtocol(t thrift.TTransport, iprot thrift.TProtocol, oprot thrift.TProtocol) *TableKVClient {
  return &TableKVClient{
    c: thrift.NewTStandardClient(iprot, oprot),
  }
}

func NewTableKVClient(c thrift.TClient) *TableKVClient {
  return &TableKVClient{
    c: c,
  }
}

// Parameters:
//  - Key
func (p *TableKVClient) Get(ctx context.Context, key []byte) (r []byte, err error) {
  var _args0 TableKVGetArgs
  _args0.Key = key
  var _result1 TableKVGetResult
  if err = p.c.Call(ctx, "Get", &_args0, &_result1); err != nil {
    return
  }
  return _result1.GetSuccess(), nil
}

// Parameters:
//  - Key
//  - Value
func (p *TableKVClient) Set(ctx context.Context, key []byte, value []byte) (err error) {
  var _args2 TableKVSetArgs
  _args2.Key = key
  _args2.Value = value
  var _result3 TableKVSetResult
  if err = p.c.Call(ctx, "Set", &_args2, &_result3); err != nil {
    return
  }
  return nil
}

// Parameters:
//  - Key
func (p *TableKVClient) Exists(ctx context.Context, key []byte) (r bool, err error) {
  var _args4 TableKVExistsArgs
  _args4.Key = key
  var _result5 TableKVExistsResult
  if err = p.c.Call(ctx, "Exists", &_args4, &_result5); err != nil {
    return
  }
  return _result5.GetSuccess(), nil
}

// Parameters:
//  - Key
func (p *TableKVClient) Delete(ctx context.Context, key []byte) (err error) {
  var _args6 TableKVDeleteArgs
  _args6.Key = key
  var _result7 TableKVDeleteResult
  if err = p.c.Call(ctx, "Delete", &_args6, &_result7); err != nil {
    return
  }
  return nil
}

// Parameters:
//  - Keys
//  - Values
func (p *TableKVClient) BatchSet(ctx context.Context, keys [][]byte, values [][]byte) (err error) {
  var _args8 TableKVBatchSetArgs
  _args8.Keys = keys
  _args8.Values = values
  var _result9 TableKVBatchSetResult
  if err = p.c.Call(ctx, "BatchSet", &_args8, &_result9); err != nil {
    return
  }
  return nil
}

type TableKVProcessor struct {
  processorMap map[string]thrift.TProcessorFunction
  handler TableKV
}

func (p *TableKVProcessor) AddToProcessorMap(key string, processor thrift.TProcessorFunction) {
  p.processorMap[key] = processor
}

func (p *TableKVProcessor) GetProcessorFunction(key string) (processor thrift.TProcessorFunction, ok bool) {
  processor, ok = p.processorMap[key]
  return processor, ok
}

func (p *TableKVProcessor) ProcessorMap() map[string]thrift.TProcessorFunction {
  return p.processorMap
}

func NewTableKVProcessor(handler TableKV) *TableKVProcessor {

  self10 := &TableKVProcessor{handler:handler, processorMap:make(map[string]thrift.TProcessorFunction)}
  self10.processorMap["Get"] = &tableKVProcessorGet{handler:handler}
  self10.processorMap["Set"] = &tableKVProcessorSet{handler:handler}
  self10.processorMap["Exists"] = &tableKVProcessorExists{handler:handler}
  self10.processorMap["Delete"] = &tableKVProcessorDelete{handler:handler}
  self10.processorMap["BatchSet"] = &tableKVProcessorBatchSet{handler:handler}
return self10
}

func (p *TableKVProcessor) Process(ctx context.Context, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
  name, _, seqId, err := iprot.ReadMessageBegin()
  if err != nil { return false, err }
  if processor, ok := p.GetProcessorFunction(name); ok {
    return processor.Process(ctx, seqId, iprot, oprot)
  }
  iprot.Skip(thrift.STRUCT)
  iprot.ReadMessageEnd()
  x11 := thrift.NewTApplicationException(thrift.UNKNOWN_METHOD, "Unknown function " + name)
  oprot.WriteMessageBegin(name, thrift.EXCEPTION, seqId)
  x11.Write(oprot)
  oprot.WriteMessageEnd()
  oprot.Flush()
  return false, x11

}

type tableKVProcessorGet struct {
  handler TableKV
}

func (p *tableKVProcessorGet) Process(ctx context.Context, seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
  args := TableKVGetArgs{}
  if err = args.Read(iprot); err != nil {
    iprot.ReadMessageEnd()
    x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
    oprot.WriteMessageBegin("Get", thrift.EXCEPTION, seqId)
    x.Write(oprot)
    oprot.WriteMessageEnd()
    oprot.Flush()
    return false, err
  }

  iprot.ReadMessageEnd()
  result := TableKVGetResult{}
var retval []byte
  var err2 error
  if retval, err2 = p.handler.Get(ctx, args.Key); err2 != nil {
    x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing Get: " + err2.Error())
    oprot.WriteMessageBegin("Get", thrift.EXCEPTION, seqId)
    x.Write(oprot)
    oprot.WriteMessageEnd()
    oprot.Flush()
    return true, err2
  } else {
    result.Success = retval
}
  if err2 = oprot.WriteMessageBegin("Get", thrift.REPLY, seqId); err2 != nil {
    err = err2
  }
  if err2 = result.Write(oprot); err == nil && err2 != nil {
    err = err2
  }
  if err2 = oprot.WriteMessageEnd(); err == nil && err2 != nil {
    err = err2
  }
  if err2 = oprot.Flush(); err == nil && err2 != nil {
    err = err2
  }
  if err != nil {
    return
  }
  return true, err
}

type tableKVProcessorSet struct {
  handler TableKV
}

func (p *tableKVProcessorSet) Process(ctx context.Context, seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
  args := TableKVSetArgs{}
  if err = args.Read(iprot); err != nil {
    iprot.ReadMessageEnd()
    x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
    oprot.WriteMessageBegin("Set", thrift.EXCEPTION, seqId)
    x.Write(oprot)
    oprot.WriteMessageEnd()
    oprot.Flush()
    return false, err
  }

  iprot.ReadMessageEnd()
  result := TableKVSetResult{}
  var err2 error
  if err2 = p.handler.Set(ctx, args.Key, args.Value); err2 != nil {
    x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing Set: " + err2.Error())
    oprot.WriteMessageBegin("Set", thrift.EXCEPTION, seqId)
    x.Write(oprot)
    oprot.WriteMessageEnd()
    oprot.Flush()
    return true, err2
  }
  if err2 = oprot.WriteMessageBegin("Set", thrift.REPLY, seqId); err2 != nil {
    err = err2
  }
  if err2 = result.Write(oprot); err == nil && err2 != nil {
    err = err2
  }
  if err2 = oprot.WriteMessageEnd(); err == nil && err2 != nil {
    err = err2
  }
  if err2 = oprot.Flush(); err == nil && err2 != nil {
    err = err2
  }
  if err != nil {
    return
  }
  return true, err
}

type tableKVProcessorExists struct {
  handler TableKV
}

func (p *tableKVProcessorExists) Process(ctx context.Context, seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
  args := TableKVExistsArgs{}
  if err = args.Read(iprot); err != nil {
    iprot.ReadMessageEnd()
    x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
    oprot.WriteMessageBegin("Exists", thrift.EXCEPTION, seqId)
    x.Write(oprot)
    oprot.WriteMessageEnd()
    oprot.Flush()
    return false, err
  }

  iprot.ReadMessageEnd()
  result := TableKVExistsResult{}
var retval bool
  var err2 error
  if retval, err2 = p.handler.Exists(ctx, args.Key); err2 != nil {
    x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing Exists: " + err2.Error())
    oprot.WriteMessageBegin("Exists", thrift.EXCEPTION, seqId)
    x.Write(oprot)
    oprot.WriteMessageEnd()
    oprot.Flush()
    return true, err2
  } else {
    result.Success = &retval
}
  if err2 = oprot.WriteMessageBegin("Exists", thrift.REPLY, seqId); err2 != nil {
    err = err2
  }
  if err2 = result.Write(oprot); err == nil && err2 != nil {
    err = err2
  }
  if err2 = oprot.WriteMessageEnd(); err == nil && err2 != nil {
    err = err2
  }
  if err2 = oprot.Flush(); err == nil && err2 != nil {
    err = err2
  }
  if err != nil {
    return
  }
  return true, err
}

type tableKVProcessorDelete struct {
  handler TableKV
}

func (p *tableKVProcessorDelete) Process(ctx context.Context, seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
  args := TableKVDeleteArgs{}
  if err = args.Read(iprot); err != nil {
    iprot.ReadMessageEnd()
    x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
    oprot.WriteMessageBegin("Delete", thrift.EXCEPTION, seqId)
    x.Write(oprot)
    oprot.WriteMessageEnd()
    oprot.Flush()
    return false, err
  }

  iprot.ReadMessageEnd()
  result := TableKVDeleteResult{}
  var err2 error
  if err2 = p.handler.Delete(ctx, args.Key); err2 != nil {
    x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing Delete: " + err2.Error())
    oprot.WriteMessageBegin("Delete", thrift.EXCEPTION, seqId)
    x.Write(oprot)
    oprot.WriteMessageEnd()
    oprot.Flush()
    return true, err2
  }
  if err2 = oprot.WriteMessageBegin("Delete", thrift.REPLY, seqId); err2 != nil {
    err = err2
  }
  if err2 = result.Write(oprot); err == nil && err2 != nil {
    err = err2
  }
  if err2 = oprot.WriteMessageEnd(); err == nil && err2 != nil {
    err = err2
  }
  if err2 = oprot.Flush(); err == nil && err2 != nil {
    err = err2
  }
  if err != nil {
    return
  }
  return true, err
}

type tableKVProcessorBatchSet struct {
  handler TableKV
}

func (p *tableKVProcessorBatchSet) Process(ctx context.Context, seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
  args := TableKVBatchSetArgs{}
  if err = args.Read(iprot); err != nil {
    iprot.ReadMessageEnd()
    x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
    oprot.WriteMessageBegin("BatchSet", thrift.EXCEPTION, seqId)
    x.Write(oprot)
    oprot.WriteMessageEnd()
    oprot.Flush()
    return false, err
  }

  iprot.ReadMessageEnd()
  result := TableKVBatchSetResult{}
  var err2 error
  if err2 = p.handler.BatchSet(ctx, args.Keys, args.Values); err2 != nil {
    x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing BatchSet: " + err2.Error())
    oprot.WriteMessageBegin("BatchSet", thrift.EXCEPTION, seqId)
    x.Write(oprot)
    oprot.WriteMessageEnd()
    oprot.Flush()
    return true, err2
  }
  if err2 = oprot.WriteMessageBegin("BatchSet", thrift.REPLY, seqId); err2 != nil {
    err = err2
  }
  if err2 = result.Write(oprot); err == nil && err2 != nil {
    err = err2
  }
  if err2 = oprot.WriteMessageEnd(); err == nil && err2 != nil {
    err = err2
  }
  if err2 = oprot.Flush(); err == nil && err2 != nil {
    err = err2
  }
  if err != nil {
    return
  }
  return true, err
}


// HELPER FUNCTIONS AND STRUCTURES

// Attributes:
//  - Key
type TableKVGetArgs struct {
  Key []byte `thrift:"key,1" db:"key" json:"key"`
}

func NewTableKVGetArgs() *TableKVGetArgs {
  return &TableKVGetArgs{}
}


func (p *TableKVGetArgs) GetKey() []byte {
  return p.Key
}
func (p *TableKVGetArgs) Read(iprot thrift.TProtocol) error {
  if _, err := iprot.ReadStructBegin(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
  }


  for {
    _, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
    if err != nil {
      return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
    }
    if fieldTypeId == thrift.STOP { break; }
    switch fieldId {
    case 1:
      if fieldTypeId == thrift.STRING {
        if err := p.ReadField1(iprot); err != nil {
          return err
        }
      } else {
        if err := iprot.Skip(fieldTypeId); err != nil {
          return err
        }
      }
    default:
      if err := iprot.Skip(fieldTypeId); err != nil {
        return err
      }
    }
    if err := iprot.ReadFieldEnd(); err != nil {
      return err
    }
  }
  if err := iprot.ReadStructEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
  }
  return nil
}

func (p *TableKVGetArgs)  ReadField1(iprot thrift.TProtocol) error {
  if v, err := iprot.ReadBinary(); err != nil {
  return thrift.PrependError("error reading field 1: ", err)
} else {
  p.Key = v
}
  return nil
}

func (p *TableKVGetArgs) Write(oprot thrift.TProtocol) error {
  if err := oprot.WriteStructBegin("Get_args"); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err) }
  if p != nil {
    if err := p.writeField1(oprot); err != nil { return err }
  }
  if err := oprot.WriteFieldStop(); err != nil {
    return thrift.PrependError("write field stop error: ", err) }
  if err := oprot.WriteStructEnd(); err != nil {
    return thrift.PrependError("write struct stop error: ", err) }
  return nil
}

func (p *TableKVGetArgs) writeField1(oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin("key", thrift.STRING, 1); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:key: ", p), err) }
  if err := oprot.WriteBinary(p.Key); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.key (1) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 1:key: ", p), err) }
  return err
}

func (p *TableKVGetArgs) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("TableKVGetArgs(%+v)", *p)
}

// Attributes:
//  - Success
type TableKVGetResult struct {
  Success []byte `thrift:"success,0" db:"success" json:"success,omitempty"`
}

func NewTableKVGetResult() *TableKVGetResult {
  return &TableKVGetResult{}
}

var TableKVGetResult_Success_DEFAULT []byte

func (p *TableKVGetResult) GetSuccess() []byte {
  return p.Success
}
func (p *TableKVGetResult) IsSetSuccess() bool {
  return p.Success != nil
}

func (p *TableKVGetResult) Read(iprot thrift.TProtocol) error {
  if _, err := iprot.ReadStructBegin(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
  }


  for {
    _, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
    if err != nil {
      return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
    }
    if fieldTypeId == thrift.STOP { break; }
    switch fieldId {
    case 0:
      if fieldTypeId == thrift.STRING {
        if err := p.ReadField0(iprot); err != nil {
          return err
        }
      } else {
        if err := iprot.Skip(fieldTypeId); err != nil {
          return err
        }
      }
    default:
      if err := iprot.Skip(fieldTypeId); err != nil {
        return err
      }
    }
    if err := iprot.ReadFieldEnd(); err != nil {
      return err
    }
  }
  if err := iprot.ReadStructEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
  }
  return nil
}

func (p *TableKVGetResult)  ReadField0(iprot thrift.TProtocol) error {
  if v, err := iprot.ReadBinary(); err != nil {
  return thrift.PrependError("error reading field 0: ", err)
} else {
  p.Success = v
}
  return nil
}

func (p *TableKVGetResult) Write(oprot thrift.TProtocol) error {
  if err := oprot.WriteStructBegin("Get_result"); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err) }
  if p != nil {
    if err := p.writeField0(oprot); err != nil { return err }
  }
  if err := oprot.WriteFieldStop(); err != nil {
    return thrift.PrependError("write field stop error: ", err) }
  if err := oprot.WriteStructEnd(); err != nil {
    return thrift.PrependError("write struct stop error: ", err) }
  return nil
}

func (p *TableKVGetResult) writeField0(oprot thrift.TProtocol) (err error) {
  if p.IsSetSuccess() {
    if err := oprot.WriteFieldBegin("success", thrift.STRING, 0); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field begin error 0:success: ", p), err) }
    if err := oprot.WriteBinary(p.Success); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T.success (0) field write error: ", p), err) }
    if err := oprot.WriteFieldEnd(); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field end error 0:success: ", p), err) }
  }
  return err
}

func (p *TableKVGetResult) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("TableKVGetResult(%+v)", *p)
}

// Attributes:
//  - Key
//  - Value
type TableKVSetArgs struct {
  Key []byte `thrift:"key,1" db:"key" json:"key"`
  Value []byte `thrift:"value,2" db:"value" json:"value"`
}

func NewTableKVSetArgs() *TableKVSetArgs {
  return &TableKVSetArgs{}
}


func (p *TableKVSetArgs) GetKey() []byte {
  return p.Key
}

func (p *TableKVSetArgs) GetValue() []byte {
  return p.Value
}
func (p *TableKVSetArgs) Read(iprot thrift.TProtocol) error {
  if _, err := iprot.ReadStructBegin(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
  }


  for {
    _, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
    if err != nil {
      return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
    }
    if fieldTypeId == thrift.STOP { break; }
    switch fieldId {
    case 1:
      if fieldTypeId == thrift.STRING {
        if err := p.ReadField1(iprot); err != nil {
          return err
        }
      } else {
        if err := iprot.Skip(fieldTypeId); err != nil {
          return err
        }
      }
    case 2:
      if fieldTypeId == thrift.STRING {
        if err := p.ReadField2(iprot); err != nil {
          return err
        }
      } else {
        if err := iprot.Skip(fieldTypeId); err != nil {
          return err
        }
      }
    default:
      if err := iprot.Skip(fieldTypeId); err != nil {
        return err
      }
    }
    if err := iprot.ReadFieldEnd(); err != nil {
      return err
    }
  }
  if err := iprot.ReadStructEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
  }
  return nil
}

func (p *TableKVSetArgs)  ReadField1(iprot thrift.TProtocol) error {
  if v, err := iprot.ReadBinary(); err != nil {
  return thrift.PrependError("error reading field 1: ", err)
} else {
  p.Key = v
}
  return nil
}

func (p *TableKVSetArgs)  ReadField2(iprot thrift.TProtocol) error {
  if v, err := iprot.ReadBinary(); err != nil {
  return thrift.PrependError("error reading field 2: ", err)
} else {
  p.Value = v
}
  return nil
}

func (p *TableKVSetArgs) Write(oprot thrift.TProtocol) error {
  if err := oprot.WriteStructBegin("Set_args"); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err) }
  if p != nil {
    if err := p.writeField1(oprot); err != nil { return err }
    if err := p.writeField2(oprot); err != nil { return err }
  }
  if err := oprot.WriteFieldStop(); err != nil {
    return thrift.PrependError("write field stop error: ", err) }
  if err := oprot.WriteStructEnd(); err != nil {
    return thrift.PrependError("write struct stop error: ", err) }
  return nil
}

func (p *TableKVSetArgs) writeField1(oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin("key", thrift.STRING, 1); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:key: ", p), err) }
  if err := oprot.WriteBinary(p.Key); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.key (1) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 1:key: ", p), err) }
  return err
}

func (p *TableKVSetArgs) writeField2(oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin("value", thrift.STRING, 2); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:value: ", p), err) }
  if err := oprot.WriteBinary(p.Value); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.value (2) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 2:value: ", p), err) }
  return err
}

func (p *TableKVSetArgs) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("TableKVSetArgs(%+v)", *p)
}

type TableKVSetResult struct {
}

func NewTableKVSetResult() *TableKVSetResult {
  return &TableKVSetResult{}
}

func (p *TableKVSetResult) Read(iprot thrift.TProtocol) error {
  if _, err := iprot.ReadStructBegin(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
  }


  for {
    _, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
    if err != nil {
      return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
    }
    if fieldTypeId == thrift.STOP { break; }
    if err := iprot.Skip(fieldTypeId); err != nil {
      return err
    }
    if err := iprot.ReadFieldEnd(); err != nil {
      return err
    }
  }
  if err := iprot.ReadStructEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
  }
  return nil
}

func (p *TableKVSetResult) Write(oprot thrift.TProtocol) error {
  if err := oprot.WriteStructBegin("Set_result"); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err) }
  if p != nil {
  }
  if err := oprot.WriteFieldStop(); err != nil {
    return thrift.PrependError("write field stop error: ", err) }
  if err := oprot.WriteStructEnd(); err != nil {
    return thrift.PrependError("write struct stop error: ", err) }
  return nil
}

func (p *TableKVSetResult) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("TableKVSetResult(%+v)", *p)
}

// Attributes:
//  - Key
type TableKVExistsArgs struct {
  Key []byte `thrift:"key,1" db:"key" json:"key"`
}

func NewTableKVExistsArgs() *TableKVExistsArgs {
  return &TableKVExistsArgs{}
}


func (p *TableKVExistsArgs) GetKey() []byte {
  return p.Key
}
func (p *TableKVExistsArgs) Read(iprot thrift.TProtocol) error {
  if _, err := iprot.ReadStructBegin(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
  }


  for {
    _, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
    if err != nil {
      return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
    }
    if fieldTypeId == thrift.STOP { break; }
    switch fieldId {
    case 1:
      if fieldTypeId == thrift.STRING {
        if err := p.ReadField1(iprot); err != nil {
          return err
        }
      } else {
        if err := iprot.Skip(fieldTypeId); err != nil {
          return err
        }
      }
    default:
      if err := iprot.Skip(fieldTypeId); err != nil {
        return err
      }
    }
    if err := iprot.ReadFieldEnd(); err != nil {
      return err
    }
  }
  if err := iprot.ReadStructEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
  }
  return nil
}

func (p *TableKVExistsArgs)  ReadField1(iprot thrift.TProtocol) error {
  if v, err := iprot.ReadBinary(); err != nil {
  return thrift.PrependError("error reading field 1: ", err)
} else {
  p.Key = v
}
  return nil
}

func (p *TableKVExistsArgs) Write(oprot thrift.TProtocol) error {
  if err := oprot.WriteStructBegin("Exists_args"); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err) }
  if p != nil {
    if err := p.writeField1(oprot); err != nil { return err }
  }
  if err := oprot.WriteFieldStop(); err != nil {
    return thrift.PrependError("write field stop error: ", err) }
  if err := oprot.WriteStructEnd(); err != nil {
    return thrift.PrependError("write struct stop error: ", err) }
  return nil
}

func (p *TableKVExistsArgs) writeField1(oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin("key", thrift.STRING, 1); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:key: ", p), err) }
  if err := oprot.WriteBinary(p.Key); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.key (1) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 1:key: ", p), err) }
  return err
}

func (p *TableKVExistsArgs) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("TableKVExistsArgs(%+v)", *p)
}

// Attributes:
//  - Success
type TableKVExistsResult struct {
  Success *bool `thrift:"success,0" db:"success" json:"success,omitempty"`
}

func NewTableKVExistsResult() *TableKVExistsResult {
  return &TableKVExistsResult{}
}

var TableKVExistsResult_Success_DEFAULT bool
func (p *TableKVExistsResult) GetSuccess() bool {
  if !p.IsSetSuccess() {
    return TableKVExistsResult_Success_DEFAULT
  }
return *p.Success
}
func (p *TableKVExistsResult) IsSetSuccess() bool {
  return p.Success != nil
}

func (p *TableKVExistsResult) Read(iprot thrift.TProtocol) error {
  if _, err := iprot.ReadStructBegin(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
  }


  for {
    _, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
    if err != nil {
      return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
    }
    if fieldTypeId == thrift.STOP { break; }
    switch fieldId {
    case 0:
      if fieldTypeId == thrift.BOOL {
        if err := p.ReadField0(iprot); err != nil {
          return err
        }
      } else {
        if err := iprot.Skip(fieldTypeId); err != nil {
          return err
        }
      }
    default:
      if err := iprot.Skip(fieldTypeId); err != nil {
        return err
      }
    }
    if err := iprot.ReadFieldEnd(); err != nil {
      return err
    }
  }
  if err := iprot.ReadStructEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
  }
  return nil
}

func (p *TableKVExistsResult)  ReadField0(iprot thrift.TProtocol) error {
  if v, err := iprot.ReadBool(); err != nil {
  return thrift.PrependError("error reading field 0: ", err)
} else {
  p.Success = &v
}
  return nil
}

func (p *TableKVExistsResult) Write(oprot thrift.TProtocol) error {
  if err := oprot.WriteStructBegin("Exists_result"); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err) }
  if p != nil {
    if err := p.writeField0(oprot); err != nil { return err }
  }
  if err := oprot.WriteFieldStop(); err != nil {
    return thrift.PrependError("write field stop error: ", err) }
  if err := oprot.WriteStructEnd(); err != nil {
    return thrift.PrependError("write struct stop error: ", err) }
  return nil
}

func (p *TableKVExistsResult) writeField0(oprot thrift.TProtocol) (err error) {
  if p.IsSetSuccess() {
    if err := oprot.WriteFieldBegin("success", thrift.BOOL, 0); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field begin error 0:success: ", p), err) }
    if err := oprot.WriteBool(bool(*p.Success)); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T.success (0) field write error: ", p), err) }
    if err := oprot.WriteFieldEnd(); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field end error 0:success: ", p), err) }
  }
  return err
}

func (p *TableKVExistsResult) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("TableKVExistsResult(%+v)", *p)
}

// Attributes:
//  - Key
type TableKVDeleteArgs struct {
  Key []byte `thrift:"key,1" db:"key" json:"key"`
}

func NewTableKVDeleteArgs() *TableKVDeleteArgs {
  return &TableKVDeleteArgs{}
}


func (p *TableKVDeleteArgs) GetKey() []byte {
  return p.Key
}
func (p *TableKVDeleteArgs) Read(iprot thrift.TProtocol) error {
  if _, err := iprot.ReadStructBegin(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
  }


  for {
    _, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
    if err != nil {
      return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
    }
    if fieldTypeId == thrift.STOP { break; }
    switch fieldId {
    case 1:
      if fieldTypeId == thrift.STRING {
        if err := p.ReadField1(iprot); err != nil {
          return err
        }
      } else {
        if err := iprot.Skip(fieldTypeId); err != nil {
          return err
        }
      }
    default:
      if err := iprot.Skip(fieldTypeId); err != nil {
        return err
      }
    }
    if err := iprot.ReadFieldEnd(); err != nil {
      return err
    }
  }
  if err := iprot.ReadStructEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
  }
  return nil
}

func (p *TableKVDeleteArgs)  ReadField1(iprot thrift.TProtocol) error {
  if v, err := iprot.ReadBinary(); err != nil {
  return thrift.PrependError("error reading field 1: ", err)
} else {
  p.Key = v
}
  return nil
}

func (p *TableKVDeleteArgs) Write(oprot thrift.TProtocol) error {
  if err := oprot.WriteStructBegin("Delete_args"); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err) }
  if p != nil {
    if err := p.writeField1(oprot); err != nil { return err }
  }
  if err := oprot.WriteFieldStop(); err != nil {
    return thrift.PrependError("write field stop error: ", err) }
  if err := oprot.WriteStructEnd(); err != nil {
    return thrift.PrependError("write struct stop error: ", err) }
  return nil
}

func (p *TableKVDeleteArgs) writeField1(oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin("key", thrift.STRING, 1); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:key: ", p), err) }
  if err := oprot.WriteBinary(p.Key); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.key (1) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 1:key: ", p), err) }
  return err
}

func (p *TableKVDeleteArgs) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("TableKVDeleteArgs(%+v)", *p)
}

type TableKVDeleteResult struct {
}

func NewTableKVDeleteResult() *TableKVDeleteResult {
  return &TableKVDeleteResult{}
}

func (p *TableKVDeleteResult) Read(iprot thrift.TProtocol) error {
  if _, err := iprot.ReadStructBegin(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
  }


  for {
    _, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
    if err != nil {
      return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
    }
    if fieldTypeId == thrift.STOP { break; }
    if err := iprot.Skip(fieldTypeId); err != nil {
      return err
    }
    if err := iprot.ReadFieldEnd(); err != nil {
      return err
    }
  }
  if err := iprot.ReadStructEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
  }
  return nil
}

func (p *TableKVDeleteResult) Write(oprot thrift.TProtocol) error {
  if err := oprot.WriteStructBegin("Delete_result"); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err) }
  if p != nil {
  }
  if err := oprot.WriteFieldStop(); err != nil {
    return thrift.PrependError("write field stop error: ", err) }
  if err := oprot.WriteStructEnd(); err != nil {
    return thrift.PrependError("write struct stop error: ", err) }
  return nil
}

func (p *TableKVDeleteResult) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("TableKVDeleteResult(%+v)", *p)
}

// Attributes:
//  - Keys
//  - Values
type TableKVBatchSetArgs struct {
  Keys [][]byte `thrift:"keys,1" db:"keys" json:"keys"`
  Values [][]byte `thrift:"values,2" db:"values" json:"values"`
}

func NewTableKVBatchSetArgs() *TableKVBatchSetArgs {
  return &TableKVBatchSetArgs{}
}


func (p *TableKVBatchSetArgs) GetKeys() [][]byte {
  return p.Keys
}

func (p *TableKVBatchSetArgs) GetValues() [][]byte {
  return p.Values
}
func (p *TableKVBatchSetArgs) Read(iprot thrift.TProtocol) error {
  if _, err := iprot.ReadStructBegin(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
  }


  for {
    _, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
    if err != nil {
      return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
    }
    if fieldTypeId == thrift.STOP { break; }
    switch fieldId {
    case 1:
      if fieldTypeId == thrift.LIST {
        if err := p.ReadField1(iprot); err != nil {
          return err
        }
      } else {
        if err := iprot.Skip(fieldTypeId); err != nil {
          return err
        }
      }
    case 2:
      if fieldTypeId == thrift.LIST {
        if err := p.ReadField2(iprot); err != nil {
          return err
        }
      } else {
        if err := iprot.Skip(fieldTypeId); err != nil {
          return err
        }
      }
    default:
      if err := iprot.Skip(fieldTypeId); err != nil {
        return err
      }
    }
    if err := iprot.ReadFieldEnd(); err != nil {
      return err
    }
  }
  if err := iprot.ReadStructEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
  }
  return nil
}

func (p *TableKVBatchSetArgs)  ReadField1(iprot thrift.TProtocol) error {
  _, size, err := iprot.ReadListBegin()
  if err != nil {
    return thrift.PrependError("error reading list begin: ", err)
  }
  tSlice := make([][]byte, 0, size)
  p.Keys =  tSlice
  for i := 0; i < size; i ++ {
var _elem12 []byte
    if v, err := iprot.ReadBinary(); err != nil {
    return thrift.PrependError("error reading field 0: ", err)
} else {
    _elem12 = v
}
    p.Keys = append(p.Keys, _elem12)
  }
  if err := iprot.ReadListEnd(); err != nil {
    return thrift.PrependError("error reading list end: ", err)
  }
  return nil
}

func (p *TableKVBatchSetArgs)  ReadField2(iprot thrift.TProtocol) error {
  _, size, err := iprot.ReadListBegin()
  if err != nil {
    return thrift.PrependError("error reading list begin: ", err)
  }
  tSlice := make([][]byte, 0, size)
  p.Values =  tSlice
  for i := 0; i < size; i ++ {
var _elem13 []byte
    if v, err := iprot.ReadBinary(); err != nil {
    return thrift.PrependError("error reading field 0: ", err)
} else {
    _elem13 = v
}
    p.Values = append(p.Values, _elem13)
  }
  if err := iprot.ReadListEnd(); err != nil {
    return thrift.PrependError("error reading list end: ", err)
  }
  return nil
}

func (p *TableKVBatchSetArgs) Write(oprot thrift.TProtocol) error {
  if err := oprot.WriteStructBegin("BatchSet_args"); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err) }
  if p != nil {
    if err := p.writeField1(oprot); err != nil { return err }
    if err := p.writeField2(oprot); err != nil { return err }
  }
  if err := oprot.WriteFieldStop(); err != nil {
    return thrift.PrependError("write field stop error: ", err) }
  if err := oprot.WriteStructEnd(); err != nil {
    return thrift.PrependError("write struct stop error: ", err) }
  return nil
}

func (p *TableKVBatchSetArgs) writeField1(oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin("keys", thrift.LIST, 1); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:keys: ", p), err) }
  if err := oprot.WriteListBegin(thrift.STRING, len(p.Keys)); err != nil {
    return thrift.PrependError("error writing list begin: ", err)
  }
  for _, v := range p.Keys {
    if err := oprot.WriteBinary(v); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T. (0) field write error: ", p), err) }
  }
  if err := oprot.WriteListEnd(); err != nil {
    return thrift.PrependError("error writing list end: ", err)
  }
  if err := oprot.WriteFieldEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 1:keys: ", p), err) }
  return err
}

func (p *TableKVBatchSetArgs) writeField2(oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin("values", thrift.LIST, 2); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:values: ", p), err) }
  if err := oprot.WriteListBegin(thrift.STRING, len(p.Values)); err != nil {
    return thrift.PrependError("error writing list begin: ", err)
  }
  for _, v := range p.Values {
    if err := oprot.WriteBinary(v); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T. (0) field write error: ", p), err) }
  }
  if err := oprot.WriteListEnd(); err != nil {
    return thrift.PrependError("error writing list end: ", err)
  }
  if err := oprot.WriteFieldEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 2:values: ", p), err) }
  return err
}

func (p *TableKVBatchSetArgs) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("TableKVBatchSetArgs(%+v)", *p)
}

type TableKVBatchSetResult struct {
}

func NewTableKVBatchSetResult() *TableKVBatchSetResult {
  return &TableKVBatchSetResult{}
}

func (p *TableKVBatchSetResult) Read(iprot thrift.TProtocol) error {
  if _, err := iprot.ReadStructBegin(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
  }


  for {
    _, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
    if err != nil {
      return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
    }
    if fieldTypeId == thrift.STOP { break; }
    if err := iprot.Skip(fieldTypeId); err != nil {
      return err
    }
    if err := iprot.ReadFieldEnd(); err != nil {
      return err
    }
  }
  if err := iprot.ReadStructEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
  }
  return nil
}

func (p *TableKVBatchSetResult) Write(oprot thrift.TProtocol) error {
  if err := oprot.WriteStructBegin("BatchSet_result"); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err) }
  if p != nil {
  }
  if err := oprot.WriteFieldStop(); err != nil {
    return thrift.PrependError("write field stop error: ", err) }
  if err := oprot.WriteStructEnd(); err != nil {
    return thrift.PrependError("write struct stop error: ", err) }
  return nil
}

func (p *TableKVBatchSetResult) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("TableKVBatchSetResult(%+v)", *p)
}


