// Code generated by protoc-gen-go. DO NOT EDIT.
// source: streamsystemaccountsrequest/streamsystemaccountsrequest_.proto

/*
Package streamsystemaccountsrequest is a generated protocol buffer package.

It is generated from these files:
	streamsystemaccountsrequest/streamsystemaccountsrequest_.proto

It has these top-level messages:
	StreamSystemAccountsRequest
*/
package streamsystemaccountsrequest

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import requestmeta "github.com/21stio/go-playground/grpc/schema/requestmeta"
import streamkind "github.com/21stio/go-playground/grpc/schema/streamkind"
import durationscalar "github.com/21stio/go-playground/grpc/schema/durationscalar"
import servicefilter "github.com/21stio/go-playground/grpc/schema/servicefilter"
import cache "github.com/21stio/go-playground/grpc/schema/cache"
import locationquery "github.com/21stio/go-playground/grpc/schema/locationquery"
import streamsystemaccountsresponseselect "github.com/21stio/go-playground/grpc/schema/streamsystemaccountsresponseselect"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type StreamSystemAccountsRequest struct {
	Meta             *requestmeta.RequestMeta                                               `protobuf:"bytes,1,opt,name=meta" json:"meta,omitempty"`
	StreamKind       *streamkind.StreamKind                                                 `protobuf:"varint,2,opt,name=streamKind,enum=streamkind.StreamKind" json:"streamKind,omitempty"`
	PollInterval     *durationscalar.DurationScalar                                         `protobuf:"bytes,3,opt,name=pollInterval" json:"pollInterval,omitempty"`
	ServiceFilter    *servicefilter.ServiceFilter                                           `protobuf:"bytes,4,opt,name=serviceFilter" json:"serviceFilter,omitempty"`
	Cache            *cache.Cache                                                           `protobuf:"bytes,5,opt,name=cache" json:"cache,omitempty"`
	Query            *string                                                                `protobuf:"bytes,6,opt,name=query" json:"query,omitempty"`
	LocationQuery    *locationquery.LocationQuery                                           `protobuf:"bytes,7,opt,name=locationQuery" json:"locationQuery,omitempty"`
	Select           *streamsystemaccountsresponseselect.StreamSystemAccountsResponseSelect `protobuf:"bytes,8,opt,name=select" json:"select,omitempty"`
	Hash             *string                                                                `protobuf:"bytes,9,opt,name=hash" json:"hash,omitempty"`
	XXX_unrecognized []byte                                                                 `json:"-"`
}

func (m *StreamSystemAccountsRequest) Reset()                    { *m = StreamSystemAccountsRequest{} }
func (m *StreamSystemAccountsRequest) String() string            { return proto.CompactTextString(m) }
func (*StreamSystemAccountsRequest) ProtoMessage()               {}
func (*StreamSystemAccountsRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *StreamSystemAccountsRequest) GetMeta() *requestmeta.RequestMeta {
	if m != nil {
		return m.Meta
	}
	return nil
}

func (m *StreamSystemAccountsRequest) GetStreamKind() streamkind.StreamKind {
	if m != nil && m.StreamKind != nil {
		return *m.StreamKind
	}
	return streamkind.StreamKind_WATCH
}

func (m *StreamSystemAccountsRequest) GetPollInterval() *durationscalar.DurationScalar {
	if m != nil {
		return m.PollInterval
	}
	return nil
}

func (m *StreamSystemAccountsRequest) GetServiceFilter() *servicefilter.ServiceFilter {
	if m != nil {
		return m.ServiceFilter
	}
	return nil
}

func (m *StreamSystemAccountsRequest) GetCache() *cache.Cache {
	if m != nil {
		return m.Cache
	}
	return nil
}

func (m *StreamSystemAccountsRequest) GetQuery() string {
	if m != nil && m.Query != nil {
		return *m.Query
	}
	return ""
}

func (m *StreamSystemAccountsRequest) GetLocationQuery() *locationquery.LocationQuery {
	if m != nil {
		return m.LocationQuery
	}
	return nil
}

func (m *StreamSystemAccountsRequest) GetSelect() *streamsystemaccountsresponseselect.StreamSystemAccountsResponseSelect {
	if m != nil {
		return m.Select
	}
	return nil
}

func (m *StreamSystemAccountsRequest) GetHash() string {
	if m != nil && m.Hash != nil {
		return *m.Hash
	}
	return ""
}

func init() {
	proto.RegisterType((*StreamSystemAccountsRequest)(nil), "streamsystemaccountsrequest.StreamSystemAccountsRequest")
}

