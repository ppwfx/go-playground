// Code generated by protoc-gen-go. DO NOT EDIT.
// source: passcreatefeedsrequestendpointselect/passcreatefeedsrequestendpointselect_.proto

/*
Package passcreatefeedsrequestendpointselect is a generated protocol buffer package.

It is generated from these files:
	passcreatefeedsrequestendpointselect/passcreatefeedsrequestendpointselect_.proto

It has these top-level messages:
	PassCreateFeedsRequestEndpointSelect
*/
package passcreatefeedsrequestendpointselect

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

type PassCreateFeedsRequestEndpointSelect struct {
	SelectAll        *bool   `protobuf:"varint,1,opt,name=selectAll" json:"selectAll,omitempty"`
	SelectHash       *bool   `protobuf:"varint,2,opt,name=selectHash" json:"selectHash,omitempty"`
	Hash             *string `protobuf:"bytes,3,opt,name=hash" json:"hash,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *PassCreateFeedsRequestEndpointSelect) Reset()         { *m = PassCreateFeedsRequestEndpointSelect{} }
func (m *PassCreateFeedsRequestEndpointSelect) String() string { return proto.CompactTextString(m) }
func (*PassCreateFeedsRequestEndpointSelect) ProtoMessage()    {}
func (*PassCreateFeedsRequestEndpointSelect) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{0}
}

func (m *PassCreateFeedsRequestEndpointSelect) GetSelectAll() bool {
	if m != nil && m.SelectAll != nil {
		return *m.SelectAll
	}
	return false
}

func (m *PassCreateFeedsRequestEndpointSelect) GetSelectHash() bool {
	if m != nil && m.SelectHash != nil {
		return *m.SelectHash
	}
	return false
}

func (m *PassCreateFeedsRequestEndpointSelect) GetHash() string {
	if m != nil && m.Hash != nil {
		return *m.Hash
	}
	return ""
}

func init() {
	proto.RegisterType((*PassCreateFeedsRequestEndpointSelect)(nil), "passcreatefeedsrequestendpointselect.PassCreateFeedsRequestEndpointSelect")
}

func init() {
	proto.RegisterFile("passcreatefeedsrequestendpointselect/passcreatefeedsrequestendpointselect_.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 197 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0xcf, 0xb1, 0x6e, 0xc2, 0x30,
	0x10, 0x06, 0x60, 0xa5, 0xed, 0xd0, 0x78, 0xf4, 0x94, 0xa1, 0xaa, 0xa2, 0x2a, 0x43, 0x96, 0xc6,
	0x82, 0x37, 0x00, 0x04, 0x62, 0x8c, 0xcc, 0xc6, 0x82, 0x8c, 0x73, 0xd8, 0x91, 0x9c, 0xd8, 0xf8,
	0x1c, 0x09, 0xde, 0x1e, 0xc5, 0x41, 0x82, 0x31, 0xdb, 0xdd, 0xff, 0x7f, 0x37, 0x1c, 0xa9, 0x9d,
	0x40, 0x94, 0x1e, 0x44, 0x80, 0x0b, 0x40, 0x83, 0x1e, 0xae, 0x03, 0x60, 0x80, 0xbe, 0x71, 0xb6,
	0xed, 0x03, 0x82, 0x01, 0x19, 0xd8, 0x1c, 0x74, 0xaa, 0x9c, 0xb7, 0xc1, 0xd2, 0x62, 0x0e, 0xfe,
	0xbb, 0x91, 0xa2, 0x16, 0x88, 0x9b, 0xe8, 0x76, 0xa3, 0xe3, 0x93, 0xdb, 0x3e, 0xdd, 0x21, 0x3a,
	0xfa, 0x43, 0xd2, 0xe9, 0x62, 0x65, 0x4c, 0x96, 0xe4, 0x49, 0xf9, 0xcd, 0x5f, 0x01, 0xfd, 0x25,
	0x64, 0x5a, 0xf6, 0x02, 0x75, 0xf6, 0x11, 0xeb, 0xb7, 0x84, 0x52, 0xf2, 0xa5, 0xc7, 0xe6, 0x33,
	0x4f, 0xca, 0x94, 0xc7, 0x79, 0xcd, 0x8f, 0xb5, 0x6a, 0x83, 0x1e, 0xce, 0x95, 0xb4, 0x1d, 0x5b,
	0x2e, 0x30, 0xb4, 0x96, 0x29, 0xfb, 0xef, 0x8c, 0xb8, 0x2b, 0x6f, 0x87, 0xbe, 0x61, 0xca, 0x3b,
	0xc9, 0x50, 0x6a, 0xe8, 0xc4, 0xac, 0xd7, 0x1f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x0d, 0xe6, 0x67,
	0xc7, 0x46, 0x01, 0x00, 0x00,
}
