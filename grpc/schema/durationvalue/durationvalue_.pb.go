// Code generated by protoc-gen-go. DO NOT EDIT.
// source: durationvalue/durationvalue_.proto

/*
Package durationvalue is a generated protocol buffer package.

It is generated from these files:
	durationvalue/durationvalue_.proto

It has these top-level messages:
	DurationValue
*/
package durationvalue

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import valuekind "github.com/21stio/go-playground/grpc/schema/valuekind"
import durationunit "github.com/21stio/go-playground/grpc/schema/durationunit"
import floatrange "github.com/21stio/go-playground/grpc/schema/floatrange"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type DurationValue struct {
	Kind             *valuekind.ValueKind       `protobuf:"varint,1,opt,name=kind,enum=valuekind.ValueKind" json:"kind,omitempty"`
	Unit             *durationunit.DurationUnit `protobuf:"varint,2,opt,name=unit,enum=durationunit.DurationUnit" json:"unit,omitempty"`
	Range            *floatrange.FloatRange     `protobuf:"bytes,3,opt,name=range" json:"range,omitempty"`
	Value            *float64                   `protobuf:"fixed64,4,opt,name=value" json:"value,omitempty"`
	IsEstimate       *bool                      `protobuf:"varint,5,opt,name=isEstimate" json:"isEstimate,omitempty"`
	Hash             *string                    `protobuf:"bytes,6,opt,name=hash" json:"hash,omitempty"`
	XXX_unrecognized []byte                     `json:"-"`
}

func (m *DurationValue) Reset()                    { *m = DurationValue{} }
func (m *DurationValue) String() string            { return proto.CompactTextString(m) }
func (*DurationValue) ProtoMessage()               {}
func (*DurationValue) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *DurationValue) GetKind() valuekind.ValueKind {
	if m != nil && m.Kind != nil {
		return *m.Kind
	}
	return valuekind.ValueKind_VALUE
}

func (m *DurationValue) GetUnit() durationunit.DurationUnit {
	if m != nil && m.Unit != nil {
		return *m.Unit
	}
	return durationunit.DurationUnit_NANOSECOND
}

func (m *DurationValue) GetRange() *floatrange.FloatRange {
	if m != nil {
		return m.Range
	}
	return nil
}

func (m *DurationValue) GetValue() float64 {
	if m != nil && m.Value != nil {
		return *m.Value
	}
	return 0
}

func (m *DurationValue) GetIsEstimate() bool {
	if m != nil && m.IsEstimate != nil {
		return *m.IsEstimate
	}
	return false
}

func (m *DurationValue) GetHash() string {
	if m != nil && m.Hash != nil {
		return *m.Hash
	}
	return ""
}

func init() {
	proto.RegisterType((*DurationValue)(nil), "durationvalue.DurationValue")
}

func init() { proto.RegisterFile("durationvalue/durationvalue_.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 274 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x91, 0x31, 0x4b, 0xc4, 0x30,
	0x14, 0xc7, 0x89, 0xb6, 0xa2, 0x91, 0x73, 0x08, 0x87, 0x84, 0x22, 0x12, 0x6e, 0xca, 0xa0, 0x09,
	0x76, 0x13, 0x37, 0x51, 0x17, 0xb7, 0x80, 0x0e, 0x2e, 0x12, 0xdb, 0xda, 0x06, 0xdb, 0xa4, 0x34,
	0x89, 0xe0, 0xc7, 0xf5, 0x9b, 0x1c, 0xc9, 0xb5, 0x77, 0xed, 0x94, 0xff, 0x7b, 0xef, 0x97, 0xf7,
	0x0b, 0x04, 0x6e, 0x4a, 0x3f, 0x48, 0xa7, 0x8c, 0xfe, 0x95, 0xad, 0xaf, 0xf8, 0xa2, 0xfa, 0x64,
	0xfd, 0x60, 0x9c, 0x41, 0xab, 0x45, 0x37, 0xcb, 0xe2, 0xf1, 0xa3, 0x74, 0xc9, 0xf7, 0x69, 0x44,
	0x33, 0x32, 0xa1, 0x5e, 0x2b, 0xc7, 0xe7, 0xc5, 0x44, 0x5c, 0x7d, 0xb7, 0x46, 0xba, 0x41, 0xea,
	0xba, 0xe2, 0x87, 0x38, 0x4e, 0x37, 0xff, 0x00, 0xae, 0x9e, 0xc6, 0x5b, 0xef, 0x61, 0x39, 0xa2,
	0x30, 0x09, 0x02, 0x0c, 0x08, 0xa0, 0x17, 0xf9, 0x9a, 0xed, 0x95, 0x2c, 0xce, 0x5f, 0x95, 0x2e,
	0x45, 0x24, 0x10, 0x83, 0x49, 0x10, 0xe1, 0xa3, 0x48, 0x66, 0x6c, 0x6e, 0x67, 0xd3, 0xd2, 0x37,
	0xad, 0x9c, 0x88, 0x1c, 0xba, 0x81, 0x69, 0x74, 0xe3, 0x63, 0x02, 0xe8, 0x79, 0x7e, 0xc9, 0x0e,
	0xcf, 0x61, 0x2f, 0x21, 0x8a, 0x10, 0xc5, 0x0e, 0x42, 0x6b, 0x98, 0x46, 0x35, 0x4e, 0x08, 0xa0,
	0x40, 0xec, 0x0a, 0x74, 0x0d, 0xa1, 0xb2, 0xcf, 0xd6, 0xa9, 0x4e, 0xba, 0x0a, 0xa7, 0x04, 0xd0,
	0x53, 0x31, 0xeb, 0x20, 0x04, 0x93, 0x46, 0xda, 0x06, 0x9f, 0x10, 0x40, 0xcf, 0x44, 0xcc, 0x8f,
	0x0f, 0x1f, 0xf7, 0xb5, 0x72, 0x8d, 0xff, 0x62, 0x85, 0xe9, 0x78, 0x7e, 0x67, 0x9d, 0x32, 0xbc,
	0x36, 0xb7, 0x7d, 0x2b, 0xff, 0xea, 0xc1, 0x78, 0x5d, 0xf2, 0x7a, 0xe8, 0x0b, 0x6e, 0x8b, 0xa6,
	0xea, 0xe4, 0xf2, 0x4b, 0xb6, 0x01, 0x00, 0x00, 0xff, 0xff, 0x49, 0x39, 0x70, 0x47, 0xb0, 0x01,
	0x00, 0x00,
}
