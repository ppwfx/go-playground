// Code generated by protoc-gen-go. DO NOT EDIT.
// source: verifytokenendpointselect/verifytokenendpointselect_.proto

/*
Package verifytokenendpointselect is a generated protocol buffer package.

It is generated from these files:
	verifytokenendpointselect/verifytokenendpointselect_.proto

It has these top-level messages:
	VerifyTokenEndpointSelect
*/
package verifytokenendpointselect

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

type VerifyTokenEndpointSelect struct {
	SelectAll        *bool   `protobuf:"varint,1,opt,name=selectAll" json:"selectAll,omitempty"`
	SelectHash       *bool   `protobuf:"varint,2,opt,name=selectHash" json:"selectHash,omitempty"`
	Hash             *string `protobuf:"bytes,3,opt,name=hash" json:"hash,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *VerifyTokenEndpointSelect) Reset()                    { *m = VerifyTokenEndpointSelect{} }
func (m *VerifyTokenEndpointSelect) String() string            { return proto.CompactTextString(m) }
func (*VerifyTokenEndpointSelect) ProtoMessage()               {}
func (*VerifyTokenEndpointSelect) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *VerifyTokenEndpointSelect) GetSelectAll() bool {
	if m != nil && m.SelectAll != nil {
		return *m.SelectAll
	}
	return false
}

func (m *VerifyTokenEndpointSelect) GetSelectHash() bool {
	if m != nil && m.SelectHash != nil {
		return *m.SelectHash
	}
	return false
}

func (m *VerifyTokenEndpointSelect) GetHash() string {
	if m != nil && m.Hash != nil {
		return *m.Hash
	}
	return ""
}

func init() {
	proto.RegisterType((*VerifyTokenEndpointSelect)(nil), "verifytokenendpointselect.VerifyTokenEndpointSelect")
}

func init() {
	proto.RegisterFile("verifytokenendpointselect/verifytokenendpointselect_.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 179 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xb2, 0x2a, 0x4b, 0x2d, 0xca,
	0x4c, 0xab, 0x2c, 0xc9, 0xcf, 0x4e, 0xcd, 0x4b, 0xcd, 0x4b, 0x29, 0xc8, 0xcf, 0xcc, 0x2b, 0x29,
	0x4e, 0xcd, 0x49, 0x4d, 0x2e, 0xd1, 0xc7, 0x29, 0x13, 0xaf, 0x57, 0x50, 0x94, 0x5f, 0x92, 0x2f,
	0x24, 0x89, 0x53, 0x85, 0x52, 0x2e, 0x97, 0x64, 0x18, 0x58, 0x32, 0x04, 0x24, 0xe9, 0x0a, 0x95,
	0x0c, 0x06, 0x4b, 0x0a, 0xc9, 0x70, 0x71, 0x42, 0x94, 0x39, 0xe6, 0xe4, 0x48, 0x30, 0x2a, 0x30,
	0x6a, 0x70, 0x04, 0x21, 0x04, 0x84, 0xe4, 0xb8, 0xb8, 0x20, 0x1c, 0x8f, 0xc4, 0xe2, 0x0c, 0x09,
	0x26, 0xb0, 0x34, 0x92, 0x88, 0x90, 0x10, 0x17, 0x4b, 0x06, 0x48, 0x86, 0x59, 0x81, 0x51, 0x83,
	0x33, 0x08, 0xcc, 0x76, 0x72, 0x8f, 0x72, 0x4d, 0xcf, 0x2c, 0xc9, 0x28, 0x4d, 0xd2, 0x4b, 0xce,
	0xcf, 0xd5, 0x37, 0x32, 0x2c, 0x2e, 0xc9, 0xcc, 0xd7, 0x4f, 0xcf, 0xd7, 0x2d, 0xc8, 0x49, 0xac,
	0x4c, 0x2f, 0xca, 0x2f, 0xcd, 0x4b, 0xd1, 0x4f, 0x2f, 0x2a, 0x48, 0xd6, 0x2f, 0x4e, 0xce, 0x48,
	0xcd, 0x4d, 0xc4, 0xed, 0x33, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x35, 0xf3, 0xf3, 0x39, 0x0f,
	0x01, 0x00, 0x00,
}
