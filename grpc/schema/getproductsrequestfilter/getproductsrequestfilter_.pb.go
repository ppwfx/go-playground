// Code generated by protoc-gen-go. DO NOT EDIT.
// source: getproductsrequestfilter/getproductsrequestfilter_.proto

/*
Package getproductsrequestfilter is a generated protocol buffer package.

It is generated from these files:
	getproductsrequestfilter/getproductsrequestfilter_.proto

It has these top-level messages:
	GetProductsRequestFilter
*/
package getproductsrequestfilter

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import requestmetafilter "github.com/21stio/go-playground/grpc/schema/requestmetafilter"
import locationqueryfilter "github.com/21stio/go-playground/grpc/schema/locationqueryfilter"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type GetProductsRequestFilter struct {
	Meta             *requestmetafilter.RequestMetaFilter     `protobuf:"bytes,1,opt,name=meta" json:"meta,omitempty"`
	LocationQuery    *locationqueryfilter.LocationQueryFilter `protobuf:"bytes,2,opt,name=locationQuery" json:"locationQuery,omitempty"`
	And              []*GetProductsRequestFilter              `protobuf:"bytes,3,rep,name=and" json:"and,omitempty"`
	Or               []*GetProductsRequestFilter              `protobuf:"bytes,4,rep,name=or" json:"or,omitempty"`
	Not              []*GetProductsRequestFilter              `protobuf:"bytes,5,rep,name=not" json:"not,omitempty"`
	Hash             *string                                  `protobuf:"bytes,6,opt,name=hash" json:"hash,omitempty"`
	XXX_unrecognized []byte                                   `json:"-"`
}

func (m *GetProductsRequestFilter) Reset()                    { *m = GetProductsRequestFilter{} }
func (m *GetProductsRequestFilter) String() string            { return proto.CompactTextString(m) }
func (*GetProductsRequestFilter) ProtoMessage()               {}
func (*GetProductsRequestFilter) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *GetProductsRequestFilter) GetMeta() *requestmetafilter.RequestMetaFilter {
	if m != nil {
		return m.Meta
	}
	return nil
}

func (m *GetProductsRequestFilter) GetLocationQuery() *locationqueryfilter.LocationQueryFilter {
	if m != nil {
		return m.LocationQuery
	}
	return nil
}

func (m *GetProductsRequestFilter) GetAnd() []*GetProductsRequestFilter {
	if m != nil {
		return m.And
	}
	return nil
}

func (m *GetProductsRequestFilter) GetOr() []*GetProductsRequestFilter {
	if m != nil {
		return m.Or
	}
	return nil
}

func (m *GetProductsRequestFilter) GetNot() []*GetProductsRequestFilter {
	if m != nil {
		return m.Not
	}
	return nil
}

func (m *GetProductsRequestFilter) GetHash() string {
	if m != nil && m.Hash != nil {
		return *m.Hash
	}
	return ""
}

func init() {
	proto.RegisterType((*GetProductsRequestFilter)(nil), "getproductsrequestfilter.GetProductsRequestFilter")
}

func init() {
	proto.RegisterFile("getproductsrequestfilter/getproductsrequestfilter_.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 278 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x92, 0x3f, 0x4f, 0xc3, 0x30,
	0x14, 0xc4, 0x95, 0x3f, 0x20, 0xe1, 0x8a, 0xc5, 0x93, 0xd5, 0x29, 0x42, 0x0c, 0x11, 0x12, 0x8e,
	0xc8, 0xd4, 0xb9, 0xaa, 0xca, 0x02, 0x08, 0x32, 0xb2, 0x20, 0x93, 0x18, 0x27, 0x52, 0xe2, 0x97,
	0xda, 0xcf, 0x43, 0xbf, 0x01, 0x1f, 0x1b, 0x25, 0x35, 0x12, 0xa8, 0xf5, 0xd2, 0xcd, 0x7e, 0xbe,
	0xfb, 0xdd, 0xd9, 0x32, 0x59, 0x29, 0x89, 0xa3, 0x81, 0xc6, 0xd5, 0x68, 0x8d, 0xdc, 0x39, 0x69,
	0xf1, 0xab, 0xeb, 0x51, 0x9a, 0x22, 0x74, 0xf0, 0xc1, 0x47, 0x03, 0x08, 0x94, 0x85, 0x04, 0xcb,
	0x3b, 0xbf, 0x1d, 0x24, 0x0a, 0x0f, 0x3b, 0x9a, 0x78, 0xca, 0x92, 0xf7, 0x50, 0x0b, 0xec, 0x40,
	0xef, 0x9c, 0x34, 0x7b, 0xaf, 0x3e, 0x31, 0xf3, 0xfa, 0x9b, 0xef, 0x84, 0xb0, 0x47, 0x89, 0xaf,
	0x3e, 0xb8, 0x3a, 0x70, 0xb7, 0xb3, 0x86, 0xae, 0x48, 0x3a, 0x25, 0xb0, 0x28, 0x8b, 0xf2, 0x45,
	0x79, 0xcb, 0x8f, 0x52, 0xb9, 0xd7, 0x3f, 0x4b, 0x14, 0x07, 0x4f, 0x35, 0x3b, 0xe8, 0x0b, 0xb9,
	0xfe, 0x0d, 0x7d, 0x9b, 0x42, 0x59, 0x3c, 0x23, 0xf2, 0x53, 0xf5, 0xf8, 0xd3, 0x5f, 0xa5, 0xc7,
	0xfc, 0xb7, 0xd3, 0x0d, 0x49, 0x84, 0x6e, 0x58, 0x92, 0x25, 0xf9, 0xa2, 0x2c, 0x79, 0xe8, 0xa9,
	0x78, 0xe8, 0x2a, 0xd5, 0x64, 0xa7, 0x6b, 0x12, 0x83, 0x61, 0xe9, 0xd9, 0x90, 0x18, 0xcc, 0xd4,
	0x44, 0x03, 0xb2, 0x8b, 0xf3, 0x9b, 0x68, 0x40, 0x4a, 0x49, 0xda, 0x0a, 0xdb, 0xb2, 0xcb, 0x2c,
	0xca, 0xaf, 0xaa, 0x79, 0xbd, 0xde, 0xbe, 0x6f, 0x54, 0x87, 0xad, 0xfb, 0xe4, 0x35, 0x0c, 0x45,
	0xf9, 0x60, 0xb1, 0x83, 0x42, 0xc1, 0xfd, 0xd8, 0x8b, 0xbd, 0x32, 0xe0, 0x74, 0x53, 0x28, 0x33,
	0xd6, 0x85, 0xad, 0x5b, 0x39, 0x88, 0xe0, 0x7f, 0xfa, 0x09, 0x00, 0x00, 0xff, 0xff, 0x79, 0xa4,
	0x50, 0x28, 0x83, 0x02, 0x00, 0x00,
}
