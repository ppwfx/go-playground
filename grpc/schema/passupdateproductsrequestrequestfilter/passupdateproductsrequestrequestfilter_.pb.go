// Code generated by protoc-gen-go. DO NOT EDIT.
// source: passupdateproductsrequestrequestfilter/passupdateproductsrequestrequestfilter_.proto

/*
Package passupdateproductsrequestrequestfilter is a generated protocol buffer package.

It is generated from these files:
	passupdateproductsrequestrequestfilter/passupdateproductsrequestrequestfilter_.proto

It has these top-level messages:
	PassUpdateProductsRequestRequestFilter
*/
package passupdateproductsrequestrequestfilter

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import requestmetafilter "github.com/21stio/go-playground/grpc/schema/requestmetafilter"
import updateproductsrequestfilter "github.com/21stio/go-playground/grpc/schema/updateproductsrequestfilter"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type PassUpdateProductsRequestRequestFilter struct {
	Meta                  *requestmetafilter.RequestMetaFilter                     `protobuf:"bytes,1,opt,name=meta" json:"meta,omitempty"`
	UpdateProductsRequest *updateproductsrequestfilter.UpdateProductsRequestFilter `protobuf:"bytes,2,opt,name=updateProductsRequest" json:"updateProductsRequest,omitempty"`
	And                   []*PassUpdateProductsRequestRequestFilter                `protobuf:"bytes,3,rep,name=and" json:"and,omitempty"`
	Or                    []*PassUpdateProductsRequestRequestFilter                `protobuf:"bytes,4,rep,name=or" json:"or,omitempty"`
	Not                   []*PassUpdateProductsRequestRequestFilter                `protobuf:"bytes,5,rep,name=not" json:"not,omitempty"`
	Hash                  *string                                                  `protobuf:"bytes,6,opt,name=hash" json:"hash,omitempty"`
	XXX_unrecognized      []byte                                                   `json:"-"`
}

func (m *PassUpdateProductsRequestRequestFilter) Reset() {
	*m = PassUpdateProductsRequestRequestFilter{}
}
func (m *PassUpdateProductsRequestRequestFilter) String() string { return proto.CompactTextString(m) }
func (*PassUpdateProductsRequestRequestFilter) ProtoMessage()    {}
func (*PassUpdateProductsRequestRequestFilter) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{0}
}

func (m *PassUpdateProductsRequestRequestFilter) GetMeta() *requestmetafilter.RequestMetaFilter {
	if m != nil {
		return m.Meta
	}
	return nil
}

func (m *PassUpdateProductsRequestRequestFilter) GetUpdateProductsRequest() *updateproductsrequestfilter.UpdateProductsRequestFilter {
	if m != nil {
		return m.UpdateProductsRequest
	}
	return nil
}

func (m *PassUpdateProductsRequestRequestFilter) GetAnd() []*PassUpdateProductsRequestRequestFilter {
	if m != nil {
		return m.And
	}
	return nil
}

func (m *PassUpdateProductsRequestRequestFilter) GetOr() []*PassUpdateProductsRequestRequestFilter {
	if m != nil {
		return m.Or
	}
	return nil
}

func (m *PassUpdateProductsRequestRequestFilter) GetNot() []*PassUpdateProductsRequestRequestFilter {
	if m != nil {
		return m.Not
	}
	return nil
}

func (m *PassUpdateProductsRequestRequestFilter) GetHash() string {
	if m != nil && m.Hash != nil {
		return *m.Hash
	}
	return ""
}

func init() {
	proto.RegisterType((*PassUpdateProductsRequestRequestFilter)(nil), "passupdateproductsrequestrequestfilter.PassUpdateProductsRequestRequestFilter")
}

func init() {
	proto.RegisterFile("passupdateproductsrequestrequestfilter/passupdateproductsrequestrequestfilter_.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 285 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x92, 0x3f, 0x6b, 0xc3, 0x30,
	0x10, 0xc5, 0xf1, 0x9f, 0x16, 0xaa, 0x6c, 0x82, 0x82, 0xc8, 0x64, 0x4a, 0x09, 0xa6, 0x50, 0x99,
	0x7a, 0xca, 0xd4, 0xa1, 0x43, 0xb7, 0x96, 0x60, 0xd2, 0xa5, 0x43, 0x5b, 0xd5, 0x56, 0x6d, 0x43,
	0xec, 0x53, 0xa5, 0xd3, 0xd0, 0xaf, 0xd0, 0x4f, 0x1d, 0xec, 0xdc, 0x16, 0x13, 0xbc, 0x64, 0x3a,
	0x5b, 0x7a, 0xef, 0x77, 0xef, 0x0e, 0xb1, 0xad, 0x51, 0xce, 0x79, 0x53, 0x29, 0xd4, 0xc6, 0x42,
	0xe5, 0x4b, 0x74, 0x56, 0xff, 0x7a, 0xed, 0x90, 0xca, 0x4f, 0xbb, 0x43, 0x6d, 0xb3, 0x79, 0xb2,
	0x4f, 0x69, 0x2c, 0x20, 0xf0, 0xd5, 0x3c, 0xf9, 0xf2, 0x8e, 0x7e, 0x3b, 0x8d, 0x8a, 0x1a, 0x1d,
	0x9d, 0x10, 0x73, 0xf9, 0x38, 0xc9, 0x23, 0xd7, 0x89, 0x3b, 0xf2, 0xdf, 0xfc, 0xc7, 0x6c, 0xb5,
	0x51, 0xce, 0xbd, 0x8d, 0xd2, 0x0d, 0x49, 0x8b, 0x83, 0x94, 0xca, 0xf3, 0xe8, 0xe0, 0x6b, 0x16,
	0x0f, 0xfd, 0x45, 0x90, 0x04, 0xe9, 0x22, 0xbf, 0x95, 0x47, 0x99, 0x24, 0xe9, 0x5f, 0x34, 0xaa,
	0x83, 0xa7, 0x18, 0x1d, 0xbc, 0x67, 0xd7, 0x7e, 0x8a, 0x2f, 0xc2, 0x11, 0xb5, 0x96, 0x27, 0x82,
	0xca, 0xc9, 0x64, 0x84, 0x9f, 0xc6, 0xf2, 0x2f, 0x16, 0xa9, 0xbe, 0x12, 0x51, 0x12, 0xa5, 0x8b,
	0xfc, 0x55, 0xce, 0x5b, 0xbb, 0x9c, 0xb7, 0x86, 0x62, 0x40, 0xf3, 0x0f, 0x16, 0x82, 0x15, 0xf1,
	0x59, 0x1a, 0x84, 0x60, 0x87, 0x09, 0x7a, 0x40, 0x71, 0x71, 0x9e, 0x09, 0x7a, 0x40, 0xce, 0x59,
	0xdc, 0x28, 0xd7, 0x88, 0xcb, 0x24, 0x48, 0xaf, 0x8a, 0xf1, 0xfb, 0x69, 0xfb, 0x5e, 0xd4, 0x2d,
	0x36, 0xfe, 0x5b, 0x96, 0xd0, 0x65, 0xf9, 0x83, 0xc3, 0x16, 0xb2, 0x1a, 0xee, 0xcd, 0x4e, 0xfd,
	0xd5, 0x16, 0x7c, 0x5f, 0x65, 0xb5, 0x35, 0x65, 0xe6, 0xca, 0x46, 0x77, 0x6a, 0xe6, 0xeb, 0xdf,
	0x07, 0x00, 0x00, 0xff, 0xff, 0x69, 0xfb, 0xff, 0x8a, 0x4d, 0x03, 0x00, 0x00,
}
