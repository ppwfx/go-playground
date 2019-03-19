// Code generated by protoc-gen-go. DO NOT EDIT.
// source: modifyservicelistresponseselect/modifyservicelistresponseselect_.proto

/*
Package modifyservicelistresponseselect is a generated protocol buffer package.

It is generated from these files:
	modifyservicelistresponseselect/modifyservicelistresponseselect_.proto

It has these top-level messages:
	ModifyServiceListResponseSelect
*/
package modifyservicelistresponseselect

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

type ModifyServiceListResponseSelect struct {
	SelectAll        *bool                                  `protobuf:"varint,1,opt,name=selectAll" json:"selectAll,omitempty"`
	Meta             *responsemetaselect.ResponseMetaSelect `protobuf:"bytes,2,opt,name=meta" json:"meta,omitempty"`
	SelectHash       *bool                                  `protobuf:"varint,3,opt,name=selectHash" json:"selectHash,omitempty"`
	Hash             *string                                `protobuf:"bytes,4,opt,name=hash" json:"hash,omitempty"`
	XXX_unrecognized []byte                                 `json:"-"`
}

func (m *ModifyServiceListResponseSelect) Reset()                    { *m = ModifyServiceListResponseSelect{} }
func (m *ModifyServiceListResponseSelect) String() string            { return proto.CompactTextString(m) }
func (*ModifyServiceListResponseSelect) ProtoMessage()               {}
func (*ModifyServiceListResponseSelect) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *ModifyServiceListResponseSelect) GetSelectAll() bool {
	if m != nil && m.SelectAll != nil {
		return *m.SelectAll
	}
	return false
}

func (m *ModifyServiceListResponseSelect) GetMeta() *responsemetaselect.ResponseMetaSelect {
	if m != nil {
		return m.Meta
	}
	return nil
}

func (m *ModifyServiceListResponseSelect) GetSelectHash() bool {
	if m != nil && m.SelectHash != nil {
		return *m.SelectHash
	}
	return false
}

func (m *ModifyServiceListResponseSelect) GetHash() string {
	if m != nil && m.Hash != nil {
		return *m.Hash
	}
	return ""
}

func init() {
	proto.RegisterType((*ModifyServiceListResponseSelect)(nil), "modifyservicelistresponseselect.ModifyServiceListResponseSelect")
}

func init() {
	proto.RegisterFile("modifyservicelistresponseselect/modifyservicelistresponseselect_.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 235 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0xd0, 0x31, 0x4b, 0xc4, 0x30,
	0x14, 0x07, 0x70, 0xa2, 0x1d, 0xbc, 0xb8, 0x65, 0x2a, 0x22, 0x5e, 0x71, 0x90, 0x0e, 0xda, 0xe0,
	0x8d, 0x6e, 0x3a, 0x88, 0xa0, 0x5d, 0x7a, 0x9b, 0x8b, 0xc4, 0xdc, 0xb3, 0x0d, 0xa4, 0xf7, 0x4a,
	0xde, 0x3b, 0xe1, 0x3e, 0x95, 0x5f, 0x51, 0x2e, 0xe9, 0xa1, 0xd0, 0xa1, 0x5b, 0xf2, 0x7f, 0xff,
	0xf7, 0x0b, 0x44, 0x3e, 0xf7, 0xb8, 0x71, 0x5f, 0x7b, 0x82, 0xf0, 0xed, 0x2c, 0x78, 0x47, 0x1c,
	0x80, 0x06, 0xdc, 0x12, 0x10, 0x78, 0xb0, 0xac, 0x67, 0xe6, 0x1f, 0xd5, 0x10, 0x90, 0x51, 0x2d,
	0x67, 0x7a, 0x17, 0xb7, 0xc7, 0x7b, 0x0f, 0x6c, 0x46, 0x7b, 0x1a, 0x8d, 0xdc, 0xf5, 0x8f, 0x90,
	0xcb, 0x3a, 0x8a, 0xeb, 0x24, 0xbe, 0x39, 0xe2, 0x66, 0xac, 0xaf, 0x63, 0x55, 0x5d, 0xca, 0x45,
	0x5a, 0x7a, 0xf4, 0x3e, 0x17, 0x85, 0x28, 0xcf, 0x9a, 0xbf, 0x40, 0x3d, 0xc8, 0xec, 0xc0, 0xe6,
	0x27, 0x85, 0x28, 0xcf, 0x57, 0x37, 0xd5, 0xf4, 0xad, 0xea, 0xe8, 0xd5, 0xc0, 0x26, 0x99, 0x4d,
	0xdc, 0x51, 0x57, 0x52, 0xa6, 0xca, 0x8b, 0xa1, 0x2e, 0x3f, 0x8d, 0xf4, 0xbf, 0x44, 0x29, 0x99,
	0x75, 0x87, 0x49, 0x56, 0x88, 0x72, 0xd1, 0xc4, 0xf3, 0x53, 0xfd, 0xfe, 0xda, 0x3a, 0xee, 0x76,
	0x9f, 0x95, 0xc5, 0x5e, 0xaf, 0xee, 0x89, 0x1d, 0xea, 0x16, 0xef, 0x06, 0x6f, 0xf6, 0x6d, 0xc0,
	0xdd, 0x76, 0xa3, 0xdb, 0x30, 0x58, 0x4d, 0xb6, 0x83, 0xde, 0xcc, 0x7d, 0xeb, 0x6f, 0x00, 0x00,
	0x00, 0xff, 0xff, 0xe7, 0xae, 0x76, 0x0b, 0x98, 0x01, 0x00, 0x00,
}