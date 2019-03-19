// Code generated by protoc-gen-go. DO NOT EDIT.
// source: deletecompaniesrequestfilter/deletecompaniesrequestfilter_.proto

/*
Package deletecompaniesrequestfilter is a generated protocol buffer package.

It is generated from these files:
	deletecompaniesrequestfilter/deletecompaniesrequestfilter_.proto

It has these top-level messages:
	DeleteCompaniesRequestFilter
*/
package deletecompaniesrequestfilter

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import requestmetafilter "github.com/21stio/go-playground/grpc/schema/requestmetafilter"
import idfilter "github.com/21stio/go-playground/grpc/schema/idfilter"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type DeleteCompaniesRequestFilter struct {
	Meta             *requestmetafilter.RequestMetaFilter `protobuf:"bytes,1,opt,name=meta" json:"meta,omitempty"`
	IdsSome          *idfilter.IdFilter                   `protobuf:"bytes,2,opt,name=idsSome" json:"idsSome,omitempty"`
	IdsEvery         *idfilter.IdFilter                   `protobuf:"bytes,3,opt,name=idsEvery" json:"idsEvery,omitempty"`
	IdsNone          *idfilter.IdFilter                   `protobuf:"bytes,4,opt,name=idsNone" json:"idsNone,omitempty"`
	And              []*DeleteCompaniesRequestFilter      `protobuf:"bytes,5,rep,name=and" json:"and,omitempty"`
	Or               []*DeleteCompaniesRequestFilter      `protobuf:"bytes,6,rep,name=or" json:"or,omitempty"`
	Not              []*DeleteCompaniesRequestFilter      `protobuf:"bytes,7,rep,name=not" json:"not,omitempty"`
	Hash             *string                              `protobuf:"bytes,8,opt,name=hash" json:"hash,omitempty"`
	XXX_unrecognized []byte                               `json:"-"`
}

func (m *DeleteCompaniesRequestFilter) Reset()                    { *m = DeleteCompaniesRequestFilter{} }
func (m *DeleteCompaniesRequestFilter) String() string            { return proto.CompactTextString(m) }
func (*DeleteCompaniesRequestFilter) ProtoMessage()               {}
func (*DeleteCompaniesRequestFilter) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *DeleteCompaniesRequestFilter) GetMeta() *requestmetafilter.RequestMetaFilter {
	if m != nil {
		return m.Meta
	}
	return nil
}

func (m *DeleteCompaniesRequestFilter) GetIdsSome() *idfilter.IdFilter {
	if m != nil {
		return m.IdsSome
	}
	return nil
}

func (m *DeleteCompaniesRequestFilter) GetIdsEvery() *idfilter.IdFilter {
	if m != nil {
		return m.IdsEvery
	}
	return nil
}

func (m *DeleteCompaniesRequestFilter) GetIdsNone() *idfilter.IdFilter {
	if m != nil {
		return m.IdsNone
	}
	return nil
}

func (m *DeleteCompaniesRequestFilter) GetAnd() []*DeleteCompaniesRequestFilter {
	if m != nil {
		return m.And
	}
	return nil
}

func (m *DeleteCompaniesRequestFilter) GetOr() []*DeleteCompaniesRequestFilter {
	if m != nil {
		return m.Or
	}
	return nil
}

func (m *DeleteCompaniesRequestFilter) GetNot() []*DeleteCompaniesRequestFilter {
	if m != nil {
		return m.Not
	}
	return nil
}

func (m *DeleteCompaniesRequestFilter) GetHash() string {
	if m != nil && m.Hash != nil {
		return *m.Hash
	}
	return ""
}

func init() {
	proto.RegisterType((*DeleteCompaniesRequestFilter)(nil), "deletecompaniesrequestfilter.DeleteCompaniesRequestFilter")
}

