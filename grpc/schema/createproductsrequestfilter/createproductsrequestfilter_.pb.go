// Code generated by protoc-gen-go. DO NOT EDIT.
// source: createproductsrequestfilter/createproductsrequestfilter_.proto

/*
Package createproductsrequestfilter is a generated protocol buffer package.

It is generated from these files:
	createproductsrequestfilter/createproductsrequestfilter_.proto

It has these top-level messages:
	CreateProductsRequestFilter
*/
package createproductsrequestfilter

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import requestmetafilter "github.com/21stio/go-playground/grpc/schema/requestmetafilter"
import productfilter "github.com/21stio/go-playground/grpc/schema/productfilter"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type CreateProductsRequestFilter struct {
	Meta             *requestmetafilter.RequestMetaFilter `protobuf:"bytes,1,opt,name=meta" json:"meta,omitempty"`
	ProductsSome     *productfilter.ProductFilter         `protobuf:"bytes,2,opt,name=productsSome" json:"productsSome,omitempty"`
	ProductsEvery    *productfilter.ProductFilter         `protobuf:"bytes,3,opt,name=productsEvery" json:"productsEvery,omitempty"`
	ProductsNone     *productfilter.ProductFilter         `protobuf:"bytes,4,opt,name=productsNone" json:"productsNone,omitempty"`
	And              []*CreateProductsRequestFilter       `protobuf:"bytes,5,rep,name=and" json:"and,omitempty"`
	Or               []*CreateProductsRequestFilter       `protobuf:"bytes,6,rep,name=or" json:"or,omitempty"`
	Not              []*CreateProductsRequestFilter       `protobuf:"bytes,7,rep,name=not" json:"not,omitempty"`
	Hash             *string                              `protobuf:"bytes,8,opt,name=hash" json:"hash,omitempty"`
	XXX_unrecognized []byte                               `json:"-"`
}

func (m *CreateProductsRequestFilter) Reset()                    { *m = CreateProductsRequestFilter{} }
func (m *CreateProductsRequestFilter) String() string            { return proto.CompactTextString(m) }
func (*CreateProductsRequestFilter) ProtoMessage()               {}
func (*CreateProductsRequestFilter) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *CreateProductsRequestFilter) GetMeta() *requestmetafilter.RequestMetaFilter {
	if m != nil {
		return m.Meta
	}
	return nil
}

func (m *CreateProductsRequestFilter) GetProductsSome() *productfilter.ProductFilter {
	if m != nil {
		return m.ProductsSome
	}
	return nil
}

func (m *CreateProductsRequestFilter) GetProductsEvery() *productfilter.ProductFilter {
	if m != nil {
		return m.ProductsEvery
	}
	return nil
}

func (m *CreateProductsRequestFilter) GetProductsNone() *productfilter.ProductFilter {
	if m != nil {
		return m.ProductsNone
	}
	return nil
}

func (m *CreateProductsRequestFilter) GetAnd() []*CreateProductsRequestFilter {
	if m != nil {
		return m.And
	}
	return nil
}

func (m *CreateProductsRequestFilter) GetOr() []*CreateProductsRequestFilter {
	if m != nil {
		return m.Or
	}
	return nil
}

func (m *CreateProductsRequestFilter) GetNot() []*CreateProductsRequestFilter {
	if m != nil {
		return m.Not
	}
	return nil
}

func (m *CreateProductsRequestFilter) GetHash() string {
	if m != nil && m.Hash != nil {
		return *m.Hash
	}
	return ""
}

func init() {
	proto.RegisterType((*CreateProductsRequestFilter)(nil), "createproductsrequestfilter.CreateProductsRequestFilter")
}

