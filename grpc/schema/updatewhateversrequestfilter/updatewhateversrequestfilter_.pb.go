// Code generated by protoc-gen-go. DO NOT EDIT.
// source: updatewhateversrequestfilter/updatewhateversrequestfilter_.proto

/*
Package updatewhateversrequestfilter is a generated protocol buffer package.

It is generated from these files:
	updatewhateversrequestfilter/updatewhateversrequestfilter_.proto

It has these top-level messages:
	UpdateWhateversRequestFilter
*/
package updatewhateversrequestfilter

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import requestmetafilter "github.com/21stio/go-playground/grpc/schema/requestmetafilter"
import whateverfilter "github.com/21stio/go-playground/grpc/schema/whateverfilter"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type UpdateWhateversRequestFilter struct {
	Meta             *requestmetafilter.RequestMetaFilter `protobuf:"bytes,1,opt,name=meta" json:"meta,omitempty"`
	WhateversSome    *whateverfilter.WhateverFilter       `protobuf:"bytes,2,opt,name=whateversSome" json:"whateversSome,omitempty"`
	WhateversEvery   *whateverfilter.WhateverFilter       `protobuf:"bytes,3,opt,name=whateversEvery" json:"whateversEvery,omitempty"`
	WhateversNone    *whateverfilter.WhateverFilter       `protobuf:"bytes,4,opt,name=whateversNone" json:"whateversNone,omitempty"`
	And              []*UpdateWhateversRequestFilter      `protobuf:"bytes,5,rep,name=and" json:"and,omitempty"`
	Or               []*UpdateWhateversRequestFilter      `protobuf:"bytes,6,rep,name=or" json:"or,omitempty"`
	Not              []*UpdateWhateversRequestFilter      `protobuf:"bytes,7,rep,name=not" json:"not,omitempty"`
	Hash             *string                              `protobuf:"bytes,8,opt,name=hash" json:"hash,omitempty"`
	XXX_unrecognized []byte                               `json:"-"`
}

func (m *UpdateWhateversRequestFilter) Reset()                    { *m = UpdateWhateversRequestFilter{} }
func (m *UpdateWhateversRequestFilter) String() string            { return proto.CompactTextString(m) }
func (*UpdateWhateversRequestFilter) ProtoMessage()               {}
func (*UpdateWhateversRequestFilter) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *UpdateWhateversRequestFilter) GetMeta() *requestmetafilter.RequestMetaFilter {
	if m != nil {
		return m.Meta
	}
	return nil
}

func (m *UpdateWhateversRequestFilter) GetWhateversSome() *whateverfilter.WhateverFilter {
	if m != nil {
		return m.WhateversSome
	}
	return nil
}

func (m *UpdateWhateversRequestFilter) GetWhateversEvery() *whateverfilter.WhateverFilter {
	if m != nil {
		return m.WhateversEvery
	}
	return nil
}

func (m *UpdateWhateversRequestFilter) GetWhateversNone() *whateverfilter.WhateverFilter {
	if m != nil {
		return m.WhateversNone
	}
	return nil
}

func (m *UpdateWhateversRequestFilter) GetAnd() []*UpdateWhateversRequestFilter {
	if m != nil {
		return m.And
	}
	return nil
}

func (m *UpdateWhateversRequestFilter) GetOr() []*UpdateWhateversRequestFilter {
	if m != nil {
		return m.Or
	}
	return nil
}

func (m *UpdateWhateversRequestFilter) GetNot() []*UpdateWhateversRequestFilter {
	if m != nil {
		return m.Not
	}
	return nil
}

func (m *UpdateWhateversRequestFilter) GetHash() string {
	if m != nil && m.Hash != nil {
		return *m.Hash
	}
	return ""
}

func init() {
	proto.RegisterType((*UpdateWhateversRequestFilter)(nil), "updatewhateversrequestfilter.UpdateWhateversRequestFilter")
}

