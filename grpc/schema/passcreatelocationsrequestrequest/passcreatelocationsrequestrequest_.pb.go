// Code generated by protoc-gen-go. DO NOT EDIT.
// source: passcreatelocationsrequestrequest/passcreatelocationsrequestrequest_.proto

/*
Package passcreatelocationsrequestrequest is a generated protocol buffer package.

It is generated from these files:
	passcreatelocationsrequestrequest/passcreatelocationsrequestrequest_.proto

It has these top-level messages:
	PassCreateLocationsRequestRequest
*/
package passcreatelocationsrequestrequest

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import requestmeta "github.com/21stio/go-playground/grpc/schema/requestmeta"
import servicefilter "github.com/21stio/go-playground/grpc/schema/servicefilter"
import createlocationsrequest "github.com/21stio/go-playground/grpc/schema/createlocationsrequest"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type PassCreateLocationsRequestRequest struct {
	Meta                   *requestmeta.RequestMeta                       `protobuf:"bytes,1,opt,name=meta" json:"meta,omitempty"`
	ServiceFilter          *servicefilter.ServiceFilter                   `protobuf:"bytes,2,opt,name=serviceFilter" json:"serviceFilter,omitempty"`
	CreateLocationsRequest *createlocationsrequest.CreateLocationsRequest `protobuf:"bytes,3,opt,name=createLocationsRequest" json:"createLocationsRequest,omitempty"`
	Hash                   *string                                        `protobuf:"bytes,4,opt,name=hash" json:"hash,omitempty"`
	XXX_unrecognized       []byte                                         `json:"-"`
}

func (m *PassCreateLocationsRequestRequest) Reset()         { *m = PassCreateLocationsRequestRequest{} }
func (m *PassCreateLocationsRequestRequest) String() string { return proto.CompactTextString(m) }
func (*PassCreateLocationsRequestRequest) ProtoMessage()    {}
func (*PassCreateLocationsRequestRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{0}
}

func (m *PassCreateLocationsRequestRequest) GetMeta() *requestmeta.RequestMeta {
	if m != nil {
		return m.Meta
	}
	return nil
}

func (m *PassCreateLocationsRequestRequest) GetServiceFilter() *servicefilter.ServiceFilter {
	if m != nil {
		return m.ServiceFilter
	}
	return nil
}

func (m *PassCreateLocationsRequestRequest) GetCreateLocationsRequest() *createlocationsrequest.CreateLocationsRequest {
	if m != nil {
		return m.CreateLocationsRequest
	}
	return nil
}

func (m *PassCreateLocationsRequestRequest) GetHash() string {
	if m != nil && m.Hash != nil {
		return *m.Hash
	}
	return ""
}

func init() {
	proto.RegisterType((*PassCreateLocationsRequestRequest)(nil), "passcreatelocationsrequestrequest.PassCreateLocationsRequestRequest")
}

func init() {
	proto.RegisterFile("passcreatelocationsrequestrequest/passcreatelocationsrequestrequest_.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 264 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0x41, 0x4b, 0xc3, 0x30,
	0x14, 0xc7, 0xe9, 0xec, 0xc5, 0x88, 0x97, 0x1c, 0xa4, 0x0c, 0x91, 0x6d, 0xa7, 0x1d, 0x34, 0xc1,
	0xe1, 0x27, 0x98, 0xe0, 0x41, 0x1c, 0x4a, 0xbd, 0x79, 0x91, 0x67, 0x7c, 0x6b, 0x03, 0xed, 0x52,
	0xf3, 0x5e, 0x05, 0x3f, 0x86, 0xdf, 0x58, 0xcc, 0x32, 0x68, 0xa1, 0xa3, 0xa7, 0xf7, 0x12, 0x7e,
	0xef, 0x9f, 0xdf, 0x83, 0x88, 0xc7, 0x06, 0x88, 0x8c, 0x47, 0x60, 0xac, 0x9c, 0x01, 0xb6, 0x6e,
	0x47, 0x1e, 0xbf, 0x5a, 0x24, 0x8e, 0x45, 0x8f, 0x12, 0xef, 0xaa, 0xf1, 0x8e, 0x9d, 0x9c, 0x8f,
	0x92, 0xd3, 0xab, 0xd8, 0xd4, 0xc8, 0xa0, 0x3b, 0x7d, 0x8c, 0x98, 0x2e, 0x08, 0xfd, 0xb7, 0x35,
	0xb8, 0xb5, 0x15, 0xa3, 0xd7, 0xbd, 0xd3, 0x81, 0xb9, 0x1b, 0x7e, 0x42, 0x0f, 0x5f, 0xc7, 0xa9,
	0xc5, 0xef, 0x44, 0xcc, 0x5f, 0x80, 0xe8, 0x3e, 0x50, 0x4f, 0x07, 0x2a, 0xdf, 0x53, 0xb1, 0xc8,
	0x6b, 0x91, 0xfe, 0xeb, 0x64, 0xc9, 0x2c, 0x59, 0x9e, 0xad, 0x32, 0xd5, 0x51, 0x54, 0x91, 0xd9,
	0x20, 0x43, 0x1e, 0x28, 0xb9, 0x16, 0xe7, 0xd1, 0xf0, 0x21, 0x18, 0x66, 0x93, 0x30, 0x76, 0xa9,
	0x7a, 0xde, 0xea, 0xb5, 0xcb, 0xe4, 0xfd, 0x11, 0xb9, 0x15, 0x17, 0x66, 0x50, 0x29, 0x3b, 0x09,
	0x61, 0x4a, 0x0d, 0xef, 0xa5, 0x8e, 0x2c, 0x72, 0x24, 0x4d, 0x4a, 0x91, 0x96, 0x40, 0x65, 0x96,
	0xce, 0x92, 0xe5, 0x69, 0x1e, 0xfa, 0xf5, 0xf3, 0xdb, 0xa6, 0xb0, 0x5c, 0xb6, 0x1f, 0xca, 0xb8,
	0x5a, 0xaf, 0x6e, 0x89, 0xad, 0xd3, 0x85, 0xbb, 0x69, 0x2a, 0xf8, 0x29, 0xbc, 0x6b, 0x77, 0x9f,
	0xba, 0xf0, 0x8d, 0xd1, 0x64, 0x4a, 0xac, 0x61, 0xfc, 0x23, 0xfc, 0x05, 0x00, 0x00, 0xff, 0xff,
	0x57, 0x22, 0x0c, 0x10, 0x4e, 0x02, 0x00, 0x00,
}