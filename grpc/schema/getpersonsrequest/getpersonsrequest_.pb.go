// Code generated by protoc-gen-go. DO NOT EDIT.
// source: getpersonsrequest/getpersonsrequest_.proto

/*
Package getpersonsrequest is a generated protocol buffer package.

It is generated from these files:
	getpersonsrequest/getpersonsrequest_.proto

It has these top-level messages:
	GetPersonsRequest
*/
package getpersonsrequest

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import requestmeta "github.com/21stio/go-playground/grpc/schema/requestmeta"
import servicefilter "github.com/21stio/go-playground/grpc/schema/servicefilter"
import cache "github.com/21stio/go-playground/grpc/schema/cache"
import locationquery "github.com/21stio/go-playground/grpc/schema/locationquery"
import getpersonsresponseselect "github.com/21stio/go-playground/grpc/schema/getpersonsresponseselect"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type GetPersonsRequest struct {
	Meta             *requestmeta.RequestMeta                           `protobuf:"bytes,1,opt,name=meta" json:"meta,omitempty"`
	ServiceFilter    *servicefilter.ServiceFilter                       `protobuf:"bytes,2,opt,name=serviceFilter" json:"serviceFilter,omitempty"`
	Cache            *cache.Cache                                       `protobuf:"bytes,3,opt,name=cache" json:"cache,omitempty"`
	Query            *string                                            `protobuf:"bytes,4,opt,name=query" json:"query,omitempty"`
	LocationQuery    *locationquery.LocationQuery                       `protobuf:"bytes,5,opt,name=locationQuery" json:"locationQuery,omitempty"`
	Select           *getpersonsresponseselect.GetPersonsResponseSelect `protobuf:"bytes,6,opt,name=select" json:"select,omitempty"`
	Hash             *string                                            `protobuf:"bytes,7,opt,name=hash" json:"hash,omitempty"`
	XXX_unrecognized []byte                                             `json:"-"`
}

func (m *GetPersonsRequest) Reset()                    { *m = GetPersonsRequest{} }
func (m *GetPersonsRequest) String() string            { return proto.CompactTextString(m) }
func (*GetPersonsRequest) ProtoMessage()               {}
func (*GetPersonsRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *GetPersonsRequest) GetMeta() *requestmeta.RequestMeta {
	if m != nil {
		return m.Meta
	}
	return nil
}

func (m *GetPersonsRequest) GetServiceFilter() *servicefilter.ServiceFilter {
	if m != nil {
		return m.ServiceFilter
	}
	return nil
}

func (m *GetPersonsRequest) GetCache() *cache.Cache {
	if m != nil {
		return m.Cache
	}
	return nil
}

func (m *GetPersonsRequest) GetQuery() string {
	if m != nil && m.Query != nil {
		return *m.Query
	}
	return ""
}

func (m *GetPersonsRequest) GetLocationQuery() *locationquery.LocationQuery {
	if m != nil {
		return m.LocationQuery
	}
	return nil
}

func (m *GetPersonsRequest) GetSelect() *getpersonsresponseselect.GetPersonsResponseSelect {
	if m != nil {
		return m.Select
	}
	return nil
}

func (m *GetPersonsRequest) GetHash() string {
	if m != nil && m.Hash != nil {
		return *m.Hash
	}
	return ""
}

func init() {
	proto.RegisterType((*GetPersonsRequest)(nil), "getpersonsrequest.GetPersonsRequest")
}