func init() {
	proto.RegisterFile("updatewhateversrequestfilter/updatewhateversrequestfilter_.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 308 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x92, 0x4f, 0x4b, 0xc3, 0x40,
	0x10, 0xc5, 0x69, 0x1b, 0xff, 0x6d, 0xd1, 0xc3, 0x9e, 0x96, 0x52, 0x24, 0x48, 0x0f, 0x41, 0x70,
	0x83, 0x39, 0x89, 0x27, 0x11, 0x2d, 0x52, 0xd4, 0x43, 0x44, 0x04, 0x2f, 0xb2, 0x26, 0x63, 0x12,
	0x68, 0x32, 0x71, 0x33, 0xa9, 0xf4, 0x1b, 0xf9, 0x31, 0x25, 0xe9, 0xa6, 0x10, 0x0b, 0x41, 0xe9,
	0x2d, 0x33, 0x79, 0xef, 0xf7, 0x5e, 0xb2, 0xcb, 0xae, 0xca, 0x3c, 0x54, 0x04, 0x5f, 0xb1, 0x22,
	0x58, 0x80, 0x2e, 0x34, 0x7c, 0x96, 0x50, 0xd0, 0x47, 0x32, 0x27, 0xd0, 0x6e, 0xd7, 0xcb, 0x37,
	0x99, 0x6b, 0x24, 0xe4, 0xe3, 0x2e, 0xd1, 0xe8, 0xd4, 0x8c, 0x29, 0x90, 0x32, 0xd0, 0x8d, 0x8d,
	0x21, 0x8d, 0x26, 0x0d, 0xc3, 0x08, 0xdb, 0xa3, 0x51, 0x9d, 0x7c, 0x5b, 0x6c, 0xfc, 0x5c, 0x47,
	0xbe, 0x34, 0x91, 0xfe, 0x8a, 0x38, 0xad, 0x75, 0xfc, 0x82, 0x59, 0x15, 0x5b, 0xf4, 0xec, 0x9e,
	0x33, 0xf4, 0x26, 0x72, 0x23, 0x4f, 0x1a, 0xfd, 0x03, 0x90, 0x5a, 0x79, 0xfc, 0xda, 0xc1, 0x6f,
	0xd8, 0xe1, 0xfa, 0x33, 0x9e, 0x30, 0x05, 0xd1, 0xaf, 0x11, 0xc7, 0xb2, 0xdd, 0x44, 0x36, 0xc1,
	0xc6, 0xdc, 0x36, 0xf1, 0x29, 0x3b, 0x5a, 0x2f, 0x6e, 0x17, 0xa0, 0x97, 0x62, 0xf0, 0x27, 0xcc,
	0x2f, 0x57, 0xab, 0xcd, 0x23, 0x66, 0x20, 0xac, 0x7f, 0xb6, 0xa9, 0x4c, 0xfc, 0x9e, 0x0d, 0x54,
	0x16, 0x8a, 0x1d, 0x7b, 0xe0, 0x0c, 0xbd, 0x4b, 0xd9, 0x75, 0x58, 0xb2, 0xeb, 0xb7, 0xfa, 0x15,
	0x86, 0xcf, 0x58, 0x1f, 0xb5, 0xd8, 0xdd, 0x1a, 0xd6, 0x47, 0x5d, 0x35, 0xcb, 0x90, 0xc4, 0xde,
	0xf6, 0xcd, 0x32, 0x24, 0xce, 0x99, 0x15, 0xab, 0x22, 0x16, 0xfb, 0x76, 0xcf, 0x39, 0xf0, 0xeb,
	0xe7, 0xeb, 0xd9, 0xeb, 0x5d, 0x94, 0x50, 0x5c, 0xbe, 0xcb, 0x00, 0x53, 0xd7, 0x3b, 0x2f, 0x28,
	0x41, 0x37, 0xc2, 0xb3, 0x7c, 0xae, 0x96, 0x91, 0xc6, 0x32, 0x0b, 0xdd, 0x48, 0xe7, 0x81, 0x5b,
	0x04, 0x31, 0xa4, 0xaa, 0xf3, 0xb6, 0xff, 0x04, 0x00, 0x00, 0xff, 0xff, 0x0e, 0xe3, 0x8f, 0x2e,
	0x29, 0x03, 0x00, 0x00,
}
