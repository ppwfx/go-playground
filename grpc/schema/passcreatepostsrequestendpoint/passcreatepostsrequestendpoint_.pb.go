// Code generated by protoc-gen-go. DO NOT EDIT.
// source: passcreatepostsrequestendpoint/passcreatepostsrequestendpoint_.proto

/*
Package passcreatepostsrequestendpoint is a generated protocol buffer package.

It is generated from these files:
	passcreatepostsrequestendpoint/passcreatepostsrequestendpoint_.proto

It has these top-level messages:
	PassCreatePostsRequestEndpoint
*/
package passcreatepostsrequestendpoint

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import passcreatepostsrequestrequestfilter "github.com/21stio/go-playground/grpc/schema/passcreatepostsrequestrequestfilter"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type PassCreatePostsRequestEndpoint struct {
	Filter           *passcreatepostsrequestrequestfilter.PassCreatePostsRequestRequestFilter `protobuf:"bytes,1,opt,name=filter" json:"filter,omitempty"`
	Hash             *string                                                                  `protobuf:"bytes,2,opt,name=hash" json:"hash,omitempty"`
	XXX_unrecognized []byte                                                                   `json:"-"`
}

func (m *PassCreatePostsRequestEndpoint) Reset()                    { *m = PassCreatePostsRequestEndpoint{} }
func (m *PassCreatePostsRequestEndpoint) String() string            { return proto.CompactTextString(m) }
func (*PassCreatePostsRequestEndpoint) ProtoMessage()               {}
func (*PassCreatePostsRequestEndpoint) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *PassCreatePostsRequestEndpoint) GetFilter() *passcreatepostsrequestrequestfilter.PassCreatePostsRequestRequestFilter {
	if m != nil {
		return m.Filter
	}
	return nil
}

func (m *PassCreatePostsRequestEndpoint) GetHash() string {
	if m != nil && m.Hash != nil {
		return *m.Hash
	}
	return ""
}

func init() {
	proto.RegisterType((*PassCreatePostsRequestEndpoint)(nil), "passcreatepostsrequestendpoint.PassCreatePostsRequestEndpoint")
}

func init() {
	proto.RegisterFile("passcreatepostsrequestendpoint/passcreatepostsrequestendpoint_.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 204 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x72, 0x29, 0x48, 0x2c, 0x2e,
	0x4e, 0x2e, 0x4a, 0x4d, 0x2c, 0x49, 0x2d, 0xc8, 0x2f, 0x2e, 0x29, 0x2e, 0x4a, 0x2d, 0x2c, 0x4d,
	0x2d, 0x2e, 0x49, 0xcd, 0x4b, 0x29, 0xc8, 0xcf, 0xcc, 0x2b, 0xd1, 0xc7, 0x2f, 0x1d, 0xaf, 0x57,
	0x50, 0x94, 0x5f, 0x92, 0x2f, 0x24, 0x87, 0x5f, 0x99, 0x94, 0x1f, 0x76, 0x79, 0x28, 0x95, 0x96,
	0x99, 0x53, 0x92, 0x5a, 0xa4, 0x4f, 0x84, 0x1a, 0xa8, 0x7d, 0x4a, 0xd3, 0x18, 0xb9, 0xe4, 0x02,
	0x12, 0x8b, 0x8b, 0x9d, 0xc1, 0xca, 0x03, 0x40, 0xca, 0x83, 0x20, 0xea, 0x5c, 0xa1, 0x56, 0x0a,
	0x25, 0x70, 0xb1, 0x41, 0xf4, 0x48, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x1b, 0x79, 0xe8, 0x11, 0x61,
	0xbe, 0x1e, 0x76, 0x43, 0xa1, 0x94, 0x1b, 0x58, 0x4d, 0x10, 0xd4, 0x5c, 0x21, 0x21, 0x2e, 0x96,
	0x8c, 0xc4, 0xe2, 0x0c, 0x09, 0x26, 0x05, 0x46, 0x0d, 0xce, 0x20, 0x30, 0xdb, 0xc9, 0x27, 0xca,
	0x2b, 0x3d, 0xb3, 0x24, 0xa3, 0x34, 0x49, 0x2f, 0x39, 0x3f, 0x57, 0xdf, 0xc8, 0xb0, 0xb8, 0x24,
	0x33, 0x5f, 0x3f, 0x3d, 0x5f, 0xb7, 0x20, 0x27, 0xb1, 0x32, 0xbd, 0x28, 0xbf, 0x34, 0x2f, 0x45,
	0x3f, 0xbd, 0xa8, 0x20, 0x59, 0xbf, 0x38, 0x39, 0x23, 0x35, 0x37, 0x91, 0x40, 0xe8, 0x02, 0x02,
	0x00, 0x00, 0xff, 0xff, 0x4c, 0x74, 0x3b, 0x6f, 0x9d, 0x01, 0x00, 0x00,
}
