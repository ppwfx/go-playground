// Code generated by protoc-gen-go. DO NOT EDIT.
// source: experiencelevelkindfilter/experiencelevelkindfilter_.proto

/*
Package experiencelevelkindfilter is a generated protocol buffer package.

It is generated from these files:
	experiencelevelkindfilter/experiencelevelkindfilter_.proto

It has these top-level messages:
	ExperienceLevelKindFilter
*/
package experiencelevelkindfilter

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import experiencelevelkind "github.com/21stio/go-playground/grpc/schema/experiencelevelkind"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type ExperienceLevelKindFilter struct {
	Is               *experiencelevelkind.ExperienceLevelKind  `protobuf:"varint,1,opt,name=is,enum=experiencelevelkind.ExperienceLevelKind" json:"is,omitempty"`
	Not              *experiencelevelkind.ExperienceLevelKind  `protobuf:"varint,2,opt,name=not,enum=experiencelevelkind.ExperienceLevelKind" json:"not,omitempty"`
	In               []experiencelevelkind.ExperienceLevelKind `protobuf:"varint,3,rep,name=in,enum=experiencelevelkind.ExperienceLevelKind" json:"in,omitempty"`
	NotIn            []experiencelevelkind.ExperienceLevelKind `protobuf:"varint,4,rep,name=notIn,enum=experiencelevelkind.ExperienceLevelKind" json:"notIn,omitempty"`
	Hash             *string                                   `protobuf:"bytes,5,opt,name=hash" json:"hash,omitempty"`
	XXX_unrecognized []byte                                    `json:"-"`
}

func (m *ExperienceLevelKindFilter) Reset()                    { *m = ExperienceLevelKindFilter{} }
func (m *ExperienceLevelKindFilter) String() string            { return proto.CompactTextString(m) }
func (*ExperienceLevelKindFilter) ProtoMessage()               {}
func (*ExperienceLevelKindFilter) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *ExperienceLevelKindFilter) GetIs() experiencelevelkind.ExperienceLevelKind {
	if m != nil && m.Is != nil {
		return *m.Is
	}
	return experiencelevelkind.ExperienceLevelKind_JUNIOR
}

func (m *ExperienceLevelKindFilter) GetNot() experiencelevelkind.ExperienceLevelKind {
	if m != nil && m.Not != nil {
		return *m.Not
	}
	return experiencelevelkind.ExperienceLevelKind_JUNIOR
}

func (m *ExperienceLevelKindFilter) GetIn() []experiencelevelkind.ExperienceLevelKind {
	if m != nil {
		return m.In
	}
	return nil
}

func (m *ExperienceLevelKindFilter) GetNotIn() []experiencelevelkind.ExperienceLevelKind {
	if m != nil {
		return m.NotIn
	}
	return nil
}

func (m *ExperienceLevelKindFilter) GetHash() string {
	if m != nil && m.Hash != nil {
		return *m.Hash
	}
	return ""
}

func init() {
	proto.RegisterType((*ExperienceLevelKindFilter)(nil), "experiencelevelkindfilter.ExperienceLevelKindFilter")
}

func init() {
	proto.RegisterFile("experiencelevelkindfilter/experiencelevelkindfilter_.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 221 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xb2, 0x4a, 0xad, 0x28, 0x48,
	0x2d, 0xca, 0x4c, 0xcd, 0x4b, 0x4e, 0xcd, 0x49, 0x2d, 0x4b, 0xcd, 0xc9, 0xce, 0xcc, 0x4b, 0x49,
	0xcb, 0xcc, 0x29, 0x49, 0x2d, 0xd2, 0xc7, 0x29, 0x13, 0xaf, 0x57, 0x50, 0x94, 0x5f, 0x92, 0x2f,
	0x24, 0x89, 0x53, 0x85, 0x94, 0x1e, 0x16, 0x29, 0x6c, 0x06, 0x42, 0x8d, 0x52, 0x9a, 0xcf, 0xc4,
	0x25, 0xe9, 0x0a, 0x97, 0xf6, 0x01, 0x49, 0x7b, 0x67, 0xe6, 0xa5, 0xb8, 0x81, 0x4d, 0x13, 0xb2,
	0xe0, 0x62, 0xca, 0x2c, 0x96, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x33, 0xd2, 0xc0, 0x66, 0xb4, 0x1e,
	0x16, 0xbd, 0x41, 0x4c, 0x99, 0xc5, 0x42, 0x56, 0x5c, 0xcc, 0x79, 0xf9, 0x25, 0x12, 0x4c, 0x24,
	0x6a, 0x05, 0x69, 0x02, 0xdb, 0x9a, 0x27, 0xc1, 0xac, 0xc0, 0x4c, 0xa2, 0xad, 0x79, 0x42, 0x76,
	0x5c, 0xac, 0x79, 0xf9, 0x25, 0x9e, 0x79, 0x12, 0x2c, 0x24, 0x6a, 0x86, 0x68, 0x13, 0x12, 0xe2,
	0x62, 0xc9, 0x48, 0x2c, 0xce, 0x90, 0x60, 0x55, 0x60, 0xd4, 0xe0, 0x0c, 0x02, 0xb3, 0x9d, 0xdc,
	0xa3, 0x5c, 0xd3, 0x33, 0x4b, 0x32, 0x4a, 0x93, 0xf4, 0x92, 0xf3, 0x73, 0xf5, 0x8d, 0x0c, 0x8b,
	0x4b, 0x32, 0xf3, 0xf5, 0xd3, 0xf3, 0x75, 0x0b, 0x72, 0x12, 0x2b, 0xd3, 0x8b, 0xf2, 0x4b, 0xf3,
	0x52, 0xf4, 0xd3, 0x8b, 0x0a, 0x92, 0xf5, 0x8b, 0x93, 0x33, 0x52, 0x73, 0x13, 0x71, 0x47, 0x1e,
	0x20, 0x00, 0x00, 0xff, 0xff, 0xe8, 0xd6, 0x67, 0x44, 0xf2, 0x01, 0x00, 0x00,
}
