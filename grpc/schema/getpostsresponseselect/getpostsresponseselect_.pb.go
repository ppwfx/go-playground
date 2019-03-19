// Code generated by protoc-gen-go. DO NOT EDIT.
// source: getpostsresponseselect/getpostsresponseselect_.proto

/*
Package getpostsresponseselect is a generated protocol buffer package.

It is generated from these files:
	getpostsresponseselect/getpostsresponseselect_.proto

It has these top-level messages:
	GetPostsResponseSelect
*/
package getpostsresponseselect

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import responsemetaselect "github.com/21stio/go-playground/grpc/schema/responsemetaselect"
import feedselect "github.com/21stio/go-playground/grpc/schema/feedselect"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type GetPostsResponseSelect struct {
	SelectAll        *bool                                  `protobuf:"varint,1,opt,name=selectAll" json:"selectAll,omitempty"`
	Meta             *responsemetaselect.ResponseMetaSelect `protobuf:"bytes,2,opt,name=meta" json:"meta,omitempty"`
	TotalCount       *bool                                  `protobuf:"varint,3,opt,name=totalCount" json:"totalCount,omitempty"`
	Posts            *feedselect.PostsSelect                `protobuf:"bytes,4,opt,name=posts" json:"posts,omitempty"`
	SelectHash       *bool                                  `protobuf:"varint,5,opt,name=selectHash" json:"selectHash,omitempty"`
	Hash             *string                                `protobuf:"bytes,6,opt,name=hash" json:"hash,omitempty"`
	XXX_unrecognized []byte                                 `json:"-"`
}

func (m *GetPostsResponseSelect) Reset()                    { *m = GetPostsResponseSelect{} }
func (m *GetPostsResponseSelect) String() string            { return proto.CompactTextString(m) }
func (*GetPostsResponseSelect) ProtoMessage()               {}
func (*GetPostsResponseSelect) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *GetPostsResponseSelect) GetSelectAll() bool {
	if m != nil && m.SelectAll != nil {
		return *m.SelectAll
	}
	return false
}

func (m *GetPostsResponseSelect) GetMeta() *responsemetaselect.ResponseMetaSelect {
	if m != nil {
		return m.Meta
	}
	return nil
}

func (m *GetPostsResponseSelect) GetTotalCount() bool {
	if m != nil && m.TotalCount != nil {
		return *m.TotalCount
	}
	return false
}

func (m *GetPostsResponseSelect) GetPosts() *feedselect.PostsSelect {
	if m != nil {
		return m.Posts
	}
	return nil
}

func (m *GetPostsResponseSelect) GetSelectHash() bool {
	if m != nil && m.SelectHash != nil {
		return *m.SelectHash
	}
	return false
}

func (m *GetPostsResponseSelect) GetHash() string {
	if m != nil && m.Hash != nil {
		return *m.Hash
	}
	return ""
}

func init() {
	proto.RegisterType((*GetPostsResponseSelect)(nil), "getpostsresponseselect.GetPostsResponseSelect")
}

func init() {
	proto.RegisterFile("getpostsresponseselect/getpostsresponseselect_.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 266 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x91, 0xc1, 0x4b, 0xc3, 0x30,
	0x14, 0xc6, 0xa9, 0x76, 0xe2, 0xe2, 0x2d, 0x87, 0x59, 0xc6, 0x90, 0xe2, 0x41, 0x7a, 0x70, 0x09,
	0x0e, 0x4f, 0xde, 0x9c, 0x82, 0x5e, 0x04, 0xa9, 0x37, 0x2f, 0x12, 0xbb, 0x67, 0x3a, 0x48, 0xfb,
	0x42, 0xf3, 0x7a, 0xf0, 0x0f, 0xf7, 0x2e, 0x4d, 0x32, 0x2a, 0x6c, 0xde, 0x5e, 0xbe, 0xef, 0xcb,
	0xef, 0x7d, 0x21, 0xec, 0x56, 0x03, 0x59, 0x74, 0xe4, 0x3a, 0x70, 0x16, 0x5b, 0x07, 0x0e, 0x0c,
	0x54, 0x24, 0x0f, 0xcb, 0x1f, 0xc2, 0x76, 0x48, 0xc8, 0x67, 0x87, 0xed, 0xf9, 0xf5, 0xee, 0xdc,
	0x00, 0xa9, 0x48, 0xda, 0x97, 0x22, 0x65, 0xbe, 0xf8, 0x02, 0xd8, 0xc4, 0xd4, 0x38, 0x46, 0xf7,
	0xf2, 0x27, 0x61, 0xb3, 0x27, 0xa0, 0xd7, 0x61, 0x4d, 0x19, 0x19, 0x6f, 0x3e, 0xc1, 0x17, 0x6c,
	0x1a, 0xb2, 0xf7, 0xc6, 0x64, 0x49, 0x9e, 0x14, 0xa7, 0xe5, 0x28, 0xf0, 0x3b, 0x96, 0x0e, 0xbb,
	0xb2, 0xa3, 0x3c, 0x29, 0xce, 0x56, 0x57, 0x62, 0xbf, 0x80, 0xd8, 0xf1, 0x5e, 0x80, 0x54, 0x60,
	0x96, 0xfe, 0x0e, 0xbf, 0x60, 0x8c, 0x90, 0x94, 0x79, 0xc0, 0xbe, 0xa5, 0xec, 0xd8, 0xa3, 0xff,
	0x28, 0x7c, 0xc9, 0x26, 0xfe, 0xdd, 0x59, 0xea, 0xe1, 0xe7, 0x62, 0xec, 0x2d, 0x7c, 0xd3, 0x48,
	0x0b, 0xa9, 0x01, 0x17, 0xcc, 0x67, 0xe5, 0xea, 0x6c, 0x12, 0x70, 0xa3, 0xc2, 0x39, 0x4b, 0xeb,
	0xc1, 0x39, 0xc9, 0x93, 0x62, 0x5a, 0xfa, 0x79, 0xfd, 0xf8, 0xbe, 0xd6, 0x5b, 0xaa, 0xfb, 0x4f,
	0x51, 0x61, 0x23, 0x57, 0x37, 0x8e, 0xb6, 0x28, 0x35, 0x2e, 0xad, 0x51, 0xdf, 0xba, 0xc3, 0xbe,
	0xdd, 0x48, 0xdd, 0xd9, 0x4a, 0xba, 0xaa, 0x86, 0x46, 0xfd, 0xf3, 0x51, 0xbf, 0x01, 0x00, 0x00,
	0xff, 0xff, 0x9b, 0xf1, 0xac, 0x8a, 0xd8, 0x01, 0x00, 0x00,
}
