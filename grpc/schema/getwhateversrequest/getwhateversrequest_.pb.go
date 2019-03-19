// Code generated by protoc-gen-go. DO NOT EDIT.
// source: getwhateversrequest/getwhateversrequest_.proto

/*
Package getwhateversrequest is a generated protocol buffer package.

It is generated from these files:
	getwhateversrequest/getwhateversrequest_.proto

It has these top-level messages:
	GetWhateversRequest
*/
package getwhateversrequest

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import requestmeta "github.com/21stio/go-playground/grpc/schema/requestmeta"
import servicefilter "github.com/21stio/go-playground/grpc/schema/servicefilter"
import cache "github.com/21stio/go-playground/grpc/schema/cache"
import locationquery "github.com/21stio/go-playground/grpc/schema/locationquery"
import getwhateversresponseselect "github.com/21stio/go-playground/grpc/schema/getwhateversresponseselect"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type GetWhateversRequest struct {
	Meta             *requestmeta.RequestMeta                               `protobuf:"bytes,1,opt,name=meta" json:"meta,omitempty"`
	ServiceFilter    *servicefilter.ServiceFilter                           `protobuf:"bytes,2,opt,name=serviceFilter" json:"serviceFilter,omitempty"`
	Cache            *cache.Cache                                           `protobuf:"bytes,3,opt,name=cache" json:"cache,omitempty"`
	Query            *string                                                `protobuf:"bytes,4,opt,name=query" json:"query,omitempty"`
	LocationQuery    *locationquery.LocationQuery                           `protobuf:"bytes,5,opt,name=locationQuery" json:"locationQuery,omitempty"`
	Select           *getwhateversresponseselect.GetWhateversResponseSelect `protobuf:"bytes,6,opt,name=select" json:"select,omitempty"`
	Hash             *string                                                `protobuf:"bytes,7,opt,name=hash" json:"hash,omitempty"`
	XXX_unrecognized []byte                                                 `json:"-"`
}

func (m *GetWhateversRequest) Reset()                    { *m = GetWhateversRequest{} }
func (m *GetWhateversRequest) String() string            { return proto.CompactTextString(m) }
func (*GetWhateversRequest) ProtoMessage()               {}
func (*GetWhateversRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *GetWhateversRequest) GetMeta() *requestmeta.RequestMeta {
	if m != nil {
		return m.Meta
	}
	return nil
}

func (m *GetWhateversRequest) GetServiceFilter() *servicefilter.ServiceFilter {
	if m != nil {
		return m.ServiceFilter
	}
	return nil
}

func (m *GetWhateversRequest) GetCache() *cache.Cache {
	if m != nil {
		return m.Cache
	}
	return nil
}

func (m *GetWhateversRequest) GetQuery() string {
	if m != nil && m.Query != nil {
		return *m.Query
	}
	return ""
}

func (m *GetWhateversRequest) GetLocationQuery() *locationquery.LocationQuery {
	if m != nil {
		return m.LocationQuery
	}
	return nil
}

func (m *GetWhateversRequest) GetSelect() *getwhateversresponseselect.GetWhateversResponseSelect {
	if m != nil {
		return m.Select
	}
	return nil
}

func (m *GetWhateversRequest) GetHash() string {
	if m != nil && m.Hash != nil {
		return *m.Hash
	}
	return ""
}

func init() {
	proto.RegisterType((*GetWhateversRequest)(nil), "getwhateversrequest.GetWhateversRequest")
}

