// Code generated by protoc-gen-go. DO NOT EDIT.
// source: createpostsrequestfilter/createpostsrequestfilter_.proto

/*
Package createpostsrequestfilter is a generated protocol buffer package.

It is generated from these files:
	createpostsrequestfilter/createpostsrequestfilter_.proto

It has these top-level messages:
	CreatePostsRequestFilter
*/
package createpostsrequestfilter

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import requestmetafilter "github.com/21stio/go-playground/grpc/schema/requestmetafilter"
import feedfilter "github.com/21stio/go-playground/grpc/schema/feedfilter"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type CreatePostsRequestFilter struct {
	Meta             *requestmetafilter.RequestMetaFilter `protobuf:"bytes,1,opt,name=meta" json:"meta,omitempty"`
	PostsSome        *feedfilter.PostFilter               `protobuf:"bytes,2,opt,name=postsSome" json:"postsSome,omitempty"`
	PostsEvery       *feedfilter.PostFilter               `protobuf:"bytes,3,opt,name=postsEvery" json:"postsEvery,omitempty"`
	PostsNone        *feedfilter.PostFilter               `protobuf:"bytes,4,opt,name=postsNone" json:"postsNone,omitempty"`
	And              []*CreatePostsRequestFilter          `protobuf:"bytes,5,rep,name=and" json:"and,omitempty"`
	Or               []*CreatePostsRequestFilter          `protobuf:"bytes,6,rep,name=or" json:"or,omitempty"`
	Not              []*CreatePostsRequestFilter          `protobuf:"bytes,7,rep,name=not" json:"not,omitempty"`
	Hash             *string                              `protobuf:"bytes,8,opt,name=hash" json:"hash,omitempty"`
	XXX_unrecognized []byte                               `json:"-"`
}

func (m *CreatePostsRequestFilter) Reset()                    { *m = CreatePostsRequestFilter{} }
func (m *CreatePostsRequestFilter) String() string            { return proto.CompactTextString(m) }
func (*CreatePostsRequestFilter) ProtoMessage()               {}
func (*CreatePostsRequestFilter) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *CreatePostsRequestFilter) GetMeta() *requestmetafilter.RequestMetaFilter {
	if m != nil {
		return m.Meta
	}
	return nil
}

func (m *CreatePostsRequestFilter) GetPostsSome() *feedfilter.PostFilter {
	if m != nil {
		return m.PostsSome
	}
	return nil
}

func (m *CreatePostsRequestFilter) GetPostsEvery() *feedfilter.PostFilter {
	if m != nil {
		return m.PostsEvery
	}
	return nil
}

func (m *CreatePostsRequestFilter) GetPostsNone() *feedfilter.PostFilter {
	if m != nil {
		return m.PostsNone
	}
	return nil
}

func (m *CreatePostsRequestFilter) GetAnd() []*CreatePostsRequestFilter {
	if m != nil {
		return m.And
	}
	return nil
}

func (m *CreatePostsRequestFilter) GetOr() []*CreatePostsRequestFilter {
	if m != nil {
		return m.Or
	}
	return nil
}

func (m *CreatePostsRequestFilter) GetNot() []*CreatePostsRequestFilter {
	if m != nil {
		return m.Not
	}
	return nil
}

func (m *CreatePostsRequestFilter) GetHash() string {
	if m != nil && m.Hash != nil {
		return *m.Hash
	}
	return ""
}

func init() {
	proto.RegisterType((*CreatePostsRequestFilter)(nil), "createpostsrequestfilter.CreatePostsRequestFilter")
}

func init() {
	proto.RegisterFile("createpostsrequestfilter/createpostsrequestfilter_.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 298 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x92, 0x4f, 0x4b, 0xc3, 0x40,
	0x10, 0xc5, 0xc9, 0x1f, 0xff, 0x74, 0x7b, 0xdb, 0x83, 0x2c, 0xc5, 0x43, 0x10, 0x0f, 0x41, 0x70,
	0x83, 0x41, 0xa4, 0xe7, 0x5a, 0x7b, 0x53, 0x24, 0xde, 0xbc, 0xc8, 0x9a, 0x4c, 0x93, 0x40, 0x93,
	0x89, 0xbb, 0x1b, 0xa1, 0xdf, 0xd5, 0x0f, 0x23, 0xbb, 0x5d, 0x8c, 0x50, 0x22, 0xd2, 0xdb, 0xcc,
	0xce, 0x7b, 0xef, 0x37, 0x0c, 0x4b, 0xe6, 0xb9, 0x04, 0xa1, 0xa1, 0x43, 0xa5, 0x95, 0x84, 0x8f,
	0x1e, 0x94, 0x5e, 0xd7, 0x1b, 0x0d, 0x32, 0x19, 0x1b, 0xbc, 0xf1, 0x4e, 0xa2, 0x46, 0xca, 0xc6,
	0x04, 0xb3, 0x2b, 0xd7, 0x36, 0xa0, 0x85, 0x0b, 0xdb, 0x7b, 0x71, 0x29, 0xb3, 0xf3, 0x35, 0x40,
	0xe1, 0x44, 0x43, 0xe9, 0xa6, 0x17, 0x5f, 0x01, 0x61, 0xf7, 0x16, 0xf3, 0x6c, 0x30, 0xd9, 0x2e,
	0x65, 0x65, 0x35, 0x74, 0x4e, 0x42, 0x93, 0xc7, 0xbc, 0xc8, 0x8b, 0xa7, 0xe9, 0x25, 0xdf, 0x63,
	0x70, 0xa7, 0x7f, 0x04, 0x2d, 0x76, 0x9e, 0xcc, 0x3a, 0xe8, 0x2d, 0x99, 0xd8, 0xb5, 0x5f, 0xb0,
	0x01, 0xe6, 0x5b, 0xfb, 0x19, 0x1f, 0xe8, 0xdc, 0xc0, 0x9c, 0x61, 0x10, 0xd2, 0x3b, 0x42, 0x6c,
	0xf3, 0xf0, 0x09, 0x72, 0xcb, 0x82, 0x3f, 0x6d, 0xbf, 0x94, 0x3f, 0xb4, 0x27, 0x6c, 0x81, 0x85,
	0xff, 0xa0, 0x19, 0x21, 0x5d, 0x92, 0x40, 0xb4, 0x05, 0x3b, 0x8a, 0x82, 0x78, 0x9a, 0xa6, 0x7c,
	0xec, 0xd8, 0x7c, 0xec, 0x3c, 0x99, 0xb1, 0xd3, 0x05, 0xf1, 0x51, 0xb2, 0xe3, 0x83, 0x43, 0x7c,
	0x94, 0x66, 0x93, 0x16, 0x35, 0x3b, 0x39, 0x7c, 0x93, 0x16, 0x35, 0xa5, 0x24, 0xac, 0x84, 0xaa,
	0xd8, 0x69, 0xe4, 0xc5, 0x93, 0xcc, 0xd6, 0x8b, 0xd5, 0xeb, 0xb2, 0xac, 0x75, 0xd5, 0xbf, 0xf3,
	0x1c, 0x9b, 0x24, 0xbd, 0x51, 0xba, 0xc6, 0xa4, 0xc4, 0xeb, 0x6e, 0x23, 0xb6, 0xa5, 0xc4, 0xbe,
	0x2d, 0x92, 0x52, 0x76, 0x79, 0xa2, 0xf2, 0x0a, 0x1a, 0x31, 0xfa, 0x23, 0xbf, 0x03, 0x00, 0x00,
	0xff, 0xff, 0x69, 0x18, 0xe1, 0x8d, 0xc5, 0x02, 0x00, 0x00,
}
