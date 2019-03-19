// Code generated by protoc-gen-go. DO NOT EDIT.
// source: linkfilter/linkfilter_.proto

/*
Package linkfilter is a generated protocol buffer package.

It is generated from these files:
	linkfilter/linkfilter_.proto

It has these top-level messages:
	LinkFilter
*/
package linkfilter

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import urlfilter "github.com/21stio/go-playground/grpc/schema/urlfilter"
import linkkindfilter "github.com/21stio/go-playground/grpc/schema/linkkindfilter"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type LinkFilter struct {
	Url              *urlfilter.UrlFilter           `protobuf:"bytes,1,opt,name=url" json:"url,omitempty"`
	Kind             *linkkindfilter.LinkKindFilter `protobuf:"bytes,2,opt,name=kind" json:"kind,omitempty"`
	And              []*LinkFilter                  `protobuf:"bytes,3,rep,name=and" json:"and,omitempty"`
	Or               []*LinkFilter                  `protobuf:"bytes,4,rep,name=or" json:"or,omitempty"`
	Not              []*LinkFilter                  `protobuf:"bytes,5,rep,name=not" json:"not,omitempty"`
	Hash             *string                        `protobuf:"bytes,6,opt,name=hash" json:"hash,omitempty"`
	XXX_unrecognized []byte                         `json:"-"`
}

func (m *LinkFilter) Reset()                    { *m = LinkFilter{} }
func (m *LinkFilter) String() string            { return proto.CompactTextString(m) }
func (*LinkFilter) ProtoMessage()               {}
func (*LinkFilter) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *LinkFilter) GetUrl() *urlfilter.UrlFilter {
	if m != nil {
		return m.Url
	}
	return nil
}

func (m *LinkFilter) GetKind() *linkkindfilter.LinkKindFilter {
	if m != nil {
		return m.Kind
	}
	return nil
}

func (m *LinkFilter) GetAnd() []*LinkFilter {
	if m != nil {
		return m.And
	}
	return nil
}

func (m *LinkFilter) GetOr() []*LinkFilter {
	if m != nil {
		return m.Or
	}
	return nil
}

func (m *LinkFilter) GetNot() []*LinkFilter {
	if m != nil {
		return m.Not
	}
	return nil
}

func (m *LinkFilter) GetHash() string {
	if m != nil && m.Hash != nil {
		return *m.Hash
	}
	return ""
}

func init() {
	proto.RegisterType((*LinkFilter)(nil), "linkfilter.LinkFilter")
}

func init() { proto.RegisterFile("linkfilter/linkfilter_.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 247 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0xd1, 0xbf, 0x4b, 0xc4, 0x30,
	0x14, 0xc0, 0x71, 0xfa, 0x43, 0xc1, 0x77, 0x5b, 0x10, 0x09, 0x45, 0xa4, 0x88, 0x1c, 0x5d, 0x4c,
	0xb0, 0x83, 0x38, 0x3b, 0xb8, 0xe8, 0x54, 0x70, 0x71, 0x91, 0xda, 0x9e, 0x6d, 0x68, 0x2e, 0xaf,
	0xbc, 0xa6, 0x83, 0xff, 0xb9, 0xa3, 0xa4, 0x86, 0xcb, 0xdd, 0xd2, 0xed, 0xf5, 0xf1, 0xe9, 0x97,
	0x84, 0xc0, 0xb5, 0x56, 0x66, 0xf8, 0x56, 0xda, 0xee, 0x48, 0x86, 0xf1, 0x53, 0x8c, 0x84, 0x16,
	0x19, 0x84, 0x55, 0x96, 0xcd, 0xa4, 0x3d, 0x3c, 0x4c, 0xde, 0x65, 0x77, 0xce, 0x0d, 0xca, 0xb4,
	0x47, 0xa5, 0xf0, 0xe9, 0xd5, 0xed, 0x6f, 0x04, 0xf0, 0xa6, 0xcc, 0xf0, 0xb2, 0x6c, 0xd9, 0x16,
	0x92, 0x99, 0x34, 0x8f, 0xf2, 0xa8, 0xd8, 0x94, 0x97, 0xe2, 0x10, 0x15, 0xef, 0xa4, 0xff, 0x49,
	0xe5, 0x00, 0x2b, 0x21, 0x75, 0x2d, 0x1e, 0x2f, 0xf0, 0x46, 0x9c, 0xc6, 0x85, 0x2b, 0xbe, 0x2a,
	0xd3, 0xfa, 0x5f, 0x16, 0xcb, 0x0a, 0x48, 0x6a, 0xd3, 0xf2, 0x24, 0x4f, 0x8a, 0x4d, 0x79, 0x25,
	0xc2, 0x35, 0x44, 0x38, 0x40, 0xe5, 0x08, 0xdb, 0x42, 0x8c, 0xc4, 0xd3, 0x55, 0x18, 0x23, 0xb9,
	0xa2, 0x41, 0xcb, 0xcf, 0xd6, 0x8b, 0x06, 0x2d, 0x63, 0x90, 0xf6, 0xf5, 0xd4, 0xf3, 0xf3, 0x3c,
	0x2a, 0x2e, 0xaa, 0x65, 0x7e, 0x7e, 0xfa, 0x78, 0xec, 0x94, 0xed, 0xe7, 0x2f, 0xd1, 0xe0, 0x5e,
	0x96, 0x0f, 0x93, 0x55, 0x28, 0x3b, 0xbc, 0x1f, 0x75, 0xfd, 0xd3, 0x11, 0xce, 0xa6, 0x95, 0x1d,
	0x8d, 0x8d, 0x9c, 0x9a, 0x7e, 0xb7, 0xaf, 0x8f, 0x5e, 0xe2, 0x2f, 0x00, 0x00, 0xff, 0xff, 0x03,
	0xc8, 0x9c, 0x43, 0xa1, 0x01, 0x00, 0x00,
}
