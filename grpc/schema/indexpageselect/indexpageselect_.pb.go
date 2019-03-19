// Code generated by protoc-gen-go. DO NOT EDIT.
// source: indexpageselect/indexpageselect_.proto

/*
Package indexpageselect is a generated protocol buffer package.

It is generated from these files:
	indexpageselect/indexpageselect_.proto

It has these top-level messages:
	IndexPageSelect
*/
package indexpageselect

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

type IndexPageSelect struct {
	SelectAll        *bool   `protobuf:"varint,1,opt,name=selectAll" json:"selectAll,omitempty"`
	Page             *bool   `protobuf:"varint,2,opt,name=page" json:"page,omitempty"`
	SelectHash       *bool   `protobuf:"varint,3,opt,name=selectHash" json:"selectHash,omitempty"`
	Hash             *string `protobuf:"bytes,4,opt,name=hash" json:"hash,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *IndexPageSelect) Reset()                    { *m = IndexPageSelect{} }
func (m *IndexPageSelect) String() string            { return proto.CompactTextString(m) }
func (*IndexPageSelect) ProtoMessage()               {}
func (*IndexPageSelect) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *IndexPageSelect) GetSelectAll() bool {
	if m != nil && m.SelectAll != nil {
		return *m.SelectAll
	}
	return false
}

func (m *IndexPageSelect) GetPage() bool {
	if m != nil && m.Page != nil {
		return *m.Page
	}
	return false
}

func (m *IndexPageSelect) GetSelectHash() bool {
	if m != nil && m.SelectHash != nil {
		return *m.SelectHash
	}
	return false
}

func (m *IndexPageSelect) GetHash() string {
	if m != nil && m.Hash != nil {
		return *m.Hash
	}
	return ""
}

func init() {
	proto.RegisterType((*IndexPageSelect)(nil), "indexpageselect.IndexPageSelect")
}

func init() { proto.RegisterFile("indexpageselect/indexpageselect_.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 176 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0xcb, 0xcc, 0x4b, 0x49,
	0xad, 0x28, 0x48, 0x4c, 0x4f, 0x2d, 0x4e, 0xcd, 0x49, 0x4d, 0x2e, 0xd1, 0x47, 0xe3, 0xc7, 0xeb,
	0x15, 0x14, 0xe5, 0x97, 0xe4, 0x0b, 0xf1, 0xa3, 0x89, 0x2b, 0x95, 0x73, 0xf1, 0x7b, 0x82, 0x84,
	0x02, 0x12, 0xd3, 0x53, 0x83, 0xc1, 0x42, 0x42, 0x32, 0x5c, 0x9c, 0x10, 0x49, 0xc7, 0x9c, 0x1c,
	0x09, 0x46, 0x05, 0x46, 0x0d, 0x8e, 0x20, 0x84, 0x80, 0x90, 0x10, 0x17, 0x0b, 0x48, 0xbb, 0x04,
	0x13, 0x58, 0x02, 0xcc, 0x16, 0x92, 0xe3, 0xe2, 0x82, 0x28, 0xf0, 0x48, 0x2c, 0xce, 0x90, 0x60,
	0x06, 0xcb, 0x20, 0x89, 0x80, 0xf4, 0x64, 0x80, 0x64, 0x58, 0x14, 0x18, 0x35, 0x38, 0x83, 0xc0,
	0x6c, 0x27, 0xdb, 0x28, 0xeb, 0xf4, 0xcc, 0x92, 0x8c, 0xd2, 0x24, 0xbd, 0xe4, 0xfc, 0x5c, 0x7d,
	0x23, 0xc3, 0xe2, 0x92, 0xcc, 0x7c, 0xfd, 0xf4, 0x7c, 0xdd, 0x82, 0x9c, 0xc4, 0xca, 0xf4, 0xa2,
	0xfc, 0xd2, 0xbc, 0x14, 0xfd, 0xf4, 0xa2, 0x82, 0x64, 0xfd, 0xe2, 0xe4, 0x8c, 0xd4, 0xdc, 0x44,
	0x74, 0xff, 0x00, 0x02, 0x00, 0x00, 0xff, 0xff, 0x65, 0x53, 0x86, 0x0d, 0xf1, 0x00, 0x00, 0x00,
}
