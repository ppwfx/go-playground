// Code generated by protoc-gen-go. DO NOT EDIT.
// source: passcreatewhateversrequestrequestfilter/passcreatewhateversrequestrequestfilter_.proto

/*
Package passcreatewhateversrequestrequestfilter is a generated protocol buffer package.

It is generated from these files:
	passcreatewhateversrequestrequestfilter/passcreatewhateversrequestrequestfilter_.proto

It has these top-level messages:
	PassCreateWhateversRequestRequestFilter
*/
package passcreatewhateversrequestrequestfilter

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import requestmetafilter "github.com/21stio/go-playground/grpc/schema/requestmetafilter"
import createwhateversrequestfilter "github.com/21stio/go-playground/grpc/schema/createwhateversrequestfilter"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type PassCreateWhateversRequestRequestFilter struct {
	Meta                   *requestmetafilter.RequestMetaFilter                       `protobuf:"bytes,1,opt,name=meta" json:"meta,omitempty"`
	CreateWhateversRequest *createwhateversrequestfilter.CreateWhateversRequestFilter `protobuf:"bytes,2,opt,name=createWhateversRequest" json:"createWhateversRequest,omitempty"`
	And                    []*PassCreateWhateversRequestRequestFilter                 `protobuf:"bytes,3,rep,name=and" json:"and,omitempty"`
	Or                     []*PassCreateWhateversRequestRequestFilter                 `protobuf:"bytes,4,rep,name=or" json:"or,omitempty"`
	Not                    []*PassCreateWhateversRequestRequestFilter                 `protobuf:"bytes,5,rep,name=not" json:"not,omitempty"`
	Hash                   *string                                                    `protobuf:"bytes,6,opt,name=hash" json:"hash,omitempty"`
	XXX_unrecognized       []byte                                                     `json:"-"`
}

func (m *PassCreateWhateversRequestRequestFilter) Reset() {
	*m = PassCreateWhateversRequestRequestFilter{}
}
func (m *PassCreateWhateversRequestRequestFilter) String() string { return proto.CompactTextString(m) }
func (*PassCreateWhateversRequestRequestFilter) ProtoMessage()    {}
func (*PassCreateWhateversRequestRequestFilter) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{0}
}

func (m *PassCreateWhateversRequestRequestFilter) GetMeta() *requestmetafilter.RequestMetaFilter {
	if m != nil {
		return m.Meta
	}
	return nil
}

func (m *PassCreateWhateversRequestRequestFilter) GetCreateWhateversRequest() *createwhateversrequestfilter.CreateWhateversRequestFilter {
	if m != nil {
		return m.CreateWhateversRequest
	}
	return nil
}

func (m *PassCreateWhateversRequestRequestFilter) GetAnd() []*PassCreateWhateversRequestRequestFilter {
	if m != nil {
		return m.And
	}
	return nil
}

func (m *PassCreateWhateversRequestRequestFilter) GetOr() []*PassCreateWhateversRequestRequestFilter {
	if m != nil {
		return m.Or
	}
	return nil
}

func (m *PassCreateWhateversRequestRequestFilter) GetNot() []*PassCreateWhateversRequestRequestFilter {
	if m != nil {
		return m.Not
	}
	return nil
}

func (m *PassCreateWhateversRequestRequestFilter) GetHash() string {
	if m != nil && m.Hash != nil {
		return *m.Hash
	}
	return ""
}

func init() {
	proto.RegisterType((*PassCreateWhateversRequestRequestFilter)(nil), "passcreatewhateversrequestrequestfilter.PassCreateWhateversRequestRequestFilter")
}

func init() {
	proto.RegisterFile("passcreatewhateversrequestrequestfilter/passcreatewhateversrequestrequestfilter_.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 285 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x53, 0xcd, 0x4a, 0xc3, 0x40,
	0x10, 0xa6, 0x49, 0x14, 0xdc, 0xde, 0xf6, 0x20, 0x4b, 0x4f, 0x41, 0x84, 0x06, 0xc1, 0x0d, 0xe6,
	0x24, 0x9e, 0x44, 0xc1, 0x9b, 0x50, 0x22, 0x2a, 0x78, 0xd1, 0x6d, 0x3a, 0x26, 0x81, 0x26, 0x13,
	0x67, 0x27, 0x8a, 0x0f, 0xe2, 0xfb, 0x4a, 0xd2, 0xf5, 0xd4, 0x58, 0x72, 0xe9, 0x69, 0xf6, 0xe7,
	0xfb, 0x9b, 0x81, 0x11, 0x4f, 0x8d, 0xb1, 0x36, 0x23, 0x30, 0x0c, 0x5f, 0x85, 0x61, 0xf8, 0x04,
	0xb2, 0x04, 0x1f, 0x2d, 0x58, 0x76, 0xe5, 0xbd, 0x5c, 0x33, 0x50, 0x3c, 0x12, 0xf7, 0xaa, 0x1b,
	0x42, 0x46, 0x39, 0x1f, 0x89, 0x9f, 0x9d, 0xb9, 0x6b, 0x05, 0x6c, 0x9c, 0xd5, 0xd6, 0x8b, 0x13,
	0x9d, 0x5d, 0x0f, 0x0b, 0x3a, 0xda, 0xae, 0x4f, 0xa7, 0x70, 0xf2, 0x13, 0x88, 0xf9, 0xc2, 0x58,
	0x7b, 0xdb, 0x63, 0x9f, 0xff, 0xb0, 0xe9, 0x06, 0xeb, 0xca, 0x5d, 0x4f, 0x91, 0x97, 0x22, 0xe8,
	0x22, 0xa8, 0x49, 0x38, 0x89, 0xa6, 0xc9, 0xa9, 0xde, 0x8a, 0xa5, 0x1d, 0xfe, 0x1e, 0xd8, 0x6c,
	0x38, 0x69, 0xcf, 0x90, 0x24, 0x8e, 0xb3, 0x41, 0x03, 0xe5, 0xf5, 0x5a, 0x57, 0x7a, 0x57, 0x56,
	0x3d, 0x1c, 0xce, 0x39, 0xfc, 0xa3, 0x2c, 0x97, 0xc2, 0x37, 0xf5, 0x4a, 0xf9, 0xa1, 0x1f, 0x4d,
	0x93, 0x85, 0x1e, 0x39, 0x7e, 0x3d, 0x72, 0x18, 0x69, 0x27, 0x2e, 0xdf, 0x84, 0x87, 0xa4, 0x82,
	0x3d, 0x59, 0x78, 0x48, 0x5d, 0x17, 0x35, 0xb2, 0x3a, 0xd8, 0x57, 0x17, 0x35, 0xb2, 0x94, 0x22,
	0x28, 0x8c, 0x2d, 0xd4, 0x61, 0x38, 0x89, 0x8e, 0xd2, 0xfe, 0x7c, 0xf3, 0xf8, 0xf2, 0x90, 0x97,
	0x5c, 0xb4, 0x4b, 0x9d, 0x61, 0x15, 0x27, 0x17, 0x96, 0x4b, 0x8c, 0x73, 0x3c, 0x6f, 0xd6, 0xe6,
	0x3b, 0x27, 0x6c, 0xeb, 0x55, 0x9c, 0x53, 0x93, 0xc5, 0x36, 0x2b, 0xa0, 0x32, 0x63, 0x97, 0xe1,
	0x37, 0x00, 0x00, 0xff, 0xff, 0xa5, 0xa6, 0xda, 0x5a, 0x5e, 0x03, 0x00, 0x00,
}
