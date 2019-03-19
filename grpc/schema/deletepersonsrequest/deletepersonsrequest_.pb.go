// Code generated by protoc-gen-go. DO NOT EDIT.
// source: deletepersonsrequest/deletepersonsrequest_.proto

/*
Package deletepersonsrequest is a generated protocol buffer package.

It is generated from these files:
	deletepersonsrequest/deletepersonsrequest_.proto

It has these top-level messages:
	DeletePersonsRequest
*/
package deletepersonsrequest

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import requestmeta "github.com/21stio/go-playground/grpc/schema/requestmeta"
import servicefilter "github.com/21stio/go-playground/grpc/schema/servicefilter"
import id "github.com/21stio/go-playground/grpc/schema/id"
import deletepersonsresponseselect "github.com/21stio/go-playground/grpc/schema/deletepersonsresponseselect"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type DeletePersonsRequest struct {
	Meta             *requestmeta.RequestMeta                                 `protobuf:"bytes,1,opt,name=meta" json:"meta,omitempty"`
	ServiceFilter    *servicefilter.ServiceFilter                             `protobuf:"bytes,2,opt,name=serviceFilter" json:"serviceFilter,omitempty"`
	Ids              []*id.Id                                                 `protobuf:"bytes,3,rep,name=ids" json:"ids,omitempty"`
	Select           *deletepersonsresponseselect.DeletePersonsResponseSelect `protobuf:"bytes,4,opt,name=select" json:"select,omitempty"`
	Hash             *string                                                  `protobuf:"bytes,5,opt,name=hash" json:"hash,omitempty"`
	XXX_unrecognized []byte                                                   `json:"-"`
}

func (m *DeletePersonsRequest) Reset()                    { *m = DeletePersonsRequest{} }
func (m *DeletePersonsRequest) String() string            { return proto.CompactTextString(m) }
func (*DeletePersonsRequest) ProtoMessage()               {}
func (*DeletePersonsRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *DeletePersonsRequest) GetMeta() *requestmeta.RequestMeta {
	if m != nil {
		return m.Meta
	}
	return nil
}

func (m *DeletePersonsRequest) GetServiceFilter() *servicefilter.ServiceFilter {
	if m != nil {
		return m.ServiceFilter
	}
	return nil
}

func (m *DeletePersonsRequest) GetIds() []*id.Id {
	if m != nil {
		return m.Ids
	}
	return nil
}

func (m *DeletePersonsRequest) GetSelect() *deletepersonsresponseselect.DeletePersonsResponseSelect {
	if m != nil {
		return m.Select
	}
	return nil
}

func (m *DeletePersonsRequest) GetHash() string {
	if m != nil && m.Hash != nil {
		return *m.Hash
	}
	return ""
}

func init() {
	proto.RegisterType((*DeletePersonsRequest)(nil), "deletepersonsrequest.DeletePersonsRequest")
}

func init() { proto.RegisterFile("deletepersonsrequest/deletepersonsrequest_.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 292 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x91, 0xc1, 0x4b, 0xc3, 0x30,
	0x14, 0xc6, 0xd9, 0x5a, 0x07, 0x66, 0x7a, 0x09, 0x3b, 0x84, 0x21, 0x52, 0x76, 0xea, 0x41, 0x13,
	0xed, 0xc9, 0x93, 0x48, 0x11, 0xc1, 0x83, 0x30, 0xb2, 0x9b, 0x17, 0xa9, 0xcd, 0xb3, 0x0d, 0xb4,
	0x4d, 0x4d, 0x52, 0xc1, 0xff, 0xc0, 0x3f, 0x5b, 0x96, 0x66, 0xd0, 0xce, 0xe2, 0xed, 0xbd, 0xf7,
	0x7d, 0xdf, 0xcb, 0xef, 0x11, 0x74, 0x23, 0xa0, 0x02, 0x0b, 0x2d, 0x68, 0xa3, 0x1a, 0xa3, 0xe1,
	0xb3, 0x03, 0x63, 0xd9, 0xd4, 0xf0, 0x8d, 0xb6, 0x5a, 0x59, 0x85, 0x57, 0x53, 0xe2, 0xfa, 0xd2,
	0x17, 0x35, 0xd8, 0x8c, 0x0d, 0x6a, 0x9f, 0x5a, 0x6f, 0x0c, 0xe8, 0x2f, 0x99, 0xc3, 0x87, 0xac,
	0x2c, 0x68, 0x36, 0xea, 0x0e, 0x9e, 0x33, 0x29, 0x98, 0x14, 0x87, 0xee, 0xfe, 0xe8, 0x1d, 0xd3,
	0xaa, 0xc6, 0x80, 0x81, 0x0a, 0xf2, 0x3f, 0x80, 0x43, 0xcd, 0xe7, 0x37, 0x3f, 0x73, 0xb4, 0x7a,
	0x74, 0xb6, 0x6d, 0x6f, 0xe3, 0x3d, 0x15, 0xbe, 0x42, 0xe1, 0x9e, 0x8c, 0xcc, 0xa2, 0x59, 0xbc,
	0x4c, 0x08, 0x1d, 0xd0, 0x52, 0xef, 0x79, 0x01, 0x9b, 0x71, 0xe7, 0xc2, 0x29, 0x3a, 0xf7, 0xb0,
	0x4f, 0x0e, 0x96, 0xcc, 0x5d, 0xec, 0x82, 0x8e, 0x4e, 0xa0, 0xbb, 0xa1, 0x87, 0x8f, 0x23, 0x98,
	0xa0, 0x40, 0x0a, 0x43, 0x82, 0x28, 0x88, 0x97, 0xc9, 0x82, 0x4a, 0x41, 0x9f, 0x05, 0xdf, 0x8f,
	0xf0, 0x16, 0x2d, 0x7a, 0x6a, 0x12, 0xba, 0xb5, 0x77, 0xf4, 0x9f, 0xcb, 0xe8, 0xd1, 0x39, 0xbd,
	0xb6, 0x73, 0x1a, 0xf7, 0x7b, 0x30, 0x46, 0x61, 0x99, 0x99, 0x92, 0x9c, 0x44, 0xb3, 0xf8, 0x94,
	0xbb, 0x3a, 0x4d, 0x5f, 0x1f, 0x0a, 0x69, 0xcb, 0xee, 0x9d, 0xe6, 0xaa, 0x66, 0xc9, 0xad, 0xb1,
	0x52, 0xb1, 0x42, 0x5d, 0xb7, 0x55, 0xf6, 0x5d, 0x68, 0xd5, 0x35, 0x82, 0x15, 0xba, 0xcd, 0x99,
	0xc9, 0x4b, 0xa8, 0xb3, 0xc9, 0xdf, 0xff, 0x0d, 0x00, 0x00, 0xff, 0xff, 0x40, 0xaa, 0xca, 0x89,
	0x29, 0x02, 0x00, 0x00,
}