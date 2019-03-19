// Code generated by protoc-gen-go. DO NOT EDIT.
// source: stringfilter/stringfilter_.proto

/*
Package stringfilter is a generated protocol buffer package.

It is generated from these files:
	stringfilter/stringfilter_.proto

It has these top-level messages:
	StringFilter
*/
package stringfilter

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type StringFilter struct {
	CaseSensitive    *bool    `protobuf:"varint,1,opt,name=caseSensitive" json:"caseSensitive,omitempty"`
	Set              *bool    `protobuf:"varint,2,opt,name=set" json:"set,omitempty"`
	Is               *string  `protobuf:"bytes,3,opt,name=is" json:"is,omitempty"`
	Not              *string  `protobuf:"bytes,4,opt,name=not" json:"not,omitempty"`
	Contains         *string  `protobuf:"bytes,5,opt,name=contains" json:"contains,omitempty"`
	NotContains      *string  `protobuf:"bytes,6,opt,name=notContains" json:"notContains,omitempty"`
	StartsWith       *string  `protobuf:"bytes,7,opt,name=startsWith" json:"startsWith,omitempty"`
	NotStartsWith    *string  `protobuf:"bytes,8,opt,name=notStartsWith" json:"notStartsWith,omitempty"`
	EndsWith         *string  `protobuf:"bytes,9,opt,name=endsWith" json:"endsWith,omitempty"`
	NotEndsWith      *string  `protobuf:"bytes,10,opt,name=notEndsWith" json:"notEndsWith,omitempty"`
	In               []string `protobuf:"bytes,11,rep,name=in" json:"in,omitempty"`
	NotIn            []string `protobuf:"bytes,12,rep,name=notIn" json:"notIn,omitempty"`
	Hash             *string  `protobuf:"bytes,13,opt,name=hash" json:"hash,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *StringFilter) Reset()                    { *m = StringFilter{} }
func (m *StringFilter) String() string            { return proto.CompactTextString(m) }
func (*StringFilter) ProtoMessage()               {}
func (*StringFilter) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *StringFilter) GetCaseSensitive() bool {
	if m != nil && m.CaseSensitive != nil {
		return *m.CaseSensitive
	}
	return false
}

func (m *StringFilter) GetSet() bool {
	if m != nil && m.Set != nil {
		return *m.Set
	}
	return false
}

func (m *StringFilter) GetIs() string {
	if m != nil && m.Is != nil {
		return *m.Is
	}
	return ""
}

func (m *StringFilter) GetNot() string {
	if m != nil && m.Not != nil {
		return *m.Not
	}
	return ""
}

func (m *StringFilter) GetContains() string {
	if m != nil && m.Contains != nil {
		return *m.Contains
	}
	return ""
}

func (m *StringFilter) GetNotContains() string {
	if m != nil && m.NotContains != nil {
		return *m.NotContains
	}
	return ""
}

func (m *StringFilter) GetStartsWith() string {
	if m != nil && m.StartsWith != nil {
		return *m.StartsWith
	}
	return ""
}

func (m *StringFilter) GetNotStartsWith() string {
	if m != nil && m.NotStartsWith != nil {
		return *m.NotStartsWith
	}
	return ""
}

func (m *StringFilter) GetEndsWith() string {
	if m != nil && m.EndsWith != nil {
		return *m.EndsWith
	}
	return ""
}

func (m *StringFilter) GetNotEndsWith() string {
	if m != nil && m.NotEndsWith != nil {
		return *m.NotEndsWith
	}
	return ""
}

func (m *StringFilter) GetIn() []string {
	if m != nil {
		return m.In
	}
	return nil
}

func (m *StringFilter) GetNotIn() []string {
	if m != nil {
		return m.NotIn
	}
	return nil
}

func (m *StringFilter) GetHash() string {
	if m != nil && m.Hash != nil {
		return *m.Hash
	}
	return ""
}

func init() {
	proto.RegisterType((*StringFilter)(nil), "stringfilter.StringFilter")
}

func init() { proto.RegisterFile("stringfilter/stringfilter_.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 282 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x91, 0x41, 0x4b, 0xc3, 0x40,
	0x10, 0x85, 0x69, 0xda, 0x6a, 0x3a, 0x4d, 0x45, 0x16, 0x0f, 0x8b, 0x07, 0x09, 0xe2, 0xa1, 0x17,
	0x1b, 0xf4, 0x24, 0x1e, 0x15, 0x05, 0xaf, 0xcd, 0x41, 0xf0, 0x22, 0x6b, 0xba, 0x26, 0x0b, 0xed,
	0x6c, 0xd8, 0x99, 0x0a, 0xfe, 0x5f, 0x7f, 0x88, 0x64, 0x62, 0x6a, 0x72, 0x7b, 0xef, 0x7b, 0x8f,
	0xe5, 0x0d, 0x0b, 0x29, 0x71, 0x70, 0x58, 0x7e, 0xba, 0x2d, 0xdb, 0x90, 0xf5, 0xcd, 0xfb, 0xaa,
	0x0e, 0x9e, 0xbd, 0x4a, 0xfa, 0xf0, 0xf2, 0x27, 0x82, 0x24, 0x17, 0xf0, 0x2c, 0x40, 0x5d, 0xc1,
	0xa2, 0x30, 0x64, 0x73, 0x8b, 0xe4, 0xd8, 0x7d, 0x59, 0x3d, 0x4a, 0x47, 0xcb, 0x78, 0x3d, 0x84,
	0xea, 0x14, 0xc6, 0x64, 0x59, 0x47, 0x92, 0x35, 0x52, 0x9d, 0x40, 0xe4, 0x48, 0x8f, 0xd3, 0xd1,
	0x72, 0xb6, 0x8e, 0x1c, 0x35, 0x0d, 0xf4, 0xac, 0x27, 0x02, 0x1a, 0xa9, 0xce, 0x21, 0x2e, 0x3c,
	0xb2, 0x71, 0x48, 0x7a, 0x2a, 0xf8, 0xe0, 0x55, 0x0a, 0x73, 0xf4, 0xfc, 0xd8, 0xc5, 0x47, 0x12,
	0xf7, 0x91, 0xba, 0x00, 0x20, 0x36, 0x81, 0xe9, 0xd5, 0x71, 0xa5, 0x8f, 0xa5, 0xd0, 0x23, 0xcd,
	0x6e, 0xf4, 0x9c, 0xff, 0x57, 0x62, 0xa9, 0x0c, 0x61, 0xb3, 0xc1, 0xe2, 0xa6, 0x2d, 0xcc, 0xda,
	0x0d, 0x9d, 0xff, 0xdb, 0xf0, 0xd4, 0xc5, 0x70, 0xd8, 0xd0, 0x21, 0xb9, 0x11, 0xf5, 0x3c, 0x1d,
	0xcb, 0x8d, 0xa8, 0xce, 0x60, 0x8a, 0x9e, 0x5f, 0x50, 0x27, 0x82, 0x5a, 0xa3, 0x14, 0x4c, 0x2a,
	0x43, 0x95, 0x5e, 0xc8, 0x03, 0xa2, 0x1f, 0xee, 0xdf, 0xee, 0x4a, 0xc7, 0xd5, 0xfe, 0x63, 0x55,
	0xf8, 0x5d, 0x76, 0x7b, 0x43, 0xec, 0x7c, 0x56, 0xfa, 0xeb, 0x7a, 0x6b, 0xbe, 0xcb, 0xe0, 0xf7,
	0xb8, 0xc9, 0xca, 0x50, 0x17, 0x19, 0x15, 0x95, 0xdd, 0x99, 0xc1, 0xbf, 0xfd, 0x06, 0x00, 0x00,
	0xff, 0xff, 0xd2, 0x7d, 0x29, 0x9d, 0xd3, 0x01, 0x00, 0x00,
}