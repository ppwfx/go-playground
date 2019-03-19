// Code generated by protoc-gen-go. DO NOT EDIT.
// source: deleteproductsresponse/deleteproductsresponse_.proto

/*
Package deleteproductsresponse is a generated protocol buffer package.

It is generated from these files:
	deleteproductsresponse/deleteproductsresponse_.proto

It has these top-level messages:
	DeleteProductsResponse
*/
package deleteproductsresponse

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

type DeleteProductsResponse struct {
	Meta             *responsemeta.ResponseMeta `protobuf:"bytes,1,opt,name=meta" json:"meta,omitempty"`
	Hash             *string                    `protobuf:"bytes,2,opt,name=hash" json:"hash,omitempty"`
	XXX_unrecognized []byte                     `json:"-"`
}

func (m *DeleteProductsResponse) Reset()                    { *m = DeleteProductsResponse{} }
func (m *DeleteProductsResponse) String() string            { return proto.CompactTextString(m) }
func (*DeleteProductsResponse) ProtoMessage()               {}
func (*DeleteProductsResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *DeleteProductsResponse) GetMeta() *responsemeta.ResponseMeta {
	if m != nil {
		return m.Meta
	}
	return nil
}

func (m *DeleteProductsResponse) GetHash() string {
	if m != nil && m.Hash != nil {
		return *m.Hash
	}
	return ""
}

func init() {
	proto.RegisterType((*DeleteProductsResponse)(nil), "deleteproductsresponse.DeleteProductsResponse")
}

func init() {
	proto.RegisterFile("deleteproductsresponse/deleteproductsresponse_.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 181 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x8f, 0x31, 0xcb, 0x83, 0x30,
	0x10, 0x40, 0xf1, 0xc3, 0xe5, 0x4b, 0xb7, 0x0c, 0x22, 0x4e, 0xd2, 0xc9, 0xa5, 0x09, 0x95, 0xfe,
	0x02, 0x71, 0x2d, 0x14, 0xc7, 0x52, 0x28, 0x69, 0x3c, 0x12, 0x41, 0xbd, 0x90, 0x9c, 0x43, 0xff,
	0x7d, 0xa9, 0x28, 0x38, 0xd8, 0xed, 0xdd, 0xdd, 0x7b, 0xc3, 0xb1, 0x4b, 0x0b, 0x3d, 0x10, 0x38,
	0x8f, 0xed, 0xa4, 0x29, 0x78, 0x08, 0x0e, 0xc7, 0x00, 0x72, 0x7f, 0xfd, 0x14, 0xce, 0x23, 0x21,
	0x4f, 0xf6, 0xcf, 0x59, 0xbe, 0xd2, 0x00, 0xa4, 0xe4, 0x76, 0x58, 0xca, 0xe3, 0x83, 0x25, 0xf5,
	0xdc, 0xde, 0x96, 0xb6, 0x59, 0x24, 0x2e, 0x58, 0xfc, 0x15, 0xd3, 0x28, 0x8f, 0x8a, 0x43, 0x99,
	0x89, 0x6d, 0x2d, 0x56, 0xeb, 0x0a, 0xa4, 0x9a, 0xd9, 0xe3, 0x9c, 0xc5, 0x56, 0x05, 0x9b, 0xfe,
	0xe5, 0x51, 0xf1, 0xdf, 0xcc, 0x5c, 0xd5, 0xf7, 0xca, 0x74, 0x64, 0xa7, 0x97, 0xd0, 0x38, 0xc8,
	0xf2, 0x1c, 0xa8, 0x43, 0x69, 0xf0, 0xe4, 0x7a, 0xf5, 0x36, 0x1e, 0xa7, 0xb1, 0x95, 0xc6, 0x3b,
	0x2d, 0x83, 0xb6, 0x30, 0xa8, 0x1f, 0x4f, 0x7e, 0x02, 0x00, 0x00, 0xff, 0xff, 0x94, 0x10, 0x7a,
	0xcb, 0x14, 0x01, 0x00, 0x00,
}