func init() {
	proto.RegisterFile("deletecompaniesrequestfilter/deletecompaniesrequestfilter_.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 301 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x92, 0x3f, 0x4f, 0xf3, 0x30,
	0x10, 0x87, 0xd5, 0xa6, 0xef, 0xdb, 0xe2, 0x6e, 0x9e, 0xac, 0xaa, 0x43, 0x85, 0x18, 0x22, 0x04,
	0x8e, 0xc8, 0x84, 0x98, 0x10, 0xff, 0x04, 0x15, 0x30, 0x84, 0x8d, 0x05, 0x99, 0xf8, 0x48, 0x2c,
	0x25, 0xbe, 0x60, 0x3b, 0x48, 0xfd, 0xb6, 0x7c, 0x14, 0x94, 0xd4, 0x61, 0xa9, 0xf0, 0xd2, 0xed,
	0x92, 0x7b, 0xee, 0xf9, 0x9d, 0x4e, 0x26, 0x97, 0x12, 0x2a, 0x70, 0x90, 0x63, 0xdd, 0x08, 0xad,
	0xc0, 0x1a, 0xf8, 0x6c, 0xc1, 0xba, 0x0f, 0x55, 0x39, 0x30, 0x49, 0xa8, 0xf9, 0xc6, 0x1b, 0x83,
	0x0e, 0xe9, 0x32, 0x04, 0x2d, 0x8e, 0xfd, 0x67, 0x0d, 0x4e, 0x78, 0xe9, 0xce, 0x1f, 0x6f, 0x5a,
	0x30, 0x25, 0x3d, 0x32, 0x14, 0xbe, 0x73, 0xf8, 0x1d, 0x91, 0xe5, 0x4d, 0x1f, 0x73, 0x3d, 0xc4,
	0x64, 0x5b, 0xcb, 0x5d, 0xcf, 0xd1, 0x73, 0x32, 0xe9, 0x7c, 0x6c, 0xb4, 0x1a, 0xc5, 0xf3, 0xf4,
	0x88, 0xef, 0x64, 0x70, 0xcf, 0x3f, 0x81, 0x13, 0xdb, 0x99, 0xac, 0x9f, 0xa0, 0x27, 0x64, 0xaa,
	0xa4, 0x7d, 0xc1, 0x1a, 0xd8, 0xb8, 0x1f, 0xa6, 0x7c, 0x48, 0xe7, 0x0f, 0xd2, 0xa3, 0x03, 0x42,
	0x39, 0x99, 0x29, 0x69, 0x6f, 0xbf, 0xc0, 0x6c, 0x58, 0xf4, 0x27, 0xfe, 0xcb, 0x78, 0xfb, 0x33,
	0x6a, 0x60, 0x93, 0xa0, 0xbd, 0x43, 0xe8, 0x23, 0x89, 0x84, 0x96, 0xec, 0xdf, 0x2a, 0x8a, 0xe7,
	0xe9, 0x05, 0x0f, 0x1d, 0x96, 0x87, 0xce, 0x91, 0x75, 0x1a, 0xba, 0x26, 0x63, 0x34, 0xec, 0xff,
	0xde, 0xb2, 0x31, 0x9a, 0x6e, 0x33, 0x8d, 0x8e, 0x4d, 0xf7, 0xdf, 0x4c, 0xa3, 0xa3, 0x94, 0x4c,
	0x4a, 0x61, 0x4b, 0x36, 0x5b, 0x8d, 0xe2, 0x83, 0xac, 0xaf, 0xaf, 0xd6, 0xaf, 0xf7, 0x85, 0x72,
	0x65, 0xfb, 0xce, 0x73, 0xac, 0x93, 0xf4, 0xcc, 0x3a, 0x85, 0x49, 0x81, 0xa7, 0x4d, 0x25, 0x36,
	0x85, 0xc1, 0x56, 0xcb, 0xa4, 0x30, 0x4d, 0x9e, 0xd8, 0xbc, 0x84, 0x5a, 0x04, 0x5f, 0xe6, 0x4f,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x8d, 0x3e, 0x73, 0x74, 0xd5, 0x02, 0x00, 0x00,
}
