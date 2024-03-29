// Code generated by protoc-gen-go. DO NOT EDIT.
// source: createspecdescriptionsrequest/createspecdescriptionsrequest_.proto

/*
Package createspecdescriptionsrequest is a generated protocol buffer package.

It is generated from these files:
	createspecdescriptionsrequest/createspecdescriptionsrequest_.proto

It has these top-level messages:
	CreateSpecDescriptionsRequest
*/
package createspecdescriptionsrequest

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import requestmeta "github.com/21stio/go-playground/grpc/schema/requestmeta"
import servicefilter "github.com/21stio/go-playground/grpc/schema/servicefilter"
import specdescription "github.com/21stio/go-playground/grpc/schema/specdescription"
import createspecdescriptionsresponseselect "github.com/21stio/go-playground/grpc/schema/createspecdescriptionsresponseselect"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type CreateSpecDescriptionsRequest struct {
	Meta             *requestmeta.RequestMeta                                                   `protobuf:"bytes,1,opt,name=meta" json:"meta,omitempty"`
	ServiceFilter    *servicefilter.ServiceFilter                                               `protobuf:"bytes,2,opt,name=serviceFilter" json:"serviceFilter,omitempty"`
	SpecDescriptions []*specdescription.SpecDescription                                         `protobuf:"bytes,3,rep,name=specDescriptions" json:"specDescriptions,omitempty"`
	Select           *createspecdescriptionsresponseselect.CreateSpecDescriptionsResponseSelect `protobuf:"bytes,4,opt,name=select" json:"select,omitempty"`
	Hash             *string                                                                    `protobuf:"bytes,5,opt,name=hash" json:"hash,omitempty"`
	XXX_unrecognized []byte                                                                     `json:"-"`
}

func (m *CreateSpecDescriptionsRequest) Reset()                    { *m = CreateSpecDescriptionsRequest{} }
func (m *CreateSpecDescriptionsRequest) String() string            { return proto.CompactTextString(m) }
func (*CreateSpecDescriptionsRequest) ProtoMessage()               {}
func (*CreateSpecDescriptionsRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *CreateSpecDescriptionsRequest) GetMeta() *requestmeta.RequestMeta {
	if m != nil {
		return m.Meta
	}
	return nil
}

func (m *CreateSpecDescriptionsRequest) GetServiceFilter() *servicefilter.ServiceFilter {
	if m != nil {
		return m.ServiceFilter
	}
	return nil
}

func (m *CreateSpecDescriptionsRequest) GetSpecDescriptions() []*specdescription.SpecDescription {
	if m != nil {
		return m.SpecDescriptions
	}
	return nil
}

func (m *CreateSpecDescriptionsRequest) GetSelect() *createspecdescriptionsresponseselect.CreateSpecDescriptionsResponseSelect {
	if m != nil {
		return m.Select
	}
	return nil
}

func (m *CreateSpecDescriptionsRequest) GetHash() string {
	if m != nil && m.Hash != nil {
		return *m.Hash
	}
	return ""
}

func init() {
	proto.RegisterType((*CreateSpecDescriptionsRequest)(nil), "createspecdescriptionsrequest.CreateSpecDescriptionsRequest")
}

func init() {
	proto.RegisterFile("createspecdescriptionsrequest/createspecdescriptionsrequest_.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 305 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xcd, 0x4a, 0x33, 0x31,
	0x14, 0x86, 0xe9, 0xcf, 0xf7, 0x81, 0x29, 0x82, 0x64, 0x15, 0x8a, 0x95, 0xa1, 0x0b, 0xe9, 0x42,
	0x13, 0xec, 0x25, 0x54, 0x11, 0xfc, 0x03, 0x49, 0x77, 0x6e, 0x24, 0x4d, 0x8f, 0x33, 0x81, 0xb6,
	0x89, 0x39, 0xa9, 0xe0, 0x95, 0x7a, 0x3b, 0x62, 0x26, 0x85, 0x99, 0x91, 0x19, 0xdc, 0x9d, 0x30,
	0xcf, 0x7b, 0x78, 0x9f, 0xc3, 0x90, 0x85, 0xf6, 0xa0, 0x02, 0xa0, 0x03, 0xbd, 0x06, 0xd4, 0xde,
	0xb8, 0x60, 0xec, 0x0e, 0x3d, 0xbc, 0xef, 0x01, 0x83, 0xe8, 0xfc, 0xfa, 0xca, 0x9d, 0xb7, 0xc1,
	0xd2, 0x49, 0x27, 0x35, 0x3e, 0x4b, 0xc3, 0x16, 0x82, 0x12, 0x95, 0x39, 0xc5, 0xc7, 0x53, 0x04,
	0xff, 0x61, 0x34, 0xbc, 0x99, 0x4d, 0x00, 0x2f, 0x6a, 0xaf, 0x03, 0x73, 0xde, 0x58, 0x2e, 0x1a,
	0xef, 0x03, 0xf7, 0xdc, 0x56, 0x05, 0x9d, 0xdd, 0x21, 0x20, 0x6c, 0x40, 0xb7, 0x5b, 0x55, 0xa1,
	0xb4, 0x71, 0xfa, 0xd5, 0x27, 0x93, 0xeb, 0xc8, 0x2f, 0x1d, 0xe8, 0x9b, 0x0a, 0x2f, 0x4b, 0x15,
	0x7a, 0x41, 0x86, 0x3f, 0x3a, 0xac, 0x97, 0xf5, 0x66, 0xa3, 0x39, 0xe3, 0x15, 0x45, 0x9e, 0x98,
	0x27, 0x08, 0x4a, 0x46, 0x8a, 0x2e, 0xc8, 0x71, 0x32, 0xbc, 0x8d, 0x86, 0xac, 0x1f, 0x63, 0xa7,
	0xbc, 0xe6, 0xcd, 0x97, 0x55, 0x46, 0xd6, 0x23, 0xf4, 0x91, 0x9c, 0x60, 0xa3, 0x0c, 0x1b, 0x64,
	0x83, 0xd9, 0x68, 0x9e, 0xf1, 0x86, 0x15, 0x6f, 0xb4, 0x96, 0xbf, 0x92, 0x74, 0x45, 0xfe, 0x97,
	0xca, 0x6c, 0x18, 0xab, 0xdc, 0xf3, 0xbf, 0xdc, 0x87, 0xb7, 0x1d, 0xa5, 0x84, 0x96, 0x11, 0x92,
	0x69, 0x33, 0xa5, 0x64, 0x58, 0x28, 0x2c, 0xd8, 0xbf, 0xac, 0x37, 0x3b, 0x92, 0x71, 0x5e, 0x3c,
	0xbc, 0xdc, 0xe5, 0x26, 0x14, 0xfb, 0x15, 0xd7, 0x76, 0x2b, 0xe6, 0x57, 0x18, 0x8c, 0x15, 0xb9,
	0xbd, 0x74, 0x1b, 0xf5, 0x99, 0x7b, 0xbb, 0xdf, 0xad, 0x45, 0xee, 0x9d, 0x16, 0xa8, 0x0b, 0xd8,
	0xaa, 0xee, 0x5f, 0xf1, 0x3b, 0x00, 0x00, 0xff, 0xff, 0xab, 0x2b, 0x1b, 0x02, 0xc8, 0x02, 0x00,
	0x00,
}