func init() { proto.RegisterFile("getwhateversrequest/getwhateversrequest_.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 334 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x52, 0x41, 0x4b, 0x33, 0x31,
	0x14, 0xa4, 0xfd, 0xda, 0x7e, 0x18, 0xf5, 0x92, 0x7a, 0x08, 0x45, 0xa4, 0xf4, 0xd4, 0x83, 0x26,
	0xd8, 0x83, 0x27, 0x41, 0xac, 0xa0, 0x17, 0x15, 0x4c, 0x0f, 0x82, 0x17, 0x89, 0xf1, 0xb9, 0xbb,
	0xb0, 0x6d, 0xb6, 0x49, 0xb6, 0xd2, 0x9f, 0xec, 0xbf, 0x90, 0xbe, 0xcd, 0xca, 0x46, 0xd6, 0x4b,
	0x78, 0xb3, 0x33, 0xb3, 0x79, 0x33, 0x84, 0xf0, 0x04, 0xfc, 0x67, 0xaa, 0x3c, 0x6c, 0xc0, 0x3a,
	0x0b, 0xeb, 0x12, 0x9c, 0x17, 0x2d, 0xdf, 0x5e, 0x79, 0x61, 0x8d, 0x37, 0x74, 0xd8, 0xc2, 0x8d,
	0x4e, 0xc2, 0xb0, 0x04, 0xaf, 0x44, 0x63, 0x0e, 0xa6, 0xd1, 0xc4, 0x81, 0xdd, 0x64, 0x1a, 0x3e,
	0xb2, 0xdc, 0x83, 0x15, 0x11, 0xaa, 0x35, 0x54, 0x2b, 0x9d, 0x82, 0xc0, 0xf3, 0xc7, 0x97, 0x1b,
	0xad, 0x7c, 0x66, 0x56, 0xeb, 0x12, 0xec, 0x56, 0x44, 0xa8, 0xd6, 0x5c, 0xc6, 0x0b, 0xb9, 0xc2,
	0xac, 0x1c, 0x38, 0xc8, 0x41, 0xff, 0xce, 0xd1, 0xa4, 0x82, 0x7b, 0xf2, 0xd5, 0x25, 0xc3, 0x3b,
	0xf0, 0xcf, 0xb5, 0x4a, 0x56, 0xcb, 0xd3, 0x53, 0xd2, 0xdb, 0x05, 0x60, 0x9d, 0x71, 0x67, 0xba,
	0x3f, 0x63, 0xbc, 0x11, 0x8a, 0x07, 0xcd, 0x03, 0x78, 0x25, 0x51, 0x45, 0xe7, 0xe4, 0x30, 0x64,
	0xba, 0xc5, 0x4c, 0xac, 0x8b, 0xb6, 0x63, 0x1e, 0x25, 0xe5, 0x8b, 0xa6, 0x46, 0xc6, 0x16, 0x3a,
	0x21, 0x7d, 0xcc, 0xce, 0xfe, 0xa1, 0xf7, 0x80, 0x23, 0xe2, 0x37, 0xbb, 0x53, 0x56, 0x14, 0x3d,
	0x22, 0x7d, 0xcc, 0xce, 0x7a, 0xe3, 0xce, 0x74, 0x4f, 0x56, 0x60, 0x77, 0x7b, 0xdd, 0xcc, 0x13,
	0xb2, 0xfd, 0x70, 0x7b, 0xd4, 0x17, 0xbf, 0x6f, 0x6a, 0x64, 0x6c, 0xa1, 0x8f, 0x64, 0x50, 0x15,
	0xc3, 0x06, 0x68, 0xbe, 0xe0, 0x7f, 0x77, 0xc7, 0xe3, 0xc2, 0x2a, 0x6a, 0x81, 0x94, 0x0c, 0x7f,
	0xa1, 0x94, 0xf4, 0x52, 0xe5, 0x52, 0xf6, 0x1f, 0x17, 0xc5, 0x79, 0x7e, 0xfd, 0x72, 0x95, 0x64,
	0x3e, 0x2d, 0xdf, 0xb8, 0x36, 0x4b, 0x31, 0x3b, 0x77, 0x3e, 0x33, 0x22, 0x31, 0x67, 0x45, 0xae,
	0xb6, 0x89, 0x35, 0xe5, 0xea, 0x5d, 0x24, 0xb6, 0xd0, 0xc2, 0xe9, 0x14, 0x96, 0xaa, 0xed, 0x11,
	0x7e, 0x07, 0x00, 0x00, 0xff, 0xff, 0xce, 0xb2, 0x08, 0xb2, 0xae, 0x02, 0x00, 0x00,
}
