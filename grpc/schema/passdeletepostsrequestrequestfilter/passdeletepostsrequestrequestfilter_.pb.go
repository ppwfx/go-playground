// Code generated by protoc-gen-go. DO NOT EDIT.
// source: passdeletepostsrequestrequestfilter/passdeletepostsrequestrequestfilter_.proto

/*
Package passdeletepostsrequestrequestfilter is a generated protocol buffer package.

It is generated from these files:
	passdeletepostsrequestrequestfilter/passdeletepostsrequestrequestfilter_.proto

It has these top-level messages:
	PassDeletePostsRequestRequestFilter
*/
package passdeletepostsrequestrequestfilter

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import requestmetafilter "github.com/21stio/go-playground/grpc/schema/requestmetafilter"
import deletepostsrequestfilter "github.com/21stio/go-playground/grpc/schema/deletepostsrequestfilter"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type PassDeletePostsRequestRequestFilter struct {
	Meta               *requestmetafilter.RequestMetaFilter               `protobuf:"bytes,1,opt,name=meta" json:"meta,omitempty"`
	DeletePostsRequest *deletepostsrequestfilter.DeletePostsRequestFilter `protobuf:"bytes,2,opt,name=deletePostsRequest" json:"deletePostsRequest,omitempty"`
	And                []*PassDeletePostsRequestRequestFilter             `protobuf:"bytes,3,rep,name=and" json:"and,omitempty"`
	Or                 []*PassDeletePostsRequestRequestFilter             `protobuf:"bytes,4,rep,name=or" json:"or,omitempty"`
	Not                []*PassDeletePostsRequestRequestFilter             `protobuf:"bytes,5,rep,name=not" json:"not,omitempty"`
	Hash               *string                                            `protobuf:"bytes,6,opt,name=hash" json:"hash,omitempty"`
	XXX_unrecognized   []byte                                             `json:"-"`
}

func (m *PassDeletePostsRequestRequestFilter) Reset()         { *m = PassDeletePostsRequestRequestFilter{} }
func (m *PassDeletePostsRequestRequestFilter) String() string { return proto.CompactTextString(m) }
func (*PassDeletePostsRequestRequestFilter) ProtoMessage()    {}
func (*PassDeletePostsRequestRequestFilter) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{0}
}

func (m *PassDeletePostsRequestRequestFilter) GetMeta() *requestmetafilter.RequestMetaFilter {
	if m != nil {
		return m.Meta
	}
	return nil
}

func (m *PassDeletePostsRequestRequestFilter) GetDeletePostsRequest() *deletepostsrequestfilter.DeletePostsRequestFilter {
	if m != nil {
		return m.DeletePostsRequest
	}
	return nil
}

func (m *PassDeletePostsRequestRequestFilter) GetAnd() []*PassDeletePostsRequestRequestFilter {
	if m != nil {
		return m.And
	}
	return nil
}

func (m *PassDeletePostsRequestRequestFilter) GetOr() []*PassDeletePostsRequestRequestFilter {
	if m != nil {
		return m.Or
	}
	return nil
}

func (m *PassDeletePostsRequestRequestFilter) GetNot() []*PassDeletePostsRequestRequestFilter {
	if m != nil {
		return m.Not
	}
	return nil
}

func (m *PassDeletePostsRequestRequestFilter) GetHash() string {
	if m != nil && m.Hash != nil {
		return *m.Hash
	}
	return ""
}

func init() {
	proto.RegisterType((*PassDeletePostsRequestRequestFilter)(nil), "passdeletepostsrequestrequestfilter.PassDeletePostsRequestRequestFilter")
}

func init() {
	proto.RegisterFile("passdeletepostsrequestrequestfilter/passdeletepostsrequestrequestfilter_.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 279 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x92, 0x3d, 0x4f, 0x84, 0x40,
	0x10, 0x86, 0xc3, 0x87, 0x26, 0xee, 0x75, 0x5b, 0x6d, 0xae, 0x22, 0x9e, 0x05, 0x31, 0x71, 0x89,
	0x54, 0x57, 0x1b, 0x63, 0x6c, 0xd4, 0x93, 0xca, 0x5c, 0x63, 0xf6, 0x60, 0x05, 0x12, 0x60, 0x70,
	0x67, 0x28, 0xfc, 0xc9, 0xfe, 0x0b, 0x03, 0xb7, 0xdd, 0x7a, 0xc9, 0x16, 0x57, 0x0d, 0x1f, 0xcf,
	0xfb, 0xcc, 0x0b, 0x59, 0xf6, 0x3a, 0x2a, 0xc4, 0x4a, 0x77, 0x9a, 0xf4, 0x08, 0x48, 0x68, 0xf4,
	0xf7, 0xa4, 0x91, 0xec, 0xf8, 0x6a, 0x3b, 0xd2, 0x26, 0xf3, 0x60, 0x3e, 0xe5, 0x68, 0x80, 0x80,
	0x6f, 0x3c, 0xd8, 0xf5, 0xad, 0xbd, 0xed, 0x35, 0x29, 0xbb, 0xc2, 0x79, 0x62, 0x85, 0xeb, 0xad,
	0x2b, 0xb3, 0x91, 0x53, 0x2f, 0x6c, 0xf2, 0xfa, 0x37, 0x62, 0x9b, 0x9d, 0x42, 0x7c, 0x5c, 0xb8,
	0xdd, 0xcc, 0x15, 0x47, 0xce, 0x8e, 0xa7, 0x05, 0xe7, 0x5b, 0x16, 0xcf, 0x6b, 0x45, 0x90, 0x04,
	0xe9, 0x2a, 0xbf, 0x91, 0x4e, 0x15, 0x69, 0xf9, 0x17, 0x4d, 0xea, 0x98, 0x29, 0x96, 0x04, 0x3f,
	0x30, 0x5e, 0x39, 0x72, 0x11, 0x2e, 0x9e, 0x5c, 0x9e, 0xea, 0x27, 0xdd, 0x42, 0xd6, 0xfa, 0x8f,
	0x8d, 0xef, 0x59, 0xa4, 0x86, 0x4a, 0x44, 0x49, 0x94, 0xae, 0xf2, 0x67, 0xe9, 0xf1, 0x7b, 0xa5,
	0xc7, 0x47, 0x17, 0xb3, 0x94, 0x7f, 0xb0, 0x10, 0x8c, 0x88, 0xcf, 0xac, 0x0e, 0xc1, 0xcc, 0xad,
	0x07, 0x20, 0x71, 0x71, 0xee, 0xd6, 0x03, 0x10, 0xe7, 0x2c, 0x6e, 0x14, 0x36, 0xe2, 0x32, 0x09,
	0xd2, 0xab, 0x62, 0xb9, 0x7e, 0x78, 0xdf, 0xbf, 0xd5, 0x2d, 0x35, 0xd3, 0x41, 0x96, 0xd0, 0x67,
	0xf9, 0x3d, 0x52, 0x0b, 0x59, 0x0d, 0x77, 0x63, 0xa7, 0x7e, 0x6a, 0x03, 0xd3, 0x50, 0x65, 0xb5,
	0x19, 0xcb, 0x0c, 0xcb, 0x46, 0xf7, 0xca, 0xe7, 0x40, 0xff, 0x05, 0x00, 0x00, 0xff, 0xff, 0x8a,
	0x23, 0x6c, 0xc0, 0x1a, 0x03, 0x00, 0x00,
}
