// Code generated by protoc-gen-go. DO NOT EDIT.
// source: streamwhateversendpoint/streamwhateversendpoint_.proto

/*
Package streamwhateversendpoint is a generated protocol buffer package.

It is generated from these files:
	streamwhateversendpoint/streamwhateversendpoint_.proto

It has these top-level messages:
	StreamWhateversEndpoint
*/
package streamwhateversendpoint

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import streamwhateversrequestfilter "github.com/21stio/go-playground/grpc/schema/streamwhateversrequestfilter"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type StreamWhateversEndpoint struct {
	Filter           *streamwhateversrequestfilter.StreamWhateversRequestFilter `protobuf:"bytes,1,opt,name=filter" json:"filter,omitempty"`
	Hash             *string                                                    `protobuf:"bytes,2,opt,name=hash" json:"hash,omitempty"`
	XXX_unrecognized []byte                                                     `json:"-"`
}

func (m *StreamWhateversEndpoint) Reset()                    { *m = StreamWhateversEndpoint{} }
func (m *StreamWhateversEndpoint) String() string            { return proto.CompactTextString(m) }
func (*StreamWhateversEndpoint) ProtoMessage()               {}
func (*StreamWhateversEndpoint) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *StreamWhateversEndpoint) GetFilter() *streamwhateversrequestfilter.StreamWhateversRequestFilter {
	if m != nil {
		return m.Filter
	}
	return nil
}

func (m *StreamWhateversEndpoint) GetHash() string {
	if m != nil && m.Hash != nil {
		return *m.Hash
	}
	return ""
}

func init() {
	proto.RegisterType((*StreamWhateversEndpoint)(nil), "streamwhateversendpoint.StreamWhateversEndpoint")
}

func init() {
	proto.RegisterFile("streamwhateversendpoint/streamwhateversendpoint_.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 196 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x32, 0x2b, 0x2e, 0x29, 0x4a,
	0x4d, 0xcc, 0x2d, 0xcf, 0x48, 0x2c, 0x49, 0x2d, 0x4b, 0x2d, 0x2a, 0x4e, 0xcd, 0x4b, 0x29, 0xc8,
	0xcf, 0xcc, 0x2b, 0xd1, 0xc7, 0x21, 0x1e, 0xaf, 0x57, 0x50, 0x94, 0x5f, 0x92, 0x2f, 0x24, 0x8e,
	0x43, 0x5e, 0xca, 0x01, 0x4d, 0xa2, 0x28, 0xb5, 0xb0, 0x34, 0xb5, 0xb8, 0x24, 0x2d, 0x33, 0xa7,
	0x24, 0xb5, 0x48, 0x1f, 0x9f, 0x24, 0xd4, 0x68, 0xa5, 0x46, 0x46, 0x2e, 0xf1, 0x60, 0xb0, 0xba,
	0x70, 0x98, 0x3a, 0x57, 0xa8, 0xe9, 0x42, 0x41, 0x5c, 0x6c, 0x10, 0xc5, 0x12, 0x8c, 0x0a, 0x8c,
	0x1a, 0xdc, 0x46, 0x56, 0x7a, 0xf8, 0x4c, 0xd4, 0x43, 0x33, 0x26, 0x08, 0x22, 0xe9, 0x06, 0x96,
	0x0c, 0x82, 0x9a, 0x24, 0x24, 0xc4, 0xc5, 0x92, 0x91, 0x58, 0x9c, 0x21, 0xc1, 0xa4, 0xc0, 0xa8,
	0xc1, 0x19, 0x04, 0x66, 0x3b, 0xb9, 0x46, 0x39, 0xa7, 0x67, 0x96, 0x64, 0x94, 0x26, 0xe9, 0x25,
	0xe7, 0xe7, 0xea, 0x1b, 0x19, 0x16, 0x97, 0x64, 0xe6, 0xeb, 0xa7, 0xe7, 0xeb, 0x16, 0xe4, 0x24,
	0x56, 0xa6, 0x17, 0xe5, 0x97, 0xe6, 0xa5, 0xe8, 0xa7, 0x17, 0x15, 0x24, 0xeb, 0x17, 0x27, 0x67,
	0xa4, 0xe6, 0x26, 0xe2, 0x0a, 0x2c, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x31, 0x54, 0xb0, 0x7d,
	0x5e, 0x01, 0x00, 0x00,
}
