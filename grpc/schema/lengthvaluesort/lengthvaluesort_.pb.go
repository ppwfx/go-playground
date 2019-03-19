// Code generated by protoc-gen-go. DO NOT EDIT.
// source: lengthvaluesort/lengthvaluesort_.proto

/*
Package lengthvaluesort is a generated protocol buffer package.

It is generated from these files:
	lengthvaluesort/lengthvaluesort_.proto

It has these top-level messages:
	LengthValueSort
*/
package lengthvaluesort

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import floatrangesort "github.com/21stio/go-playground/grpc/schema/floatrangesort"
import sortkind "github.com/21stio/go-playground/grpc/schema/sortkind"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type LengthValueSort struct {
	Range            *floatrangesort.FloatRangeSort `protobuf:"bytes,1,opt,name=range" json:"range,omitempty"`
	Value            *sortkind.SortKind             `protobuf:"varint,2,opt,name=value,enum=sortkind.SortKind" json:"value,omitempty"`
	IsEstimate       *sortkind.SortKind             `protobuf:"varint,3,opt,name=isEstimate,enum=sortkind.SortKind" json:"isEstimate,omitempty"`
	Hash             *string                        `protobuf:"bytes,4,opt,name=hash" json:"hash,omitempty"`
	XXX_unrecognized []byte                         `json:"-"`
}

func (m *LengthValueSort) Reset()                    { *m = LengthValueSort{} }
func (m *LengthValueSort) String() string            { return proto.CompactTextString(m) }
func (*LengthValueSort) ProtoMessage()               {}
func (*LengthValueSort) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *LengthValueSort) GetRange() *floatrangesort.FloatRangeSort {
	if m != nil {
		return m.Range
	}
	return nil
}

func (m *LengthValueSort) GetValue() sortkind.SortKind {
	if m != nil && m.Value != nil {
		return *m.Value
	}
	return sortkind.SortKind_ASC
}

func (m *LengthValueSort) GetIsEstimate() sortkind.SortKind {
	if m != nil && m.IsEstimate != nil {
		return *m.IsEstimate
	}
	return sortkind.SortKind_ASC
}

func (m *LengthValueSort) GetHash() string {
	if m != nil && m.Hash != nil {
		return *m.Hash
	}
	return ""
}

func init() {
	proto.RegisterType((*LengthValueSort)(nil), "lengthvaluesort.LengthValueSort")
}

func init() { proto.RegisterFile("lengthvaluesort/lengthvaluesort_.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 240 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x90, 0x31, 0x4b, 0xc4, 0x30,
	0x14, 0xc7, 0x89, 0xde, 0x0d, 0x46, 0xf0, 0x20, 0x53, 0xb8, 0x41, 0x8a, 0x88, 0x74, 0x31, 0xc1,
	0xe2, 0x26, 0x2e, 0x82, 0x2e, 0x3a, 0x55, 0x70, 0x70, 0x91, 0xd8, 0xd6, 0x24, 0x98, 0xe6, 0x95,
	0xe4, 0x55, 0xf0, 0xa3, 0xf9, 0xed, 0x24, 0xe9, 0x1d, 0xd8, 0x82, 0xdb, 0xfb, 0x25, 0xbf, 0xff,
	0x3f, 0xe1, 0xd1, 0x0b, 0xd7, 0x79, 0x8d, 0xe6, 0x4b, 0xb9, 0xb1, 0x8b, 0x10, 0x50, 0x2e, 0xf8,
	0x4d, 0x0c, 0x01, 0x10, 0xd8, 0x66, 0x71, 0xbe, 0x3d, 0xff, 0x70, 0xa0, 0x30, 0x28, 0xaf, 0xa7,
	0xdc, 0x1c, 0x77, 0xb1, 0x2d, 0x4f, 0xf0, 0x69, 0x7d, 0x2b, 0xf7, 0xc3, 0xee, 0xe6, 0xec, 0x87,
	0xd0, 0xcd, 0x53, 0xee, 0x7c, 0x49, 0x9d, 0xcf, 0x10, 0x90, 0x5d, 0xd3, 0x75, 0x6e, 0xe0, 0xa4,
	0x20, 0xe5, 0x71, 0x75, 0x2a, 0xe6, 0xa5, 0xe2, 0x21, 0x61, 0x9d, 0x30, 0xe9, 0xf5, 0x24, 0xb3,
	0x92, 0xae, 0xf3, 0xb7, 0xf8, 0x41, 0x41, 0xca, 0x93, 0x8a, 0x89, 0xfd, 0x53, 0x22, 0x59, 0x8f,
	0xd6, 0xb7, 0xf5, 0x24, 0xb0, 0x8a, 0x52, 0x1b, 0xef, 0x23, 0xda, 0x5e, 0x61, 0xc7, 0x0f, 0xff,
	0xd5, 0xff, 0x58, 0x8c, 0xd1, 0x95, 0x51, 0xd1, 0xf0, 0x55, 0x41, 0xca, 0xa3, 0x3a, 0xcf, 0x77,
	0xb7, 0xaf, 0x37, 0xda, 0xa2, 0x19, 0xdf, 0x45, 0x03, 0xbd, 0xac, 0xae, 0x22, 0x5a, 0x90, 0x1a,
	0x2e, 0x07, 0xa7, 0xbe, 0x75, 0x80, 0xd1, 0xb7, 0x52, 0x87, 0xa1, 0x91, 0xb1, 0x31, 0x5d, 0xaf,
	0x96, 0x2b, 0xfd, 0x0d, 0x00, 0x00, 0xff, 0xff, 0xea, 0x8f, 0xb2, 0xd4, 0x74, 0x01, 0x00, 0x00,
}