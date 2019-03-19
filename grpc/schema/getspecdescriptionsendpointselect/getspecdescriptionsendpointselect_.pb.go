// Code generated by protoc-gen-go. DO NOT EDIT.
// source: getspecdescriptionsendpointselect/getspecdescriptionsendpointselect_.proto

/*
Package getspecdescriptionsendpointselect is a generated protocol buffer package.

It is generated from these files:
	getspecdescriptionsendpointselect/getspecdescriptionsendpointselect_.proto

It has these top-level messages:
	GetSpecDescriptionsEndpointSelect
*/
package getspecdescriptionsendpointselect

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

type GetSpecDescriptionsEndpointSelect struct {
	SelectAll        *bool   `protobuf:"varint,1,opt,name=selectAll" json:"selectAll,omitempty"`
	SelectHash       *bool   `protobuf:"varint,2,opt,name=selectHash" json:"selectHash,omitempty"`
	Hash             *string `protobuf:"bytes,3,opt,name=hash" json:"hash,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *GetSpecDescriptionsEndpointSelect) Reset()         { *m = GetSpecDescriptionsEndpointSelect{} }
func (m *GetSpecDescriptionsEndpointSelect) String() string { return proto.CompactTextString(m) }
func (*GetSpecDescriptionsEndpointSelect) ProtoMessage()    {}
func (*GetSpecDescriptionsEndpointSelect) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{0}
}

func (m *GetSpecDescriptionsEndpointSelect) GetSelectAll() bool {
	if m != nil && m.SelectAll != nil {
		return *m.SelectAll
	}
	return false
}

func (m *GetSpecDescriptionsEndpointSelect) GetSelectHash() bool {
	if m != nil && m.SelectHash != nil {
		return *m.SelectHash
	}
	return false
}

func (m *GetSpecDescriptionsEndpointSelect) GetHash() string {
	if m != nil && m.Hash != nil {
		return *m.Hash
	}
	return ""
}

func init() {
	proto.RegisterType((*GetSpecDescriptionsEndpointSelect)(nil), "getspecdescriptionsendpointselect.GetSpecDescriptionsEndpointSelect")
}

func init() {
	proto.RegisterFile("getspecdescriptionsendpointselect/getspecdescriptionsendpointselect_.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 191 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xf2, 0x4a, 0x4f, 0x2d, 0x29,
	0x2e, 0x48, 0x4d, 0x4e, 0x49, 0x2d, 0x4e, 0x2e, 0xca, 0x2c, 0x28, 0xc9, 0xcc, 0xcf, 0x2b, 0x4e,
	0xcd, 0x4b, 0x29, 0xc8, 0xcf, 0xcc, 0x2b, 0x29, 0x4e, 0xcd, 0x49, 0x4d, 0x2e, 0xd1, 0x27, 0xa8,
	0x22, 0x5e, 0xaf, 0xa0, 0x28, 0xbf, 0x24, 0x5f, 0x48, 0x91, 0xa0, 0x4a, 0xa5, 0x52, 0x2e, 0x45,
	0xf7, 0xd4, 0x92, 0xe0, 0x82, 0xd4, 0x64, 0x17, 0x24, 0x45, 0xae, 0x50, 0x45, 0xc1, 0x60, 0x45,
	0x42, 0x32, 0x5c, 0x9c, 0x10, 0xe5, 0x8e, 0x39, 0x39, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x1c, 0x41,
	0x08, 0x01, 0x21, 0x39, 0x2e, 0x2e, 0x08, 0xc7, 0x23, 0xb1, 0x38, 0x43, 0x82, 0x09, 0x2c, 0x8d,
	0x24, 0x22, 0x24, 0xc4, 0xc5, 0x92, 0x01, 0x92, 0x61, 0x56, 0x60, 0xd4, 0xe0, 0x0c, 0x02, 0xb3,
	0x9d, 0xfc, 0xa3, 0x7c, 0xd3, 0x33, 0x4b, 0x32, 0x4a, 0x93, 0xf4, 0x92, 0xf3, 0x73, 0xf5, 0x8d,
	0x0c, 0x8b, 0x4b, 0x32, 0xf3, 0xf5, 0xd3, 0xf3, 0x75, 0x0b, 0x72, 0x12, 0x2b, 0xd3, 0x8b, 0xf2,
	0x4b, 0xf3, 0x52, 0xf4, 0xd3, 0x8b, 0x0a, 0x92, 0xf5, 0x8b, 0x93, 0x33, 0x52, 0x73, 0x13, 0x09,
	0xfb, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0xdc, 0xdb, 0xf5, 0x4d, 0x37, 0x01, 0x00, 0x00,
}
