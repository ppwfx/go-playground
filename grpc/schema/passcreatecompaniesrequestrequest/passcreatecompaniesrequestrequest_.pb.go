// Code generated by protoc-gen-go. DO NOT EDIT.
// source: passcreatecompaniesrequestrequest/passcreatecompaniesrequestrequest_.proto

/*
Package passcreatecompaniesrequestrequest is a generated protocol buffer package.

It is generated from these files:
	passcreatecompaniesrequestrequest/passcreatecompaniesrequestrequest_.proto

It has these top-level messages:
	PassCreateCompaniesRequestRequest
*/
package passcreatecompaniesrequestrequest

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import requestmeta "github.com/21stio/go-playground/grpc/schema/requestmeta"
import servicefilter "github.com/21stio/go-playground/grpc/schema/servicefilter"
import createcompaniesrequest "github.com/21stio/go-playground/grpc/schema/createcompaniesrequest"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type PassCreateCompaniesRequestRequest struct {
	Meta                   *requestmeta.RequestMeta                       `protobuf:"bytes,1,opt,name=meta" json:"meta,omitempty"`
	ServiceFilter          *servicefilter.ServiceFilter                   `protobuf:"bytes,2,opt,name=serviceFilter" json:"serviceFilter,omitempty"`
	CreateCompaniesRequest *createcompaniesrequest.CreateCompaniesRequest `protobuf:"bytes,3,opt,name=createCompaniesRequest" json:"createCompaniesRequest,omitempty"`
	Hash                   *string                                        `protobuf:"bytes,4,opt,name=hash" json:"hash,omitempty"`
	XXX_unrecognized       []byte                                         `json:"-"`
}

func (m *PassCreateCompaniesRequestRequest) Reset()         { *m = PassCreateCompaniesRequestRequest{} }
func (m *PassCreateCompaniesRequestRequest) String() string { return proto.CompactTextString(m) }
func (*PassCreateCompaniesRequestRequest) ProtoMessage()    {}
func (*PassCreateCompaniesRequestRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{0}
}

func (m *PassCreateCompaniesRequestRequest) GetMeta() *requestmeta.RequestMeta {
	if m != nil {
		return m.Meta
	}
	return nil
}

func (m *PassCreateCompaniesRequestRequest) GetServiceFilter() *servicefilter.ServiceFilter {
	if m != nil {
		return m.ServiceFilter
	}
	return nil
}

func (m *PassCreateCompaniesRequestRequest) GetCreateCompaniesRequest() *createcompaniesrequest.CreateCompaniesRequest {
	if m != nil {
		return m.CreateCompaniesRequest
	}
	return nil
}

func (m *PassCreateCompaniesRequestRequest) GetHash() string {
	if m != nil && m.Hash != nil {
		return *m.Hash
	}
	return ""
}

func init() {
	proto.RegisterType((*PassCreateCompaniesRequestRequest)(nil), "passcreatecompaniesrequestrequest.PassCreateCompaniesRequestRequest")
}

func init() {
	proto.RegisterFile("passcreatecompaniesrequestrequest/passcreatecompaniesrequestrequest_.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 263 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0xc1, 0x4a, 0x03, 0x31,
	0x10, 0x86, 0xd9, 0xba, 0x17, 0x23, 0x5e, 0x72, 0x90, 0x50, 0x44, 0xda, 0x9e, 0x7a, 0xd0, 0x04,
	0x8b, 0x4f, 0xd0, 0x82, 0x07, 0xa1, 0x28, 0xeb, 0xcd, 0x8b, 0x8c, 0x71, 0xba, 0x1b, 0xe8, 0x36,
	0x31, 0x93, 0x15, 0x7c, 0x0c, 0xdf, 0x58, 0x4c, 0x53, 0xd8, 0x85, 0x5d, 0xf6, 0x34, 0x93, 0xf0,
	0xcd, 0x9f, 0x6f, 0x20, 0xec, 0xc9, 0x01, 0x91, 0xf6, 0x08, 0x01, 0xb5, 0xad, 0x1d, 0x1c, 0x0c,
	0x92, 0xc7, 0xaf, 0x06, 0x29, 0xa4, 0xa2, 0x46, 0x89, 0x77, 0xe9, 0xbc, 0x0d, 0x96, 0xcf, 0x47,
	0xc9, 0xe9, 0x4d, 0x6a, 0x6a, 0x0c, 0xa0, 0x5a, 0x7d, 0x8a, 0x98, 0x2e, 0x08, 0xfd, 0xb7, 0xd1,
	0xb8, 0x33, 0xfb, 0x80, 0x5e, 0x75, 0x4e, 0x27, 0xe6, 0xa1, 0xff, 0x09, 0xd5, 0x7f, 0x9d, 0xa6,
	0x16, 0xbf, 0x13, 0x36, 0x7f, 0x01, 0xa2, 0x4d, 0xa4, 0x36, 0x27, 0xaa, 0x38, 0x52, 0xa9, 0xf0,
	0x5b, 0x96, 0xff, 0xeb, 0x88, 0x6c, 0x96, 0x2d, 0x2f, 0x56, 0x42, 0xb6, 0x14, 0x65, 0x62, 0xb6,
	0x18, 0xa0, 0x88, 0x14, 0x5f, 0xb3, 0xcb, 0x64, 0xf8, 0x18, 0x0d, 0xc5, 0x24, 0x8e, 0x5d, 0xcb,
	0x8e, 0xb7, 0x7c, 0x6d, 0x33, 0x45, 0x77, 0x84, 0xef, 0xd8, 0x95, 0xee, 0x55, 0x12, 0x67, 0x31,
	0x4c, 0xca, 0xfe, 0xbd, 0xe4, 0xc0, 0x22, 0x03, 0x69, 0x9c, 0xb3, 0xbc, 0x02, 0xaa, 0x44, 0x3e,
	0xcb, 0x96, 0xe7, 0x45, 0xec, 0xd7, 0xcf, 0x6f, 0xdb, 0xd2, 0x84, 0xaa, 0xf9, 0x90, 0xda, 0xd6,
	0x6a, 0x75, 0x4f, 0xc1, 0x58, 0x55, 0xda, 0x3b, 0xb7, 0x87, 0x9f, 0xd2, 0xdb, 0xe6, 0xf0, 0xa9,
	0x4a, 0xef, 0xb4, 0x22, 0x5d, 0x61, 0x0d, 0xe3, 0x1f, 0xe1, 0x2f, 0x00, 0x00, 0xff, 0xff, 0xca,
	0x99, 0x43, 0xfd, 0x4e, 0x02, 0x00, 0x00,
}
