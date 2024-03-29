// Code generated by protoc-gen-go. DO NOT EDIT.
// source: updatespecdescriptionsrequestfilter/updatespecdescriptionsrequestfilter_.proto

/*
Package updatespecdescriptionsrequestfilter is a generated protocol buffer package.

It is generated from these files:
	updatespecdescriptionsrequestfilter/updatespecdescriptionsrequestfilter_.proto

It has these top-level messages:
	UpdateSpecDescriptionsRequestFilter
*/
package updatespecdescriptionsrequestfilter

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import requestmetafilter "github.com/21stio/go-playground/grpc/schema/requestmetafilter"
import specdescriptionfilter "github.com/21stio/go-playground/grpc/schema/specdescriptionfilter"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type UpdateSpecDescriptionsRequestFilter struct {
	Meta                  *requestmetafilter.RequestMetaFilter         `protobuf:"bytes,1,opt,name=meta" json:"meta,omitempty"`
	SpecDescriptionsSome  *specdescriptionfilter.SpecDescriptionFilter `protobuf:"bytes,2,opt,name=specDescriptionsSome" json:"specDescriptionsSome,omitempty"`
	SpecDescriptionsEvery *specdescriptionfilter.SpecDescriptionFilter `protobuf:"bytes,3,opt,name=specDescriptionsEvery" json:"specDescriptionsEvery,omitempty"`
	SpecDescriptionsNone  *specdescriptionfilter.SpecDescriptionFilter `protobuf:"bytes,4,opt,name=specDescriptionsNone" json:"specDescriptionsNone,omitempty"`
	And                   []*UpdateSpecDescriptionsRequestFilter       `protobuf:"bytes,5,rep,name=and" json:"and,omitempty"`
	Or                    []*UpdateSpecDescriptionsRequestFilter       `protobuf:"bytes,6,rep,name=or" json:"or,omitempty"`
	Not                   []*UpdateSpecDescriptionsRequestFilter       `protobuf:"bytes,7,rep,name=not" json:"not,omitempty"`
	Hash                  *string                                      `protobuf:"bytes,8,opt,name=hash" json:"hash,omitempty"`
	XXX_unrecognized      []byte                                       `json:"-"`
}

func (m *UpdateSpecDescriptionsRequestFilter) Reset()         { *m = UpdateSpecDescriptionsRequestFilter{} }
func (m *UpdateSpecDescriptionsRequestFilter) String() string { return proto.CompactTextString(m) }
func (*UpdateSpecDescriptionsRequestFilter) ProtoMessage()    {}
func (*UpdateSpecDescriptionsRequestFilter) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{0}
}

func (m *UpdateSpecDescriptionsRequestFilter) GetMeta() *requestmetafilter.RequestMetaFilter {
	if m != nil {
		return m.Meta
	}
	return nil
}

func (m *UpdateSpecDescriptionsRequestFilter) GetSpecDescriptionsSome() *specdescriptionfilter.SpecDescriptionFilter {
	if m != nil {
		return m.SpecDescriptionsSome
	}
	return nil
}

func (m *UpdateSpecDescriptionsRequestFilter) GetSpecDescriptionsEvery() *specdescriptionfilter.SpecDescriptionFilter {
	if m != nil {
		return m.SpecDescriptionsEvery
	}
	return nil
}

func (m *UpdateSpecDescriptionsRequestFilter) GetSpecDescriptionsNone() *specdescriptionfilter.SpecDescriptionFilter {
	if m != nil {
		return m.SpecDescriptionsNone
	}
	return nil
}

func (m *UpdateSpecDescriptionsRequestFilter) GetAnd() []*UpdateSpecDescriptionsRequestFilter {
	if m != nil {
		return m.And
	}
	return nil
}

func (m *UpdateSpecDescriptionsRequestFilter) GetOr() []*UpdateSpecDescriptionsRequestFilter {
	if m != nil {
		return m.Or
	}
	return nil
}

func (m *UpdateSpecDescriptionsRequestFilter) GetNot() []*UpdateSpecDescriptionsRequestFilter {
	if m != nil {
		return m.Not
	}
	return nil
}

func (m *UpdateSpecDescriptionsRequestFilter) GetHash() string {
	if m != nil && m.Hash != nil {
		return *m.Hash
	}
	return ""
}

func init() {
	proto.RegisterType((*UpdateSpecDescriptionsRequestFilter)(nil), "updatespecdescriptionsrequestfilter.UpdateSpecDescriptionsRequestFilter")
}

