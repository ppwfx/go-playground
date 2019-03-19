// Code generated by protoc-gen-go. DO NOT EDIT.
// source: floatrangesort/floatrangesort_.proto

/*
Package floatrangesort is a generated protocol buffer package.

It is generated from these files:
	floatrangesort/floatrangesort_.proto

It has these top-level messages:
	FloatRangeSort
*/
package floatrangesort

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
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

type FloatRangeSort struct {
	From             *sortkind.SortKind `protobuf:"varint,1,opt,name=from,enum=sortkind.SortKind" json:"from,omitempty"`
	To               *sortkind.SortKind `protobuf:"varint,2,opt,name=to,enum=sortkind.SortKind" json:"to,omitempty"`
	Hash             *string            `protobuf:"bytes,3,opt,name=hash" json:"hash,omitempty"`
	XXX_unrecognized []byte             `json:"-"`
}

func (m *FloatRangeSort) Reset()                    { *m = FloatRangeSort{} }
func (m *FloatRangeSort) String() string            { return proto.CompactTextString(m) }
func (*FloatRangeSort) ProtoMessage()               {}
func (*FloatRangeSort) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *FloatRangeSort) GetFrom() sortkind.SortKind {
	if m != nil && m.From != nil {
		return *m.From
	}
	return sortkind.SortKind_ASC
}

func (m *FloatRangeSort) GetTo() sortkind.SortKind {
	if m != nil && m.To != nil {
		return *m.To
	}
	return sortkind.SortKind_ASC
}

func (m *FloatRangeSort) GetHash() string {
	if m != nil && m.Hash != nil {
		return *m.Hash
	}
	return ""
}

func init() {
	proto.RegisterType((*FloatRangeSort)(nil), "floatrangesort.FloatRangeSort")
}

func init() { proto.RegisterFile("floatrangesort/floatrangesort_.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 188 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x49, 0xcb, 0xc9, 0x4f,
	0x2c, 0x29, 0x4a, 0xcc, 0x4b, 0x4f, 0x2d, 0xce, 0x2f, 0x2a, 0xd1, 0x47, 0xe5, 0xc6, 0xeb, 0x15,
	0x14, 0xe5, 0x97, 0xe4, 0x0b, 0xf1, 0xa1, 0x0a, 0x4b, 0x49, 0x80, 0xc8, 0xec, 0xcc, 0xbc, 0x14,
	0x7d, 0x18, 0x03, 0xaa, 0x52, 0xa9, 0x80, 0x8b, 0xcf, 0x0d, 0xa4, 0x36, 0x08, 0xa4, 0x36, 0x38,
	0xbf, 0xa8, 0x44, 0x48, 0x8d, 0x8b, 0x25, 0xad, 0x28, 0x3f, 0x57, 0x82, 0x51, 0x81, 0x51, 0x83,
	0xcf, 0x48, 0x48, 0x0f, 0xa6, 0x43, 0x0f, 0x24, 0xeb, 0x9d, 0x99, 0x97, 0x12, 0x04, 0x96, 0x17,
	0x52, 0xe2, 0x62, 0x2a, 0xc9, 0x97, 0x60, 0xc2, 0xa9, 0x8a, 0xa9, 0x24, 0x5f, 0x48, 0x88, 0x8b,
	0x25, 0x23, 0xb1, 0x38, 0x43, 0x82, 0x59, 0x81, 0x51, 0x83, 0x33, 0x08, 0xcc, 0x76, 0xb2, 0x89,
	0xb2, 0x4a, 0xcf, 0x2c, 0xc9, 0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x37, 0x32, 0x2c, 0x2e,
	0xc9, 0xcc, 0xd7, 0x4f, 0xcf, 0xd7, 0x2d, 0xc8, 0x49, 0xac, 0x4c, 0x2f, 0xca, 0x2f, 0xcd, 0x4b,
	0xd1, 0x4f, 0x2f, 0x2a, 0x48, 0xd6, 0x2f, 0x4e, 0xce, 0x48, 0xcd, 0x4d, 0x44, 0xf3, 0x20, 0x20,
	0x00, 0x00, 0xff, 0xff, 0x76, 0x58, 0x0e, 0x58, 0x00, 0x01, 0x00, 0x00,
}