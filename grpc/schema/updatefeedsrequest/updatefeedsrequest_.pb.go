// Code generated by protoc-gen-go. DO NOT EDIT.
// source: updatefeedsrequest/updatefeedsrequest_.proto

/*
Package updatefeedsrequest is a generated protocol buffer package.

It is generated from these files:
	updatefeedsrequest/updatefeedsrequest_.proto

It has these top-level messages:
	UpdateFeedsRequest
*/
package updatefeedsrequest

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import requestmeta "github.com/21stio/go-playground/grpc/schema/requestmeta"
import servicefilter "github.com/21stio/go-playground/grpc/schema/servicefilter"
import feed "github.com/21stio/go-playground/grpc/schema/feed"
import updatefeedsresponseselect "github.com/21stio/go-playground/grpc/schema/updatefeedsresponseselect"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type UpdateFeedsRequest struct {
	Meta             *requestmeta.RequestMeta                             `protobuf:"bytes,1,opt,name=meta" json:"meta,omitempty"`
	ServiceFilter    *servicefilter.ServiceFilter                         `protobuf:"bytes,2,opt,name=serviceFilter" json:"serviceFilter,omitempty"`
	Feeds            []*feed.Feed                                         `protobuf:"bytes,3,rep,name=feeds" json:"feeds,omitempty"`
	Select           *updatefeedsresponseselect.UpdateFeedsResponseSelect `protobuf:"bytes,4,opt,name=select" json:"select,omitempty"`
	Hash             *string                                              `protobuf:"bytes,5,opt,name=hash" json:"hash,omitempty"`
	XXX_unrecognized []byte                                               `json:"-"`
}

func (m *UpdateFeedsRequest) Reset()                    { *m = UpdateFeedsRequest{} }
func (m *UpdateFeedsRequest) String() string            { return proto.CompactTextString(m) }
func (*UpdateFeedsRequest) ProtoMessage()               {}
func (*UpdateFeedsRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *UpdateFeedsRequest) GetMeta() *requestmeta.RequestMeta {
	if m != nil {
		return m.Meta
	}
	return nil
}

func (m *UpdateFeedsRequest) GetServiceFilter() *servicefilter.ServiceFilter {
	if m != nil {
		return m.ServiceFilter
	}
	return nil
}

func (m *UpdateFeedsRequest) GetFeeds() []*feed.Feed {
	if m != nil {
		return m.Feeds
	}
	return nil
}

func (m *UpdateFeedsRequest) GetSelect() *updatefeedsresponseselect.UpdateFeedsResponseSelect {
	if m != nil {
		return m.Select
	}
	return nil
}

func (m *UpdateFeedsRequest) GetHash() string {
	if m != nil && m.Hash != nil {
		return *m.Hash
	}
	return ""
}

func init() {
	proto.RegisterType((*UpdateFeedsRequest)(nil), "updatefeedsrequest.UpdateFeedsRequest")
}

func init() { proto.RegisterFile("updatefeedsrequest/updatefeedsrequest_.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 291 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x91, 0xc1, 0x4b, 0xc3, 0x30,
	0x14, 0xc6, 0xd9, 0xd6, 0x09, 0x66, 0x08, 0x92, 0x53, 0x18, 0x22, 0x65, 0xa7, 0x1e, 0xe6, 0x0b,
	0x16, 0x4f, 0x1e, 0x44, 0x76, 0xd8, 0x49, 0x2f, 0x19, 0x5e, 0xbc, 0x48, 0x6c, 0xdf, 0xda, 0x42,
	0xbb, 0xd4, 0x24, 0x15, 0xfc, 0x17, 0xfc, 0xab, 0x25, 0x69, 0x06, 0xad, 0xd3, 0x4b, 0x79, 0xef,
	0x7d, 0xdf, 0xf7, 0xfa, 0x7b, 0x84, 0xac, 0xbb, 0x36, 0x97, 0x16, 0xf7, 0x88, 0xb9, 0xd1, 0xf8,
	0xd1, 0xa1, 0xb1, 0xfc, 0x74, 0xf4, 0x06, 0xad, 0x56, 0x56, 0x51, 0x7a, 0x2a, 0x2d, 0xaf, 0x43,
	0xd1, 0xa0, 0x95, 0x7c, 0x50, 0x87, 0xcc, 0x72, 0x65, 0x50, 0x7f, 0x56, 0x19, 0xee, 0xab, 0xda,
	0xa2, 0xe6, 0xa3, 0xee, 0xe8, 0xb9, 0x74, 0x1b, 0xb9, 0xfb, 0x1c, 0x27, 0xf7, 0xa3, 0x3f, 0x99,
	0x56, 0x1d, 0x0c, 0x1a, 0xac, 0x31, 0xfb, 0x85, 0x37, 0x54, 0x42, 0x76, 0xf5, 0x3d, 0x25, 0xf4,
	0xc5, 0x9b, 0xb6, 0xce, 0x24, 0x7a, 0x26, 0xba, 0x26, 0x91, 0xe3, 0x62, 0x93, 0x78, 0x92, 0x2c,
	0x52, 0x06, 0x03, 0x56, 0x08, 0x9e, 0x67, 0xb4, 0x52, 0x78, 0x17, 0xdd, 0x90, 0x8b, 0x80, 0xba,
	0xf5, 0xa8, 0x6c, 0xea, 0x63, 0x57, 0x30, 0x3a, 0x00, 0x76, 0x43, 0x8f, 0x18, 0x47, 0x68, 0x4c,
	0xe6, 0x1e, 0x93, 0xcd, 0xe2, 0x59, 0xb2, 0x48, 0x09, 0xb8, 0x0e, 0x1c, 0x94, 0xe8, 0x05, 0xfa,
	0x44, 0xce, 0x7a, 0x76, 0x16, 0xf9, 0xf5, 0x77, 0xf0, 0xef, 0x75, 0x30, 0x3a, 0xa9, 0x57, 0x76,
	0x5e, 0x11, 0x61, 0x07, 0xa5, 0x24, 0x2a, 0xa5, 0x29, 0xd9, 0x3c, 0x9e, 0x24, 0xe7, 0xc2, 0xd7,
	0x9b, 0xc7, 0xd7, 0x87, 0xa2, 0xb2, 0x65, 0xf7, 0x0e, 0x99, 0x6a, 0x78, 0x7a, 0x6b, 0x6c, 0xa5,
	0x78, 0xa1, 0x6e, 0xda, 0x5a, 0x7e, 0x15, 0x5a, 0x75, 0x87, 0x9c, 0x17, 0xba, 0xcd, 0xb8, 0xc9,
	0x4a, 0x6c, 0xe4, 0x1f, 0x6f, 0xff, 0x13, 0x00, 0x00, 0xff, 0xff, 0xe6, 0x4e, 0x8b, 0xa5, 0x23,
	0x02, 0x00, 0x00,
}
