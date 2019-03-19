// Code generated by protoc-gen-go. DO NOT EDIT.
// source: passdeletesystemaccountsrequestendpointselect/passdeletesystemaccountsrequestendpointselect_.proto

/*
Package passdeletesystemaccountsrequestendpointselect is a generated protocol buffer package.

It is generated from these files:
	passdeletesystemaccountsrequestendpointselect/passdeletesystemaccountsrequestendpointselect_.proto

It has these top-level messages:
	PassDeleteSystemAccountsRequestEndpointSelect
*/
package passdeletesystemaccountsrequestendpointselect

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

type PassDeleteSystemAccountsRequestEndpointSelect struct {
	SelectAll        *bool   `protobuf:"varint,1,opt,name=selectAll" json:"selectAll,omitempty"`
	SelectHash       *bool   `protobuf:"varint,2,opt,name=selectHash" json:"selectHash,omitempty"`
	Hash             *string `protobuf:"bytes,3,opt,name=hash" json:"hash,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *PassDeleteSystemAccountsRequestEndpointSelect) Reset() {
	*m = PassDeleteSystemAccountsRequestEndpointSelect{}
}
func (m *PassDeleteSystemAccountsRequestEndpointSelect) String() string {
	return proto.CompactTextString(m)
}
func (*PassDeleteSystemAccountsRequestEndpointSelect) ProtoMessage() {}
func (*PassDeleteSystemAccountsRequestEndpointSelect) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{0}
}

func (m *PassDeleteSystemAccountsRequestEndpointSelect) GetSelectAll() bool {
	if m != nil && m.SelectAll != nil {
		return *m.SelectAll
	}
	return false
}

func (m *PassDeleteSystemAccountsRequestEndpointSelect) GetSelectHash() bool {
	if m != nil && m.SelectHash != nil {
		return *m.SelectHash
	}
	return false
}

func (m *PassDeleteSystemAccountsRequestEndpointSelect) GetHash() string {
	if m != nil && m.Hash != nil {
		return *m.Hash
	}
	return ""
}

func init() {
	proto.RegisterType((*PassDeleteSystemAccountsRequestEndpointSelect)(nil), "passdeletesystemaccountsrequestendpointselect.PassDeleteSystemAccountsRequestEndpointSelect")
}

func init() {
	proto.RegisterFile("passdeletesystemaccountsrequestendpointselect/passdeletesystemaccountsrequestendpointselect_.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 206 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0xd0, 0xb1, 0x4e, 0xc4, 0x30,
	0x0c, 0x06, 0x60, 0x05, 0x18, 0xb8, 0x8c, 0x99, 0x6e, 0x40, 0xe8, 0xc4, 0x74, 0x4b, 0x2f, 0x82,
	0x37, 0x38, 0x04, 0x12, 0x23, 0x6a, 0x27, 0x60, 0x40, 0x69, 0x6a, 0x25, 0x95, 0xd2, 0x38, 0xc4,
	0xce, 0xd0, 0x91, 0x37, 0x47, 0xa4, 0x95, 0x60, 0xed, 0x66, 0xff, 0xbf, 0xbf, 0xc5, 0xb2, 0x4f,
	0x86, 0x68, 0x80, 0x00, 0x0c, 0x34, 0x13, 0xc3, 0x64, 0xac, 0xc5, 0x12, 0x99, 0x32, 0x7c, 0x15,
	0x20, 0x86, 0x38, 0x24, 0x1c, 0x23, 0x13, 0x04, 0xb0, 0xac, 0x37, 0x5d, 0x7f, 0x9e, 0x52, 0x46,
	0x46, 0xd5, 0x6c, 0x52, 0x77, 0xdf, 0x42, 0x36, 0xaf, 0x86, 0xe8, 0xa9, 0x8a, 0xae, 0x8a, 0xf3,
	0x2a, 0xda, 0x45, 0x3c, 0xaf, 0xa2, 0xab, 0x42, 0xdd, 0xc8, 0xdd, 0x62, 0xcf, 0x21, 0xec, 0xc5,
	0x41, 0x1c, 0xaf, 0xdb, 0xbf, 0x40, 0xdd, 0x4a, 0xb9, 0x2c, 0x2f, 0x86, 0xfc, 0xfe, 0xa2, 0xd6,
	0xff, 0x12, 0xa5, 0xe4, 0x95, 0xff, 0x6d, 0x2e, 0x0f, 0xe2, 0xb8, 0x6b, 0xeb, 0xfc, 0xf8, 0xf1,
	0xfe, 0xe6, 0x46, 0xf6, 0xa5, 0x3f, 0x59, 0x9c, 0xf4, 0xc3, 0x3d, 0xf1, 0x88, 0xda, 0x61, 0x93,
	0x82, 0x99, 0x5d, 0xc6, 0x12, 0x07, 0xed, 0x72, 0xb2, 0x9a, 0xac, 0x87, 0xc9, 0x6c, 0x7b, 0xcb,
	0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x69, 0x42, 0x54, 0xe7, 0x74, 0x01, 0x00, 0x00,
}
