// Code generated by protoc-gen-go. DO NOT EDIT.
// source: createproductsresponse/createproductsresponse_.proto

/*
Package createproductsresponse is a generated protocol buffer package.

It is generated from these files:
	createproductsresponse/createproductsresponse_.proto

It has these top-level messages:
	CreateProductsResponse
*/
package createproductsresponse

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

type CreateProductsResponse struct {
	Meta             *responsemeta.ResponseMeta `protobuf:"bytes,1,opt,name=meta" json:"meta,omitempty"`
	Hash             *string                    `protobuf:"bytes,2,opt,name=hash" json:"hash,omitempty"`
	XXX_unrecognized []byte                     `json:"-"`
}

func (m *CreateProductsResponse) Reset()                    { *m = CreateProductsResponse{} }
func (m *CreateProductsResponse) String() string            { return proto.CompactTextString(m) }
func (*CreateProductsResponse) ProtoMessage()               {}
func (*CreateProductsResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *CreateProductsResponse) GetMeta() *responsemeta.ResponseMeta {
	if m != nil {
		return m.Meta
	}
	return nil
}

func (m *CreateProductsResponse) GetHash() string {
	if m != nil && m.Hash != nil {
		return *m.Hash
	}
	return ""
}

func init() {
	proto.RegisterType((*CreateProductsResponse)(nil), "createproductsresponse.CreateProductsResponse")
}

func init() {
	proto.RegisterFile("createproductsresponse/createproductsresponse_.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 180 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x8f, 0xb1, 0x0a, 0xc2, 0x30,
	0x10, 0x40, 0xa9, 0x74, 0x31, 0x6e, 0x19, 0x4a, 0xe9, 0x54, 0x9c, 0xba, 0x98, 0x60, 0xf1, 0x0b,
	0xaa, 0xab, 0x20, 0x1d, 0x45, 0x90, 0x98, 0x1e, 0x49, 0xc1, 0xf6, 0x42, 0x72, 0x1d, 0xfc, 0x7b,
	0xb1, 0xb4, 0xd0, 0xa1, 0x6e, 0xef, 0xee, 0xde, 0x1b, 0x8e, 0x9d, 0xb4, 0x07, 0x45, 0xe0, 0x3c,
	0x36, 0x83, 0xa6, 0xe0, 0x21, 0x38, 0xec, 0x03, 0xc8, 0xf5, 0xf5, 0x53, 0x38, 0x8f, 0x84, 0x3c,
	0x59, 0x3f, 0x67, 0xf9, 0x4c, 0x1d, 0x90, 0x92, 0xcb, 0x61, 0x2a, 0xf7, 0x0f, 0x96, 0x9c, 0xc7,
	0xf6, 0x36, 0xb5, 0xf5, 0x24, 0x71, 0xc1, 0xe2, 0x9f, 0x98, 0x46, 0x79, 0x54, 0xec, 0xca, 0x4c,
	0x2c, 0x6b, 0x31, 0x5b, 0x57, 0x20, 0x55, 0x8f, 0x1e, 0xe7, 0x2c, 0xb6, 0x2a, 0xd8, 0x74, 0x93,
	0x47, 0xc5, 0xb6, 0x1e, 0xb9, 0xba, 0xdc, 0x2b, 0xd3, 0x92, 0x1d, 0x5e, 0x42, 0x63, 0x27, 0xcb,
	0x63, 0xa0, 0x16, 0xa5, 0xc1, 0x83, 0x7b, 0xab, 0x8f, 0xf1, 0x38, 0xf4, 0x8d, 0x34, 0xde, 0x69,
	0x19, 0xb4, 0x85, 0x4e, 0xfd, 0x79, 0xf2, 0x1b, 0x00, 0x00, 0xff, 0xff, 0x79, 0x1f, 0xa8, 0x6e,
	0x14, 0x01, 0x00, 0x00,
}