func init() { proto.RegisterFile("getpersonsrequest/getpersonsrequest_.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 329 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x52, 0xc1, 0x4a, 0x03, 0x31,
	0x14, 0xa4, 0xb5, 0xad, 0x18, 0xf5, 0xd0, 0xe0, 0x21, 0x14, 0x91, 0xd2, 0x53, 0x11, 0x4d, 0xb0,
	0x27, 0x2f, 0x22, 0x54, 0x50, 0x10, 0x05, 0x4d, 0x6f, 0x5e, 0x24, 0xc6, 0xe7, 0xee, 0xc2, 0x76,
	0xb3, 0x4d, 0xb2, 0x42, 0xbf, 0xd7, 0x1f, 0x91, 0xbe, 0x4d, 0x25, 0xa1, 0x78, 0x09, 0xef, 0xed,
	0xcc, 0x6c, 0x66, 0x86, 0x90, 0xf3, 0x0c, 0x7c, 0x0d, 0xd6, 0x99, 0xca, 0x59, 0x58, 0x35, 0xe0,
	0xbc, 0xd8, 0xf9, 0xf2, 0xce, 0x6b, 0x6b, 0xbc, 0xa1, 0xc3, 0x1d, 0x64, 0x74, 0x16, 0x86, 0x25,
	0x78, 0x25, 0xa2, 0x39, 0x48, 0x46, 0x13, 0x07, 0xf6, 0xbb, 0xd0, 0xf0, 0x55, 0x94, 0x1e, 0xac,
	0x48, 0xb6, 0x2d, 0x87, 0x6a, 0xa5, 0x73, 0x10, 0x78, 0xfe, 0xe9, 0x4a, 0xa3, 0x95, 0x2f, 0x4c,
	0xb5, 0x6a, 0xc0, 0xae, 0x45, 0xb2, 0x6d, 0x39, 0xd7, 0xb1, 0x1d, 0x57, 0x9b, 0xca, 0x81, 0x83,
	0x12, 0x74, 0x9a, 0x20, 0x06, 0x82, 0x72, 0xf2, 0xd3, 0x25, 0xc3, 0x07, 0xf0, 0x2f, 0x2d, 0x47,
	0xb6, 0xb6, 0xe9, 0x05, 0xe9, 0x6d, 0xac, 0xb3, 0xce, 0xb8, 0x33, 0x3d, 0x9c, 0x31, 0x1e, 0xc5,
	0xe1, 0x81, 0xf3, 0x0c, 0x5e, 0x49, 0x64, 0xd1, 0x39, 0x39, 0x0e, 0x69, 0xee, 0x31, 0x0d, 0xeb,
	0xa2, 0xec, 0x94, 0x27, 0x19, 0xf9, 0x22, 0xe6, 0xc8, 0x54, 0x42, 0x27, 0xa4, 0x8f, 0xa9, 0xd9,
	0x1e, 0x6a, 0x8f, 0x38, 0x6e, 0xfc, 0x6e, 0x73, 0xca, 0x16, 0xa2, 0x27, 0xa4, 0x8f, 0xa9, 0x59,
	0x6f, 0xdc, 0x99, 0x1e, 0xc8, 0x76, 0xd9, 0xdc, 0xbe, 0xed, 0xe4, 0x15, 0xd1, 0x7e, 0xb8, 0x3d,
	0x69, 0x8a, 0x3f, 0xc5, 0x1c, 0x99, 0x4a, 0xe8, 0x23, 0x19, 0xb4, 0xb5, 0xb0, 0x01, 0x8a, 0x67,
	0xfc, 0xbf, 0xde, 0x78, 0x5c, 0x56, 0x0b, 0x2c, 0x10, 0x90, 0xe1, 0x0f, 0x94, 0x92, 0x5e, 0xae,
	0x5c, 0xce, 0xf6, 0xd1, 0x24, 0xce, 0xf3, 0xdb, 0xb7, 0x9b, 0xac, 0xf0, 0x79, 0xf3, 0xc1, 0xb5,
	0x59, 0x8a, 0xd9, 0x95, 0xf3, 0x85, 0x11, 0x99, 0xb9, 0xac, 0x4b, 0xb5, 0xce, 0xac, 0x69, 0xaa,
	0x4f, 0x91, 0xd9, 0x5a, 0x0b, 0xa7, 0x73, 0x58, 0xaa, 0xdd, 0x67, 0xf7, 0x1b, 0x00, 0x00, 0xff,
	0xff, 0xc9, 0x4f, 0x04, 0xef, 0x9c, 0x02, 0x00, 0x00,
}