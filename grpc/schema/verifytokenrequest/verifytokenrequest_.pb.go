// Code generated by protoc-gen-go. DO NOT EDIT.
// source: verifytokenrequest/verifytokenrequest_.proto

/*
Package verifytokenrequest is a generated protocol buffer package.

It is generated from these files:
	verifytokenrequest/verifytokenrequest_.proto

It has these top-level messages:
	VerifyTokenRequest
*/
package verifytokenrequest

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import requestmeta "github.com/21stio/go-playground/grpc/schema/requestmeta"
import servicefilter "github.com/21stio/go-playground/grpc/schema/servicefilter"
import verifytokenresponseselect "github.com/21stio/go-playground/grpc/schema/verifytokenresponseselect"
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

type VerifyTokenRequest struct {
	Meta             *requestmeta.RequestMeta                             `protobuf:"bytes,1,opt,name=meta" json:"meta,omitempty"`
	ServiceFilter    *servicefilter.ServiceFilter                         `protobuf:"bytes,2,opt,name=serviceFilter" json:"serviceFilter,omitempty"`
	Select           *verifytokenresponseselect.VerifyTokenResponseSelect `protobuf:"bytes,3,opt,name=select" json:"select,omitempty"`
	Token            *token.Token                                         `protobuf:"bytes,4,opt,name=token" json:"token,omitempty"`
	Hash             *string                                              `protobuf:"bytes,5,opt,name=hash" json:"hash,omitempty"`
	XXX_unrecognized []byte                                               `json:"-"`
}

func (m *VerifyTokenRequest) Reset()                    { *m = VerifyTokenRequest{} }
func (m *VerifyTokenRequest) String() string            { return proto.CompactTextString(m) }
func (*VerifyTokenRequest) ProtoMessage()               {}
func (*VerifyTokenRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *VerifyTokenRequest) GetMeta() *requestmeta.RequestMeta {
	if m != nil {
		return m.Meta
	}
	return nil
}

func (m *VerifyTokenRequest) GetServiceFilter() *servicefilter.ServiceFilter {
	if m != nil {
		return m.ServiceFilter
	}
	return nil
}

func (m *VerifyTokenRequest) GetSelect() *verifytokenresponseselect.VerifyTokenResponseSelect {
	if m != nil {
		return m.Select
	}
	return nil
}

func (m *VerifyTokenRequest) GetToken() *token.Token {
	if m != nil {
		return m.Token
	}
	return nil
}

func (m *VerifyTokenRequest) GetHash() string {
	if m != nil && m.Hash != nil {
		return *m.Hash
	}
	return ""
}

func init() {
	proto.RegisterType((*VerifyTokenRequest)(nil), "verifytokenrequest.VerifyTokenRequest")
}

func init() { proto.RegisterFile("verifytokenrequest/verifytokenrequest_.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 286 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x91, 0xc1, 0x4a, 0xc3, 0x40,
	0x10, 0x86, 0x49, 0x4d, 0x05, 0x57, 0xbd, 0xcc, 0x69, 0x29, 0x22, 0x25, 0xa7, 0x1e, 0xea, 0x2c,
	0x16, 0x4f, 0x1e, 0x44, 0x7a, 0xf0, 0xa4, 0x97, 0x54, 0x3c, 0x78, 0x91, 0x18, 0xa7, 0x49, 0x30,
	0xc9, 0xc6, 0xdd, 0x4d, 0xa1, 0xef, 0xe0, 0x43, 0x4b, 0x27, 0x5b, 0x48, 0xac, 0x5e, 0xc2, 0xec,
	0xfc, 0xf3, 0x0d, 0xdf, 0x10, 0x31, 0xdf, 0x90, 0x29, 0xd6, 0x5b, 0xa7, 0x3f, 0xa9, 0x36, 0xf4,
	0xd5, 0x92, 0x75, 0xea, 0xb0, 0xf5, 0x86, 0x8d, 0xd1, 0x4e, 0x03, 0x1c, 0x46, 0x93, 0x4b, 0x5f,
	0x54, 0xe4, 0x12, 0xd5, 0xab, 0x3d, 0x33, 0x89, 0x2c, 0x99, 0x4d, 0x91, 0xd2, 0xba, 0x28, 0x1d,
	0x19, 0x35, 0x78, 0xed, 0x67, 0x6e, 0x07, 0x7b, 0x6d, 0xa3, 0x6b, 0x4b, 0x96, 0x4a, 0x4a, 0x7f,
	0xc9, 0xf4, 0x93, 0x3d, 0x0b, 0x9c, 0x29, 0xfe, 0xfa, 0x5e, 0xf4, 0x3d, 0x12, 0xf0, 0xc2, 0xe0,
	0xf3, 0xae, 0x1d, 0x77, 0x56, 0x30, 0x17, 0xe1, 0xce, 0x4c, 0x06, 0xd3, 0x60, 0x76, 0xba, 0x90,
	0xd8, 0xb3, 0x45, 0x3f, 0xf3, 0x44, 0x2e, 0x89, 0x79, 0x0a, 0x96, 0xe2, 0xdc, 0xcb, 0x3e, 0xb0,
	0xac, 0x1c, 0x31, 0x76, 0x81, 0x83, 0x13, 0x70, 0xd5, 0x9f, 0x89, 0x87, 0x08, 0x3c, 0x8a, 0xe3,
	0xce, 0x56, 0x1e, 0x31, 0x7c, 0x83, 0xff, 0xde, 0x83, 0x03, 0xe1, 0x2e, 0x59, 0x71, 0x12, 0xfb,
	0x1d, 0x10, 0x89, 0x31, 0x83, 0x32, 0xe4, 0x65, 0x67, 0xc8, 0x2f, 0xec, 0x90, 0x2e, 0x02, 0x10,
	0x61, 0x9e, 0xd8, 0x5c, 0x8e, 0xa7, 0xc1, 0xec, 0x24, 0xe6, 0x7a, 0x79, 0xff, 0x7a, 0x97, 0x15,
	0x2e, 0x6f, 0xdf, 0x31, 0xd5, 0x95, 0x5a, 0x5c, 0x5b, 0x57, 0x68, 0x95, 0xe9, 0xab, 0xa6, 0x4c,
	0xb6, 0x99, 0xd1, 0x6d, 0xfd, 0xa1, 0x32, 0xd3, 0xa4, 0xca, 0xa6, 0x39, 0x55, 0xc9, 0x1f, 0xff,
	0xff, 0x27, 0x00, 0x00, 0xff, 0xff, 0xc2, 0xca, 0x34, 0x13, 0x27, 0x02, 0x00, 0x00,
}
