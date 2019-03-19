// Code generated by protoc-gen-go. DO NOT EDIT.
// source: modifysocialaccountlistresponseselect/modifysocialaccountlistresponseselect_.proto

/*
Package modifysocialaccountlistresponseselect is a generated protocol buffer package.

It is generated from these files:
	modifysocialaccountlistresponseselect/modifysocialaccountlistresponseselect_.proto

It has these top-level messages:
	ModifySocialAccountListResponseSelect
*/
package modifysocialaccountlistresponseselect

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

type ModifySocialAccountListResponseSelect struct {
	SelectAll        *bool                                  `protobuf:"varint,1,opt,name=selectAll" json:"selectAll,omitempty"`
	Meta             *responsemetaselect.ResponseMetaSelect `protobuf:"bytes,2,opt,name=meta" json:"meta,omitempty"`
	SelectHash       *bool                                  `protobuf:"varint,3,opt,name=selectHash" json:"selectHash,omitempty"`
	Hash             *string                                `protobuf:"bytes,4,opt,name=hash" json:"hash,omitempty"`
	XXX_unrecognized []byte                                 `json:"-"`
}

func (m *ModifySocialAccountListResponseSelect) Reset()         { *m = ModifySocialAccountListResponseSelect{} }
func (m *ModifySocialAccountListResponseSelect) String() string { return proto.CompactTextString(m) }
func (*ModifySocialAccountListResponseSelect) ProtoMessage()    {}
func (*ModifySocialAccountListResponseSelect) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{0}
}

func (m *ModifySocialAccountListResponseSelect) GetSelectAll() bool {
	if m != nil && m.SelectAll != nil {
		return *m.SelectAll
	}
	return false
}

func (m *ModifySocialAccountListResponseSelect) GetMeta() *responsemetaselect.ResponseMetaSelect {
	if m != nil {
		return m.Meta
	}
	return nil
}

func (m *ModifySocialAccountListResponseSelect) GetSelectHash() bool {
	if m != nil && m.SelectHash != nil {
		return *m.SelectHash
	}
	return false
}

func (m *ModifySocialAccountListResponseSelect) GetHash() string {
	if m != nil && m.Hash != nil {
		return *m.Hash
	}
	return ""
}

func init() {
	proto.RegisterType((*ModifySocialAccountListResponseSelect)(nil), "modifysocialaccountlistresponseselect.ModifySocialAccountListResponseSelect")
}

func init() {
	proto.RegisterFile("modifysocialaccountlistresponseselect/modifysocialaccountlistresponseselect_.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 240 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x90, 0x31, 0x4b, 0xc4, 0x30,
	0x18, 0x86, 0x89, 0x76, 0xf0, 0xe2, 0x96, 0xa9, 0x88, 0x48, 0x11, 0x4e, 0x3a, 0x68, 0x83, 0x37,
	0xba, 0x9d, 0x93, 0x83, 0x37, 0x98, 0x6e, 0x2e, 0x12, 0x73, 0xb1, 0x0d, 0xa4, 0xfd, 0x4a, 0xbe,
	0xaf, 0xc3, 0xfd, 0x30, 0xff, 0x9f, 0x5c, 0xd2, 0xc3, 0x83, 0x2e, 0xdd, 0x92, 0x37, 0xef, 0xf3,
	0xbc, 0x10, 0xae, 0x3a, 0xd8, 0xbb, 0x9f, 0x03, 0x82, 0x71, 0xda, 0x6b, 0x63, 0x60, 0xec, 0xc9,
	0x3b, 0xa4, 0x60, 0x71, 0x80, 0x1e, 0x2d, 0x5a, 0x6f, 0x0d, 0xc9, 0x45, 0xad, 0xaf, 0x6a, 0x08,
	0x40, 0x20, 0xd6, 0x8b, 0xda, 0x37, 0x8f, 0xa7, 0x7b, 0x67, 0x49, 0x4f, 0x3b, 0xf3, 0x68, 0x92,
	0xde, 0xff, 0x32, 0xbe, 0xde, 0x45, 0x6f, 0x1d, 0xbd, 0xdb, 0xe4, 0x7d, 0x77, 0x48, 0x6a, 0x82,
	0xea, 0x08, 0x88, 0x5b, 0xbe, 0x4a, 0xe8, 0xd6, 0xfb, 0x9c, 0x15, 0xac, 0xbc, 0x52, 0xff, 0x81,
	0x78, 0xe1, 0xd9, 0x51, 0x9e, 0x5f, 0x14, 0xac, 0xbc, 0xde, 0x3c, 0x54, 0xf3, 0xc5, 0xea, 0xe4,
	0xdb, 0x59, 0xd2, 0xc9, 0xa9, 0x22, 0x23, 0xee, 0x38, 0x4f, 0x95, 0x37, 0x8d, 0x6d, 0x7e, 0x19,
	0xd5, 0x67, 0x89, 0x10, 0x3c, 0x6b, 0x8f, 0x2f, 0x59, 0xc1, 0xca, 0x95, 0x8a, 0xe7, 0xd7, 0xfa,
	0xf3, 0xa3, 0x71, 0xd4, 0x8e, 0xdf, 0x95, 0x81, 0x4e, 0x6e, 0x9e, 0x91, 0x1c, 0xc8, 0x06, 0x9e,
	0x06, 0xaf, 0x0f, 0x4d, 0x80, 0xb1, 0xdf, 0xcb, 0x26, 0x0c, 0x46, 0xa2, 0x69, 0x6d, 0xa7, 0x97,
	0x7d, 0xf4, 0x5f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xb3, 0x56, 0x5b, 0x52, 0xb6, 0x01, 0x00, 0x00,
}
