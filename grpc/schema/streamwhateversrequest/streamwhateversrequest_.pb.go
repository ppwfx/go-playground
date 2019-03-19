// Code generated by protoc-gen-go. DO NOT EDIT.
// source: streamwhateversrequest/streamwhateversrequest_.proto

/*
Package streamwhateversrequest is a generated protocol buffer package.

It is generated from these files:
	streamwhateversrequest/streamwhateversrequest_.proto

It has these top-level messages:
	StreamWhateversRequest
*/
package streamwhateversrequest

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import requestmeta "github.com/21stio/go-playground/grpc/schema/requestmeta"
import streamkind "github.com/21stio/go-playground/grpc/schema/streamkind"
import durationscalar "github.com/21stio/go-playground/grpc/schema/durationscalar"
import servicefilter "github.com/21stio/go-playground/grpc/schema/servicefilter"
import cache "github.com/21stio/go-playground/grpc/schema/cache"
import locationquery "github.com/21stio/go-playground/grpc/schema/locationquery"
import streamwhateversresponseselect "github.com/21stio/go-playground/grpc/schema/streamwhateversresponseselect"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type StreamWhateversRequest struct {
	Meta             *requestmeta.RequestMeta                                     `protobuf:"bytes,1,opt,name=meta" json:"meta,omitempty"`
	StreamKind       *streamkind.StreamKind                                       `protobuf:"varint,2,opt,name=streamKind,enum=streamkind.StreamKind" json:"streamKind,omitempty"`
	PollInterval     *durationscalar.DurationScalar                               `protobuf:"bytes,3,opt,name=pollInterval" json:"pollInterval,omitempty"`
	ServiceFilter    *servicefilter.ServiceFilter                                 `protobuf:"bytes,4,opt,name=serviceFilter" json:"serviceFilter,omitempty"`
	Cache            *cache.Cache                                                 `protobuf:"bytes,5,opt,name=cache" json:"cache,omitempty"`
	Query            *string                                                      `protobuf:"bytes,6,opt,name=query" json:"query,omitempty"`
	LocationQuery    *locationquery.LocationQuery                                 `protobuf:"bytes,7,opt,name=locationQuery" json:"locationQuery,omitempty"`
	Select           *streamwhateversresponseselect.StreamWhateversResponseSelect `protobuf:"bytes,8,opt,name=select" json:"select,omitempty"`
	Hash             *string                                                      `protobuf:"bytes,9,opt,name=hash" json:"hash,omitempty"`
	XXX_unrecognized []byte                                                       `json:"-"`
}

func (m *StreamWhateversRequest) Reset()                    { *m = StreamWhateversRequest{} }
func (m *StreamWhateversRequest) String() string            { return proto.CompactTextString(m) }
func (*StreamWhateversRequest) ProtoMessage()               {}
func (*StreamWhateversRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *StreamWhateversRequest) GetMeta() *requestmeta.RequestMeta {
	if m != nil {
		return m.Meta
	}
	return nil
}

func (m *StreamWhateversRequest) GetStreamKind() streamkind.StreamKind {
	if m != nil && m.StreamKind != nil {
		return *m.StreamKind
	}
	return streamkind.StreamKind_WATCH
}

func (m *StreamWhateversRequest) GetPollInterval() *durationscalar.DurationScalar {
	if m != nil {
		return m.PollInterval
	}
	return nil
}

func (m *StreamWhateversRequest) GetServiceFilter() *servicefilter.ServiceFilter {
	if m != nil {
		return m.ServiceFilter
	}
	return nil
}

func (m *StreamWhateversRequest) GetCache() *cache.Cache {
	if m != nil {
		return m.Cache
	}
	return nil
}

func (m *StreamWhateversRequest) GetQuery() string {
	if m != nil && m.Query != nil {
		return *m.Query
	}
	return ""
}

func (m *StreamWhateversRequest) GetLocationQuery() *locationquery.LocationQuery {
	if m != nil {
		return m.LocationQuery
	}
	return nil
}

func (m *StreamWhateversRequest) GetSelect() *streamwhateversresponseselect.StreamWhateversResponseSelect {
	if m != nil {
		return m.Select
	}
	return nil
}

func (m *StreamWhateversRequest) GetHash() string {
	if m != nil && m.Hash != nil {
		return *m.Hash
	}
	return ""
}

func init() {
	proto.RegisterType((*StreamWhateversRequest)(nil), "streamwhateversrequest.StreamWhateversRequest")
}