func init() {
	proto.RegisterFile("updatespecdescriptionsrequestfilter/updatespecdescriptionsrequestfilter_.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 316 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x93, 0xcf, 0x4b, 0xc3, 0x30,
	0x14, 0xc7, 0xd9, 0x2f, 0x7f, 0x64, 0xb7, 0xa0, 0x10, 0x76, 0x1a, 0xce, 0xc3, 0x10, 0x4d, 0xb1,
	0x27, 0xcf, 0xa2, 0xe2, 0xc5, 0x89, 0x1d, 0x82, 0xec, 0xa2, 0x59, 0xfa, 0x6c, 0x0b, 0x6b, 0x5e,
	0x4c, 0x52, 0x61, 0x7f, 0x94, 0xff, 0xa3, 0xb4, 0xcb, 0x41, 0xba, 0x1e, 0x8a, 0xf4, 0xd6, 0x3e,
	0xbe, 0xef, 0xf3, 0xfd, 0x84, 0x36, 0x64, 0x51, 0xe8, 0x58, 0x38, 0xb0, 0x1a, 0x64, 0x0c, 0x56,
	0x9a, 0x4c, 0xbb, 0x0c, 0x95, 0x35, 0xf0, 0x55, 0x80, 0x75, 0x9f, 0xd9, 0xc6, 0x81, 0x09, 0x5a,
	0x64, 0xde, 0xb9, 0x36, 0xe8, 0x90, 0xce, 0x5a, 0x64, 0x27, 0x17, 0xfe, 0x35, 0x07, 0x27, 0x7c,
	0xc5, 0xde, 0xc4, 0x03, 0x27, 0x61, 0x0d, 0xe5, 0xf3, 0x8d, 0x53, 0xbf, 0x73, 0xf6, 0x33, 0x22,
	0xb3, 0xd7, 0xca, 0x63, 0xa9, 0x41, 0xde, 0xfd, 0xf1, 0x88, 0x76, 0x35, 0x0f, 0x55, 0x9c, 0xde,
	0x90, 0x61, 0x59, 0xc8, 0x7a, 0xd3, 0xde, 0x7c, 0x1c, 0x9e, 0xf3, 0x3d, 0x09, 0xee, 0xf3, 0x4f,
	0xe0, 0xc4, 0x6e, 0x27, 0xaa, 0x36, 0xe8, 0x07, 0x39, 0xb1, 0x35, 0xf4, 0x12, 0x73, 0x60, 0xfd,
	0x8a, 0x74, 0xc9, 0x1b, 0xf5, 0x78, 0xcd, 0xc6, 0x13, 0x1b, 0x49, 0x74, 0x4d, 0x4e, 0xeb, 0xf3,
	0xfb, 0x6f, 0x30, 0x5b, 0x36, 0xf8, 0x47, 0x45, 0x33, 0xaa, 0xe9, 0x14, 0x0b, 0x54, 0xc0, 0x86,
	0x5d, 0x9c, 0xa2, 0x24, 0xd1, 0x15, 0x19, 0x08, 0x15, 0xb3, 0xd1, 0x74, 0x30, 0x1f, 0x87, 0x8f,
	0xbc, 0xc5, 0xcf, 0xc1, 0x5b, 0x7c, 0xb8, 0xa8, 0x84, 0xd2, 0x37, 0xd2, 0x47, 0xc3, 0x0e, 0x3a,
	0x46, 0xf7, 0xd1, 0x94, 0xd6, 0x0a, 0x1d, 0x3b, 0xec, 0xda, 0x5a, 0xa1, 0xa3, 0x94, 0x0c, 0x53,
	0x61, 0x53, 0x76, 0x34, 0xed, 0xcd, 0x8f, 0xa3, 0xea, 0xf9, 0xf6, 0x65, 0xf5, 0x9c, 0x64, 0x2e,
	0x2d, 0xd6, 0x5c, 0x62, 0x1e, 0x84, 0xd7, 0xd6, 0x65, 0x18, 0x24, 0x78, 0xa5, 0x37, 0x62, 0x9b,
	0x18, 0x2c, 0x54, 0x1c, 0x24, 0x46, 0xcb, 0xc0, 0xca, 0x14, 0x72, 0xd1, 0xe6, 0x3a, 0xfe, 0x06,
	0x00, 0x00, 0xff, 0xff, 0xb7, 0xe2, 0x6b, 0x7d, 0xd8, 0x03, 0x00, 0x00,
}
