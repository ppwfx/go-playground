// Code generated by protoc-gen-go. DO NOT EDIT.
// source: getlocationbyidendpointselect/getlocationbyidendpointselect_.proto

/*
Package getlocationbyidendpointselect is a generated protocol buffer package.

It is generated from these files:
	getlocationbyidendpointselect/getlocationbyidendpointselect_.proto

It has these top-level messages:
	GetLocationByIdEndpointSelect
*/
package getlocationbyidendpointselect

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

type GetLocationByIdEndpointSelect struct {
	SelectAll        *bool   `protobuf:"varint,1,opt,name=selectAll" json:"selectAll,omitempty"`
	SelectHash       *bool   `protobuf:"varint,2,opt,name=selectHash" json:"selectHash,omitempty"`
	Hash             *string `protobuf:"bytes,3,opt,name=hash" json:"hash,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *GetLocationByIdEndpointSelect) Reset()                    { *m = GetLocationByIdEndpointSelect{} }
func (m *GetLocationByIdEndpointSelect) String() string            { return proto.CompactTextString(m) }
func (*GetLocationByIdEndpointSelect) ProtoMessage()               {}
func (*GetLocationByIdEndpointSelect) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *GetLocationByIdEndpointSelect) GetSelectAll() bool {
	if m != nil && m.SelectAll != nil {
		return *m.SelectAll
	}
	return false
}

func (m *GetLocationByIdEndpointSelect) GetSelectHash() bool {
	if m != nil && m.SelectHash != nil {
		return *m.SelectHash
	}
	return false
}

func (m *GetLocationByIdEndpointSelect) GetHash() string {
	if m != nil && m.Hash != nil {
		return *m.Hash
	}
	return ""
}

func init() {
	proto.RegisterType((*GetLocationByIdEndpointSelect)(nil), "getlocationbyidendpointselect.GetLocationByIdEndpointSelect")
}

func init() {
	proto.RegisterFile("getlocationbyidendpointselect/getlocationbyidendpointselect_.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 187 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x72, 0x4a, 0x4f, 0x2d, 0xc9,
	0xc9, 0x4f, 0x4e, 0x2c, 0xc9, 0xcc, 0xcf, 0x4b, 0xaa, 0xcc, 0x4c, 0x49, 0xcd, 0x4b, 0x29, 0xc8,
	0xcf, 0xcc, 0x2b, 0x29, 0x4e, 0xcd, 0x49, 0x4d, 0x2e, 0xd1, 0xc7, 0x2b, 0x1b, 0xaf, 0x57, 0x50,
	0x94, 0x5f, 0x92, 0x2f, 0x24, 0x8b, 0x57, 0x95, 0x52, 0x21, 0x97, 0xac, 0x7b, 0x6a, 0x89, 0x0f,
	0x54, 0x81, 0x53, 0xa5, 0x67, 0x8a, 0x2b, 0x54, 0x41, 0x30, 0x58, 0x81, 0x90, 0x0c, 0x17, 0x27,
	0x44, 0xa9, 0x63, 0x4e, 0x8e, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0x47, 0x10, 0x42, 0x40, 0x48, 0x8e,
	0x8b, 0x0b, 0xc2, 0xf1, 0x48, 0x2c, 0xce, 0x90, 0x60, 0x02, 0x4b, 0x23, 0x89, 0x08, 0x09, 0x71,
	0xb1, 0x64, 0x80, 0x64, 0x98, 0x15, 0x18, 0x35, 0x38, 0x83, 0xc0, 0x6c, 0x27, 0xef, 0x28, 0xcf,
	0xf4, 0xcc, 0x92, 0x8c, 0xd2, 0x24, 0xbd, 0xe4, 0xfc, 0x5c, 0x7d, 0x23, 0xc3, 0xe2, 0x92, 0xcc,
	0x7c, 0xfd, 0xf4, 0x7c, 0xdd, 0x82, 0x9c, 0xc4, 0xca, 0xf4, 0xa2, 0xfc, 0xd2, 0xbc, 0x14, 0xfd,
	0xf4, 0xa2, 0x82, 0x64, 0xfd, 0xe2, 0xe4, 0x8c, 0xd4, 0xdc, 0x44, 0xfc, 0xbe, 0x04, 0x04, 0x00,
	0x00, 0xff, 0xff, 0x1e, 0xd9, 0x67, 0xaa, 0x23, 0x01, 0x00, 0x00,
}
