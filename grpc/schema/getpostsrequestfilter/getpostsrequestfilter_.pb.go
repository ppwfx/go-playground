// Code generated by protoc-gen-go. DO NOT EDIT.
// source: getpostsrequestfilter/getpostsrequestfilter_.proto

/*
Package getpostsrequestfilter is a generated protocol buffer package.

It is generated from these files:
	getpostsrequestfilter/getpostsrequestfilter_.proto

It has these top-level messages:
	GetPostsRequestFilter
*/
package getpostsrequestfilter

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import requestmetafilter "github.com/21stio/go-playground/grpc/schema/requestmetafilter"
import locationqueryfilter "github.com/21stio/go-playground/grpc/schema/locationqueryfilter"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type GetPostsRequestFilter struct {
	Meta             *requestmetafilter.RequestMetaFilter     `protobuf:"bytes,1,opt,name=meta" json:"meta,omitempty"`
	LocationQuery    *locationqueryfilter.LocationQueryFilter `protobuf:"bytes,2,opt,name=locationQuery" json:"locationQuery,omitempty"`
	And              []*GetPostsRequestFilter                 `protobuf:"bytes,3,rep,name=and" json:"and,omitempty"`
	Or               []*GetPostsRequestFilter                 `protobuf:"bytes,4,rep,name=or" json:"or,omitempty"`
	Not              []*GetPostsRequestFilter                 `protobuf:"bytes,5,rep,name=not" json:"not,omitempty"`
	Hash             *string                                  `protobuf:"bytes,6,opt,name=hash" json:"hash,omitempty"`
	XXX_unrecognized []byte                                   `json:"-"`
}

func (m *GetPostsRequestFilter) Reset()                    { *m = GetPostsRequestFilter{} }
func (m *GetPostsRequestFilter) String() string            { return proto.CompactTextString(m) }
func (*GetPostsRequestFilter) ProtoMessage()               {}
func (*GetPostsRequestFilter) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *GetPostsRequestFilter) GetMeta() *requestmetafilter.RequestMetaFilter {
	if m != nil {
		return m.Meta
	}
	return nil
}

func (m *GetPostsRequestFilter) GetLocationQuery() *locationqueryfilter.LocationQueryFilter {
	if m != nil {
		return m.LocationQuery
	}
	return nil
}

func (m *GetPostsRequestFilter) GetAnd() []*GetPostsRequestFilter {
	if m != nil {
		return m.And
	}
	return nil
}

func (m *GetPostsRequestFilter) GetOr() []*GetPostsRequestFilter {
	if m != nil {
		return m.Or
	}
	return nil
}

func (m *GetPostsRequestFilter) GetNot() []*GetPostsRequestFilter {
	if m != nil {
		return m.Not
	}
	return nil
}

func (m *GetPostsRequestFilter) GetHash() string {
	if m != nil && m.Hash != nil {
		return *m.Hash
	}
	return ""
}

func init() {
	proto.RegisterType((*GetPostsRequestFilter)(nil), "getpostsrequestfilter.GetPostsRequestFilter")
}

func init() { proto.RegisterFile("getpostsrequestfilter/getpostsrequestfilter_.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 277 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x92, 0x31, 0x4b, 0xc4, 0x30,
	0x18, 0x86, 0xb9, 0xb6, 0x0a, 0xe6, 0x70, 0x09, 0x1c, 0x84, 0x9b, 0x8a, 0x38, 0x14, 0xd1, 0x04,
	0x3b, 0x39, 0x88, 0xa0, 0x82, 0x2e, 0x2a, 0xda, 0xd1, 0x45, 0x62, 0x2f, 0xa6, 0x85, 0x36, 0x5f,
	0x2f, 0xf9, 0x3a, 0xdc, 0xef, 0xf6, 0x0f, 0x48, 0x7a, 0x11, 0x94, 0xcb, 0x74, 0x5b, 0xfb, 0xe5,
	0x7d, 0x9f, 0x27, 0x09, 0x21, 0xa5, 0x56, 0x38, 0x80, 0x43, 0x67, 0xd5, 0x7a, 0x54, 0x0e, 0xbf,
	0xda, 0x0e, 0x95, 0x15, 0xd1, 0xe9, 0x07, 0x1f, 0x2c, 0x20, 0xd0, 0x45, 0x74, 0x75, 0x79, 0x16,
	0x7e, 0x7b, 0x85, 0x32, 0x60, 0x76, 0x26, 0x01, 0xb1, 0xe4, 0x1d, 0xd4, 0x12, 0x5b, 0x30, 0xeb,
	0x51, 0xd9, 0x4d, 0x48, 0x47, 0x66, 0x21, 0x7f, 0xf2, 0x9d, 0x90, 0xc5, 0xa3, 0xc2, 0x57, 0x6f,
	0xad, 0xb6, 0xd0, 0x87, 0x29, 0x40, 0xaf, 0x48, 0xe6, 0xf1, 0x6c, 0x96, 0xcf, 0x8a, 0x79, 0x79,
	0xca, 0x77, 0x94, 0x3c, 0xe4, 0x9f, 0x15, 0xca, 0x6d, 0xa7, 0x9a, 0x1a, 0xf4, 0x85, 0x1c, 0xff,
	0x1a, 0xdf, 0xbc, 0x91, 0x25, 0x13, 0xa2, 0x88, 0xed, 0x8d, 0x3f, 0xfd, 0x4d, 0x06, 0xcc, 0xff,
	0x3a, 0xbd, 0x21, 0xa9, 0x34, 0x2b, 0x96, 0xe6, 0x69, 0x31, 0x2f, 0xcf, 0x79, 0xf4, 0x92, 0x78,
	0xf4, 0x10, 0x95, 0x2f, 0xd2, 0x6b, 0x92, 0x80, 0x65, 0xd9, 0x1e, 0xf5, 0x04, 0xac, 0xb7, 0x1b,
	0x40, 0x76, 0xb0, 0x8f, 0xdd, 0x00, 0x52, 0x4a, 0xb2, 0x46, 0xba, 0x86, 0x1d, 0xe6, 0xb3, 0xe2,
	0xa8, 0x9a, 0xbe, 0xef, 0xee, 0xdf, 0x6f, 0x75, 0x8b, 0xcd, 0xf8, 0xc9, 0x6b, 0xe8, 0x45, 0x79,
	0xe9, 0xb0, 0x05, 0xa1, 0xe1, 0x62, 0xe8, 0xe4, 0x46, 0x5b, 0x18, 0xcd, 0x4a, 0x68, 0x3b, 0xd4,
	0xc2, 0xd5, 0x8d, 0xea, 0x65, 0xfc, 0xd1, 0xfc, 0x04, 0x00, 0x00, 0xff, 0xff, 0x1d, 0x76, 0x6a,
	0x26, 0x62, 0x02, 0x00, 0x00,
}