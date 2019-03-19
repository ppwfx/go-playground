// Code generated by protoc-gen-go. DO NOT EDIT.
// source: cursorpageselect/cursorpageselect_.proto

/*
Package cursorpageselect is a generated protocol buffer package.

It is generated from these files:
	cursorpageselect/cursorpageselect_.proto

It has these top-level messages:
	CursorPageSelect
*/
package cursorpageselect

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

type CursorPageSelect struct {
	SelectAll        *bool   `protobuf:"varint,1,opt,name=selectAll" json:"selectAll,omitempty"`
	Prev             *bool   `protobuf:"varint,2,opt,name=prev" json:"prev,omitempty"`
	Next             *bool   `protobuf:"varint,3,opt,name=next" json:"next,omitempty"`
	SelectHash       *bool   `protobuf:"varint,4,opt,name=selectHash" json:"selectHash,omitempty"`
	Hash             *string `protobuf:"bytes,5,opt,name=hash" json:"hash,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *CursorPageSelect) Reset()                    { *m = CursorPageSelect{} }
func (m *CursorPageSelect) String() string            { return proto.CompactTextString(m) }
func (*CursorPageSelect) ProtoMessage()               {}
func (*CursorPageSelect) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *CursorPageSelect) GetSelectAll() bool {
	if m != nil && m.SelectAll != nil {
		return *m.SelectAll
	}
	return false
}

func (m *CursorPageSelect) GetPrev() bool {
	if m != nil && m.Prev != nil {
		return *m.Prev
	}
	return false
}

func (m *CursorPageSelect) GetNext() bool {
	if m != nil && m.Next != nil {
		return *m.Next
	}
	return false
}

func (m *CursorPageSelect) GetSelectHash() bool {
	if m != nil && m.SelectHash != nil {
		return *m.SelectHash
	}
	return false
}

func (m *CursorPageSelect) GetHash() string {
	if m != nil && m.Hash != nil {
		return *m.Hash
	}
	return ""
}

func init() {
	proto.RegisterType((*CursorPageSelect)(nil), "cursorpageselect.CursorPageSelect")
}

func init() { proto.RegisterFile("cursorpageselect/cursorpageselect_.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 191 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0xcf, 0xb1, 0x0a, 0xc2, 0x30,
	0x14, 0x85, 0x61, 0xa2, 0x15, 0x6c, 0xa6, 0x92, 0x29, 0x83, 0x48, 0x71, 0xea, 0x62, 0x83, 0xce,
	0x22, 0xa8, 0x8b, 0xa3, 0xd4, 0xcd, 0x45, 0x62, 0x0c, 0x49, 0x21, 0x6d, 0x42, 0x92, 0x8a, 0xbe,
	0x83, 0x0f, 0x2d, 0x49, 0x07, 0xa5, 0x6e, 0x87, 0xef, 0xe7, 0x0e, 0x17, 0x16, 0xac, 0xb3, 0x4e,
	0x5b, 0x43, 0x05, 0x77, 0x5c, 0x71, 0xe6, 0xc9, 0x10, 0xae, 0xa5, 0xb1, 0xda, 0x6b, 0x94, 0x0d,
	0xc3, 0xe2, 0x0d, 0x60, 0x76, 0x88, 0x78, 0xa2, 0x82, 0x9f, 0x23, 0xa2, 0x19, 0x4c, 0xfb, 0xbc,
	0x53, 0x0a, 0x83, 0x1c, 0x14, 0xd3, 0xea, 0x0b, 0x08, 0xc1, 0xc4, 0x58, 0xfe, 0xc0, 0xa3, 0x18,
	0xe2, 0x0e, 0xd6, 0xf2, 0xa7, 0xc7, 0xe3, 0xde, 0xc2, 0x46, 0x73, 0x08, 0xfb, 0xa3, 0x23, 0x75,
	0x12, 0x27, 0xb1, 0xfc, 0x48, 0xb8, 0x91, 0xa1, 0x4c, 0x72, 0x50, 0xa4, 0x55, 0xdc, 0xfb, 0xed,
	0x65, 0x23, 0x6a, 0x2f, 0xbb, 0x5b, 0xc9, 0x74, 0x43, 0xd6, 0x2b, 0xe7, 0x6b, 0x4d, 0x84, 0x5e,
	0x1a, 0x45, 0x5f, 0xc2, 0xea, 0xae, 0xbd, 0x13, 0x61, 0x0d, 0x23, 0x8e, 0x49, 0xde, 0xd0, 0xbf,
	0x3f, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xde, 0xd1, 0xd3, 0x8e, 0x0b, 0x01, 0x00, 0x00,
}
