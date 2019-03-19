// Code generated by protoc-gen-go. DO NOT EDIT.
// source: deletecompaniesresponse/deletecompaniesresponse_.proto

/*
Package deletecompaniesresponse is a generated protocol buffer package.

It is generated from these files:
	deletecompaniesresponse/deletecompaniesresponse_.proto

It has these top-level messages:
	DeleteCompaniesResponse
*/
package deletecompaniesresponse

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

type DeleteCompaniesResponse struct {
	Meta             *responsemeta.ResponseMeta `protobuf:"bytes,1,opt,name=meta" json:"meta,omitempty"`
	Hash             *string                    `protobuf:"bytes,2,opt,name=hash" json:"hash,omitempty"`
	XXX_unrecognized []byte                     `json:"-"`
}

func (m *DeleteCompaniesResponse) Reset()                    { *m = DeleteCompaniesResponse{} }
func (m *DeleteCompaniesResponse) String() string            { return proto.CompactTextString(m) }
func (*DeleteCompaniesResponse) ProtoMessage()               {}
func (*DeleteCompaniesResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *DeleteCompaniesResponse) GetMeta() *responsemeta.ResponseMeta {
	if m != nil {
		return m.Meta
	}
	return nil
}

func (m *DeleteCompaniesResponse) GetHash() string {
	if m != nil && m.Hash != nil {
		return *m.Hash
	}
	return ""
}

func init() {
	proto.RegisterType((*DeleteCompaniesResponse)(nil), "deletecompaniesresponse.DeleteCompaniesResponse")
}

func init() {
	proto.RegisterFile("deletecompaniesresponse/deletecompaniesresponse_.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 181 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x8f, 0xb1, 0x0e, 0x82, 0x30,
	0x10, 0x40, 0x83, 0x61, 0xb1, 0x6e, 0x5d, 0x20, 0x4c, 0xc4, 0x89, 0xc5, 0x36, 0x32, 0xf8, 0x01,
	0xa2, 0xa3, 0x0b, 0xa3, 0x89, 0x31, 0xb5, 0x5c, 0x5a, 0x12, 0xca, 0x35, 0x6d, 0x19, 0xfc, 0x7b,
	0x23, 0x81, 0x84, 0xa5, 0xdb, 0xbb, 0xbb, 0xf7, 0x86, 0x23, 0x97, 0x0e, 0x06, 0x08, 0x20, 0xd1,
	0x58, 0x31, 0xf6, 0xe0, 0x1d, 0x78, 0x8b, 0xa3, 0x07, 0x1e, 0xd9, 0xbf, 0x99, 0x75, 0x18, 0x90,
	0x66, 0x91, 0x7b, 0x51, 0xae, 0x64, 0x20, 0x08, 0xbe, 0x1d, 0x96, 0xf4, 0xf8, 0x22, 0xd9, 0x6d,
	0x8e, 0x9b, 0x35, 0x6e, 0x17, 0x8b, 0x32, 0x92, 0xfe, 0xcd, 0x3c, 0x29, 0x93, 0xea, 0x50, 0x17,
	0x6c, 0x9b, 0xb3, 0xd5, 0x7a, 0x40, 0x10, 0xed, 0xec, 0x51, 0x4a, 0x52, 0x2d, 0xbc, 0xce, 0x77,
	0x65, 0x52, 0xed, 0xdb, 0x99, 0xaf, 0xf7, 0x67, 0xa3, 0xfa, 0xa0, 0xa7, 0x0f, 0x93, 0x68, 0x78,
	0x7d, 0xf6, 0xa1, 0x47, 0xae, 0xf0, 0x64, 0x07, 0xf1, 0x55, 0x0e, 0xa7, 0xb1, 0xe3, 0xca, 0x59,
	0xc9, 0xbd, 0xd4, 0x60, 0x44, 0xec, 0xcf, 0x5f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x8f, 0x62, 0xb9,
	0xe6, 0x19, 0x01, 0x00, 0x00,
}