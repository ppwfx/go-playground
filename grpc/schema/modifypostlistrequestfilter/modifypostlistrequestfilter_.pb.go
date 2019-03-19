// Code generated by protoc-gen-go. DO NOT EDIT.
// source: modifypostlistrequestfilter/modifypostlistrequestfilter_.proto

/*
Package modifypostlistrequestfilter is a generated protocol buffer package.

It is generated from these files:
	modifypostlistrequestfilter/modifypostlistrequestfilter_.proto

It has these top-level messages:
	ModifyPostListRequestFilter
*/
package modifypostlistrequestfilter

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import requestmetafilter "github.com/21stio/go-playground/grpc/schema/requestmetafilter"
import postskindfilter "github.com/21stio/go-playground/grpc/schema/postskindfilter"
import idfilter "github.com/21stio/go-playground/grpc/schema/idfilter"
import listmodifierfilter "github.com/21stio/go-playground/grpc/schema/listmodifierfilter"
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

type ModifyPostListRequestFilter struct {
	Meta             *requestmetafilter.RequestMetaFilter   `protobuf:"bytes,1,opt,name=meta" json:"meta,omitempty"`
	Kind             *postskindfilter.PostsKindFilter       `protobuf:"bytes,2,opt,name=kind" json:"kind,omitempty"`
	Id               *idfilter.IdFilter                     `protobuf:"bytes,3,opt,name=id" json:"id,omitempty"`
	Modifier         *listmodifierfilter.ListModifierFilter `protobuf:"bytes,4,opt,name=modifier" json:"modifier,omitempty"`
	PostsSome        *feedfilter.PostFilter                 `protobuf:"bytes,5,opt,name=postsSome" json:"postsSome,omitempty"`
	PostsEvery       *feedfilter.PostFilter                 `protobuf:"bytes,6,opt,name=postsEvery" json:"postsEvery,omitempty"`
	PostsNone        *feedfilter.PostFilter                 `protobuf:"bytes,7,opt,name=postsNone" json:"postsNone,omitempty"`
	And              []*ModifyPostListRequestFilter         `protobuf:"bytes,8,rep,name=and" json:"and,omitempty"`
	Or               []*ModifyPostListRequestFilter         `protobuf:"bytes,9,rep,name=or" json:"or,omitempty"`
	Not              []*ModifyPostListRequestFilter         `protobuf:"bytes,10,rep,name=not" json:"not,omitempty"`
	Hash             *string                                `protobuf:"bytes,11,opt,name=hash" json:"hash,omitempty"`
	XXX_unrecognized []byte                                 `json:"-"`
}

func (m *ModifyPostListRequestFilter) Reset()                    { *m = ModifyPostListRequestFilter{} }
func (m *ModifyPostListRequestFilter) String() string            { return proto.CompactTextString(m) }
func (*ModifyPostListRequestFilter) ProtoMessage()               {}
func (*ModifyPostListRequestFilter) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *ModifyPostListRequestFilter) GetMeta() *requestmetafilter.RequestMetaFilter {
	if m != nil {
		return m.Meta
	}
	return nil
}

func (m *ModifyPostListRequestFilter) GetKind() *postskindfilter.PostsKindFilter {
	if m != nil {
		return m.Kind
	}
	return nil
}

func (m *ModifyPostListRequestFilter) GetId() *idfilter.IdFilter {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *ModifyPostListRequestFilter) GetModifier() *listmodifierfilter.ListModifierFilter {
	if m != nil {
		return m.Modifier
	}
	return nil
}

func (m *ModifyPostListRequestFilter) GetPostsSome() *feedfilter.PostFilter {
	if m != nil {
		return m.PostsSome
	}
	return nil
}

func (m *ModifyPostListRequestFilter) GetPostsEvery() *feedfilter.PostFilter {
	if m != nil {
		return m.PostsEvery
	}
	return nil
}

func (m *ModifyPostListRequestFilter) GetPostsNone() *feedfilter.PostFilter {
	if m != nil {
		return m.PostsNone
	}
	return nil
}

func (m *ModifyPostListRequestFilter) GetAnd() []*ModifyPostListRequestFilter {
	if m != nil {
		return m.And
	}
	return nil
}

func (m *ModifyPostListRequestFilter) GetOr() []*ModifyPostListRequestFilter {
	if m != nil {
		return m.Or
	}
	return nil
}

func (m *ModifyPostListRequestFilter) GetNot() []*ModifyPostListRequestFilter {
	if m != nil {
		return m.Not
	}
	return nil
}

func (m *ModifyPostListRequestFilter) GetHash() string {
	if m != nil && m.Hash != nil {
		return *m.Hash
	}
	return ""
}

func init() {
	proto.RegisterType((*ModifyPostListRequestFilter)(nil), "modifypostlistrequestfilter.ModifyPostListRequestFilter")
}

