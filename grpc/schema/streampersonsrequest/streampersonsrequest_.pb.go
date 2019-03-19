// Code generated by protoc-gen-go. DO NOT EDIT.
// source: streampersonsrequest/streampersonsrequest_.proto

/*
Package streampersonsrequest is a generated protocol buffer package.

It is generated from these files:
	streampersonsrequest/streampersonsrequest_.proto

It has these top-level messages:
	StreamPersonsRequest
*/
package streampersonsrequest

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import requestmeta "github.com/21stio/go-playground/grpc/schema/requestmeta"
import streamkind "github.com/21stio/go-playground/grpc/schema/streamkind"
import durationscalar "github.com/21stio/go-playground/grpc/schema/durationscalar"
import servicefilter "github.com/21stio/go-playground/grpc/schema/servicefilter"
import cache "github.com/21stio/go-playground/grpc/schema/cache"
import locationquery "github.com/21stio/go-playground/grpc/schema/locationquery"
import streampersonsresponseselect "github.com/21stio/go-playground/grpc/schema/streampersonsresponseselect"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type StreamPersonsRequest struct {
	Meta             *requestmeta.RequestMeta                                 `protobuf:"bytes,1,opt,name=meta" json:"meta,omitempty"`
	StreamKind       *streamkind.StreamKind                                   `protobuf:"varint,2,opt,name=streamKind,enum=streamkind.StreamKind" json:"streamKind,omitempty"`
	PollInterval     *durationscalar.DurationScalar                           `protobuf:"bytes,3,opt,name=pollInterval" json:"pollInterval,omitempty"`
	ServiceFilter    *servicefilter.ServiceFilter                             `protobuf:"bytes,4,opt,name=serviceFilter" json:"serviceFilter,omitempty"`
	Cache            *cache.Cache                                             `protobuf:"bytes,5,opt,name=cache" json:"cache,omitempty"`
	Query            *string                                                  `protobuf:"bytes,6,opt,name=query" json:"query,omitempty"`
	LocationQuery    *locationquery.LocationQuery                             `protobuf:"bytes,7,opt,name=locationQuery" json:"locationQuery,omitempty"`
	Select           *streampersonsresponseselect.StreamPersonsResponseSelect `protobuf:"bytes,8,opt,name=select" json:"select,omitempty"`
	Hash             *string                                                  `protobuf:"bytes,9,opt,name=hash" json:"hash,omitempty"`
	XXX_unrecognized []byte                                                   `json:"-"`
}

func (m *StreamPersonsRequest) Reset()                    { *m = StreamPersonsRequest{} }
func (m *StreamPersonsRequest) String() string            { return proto.CompactTextString(m) }
func (*StreamPersonsRequest) ProtoMessage()               {}
func (*StreamPersonsRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *StreamPersonsRequest) GetMeta() *requestmeta.RequestMeta {
	if m != nil {
		return m.Meta
	}
	return nil
}

func (m *StreamPersonsRequest) GetStreamKind() streamkind.StreamKind {
	if m != nil && m.StreamKind != nil {
		return *m.StreamKind
	}
	return streamkind.StreamKind_WATCH
}

func (m *StreamPersonsRequest) GetPollInterval() *durationscalar.DurationScalar {
	if m != nil {
		return m.PollInterval
	}
	return nil
}

func (m *StreamPersonsRequest) GetServiceFilter() *servicefilter.ServiceFilter {
	if m != nil {
		return m.ServiceFilter
	}
	return nil
}

func (m *StreamPersonsRequest) GetCache() *cache.Cache {
	if m != nil {
		return m.Cache
	}
	return nil
}

func (m *StreamPersonsRequest) GetQuery() string {
	if m != nil && m.Query != nil {
		return *m.Query
	}
	return ""
}

func (m *StreamPersonsRequest) GetLocationQuery() *locationquery.LocationQuery {
	if m != nil {
		return m.LocationQuery
	}
	return nil
}

func (m *StreamPersonsRequest) GetSelect() *streampersonsresponseselect.StreamPersonsResponseSelect {
	if m != nil {
		return m.Select
	}
	return nil
}

func (m *StreamPersonsRequest) GetHash() string {
	if m != nil && m.Hash != nil {
		return *m.Hash
	}
	return ""
}

func init() {
	proto.RegisterType((*StreamPersonsRequest)(nil), "streampersonsrequest.StreamPersonsRequest")
}

