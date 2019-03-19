// Code generated by protoc-gen-go. DO NOT EDIT.
// source: getwhateverlistbyidrequestfilter/getwhateverlistbyidrequestfilter_.proto

/*
Package getwhateverlistbyidrequestfilter is a generated protocol buffer package.

It is generated from these files:
	getwhateverlistbyidrequestfilter/getwhateverlistbyidrequestfilter_.proto

It has these top-level messages:
	GetWhateverListByIdRequestFilter
*/
package getwhateverlistbyidrequestfilter

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import requestmetafilter "github.com/21stio/go-playground/grpc/schema/requestmetafilter"
import idfilter "github.com/21stio/go-playground/grpc/schema/idfilter"
import whateverskindfilter "github.com/21stio/go-playground/grpc/schema/whateverskindfilter"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type GetWhateverListByIdRequestFilter struct {
	Meta             *requestmetafilter.RequestMetaFilter     `protobuf:"bytes,1,opt,name=meta" json:"meta,omitempty"`
	Id               *idfilter.IdFilter                       `protobuf:"bytes,2,opt,name=id" json:"id,omitempty"`
	Kind             *whateverskindfilter.WhateversKindFilter `protobuf:"bytes,3,opt,name=kind" json:"kind,omitempty"`
	And              []*GetWhateverListByIdRequestFilter      `protobuf:"bytes,4,rep,name=and" json:"and,omitempty"`
	Or               []*GetWhateverListByIdRequestFilter      `protobuf:"bytes,5,rep,name=or" json:"or,omitempty"`
	Not              []*GetWhateverListByIdRequestFilter      `protobuf:"bytes,6,rep,name=not" json:"not,omitempty"`
	Hash             *string                                  `protobuf:"bytes,7,opt,name=hash" json:"hash,omitempty"`
	XXX_unrecognized []byte                                   `json:"-"`
}

func (m *GetWhateverListByIdRequestFilter) Reset()         { *m = GetWhateverListByIdRequestFilter{} }
func (m *GetWhateverListByIdRequestFilter) String() string { return proto.CompactTextString(m) }
func (*GetWhateverListByIdRequestFilter) ProtoMessage()    {}
func (*GetWhateverListByIdRequestFilter) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{0}
}

func (m *GetWhateverListByIdRequestFilter) GetMeta() *requestmetafilter.RequestMetaFilter {
	if m != nil {
		return m.Meta
	}
	return nil
}

func (m *GetWhateverListByIdRequestFilter) GetId() *idfilter.IdFilter {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *GetWhateverListByIdRequestFilter) GetKind() *whateverskindfilter.WhateversKindFilter {
	if m != nil {
		return m.Kind
	}
	return nil
}

func (m *GetWhateverListByIdRequestFilter) GetAnd() []*GetWhateverListByIdRequestFilter {
	if m != nil {
		return m.And
	}
	return nil
}

func (m *GetWhateverListByIdRequestFilter) GetOr() []*GetWhateverListByIdRequestFilter {
	if m != nil {
		return m.Or
	}
	return nil
}

func (m *GetWhateverListByIdRequestFilter) GetNot() []*GetWhateverListByIdRequestFilter {
	if m != nil {
		return m.Not
	}
	return nil
}

func (m *GetWhateverListByIdRequestFilter) GetHash() string {
	if m != nil && m.Hash != nil {
		return *m.Hash
	}
	return ""
}

func init() {
	proto.RegisterType((*GetWhateverListByIdRequestFilter)(nil), "getwhateverlistbyidrequestfilter.GetWhateverListByIdRequestFilter")
}

