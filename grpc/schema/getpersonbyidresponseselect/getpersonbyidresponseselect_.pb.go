// Code generated by protoc-gen-go. DO NOT EDIT.
// source: getpersonbyidresponseselect/getpersonbyidresponseselect_.proto

/*
Package getpersonbyidresponseselect is a generated protocol buffer package.

It is generated from these files:
	getpersonbyidresponseselect/getpersonbyidresponseselect_.proto

It has these top-level messages:
	GetPersonByIdResponseSelect
*/
package getpersonbyidresponseselect

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import responsemetaselect "github.com/21stio/go-playground/grpc/schema/responsemetaselect"
import personselect "github.com/21stio/go-playground/grpc/schema/personselect"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type GetPersonByIdResponseSelect struct {
	SelectAll        *bool                                  `protobuf:"varint,1,opt,name=selectAll" json:"selectAll,omitempty"`
	Meta             *responsemetaselect.ResponseMetaSelect `protobuf:"bytes,2,opt,name=meta" json:"meta,omitempty"`
	Person           *personselect.PersonSelect             `protobuf:"bytes,3,opt,name=person" json:"person,omitempty"`
	SelectHash       *bool                                  `protobuf:"varint,4,opt,name=selectHash" json:"selectHash,omitempty"`
	Hash             *string                                `protobuf:"bytes,5,opt,name=hash" json:"hash,omitempty"`
	XXX_unrecognized []byte                                 `json:"-"`
}

func (m *GetPersonByIdResponseSelect) Reset()                    { *m = GetPersonByIdResponseSelect{} }
func (m *GetPersonByIdResponseSelect) String() string            { return proto.CompactTextString(m) }
func (*GetPersonByIdResponseSelect) ProtoMessage()               {}
func (*GetPersonByIdResponseSelect) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *GetPersonByIdResponseSelect) GetSelectAll() bool {
	if m != nil && m.SelectAll != nil {
		return *m.SelectAll
	}
	return false
}

func (m *GetPersonByIdResponseSelect) GetMeta() *responsemetaselect.ResponseMetaSelect {
	if m != nil {
		return m.Meta
	}
	return nil
}

func (m *GetPersonByIdResponseSelect) GetPerson() *personselect.PersonSelect {
	if m != nil {
		return m.Person
	}
	return nil
}

func (m *GetPersonByIdResponseSelect) GetSelectHash() bool {
	if m != nil && m.SelectHash != nil {
		return *m.SelectHash
	}
	return false
}

func (m *GetPersonByIdResponseSelect) GetHash() string {
	if m != nil && m.Hash != nil {
		return *m.Hash
	}
	return ""
}

func init() {
	proto.RegisterType((*GetPersonByIdResponseSelect)(nil), "getpersonbyidresponseselect.GetPersonByIdResponseSelect")
}

func init() {
	proto.RegisterFile("getpersonbyidresponseselect/getpersonbyidresponseselect_.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 257 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x91, 0x3f, 0x4b, 0xc4, 0x30,
	0x18, 0xc6, 0x89, 0x56, 0xf1, 0xe2, 0x96, 0xa9, 0xf4, 0x44, 0x8a, 0x83, 0x74, 0xd0, 0x04, 0x3b,
	0x3a, 0x08, 0xde, 0x72, 0xde, 0x20, 0x48, 0xdd, 0x5c, 0x24, 0xd7, 0xbe, 0xa4, 0x85, 0xb4, 0x09,
	0x49, 0x6e, 0xe8, 0xa7, 0xf5, 0xab, 0xc8, 0x25, 0x39, 0xae, 0x50, 0xe8, 0xd6, 0xf7, 0xf9, 0xf3,
	0xeb, 0x03, 0xc1, 0x6f, 0x02, 0x9c, 0x06, 0x63, 0xd5, 0xb0, 0x1f, 0xbb, 0xc6, 0x80, 0xd5, 0x6a,
	0xb0, 0x60, 0x41, 0x42, 0xed, 0xd8, 0x82, 0xf7, 0x4b, 0xb5, 0x51, 0x4e, 0x91, 0xf5, 0x42, 0x26,
	0x7b, 0x3a, 0xdd, 0x3d, 0x38, 0x1e, 0x99, 0x73, 0x29, 0xa2, 0xb2, 0x3c, 0x70, 0x62, 0x6e, 0x7a,
	0xc4, 0xc4, 0xc3, 0x1f, 0xc2, 0xeb, 0x2d, 0xb8, 0x2f, 0x6f, 0x6d, 0xc6, 0x5d, 0x53, 0x45, 0xd8,
	0xb7, 0x8f, 0x91, 0x3b, 0xbc, 0x0a, 0x85, 0x77, 0x29, 0x53, 0x94, 0xa3, 0xe2, 0xa6, 0x3a, 0x0b,
	0xe4, 0x15, 0x27, 0xc7, 0x9f, 0xa6, 0x17, 0x39, 0x2a, 0x6e, 0xcb, 0x47, 0x3a, 0x5f, 0x42, 0x4f,
	0xbc, 0x4f, 0x70, 0x3c, 0x30, 0x2b, 0xdf, 0x21, 0x25, 0xbe, 0x0e, 0x83, 0xd2, 0x4b, 0xdf, 0xce,
	0xe8, 0x74, 0x1f, 0x0d, 0x8b, 0x62, 0x23, 0x26, 0xc9, 0x3d, 0xc6, 0xc1, 0xfe, 0xe0, 0xb6, 0x4d,
	0x13, 0x3f, 0x67, 0xa2, 0x10, 0x82, 0x93, 0xf6, 0xe8, 0x5c, 0xe5, 0xa8, 0x58, 0x55, 0xfe, 0x7b,
	0xb3, 0xfb, 0xd9, 0x8a, 0xce, 0xb5, 0x87, 0x3d, 0xad, 0x55, 0xcf, 0xca, 0x17, 0xeb, 0x3a, 0xc5,
	0x84, 0x7a, 0xd6, 0x92, 0x8f, 0xc2, 0xa8, 0xc3, 0xd0, 0x30, 0x61, 0x74, 0xcd, 0x6c, 0xdd, 0x42,
	0xcf, 0x97, 0x1e, 0xe8, 0x3f, 0x00, 0x00, 0xff, 0xff, 0x41, 0x86, 0xa3, 0x16, 0xda, 0x01, 0x00,
	0x00,
}