func init() {
	proto.RegisterFile("createproductsrequestfilter/createproductsrequestfilter_.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 303 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x92, 0x4b, 0x4b, 0xc3, 0x40,
	0x14, 0x85, 0xc9, 0xc3, 0xd7, 0x54, 0x37, 0xb3, 0x1a, 0x5a, 0x17, 0xa1, 0xb8, 0x08, 0x82, 0x13,
	0xcc, 0xaa, 0x2b, 0x91, 0x8a, 0x4f, 0x50, 0x24, 0xee, 0xdc, 0xc8, 0x98, 0x5c, 0x93, 0x40, 0x93,
	0x1b, 0x27, 0x37, 0x42, 0x7f, 0x88, 0xff, 0x57, 0x92, 0x4e, 0xc1, 0x50, 0x08, 0x4a, 0x77, 0xb9,
	0x99, 0x73, 0xbe, 0x73, 0xe6, 0xc1, 0x2e, 0x62, 0x0d, 0x8a, 0xa0, 0xd2, 0x98, 0x34, 0x31, 0xd5,
	0x1a, 0x3e, 0x1b, 0xa8, 0xe9, 0x23, 0x5f, 0x10, 0xe8, 0x60, 0x60, 0xed, 0x4d, 0x56, 0x1a, 0x09,
	0xf9, 0x64, 0x40, 0x33, 0x3e, 0x35, 0x63, 0x01, 0xa4, 0x0c, 0x72, 0xe3, 0x8f, 0x01, 0x8d, 0xa7,
	0x06, 0x61, 0x74, 0xbd, 0xc9, 0x68, 0xa6, 0xdf, 0x2e, 0x9b, 0x5c, 0x75, 0x79, 0xcf, 0x26, 0x2f,
	0x5a, 0xe1, 0x6e, 0x3a, 0x19, 0x9f, 0x31, 0xb7, 0x05, 0x0b, 0xcb, 0xb3, 0xfc, 0x51, 0x78, 0x22,
	0x37, 0xc2, 0xa4, 0xd1, 0x3f, 0x02, 0xa9, 0x95, 0x27, 0xea, 0x1c, 0xfc, 0x92, 0x1d, 0xae, 0xb7,
	0xf0, 0x82, 0x05, 0x08, 0xbb, 0x23, 0x1c, 0xcb, 0x5e, 0x0d, 0x69, 0x52, 0x8d, 0xb3, 0xe7, 0xe0,
	0x73, 0x76, 0xb4, 0x9e, 0xaf, 0xbf, 0x40, 0x2f, 0x85, 0xf3, 0x07, 0x44, 0xdf, 0xf2, 0xbb, 0xc5,
	0x13, 0x96, 0x20, 0xdc, 0xff, 0xb4, 0x68, 0x1d, 0xfc, 0x81, 0x39, 0xaa, 0x4c, 0xc4, 0x8e, 0xe7,
	0xf8, 0xa3, 0x70, 0x26, 0x07, 0x2e, 0x47, 0x0e, 0x1c, 0x64, 0xd4, 0x42, 0xf8, 0x1d, 0xb3, 0x51,
	0x8b, 0xdd, 0x2d, 0x51, 0x36, 0xea, 0xb6, 0x55, 0x89, 0x24, 0xf6, 0xb6, 0x6d, 0x55, 0x22, 0x71,
	0xce, 0xdc, 0x4c, 0xd5, 0x99, 0xd8, 0xf7, 0x2c, 0xff, 0x20, 0xea, 0xbe, 0xe7, 0xf7, 0xaf, 0xb7,
	0x69, 0x4e, 0x59, 0xf3, 0x2e, 0x63, 0x2c, 0x82, 0xf0, 0xbc, 0xa6, 0x1c, 0x83, 0x14, 0xcf, 0xaa,
	0x85, 0x5a, 0xa6, 0x1a, 0x9b, 0x32, 0x09, 0x52, 0x5d, 0xc5, 0x41, 0x1d, 0x67, 0x50, 0xa8, 0xa1,
	0x67, 0xfd, 0x13, 0x00, 0x00, 0xff, 0xff, 0x58, 0x30, 0x79, 0x0a, 0x10, 0x03, 0x00, 0x00,
}
