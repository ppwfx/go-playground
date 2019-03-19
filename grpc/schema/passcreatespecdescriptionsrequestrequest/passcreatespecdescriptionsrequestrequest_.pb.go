// Code generated by protoc-gen-go. DO NOT EDIT.
// source: passcreatespecdescriptionsrequestrequest/passcreatespecdescriptionsrequestrequest_.proto

/*
Package passcreatespecdescriptionsrequestrequest is a generated protocol buffer package.

It is generated from these files:
	passcreatespecdescriptionsrequestrequest/passcreatespecdescriptionsrequestrequest_.proto

It has these top-level messages:
	PassCreateSpecDescriptionsRequestRequest
*/
package passcreatespecdescriptionsrequestrequest

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import requestmeta "github.com/21stio/go-playground/grpc/schema/requestmeta"
import servicefilter "github.com/21stio/go-playground/grpc/schema/servicefilter"
import createspecdescriptionsrequest "github.com/21stio/go-playground/grpc/schema/createspecdescriptionsrequest"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type PassCreateSpecDescriptionsRequestRequest struct {
	Meta                          *requestmeta.RequestMeta                                     `protobuf:"bytes,1,opt,name=meta" json:"meta,omitempty"`
	ServiceFilter                 *servicefilter.ServiceFilter                                 `protobuf:"bytes,2,opt,name=serviceFilter" json:"serviceFilter,omitempty"`
	CreateSpecDescriptionsRequest *createspecdescriptionsrequest.CreateSpecDescriptionsRequest `protobuf:"bytes,3,opt,name=createSpecDescriptionsRequest" json:"createSpecDescriptionsRequest,omitempty"`
	Hash                          *string                                                      `protobuf:"bytes,4,opt,name=hash" json:"hash,omitempty"`
	XXX_unrecognized              []byte                                                       `json:"-"`
}

func (m *PassCreateSpecDescriptionsRequestRequest) Reset() {
	*m = PassCreateSpecDescriptionsRequestRequest{}
}
func (m *PassCreateSpecDescriptionsRequestRequest) String() string { return proto.CompactTextString(m) }
func (*PassCreateSpecDescriptionsRequestRequest) ProtoMessage()    {}
func (*PassCreateSpecDescriptionsRequestRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{0}
}

func (m *PassCreateSpecDescriptionsRequestRequest) GetMeta() *requestmeta.RequestMeta {
	if m != nil {
		return m.Meta
	}
	return nil
}

func (m *PassCreateSpecDescriptionsRequestRequest) GetServiceFilter() *servicefilter.ServiceFilter {
	if m != nil {
		return m.ServiceFilter
	}
	return nil
}

func (m *PassCreateSpecDescriptionsRequestRequest) GetCreateSpecDescriptionsRequest() *createspecdescriptionsrequest.CreateSpecDescriptionsRequest {
	if m != nil {
		return m.CreateSpecDescriptionsRequest
	}
	return nil
}

func (m *PassCreateSpecDescriptionsRequestRequest) GetHash() string {
	if m != nil && m.Hash != nil {
		return *m.Hash
	}
	return ""
}

func init() {
	proto.RegisterType((*PassCreateSpecDescriptionsRequestRequest)(nil), "passcreatespecdescriptionsrequestrequest.PassCreateSpecDescriptionsRequestRequest")
}

func init() {
	proto.RegisterFile("passcreatespecdescriptionsrequestrequest/passcreatespecdescriptionsrequestrequest_.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 272 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0x41, 0x4b, 0xc4, 0x30,
	0x10, 0x85, 0xe9, 0xba, 0x17, 0x23, 0x5e, 0x72, 0x0a, 0x8b, 0xca, 0xb2, 0xa7, 0x1e, 0x34, 0xc1,
	0x3d, 0x7b, 0xaa, 0xe2, 0x4d, 0x90, 0xae, 0x88, 0x78, 0x91, 0x98, 0x1d, 0xdb, 0xc0, 0x76, 0x13,
	0x33, 0x53, 0xc1, 0xab, 0xff, 0xc9, 0xff, 0x27, 0xc6, 0x08, 0xed, 0x25, 0xf4, 0x34, 0x53, 0xfa,
	0xbd, 0x37, 0xef, 0x41, 0xd8, 0x93, 0xd7, 0x88, 0x26, 0x80, 0x26, 0x40, 0x0f, 0x66, 0x0b, 0x68,
	0x82, 0xf5, 0x64, 0xdd, 0x1e, 0x03, 0xbc, 0xf7, 0x80, 0x94, 0x86, 0x9a, 0x0a, 0xbe, 0x48, 0x1f,
	0x1c, 0x39, 0x5e, 0x4e, 0x15, 0x2c, 0xce, 0xd2, 0xd2, 0x01, 0x69, 0x35, 0xd8, 0x93, 0xd3, 0x62,
	0x85, 0x10, 0x3e, 0xac, 0x81, 0x37, 0xbb, 0x23, 0x08, 0x6a, 0xf4, 0xf5, 0xcf, 0x54, 0xd9, 0x4b,
	0x2a, 0xfb, 0x37, 0x79, 0xac, 0xbe, 0x67, 0xac, 0xbc, 0xd7, 0x88, 0xd7, 0x11, 0xde, 0x78, 0x30,
	0x37, 0x03, 0xb8, 0xfe, 0x83, 0xd3, 0xe0, 0xe7, 0x6c, 0xfe, 0x9b, 0x51, 0x14, 0xcb, 0xa2, 0x3c,
	0x5a, 0x0b, 0x39, 0xc8, 0x2d, 0x13, 0x73, 0x07, 0xa4, 0xeb, 0x48, 0xf1, 0x8a, 0x1d, 0xa7, 0xd8,
	0xb7, 0x31, 0xb6, 0x98, 0x45, 0xd9, 0x89, 0x1c, 0x95, 0x91, 0x9b, 0x21, 0x53, 0x8f, 0x25, 0xfc,
	0xab, 0x60, 0xa7, 0x26, 0x17, 0x4d, 0x1c, 0x44, 0xd3, 0x2b, 0x99, 0x6d, 0x2b, 0xf3, 0xf5, 0xf2,
	0x27, 0x38, 0x67, 0xf3, 0x56, 0x63, 0x2b, 0xe6, 0xcb, 0xa2, 0x3c, 0xac, 0xe3, 0x5e, 0x3d, 0x3e,
	0x3f, 0x34, 0x96, 0xda, 0xfe, 0x55, 0x1a, 0xd7, 0xa9, 0xf5, 0x25, 0x92, 0x75, 0xaa, 0x71, 0x17,
	0x7e, 0xa7, 0x3f, 0x9b, 0xe0, 0xfa, 0xfd, 0x56, 0x35, 0xc1, 0x1b, 0x85, 0xa6, 0x85, 0x4e, 0x4f,
	0x7e, 0x48, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xd7, 0xc1, 0xd7, 0xed, 0x9c, 0x02, 0x00, 0x00,
}