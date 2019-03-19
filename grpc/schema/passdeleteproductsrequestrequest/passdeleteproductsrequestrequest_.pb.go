// Code generated by protoc-gen-go. DO NOT EDIT.
// source: passdeleteproductsrequestrequest/passdeleteproductsrequestrequest_.proto

/*
Package passdeleteproductsrequestrequest is a generated protocol buffer package.

It is generated from these files:
	passdeleteproductsrequestrequest/passdeleteproductsrequestrequest_.proto

It has these top-level messages:
	PassDeleteProductsRequestRequest
*/
package passdeleteproductsrequestrequest

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import requestmeta "github.com/21stio/go-playground/grpc/schema/requestmeta"
import servicefilter "github.com/21stio/go-playground/grpc/schema/servicefilter"
import deleteproductsrequest "github.com/21stio/go-playground/grpc/schema/deleteproductsrequest"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type PassDeleteProductsRequestRequest struct {
	Meta                  *requestmeta.RequestMeta                     `protobuf:"bytes,1,opt,name=meta" json:"meta,omitempty"`
	ServiceFilter         *servicefilter.ServiceFilter                 `protobuf:"bytes,2,opt,name=serviceFilter" json:"serviceFilter,omitempty"`
	DeleteProductsRequest *deleteproductsrequest.DeleteProductsRequest `protobuf:"bytes,3,opt,name=deleteProductsRequest" json:"deleteProductsRequest,omitempty"`
	Hash                  *string                                      `protobuf:"bytes,4,opt,name=hash" json:"hash,omitempty"`
	XXX_unrecognized      []byte                                       `json:"-"`
}

func (m *PassDeleteProductsRequestRequest) Reset()         { *m = PassDeleteProductsRequestRequest{} }
func (m *PassDeleteProductsRequestRequest) String() string { return proto.CompactTextString(m) }
func (*PassDeleteProductsRequestRequest) ProtoMessage()    {}
func (*PassDeleteProductsRequestRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{0}
}

func (m *PassDeleteProductsRequestRequest) GetMeta() *requestmeta.RequestMeta {
	if m != nil {
		return m.Meta
	}
	return nil
}

func (m *PassDeleteProductsRequestRequest) GetServiceFilter() *servicefilter.ServiceFilter {
	if m != nil {
		return m.ServiceFilter
	}
	return nil
}

func (m *PassDeleteProductsRequestRequest) GetDeleteProductsRequest() *deleteproductsrequest.DeleteProductsRequest {
	if m != nil {
		return m.DeleteProductsRequest
	}
	return nil
}

func (m *PassDeleteProductsRequestRequest) GetHash() string {
	if m != nil && m.Hash != nil {
		return *m.Hash
	}
	return ""
}

func init() {
	proto.RegisterType((*PassDeleteProductsRequestRequest)(nil), "passdeleteproductsrequestrequest.PassDeleteProductsRequestRequest")
}

func init() {
	proto.RegisterFile("passdeleteproductsrequestrequest/passdeleteproductsrequestrequest_.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 260 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x91, 0xc1, 0x4a, 0xc4, 0x30,
	0x10, 0x86, 0xe9, 0xda, 0x8b, 0x11, 0x2f, 0x01, 0x21, 0x2c, 0x22, 0x65, 0x4f, 0x7b, 0x58, 0x13,
	0xec, 0x23, 0x2c, 0x22, 0x1e, 0x54, 0x96, 0x7a, 0xf3, 0x22, 0xd9, 0x74, 0x6c, 0x0b, 0xad, 0x89,
	0x99, 0xa9, 0xe0, 0x3b, 0xf8, 0xd0, 0x62, 0x8c, 0xd0, 0x42, 0xa0, 0xa7, 0x99, 0x84, 0x6f, 0xfe,
	0x7c, 0x43, 0xd8, 0xbd, 0xd3, 0x88, 0x35, 0xf4, 0x40, 0xe0, 0xbc, 0xad, 0x47, 0x43, 0xe8, 0xe1,
	0x63, 0x04, 0xa4, 0x58, 0xd4, 0x12, 0xf0, 0x2a, 0x9d, 0xb7, 0x64, 0x79, 0xb1, 0x04, 0xae, 0xaf,
	0x62, 0x33, 0x00, 0x69, 0x35, 0xe9, 0x63, 0xc2, 0x7a, 0x83, 0xe0, 0x3f, 0x3b, 0x03, 0x6f, 0x5d,
	0x4f, 0xe0, 0xd5, 0xec, 0xf4, 0xcf, 0x94, 0xc9, 0x17, 0x54, 0xf2, 0x36, 0xce, 0x6c, 0xbe, 0x57,
	0xac, 0x38, 0x68, 0xc4, 0xdb, 0x00, 0x1d, 0x22, 0x54, 0xfd, 0x41, 0xb1, 0xf0, 0x1d, 0xcb, 0x7f,
	0x5d, 0x44, 0x56, 0x64, 0xdb, 0xb3, 0x52, 0xc8, 0x89, 0x9f, 0x8c, 0xcc, 0x23, 0x90, 0xae, 0x02,
	0xc5, 0xf7, 0xec, 0x3c, 0xea, 0xdd, 0x05, 0x3d, 0xb1, 0x0a, 0x63, 0x97, 0x72, 0x26, 0x2d, 0x9f,
	0xa7, 0x4c, 0x35, 0x1f, 0xe1, 0x47, 0x76, 0x51, 0xa7, 0x8c, 0xc4, 0x49, 0xc8, 0xda, 0xc9, 0xe4,
	0x52, 0x32, 0xbd, 0x45, 0x3a, 0x8a, 0x73, 0x96, 0xb7, 0x1a, 0x5b, 0x91, 0x17, 0xd9, 0xf6, 0xb4,
	0x0a, 0xfd, 0xfe, 0xe9, 0xe5, 0xa1, 0xe9, 0xa8, 0x1d, 0x8f, 0xd2, 0xd8, 0x41, 0x95, 0x37, 0x48,
	0x9d, 0x55, 0x8d, 0xbd, 0x76, 0xbd, 0xfe, 0x6a, 0xbc, 0x1d, 0xdf, 0x6b, 0xd5, 0x78, 0x67, 0x14,
	0x9a, 0x16, 0x06, 0xbd, 0xf8, 0xff, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x79, 0x24, 0x6c, 0x23,
	0x43, 0x02, 0x00, 0x00,
}
