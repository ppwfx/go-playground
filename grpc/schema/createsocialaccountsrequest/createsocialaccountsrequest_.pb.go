// Code generated by protoc-gen-go. DO NOT EDIT.
// source: createsocialaccountsrequest/createsocialaccountsrequest_.proto

/*
Package createsocialaccountsrequest is a generated protocol buffer package.

It is generated from these files:
	createsocialaccountsrequest/createsocialaccountsrequest_.proto

It has these top-level messages:
	CreateSocialAccountsRequest
*/
package createsocialaccountsrequest

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import requestmeta "github.com/21stio/go-playground/grpc/schema/requestmeta"
import servicefilter "github.com/21stio/go-playground/grpc/schema/servicefilter"
import feed "github.com/21stio/go-playground/grpc/schema/feed"
import createsocialaccountsresponseselect "github.com/21stio/go-playground/grpc/schema/createsocialaccountsresponseselect"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type CreateSocialAccountsRequest struct {
	Meta             *requestmeta.RequestMeta                                               `protobuf:"bytes,1,opt,name=meta" json:"meta,omitempty"`
	ServiceFilter    *servicefilter.ServiceFilter                                           `protobuf:"bytes,2,opt,name=serviceFilter" json:"serviceFilter,omitempty"`
	SocialAccounts   []*feed.SocialAccount                                                  `protobuf:"bytes,3,rep,name=socialAccounts" json:"socialAccounts,omitempty"`
	Select           *createsocialaccountsresponseselect.CreateSocialAccountsResponseSelect `protobuf:"bytes,4,opt,name=select" json:"select,omitempty"`
	Hash             *string                                                                `protobuf:"bytes,5,opt,name=hash" json:"hash,omitempty"`
	XXX_unrecognized []byte                                                                 `json:"-"`
}

func (m *CreateSocialAccountsRequest) Reset()                    { *m = CreateSocialAccountsRequest{} }
func (m *CreateSocialAccountsRequest) String() string            { return proto.CompactTextString(m) }
func (*CreateSocialAccountsRequest) ProtoMessage()               {}
func (*CreateSocialAccountsRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *CreateSocialAccountsRequest) GetMeta() *requestmeta.RequestMeta {
	if m != nil {
		return m.Meta
	}
	return nil
}

func (m *CreateSocialAccountsRequest) GetServiceFilter() *servicefilter.ServiceFilter {
	if m != nil {
		return m.ServiceFilter
	}
	return nil
}

func (m *CreateSocialAccountsRequest) GetSocialAccounts() []*feed.SocialAccount {
	if m != nil {
		return m.SocialAccounts
	}
	return nil
}

func (m *CreateSocialAccountsRequest) GetSelect() *createsocialaccountsresponseselect.CreateSocialAccountsResponseSelect {
	if m != nil {
		return m.Select
	}
	return nil
}

func (m *CreateSocialAccountsRequest) GetHash() string {
	if m != nil && m.Hash != nil {
		return *m.Hash
	}
	return ""
}

func init() {
	proto.RegisterType((*CreateSocialAccountsRequest)(nil), "createsocialaccountsrequest.CreateSocialAccountsRequest")
}

func init() {
	proto.RegisterFile("createsocialaccountsrequest/createsocialaccountsrequest_.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 306 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0x4f, 0x4b, 0xc4, 0x30,
	0x10, 0xc5, 0xd9, 0x3f, 0x0a, 0x66, 0x51, 0x24, 0x5e, 0xc2, 0xae, 0x48, 0xd9, 0x53, 0x0f, 0x9a,
	0x60, 0xaf, 0x82, 0xe0, 0x0a, 0x2b, 0x82, 0x5e, 0xd2, 0x9b, 0x07, 0x25, 0x66, 0x67, 0xdb, 0x42,
	0xdb, 0xd4, 0x24, 0x15, 0xfc, 0x76, 0x7e, 0x34, 0x69, 0x9a, 0x85, 0x56, 0xd6, 0x7a, 0x29, 0xd3,
	0xe4, 0xf7, 0x66, 0xde, 0x1b, 0x82, 0x6e, 0xa5, 0x06, 0x61, 0xc1, 0x28, 0x99, 0x89, 0x5c, 0x48,
	0xa9, 0xea, 0xd2, 0x1a, 0x0d, 0x1f, 0x35, 0x18, 0xcb, 0x06, 0xee, 0xde, 0x68, 0xa5, 0x95, 0x55,
	0x78, 0x31, 0xc0, 0xcc, 0x2f, 0x7c, 0x51, 0x80, 0x15, 0xac, 0x53, 0x7b, 0xf1, 0x7c, 0x69, 0x40,
	0x7f, 0x66, 0x12, 0xb6, 0x59, 0x6e, 0x41, 0xb3, 0xde, 0xdf, 0x8e, 0x39, 0xdd, 0x02, 0x6c, 0x58,
	0xf3, 0xd9, 0x9d, 0x3c, 0xed, 0x1f, 0x69, 0x2a, 0x55, 0x1a, 0x30, 0x90, 0x83, 0xfc, 0xcb, 0x79,
	0x17, 0xf1, 0xdd, 0x96, 0xdf, 0x63, 0xb4, 0xb8, 0x77, 0x74, 0xec, 0xe8, 0x3b, 0x4f, 0xf3, 0xd6,
	0x2e, 0xbe, 0x44, 0xd3, 0xc6, 0x32, 0x19, 0x05, 0xa3, 0x70, 0x16, 0x11, 0xda, 0x89, 0x41, 0x3d,
	0xf3, 0x0c, 0x56, 0x70, 0x47, 0xe1, 0x15, 0x3a, 0xf6, 0x29, 0xd6, 0x2e, 0x05, 0x19, 0x3b, 0xd9,
	0x39, 0xed, 0x65, 0xa3, 0x71, 0x97, 0xe1, 0x7d, 0x09, 0xbe, 0x41, 0x27, 0xa6, 0x67, 0x85, 0x4c,
	0x82, 0x49, 0x38, 0x8b, 0xce, 0x68, 0xb3, 0x05, 0xda, 0xb3, 0xc9, 0x7f, 0xa1, 0xf8, 0x15, 0x1d,
	0xb6, 0xf9, 0xc8, 0xd4, 0x4d, 0x5e, 0xd3, 0xff, 0x57, 0x41, 0xf7, 0xe7, 0x6f, 0x91, 0xd8, 0x21,
	0xdc, 0x77, 0xc5, 0x18, 0x4d, 0x53, 0x61, 0x52, 0x72, 0x10, 0x8c, 0xc2, 0x23, 0xee, 0xea, 0xd5,
	0xe3, 0xcb, 0x43, 0x92, 0xd9, 0xb4, 0x7e, 0xa7, 0x52, 0x15, 0x2c, 0xba, 0x36, 0x36, 0x53, 0x2c,
	0x51, 0x57, 0x55, 0x2e, 0xbe, 0x12, 0xad, 0xea, 0x72, 0xc3, 0x12, 0x5d, 0x49, 0x66, 0x64, 0x0a,
	0x85, 0x18, 0x7a, 0x55, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xf4, 0xf4, 0x39, 0x52, 0x8f, 0x02,
	0x00, 0x00,
}
