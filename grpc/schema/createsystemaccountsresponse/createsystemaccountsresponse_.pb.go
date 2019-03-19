// Code generated by protoc-gen-go. DO NOT EDIT.
// source: createsystemaccountsresponse/createsystemaccountsresponse_.proto

/*
Package createsystemaccountsresponse is a generated protocol buffer package.

It is generated from these files:
	createsystemaccountsresponse/createsystemaccountsresponse_.proto

It has these top-level messages:
	CreateSystemAccountsResponse
*/
package createsystemaccountsresponse

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import responsemeta "github.com/21stio/go-playground/grpc/schema/responsemeta"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type CreateSystemAccountsResponse struct {
	Meta             *responsemeta.ResponseMeta `protobuf:"bytes,1,opt,name=meta" json:"meta,omitempty"`
	Hash             *string                    `protobuf:"bytes,2,opt,name=hash" json:"hash,omitempty"`
	XXX_unrecognized []byte                     `json:"-"`
}

func (m *CreateSystemAccountsResponse) Reset()                    { *m = CreateSystemAccountsResponse{} }
func (m *CreateSystemAccountsResponse) String() string            { return proto.CompactTextString(m) }
func (*CreateSystemAccountsResponse) ProtoMessage()               {}
func (*CreateSystemAccountsResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *CreateSystemAccountsResponse) GetMeta() *responsemeta.ResponseMeta {
	if m != nil {
		return m.Meta
	}
	return nil
}

func (m *CreateSystemAccountsResponse) GetHash() string {
	if m != nil && m.Hash != nil {
		return *m.Hash
	}
	return ""
}

func init() {
	proto.RegisterType((*CreateSystemAccountsResponse)(nil), "createsystemaccountsresponse.CreateSystemAccountsResponse")
}

func init() {
	proto.RegisterFile("createsystemaccountsresponse/createsystemaccountsresponse_.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 190 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x8f, 0xbf, 0x6b, 0x85, 0x30,
	0x10, 0x80, 0xb1, 0xb8, 0x34, 0xdd, 0x32, 0x89, 0x38, 0x48, 0x27, 0x97, 0x26, 0xd4, 0xbf, 0xa0,
	0x3f, 0x96, 0x52, 0xe8, 0x62, 0xb7, 0x2e, 0x25, 0xe6, 0x1d, 0x89, 0xf0, 0xe2, 0x85, 0xdc, 0x39,
	0xf8, 0xdf, 0x3f, 0x5e, 0x50, 0x70, 0x72, 0xfb, 0x8e, 0xfb, 0xee, 0x83, 0x13, 0x6f, 0x36, 0x81,
	0x61, 0xa0, 0x95, 0x18, 0x82, 0xb1, 0x16, 0x97, 0x99, 0x29, 0x01, 0x45, 0x9c, 0x09, 0xf4, 0xd9,
	0xf2, 0x5f, 0xc5, 0x84, 0x8c, 0xb2, 0x39, 0x93, 0xea, 0x76, 0xa7, 0x00, 0x6c, 0xf4, 0x71, 0xd8,
	0xee, 0x9f, 0x47, 0xd1, 0x7c, 0xe6, 0xc2, 0x6f, 0x2e, 0xbc, 0x6f, 0x85, 0x61, 0x53, 0xa5, 0x12,
	0xe5, 0x5d, 0xaf, 0x8a, 0xb6, 0xe8, 0x9e, 0xfa, 0x5a, 0x1d, 0x1b, 0x6a, 0xb7, 0x7e, 0x80, 0xcd,
	0x90, 0x3d, 0x29, 0x45, 0xe9, 0x0d, 0xf9, 0xea, 0xa1, 0x2d, 0xba, 0xc7, 0x21, 0xf3, 0xc7, 0xf7,
	0xdf, 0x97, 0x9b, 0xd8, 0x2f, 0xa3, 0xb2, 0x18, 0x74, 0xff, 0x4a, 0x3c, 0xa1, 0x76, 0xf8, 0x12,
	0xaf, 0x66, 0x75, 0x09, 0x97, 0xf9, 0xa2, 0x5d, 0x8a, 0x56, 0x93, 0xf5, 0x10, 0xcc, 0xe9, 0xdb,
	0xb7, 0x00, 0x00, 0x00, 0xff, 0xff, 0xa0, 0xc3, 0xc1, 0x99, 0x32, 0x01, 0x00, 0x00,
}