func init() {
	proto.RegisterFile("streamsystemaccountsrequest/streamsystemaccountsrequest_.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 422 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x93, 0xc1, 0x6b, 0x14, 0x31,
	0x14, 0xc6, 0x19, 0xbb, 0x5b, 0x6d, 0xac, 0x1e, 0x82, 0x48, 0x68, 0x4b, 0x59, 0x16, 0x0f, 0x7b,
	0xd0, 0x04, 0xf7, 0xe0, 0x51, 0x70, 0x95, 0x4a, 0xb1, 0x1e, 0xcc, 0xdc, 0x3c, 0x28, 0x31, 0xf3,
	0x9c, 0x19, 0xcc, 0x4c, 0xa6, 0x49, 0xa6, 0xb0, 0x7f, 0xbc, 0x20, 0xfb, 0x26, 0xd3, 0x4e, 0xa0,
	0x6e, 0x2f, 0xe1, 0x3d, 0xde, 0xef, 0x7b, 0x6f, 0xf7, 0xfb, 0x18, 0xf2, 0xde, 0x07, 0x07, 0xaa,
	0xf1, 0x5b, 0x1f, 0xa0, 0x51, 0x5a, 0xdb, 0xbe, 0x0d, 0xde, 0xc1, 0x75, 0x0f, 0x3e, 0x88, 0x3d,
	0xb3, 0x9f, 0xbc, 0x73, 0x36, 0x58, 0x7a, 0xba, 0x87, 0x39, 0x39, 0x8f, 0x45, 0x03, 0x41, 0x89,
	0x49, 0x1d, 0xc5, 0x27, 0x67, 0x83, 0xf8, 0x4f, 0xdd, 0x16, 0xe2, 0xae, 0x1c, 0xa7, 0xaf, 0x8a,
	0xde, 0xa9, 0x50, 0xdb, 0xd6, 0x6b, 0x65, 0x94, 0x13, 0x69, 0x3b, 0x52, 0x4b, 0x0f, 0xee, 0xa6,
	0xd6, 0xf0, 0xbb, 0x36, 0x01, 0x9c, 0x48, 0xba, 0x91, 0xa1, 0x5a, 0xe9, 0x0a, 0x04, 0xbe, 0xb7,
	0x3a, 0x63, 0x35, 0xae, 0xbb, 0xee, 0xc1, 0x6d, 0x45, 0xd2, 0x8d, 0xcc, 0xd5, 0xfd, 0x7f, 0xce,
	0x77, 0xb6, 0xf5, 0xe0, 0xc1, 0x80, 0xfe, 0x9f, 0x47, 0x53, 0x24, 0x6e, 0x5b, 0xfe, 0x3d, 0x20,
	0xa7, 0x39, 0xd2, 0x39, 0xd2, 0x1f, 0x22, 0x2d, 0x07, 0x63, 0xe8, 0x6b, 0x32, 0xdb, 0x99, 0xc3,
	0xb2, 0x45, 0xb6, 0x7a, 0xba, 0x66, 0x7c, 0x62, 0x18, 0x8f, 0xcc, 0x57, 0x08, 0x4a, 0x22, 0x45,
	0xdf, 0x11, 0x32, 0x9c, 0xfe, 0x52, 0xb7, 0x05, 0x7b, 0xb4, 0xc8, 0x56, 0xcf, 0xd7, 0x2f, 0xf9,
	0x9d, 0x8b, 0x3c, 0xbf, 0x9d, 0xca, 0x09, 0x49, 0x37, 0xe4, 0xb8, 0xb3, 0xc6, 0x5c, 0xb6, 0x01,
	0xdc, 0x8d, 0x32, 0xec, 0x00, 0xaf, 0x9d, 0xf3, 0xd4, 0x5d, 0xfe, 0x29, 0xb6, 0x39, 0xb6, 0x32,
	0xd1, 0xd0, 0x0d, 0x79, 0x16, 0x7d, 0xbe, 0x40, 0x9f, 0xd9, 0x0c, 0x97, 0x9c, 0xf1, 0xc4, 0x7d,
	0x9e, 0x4f, 0x19, 0x99, 0x4a, 0xe8, 0x92, 0xcc, 0x31, 0x0f, 0x36, 0x47, 0xed, 0x31, 0xc7, 0x8e,
	0x7f, 0xdc, 0xbd, 0x72, 0x18, 0xd1, 0x17, 0x64, 0x8e, 0x79, 0xb0, 0xc3, 0x45, 0xb6, 0x3a, 0x92,
	0x43, 0xb3, 0xbb, 0x3e, 0xa6, 0xf5, 0x0d, 0xa7, 0x8f, 0xe3, 0xf5, 0x24, 0x43, 0x7e, 0x35, 0x65,
	0x64, 0x2a, 0xa1, 0x3f, 0xc8, 0xe1, 0x10, 0x0e, 0x7b, 0x82, 0xe2, 0x0b, 0xfe, 0x70, 0x8e, 0xfc,
	0xfe, 0xf0, 0x06, 0x24, 0x47, 0x44, 0xc6, 0xad, 0x94, 0x92, 0x59, 0xa5, 0x7c, 0xc5, 0x8e, 0xf0,
	0x87, 0x63, 0xbd, 0xb9, 0xfc, 0xfe, 0xb9, 0xac, 0x43, 0xd5, 0xff, 0xe2, 0xda, 0x36, 0x62, 0xfd,
	0xd6, 0x87, 0xda, 0x8a, 0xd2, 0xbe, 0xe9, 0x8c, 0xda, 0x96, 0xce, 0xf6, 0x6d, 0x21, 0x4a, 0xd7,
	0x69, 0xe1, 0x75, 0x05, 0x8d, 0xda, 0xf7, 0xf1, 0xfd, 0x0b, 0x00, 0x00, 0xff, 0xff, 0x5a, 0x56,
	0xa2, 0x59, 0xb6, 0x03, 0x00, 0x00,
}
