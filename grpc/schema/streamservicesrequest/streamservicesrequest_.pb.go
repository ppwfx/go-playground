// Code generated by protoc-gen-go. DO NOT EDIT.
// source: streamservicesrequest/streamservicesrequest_.proto

/*
Package streamservicesrequest is a generated protocol buffer package.

It is generated from these files:
	streamservicesrequest/streamservicesrequest_.proto

It has these top-level messages:
	StreamServicesRequest
*/
package streamservicesrequest

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import requestmeta "github.com/21stio/go-playground/grpc/schema/requestmeta"
import streamkind "github.com/21stio/go-playground/grpc/schema/streamkind"
import durationscalar "github.com/21stio/go-playground/grpc/schema/durationscalar"
import servicefilter "github.com/21stio/go-playground/grpc/schema/servicefilter"
import cache "github.com/21stio/go-playground/grpc/schema/cache"
import locationquery "github.com/21stio/go-playground/grpc/schema/locationquery"
import streamservicesresponseselect "github.com/21stio/go-playground/grpc/schema/streamservicesresponseselect"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type StreamServicesRequest struct {
	Meta             *requestmeta.RequestMeta                                   `protobuf:"bytes,1,opt,name=meta" json:"meta,omitempty"`
	StreamKind       *streamkind.StreamKind                                     `protobuf:"varint,2,opt,name=streamKind,enum=streamkind.StreamKind" json:"streamKind,omitempty"`
	PollInterval     *durationscalar.DurationScalar                             `protobuf:"bytes,3,opt,name=pollInterval" json:"pollInterval,omitempty"`
	ServiceFilter    *servicefilter.ServiceFilter                               `protobuf:"bytes,4,opt,name=serviceFilter" json:"serviceFilter,omitempty"`
	Cache            *cache.Cache                                               `protobuf:"bytes,5,opt,name=cache" json:"cache,omitempty"`
	Query            *string                                                    `protobuf:"bytes,6,opt,name=query" json:"query,omitempty"`
	LocationQuery    *locationquery.LocationQuery                               `protobuf:"bytes,7,opt,name=locationQuery" json:"locationQuery,omitempty"`
	Select           *streamservicesresponseselect.StreamServicesResponseSelect `protobuf:"bytes,8,opt,name=select" json:"select,omitempty"`
	Hash             *string                                                    `protobuf:"bytes,9,opt,name=hash" json:"hash,omitempty"`
	XXX_unrecognized []byte                                                     `json:"-"`
}

func (m *StreamServicesRequest) Reset()                    { *m = StreamServicesRequest{} }
func (m *StreamServicesRequest) String() string            { return proto.CompactTextString(m) }
func (*StreamServicesRequest) ProtoMessage()               {}
func (*StreamServicesRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *StreamServicesRequest) GetMeta() *requestmeta.RequestMeta {
	if m != nil {
		return m.Meta
	}
	return nil
}

func (m *StreamServicesRequest) GetStreamKind() streamkind.StreamKind {
	if m != nil && m.StreamKind != nil {
		return *m.StreamKind
	}
	return streamkind.StreamKind_WATCH
}

func (m *StreamServicesRequest) GetPollInterval() *durationscalar.DurationScalar {
	if m != nil {
		return m.PollInterval
	}
	return nil
}

func (m *StreamServicesRequest) GetServiceFilter() *servicefilter.ServiceFilter {
	if m != nil {
		return m.ServiceFilter
	}
	return nil
}

func (m *StreamServicesRequest) GetCache() *cache.Cache {
	if m != nil {
		return m.Cache
	}
	return nil
}

func (m *StreamServicesRequest) GetQuery() string {
	if m != nil && m.Query != nil {
		return *m.Query
	}
	return ""
}

func (m *StreamServicesRequest) GetLocationQuery() *locationquery.LocationQuery {
	if m != nil {
		return m.LocationQuery
	}
	return nil
}

func (m *StreamServicesRequest) GetSelect() *streamservicesresponseselect.StreamServicesResponseSelect {
	if m != nil {
		return m.Select
	}
	return nil
}

func (m *StreamServicesRequest) GetHash() string {
	if m != nil && m.Hash != nil {
		return *m.Hash
	}
	return ""
}

func init() {
	proto.RegisterType((*StreamServicesRequest)(nil), "streamservicesrequest.StreamServicesRequest")
}

