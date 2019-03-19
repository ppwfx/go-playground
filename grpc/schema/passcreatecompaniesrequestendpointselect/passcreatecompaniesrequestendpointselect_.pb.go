// Code generated by protoc-gen-go. DO NOT EDIT.
// source: passcreatecompaniesrequestendpointselect/passcreatecompaniesrequestendpointselect_.proto

/*
Package passcreatecompaniesrequestendpointselect is a generated protocol buffer package.

It is generated from these files:
	passcreatecompaniesrequestendpointselect/passcreatecompaniesrequestendpointselect_.proto

It has these top-level messages:
	PassCreateCompaniesRequestEndpointSelect
*/
package passcreatecompaniesrequestendpointselect

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

type PassCreateCompaniesRequestEndpointSelect struct {
	SelectAll        *bool   `protobuf:"varint,1,opt,name=selectAll" json:"selectAll,omitempty"`
	SelectHash       *bool   `protobuf:"varint,2,opt,name=selectHash" json:"selectHash,omitempty"`
	Hash             *string `protobuf:"bytes,3,opt,name=hash" json:"hash,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *PassCreateCompaniesRequestEndpointSelect) Reset() {
	*m = PassCreateCompaniesRequestEndpointSelect{}
}
func (m *PassCreateCompaniesRequestEndpointSelect) String() string { return proto.CompactTextString(m) }
func (*PassCreateCompaniesRequestEndpointSelect) ProtoMessage()    {}
func (*PassCreateCompaniesRequestEndpointSelect) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{0}
}

func (m *PassCreateCompaniesRequestEndpointSelect) GetSelectAll() bool {
	if m != nil && m.SelectAll != nil {
		return *m.SelectAll
	}
	return false
}

func (m *PassCreateCompaniesRequestEndpointSelect) GetSelectHash() bool {
	if m != nil && m.SelectHash != nil {
		return *m.SelectHash
	}
	return false
}

func (m *PassCreateCompaniesRequestEndpointSelect) GetHash() string {
	if m != nil && m.Hash != nil {
		return *m.Hash
	}
	return ""
}

func init() {
	proto.RegisterType((*PassCreateCompaniesRequestEndpointSelect)(nil), "passcreatecompaniesrequestendpointselect.PassCreateCompaniesRequestEndpointSelect")
}

func init() {
	proto.RegisterFile("passcreatecompaniesrequestendpointselect/passcreatecompaniesrequestendpointselect_.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 201 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0xd0, 0xb1, 0x4b, 0xc4, 0x30,
	0x14, 0x06, 0x70, 0xaa, 0x0e, 0x36, 0x63, 0xa6, 0x0e, 0x22, 0xc5, 0x29, 0x8b, 0x0d, 0xfa, 0x1f,
	0x68, 0x11, 0x1c, 0xa5, 0x8a, 0x88, 0x8b, 0xc4, 0xf4, 0x91, 0x04, 0xd2, 0xbc, 0x5c, 0x5e, 0x3a,
	0x1c, 0xdc, 0x1f, 0x7f, 0x5c, 0x7a, 0xc7, 0xdd, 0xd8, 0x2d, 0xf9, 0xbe, 0xdf, 0xb7, 0x3c, 0xf6,
	0x13, 0x15, 0x91, 0x4e, 0xa0, 0x32, 0x68, 0x9c, 0xa2, 0x0a, 0x0e, 0x28, 0xc1, 0x66, 0x06, 0xca,
	0x10, 0xc6, 0x88, 0x2e, 0x64, 0x02, 0x0f, 0x3a, 0xcb, 0xb5, 0xf0, 0xaf, 0x8b, 0x09, 0x33, 0x72,
	0xb1, 0x76, 0xf0, 0xb0, 0x63, 0xe2, 0x43, 0x11, 0xf5, 0xc5, 0xf6, 0x27, 0x3b, 0x2c, 0xf6, 0xed,
	0x68, 0x3f, 0x8b, 0xe5, 0x77, 0xac, 0x5e, 0x56, 0x2f, 0xde, 0x37, 0x55, 0x5b, 0x89, 0xdb, 0xe1,
	0x1c, 0xf0, 0x7b, 0xc6, 0x96, 0xcf, 0xbb, 0x22, 0xdb, 0x5c, 0x95, 0xfa, 0x22, 0xe1, 0x9c, 0xdd,
	0xd8, 0x43, 0x73, 0xdd, 0x56, 0xa2, 0x1e, 0xca, 0xfb, 0xf5, 0xfb, 0xf7, 0xcb, 0xb8, 0x6c, 0xe7,
	0xff, 0x4e, 0xe3, 0x24, 0x9f, 0x9f, 0x28, 0x3b, 0x94, 0x06, 0x1f, 0xa3, 0x57, 0x5b, 0x93, 0x70,
	0x0e, 0xa3, 0x34, 0x29, 0x6a, 0x49, 0xda, 0xc2, 0xa4, 0x56, 0x9f, 0x61, 0x1f, 0x00, 0x00, 0xff,
	0xff, 0x9b, 0x1a, 0x1b, 0xf3, 0x5a, 0x01, 0x00, 0x00,
}
