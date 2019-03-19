// Code generated by protoc-gen-go. DO NOT EDIT.
// source: deletecompaniesresponseselect/deletecompaniesresponseselect_.proto

/*
Package deletecompaniesresponseselect is a generated protocol buffer package.

It is generated from these files:
	deletecompaniesresponseselect/deletecompaniesresponseselect_.proto

It has these top-level messages:
	DeleteCompaniesResponseSelect
*/
package deletecompaniesresponseselect

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

type DeleteCompaniesResponseSelect struct {
	SelectAll        *bool                                  `protobuf:"varint,1,opt,name=selectAll" json:"selectAll,omitempty"`
	Meta             *responsemetaselect.ResponseMetaSelect `protobuf:"bytes,2,opt,name=meta" json:"meta,omitempty"`
	SelectHash       *bool                                  `protobuf:"varint,3,opt,name=selectHash" json:"selectHash,omitempty"`
	Hash             *string                                `protobuf:"bytes,4,opt,name=hash" json:"hash,omitempty"`
	XXX_unrecognized []byte                                 `json:"-"`
}

func (m *DeleteCompaniesResponseSelect) Reset()                    { *m = DeleteCompaniesResponseSelect{} }
func (m *DeleteCompaniesResponseSelect) String() string            { return proto.CompactTextString(m) }
func (*DeleteCompaniesResponseSelect) ProtoMessage()               {}
func (*DeleteCompaniesResponseSelect) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *DeleteCompaniesResponseSelect) GetSelectAll() bool {
	if m != nil && m.SelectAll != nil {
		return *m.SelectAll
	}
	return false
}

func (m *DeleteCompaniesResponseSelect) GetMeta() *responsemetaselect.ResponseMetaSelect {
	if m != nil {
		return m.Meta
	}
	return nil
}

func (m *DeleteCompaniesResponseSelect) GetSelectHash() bool {
	if m != nil && m.SelectHash != nil {
		return *m.SelectHash
	}
	return false
}

func (m *DeleteCompaniesResponseSelect) GetHash() string {
	if m != nil && m.Hash != nil {
		return *m.Hash
	}
	return ""
}

func init() {
	proto.RegisterType((*DeleteCompaniesResponseSelect)(nil), "deletecompaniesresponseselect.DeleteCompaniesResponseSelect")
}

func init() {
	proto.RegisterFile("deletecompaniesresponseselect/deletecompaniesresponseselect_.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 228 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0xd0, 0x31, 0x4b, 0xc4, 0x30,
	0x14, 0x07, 0x70, 0xa2, 0x1d, 0xbc, 0xb8, 0x65, 0x2a, 0xe2, 0x49, 0x71, 0x90, 0x0e, 0x9a, 0xe0,
	0x8d, 0x6e, 0x9e, 0x0e, 0x8a, 0xb8, 0xc4, 0xcd, 0x45, 0x62, 0xee, 0x91, 0x14, 0x92, 0xbe, 0x90,
	0xe4, 0x06, 0x3f, 0x94, 0xdf, 0x51, 0x2e, 0x6d, 0xf1, 0xa0, 0xd0, 0x2d, 0xf9, 0xbf, 0xff, 0xfb,
	0x05, 0x42, 0xb7, 0x3b, 0x70, 0x90, 0x41, 0xa3, 0x0f, 0xaa, 0xef, 0x20, 0x45, 0x48, 0x01, 0xfb,
	0x04, 0x09, 0x1c, 0xe8, 0x2c, 0x16, 0xa7, 0x5f, 0x3c, 0x44, 0xcc, 0xc8, 0xd6, 0x8b, 0xad, 0x8b,
	0xdb, 0xe9, 0xee, 0x21, 0xab, 0xd1, 0x9d, 0x47, 0x23, 0x76, 0xfd, 0x4b, 0xe8, 0xfa, 0xb9, 0x78,
	0x4f, 0x93, 0x27, 0xc7, 0xf2, 0x47, 0x29, 0xb2, 0x4b, 0xba, 0x1a, 0x56, 0x1e, 0x9d, 0xab, 0x49,
	0x43, 0xda, 0x33, 0xf9, 0x1f, 0xb0, 0x07, 0x5a, 0x1d, 0xd0, 0xfa, 0xa4, 0x21, 0xed, 0xf9, 0xe6,
	0x86, 0xcf, 0x5f, 0xe2, 0x93, 0xf7, 0x0e, 0x59, 0x0d, 0xa6, 0x2c, 0x3b, 0xec, 0x8a, 0xd2, 0xa1,
	0xf2, 0xa2, 0x92, 0xad, 0x4f, 0x0b, 0x7d, 0x94, 0x30, 0x46, 0x2b, 0x7b, 0x98, 0x54, 0x0d, 0x69,
	0x57, 0xb2, 0x9c, 0xb7, 0x6f, 0x9f, 0xaf, 0xa6, 0xcb, 0x76, 0xff, 0xcd, 0x35, 0x7a, 0xb1, 0xb9,
	0x4f, 0xb9, 0x43, 0x61, 0xf0, 0x2e, 0x38, 0xf5, 0x63, 0x22, 0xee, 0xfb, 0x9d, 0x30, 0x31, 0x68,
	0x91, 0xb4, 0x05, 0xaf, 0x96, 0x3f, 0xf4, 0x2f, 0x00, 0x00, 0xff, 0xff, 0x6b, 0xa5, 0xe6, 0x40,
	0x8e, 0x01, 0x00, 0x00,
}