func init() { proto.RegisterFile("streamservicesrequest/streamservicesrequest_.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 411 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x53, 0xc1, 0x8e, 0xd3, 0x30,
	0x10, 0x55, 0xd8, 0x76, 0x61, 0xcd, 0xc2, 0xc1, 0x62, 0x91, 0xb5, 0xaa, 0xaa, 0xaa, 0xe2, 0xd0,
	0x03, 0xd8, 0x22, 0x07, 0x0e, 0x9c, 0xa0, 0x45, 0x48, 0x08, 0x38, 0xe0, 0xdc, 0xb8, 0x20, 0xe3,
	0x98, 0x24, 0xc2, 0x89, 0x53, 0xdb, 0xa9, 0xd4, 0x6f, 0xe6, 0x27, 0x50, 0x27, 0x4e, 0x1b, 0x97,
	0x8a, 0x8b, 0x35, 0x2f, 0xef, 0xbd, 0x99, 0xf8, 0x8d, 0x8c, 0x52, 0xe7, 0xad, 0x12, 0xb5, 0x53,
	0x76, 0x57, 0x49, 0xe5, 0xac, 0xda, 0x76, 0xca, 0x79, 0x76, 0xf1, 0xeb, 0x0f, 0xda, 0x5a, 0xe3,
	0x0d, 0xbe, 0xbb, 0xc8, 0xde, 0xcf, 0x43, 0x51, 0x2b, 0x2f, 0xd8, 0xa8, 0x0e, 0xb6, 0xfb, 0x59,
	0x6f, 0xfb, 0x5d, 0x35, 0x39, 0x3b, 0x95, 0x03, 0xfb, 0x22, 0xef, 0xac, 0xf0, 0x95, 0x69, 0x9c,
	0x14, 0x5a, 0x58, 0x16, 0xc3, 0x41, 0xb5, 0x0c, 0x43, 0x7f, 0x55, 0xda, 0x2b, 0xcb, 0x22, 0x34,
	0x68, 0xb0, 0x14, 0xb2, 0x54, 0x0c, 0xce, 0xa3, 0x4f, 0x1b, 0x09, 0xed, 0xb6, 0x9d, 0xb2, 0x7b,
	0x16, 0xa1, 0x41, 0xf3, 0xee, 0xfc, 0x5a, 0xae, 0x35, 0x8d, 0x53, 0x4e, 0x69, 0x25, 0xff, 0x4d,
	0x64, 0x4c, 0x86, 0x0e, 0xcb, 0x3f, 0x57, 0xe8, 0x2e, 0x03, 0x5d, 0x16, 0x74, 0xbc, 0x8f, 0x01,
	0xbf, 0x44, 0x93, 0x43, 0x14, 0x24, 0x59, 0x24, 0xab, 0xc7, 0x29, 0xa1, 0xa3, 0x78, 0x68, 0xd0,
	0x7c, 0x55, 0x5e, 0x70, 0x50, 0xe1, 0x37, 0x08, 0xf5, 0xe3, 0x3e, 0x57, 0x4d, 0x4e, 0x1e, 0x2c,
	0x92, 0xd5, 0xd3, 0xf4, 0x39, 0x3d, 0x65, 0x46, 0xb3, 0x23, 0xcb, 0x47, 0x4a, 0xbc, 0x46, 0xb7,
	0xad, 0xd1, 0xfa, 0x53, 0xe3, 0x95, 0xdd, 0x09, 0x4d, 0xae, 0x60, 0xda, 0x9c, 0xc6, 0x59, 0xd2,
	0x0f, 0x01, 0x66, 0x00, 0x79, 0xe4, 0xc1, 0x6b, 0xf4, 0x24, 0x5c, 0xf2, 0x23, 0xa4, 0x4a, 0x26,
	0xd0, 0x64, 0x46, 0xa3, 0xac, 0x69, 0x36, 0xd6, 0xf0, 0xd8, 0x82, 0x97, 0x68, 0x0a, 0xe9, 0x93,
	0x29, 0x78, 0x6f, 0x29, 0x20, 0xba, 0x39, 0x9c, 0xbc, 0xa7, 0xf0, 0x33, 0x34, 0x85, 0xf4, 0xc9,
	0xf5, 0x22, 0x59, 0xdd, 0xf0, 0x1e, 0x1c, 0xa6, 0x0f, 0xbb, 0xf9, 0x06, 0xec, 0xc3, 0x30, 0x3d,
	0xda, 0x18, 0xfd, 0x32, 0xd6, 0xf0, 0xd8, 0x82, 0x39, 0xba, 0xee, 0xd7, 0x42, 0x1e, 0x81, 0xf9,
	0x2d, 0xfd, 0xdf, 0xee, 0xe8, 0xf9, 0xc2, 0x7a, 0x32, 0x03, 0x92, 0x87, 0x4e, 0x18, 0xa3, 0x49,
	0x29, 0x5c, 0x49, 0x6e, 0xe0, 0x67, 0xa1, 0x5e, 0x6f, 0xbe, 0xbf, 0x2f, 0x2a, 0x5f, 0x76, 0x3f,
	0xa9, 0x34, 0x35, 0x4b, 0x5f, 0x3b, 0x5f, 0x19, 0x56, 0x98, 0x57, 0xad, 0x16, 0xfb, 0xc2, 0x9a,
	0xae, 0xc9, 0x59, 0x61, 0x5b, 0xc9, 0x9c, 0x2c, 0x55, 0x2d, 0x2e, 0x3f, 0xa9, 0xbf, 0x01, 0x00,
	0x00, 0xff, 0xff, 0x10, 0x4e, 0x52, 0x8f, 0x80, 0x03, 0x00, 0x00,
}
