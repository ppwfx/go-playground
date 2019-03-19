// Code generated by protoc-gen-go. DO NOT EDIT.
// source: refreshtokenrequest/refreshtokenrequest_.proto

/*
Package refreshtokenrequest is a generated protocol buffer package.

It is generated from these files:
	refreshtokenrequest/refreshtokenrequest_.proto

It has these top-level messages:
	RefreshTokenRequest
*/
package refreshtokenrequest

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import requestmeta "github.com/21stio/go-playground/grpc/schema/requestmeta"
import servicefilter "github.com/21stio/go-playground/grpc/schema/servicefilter"
import refreshtokenresponseselect "github.com/21stio/go-playground/grpc/schema/refreshtokenresponseselect"
import token "github.com/21stio/go-playground/grpc/schema/token"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type RefreshTokenRequest struct {
	Meta             *requestmeta.RequestMeta                               `protobuf:"bytes,1,opt,name=meta" json:"meta,omitempty"`
	ServiceFilter    *servicefilter.ServiceFilter                           `protobuf:"bytes,2,opt,name=serviceFilter" json:"serviceFilter,omitempty"`
	Select           *refreshtokenresponseselect.RefreshTokenResponseSelect `protobuf:"bytes,3,opt,name=select" json:"select,omitempty"`
	Token            *token.Token                                           `protobuf:"bytes,4,opt,name=token" json:"token,omitempty"`
	Hash             *string                                                `protobuf:"bytes,5,opt,name=hash" json:"hash,omitempty"`
	XXX_unrecognized []byte                                                 `json:"-"`
}

func (m *RefreshTokenRequest) Reset()                    { *m = RefreshTokenRequest{} }
func (m *RefreshTokenRequest) String() string            { return proto.CompactTextString(m) }
func (*RefreshTokenRequest) ProtoMessage()               {}
func (*RefreshTokenRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *RefreshTokenRequest) GetMeta() *requestmeta.RequestMeta {
	if m != nil {
		return m.Meta
	}
	return nil
}

func (m *RefreshTokenRequest) GetServiceFilter() *servicefilter.ServiceFilter {
	if m != nil {
		return m.ServiceFilter
	}
	return nil
}

func (m *RefreshTokenRequest) GetSelect() *refreshtokenresponseselect.RefreshTokenResponseSelect {
	if m != nil {
		return m.Select
	}
	return nil
}

func (m *RefreshTokenRequest) GetToken() *token.Token {
	if m != nil {
		return m.Token
	}
	return nil
}

func (m *RefreshTokenRequest) GetHash() string {
	if m != nil && m.Hash != nil {
		return *m.Hash
	}
	return ""
}

func init() {
	proto.RegisterType((*RefreshTokenRequest)(nil), "refreshtokenrequest.RefreshTokenRequest")
}

func init() { proto.RegisterFile("refreshtokenrequest/refreshtokenrequest_.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 287 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x92, 0xc1, 0x4a, 0xc3, 0x40,
	0x10, 0x86, 0x49, 0x4d, 0x05, 0x57, 0xbd, 0x6c, 0x2f, 0x4b, 0x11, 0x29, 0x39, 0xf5, 0xa0, 0xb3,
	0xd8, 0x83, 0x27, 0x41, 0xec, 0xc1, 0x9b, 0x1e, 0xb6, 0x9e, 0xbc, 0x48, 0x8c, 0xd3, 0x24, 0x98,
	0x64, 0xe3, 0xee, 0x46, 0xf0, 0x29, 0x7c, 0x65, 0xe9, 0xec, 0x16, 0x12, 0x49, 0x2f, 0x61, 0x76,
	0xfe, 0xf9, 0x86, 0x6f, 0x20, 0x0c, 0x0c, 0x6e, 0x0d, 0xda, 0xc2, 0xe9, 0x4f, 0x6c, 0x0c, 0x7e,
	0x75, 0x68, 0x9d, 0x1c, 0xe9, 0xbd, 0x41, 0x6b, 0xb4, 0xd3, 0x7c, 0x36, 0x92, 0xcd, 0x2f, 0x43,
	0x51, 0xa3, 0x4b, 0x65, 0xaf, 0x0e, 0xd0, 0x3c, 0xb1, 0x68, 0xbe, 0xcb, 0x0c, 0xb7, 0x65, 0xe5,
	0xd0, 0xc8, 0xc1, 0x6b, 0x3f, 0x73, 0x37, 0x5c, 0x6c, 0x5b, 0xdd, 0x58, 0xb4, 0x58, 0x61, 0xf6,
	0xdf, 0xa7, 0x1f, 0xed, 0x69, 0x4e, 0x99, 0xa4, 0x6f, 0xe8, 0x25, 0xbf, 0x13, 0x36, 0x53, 0x9e,
	0x7c, 0xd9, 0xf5, 0x95, 0x17, 0xe3, 0x57, 0x2c, 0xde, 0xc9, 0x89, 0x68, 0x11, 0x2d, 0x4f, 0x57,
	0x02, 0x7a, 0xc2, 0x10, 0x66, 0x9e, 0xd0, 0xa5, 0x8a, 0xa6, 0xf8, 0x9a, 0x9d, 0x07, 0xdf, 0x47,
	0xf2, 0x15, 0x13, 0xc2, 0x2e, 0x60, 0x70, 0x05, 0x6c, 0xfa, 0x33, 0x6a, 0x88, 0xf0, 0x67, 0x76,
	0xec, 0x75, 0xc5, 0x11, 0xc1, 0xb7, 0x70, 0xf8, 0x22, 0x18, 0x2a, 0xfb, 0x68, 0x43, 0x91, 0x0a,
	0x5b, 0x78, 0xc2, 0xa6, 0x44, 0x8a, 0x98, 0xd6, 0x9d, 0x01, 0xbd, 0xc0, 0x23, 0x3e, 0xe2, 0x9c,
	0xc5, 0x45, 0x6a, 0x0b, 0x31, 0x5d, 0x44, 0xcb, 0x13, 0x45, 0xf5, 0xfa, 0xe1, 0xf5, 0x3e, 0x2f,
	0x5d, 0xd1, 0xbd, 0x43, 0xa6, 0x6b, 0xb9, 0xba, 0xb1, 0xae, 0xd4, 0x32, 0xd7, 0xd7, 0x6d, 0x95,
	0xfe, 0xe4, 0x46, 0x77, 0xcd, 0x87, 0xcc, 0x4d, 0x9b, 0x49, 0x9b, 0x15, 0x58, 0xa7, 0x63, 0xbf,
	0xc1, 0x5f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xdd, 0x12, 0x39, 0x22, 0x30, 0x02, 0x00, 0x00,
}
