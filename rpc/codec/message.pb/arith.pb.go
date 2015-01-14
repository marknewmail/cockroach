// Code generated by protoc-gen-gogo.
// source: arith.proto
// DO NOT EDIT!

/*
Package message is a generated protocol buffer package.

It is generated from these files:
	arith.proto
	echo.proto

It has these top-level messages:
	ArithRequest
	ArithResponse
*/
package message

import proto "github.com/gogo/protobuf/proto"
import math "math"

// discarding unused import gogoproto "code.google.com/p/gogoprotobuf/gogoproto/gogo.pb"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

type ArithRequest struct {
	A                int32  `protobuf:"varint,1,opt,name=a" json:"a"`
	B                int32  `protobuf:"varint,2,opt,name=b" json:"b"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *ArithRequest) Reset()         { *m = ArithRequest{} }
func (m *ArithRequest) String() string { return proto.CompactTextString(m) }
func (*ArithRequest) ProtoMessage()    {}

func (m *ArithRequest) GetA() int32 {
	if m != nil {
		return m.A
	}
	return 0
}

func (m *ArithRequest) GetB() int32 {
	if m != nil {
		return m.B
	}
	return 0
}

type ArithResponse struct {
	C                int32  `protobuf:"varint,1,opt,name=c" json:"c"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *ArithResponse) Reset()         { *m = ArithResponse{} }
func (m *ArithResponse) String() string { return proto.CompactTextString(m) }
func (*ArithResponse) ProtoMessage()    {}

func (m *ArithResponse) GetC() int32 {
	if m != nil {
		return m.C
	}
	return 0
}

func init() {
}
