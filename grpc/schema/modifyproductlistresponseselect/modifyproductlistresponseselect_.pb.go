// Code generated by protoc-gen-go. DO NOT EDIT.
// source: modifyproductlistresponseselect/modifyproductlistresponseselect_.proto

/*
Package modifyproductlistresponseselect is a generated protocol buffer package.

It is generated from these files:
	modifyproductlistresponseselect/modifyproductlistresponseselect_.proto

It has these top-level messages:
	ModifyProductListResponseSelect
*/
package modifyproductlistresponseselect

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import responsemetaselect "github.com/21stio/go-playground/grpc/schema/responsemetaselect"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type ModifyProductListResponseSelect struct {
	SelectAll        *bool                                  `protobuf:"varint,1,opt,name=selectAll" json:"selectAll,omitempty"`
	Meta             *responsemetaselect.ResponseMetaSelect `protobuf:"bytes,2,opt,name=meta" json:"meta,omitempty"`
	SelectHash       *bool                                  `protobuf:"varint,3,opt,name=selectHash" json:"selectHash,omitempty"`
	Hash             *string                                `protobuf:"bytes,4,opt,name=hash" json:"hash,omitempty"`
	XXX_unrecognized []byte                                 `json:"-"`
}

func (m *ModifyProductListResponseSelect) Reset()                    { *m = ModifyProductListResponseSelect{} }
func (m *ModifyProductListResponseSelect) String() string            { return proto.CompactTextString(m) }
func (*ModifyProductListResponseSelect) ProtoMessage()               {}
func (*ModifyProductListResponseSelect) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *ModifyProductListResponseSelect) GetSelectAll() bool {
	if m != nil && m.SelectAll != nil {
		return *m.SelectAll
	}
	return false
}

func (m *ModifyProductListResponseSelect) GetMeta() *responsemetaselect.ResponseMetaSelect {
	if m != nil {
		return m.Meta
	}
	return nil
}

func (m *ModifyProductListResponseSelect) GetSelectHash() bool {
	if m != nil && m.SelectHash != nil {
		return *m.SelectHash
	}
	return false
}

func (m *ModifyProductListResponseSelect) GetHash() string {
	if m != nil && m.Hash != nil {
		return *m.Hash
	}
	return ""
}

func init() {
	proto.RegisterType((*ModifyProductListResponseSelect)(nil), "modifyproductlistresponseselect.ModifyProductListResponseSelect")
}

func init() {
	proto.RegisterFile("modifyproductlistresponseselect/modifyproductlistresponseselect_.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 234 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0xd0, 0x31, 0x4b, 0xc4, 0x30,
	0x14, 0x07, 0x70, 0xa2, 0x1d, 0xbc, 0xb8, 0x65, 0x2a, 0x22, 0x5e, 0x71, 0x90, 0x0e, 0xda, 0xe0,
	0x8d, 0x6e, 0x3a, 0x88, 0xa0, 0x05, 0xa9, 0x9b, 0x8b, 0xc4, 0x34, 0xb6, 0x81, 0xf4, 0x5e, 0xc8,
	0x7b, 0x1d, 0xee, 0x53, 0xf9, 0x15, 0xe5, 0x92, 0x1e, 0x1e, 0x74, 0xe8, 0x96, 0xfc, 0xdf, 0xff,
	0xfd, 0x02, 0xe1, 0xcf, 0x03, 0xb4, 0xf6, 0x67, 0xe7, 0x03, 0xb4, 0xa3, 0x26, 0x67, 0x91, 0x82,
	0x41, 0x0f, 0x5b, 0x34, 0x68, 0x9c, 0xd1, 0x24, 0x17, 0xe6, 0x5f, 0x95, 0x0f, 0x40, 0x20, 0xd6,
	0x0b, 0xbd, 0x8b, 0xdb, 0xc3, 0x7d, 0x30, 0xa4, 0x26, 0x7b, 0x1e, 0x4d, 0xdc, 0xf5, 0x2f, 0xe3,
	0xeb, 0x3a, 0x8a, 0xef, 0x49, 0x7c, 0xb3, 0x48, 0xcd, 0x54, 0xff, 0x88, 0x55, 0x71, 0xc9, 0x57,
	0x69, 0xe9, 0xd1, 0xb9, 0x9c, 0x15, 0xac, 0x3c, 0x6b, 0xfe, 0x03, 0xf1, 0xc0, 0xb3, 0x3d, 0x9b,
	0x9f, 0x14, 0xac, 0x3c, 0xdf, 0xdc, 0x54, 0xf3, 0xb7, 0xaa, 0x83, 0x57, 0x1b, 0x52, 0xc9, 0x6c,
	0xe2, 0x8e, 0xb8, 0xe2, 0x3c, 0x55, 0x5e, 0x14, 0xf6, 0xf9, 0x69, 0xa4, 0x8f, 0x12, 0x21, 0x78,
	0xd6, 0xef, 0x27, 0x59, 0xc1, 0xca, 0x55, 0x13, 0xcf, 0x4f, 0xf5, 0xe7, 0x6b, 0x67, 0xa9, 0x1f,
	0xbf, 0x2b, 0x0d, 0x83, 0xdc, 0xdc, 0x23, 0x59, 0x90, 0x1d, 0xdc, 0x79, 0xa7, 0x76, 0x5d, 0x80,
	0x71, 0xdb, 0xca, 0x2e, 0x78, 0x2d, 0x51, 0xf7, 0x66, 0x50, 0x4b, 0xdf, 0xfa, 0x17, 0x00, 0x00,
	0xff, 0xff, 0xc9, 0x58, 0x2b, 0x3a, 0x98, 0x01, 0x00, 0x00,
}
