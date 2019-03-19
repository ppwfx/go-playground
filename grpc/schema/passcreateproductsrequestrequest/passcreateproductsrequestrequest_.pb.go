// Code generated by protoc-gen-go. DO NOT EDIT.
// source: passcreateproductsrequestrequest/passcreateproductsrequestrequest_.proto

/*
Package passcreateproductsrequestrequest is a generated protocol buffer package.

It is generated from these files:
	passcreateproductsrequestrequest/passcreateproductsrequestrequest_.proto

It has these top-level messages:
	PassCreateProductsRequestRequest
*/
package passcreateproductsrequestrequest

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import requestmeta "github.com/21stio/go-playground/grpc/schema/requestmeta"
import servicefilter "github.com/21stio/go-playground/grpc/schema/servicefilter"
import createproductsrequest "github.com/21stio/go-playground/grpc/schema/createproductsrequest"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type PassCreateProductsRequestRequest struct {
	Meta                  *requestmeta.RequestMeta                     `protobuf:"bytes,1,opt,name=meta" json:"meta,omitempty"`
	ServiceFilter         *servicefilter.ServiceFilter                 `protobuf:"bytes,2,opt,name=serviceFilter" json:"serviceFilter,omitempty"`
	CreateProductsRequest *createproductsrequest.CreateProductsRequest `protobuf:"bytes,3,opt,name=createProductsRequest" json:"createProductsRequest,omitempty"`
	Hash                  *string                                      `protobuf:"bytes,4,opt,name=hash" json:"hash,omitempty"`
	XXX_unrecognized      []byte                                       `json:"-"`
}

func (m *PassCreateProductsRequestRequest) Reset()         { *m = PassCreateProductsRequestRequest{} }
func (m *PassCreateProductsRequestRequest) String() string { return proto.CompactTextString(m) }
func (*PassCreateProductsRequestRequest) ProtoMessage()    {}
func (*PassCreateProductsRequestRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{0}
}

func (m *PassCreateProductsRequestRequest) GetMeta() *requestmeta.RequestMeta {
	if m != nil {
		return m.Meta
	}
	return nil
}

func (m *PassCreateProductsRequestRequest) GetServiceFilter() *servicefilter.ServiceFilter {
	if m != nil {
		return m.ServiceFilter
	}
	return nil
}

func (m *PassCreateProductsRequestRequest) GetCreateProductsRequest() *createproductsrequest.CreateProductsRequest {
	if m != nil {
		return m.CreateProductsRequest
	}
	return nil
}

func (m *PassCreateProductsRequestRequest) GetHash() string {
	if m != nil && m.Hash != nil {
		return *m.Hash
	}
	return ""
}

func init() {
	proto.RegisterType((*PassCreateProductsRequestRequest)(nil), "passcreateproductsrequestrequest.PassCreateProductsRequestRequest")
}

func init() {
	proto.RegisterFile("passcreateproductsrequestrequest/passcreateproductsrequestrequest_.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 260 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x91, 0xc1, 0x4a, 0xc4, 0x30,
	0x10, 0x86, 0xe9, 0xda, 0x8b, 0x11, 0x2f, 0x01, 0x21, 0x2c, 0x22, 0x65, 0x4f, 0x7b, 0x58, 0x13,
	0xec, 0x23, 0xac, 0x20, 0x1e, 0x54, 0x96, 0x7a, 0xf3, 0x22, 0xd9, 0xec, 0xd8, 0x16, 0xb6, 0x26,
	0x66, 0x26, 0x82, 0xef, 0xe0, 0x43, 0x8b, 0x31, 0x42, 0x0b, 0x81, 0x9e, 0x66, 0x12, 0xbe, 0xf9,
	0xf3, 0x0d, 0x61, 0xf7, 0x4e, 0x23, 0x1a, 0x0f, 0x9a, 0xc0, 0x79, 0x7b, 0x08, 0x86, 0xd0, 0xc3,
	0x47, 0x00, 0xa4, 0x54, 0xd4, 0x1c, 0xf0, 0x2a, 0x9d, 0xb7, 0x64, 0x79, 0x35, 0x07, 0x2e, 0xaf,
	0x52, 0x33, 0x00, 0x69, 0x35, 0xea, 0x53, 0xc2, 0x72, 0x85, 0xe0, 0x3f, 0x7b, 0x03, 0x6f, 0xfd,
	0x91, 0xc0, 0xab, 0xc9, 0xe9, 0x9f, 0xa9, 0xb3, 0x2f, 0xa8, 0xec, 0x6d, 0x9a, 0x59, 0x7d, 0x2f,
	0x58, 0xb5, 0xd3, 0x88, 0xb7, 0x11, 0xda, 0x25, 0xa8, 0xf9, 0x83, 0x52, 0xe1, 0x1b, 0x56, 0xfe,
	0xba, 0x88, 0xa2, 0x2a, 0xd6, 0x67, 0xb5, 0x90, 0x23, 0x3f, 0x99, 0x98, 0x47, 0x20, 0xdd, 0x44,
	0x8a, 0x6f, 0xd9, 0x79, 0xd2, 0xbb, 0x8b, 0x7a, 0x62, 0x11, 0xc7, 0x2e, 0xe5, 0x44, 0x5a, 0x3e,
	0x8f, 0x99, 0x66, 0x3a, 0xc2, 0xf7, 0xec, 0xc2, 0xe4, 0x8c, 0xc4, 0x49, 0xcc, 0xda, 0xc8, 0xec,
	0x52, 0x32, 0xbf, 0x45, 0x3e, 0x8a, 0x73, 0x56, 0x76, 0x1a, 0x3b, 0x51, 0x56, 0xc5, 0xfa, 0xb4,
	0x89, 0xfd, 0xf6, 0xe9, 0xe5, 0xa1, 0xed, 0xa9, 0x0b, 0x7b, 0x69, 0xec, 0xa0, 0xea, 0x1b, 0xa4,
	0xde, 0xaa, 0xd6, 0x5e, 0xbb, 0xa3, 0xfe, 0x6a, 0xbd, 0x0d, 0xef, 0x07, 0xd5, 0x7a, 0x67, 0x14,
	0x9a, 0x0e, 0x06, 0x3d, 0xfb, 0xff, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x90, 0x43, 0x95, 0x8b,
	0x43, 0x02, 0x00, 0x00,
}
