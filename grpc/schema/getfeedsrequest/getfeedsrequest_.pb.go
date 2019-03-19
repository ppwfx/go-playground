// Code generated by protoc-gen-go. DO NOT EDIT.
// source: getfeedsrequest/getfeedsrequest_.proto

/*
Package getfeedsrequest is a generated protocol buffer package.

It is generated from these files:
	getfeedsrequest/getfeedsrequest_.proto

It has these top-level messages:
	GetFeedsRequest
*/
package getfeedsrequest

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import requestmeta "github.com/21stio/go-playground/grpc/schema/requestmeta"
import servicefilter "github.com/21stio/go-playground/grpc/schema/servicefilter"
import cache "github.com/21stio/go-playground/grpc/schema/cache"
import locationquery "github.com/21stio/go-playground/grpc/schema/locationquery"
import getfeedsresponseselect "github.com/21stio/go-playground/grpc/schema/getfeedsresponseselect"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type GetFeedsRequest struct {
	Meta             *requestmeta.RequestMeta                       `protobuf:"bytes,1,opt,name=meta" json:"meta,omitempty"`
	ServiceFilter    *servicefilter.ServiceFilter                   `protobuf:"bytes,2,opt,name=serviceFilter" json:"serviceFilter,omitempty"`
	Cache            *cache.Cache                                   `protobuf:"bytes,3,opt,name=cache" json:"cache,omitempty"`
	Query            *string                                        `protobuf:"bytes,4,opt,name=query" json:"query,omitempty"`
	LocationQuery    *locationquery.LocationQuery                   `protobuf:"bytes,5,opt,name=locationQuery" json:"locationQuery,omitempty"`
	Select           *getfeedsresponseselect.GetFeedsResponseSelect `protobuf:"bytes,6,opt,name=select" json:"select,omitempty"`
	Hash             *string                                        `protobuf:"bytes,7,opt,name=hash" json:"hash,omitempty"`
	XXX_unrecognized []byte                                         `json:"-"`
}

func (m *GetFeedsRequest) Reset()                    { *m = GetFeedsRequest{} }
func (m *GetFeedsRequest) String() string            { return proto.CompactTextString(m) }
func (*GetFeedsRequest) ProtoMessage()               {}
func (*GetFeedsRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *GetFeedsRequest) GetMeta() *requestmeta.RequestMeta {
	if m != nil {
		return m.Meta
	}
	return nil
}

func (m *GetFeedsRequest) GetServiceFilter() *servicefilter.ServiceFilter {
	if m != nil {
		return m.ServiceFilter
	}
	return nil
}

func (m *GetFeedsRequest) GetCache() *cache.Cache {
	if m != nil {
		return m.Cache
	}
	return nil
}

func (m *GetFeedsRequest) GetQuery() string {
	if m != nil && m.Query != nil {
		return *m.Query
	}
	return ""
}

func (m *GetFeedsRequest) GetLocationQuery() *locationquery.LocationQuery {
	if m != nil {
		return m.LocationQuery
	}
	return nil
}

func (m *GetFeedsRequest) GetSelect() *getfeedsresponseselect.GetFeedsResponseSelect {
	if m != nil {
		return m.Select
	}
	return nil
}

func (m *GetFeedsRequest) GetHash() string {
	if m != nil && m.Hash != nil {
		return *m.Hash
	}
	return ""
}

func init() {
	proto.RegisterType((*GetFeedsRequest)(nil), "getfeedsrequest.GetFeedsRequest")
}

