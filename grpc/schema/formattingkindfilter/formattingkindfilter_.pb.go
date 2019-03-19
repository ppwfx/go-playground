// Code generated by protoc-gen-go. DO NOT EDIT.
// source: formattingkindfilter/formattingkindfilter_.proto

/*
Package formattingkindfilter is a generated protocol buffer package.

It is generated from these files:
	formattingkindfilter/formattingkindfilter_.proto

It has these top-level messages:
	FormattingKindFilter
*/
package formattingkindfilter

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import formattingkind "github.com/21stio/go-playground/grpc/schema/formattingkind"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type FormattingKindFilter struct {
	Is               *formattingkind.FormattingKind  `protobuf:"varint,1,opt,name=is,enum=formattingkind.FormattingKind" json:"is,omitempty"`
	Not              *formattingkind.FormattingKind  `protobuf:"varint,2,opt,name=not,enum=formattingkind.FormattingKind" json:"not,omitempty"`
	In               []formattingkind.FormattingKind `protobuf:"varint,3,rep,name=in,enum=formattingkind.FormattingKind" json:"in,omitempty"`
	NotIn            []formattingkind.FormattingKind `protobuf:"varint,4,rep,name=notIn,enum=formattingkind.FormattingKind" json:"notIn,omitempty"`
	Hash             *string                         `protobuf:"bytes,5,opt,name=hash" json:"hash,omitempty"`
	XXX_unrecognized []byte                          `json:"-"`
}

func (m *FormattingKindFilter) Reset()                    { *m = FormattingKindFilter{} }
func (m *FormattingKindFilter) String() string            { return proto.CompactTextString(m) }
func (*FormattingKindFilter) ProtoMessage()               {}
func (*FormattingKindFilter) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *FormattingKindFilter) GetIs() formattingkind.FormattingKind {
	if m != nil && m.Is != nil {
		return *m.Is
	}
	return formattingkind.FormattingKind_HTML
}

func (m *FormattingKindFilter) GetNot() formattingkind.FormattingKind {
	if m != nil && m.Not != nil {
		return *m.Not
	}
	return formattingkind.FormattingKind_HTML
}

func (m *FormattingKindFilter) GetIn() []formattingkind.FormattingKind {
	if m != nil {
		return m.In
	}
	return nil
}

func (m *FormattingKindFilter) GetNotIn() []formattingkind.FormattingKind {
	if m != nil {
		return m.NotIn
	}
	return nil
}

func (m *FormattingKindFilter) GetHash() string {
	if m != nil && m.Hash != nil {
		return *m.Hash
	}
	return ""
}

func init() {
	proto.RegisterType((*FormattingKindFilter)(nil), "formattingkindfilter.FormattingKindFilter")
}

func init() { proto.RegisterFile("formattingkindfilter/formattingkindfilter_.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 214 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x32, 0x48, 0xcb, 0x2f, 0xca,
	0x4d, 0x2c, 0x29, 0xc9, 0xcc, 0x4b, 0xcf, 0xce, 0xcc, 0x4b, 0x49, 0xcb, 0xcc, 0x29, 0x49, 0x2d,
	0xd2, 0xc7, 0x26, 0x18, 0xaf, 0x57, 0x50, 0x94, 0x5f, 0x92, 0x2f, 0x24, 0x82, 0x4d, 0x52, 0x4a,
	0x05, 0x55, 0x14, 0xcd, 0x04, 0xa8, 0x5e, 0xa5, 0x4f, 0x8c, 0x5c, 0x22, 0x6e, 0x70, 0x19, 0xef,
	0xcc, 0xbc, 0x14, 0x37, 0xb0, 0x76, 0x21, 0x3d, 0x2e, 0xa6, 0xcc, 0x62, 0x09, 0x46, 0x05, 0x46,
	0x0d, 0x3e, 0x23, 0x39, 0x3d, 0x54, 0xcd, 0x7a, 0xa8, 0x3a, 0x82, 0x98, 0x32, 0x8b, 0x85, 0x0c,
	0xb8, 0x98, 0xf3, 0xf2, 0x4b, 0x24, 0x98, 0x88, 0xd2, 0x00, 0x52, 0x0a, 0xb6, 0x21, 0x4f, 0x82,
	0x59, 0x81, 0x99, 0x28, 0x1b, 0xf2, 0x84, 0x4c, 0xb8, 0x58, 0xf3, 0xf2, 0x4b, 0x3c, 0xf3, 0x24,
	0x58, 0x88, 0xd2, 0x02, 0x51, 0x2c, 0x24, 0xc4, 0xc5, 0x92, 0x91, 0x58, 0x9c, 0x21, 0xc1, 0xaa,
	0xc0, 0xa8, 0xc1, 0x19, 0x04, 0x66, 0x3b, 0x39, 0x45, 0x39, 0xa4, 0x67, 0x96, 0x64, 0x94, 0x26,
	0xe9, 0x25, 0xe7, 0xe7, 0xea, 0x1b, 0x19, 0x16, 0x97, 0x64, 0xe6, 0xeb, 0xa7, 0xe7, 0xeb, 0x16,
	0xe4, 0x24, 0x56, 0xa6, 0x17, 0xe5, 0x97, 0xe6, 0xa5, 0xe8, 0xa7, 0x17, 0x15, 0x24, 0xeb, 0x17,
	0x27, 0x67, 0xa4, 0xe6, 0x26, 0x62, 0x0d, 0x7b, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x68, 0x5e,
	0xc8, 0xd6, 0xa7, 0x01, 0x00, 0x00,
}
