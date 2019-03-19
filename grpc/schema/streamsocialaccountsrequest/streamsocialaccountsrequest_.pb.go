// Code generated by protoc-gen-go. DO NOT EDIT.
// source: streamsocialaccountsrequest/streamsocialaccountsrequest_.proto

/*
Package streamsocialaccountsrequest is a generated protocol buffer package.

It is generated from these files:
	streamsocialaccountsrequest/streamsocialaccountsrequest_.proto

It has these top-level messages:
	StreamSocialAccountsRequest
*/
package streamsocialaccountsrequest

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import requestmeta "github.com/21stio/go-playground/grpc/schema/requestmeta"
import streamkind "github.com/21stio/go-playground/grpc/schema/streamkind"
import durationscalar "github.com/21stio/go-playground/grpc/schema/durationscalar"
import servicefilter "github.com/21stio/go-playground/grpc/schema/servicefilter"
import cache "github.com/21stio/go-playground/grpc/schema/cache"
import locationquery "github.com/21stio/go-playground/grpc/schema/locationquery"
import streamsocialaccountsresponseselect "github.com/21stio/go-playground/grpc/schema/streamsocialaccountsresponseselect"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type StreamSocialAccountsRequest struct {
	Meta             *requestmeta.RequestMeta                                               `protobuf:"bytes,1,opt,name=meta" json:"meta,omitempty"`
	StreamKind       *streamkind.StreamKind                                                 `protobuf:"varint,2,opt,name=streamKind,enum=streamkind.StreamKind" json:"streamKind,omitempty"`
	PollInterval     *durationscalar.DurationScalar                                         `protobuf:"bytes,3,opt,name=pollInterval" json:"pollInterval,omitempty"`
	ServiceFilter    *servicefilter.ServiceFilter                                           `protobuf:"bytes,4,opt,name=serviceFilter" json:"serviceFilter,omitempty"`
	Cache            *cache.Cache                                                           `protobuf:"bytes,5,opt,name=cache" json:"cache,omitempty"`
	Query            *string                                                                `protobuf:"bytes,6,opt,name=query" json:"query,omitempty"`
	LocationQuery    *locationquery.LocationQuery                                           `protobuf:"bytes,7,opt,name=locationQuery" json:"locationQuery,omitempty"`
	Select           *streamsocialaccountsresponseselect.StreamSocialAccountsResponseSelect `protobuf:"bytes,8,opt,name=select" json:"select,omitempty"`
	Hash             *string                                                                `protobuf:"bytes,9,opt,name=hash" json:"hash,omitempty"`
	XXX_unrecognized []byte                                                                 `json:"-"`
}

func (m *StreamSocialAccountsRequest) Reset()                    { *m = StreamSocialAccountsRequest{} }
func (m *StreamSocialAccountsRequest) String() string            { return proto.CompactTextString(m) }
func (*StreamSocialAccountsRequest) ProtoMessage()               {}
func (*StreamSocialAccountsRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *StreamSocialAccountsRequest) GetMeta() *requestmeta.RequestMeta {
	if m != nil {
		return m.Meta
	}
	return nil
}

func (m *StreamSocialAccountsRequest) GetStreamKind() streamkind.StreamKind {
	if m != nil && m.StreamKind != nil {
		return *m.StreamKind
	}
	return streamkind.StreamKind_WATCH
}

func (m *StreamSocialAccountsRequest) GetPollInterval() *durationscalar.DurationScalar {
	if m != nil {
		return m.PollInterval
	}
	return nil
}

func (m *StreamSocialAccountsRequest) GetServiceFilter() *servicefilter.ServiceFilter {
	if m != nil {
		return m.ServiceFilter
	}
	return nil
}

func (m *StreamSocialAccountsRequest) GetCache() *cache.Cache {
	if m != nil {
		return m.Cache
	}
	return nil
}

func (m *StreamSocialAccountsRequest) GetQuery() string {
	if m != nil && m.Query != nil {
		return *m.Query
	}
	return ""
}

func (m *StreamSocialAccountsRequest) GetLocationQuery() *locationquery.LocationQuery {
	if m != nil {
		return m.LocationQuery
	}
	return nil
}

func (m *StreamSocialAccountsRequest) GetSelect() *streamsocialaccountsresponseselect.StreamSocialAccountsResponseSelect {
	if m != nil {
		return m.Select
	}
	return nil
}

func (m *StreamSocialAccountsRequest) GetHash() string {
	if m != nil && m.Hash != nil {
		return *m.Hash
	}
	return ""
}

func init() {
	proto.RegisterType((*StreamSocialAccountsRequest)(nil), "streamsocialaccountsrequest.StreamSocialAccountsRequest")
}

