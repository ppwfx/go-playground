// Code generated by protoc-gen-go. DO NOT EDIT.
// source: passupdatewhateversrequestrequestfilter/passupdatewhateversrequestrequestfilter_.proto

/*
Package passupdatewhateversrequestrequestfilter is a generated protocol buffer package.

It is generated from these files:
	passupdatewhateversrequestrequestfilter/passupdatewhateversrequestrequestfilter_.proto

It has these top-level messages:
	PassUpdateWhateversRequestRequestFilter
*/
package passupdatewhateversrequestrequestfilter

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import requestmetafilter "github.com/21stio/go-playground/grpc/schema/requestmetafilter"
import updatewhateversrequestfilter "github.com/21stio/go-playground/grpc/schema/updatewhateversrequestfilter"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type PassUpdateWhateversRequestRequestFilter struct {
	Meta                   *requestmetafilter.RequestMetaFilter                       `protobuf:"bytes,1,opt,name=meta" json:"meta,omitempty"`
	UpdateWhateversRequest *updatewhateversrequestfilter.UpdateWhateversRequestFilter `protobuf:"bytes,2,opt,name=updateWhateversRequest" json:"updateWhateversRequest,omitempty"`
	And                    []*PassUpdateWhateversRequestRequestFilter                 `protobuf:"bytes,3,rep,name=and" json:"and,omitempty"`
	Or                     []*PassUpdateWhateversRequestRequestFilter                 `protobuf:"bytes,4,rep,name=or" json:"or,omitempty"`
	Not                    []*PassUpdateWhateversRequestRequestFilter                 `protobuf:"bytes,5,rep,name=not" json:"not,omitempty"`
	Hash                   *string                                                    `protobuf:"bytes,6,opt,name=hash" json:"hash,omitempty"`
	XXX_unrecognized       []byte                                                     `json:"-"`
}

func (m *PassUpdateWhateversRequestRequestFilter) Reset() {
	*m = PassUpdateWhateversRequestRequestFilter{}
}
func (m *PassUpdateWhateversRequestRequestFilter) String() string { return proto.CompactTextString(m) }
func (*PassUpdateWhateversRequestRequestFilter) ProtoMessage()    {}
func (*PassUpdateWhateversRequestRequestFilter) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{0}
}

func (m *PassUpdateWhateversRequestRequestFilter) GetMeta() *requestmetafilter.RequestMetaFilter {
	if m != nil {
		return m.Meta
	}
	return nil
}

func (m *PassUpdateWhateversRequestRequestFilter) GetUpdateWhateversRequest() *updatewhateversrequestfilter.UpdateWhateversRequestFilter {
	if m != nil {
		return m.UpdateWhateversRequest
	}
	return nil
}

func (m *PassUpdateWhateversRequestRequestFilter) GetAnd() []*PassUpdateWhateversRequestRequestFilter {
	if m != nil {
		return m.And
	}
	return nil
}

func (m *PassUpdateWhateversRequestRequestFilter) GetOr() []*PassUpdateWhateversRequestRequestFilter {
	if m != nil {
		return m.Or
	}
	return nil
}

func (m *PassUpdateWhateversRequestRequestFilter) GetNot() []*PassUpdateWhateversRequestRequestFilter {
	if m != nil {
		return m.Not
	}
	return nil
}

func (m *PassUpdateWhateversRequestRequestFilter) GetHash() string {
	if m != nil && m.Hash != nil {
		return *m.Hash
	}
	return ""
}

func init() {
	proto.RegisterType((*PassUpdateWhateversRequestRequestFilter)(nil), "passupdatewhateversrequestrequestfilter.PassUpdateWhateversRequestRequestFilter")
}

func init() {
	proto.RegisterFile("passupdatewhateversrequestrequestfilter/passupdatewhateversrequestrequestfilter_.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 286 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x53, 0x3b, 0x4b, 0xc4, 0x40,
	0x10, 0x26, 0x0f, 0x05, 0xf7, 0xba, 0x2d, 0x64, 0xb9, 0x2a, 0x88, 0x70, 0x41, 0x70, 0x83, 0xa9,
	0xc4, 0x4a, 0x2c, 0xec, 0x84, 0x23, 0x72, 0x0a, 0x36, 0x3a, 0x97, 0xac, 0x49, 0xe0, 0x92, 0x89,
	0xbb, 0x13, 0xc5, 0x1f, 0xe2, 0xff, 0x95, 0xe4, 0xd6, 0xea, 0xe2, 0xb1, 0xcd, 0x55, 0xb3, 0x8f,
	0xef, 0x35, 0x03, 0xc3, 0x9e, 0x3a, 0x30, 0xa6, 0xef, 0x0a, 0x20, 0xf5, 0x55, 0x01, 0xa9, 0x4f,
	0xa5, 0x8d, 0x56, 0x1f, 0xbd, 0x32, 0x64, 0xcb, 0x7b, 0xbd, 0x21, 0xa5, 0x13, 0x47, 0xdc, 0xab,
	0xec, 0x34, 0x12, 0xf2, 0x85, 0x23, 0x7e, 0x7e, 0x61, 0xaf, 0x8d, 0x22, 0xb0, 0x56, 0x3b, 0x2f,
	0x56, 0x74, 0x7e, 0x3b, 0x2d, 0x68, 0x69, 0xfb, 0x3e, 0xad, 0xc2, 0xd9, 0x4f, 0xc8, 0x16, 0x4b,
	0x30, 0x66, 0x35, 0x62, 0x9f, 0xff, 0xb0, 0xd9, 0x16, 0x6b, 0xcb, 0xfd, 0x48, 0xe1, 0xd7, 0x2c,
	0x1c, 0x22, 0x08, 0x2f, 0xf2, 0xe2, 0x59, 0x7a, 0x2e, 0x77, 0x62, 0x49, 0x8b, 0x7f, 0x50, 0x04,
	0x5b, 0x4e, 0x36, 0x32, 0xb8, 0x66, 0xa7, 0xfd, 0xa4, 0x81, 0xf0, 0x47, 0xad, 0x1b, 0xb9, 0x2f,
	0xab, 0x9c, 0x0e, 0x67, 0x1d, 0xfe, 0x51, 0xe6, 0x6b, 0x16, 0x40, 0x5b, 0x88, 0x20, 0x0a, 0xe2,
	0x59, 0xba, 0x94, 0x8e, 0xe3, 0x97, 0x8e, 0xc3, 0xc8, 0x06, 0x71, 0xfe, 0xc6, 0x7c, 0xd4, 0x22,
	0x3c, 0x90, 0x85, 0x8f, 0x7a, 0xe8, 0xa2, 0x45, 0x12, 0x47, 0x87, 0xea, 0xa2, 0x45, 0xe2, 0x9c,
	0x85, 0x15, 0x98, 0x4a, 0x1c, 0x47, 0x5e, 0x7c, 0x92, 0x8d, 0xe7, 0xbb, 0xd5, 0xcb, 0x63, 0x59,
	0x53, 0xd5, 0xaf, 0x65, 0x8e, 0x4d, 0x92, 0x5e, 0x19, 0xaa, 0x31, 0x29, 0xf1, 0xb2, 0xdb, 0xc0,
	0x77, 0xa9, 0xb1, 0x6f, 0x8b, 0xa4, 0xd4, 0x5d, 0x9e, 0x98, 0xbc, 0x52, 0x0d, 0xb8, 0x2e, 0xc3,
	0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x36, 0x47, 0x63, 0x55, 0x5e, 0x03, 0x00, 0x00,
}