func init() {
	proto.RegisterFile("getwhateverlistbyidrequestfilter/getwhateverlistbyidrequestfilter_.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 312 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x92, 0x31, 0x4f, 0xfb, 0x30,
	0x10, 0xc5, 0xd5, 0x24, 0xff, 0x3f, 0xc2, 0xdd, 0x3c, 0x59, 0x9d, 0xa2, 0x8a, 0x21, 0x42, 0xc2,
	0x11, 0x9d, 0x18, 0x98, 0x3a, 0x00, 0x15, 0x85, 0x21, 0x42, 0x42, 0x62, 0x41, 0x6e, 0x6c, 0x12,
	0x8b, 0xc4, 0x0e, 0xf6, 0x05, 0x94, 0xaf, 0xc9, 0x27, 0x42, 0x71, 0xed, 0xa9, 0x95, 0xb2, 0x74,
	0x3b, 0x9f, 0x7f, 0xef, 0xdd, 0x3b, 0xcb, 0xe8, 0xa1, 0x12, 0xf0, 0x53, 0x33, 0x10, 0xdf, 0xc2,
	0x34, 0xd2, 0xc2, 0x6e, 0x90, 0xdc, 0x88, 0xaf, 0x5e, 0x58, 0xf8, 0x90, 0x0d, 0x08, 0x93, 0x4f,
	0x01, 0xef, 0xb4, 0x33, 0x1a, 0x34, 0x4e, 0xa7, 0xc0, 0xc5, 0xa5, 0x3f, 0xb6, 0x02, 0x98, 0x37,
	0x3f, 0xe8, 0x78, 0xb7, 0x05, 0x91, 0xdc, 0x23, 0xa1, 0x08, 0x37, 0x34, 0x0c, 0xb1, 0x9f, 0x52,
	0x05, 0xe8, 0x48, 0xcf, 0xf3, 0xcb, 0xdf, 0x18, 0xa5, 0xf7, 0x02, 0x5e, 0x3d, 0xb1, 0x95, 0x16,
	0xd6, 0xc3, 0x86, 0x17, 0xfb, 0xc9, 0x77, 0x8e, 0xc5, 0x37, 0x28, 0x19, 0x33, 0x90, 0x59, 0x3a,
	0xcb, 0xe6, 0xab, 0x0b, 0x7a, 0x90, 0x8b, 0x7a, 0xfe, 0x49, 0x00, 0xdb, 0x6b, 0x0a, 0xa7, 0xc0,
	0x4b, 0x14, 0x49, 0x4e, 0x22, 0xa7, 0xc3, 0x34, 0x84, 0xa5, 0x1b, 0xee, 0xa9, 0x48, 0x72, 0x7c,
	0x8b, 0x92, 0x31, 0x17, 0x89, 0x1d, 0x95, 0x1d, 0xdb, 0x80, 0x86, 0x7c, 0xf6, 0x51, 0xaa, 0xa0,
	0x75, 0x2a, 0xfc, 0x82, 0x62, 0xa6, 0x38, 0x49, 0xd2, 0x38, 0x9b, 0xaf, 0xd6, 0x74, 0xea, 0x99,
	0xe9, 0xd4, 0xb2, 0xc5, 0x68, 0x87, 0x0b, 0x14, 0x69, 0x43, 0xfe, 0x9d, 0xcc, 0x34, 0xd2, 0x66,
	0x4c, 0xaa, 0x34, 0x90, 0xff, 0xa7, 0x4b, 0xaa, 0x34, 0x60, 0x8c, 0x92, 0x9a, 0xd9, 0x9a, 0x9c,
	0xa5, 0xb3, 0xec, 0xbc, 0x70, 0xf5, 0xfa, 0xf9, 0x6d, 0x5b, 0x49, 0xa8, 0xfb, 0x1d, 0x2d, 0x75,
	0x9b, 0xaf, 0xae, 0x2d, 0x48, 0x9d, 0x57, 0xfa, 0xaa, 0x6b, 0xd8, 0x50, 0x19, 0xdd, 0x2b, 0x9e,
	0x57, 0xa6, 0x2b, 0x73, 0x5b, 0xd6, 0xa2, 0x65, 0x93, 0x7f, 0xf8, 0x2f, 0x00, 0x00, 0xff, 0xff,
	0x95, 0xc7, 0x90, 0xd2, 0x07, 0x03, 0x00, 0x00,
}
