// Code generated by protoc-gen-go. DO NOT EDIT.
// source: passcreateservicesrequestrequest/passcreateservicesrequestrequest_.proto

/*
Package passcreateservicesrequestrequest is a generated protocol buffer package.

It is generated from these files:
	passcreateservicesrequestrequest/passcreateservicesrequestrequest_.proto

It has these top-level messages:
	PassCreateServicesRequestRequest
*/
package passcreateservicesrequestrequest

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import requestmeta "github.com/21stio/go-playground/grpc/schema/requestmeta"
import servicefilter "github.com/21stio/go-playground/grpc/schema/servicefilter"
import createservicesrequest "github.com/21stio/go-playground/grpc/schema/createservicesrequest"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type PassCreateServicesRequestRequest struct {
	Meta                  *requestmeta.RequestMeta                     `protobuf:"bytes,1,opt,name=meta" json:"meta,omitempty"`
	ServiceFilter         *servicefilter.ServiceFilter                 `protobuf:"bytes,2,opt,name=serviceFilter" json:"serviceFilter,omitempty"`
	CreateServicesRequest *createservicesrequest.CreateServicesRequest `protobuf:"bytes,3,opt,name=createServicesRequest" json:"createServicesRequest,omitempty"`
	Hash                  *string                                      `protobuf:"bytes,4,opt,name=hash" json:"hash,omitempty"`
	XXX_unrecognized      []byte                                       `json:"-"`
}

func (m *PassCreateServicesRequestRequest) Reset()         { *m = PassCreateServicesRequestRequest{} }
func (m *PassCreateServicesRequestRequest) String() string { return proto.CompactTextString(m) }
func (*PassCreateServicesRequestRequest) ProtoMessage()    {}
func (*PassCreateServicesRequestRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{0}
}

func (m *PassCreateServicesRequestRequest) GetMeta() *requestmeta.RequestMeta {
	if m != nil {
		return m.Meta
	}
	return nil
}

func (m *PassCreateServicesRequestRequest) GetServiceFilter() *servicefilter.ServiceFilter {
	if m != nil {
		return m.ServiceFilter
	}
	return nil
}

func (m *PassCreateServicesRequestRequest) GetCreateServicesRequest() *createservicesrequest.CreateServicesRequest {
	if m != nil {
		return m.CreateServicesRequest
	}
	return nil
}

func (m *PassCreateServicesRequestRequest) GetHash() string {
	if m != nil && m.Hash != nil {
		return *m.Hash
	}
	return ""
}

func init() {
	proto.RegisterType((*PassCreateServicesRequestRequest)(nil), "passcreateservicesrequestrequest.PassCreateServicesRequestRequest")
}

func init() {
	proto.RegisterFile("passcreateservicesrequestrequest/passcreateservicesrequestrequest_.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 257 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x91, 0x41, 0x4b, 0xc4, 0x30,
	0x10, 0x85, 0xe9, 0xda, 0x8b, 0x11, 0x2f, 0x01, 0x21, 0x2c, 0x22, 0x65, 0x4f, 0x7b, 0x58, 0x13,
	0xec, 0x4f, 0x58, 0x41, 0x3c, 0xa8, 0x48, 0xf7, 0xe6, 0x45, 0xb2, 0x71, 0x6c, 0x0b, 0x5b, 0x53,
	0x33, 0x53, 0xc1, 0xff, 0xe0, 0x8f, 0x16, 0xe3, 0x2c, 0xb4, 0x10, 0xe8, 0x69, 0x5e, 0xcb, 0xf7,
	0xde, 0xbc, 0x21, 0xe2, 0xbe, 0xb7, 0x88, 0x2e, 0x80, 0x25, 0x40, 0x08, 0x5f, 0xad, 0x03, 0x0c,
	0xf0, 0x39, 0x00, 0x12, 0x0f, 0x33, 0x07, 0xbc, 0xea, 0x3e, 0x78, 0xf2, 0xb2, 0x98, 0x03, 0x97,
	0x57, 0x2c, 0x3a, 0x20, 0x6b, 0x46, 0x9a, 0x13, 0x96, 0x2b, 0xf6, 0xbd, 0xb7, 0x07, 0x82, 0x60,
	0x26, 0x5f, 0x47, 0xa6, 0x4c, 0x6e, 0x30, 0xc9, 0xbf, 0xec, 0x59, 0xfd, 0x2c, 0x44, 0xf1, 0x6c,
	0x11, 0x6f, 0x23, 0xb4, 0x63, 0xa8, 0xfa, 0x87, 0x78, 0xc8, 0x8d, 0xc8, 0xff, 0xba, 0xa8, 0xac,
	0xc8, 0xd6, 0x67, 0xa5, 0xd2, 0xa3, 0x7e, 0x9a, 0x99, 0x47, 0x20, 0x5b, 0x45, 0x4a, 0x6e, 0xc5,
	0x39, 0x2f, 0xbb, 0x8b, 0xf5, 0xd4, 0x22, 0xda, 0x2e, 0xf5, 0xa4, 0xb4, 0xde, 0x8d, 0x99, 0x6a,
	0x6a, 0x91, 0x7b, 0x71, 0xe1, 0x52, 0x8d, 0xd4, 0x49, 0xcc, 0xda, 0xe8, 0xe4, 0x51, 0x3a, 0x7d,
	0x45, 0x3a, 0x4a, 0x4a, 0x91, 0x37, 0x16, 0x1b, 0x95, 0x17, 0xd9, 0xfa, 0xb4, 0x8a, 0x7a, 0xfb,
	0xf4, 0xf2, 0x50, 0xb7, 0xd4, 0x0c, 0x7b, 0xed, 0x7c, 0x67, 0xca, 0x1b, 0xa4, 0xd6, 0x9b, 0xda,
	0x5f, 0xf7, 0x07, 0xfb, 0x5d, 0x07, 0x3f, 0x7c, 0xbc, 0x99, 0x3a, 0xf4, 0xce, 0xa0, 0x6b, 0xa0,
	0xb3, 0xb3, 0xef, 0xff, 0x1b, 0x00, 0x00, 0xff, 0xff, 0x78, 0xb6, 0xaa, 0xcf, 0x43, 0x02, 0x00,
	0x00,
}
