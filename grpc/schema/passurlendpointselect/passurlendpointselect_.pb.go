// Code generated by protoc-gen-go. DO NOT EDIT.
// source: passurlendpointselect/passurlendpointselect_.proto

/*
Package passurlendpointselect is a generated protocol buffer package.

It is generated from these files:
	passurlendpointselect/passurlendpointselect_.proto

It has these top-level messages:
	PassUrlEndpointSelect
*/
package passurlendpointselect

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

type PassUrlEndpointSelect struct {
	SelectAll        *bool   `protobuf:"varint,1,opt,name=selectAll" json:"selectAll,omitempty"`
	SelectHash       *bool   `protobuf:"varint,2,opt,name=selectHash" json:"selectHash,omitempty"`
	Hash             *string `protobuf:"bytes,3,opt,name=hash" json:"hash,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *PassUrlEndpointSelect) Reset()                    { *m = PassUrlEndpointSelect{} }
func (m *PassUrlEndpointSelect) String() string            { return proto.CompactTextString(m) }
func (*PassUrlEndpointSelect) ProtoMessage()               {}
func (*PassUrlEndpointSelect) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *PassUrlEndpointSelect) GetSelectAll() bool {
	if m != nil && m.SelectAll != nil {
		return *m.SelectAll
	}
	return false
}

func (m *PassUrlEndpointSelect) GetSelectHash() bool {
	if m != nil && m.SelectHash != nil {
		return *m.SelectHash
	}
	return false
}

func (m *PassUrlEndpointSelect) GetHash() string {
	if m != nil && m.Hash != nil {
		return *m.Hash
	}
	return ""
}

func init() {
	proto.RegisterType((*PassUrlEndpointSelect)(nil), "passurlendpointselect.PassUrlEndpointSelect")
}

func init() { proto.RegisterFile("passurlendpointselect/passurlendpointselect_.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 176 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x32, 0x2a, 0x48, 0x2c, 0x2e,
	0x2e, 0x2d, 0xca, 0x49, 0xcd, 0x4b, 0x29, 0xc8, 0xcf, 0xcc, 0x2b, 0x29, 0x4e, 0xcd, 0x49, 0x4d,
	0x2e, 0xd1, 0xc7, 0x2a, 0x1a, 0xaf, 0x57, 0x50, 0x94, 0x5f, 0x92, 0x2f, 0x24, 0x8a, 0x55, 0x56,
	0x29, 0x93, 0x4b, 0x34, 0x20, 0xb1, 0xb8, 0x38, 0xb4, 0x28, 0xc7, 0x15, 0x2a, 0x11, 0x0c, 0x96,
	0x10, 0x92, 0xe1, 0xe2, 0x84, 0x28, 0x71, 0xcc, 0xc9, 0x91, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x08,
	0x42, 0x08, 0x08, 0xc9, 0x71, 0x71, 0x41, 0x38, 0x1e, 0x89, 0xc5, 0x19, 0x12, 0x4c, 0x60, 0x69,
	0x24, 0x11, 0x21, 0x21, 0x2e, 0x96, 0x0c, 0x90, 0x0c, 0xb3, 0x02, 0xa3, 0x06, 0x67, 0x10, 0x98,
	0xed, 0xe4, 0x1c, 0xe5, 0x98, 0x9e, 0x59, 0x92, 0x51, 0x9a, 0xa4, 0x97, 0x9c, 0x9f, 0xab, 0x6f,
	0x64, 0x58, 0x5c, 0x92, 0x99, 0xaf, 0x9f, 0x9e, 0xaf, 0x5b, 0x90, 0x93, 0x58, 0x99, 0x5e, 0x94,
	0x5f, 0x9a, 0x97, 0xa2, 0x9f, 0x5e, 0x54, 0x90, 0xac, 0x5f, 0x9c, 0x9c, 0x91, 0x9a, 0x9b, 0x88,
	0xdd, 0x37, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0xa1, 0x82, 0xb8, 0x46, 0xfb, 0x00, 0x00, 0x00,
}