func init() {
	proto.RegisterFile("modifypostlistrequestfilter/modifypostlistrequestfilter_.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 399 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x93, 0xcf, 0x6b, 0x22, 0x31,
	0x1c, 0xc5, 0x71, 0x1c, 0x77, 0x35, 0xde, 0x72, 0x58, 0x82, 0xee, 0x61, 0x90, 0x45, 0x64, 0xd9,
	0xcd, 0xb0, 0x22, 0x8b, 0xa7, 0x1e, 0x84, 0xfe, 0xb0, 0xad, 0xa5, 0x4c, 0x6f, 0xbd, 0x94, 0xa9,
	0x13, 0x67, 0x42, 0x9d, 0xf9, 0x4e, 0x93, 0x58, 0xf0, 0xaf, 0xe8, 0xbf, 0x5c, 0x12, 0x13, 0x2d,
	0x0e, 0x0c, 0x05, 0x6f, 0x49, 0xbe, 0x9f, 0xf7, 0xf2, 0xde, 0xfc, 0x40, 0x67, 0x39, 0x24, 0x7c,
	0xb5, 0x2d, 0x41, 0xaa, 0x35, 0x97, 0x4a, 0xb0, 0xd7, 0x0d, 0x93, 0x6a, 0xc5, 0xd7, 0x8a, 0x89,
	0xb0, 0x66, 0xf6, 0x44, 0x4b, 0x01, 0x0a, 0x70, 0xbf, 0x86, 0xe9, 0xfd, 0xb6, 0xdb, 0x9c, 0xa9,
	0xd8, 0x5a, 0x56, 0x4e, 0xac, 0x51, 0x6f, 0xa8, 0x2d, 0xe4, 0x0b, 0x2f, 0x12, 0x4b, 0x1e, 0xed,
	0x1d, 0x47, 0xb8, 0x03, 0xf8, 0xd1, 0xe4, 0x8f, 0x0e, 0x60, 0xe2, 0x70, 0x26, 0x2c, 0x53, 0x3d,
	0x72, 0xf4, 0xcf, 0x15, 0x63, 0xce, 0xe9, 0xb0, 0xb4, 0xd3, 0xc1, 0x7b, 0x0b, 0xf5, 0x17, 0xa6,
	0xd9, 0x3d, 0x48, 0x75, 0xcb, 0xa5, 0x8a, 0x76, 0xc1, 0x2f, 0x0c, 0x86, 0xa7, 0xc8, 0xd7, 0x15,
	0x48, 0x23, 0x68, 0x8c, 0xba, 0xe3, 0x5f, 0xb4, 0x52, 0x8b, 0x5a, 0x7e, 0xc1, 0x54, 0xbc, 0xd3,
	0x44, 0x46, 0x81, 0x27, 0xc8, 0xd7, 0xa5, 0x88, 0x67, 0x94, 0x01, 0x3d, 0xaa, 0x49, 0xf5, 0x7d,
	0xf2, 0x86, 0x17, 0x89, 0x53, 0xe9, 0x19, 0x1e, 0x20, 0x8f, 0x27, 0xa4, 0x69, 0x34, 0x98, 0xba,
	0xe6, 0x74, 0xee, 0x28, 0x8f, 0x27, 0x78, 0x86, 0xda, 0xae, 0x2a, 0xf1, 0x0d, 0x39, 0xa4, 0xd5,
	0xfe, 0x54, 0x97, 0x59, 0xd8, 0x23, 0xab, 0xde, 0xeb, 0xf0, 0x04, 0x75, 0x4c, 0xa0, 0x07, 0xc8,
	0x19, 0x69, 0x19, 0x93, 0x1f, 0xf4, 0xf0, 0x78, 0x4c, 0x3a, 0x2b, 0x3a, 0x80, 0xf8, 0x3f, 0x42,
	0x66, 0x73, 0xfe, 0xc6, 0xc4, 0x96, 0x7c, 0xab, 0x95, 0x7d, 0x22, 0xf7, 0xb7, 0xdd, 0x41, 0xc1,
	0xc8, 0xf7, 0x2f, 0xdc, 0xa6, 0x41, 0x7c, 0x8d, 0x9a, 0x71, 0x91, 0x90, 0x76, 0xd0, 0x1c, 0x75,
	0xc7, 0x53, 0x5a, 0xf3, 0x01, 0xd2, 0x9a, 0x57, 0x18, 0x69, 0x13, 0x7c, 0x85, 0x3c, 0x10, 0xa4,
	0x73, 0xa2, 0x95, 0x07, 0x42, 0xa7, 0x2a, 0x40, 0x11, 0x74, 0x6a, 0xaa, 0x02, 0x14, 0xc6, 0xc8,
	0xcf, 0x62, 0x99, 0x91, 0x6e, 0xd0, 0x18, 0x75, 0x22, 0xb3, 0x9e, 0xcd, 0x1f, 0x2f, 0x53, 0xae,
	0xb2, 0xcd, 0x33, 0x5d, 0x42, 0x1e, 0x8e, 0xff, 0x49, 0xc5, 0x21, 0x4c, 0xe1, 0x6f, 0xb9, 0x8e,
	0xb7, 0xa9, 0x80, 0x4d, 0x91, 0x84, 0xa9, 0x28, 0x97, 0xa1, 0x5c, 0x66, 0x2c, 0x8f, 0xeb, 0x7e,
	0xdd, 0x8f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x87, 0x52, 0x04, 0xe9, 0xf4, 0x03, 0x00, 0x00,
}
