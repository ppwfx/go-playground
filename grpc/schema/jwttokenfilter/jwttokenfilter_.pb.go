// Code generated by protoc-gen-go. DO NOT EDIT.
// source: jwttokenfilter/jwttokenfilter_.proto

/*
Package jwttokenfilter is a generated protocol buffer package.

It is generated from these files:
	jwttokenfilter/jwttokenfilter_.proto

It has these top-level messages:
	JwtTokenFilter
*/
package jwttokenfilter

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

type JwtTokenFilter struct {
	Value            *stringfilter.StringFilter `protobuf:"bytes,1,opt,name=value" json:"value,omitempty"`
	And              []*JwtTokenFilter          `protobuf:"bytes,2,rep,name=and" json:"and,omitempty"`
	Or               []*JwtTokenFilter          `protobuf:"bytes,3,rep,name=or" json:"or,omitempty"`
	Not              []*JwtTokenFilter          `protobuf:"bytes,4,rep,name=not" json:"not,omitempty"`
	Hash             *string                    `protobuf:"bytes,5,opt,name=hash" json:"hash,omitempty"`
	XXX_unrecognized []byte                     `json:"-"`
}

func (m *JwtTokenFilter) Reset()                    { *m = JwtTokenFilter{} }
func (m *JwtTokenFilter) String() string            { return proto.CompactTextString(m) }
func (*JwtTokenFilter) ProtoMessage()               {}
func (*JwtTokenFilter) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *JwtTokenFilter) GetValue() *stringfilter.StringFilter {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *JwtTokenFilter) GetAnd() []*JwtTokenFilter {
	if m != nil {
		return m.And
	}
	return nil
}

func (m *JwtTokenFilter) GetOr() []*JwtTokenFilter {
	if m != nil {
		return m.Or
	}
	return nil
}

func (m *JwtTokenFilter) GetNot() []*JwtTokenFilter {
	if m != nil {
		return m.Not
	}
	return nil
}

func (m *JwtTokenFilter) GetHash() string {
	if m != nil && m.Hash != nil {
		return *m.Hash
	}
	return ""
}

func init() {
	proto.RegisterType((*JwtTokenFilter)(nil), "jwttokenfilter.JwtTokenFilter")
}

func init() { proto.RegisterFile("jwttokenfilter/jwttokenfilter_.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 228 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0xc9, 0x2a, 0x2f, 0x29,
	0xc9, 0xcf, 0x4e, 0xcd, 0x4b, 0xcb, 0xcc, 0x29, 0x49, 0x2d, 0xd2, 0x47, 0xe5, 0xc6, 0xeb, 0x15,
	0x14, 0xe5, 0x97, 0xe4, 0x0b, 0xf1, 0xa1, 0x0a, 0x4b, 0x29, 0x14, 0x97, 0x14, 0x65, 0xe6, 0xa5,
	0x43, 0xf5, 0x20, 0x73, 0xa0, 0x3a, 0x94, 0x5e, 0x31, 0x72, 0xf1, 0x79, 0x95, 0x97, 0x84, 0x80,
	0x34, 0xb9, 0x81, 0x65, 0x84, 0x0c, 0xb8, 0x58, 0xcb, 0x12, 0x73, 0x4a, 0x53, 0x25, 0x18, 0x15,
	0x18, 0x35, 0xb8, 0x8d, 0xa4, 0xf4, 0x90, 0xf5, 0xe9, 0x05, 0x83, 0x39, 0x10, 0xa5, 0x41, 0x10,
	0x85, 0x42, 0x06, 0x5c, 0xcc, 0x89, 0x79, 0x29, 0x12, 0x4c, 0x0a, 0xcc, 0x1a, 0xdc, 0x46, 0x72,
	0x7a, 0xa8, 0x8e, 0xd0, 0x43, 0x35, 0x3e, 0x08, 0xa4, 0x54, 0x48, 0x8f, 0x8b, 0x29, 0xbf, 0x48,
	0x82, 0x99, 0x28, 0x0d, 0x4c, 0xf9, 0x20, 0x37, 0x31, 0xe7, 0xe5, 0x97, 0x48, 0xb0, 0x10, 0x67,
	0x43, 0x5e, 0x7e, 0x89, 0x90, 0x10, 0x17, 0x4b, 0x46, 0x62, 0x71, 0x86, 0x04, 0xab, 0x02, 0xa3,
	0x06, 0x67, 0x10, 0x98, 0xed, 0x64, 0x13, 0x65, 0x95, 0x9e, 0x59, 0x92, 0x51, 0x9a, 0xa4, 0x97,
	0x9c, 0x9f, 0xab, 0x6f, 0x64, 0x58, 0x5c, 0x92, 0x99, 0xaf, 0x9f, 0x9e, 0xaf, 0x5b, 0x90, 0x93,
	0x58, 0x99, 0x5e, 0x94, 0x5f, 0x9a, 0x97, 0xa2, 0x9f, 0x5e, 0x54, 0x90, 0xac, 0x5f, 0x9c, 0x9c,
	0x91, 0x9a, 0x9b, 0x88, 0x16, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0xb1, 0x8d, 0xfe, 0xd4,
	0x83, 0x01, 0x00, 0x00,
}
