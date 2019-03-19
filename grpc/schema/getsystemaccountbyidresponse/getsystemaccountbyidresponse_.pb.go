// Code generated by protoc-gen-go. DO NOT EDIT.
// source: getsystemaccountbyidresponse/getsystemaccountbyidresponse_.proto

/*
Package getsystemaccountbyidresponse is a generated protocol buffer package.

It is generated from these files:
	getsystemaccountbyidresponse/getsystemaccountbyidresponse_.proto

It has these top-level messages:
	GetSystemAccountByIdResponse
*/
package getsystemaccountbyidresponse

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import responsemeta "github.com/21stio/go-playground/grpc/schema/responsemeta"
import systemaccount "github.com/21stio/go-playground/grpc/schema/systemaccount"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type GetSystemAccountByIdResponse struct {
	Meta             *responsemeta.ResponseMeta   `protobuf:"bytes,1,opt,name=meta" json:"meta,omitempty"`
	SystemAccount    *systemaccount.SystemAccount `protobuf:"bytes,2,opt,name=systemAccount" json:"systemAccount,omitempty"`
	Hash             *string                      `protobuf:"bytes,3,opt,name=hash" json:"hash,omitempty"`
	XXX_unrecognized []byte                       `json:"-"`
}

func (m *GetSystemAccountByIdResponse) Reset()                    { *m = GetSystemAccountByIdResponse{} }
func (m *GetSystemAccountByIdResponse) String() string            { return proto.CompactTextString(m) }
func (*GetSystemAccountByIdResponse) ProtoMessage()               {}
func (*GetSystemAccountByIdResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *GetSystemAccountByIdResponse) GetMeta() *responsemeta.ResponseMeta {
	if m != nil {
		return m.Meta
	}
	return nil
}

func (m *GetSystemAccountByIdResponse) GetSystemAccount() *systemaccount.SystemAccount {
	if m != nil {
		return m.SystemAccount
	}
	return nil
}

func (m *GetSystemAccountByIdResponse) GetHash() string {
	if m != nil && m.Hash != nil {
		return *m.Hash
	}
	return ""
}

func init() {
	proto.RegisterType((*GetSystemAccountByIdResponse)(nil), "getsystemaccountbyidresponse.GetSystemAccountByIdResponse")
}

func init() {
	proto.RegisterFile("getsystemaccountbyidresponse/getsystemaccountbyidresponse_.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 225 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0xbd, 0x4e, 0xc4, 0x30,
	0x10, 0x84, 0x15, 0xb8, 0x06, 0x23, 0x1a, 0x57, 0x51, 0x94, 0x22, 0xba, 0xea, 0x1a, 0x6c, 0x71,
	0x4f, 0x00, 0x69, 0xf8, 0x91, 0x68, 0x42, 0x47, 0x83, 0x7c, 0xce, 0xca, 0x89, 0x44, 0xb2, 0x91,
	0x77, 0x53, 0xf8, 0x85, 0x78, 0x4e, 0x84, 0x49, 0xa4, 0xb8, 0x49, 0xb7, 0xa3, 0x99, 0x6f, 0x46,
	0x5a, 0xf1, 0xe8, 0x80, 0x29, 0x10, 0xc3, 0x60, 0xac, 0xc5, 0x79, 0xe4, 0x4b, 0xe8, 0x5b, 0x0f,
	0x34, 0xe1, 0x48, 0xa0, 0xf7, 0xcc, 0x2f, 0x35, 0x79, 0x64, 0x94, 0xe5, 0x5e, 0xa8, 0xa8, 0xd6,
	0x6b, 0x00, 0x36, 0x7a, 0x2b, 0x16, 0xbe, 0x38, 0x26, 0xb0, 0x4e, 0xd4, 0x92, 0x39, 0xfe, 0x64,
	0xa2, 0x7c, 0x06, 0xfe, 0x88, 0xde, 0xd3, 0xbf, 0x57, 0x87, 0xd7, 0xb6, 0x59, 0xfa, 0xa4, 0x12,
	0x87, 0xbf, 0xce, 0x3c, 0xab, 0xb2, 0xd3, 0xed, 0xb9, 0x50, 0xdb, 0x21, 0xb5, 0xa6, 0xde, 0x81,
	0x4d, 0x13, 0x73, 0xb2, 0x16, 0x77, 0xb4, 0x2d, 0xcb, 0xaf, 0x22, 0x58, 0xaa, 0x64, 0x5e, 0x25,
	0x83, 0x4d, 0x8a, 0x48, 0x29, 0x0e, 0x9d, 0xa1, 0x2e, 0xbf, 0xae, 0xb2, 0xd3, 0x4d, 0x13, 0xef,
	0xfa, 0xed, 0xf3, 0xc5, 0xf5, 0xdc, 0xcd, 0x17, 0x65, 0x71, 0xd0, 0xe7, 0x07, 0xe2, 0x1e, 0xb5,
	0xc3, 0xfb, 0xe9, 0xdb, 0x04, 0xe7, 0x71, 0x1e, 0x5b, 0xed, 0xfc, 0x64, 0x35, 0xd9, 0x0e, 0x06,
	0xb3, 0xfb, 0xdf, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x41, 0xc9, 0x7c, 0x3d, 0x9b, 0x01, 0x00,
	0x00,
}
