// Code generated by protoc-gen-go. DO NOT EDIT.
// source: rating/rating_.proto

/*
Package rating is a generated protocol buffer package.

It is generated from these files:
	rating/rating_.proto

It has these top-level messages:
	Rating
*/
package rating

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

type Rating struct {
	Value            *float64 `protobuf:"fixed64,1,opt,name=value" json:"value,omitempty"`
	Of               *float64 `protobuf:"fixed64,2,opt,name=of" json:"of,omitempty"`
	Count            *int32   `protobuf:"varint,3,opt,name=count" json:"count,omitempty"`
	Hash             *string  `protobuf:"bytes,4,opt,name=hash" json:"hash,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *Rating) Reset()                    { *m = Rating{} }
func (m *Rating) String() string            { return proto.CompactTextString(m) }
func (*Rating) ProtoMessage()               {}
func (*Rating) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Rating) GetValue() float64 {
	if m != nil && m.Value != nil {
		return *m.Value
	}
	return 0
}

func (m *Rating) GetOf() float64 {
	if m != nil && m.Of != nil {
		return *m.Of
	}
	return 0
}

func (m *Rating) GetCount() int32 {
	if m != nil && m.Count != nil {
		return *m.Count
	}
	return 0
}

func (m *Rating) GetHash() string {
	if m != nil && m.Hash != nil {
		return *m.Hash
	}
	return ""
}

func init() {
	proto.RegisterType((*Rating)(nil), "rating.Rating")
}

func init() { proto.RegisterFile("rating/rating_.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 164 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x2c, 0x8e, 0xbf, 0x0e, 0xc2, 0x20,
	0x1c, 0x84, 0x43, 0x6d, 0x9b, 0xc8, 0xe0, 0x40, 0x3a, 0x30, 0x36, 0x4e, 0x5d, 0x2c, 0xb1, 0xf1,
	0x09, 0x7c, 0x04, 0x26, 0xe3, 0x62, 0x10, 0x5b, 0x20, 0x69, 0xf9, 0x11, 0xfe, 0x98, 0xf8, 0xf6,
	0x46, 0x70, 0xba, 0xef, 0xbb, 0xdc, 0x70, 0xb8, 0xf3, 0x22, 0x1a, 0xab, 0x58, 0x89, 0xc7, 0xe8,
	0x3c, 0x44, 0x20, 0x6d, 0xd1, 0xe3, 0x0d, 0xb7, 0x3c, 0x13, 0xe9, 0x70, 0xf3, 0x16, 0x6b, 0x9a,
	0x29, 0xea, 0xd1, 0x80, 0x78, 0x11, 0x72, 0xc0, 0x15, 0x2c, 0xb4, 0xca, 0x55, 0x05, 0xcb, 0x6f,
	0x25, 0x21, 0xd9, 0x48, 0x77, 0x3d, 0x1a, 0x1a, 0x5e, 0x84, 0x10, 0x5c, 0x6b, 0x11, 0x34, 0xad,
	0x7b, 0x34, 0xec, 0x79, 0xe6, 0xeb, 0xe5, 0x3e, 0x29, 0x13, 0x75, 0x7a, 0x8e, 0x12, 0x36, 0x36,
	0x9d, 0x43, 0x34, 0xc0, 0x14, 0x9c, 0xdc, 0x2a, 0x3e, 0xca, 0x43, 0xb2, 0x2f, 0xa6, 0xbc, 0x93,
	0x2c, 0x48, 0x3d, 0x6f, 0xe2, 0x7f, 0xef, 0x1b, 0x00, 0x00, 0xff, 0xff, 0x59, 0x92, 0x99, 0x29,
	0xae, 0x00, 0x00, 0x00,
}