func init() { proto.RegisterFile("getfeedsrequest/getfeedsrequest_.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 329 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x52, 0x41, 0x4b, 0x33, 0x31,
	0x14, 0xa4, 0xfd, 0xda, 0x7e, 0x18, 0x95, 0x42, 0xf0, 0x10, 0x8a, 0x48, 0xe9, 0x41, 0x7a, 0xd0,
	0x04, 0x8b, 0x37, 0xf1, 0x52, 0xa1, 0x5e, 0xf4, 0x60, 0x7a, 0xf3, 0x22, 0x31, 0x7d, 0xdd, 0x5d,
	0xd8, 0x36, 0xdb, 0x24, 0x2b, 0xf4, 0xc7, 0xfa, 0x5f, 0x64, 0xdf, 0xa6, 0x76, 0xb3, 0x78, 0x09,
	0x6f, 0x5e, 0x66, 0xc8, 0xcc, 0x10, 0x72, 0x9d, 0x80, 0x5f, 0x03, 0xac, 0x9c, 0x85, 0x5d, 0x09,
	0xce, 0x8b, 0x16, 0xfe, 0xe0, 0x85, 0x35, 0xde, 0xd0, 0x61, 0x6b, 0x3f, 0xba, 0x0a, 0xc3, 0x06,
	0xbc, 0x12, 0x8d, 0x39, 0x08, 0x46, 0x13, 0x07, 0xf6, 0x2b, 0xd3, 0xb0, 0xce, 0x72, 0x0f, 0x56,
	0x44, 0xe8, 0xc0, 0xa1, 0x5a, 0xe9, 0x14, 0x04, 0x9e, 0xbf, 0xba, 0xdc, 0x68, 0xe5, 0x33, 0xb3,
	0xdd, 0x95, 0x60, 0xf7, 0x22, 0x42, 0x07, 0xce, 0xfd, 0xd1, 0x8c, 0x2b, 0xcc, 0xd6, 0x81, 0x83,
	0x1c, 0x74, 0xd3, 0x7b, 0x73, 0x1d, 0x54, 0x93, 0xef, 0x2e, 0x19, 0x3e, 0x83, 0x5f, 0x54, 0x0c,
	0x59, 0x1b, 0xa6, 0x37, 0xa4, 0x57, 0x99, 0x66, 0x9d, 0x71, 0x67, 0x7a, 0x3a, 0x63, 0xbc, 0x11,
	0x84, 0x07, 0xce, 0x2b, 0x78, 0x25, 0x91, 0x45, 0xe7, 0xe4, 0x3c, 0xe4, 0x58, 0x60, 0x0e, 0xd6,
	0x45, 0xd9, 0x25, 0x8f, 0xd2, 0xf1, 0x65, 0x93, 0x23, 0x63, 0x09, 0x9d, 0x90, 0x3e, 0xe6, 0x65,
	0xff, 0x50, 0x7b, 0xc6, 0x11, 0xf1, 0xa7, 0xea, 0x94, 0xf5, 0x15, 0xbd, 0x20, 0x7d, 0xcc, 0xcb,
	0x7a, 0xe3, 0xce, 0xf4, 0x44, 0xd6, 0xa0, 0x7a, 0xfd, 0xd0, 0xc6, 0x1b, 0xde, 0xf6, 0xc3, 0xeb,
	0x51, 0x47, 0xfc, 0xa5, 0xc9, 0x91, 0xb1, 0x84, 0x2e, 0xc8, 0xa0, 0x2e, 0x85, 0x0d, 0x50, 0xcc,
	0xf9, 0xdf, 0x9d, 0xf1, 0x63, 0x51, 0xf5, 0x7a, 0x89, 0x6b, 0x19, 0xd4, 0x94, 0x92, 0x5e, 0xaa,
	0x5c, 0xca, 0xfe, 0xa3, 0x41, 0x9c, 0xe7, 0x8f, 0xef, 0x0f, 0x49, 0xe6, 0xd3, 0xf2, 0x93, 0x6b,
	0xb3, 0x11, 0xb3, 0x3b, 0xe7, 0x33, 0x23, 0x12, 0x73, 0x5b, 0xe4, 0x6a, 0x9f, 0x58, 0x53, 0x6e,
	0x57, 0x22, 0xb1, 0x85, 0x16, 0x4e, 0xa7, 0xb0, 0x51, 0xed, 0x8f, 0xf6, 0x13, 0x00, 0x00, 0xff,
	0xff, 0xd2, 0x1c, 0x60, 0x31, 0x8a, 0x02, 0x00, 0x00,
}