// Code generated by protoc-gen-go. DO NOT EDIT.
// source: passcreatesocialaccountsrequestrequest/passcreatesocialaccountsrequestrequest_.proto

/*
Package passcreatesocialaccountsrequestrequest is a generated protocol buffer package.

It is generated from these files:
	passcreatesocialaccountsrequestrequest/passcreatesocialaccountsrequestrequest_.proto

It has these top-level messages:
	PassCreateSocialAccountsRequestRequest
*/
package passcreatesocialaccountsrequestrequest

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import requestmeta "github.com/21stio/go-playground/grpc/schema/requestmeta"
import servicefilter "github.com/21stio/go-playground/grpc/schema/servicefilter"
import createsocialaccountsrequest "github.com/21stio/go-playground/grpc/schema/createsocialaccountsrequest"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type PassCreateSocialAccountsRequestRequest struct {
	Meta                        *requestmeta.RequestMeta                                 `protobuf:"bytes,1,opt,name=meta" json:"meta,omitempty"`
	ServiceFilter               *servicefilter.ServiceFilter                             `protobuf:"bytes,2,opt,name=serviceFilter" json:"serviceFilter,omitempty"`
	CreateSocialAccountsRequest *createsocialaccountsrequest.CreateSocialAccountsRequest `protobuf:"bytes,3,opt,name=createSocialAccountsRequest" json:"createSocialAccountsRequest,omitempty"`
	Hash                        *string                                                  `protobuf:"bytes,4,opt,name=hash" json:"hash,omitempty"`
	XXX_unrecognized            []byte                                                   `json:"-"`
}

func (m *PassCreateSocialAccountsRequestRequest) Reset() {
	*m = PassCreateSocialAccountsRequestRequest{}
}
func (m *PassCreateSocialAccountsRequestRequest) String() string { return proto.CompactTextString(m) }
func (*PassCreateSocialAccountsRequestRequest) ProtoMessage()    {}
func (*PassCreateSocialAccountsRequestRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{0}
}

func (m *PassCreateSocialAccountsRequestRequest) GetMeta() *requestmeta.RequestMeta {
	if m != nil {
		return m.Meta
	}
	return nil
}

func (m *PassCreateSocialAccountsRequestRequest) GetServiceFilter() *servicefilter.ServiceFilter {
	if m != nil {
		return m.ServiceFilter
	}
	return nil
}

func (m *PassCreateSocialAccountsRequestRequest) GetCreateSocialAccountsRequest() *createsocialaccountsrequest.CreateSocialAccountsRequest {
	if m != nil {
		return m.CreateSocialAccountsRequest
	}
	return nil
}

func (m *PassCreateSocialAccountsRequestRequest) GetHash() string {
	if m != nil && m.Hash != nil {
		return *m.Hash
	}
	return ""
}

func init() {
	proto.RegisterType((*PassCreateSocialAccountsRequestRequest)(nil), "passcreatesocialaccountsrequestrequest.PassCreateSocialAccountsRequestRequest")
}

func init() {
	proto.RegisterFile("passcreatesocialaccountsrequestrequest/passcreatesocialaccountsrequestrequest_.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 270 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xb1, 0x4b, 0x03, 0x31,
	0x14, 0xc6, 0x69, 0xbd, 0xc5, 0x88, 0x4b, 0xa6, 0x50, 0x45, 0x4a, 0x87, 0xd2, 0x41, 0x13, 0xec,
	0xe4, 0x24, 0x58, 0xc1, 0x4d, 0x90, 0xb4, 0x93, 0x8b, 0x3c, 0xe3, 0xf3, 0x2e, 0x70, 0xd7, 0x9c,
	0x79, 0x39, 0x41, 0xff, 0x24, 0xff, 0x4a, 0x31, 0x46, 0xb8, 0x1b, 0x0c, 0x37, 0xbd, 0x17, 0xf2,
	0xfb, 0xbe, 0x7c, 0x1f, 0x84, 0xed, 0x5a, 0x20, 0x32, 0x1e, 0x21, 0x20, 0x39, 0x63, 0xa1, 0x06,
	0x63, 0x5c, 0xb7, 0x0f, 0xe4, 0xf1, 0xad, 0x43, 0x0a, 0x69, 0xa8, 0x71, 0xd8, 0x93, 0x6c, 0xbd,
	0x0b, 0x8e, 0x2f, 0xc7, 0xe1, 0xb3, 0xb3, 0xb4, 0x34, 0x18, 0x40, 0xf5, 0xf6, 0xe4, 0x33, 0x5b,
	0x10, 0xfa, 0x77, 0x6b, 0xf0, 0xd5, 0xd6, 0x01, 0xbd, 0x1a, 0x9c, 0xfe, 0x98, 0xeb, 0xcc, 0x3b,
	0x2a, 0x73, 0x97, 0xf4, 0x8b, 0xaf, 0x29, 0x5b, 0x3e, 0x00, 0xd1, 0x6d, 0x44, 0xb7, 0x11, 0xbd,
	0x49, 0xa8, 0xfe, 0x45, 0xd3, 0xe0, 0xe7, 0xac, 0xf8, 0x49, 0x27, 0x26, 0xf3, 0xc9, 0xea, 0x68,
	0x2d, 0x64, 0x2f, 0xb1, 0x4c, 0xcc, 0x3d, 0x06, 0xd0, 0x91, 0xe2, 0x1b, 0x76, 0x9c, 0x02, 0xdf,
	0xc5, 0xc0, 0x62, 0x1a, 0x65, 0xa7, 0x72, 0x50, 0x43, 0x6e, 0xfb, 0x8c, 0x1e, 0x4a, 0xf8, 0x27,
	0x3b, 0x31, 0xff, 0xe7, 0x12, 0x07, 0xd1, 0xf1, 0x4a, 0x66, 0x6a, 0xca, 0x5c, 0xaf, 0x9c, 0x39,
	0xe7, 0xac, 0xa8, 0x80, 0x2a, 0x51, 0xcc, 0x27, 0xab, 0x43, 0x1d, 0xf7, 0xcd, 0xee, 0x51, 0x97,
	0x36, 0x54, 0xdd, 0xb3, 0x34, 0xae, 0x51, 0xeb, 0x4b, 0x0a, 0xd6, 0xa9, 0xd2, 0x5d, 0xb4, 0x35,
	0x7c, 0x94, 0xde, 0x75, 0xfb, 0x17, 0x55, 0xfa, 0xd6, 0x28, 0x32, 0x15, 0x36, 0x30, 0xf2, 0xd7,
	0x7c, 0x07, 0x00, 0x00, 0xff, 0xff, 0xb1, 0x29, 0x2c, 0x1c, 0x85, 0x02, 0x00, 0x00,
}
