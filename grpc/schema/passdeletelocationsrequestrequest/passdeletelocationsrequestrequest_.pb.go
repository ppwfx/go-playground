// Code generated by protoc-gen-go. DO NOT EDIT.
// source: passdeletelocationsrequestrequest/passdeletelocationsrequestrequest_.proto

/*
Package passdeletelocationsrequestrequest is a generated protocol buffer package.

It is generated from these files:
	passdeletelocationsrequestrequest/passdeletelocationsrequestrequest_.proto

It has these top-level messages:
	PassDeleteLocationsRequestRequest
*/
package passdeletelocationsrequestrequest

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import requestmeta "github.com/21stio/go-playground/grpc/schema/requestmeta"
import servicefilter "github.com/21stio/go-playground/grpc/schema/servicefilter"
import deletelocationsrequest "github.com/21stio/go-playground/grpc/schema/deletelocationsrequest"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type PassDeleteLocationsRequestRequest struct {
	Meta                   *requestmeta.RequestMeta                       `protobuf:"bytes,1,opt,name=meta" json:"meta,omitempty"`
	ServiceFilter          *servicefilter.ServiceFilter                   `protobuf:"bytes,2,opt,name=serviceFilter" json:"serviceFilter,omitempty"`
	DeleteLocationsRequest *deletelocationsrequest.DeleteLocationsRequest `protobuf:"bytes,3,opt,name=deleteLocationsRequest" json:"deleteLocationsRequest,omitempty"`
	Hash                   *string                                        `protobuf:"bytes,4,opt,name=hash" json:"hash,omitempty"`
	XXX_unrecognized       []byte                                         `json:"-"`
}

func (m *PassDeleteLocationsRequestRequest) Reset()         { *m = PassDeleteLocationsRequestRequest{} }
func (m *PassDeleteLocationsRequestRequest) String() string { return proto.CompactTextString(m) }
func (*PassDeleteLocationsRequestRequest) ProtoMessage()    {}
func (*PassDeleteLocationsRequestRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{0}
}

func (m *PassDeleteLocationsRequestRequest) GetMeta() *requestmeta.RequestMeta {
	if m != nil {
		return m.Meta
	}
	return nil
}

func (m *PassDeleteLocationsRequestRequest) GetServiceFilter() *servicefilter.ServiceFilter {
	if m != nil {
		return m.ServiceFilter
	}
	return nil
}

func (m *PassDeleteLocationsRequestRequest) GetDeleteLocationsRequest() *deletelocationsrequest.DeleteLocationsRequest {
	if m != nil {
		return m.DeleteLocationsRequest
	}
	return nil
}

func (m *PassDeleteLocationsRequestRequest) GetHash() string {
	if m != nil && m.Hash != nil {
		return *m.Hash
	}
	return ""
}

func init() {
	proto.RegisterType((*PassDeleteLocationsRequestRequest)(nil), "passdeletelocationsrequestrequest.PassDeleteLocationsRequestRequest")
}

func init() {
	proto.RegisterFile("passdeletelocationsrequestrequest/passdeletelocationsrequestrequest_.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 264 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0xc1, 0x4a, 0xc4, 0x30,
	0x10, 0x86, 0xe9, 0xda, 0x8b, 0x11, 0x2f, 0x39, 0x48, 0x58, 0x44, 0x76, 0xf7, 0xb4, 0x07, 0x4d,
	0x70, 0xf1, 0x09, 0x16, 0xf1, 0x20, 0x2e, 0x4a, 0xbd, 0x79, 0x91, 0x98, 0xce, 0xb6, 0x81, 0x76,
	0x53, 0x33, 0x53, 0xc1, 0xc7, 0xf0, 0x8d, 0xc5, 0x98, 0x85, 0x16, 0x5a, 0x7a, 0x9a, 0x49, 0xf8,
	0xe6, 0xcf, 0x37, 0x10, 0xf6, 0xd8, 0x68, 0xc4, 0x1c, 0x2a, 0x20, 0xa8, 0x9c, 0xd1, 0x64, 0xdd,
	0x01, 0x3d, 0x7c, 0xb6, 0x80, 0x14, 0x8b, 0x9a, 0x24, 0xde, 0x65, 0xe3, 0x1d, 0x39, 0xbe, 0x9c,
	0x24, 0xe7, 0x57, 0xb1, 0xa9, 0x81, 0xb4, 0xea, 0xf4, 0x31, 0x62, 0xbe, 0x42, 0xf0, 0x5f, 0xd6,
	0xc0, 0xde, 0x56, 0x04, 0x5e, 0xf5, 0x4e, 0x47, 0xe6, 0x6e, 0xf8, 0x09, 0x35, 0x7c, 0x1d, 0xa7,
	0x56, 0x3f, 0x33, 0xb6, 0x7c, 0xd1, 0x88, 0xf7, 0x81, 0x7a, 0x3a, 0x52, 0xd9, 0x3f, 0x15, 0x0b,
	0xbf, 0x66, 0xe9, 0x9f, 0x8e, 0x48, 0x16, 0xc9, 0xfa, 0x6c, 0x23, 0x64, 0x47, 0x51, 0x46, 0x66,
	0x07, 0xa4, 0xb3, 0x40, 0xf1, 0x2d, 0x3b, 0x8f, 0x86, 0x0f, 0xc1, 0x50, 0xcc, 0xc2, 0xd8, 0xa5,
	0xec, 0x79, 0xcb, 0xd7, 0x2e, 0x93, 0xf5, 0x47, 0xf8, 0x9e, 0x5d, 0xe4, 0x83, 0x4a, 0xe2, 0x24,
	0x84, 0x49, 0x39, 0xbc, 0x97, 0x1c, 0x59, 0x64, 0x24, 0x8d, 0x73, 0x96, 0x96, 0x1a, 0x4b, 0x91,
	0x2e, 0x92, 0xf5, 0x69, 0x16, 0xfa, 0xed, 0xf3, 0xdb, 0xae, 0xb0, 0x54, 0xb6, 0x1f, 0xd2, 0xb8,
	0x5a, 0x6d, 0x6e, 0x91, 0xac, 0x53, 0x85, 0xbb, 0x69, 0x2a, 0xfd, 0x5d, 0x78, 0xd7, 0x1e, 0x72,
	0x55, 0xf8, 0xc6, 0x28, 0x34, 0x25, 0xd4, 0x7a, 0xfa, 0x23, 0xfc, 0x06, 0x00, 0x00, 0xff, 0xff,
	0x60, 0x6b, 0x4d, 0x96, 0x4e, 0x02, 0x00, 0x00,
}