func init() {
	proto.RegisterFile("streamsocialaccountsrequest/streamsocialaccountsrequest_.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 422 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x93, 0xc1, 0x8b, 0x13, 0x31,
	0x14, 0xc6, 0x19, 0xb7, 0x5d, 0xdd, 0xb8, 0x7a, 0x08, 0x22, 0x61, 0x77, 0x59, 0x4a, 0xf1, 0xd0,
	0x83, 0x26, 0xd8, 0x83, 0x47, 0xc1, 0x2a, 0x2b, 0x8b, 0xeb, 0xc1, 0xcc, 0xcd, 0x83, 0x12, 0x33,
	0x71, 0x26, 0x98, 0x99, 0xcc, 0x26, 0x99, 0x42, 0xff, 0x78, 0x41, 0xfa, 0x26, 0xd3, 0x4e, 0xa0,
	0x76, 0x2f, 0xe1, 0x3d, 0xde, 0xef, 0x7b, 0xaf, 0xfd, 0x3e, 0x06, 0xbd, 0xf7, 0xc1, 0x29, 0x51,
	0x7b, 0x2b, 0xb5, 0x30, 0x42, 0x4a, 0xdb, 0x35, 0xc1, 0x3b, 0x75, 0xdf, 0x29, 0x1f, 0xd8, 0x91,
	0xd9, 0x4f, 0xda, 0x3a, 0x1b, 0x2c, 0xbe, 0x3c, 0xc2, 0x5c, 0x5c, 0xc7, 0xa2, 0x56, 0x41, 0xb0,
	0x51, 0x1d, 0xc5, 0x17, 0x57, 0xbd, 0xf8, 0x8f, 0x6e, 0x0a, 0xb6, 0x2f, 0x87, 0xe9, 0xab, 0xa2,
	0x73, 0x22, 0x68, 0xdb, 0x78, 0x29, 0x8c, 0x70, 0x2c, 0x6d, 0x07, 0x6a, 0xee, 0x95, 0x5b, 0x6b,
	0xa9, 0x7e, 0x6b, 0x13, 0x94, 0x63, 0x49, 0x37, 0x30, 0x58, 0x0a, 0x59, 0x29, 0x06, 0xef, 0x4e,
	0x67, 0xac, 0x84, 0x75, 0xf7, 0x9d, 0x72, 0x1b, 0x96, 0x74, 0x03, 0x73, 0x77, 0xf8, 0xcf, 0xf9,
	0xd6, 0x36, 0x5e, 0x79, 0x65, 0x94, 0xfc, 0x9f, 0x47, 0x63, 0x24, 0x6e, 0x9b, 0xff, 0x3d, 0x41,
	0x97, 0x39, 0xd0, 0x39, 0xd0, 0x1f, 0x22, 0xcd, 0x7b, 0x63, 0xf0, 0x6b, 0x34, 0xd9, 0x9a, 0x43,
	0xb2, 0x59, 0xb6, 0x78, 0xba, 0x24, 0x74, 0x64, 0x18, 0x8d, 0xcc, 0x57, 0x15, 0x04, 0x07, 0x0a,
	0xbf, 0x43, 0xa8, 0x3f, 0xfd, 0x45, 0x37, 0x05, 0x79, 0x34, 0xcb, 0x16, 0xcf, 0x97, 0x2f, 0xe9,
	0xde, 0x45, 0x9a, 0xef, 0xa6, 0x7c, 0x44, 0xe2, 0x15, 0x3a, 0x6f, 0xad, 0x31, 0xb7, 0x4d, 0x50,
	0x6e, 0x2d, 0x0c, 0x39, 0x81, 0x6b, 0xd7, 0x34, 0x75, 0x97, 0x7e, 0x8a, 0x6d, 0x0e, 0x2d, 0x4f,
	0x34, 0x78, 0x85, 0x9e, 0x45, 0x9f, 0x6f, 0xc0, 0x67, 0x32, 0x81, 0x25, 0x57, 0x34, 0x71, 0x9f,
	0xe6, 0x63, 0x86, 0xa7, 0x12, 0x3c, 0x47, 0x53, 0xc8, 0x83, 0x4c, 0x41, 0x7b, 0x4e, 0xa1, 0xa3,
	0x1f, 0xb7, 0x2f, 0xef, 0x47, 0xf8, 0x05, 0x9a, 0x42, 0x1e, 0xe4, 0x74, 0x96, 0x2d, 0xce, 0x78,
	0xdf, 0x6c, 0xaf, 0x0f, 0x69, 0x7d, 0x83, 0xe9, 0xe3, 0x78, 0x3d, 0xc9, 0x90, 0xde, 0x8d, 0x19,
	0x9e, 0x4a, 0xf0, 0x0f, 0x74, 0xda, 0x87, 0x43, 0x9e, 0x80, 0xf8, 0x86, 0x3e, 0x9c, 0x23, 0x3d,
	0x1c, 0x5e, 0x8f, 0xe4, 0x80, 0xf0, 0xb8, 0x15, 0x63, 0x34, 0xa9, 0x84, 0xaf, 0xc8, 0x19, 0xfc,
	0x70, 0xa8, 0x57, 0xb7, 0xdf, 0x3f, 0x97, 0x3a, 0x54, 0xdd, 0x2f, 0x2a, 0x6d, 0xcd, 0x96, 0x6f,
	0x7d, 0xd0, 0x96, 0x95, 0xf6, 0x4d, 0x6b, 0xc4, 0xa6, 0x74, 0xb6, 0x6b, 0x0a, 0x56, 0xba, 0x56,
	0x32, 0x2f, 0x2b, 0x55, 0x8b, 0x63, 0x1f, 0xdf, 0xbf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xcc, 0x8c,
	0x95, 0x0c, 0xb6, 0x03, 0x00, 0x00,
}
