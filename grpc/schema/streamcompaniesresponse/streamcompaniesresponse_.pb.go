// Code generated by protoc-gen-go. DO NOT EDIT.
// source: streamcompaniesresponse/streamcompaniesresponse_.proto

/*
Package streamcompaniesresponse is a generated protocol buffer package.

It is generated from these files:
	streamcompaniesresponse/streamcompaniesresponse_.proto

It has these top-level messages:
	StreamCompaniesResponse
*/
package streamcompaniesresponse

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import responsemeta "github.com/21stio/go-playground/grpc/schema/responsemeta"
import company "github.com/21stio/go-playground/grpc/schema/company"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type StreamCompaniesResponse struct {
	Meta             *responsemeta.ResponseMeta `protobuf:"bytes,1,opt,name=meta" json:"meta,omitempty"`
	TotalCount       *int32                     `protobuf:"varint,2,opt,name=totalCount" json:"totalCount,omitempty"`
	Companies        []*company.Company         `protobuf:"bytes,3,rep,name=companies" json:"companies,omitempty"`
	Hash             *string                    `protobuf:"bytes,4,opt,name=hash" json:"hash,omitempty"`
	XXX_unrecognized []byte                     `json:"-"`
}

func (m *StreamCompaniesResponse) Reset()                    { *m = StreamCompaniesResponse{} }
func (m *StreamCompaniesResponse) String() string            { return proto.CompactTextString(m) }
func (*StreamCompaniesResponse) ProtoMessage()               {}
func (*StreamCompaniesResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *StreamCompaniesResponse) GetMeta() *responsemeta.ResponseMeta {
	if m != nil {
		return m.Meta
	}
	return nil
}

func (m *StreamCompaniesResponse) GetTotalCount() int32 {
	if m != nil && m.TotalCount != nil {
		return *m.TotalCount
	}
	return 0
}

func (m *StreamCompaniesResponse) GetCompanies() []*company.Company {
	if m != nil {
		return m.Companies
	}
	return nil
}

func (m *StreamCompaniesResponse) GetHash() string {
	if m != nil && m.Hash != nil {
		return *m.Hash
	}
	return ""
}

func init() {
	proto.RegisterType((*StreamCompaniesResponse)(nil), "streamcompaniesresponse.StreamCompaniesResponse")
}

func init() {
	proto.RegisterFile("streamcompaniesresponse/streamcompaniesresponse_.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 235 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x90, 0x31, 0x4f, 0xc3, 0x30,
	0x10, 0x85, 0x65, 0x1a, 0x86, 0xba, 0x0b, 0xf2, 0x40, 0xa3, 0x0c, 0xc8, 0x62, 0xca, 0x82, 0x2d,
	0x32, 0xf0, 0x03, 0x88, 0x18, 0x59, 0xcc, 0xc6, 0x82, 0x8e, 0x60, 0x25, 0x91, 0x1a, 0x9f, 0x65,
	0x5f, 0x86, 0xfc, 0x28, 0xfe, 0x23, 0x6a, 0xea, 0xd0, 0x2c, 0x99, 0x7c, 0x7e, 0x7e, 0xef, 0xf3,
	0xd3, 0xf1, 0x97, 0x48, 0xc1, 0xc2, 0xd0, 0xe0, 0xe0, 0xc1, 0xf5, 0x36, 0x06, 0x1b, 0x3d, 0xba,
	0x68, 0xf5, 0x86, 0xfe, 0xa5, 0x7c, 0x40, 0x42, 0x71, 0xdc, 0x78, 0x2f, 0xe4, 0x32, 0x0d, 0x96,
	0x40, 0xaf, 0x2f, 0x29, 0x5a, 0xdc, 0x5f, 0x42, 0x93, 0x4e, 0x67, 0xd2, 0x1f, 0x7f, 0x19, 0x3f,
	0x7e, 0xcc, 0xd4, 0x7a, 0xa1, 0x9a, 0x14, 0x17, 0x8a, 0x67, 0x67, 0x44, 0xce, 0x24, 0x2b, 0x0f,
	0x55, 0xa1, 0xd6, 0x5c, 0xb5, 0xb8, 0xde, 0x2d, 0x81, 0x99, 0x7d, 0xe2, 0x81, 0x73, 0x42, 0x82,
	0x53, 0x8d, 0xa3, 0xa3, 0xfc, 0x46, 0xb2, 0xf2, 0xd6, 0xac, 0x14, 0xa1, 0xf8, 0xfe, 0xbf, 0x7a,
	0xbe, 0x93, 0xbb, 0xf2, 0x50, 0xdd, 0xa9, 0xd4, 0x47, 0x5d, 0xbe, 0x9f, 0xcc, 0xd5, 0x22, 0x04,
	0xcf, 0x3a, 0x88, 0x5d, 0x9e, 0x49, 0x56, 0xee, 0xcd, 0x3c, 0xbf, 0xbe, 0x7d, 0xd6, 0x6d, 0x4f,
	0xdd, 0xf8, 0x7d, 0x0e, 0xea, 0xea, 0x39, 0x52, 0x8f, 0xba, 0xc5, 0x27, 0x7f, 0x82, 0xa9, 0x0d,
	0x38, 0xba, 0x1f, 0xdd, 0x06, 0xdf, 0xe8, 0xd8, 0x74, 0x76, 0x80, 0xad, 0x85, 0xfe, 0x05, 0x00,
	0x00, 0xff, 0xff, 0xcb, 0xa8, 0x76, 0x44, 0x82, 0x01, 0x00, 0x00,
}
