// Code generated by protoc-gen-go. DO NOT EDIT.
// source: createpostsresponse/createpostsresponse_.proto

/*
Package createpostsresponse is a generated protocol buffer package.

It is generated from these files:
	createpostsresponse/createpostsresponse_.proto

It has these top-level messages:
	CreatePostsResponse
*/
package createpostsresponse

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import responsemeta "github.com/21stio/go-playground/grpc/schema/responsemeta"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type CreatePostsResponse struct {
	Meta             *responsemeta.ResponseMeta `protobuf:"bytes,1,opt,name=meta" json:"meta,omitempty"`
	Hash             *string                    `protobuf:"bytes,2,opt,name=hash" json:"hash,omitempty"`
	XXX_unrecognized []byte                     `json:"-"`
}

func (m *CreatePostsResponse) Reset()                    { *m = CreatePostsResponse{} }
func (m *CreatePostsResponse) String() string            { return proto.CompactTextString(m) }
func (*CreatePostsResponse) ProtoMessage()               {}
func (*CreatePostsResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *CreatePostsResponse) GetMeta() *responsemeta.ResponseMeta {
	if m != nil {
		return m.Meta
	}
	return nil
}

func (m *CreatePostsResponse) GetHash() string {
	if m != nil && m.Hash != nil {
		return *m.Hash
	}
	return ""
}

func init() {
	proto.RegisterType((*CreatePostsResponse)(nil), "createpostsresponse.CreatePostsResponse")
}

func init() { proto.RegisterFile("createpostsresponse/createpostsresponse_.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 178 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0x4b, 0x2e, 0x4a, 0x4d,
	0x2c, 0x49, 0x2d, 0xc8, 0x2f, 0x2e, 0x29, 0x2e, 0x4a, 0x2d, 0x2e, 0xc8, 0xcf, 0x2b, 0x4e, 0xd5,
	0xc7, 0x22, 0x16, 0xaf, 0x57, 0x50, 0x94, 0x5f, 0x92, 0x2f, 0x24, 0x8c, 0x45, 0x4e, 0x4a, 0x01,
	0xc6, 0xca, 0x4d, 0x2d, 0x49, 0xd4, 0x47, 0xe6, 0x40, 0xb5, 0x29, 0x45, 0x72, 0x09, 0x3b, 0x83,
	0x35, 0x06, 0x80, 0x34, 0x06, 0x41, 0x55, 0x08, 0xe9, 0x71, 0xb1, 0x80, 0x54, 0x49, 0x30, 0x2a,
	0x30, 0x6a, 0x70, 0x1b, 0x49, 0xe9, 0x21, 0x6b, 0xd5, 0x83, 0xa9, 0xf2, 0x4d, 0x2d, 0x49, 0x0c,
	0x02, 0xab, 0x13, 0x12, 0xe2, 0x62, 0xc9, 0x48, 0x2c, 0xce, 0x90, 0x60, 0x52, 0x60, 0xd4, 0xe0,
	0x0c, 0x02, 0xb3, 0x9d, 0x1c, 0xa3, 0xec, 0xd3, 0x33, 0x4b, 0x32, 0x4a, 0x93, 0xf4, 0x92, 0xf3,
	0x73, 0xf5, 0x8d, 0x0c, 0x8b, 0x4b, 0x32, 0xf3, 0xf5, 0xd3, 0xf3, 0x75, 0x0b, 0x72, 0x12, 0x2b,
	0xd3, 0x8b, 0xf2, 0x4b, 0xf3, 0x52, 0xf4, 0xd3, 0x8b, 0x0a, 0x92, 0xf5, 0x8b, 0x93, 0x33, 0x52,
	0x73, 0x13, 0xb1, 0xf9, 0x0d, 0x10, 0x00, 0x00, 0xff, 0xff, 0xf6, 0x67, 0x92, 0x89, 0x05, 0x01,
	0x00, 0x00,
}
