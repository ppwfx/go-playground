// Code generated by protoc-gen-go. DO NOT EDIT.
// source: floatrange/floatrange_.proto

/*
Package floatrange is a generated protocol buffer package.

It is generated from these files:
	floatrange/floatrange_.proto

It has these top-level messages:
	FloatRange
*/
package floatrange

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type FloatRange struct {
	From             *float64 `protobuf:"fixed64,1,opt,name=from" json:"from,omitempty"`
	To               *float64 `protobuf:"fixed64,2,opt,name=to" json:"to,omitempty"`
	Hash             *string  `protobuf:"bytes,3,opt,name=hash" json:"hash,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *FloatRange) Reset()                    { *m = FloatRange{} }
func (m *FloatRange) String() string            { return proto.CompactTextString(m) }
func (*FloatRange) ProtoMessage()               {}
func (*FloatRange) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *FloatRange) GetFrom() float64 {
	if m != nil && m.From != nil {
		return *m.From
	}
	return 0
}

func (m *FloatRange) GetTo() float64 {
	if m != nil && m.To != nil {
		return *m.To
	}
	return 0
}

func (m *FloatRange) GetHash() string {
	if m != nil && m.Hash != nil {
		return *m.Hash
	}
	return ""
}

func init() {
	proto.RegisterType((*FloatRange)(nil), "floatrange.FloatRange")
}

func init() { proto.RegisterFile("floatrange/floatrange_.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 153 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x49, 0xcb, 0xc9, 0x4f,
	0x2c, 0x29, 0x4a, 0xcc, 0x4b, 0x4f, 0xd5, 0x47, 0x30, 0xe3, 0xf5, 0x0a, 0x8a, 0xf2, 0x4b, 0xf2,
	0x85, 0xb8, 0x10, 0x42, 0x4a, 0x2e, 0x5c, 0x5c, 0x6e, 0x20, 0x5e, 0x10, 0x88, 0x27, 0x24, 0xc4,
	0xc5, 0x92, 0x56, 0x94, 0x9f, 0x2b, 0xc1, 0xa8, 0xc0, 0xa8, 0xc1, 0x18, 0x04, 0x66, 0x0b, 0xf1,
	0x71, 0x31, 0x95, 0xe4, 0x4b, 0x30, 0x81, 0x45, 0x98, 0x4a, 0xf2, 0x41, 0x6a, 0x32, 0x12, 0x8b,
	0x33, 0x24, 0x98, 0x15, 0x18, 0x35, 0x38, 0x83, 0xc0, 0x6c, 0x27, 0x8b, 0x28, 0xb3, 0xf4, 0xcc,
	0x92, 0x8c, 0xd2, 0x24, 0xbd, 0xe4, 0xfc, 0x5c, 0x7d, 0x23, 0xc3, 0xe2, 0x92, 0xcc, 0x7c, 0xfd,
	0xf4, 0x7c, 0xdd, 0x82, 0x9c, 0xc4, 0xca, 0xf4, 0xa2, 0xfc, 0xd2, 0xbc, 0x14, 0xfd, 0xf4, 0xa2,
	0x82, 0x64, 0xfd, 0xe2, 0xe4, 0x8c, 0xd4, 0xdc, 0x44, 0x24, 0x27, 0x01, 0x02, 0x00, 0x00, 0xff,
	0xff, 0xd2, 0xd3, 0x5c, 0x43, 0xaa, 0x00, 0x00, 0x00,
}
