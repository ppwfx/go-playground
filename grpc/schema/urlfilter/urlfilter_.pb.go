// Code generated by protoc-gen-go. DO NOT EDIT.
// source: urlfilter/urlfilter_.proto

/*
Package urlfilter is a generated protocol buffer package.

It is generated from these files:
	urlfilter/urlfilter_.proto

It has these top-level messages:
	UrlFilter
*/
package urlfilter

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import stringfilter "github.com/21stio/go-playground/grpc/schema/stringfilter"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type UrlFilter struct {
	Value            *stringfilter.StringFilter `protobuf:"bytes,1,opt,name=value" json:"value,omitempty"`
	And              []*UrlFilter               `protobuf:"bytes,2,rep,name=and" json:"and,omitempty"`
	Or               []*UrlFilter               `protobuf:"bytes,3,rep,name=or" json:"or,omitempty"`
	Not              []*UrlFilter               `protobuf:"bytes,4,rep,name=not" json:"not,omitempty"`
	Hash             *string                    `protobuf:"bytes,5,opt,name=hash" json:"hash,omitempty"`
	XXX_unrecognized []byte                     `json:"-"`
}

func (m *UrlFilter) Reset()                    { *m = UrlFilter{} }
func (m *UrlFilter) String() string            { return proto.CompactTextString(m) }
func (*UrlFilter) ProtoMessage()               {}
func (*UrlFilter) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *UrlFilter) GetValue() *stringfilter.StringFilter {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *UrlFilter) GetAnd() []*UrlFilter {
	if m != nil {
		return m.And
	}
	return nil
}

func (m *UrlFilter) GetOr() []*UrlFilter {
	if m != nil {
		return m.Or
	}
	return nil
}

func (m *UrlFilter) GetNot() []*UrlFilter {
	if m != nil {
		return m.Not
	}
	return nil
}

func (m *UrlFilter) GetHash() string {
	if m != nil && m.Hash != nil {
		return *m.Hash
	}
	return ""
}

func init() {
	proto.RegisterType((*UrlFilter)(nil), "urlfilter.UrlFilter")
}

func init() { proto.RegisterFile("urlfilter/urlfilter_.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 218 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0x31, 0x4b, 0xc4, 0x30,
	0x18, 0x86, 0x49, 0x7b, 0x37, 0x34, 0xb7, 0x05, 0x87, 0xd0, 0x29, 0x88, 0x48, 0x17, 0x13, 0x2d,
	0x88, 0xbb, 0x83, 0x3f, 0xa0, 0xe2, 0xe2, 0x22, 0xb1, 0x57, 0x93, 0x40, 0x2e, 0x5f, 0xf9, 0x92,
	0x08, 0xfe, 0x42, 0xff, 0x96, 0x5c, 0xef, 0x88, 0x37, 0x75, 0x7b, 0xdf, 0x97, 0x27, 0x4f, 0xe0,
	0xa3, 0x6d, 0x46, 0xff, 0xe5, 0x7c, 0x9a, 0x50, 0x95, 0xf4, 0x21, 0x67, 0x84, 0x04, 0xac, 0x29,
	0x4b, 0x2b, 0x62, 0x42, 0x17, 0xcc, 0x99, 0xbc, 0x2c, 0x67, 0xf8, 0xfa, 0x97, 0xd0, 0xe6, 0x0d,
	0xfd, 0xcb, 0x32, 0xb2, 0x7b, 0xba, 0xfd, 0xd6, 0x3e, 0x4f, 0x9c, 0x08, 0xd2, 0xed, 0xfa, 0x56,
	0x5e, 0x3e, 0x91, 0xaf, 0x4b, 0x39, 0xa1, 0xc3, 0x09, 0x64, 0xb7, 0xb4, 0xd6, 0x61, 0xcf, 0x2b,
	0x51, 0x77, 0xbb, 0xfe, 0x4a, 0x96, 0xaf, 0x65, 0x91, 0x0e, 0x47, 0x80, 0xdd, 0xd0, 0x0a, 0x90,
	0xd7, 0x2b, 0x58, 0x05, 0x78, 0xb4, 0x05, 0x48, 0x7c, 0xb3, 0x66, 0x0b, 0x90, 0x18, 0xa3, 0x1b,
	0xab, 0xa3, 0xe5, 0x5b, 0x41, 0xba, 0x66, 0x58, 0xf2, 0xf3, 0xd3, 0xfb, 0xa3, 0x71, 0xc9, 0xe6,
	0x4f, 0x39, 0xc2, 0x41, 0xf5, 0x0f, 0x31, 0x39, 0x50, 0x06, 0xee, 0x66, 0xaf, 0x7f, 0x0c, 0x42,
	0x0e, 0x7b, 0x65, 0x70, 0x1e, 0x55, 0x1c, 0xed, 0x74, 0xd0, 0xff, 0x67, 0xfb, 0x0b, 0x00, 0x00,
	0xff, 0xff, 0x6f, 0x8c, 0xe0, 0x75, 0x4c, 0x01, 0x00, 0x00,
}