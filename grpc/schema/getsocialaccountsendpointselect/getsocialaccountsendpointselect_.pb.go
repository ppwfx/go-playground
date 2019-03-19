// Code generated by protoc-gen-go. DO NOT EDIT.
// source: getsocialaccountsendpointselect/getsocialaccountsendpointselect_.proto

/*
Package getsocialaccountsendpointselect is a generated protocol buffer package.

It is generated from these files:
	getsocialaccountsendpointselect/getsocialaccountsendpointselect_.proto

It has these top-level messages:
	GetSocialAccountsEndpointSelect
*/
package getsocialaccountsendpointselect

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

type GetSocialAccountsEndpointSelect struct {
	SelectAll        *bool   `protobuf:"varint,1,opt,name=selectAll" json:"selectAll,omitempty"`
	SelectHash       *bool   `protobuf:"varint,2,opt,name=selectHash" json:"selectHash,omitempty"`
	Hash             *string `protobuf:"bytes,3,opt,name=hash" json:"hash,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *GetSocialAccountsEndpointSelect) Reset()                    { *m = GetSocialAccountsEndpointSelect{} }
func (m *GetSocialAccountsEndpointSelect) String() string            { return proto.CompactTextString(m) }
func (*GetSocialAccountsEndpointSelect) ProtoMessage()               {}
func (*GetSocialAccountsEndpointSelect) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *GetSocialAccountsEndpointSelect) GetSelectAll() bool {
	if m != nil && m.SelectAll != nil {
		return *m.SelectAll
	}
	return false
}

func (m *GetSocialAccountsEndpointSelect) GetSelectHash() bool {
	if m != nil && m.SelectHash != nil {
		return *m.SelectHash
	}
	return false
}

func (m *GetSocialAccountsEndpointSelect) GetHash() string {
	if m != nil && m.Hash != nil {
		return *m.Hash
	}
	return ""
}

func init() {
	proto.RegisterType((*GetSocialAccountsEndpointSelect)(nil), "getsocialaccountsendpointselect.GetSocialAccountsEndpointSelect")
}

func init() {
	proto.RegisterFile("getsocialaccountsendpointselect/getsocialaccountsendpointselect_.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 186 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x72, 0x4b, 0x4f, 0x2d, 0x29,
	0xce, 0x4f, 0xce, 0x4c, 0xcc, 0x49, 0x4c, 0x4e, 0xce, 0x2f, 0xcd, 0x2b, 0x29, 0x4e, 0xcd, 0x4b,
	0x29, 0xc8, 0xcf, 0x04, 0x31, 0x72, 0x52, 0x93, 0x4b, 0xf4, 0x09, 0xc8, 0xc7, 0xeb, 0x15, 0x14,
	0xe5, 0x97, 0xe4, 0x0b, 0xc9, 0x13, 0x50, 0xa7, 0x54, 0xcc, 0x25, 0xef, 0x9e, 0x5a, 0x12, 0x0c,
	0x56, 0xe2, 0x08, 0x55, 0xe2, 0x0a, 0x55, 0x12, 0x0c, 0x56, 0x22, 0x24, 0xc3, 0xc5, 0x09, 0x51,
	0xec, 0x98, 0x93, 0x23, 0xc1, 0xa8, 0xc0, 0xa8, 0xc1, 0x11, 0x84, 0x10, 0x10, 0x92, 0xe3, 0xe2,
	0x82, 0x70, 0x3c, 0x12, 0x8b, 0x33, 0x24, 0x98, 0xc0, 0xd2, 0x48, 0x22, 0x42, 0x42, 0x5c, 0x2c,
	0x19, 0x20, 0x19, 0x66, 0x05, 0x46, 0x0d, 0xce, 0x20, 0x30, 0xdb, 0xc9, 0x37, 0xca, 0x3b, 0x3d,
	0xb3, 0x24, 0xa3, 0x34, 0x49, 0x2f, 0x39, 0x3f, 0x57, 0xdf, 0xc8, 0xb0, 0xb8, 0x24, 0x33, 0x5f,
	0x3f, 0x3d, 0x5f, 0xb7, 0x20, 0x27, 0xb1, 0x32, 0xbd, 0x28, 0xbf, 0x34, 0x2f, 0x45, 0x3f, 0xbd,
	0xa8, 0x20, 0x59, 0xbf, 0x38, 0x39, 0x23, 0x35, 0x37, 0x91, 0x90, 0x5f, 0x01, 0x01, 0x00, 0x00,
	0xff, 0xff, 0xa4, 0x1b, 0xc9, 0x11, 0x2d, 0x01, 0x00, 0x00,
}