func init() {
	proto.RegisterFile("streamwhateversrequest/streamwhateversrequest_.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 417 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x53, 0x41, 0x8b, 0x13, 0x31,
	0x18, 0xa5, 0x6e, 0xbb, 0xba, 0x71, 0xf5, 0x10, 0xa4, 0x84, 0x65, 0x59, 0x4a, 0xf1, 0xd0, 0x83,
	0x26, 0x58, 0xc4, 0x93, 0xa7, 0x71, 0x11, 0x44, 0x3d, 0x98, 0x11, 0x04, 0x2f, 0x12, 0x33, 0x9f,
	0x33, 0x83, 0x99, 0xc9, 0x6c, 0x92, 0xa9, 0xec, 0x9f, 0xf6, 0x37, 0x48, 0xbf, 0xc9, 0xec, 0x4e,
	0xd8, 0xba, 0x97, 0xf0, 0xbd, 0xbe, 0xf7, 0xbe, 0xaf, 0x7d, 0x8f, 0x92, 0xd7, 0x3e, 0x38, 0x50,
	0xcd, 0x9f, 0x4a, 0x05, 0xd8, 0x81, 0xf3, 0x0e, 0xae, 0x7a, 0xf0, 0x41, 0x1c, 0xfe, 0xf8, 0x07,
	0xef, 0x9c, 0x0d, 0x96, 0x2e, 0x0f, 0xd3, 0x67, 0x17, 0x71, 0x68, 0x20, 0x28, 0x31, 0x99, 0xa3,
	0xef, 0xec, 0x7c, 0xf0, 0xfd, 0xae, 0xdb, 0x42, 0xdc, 0x8e, 0x23, 0xfb, 0xbc, 0xe8, 0x9d, 0x0a,
	0xb5, 0x6d, 0xbd, 0x56, 0x46, 0x39, 0x91, 0xc2, 0x51, 0xb5, 0xf6, 0xe0, 0x76, 0xb5, 0x86, 0x5f,
	0xb5, 0x09, 0xe0, 0x44, 0x82, 0x46, 0x0d, 0xd5, 0x4a, 0x57, 0x20, 0xf0, 0xbd, 0xf1, 0x19, 0xab,
	0x71, 0xdd, 0x55, 0x0f, 0xee, 0x5a, 0x24, 0x68, 0xd4, 0x64, 0x77, 0x7e, 0x97, 0xef, 0x6c, 0xeb,
	0xc1, 0x83, 0x01, 0x7d, 0x20, 0x94, 0x29, 0x1b, 0x77, 0xac, 0xff, 0x1e, 0x91, 0x65, 0x8e, 0xc2,
	0x6f, 0xa3, 0x50, 0x0e, 0x49, 0xd0, 0x17, 0x64, 0xbe, 0x4f, 0x83, 0xcd, 0x56, 0xb3, 0xcd, 0xe3,
	0x2d, 0xe3, 0x93, 0x84, 0x78, 0xd4, 0x7c, 0x86, 0xa0, 0x24, 0xaa, 0xe8, 0x1b, 0x42, 0x86, 0x83,
	0x1f, 0xeb, 0xb6, 0x60, 0x0f, 0x56, 0xb3, 0xcd, 0xd3, 0xed, 0x92, 0xdf, 0xc6, 0xc6, 0xf3, 0x1b,
	0x56, 0x4e, 0x94, 0x34, 0x23, 0xa7, 0x9d, 0x35, 0xe6, 0x43, 0x1b, 0xc0, 0xed, 0x94, 0x61, 0x47,
	0x78, 0xed, 0x82, 0xa7, 0x71, 0xf2, 0xcb, 0x08, 0x73, 0x84, 0x32, 0xf1, 0xd0, 0x8c, 0x3c, 0x89,
	0xc1, 0xbe, 0xc7, 0x60, 0xd9, 0x1c, 0x97, 0x9c, 0xf3, 0x24, 0x6e, 0x9e, 0x4f, 0x35, 0x32, 0xb5,
	0xd0, 0x35, 0x59, 0x60, 0x01, 0x6c, 0x81, 0xde, 0x53, 0x8e, 0x88, 0xbf, 0xdb, 0xbf, 0x72, 0xa0,
	0xe8, 0x33, 0xb2, 0xc0, 0x02, 0xd8, 0xf1, 0x6a, 0xb6, 0x39, 0x91, 0x03, 0xd8, 0x5f, 0x1f, 0xeb,
	0xf9, 0x82, 0xec, 0xc3, 0x78, 0x3d, 0x29, 0x8d, 0x7f, 0x9a, 0x6a, 0x64, 0x6a, 0xa1, 0x5f, 0xc9,
	0xf1, 0xd0, 0x0b, 0x7b, 0x84, 0xe6, 0xb7, 0xfc, 0xde, 0xf6, 0xf8, 0x9d, 0xca, 0x06, 0x36, 0x47,
	0x56, 0xc6, 0x5d, 0x94, 0x92, 0x79, 0xa5, 0x7c, 0xc5, 0x4e, 0xf0, 0xeb, 0xe2, 0x9c, 0x5d, 0x7e,
	0xcf, 0xca, 0x3a, 0x54, 0xfd, 0x4f, 0xae, 0x6d, 0x23, 0xb6, 0xaf, 0x7c, 0xa8, 0xad, 0x28, 0xed,
	0xcb, 0xce, 0xa8, 0xeb, 0xd2, 0xd9, 0xbe, 0x2d, 0x44, 0xe9, 0x3a, 0x2d, 0xbc, 0xae, 0xa0, 0x51,
	0xff, 0xf9, 0x67, 0xfd, 0x0b, 0x00, 0x00, 0xff, 0xff, 0xa8, 0x1e, 0x7d, 0xf4, 0x89, 0x03, 0x00,
	0x00,
}