func init() { proto.RegisterFile("streampersonsrequest/streampersonsrequest_.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 415 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x53, 0x41, 0x6b, 0x14, 0x31,
	0x18, 0x65, 0xed, 0x6e, 0xb5, 0xb1, 0x7a, 0x08, 0x8b, 0x84, 0x52, 0xca, 0xb2, 0x78, 0xd8, 0x83,
	0x26, 0xba, 0x07, 0xf1, 0x24, 0xb2, 0x8a, 0x20, 0x2a, 0xd4, 0xcc, 0xcd, 0x8b, 0xc4, 0xec, 0xe7,
	0xcc, 0x60, 0x66, 0x32, 0x4d, 0x32, 0x85, 0xfe, 0x63, 0x7f, 0x86, 0xec, 0x97, 0x4c, 0x3b, 0x59,
	0x97, 0x5e, 0xc2, 0xf7, 0xe6, 0xbd, 0x97, 0x6f, 0xe6, 0x3d, 0x86, 0xbc, 0xf2, 0xc1, 0x81, 0x6a,
	0x3a, 0x70, 0xde, 0xb6, 0xde, 0xc1, 0x55, 0x0f, 0x3e, 0x88, 0x43, 0x0f, 0x7f, 0xf2, 0xce, 0xd9,
	0x60, 0xe9, 0xfc, 0x10, 0x79, 0x76, 0x91, 0x86, 0x06, 0x82, 0x12, 0xa3, 0x39, 0xb9, 0xce, 0xce,
	0xa3, 0xeb, 0x4f, 0xdd, 0x6e, 0xc5, 0xdd, 0x38, 0xb0, 0xcf, 0xb7, 0xbd, 0x53, 0xa1, 0xb6, 0xad,
	0xd7, 0xca, 0x28, 0x27, 0x72, 0x38, 0xa8, 0x96, 0x1e, 0xdc, 0x75, 0xad, 0xe1, 0x77, 0x6d, 0x02,
	0x38, 0x91, 0xa1, 0x41, 0x43, 0xb5, 0xd2, 0x15, 0x08, 0x3c, 0x6f, 0x7d, 0xc6, 0x6a, 0xbc, 0xee,
	0xaa, 0x07, 0x77, 0x23, 0x32, 0x34, 0x68, 0xde, 0xed, 0x7d, 0x95, 0xef, 0x6c, 0xeb, 0xc1, 0x83,
	0x01, 0xfd, 0x5f, 0x1c, 0x63, 0x2e, 0xf9, 0x97, 0x7f, 0x8f, 0xc8, 0xbc, 0x40, 0xd9, 0x65, 0x94,
	0xc9, 0x98, 0x01, 0x7d, 0x41, 0xa6, 0xbb, 0x1c, 0xd8, 0x64, 0x31, 0x59, 0x3d, 0x5e, 0x33, 0x3e,
	0xca, 0x86, 0x27, 0xcd, 0x37, 0x08, 0x4a, 0xa2, 0x8a, 0xbe, 0x21, 0x24, 0x2e, 0xfb, 0x52, 0xb7,
	0x5b, 0xf6, 0x60, 0x31, 0x59, 0x3d, 0x5d, 0x3f, 0xe3, 0x77, 0x81, 0xf1, 0xe2, 0x96, 0x95, 0x23,
	0x25, 0xdd, 0x90, 0xd3, 0xce, 0x1a, 0xf3, 0xb9, 0x0d, 0xe0, 0xae, 0x95, 0x61, 0x47, 0xb8, 0xed,
	0x82, 0xe7, 0x41, 0xf2, 0x8f, 0x09, 0x16, 0x08, 0x65, 0xe6, 0xa1, 0x1b, 0xf2, 0x24, 0x45, 0xfa,
	0x09, 0x23, 0x65, 0x53, 0xbc, 0xe4, 0x9c, 0x67, 0x41, 0xf3, 0x62, 0xac, 0x91, 0xb9, 0x85, 0x2e,
	0xc9, 0x0c, 0xa3, 0x67, 0x33, 0xf4, 0x9e, 0x72, 0x44, 0xfc, 0xc3, 0xee, 0x94, 0x91, 0xa2, 0x73,
	0x32, 0xc3, 0xe8, 0xd9, 0xf1, 0x62, 0xb2, 0x3a, 0x91, 0x11, 0xec, 0xb6, 0x0f, 0xc5, 0x7c, 0x47,
	0xf6, 0x61, 0xda, 0x9e, 0xd5, 0xc5, 0xbf, 0x8e, 0x35, 0x32, 0xb7, 0xd0, 0x4b, 0x72, 0x1c, 0x5b,
	0x61, 0x8f, 0xd0, 0xfc, 0x96, 0xdf, 0xd3, 0x1c, 0xdf, 0xab, 0x2b, 0x72, 0x05, 0x72, 0x32, 0xdd,
	0x43, 0x29, 0x99, 0x56, 0xca, 0x57, 0xec, 0x04, 0x5f, 0x15, 0xe7, 0xcd, 0xe6, 0xc7, 0xfb, 0xb2,
	0x0e, 0x55, 0xff, 0x8b, 0x6b, 0xdb, 0x88, 0xf5, 0x6b, 0x1f, 0x6a, 0x2b, 0x4a, 0xfb, 0xb2, 0x33,
	0xea, 0xa6, 0x74, 0xb6, 0x6f, 0xb7, 0xa2, 0x74, 0x9d, 0x16, 0x5e, 0x57, 0xd0, 0xa8, 0x83, 0xff,
	0xd2, 0xbf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x42, 0xae, 0x9d, 0x75, 0x77, 0x03, 0x00, 0x00,
}
