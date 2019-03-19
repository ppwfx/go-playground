// Code generated by protoc-gen-go. DO NOT EDIT.
// source: streamfeedsresponseselect/streamfeedsresponseselect_.proto

/*
Package streamfeedsresponseselect is a generated protocol buffer package.

It is generated from these files:
	streamfeedsresponseselect/streamfeedsresponseselect_.proto

It has these top-level messages:
	StreamFeedsResponseSelect
*/
package streamfeedsresponseselect

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import responsemetaselect "github.com/21stio/go-playground/grpc/schema/responsemetaselect"
import feedsselect "github.com/21stio/go-playground/grpc/schema/feedsselect"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type StreamFeedsResponseSelect struct {
	SelectAll        *bool                                  `protobuf:"varint,1,opt,name=selectAll" json:"selectAll,omitempty"`
	Meta             *responsemetaselect.ResponseMetaSelect `protobuf:"bytes,2,opt,name=meta" json:"meta,omitempty"`
	TotalCount       *bool                                  `protobuf:"varint,3,opt,name=totalCount" json:"totalCount,omitempty"`
	Feeds            *feedsselect.FeedsSelect               `protobuf:"bytes,4,opt,name=feeds" json:"feeds,omitempty"`
	SelectHash       *bool                                  `protobuf:"varint,5,opt,name=selectHash" json:"selectHash,omitempty"`
	Hash             *string                                `protobuf:"bytes,6,opt,name=hash" json:"hash,omitempty"`
	XXX_unrecognized []byte                                 `json:"-"`
}

func (m *StreamFeedsResponseSelect) Reset()                    { *m = StreamFeedsResponseSelect{} }
func (m *StreamFeedsResponseSelect) String() string            { return proto.CompactTextString(m) }
func (*StreamFeedsResponseSelect) ProtoMessage()               {}
func (*StreamFeedsResponseSelect) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *StreamFeedsResponseSelect) GetSelectAll() bool {
	if m != nil && m.SelectAll != nil {
		return *m.SelectAll
	}
	return false
}

func (m *StreamFeedsResponseSelect) GetMeta() *responsemetaselect.ResponseMetaSelect {
	if m != nil {
		return m.Meta
	}
	return nil
}

func (m *StreamFeedsResponseSelect) GetTotalCount() bool {
	if m != nil && m.TotalCount != nil {
		return *m.TotalCount
	}
	return false
}

func (m *StreamFeedsResponseSelect) GetFeeds() *feedsselect.FeedsSelect {
	if m != nil {
		return m.Feeds
	}
	return nil
}

func (m *StreamFeedsResponseSelect) GetSelectHash() bool {
	if m != nil && m.SelectHash != nil {
		return *m.SelectHash
	}
	return false
}

func (m *StreamFeedsResponseSelect) GetHash() string {
	if m != nil && m.Hash != nil {
		return *m.Hash
	}
	return ""
}

func init() {
	proto.RegisterType((*StreamFeedsResponseSelect)(nil), "streamfeedsresponseselect.StreamFeedsResponseSelect")
}

func init() {
	proto.RegisterFile("streamfeedsresponseselect/streamfeedsresponseselect_.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 265 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x91, 0x3f, 0x4f, 0xc3, 0x30,
	0x10, 0xc5, 0x15, 0x48, 0x11, 0x35, 0x9b, 0x27, 0xb7, 0x42, 0x55, 0xc4, 0x80, 0x32, 0x80, 0x2d,
	0x3a, 0x76, 0x03, 0xc4, 0x9f, 0x85, 0x25, 0xdd, 0x58, 0x90, 0x49, 0x8f, 0xa4, 0x92, 0x13, 0x47,
	0xf6, 0x65, 0xe0, 0xa3, 0xb3, 0xa1, 0x9c, 0x5d, 0x35, 0x52, 0x95, 0xed, 0xfc, 0xde, 0xf3, 0xef,
	0x9e, 0x74, 0x6c, 0xe3, 0xd1, 0x81, 0x6e, 0x7e, 0x00, 0x76, 0xde, 0x81, 0xef, 0x6c, 0xeb, 0xc1,
	0x83, 0x81, 0x12, 0xd5, 0xa4, 0xf3, 0x25, 0x3b, 0x67, 0xd1, 0xf2, 0xc5, 0x64, 0x62, 0x79, 0x77,
	0x78, 0x37, 0x80, 0x3a, 0xf2, 0x4e, 0xa5, 0x08, 0x5a, 0xae, 0x08, 0x11, 0x63, 0xa3, 0x39, 0xfa,
	0x37, 0x7f, 0x09, 0x5b, 0x6c, 0x69, 0xd7, 0xeb, 0x60, 0x16, 0x11, 0xb4, 0xa5, 0x10, 0xbf, 0x66,
	0xf3, 0x10, 0x7f, 0x34, 0x46, 0x24, 0x59, 0x92, 0x5f, 0x16, 0x47, 0x81, 0x6f, 0x58, 0x3a, 0x2c,
	0x14, 0x67, 0x59, 0x92, 0x5f, 0xad, 0x6f, 0xe5, 0x69, 0x0b, 0x79, 0xe0, 0x7d, 0x00, 0xea, 0xc0,
	0x2c, 0xe8, 0x0f, 0x5f, 0x31, 0x86, 0x16, 0xb5, 0x79, 0xb6, 0x7d, 0x8b, 0xe2, 0x9c, 0xd0, 0x23,
	0x85, 0x4b, 0x36, 0xa3, 0xb6, 0x22, 0x25, 0xb8, 0x90, 0xa3, 0xee, 0x92, 0xaa, 0x46, 0x5c, 0x88,
	0x0d, 0xbc, 0x60, 0xbe, 0x6b, 0x5f, 0x8b, 0x59, 0xe0, 0x1d, 0x15, 0xce, 0x59, 0x5a, 0x0f, 0xce,
	0x45, 0x96, 0xe4, 0xf3, 0x82, 0xe6, 0xa7, 0xb7, 0xcf, 0x97, 0x6a, 0x8f, 0x75, 0xff, 0x2d, 0x4b,
	0xdb, 0xa8, 0xf5, 0x83, 0xc7, 0xbd, 0x55, 0x95, 0xbd, 0xef, 0x8c, 0xfe, 0xad, 0x9c, 0xed, 0xdb,
	0x9d, 0xaa, 0x5c, 0x57, 0x2a, 0x5f, 0xd6, 0xd0, 0xe8, 0xe9, 0xa3, 0xfd, 0x07, 0x00, 0x00, 0xff,
	0xff, 0x8d, 0x6b, 0x21, 0xa0, 0xea, 0x01, 0x00, 0x00,
}
