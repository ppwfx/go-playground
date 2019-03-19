// Code generated by protoc-gen-go. DO NOT EDIT.
// source: getpersonlistbyidrequestfilter/getpersonlistbyidrequestfilter_.proto

/*
Package getpersonlistbyidrequestfilter is a generated protocol buffer package.

It is generated from these files:
	getpersonlistbyidrequestfilter/getpersonlistbyidrequestfilter_.proto

It has these top-level messages:
	GetPersonListByIdRequestFilter
*/
package getpersonlistbyidrequestfilter

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import requestmetafilter "github.com/21stio/go-playground/grpc/schema/requestmetafilter"
import idfilter "github.com/21stio/go-playground/grpc/schema/idfilter"
import personskindfilter "github.com/21stio/go-playground/grpc/schema/personskindfilter"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type GetPersonListByIdRequestFilter struct {
	Meta             *requestmetafilter.RequestMetaFilter `protobuf:"bytes,1,opt,name=meta" json:"meta,omitempty"`
	Id               *idfilter.IdFilter                   `protobuf:"bytes,2,opt,name=id" json:"id,omitempty"`
	Kind             *personskindfilter.PersonsKindFilter `protobuf:"bytes,3,opt,name=kind" json:"kind,omitempty"`
	And              []*GetPersonListByIdRequestFilter    `protobuf:"bytes,4,rep,name=and" json:"and,omitempty"`
	Or               []*GetPersonListByIdRequestFilter    `protobuf:"bytes,5,rep,name=or" json:"or,omitempty"`
	Not              []*GetPersonListByIdRequestFilter    `protobuf:"bytes,6,rep,name=not" json:"not,omitempty"`
	Hash             *string                              `protobuf:"bytes,7,opt,name=hash" json:"hash,omitempty"`
	XXX_unrecognized []byte                               `json:"-"`
}

func (m *GetPersonListByIdRequestFilter) Reset()                    { *m = GetPersonListByIdRequestFilter{} }
func (m *GetPersonListByIdRequestFilter) String() string            { return proto.CompactTextString(m) }
func (*GetPersonListByIdRequestFilter) ProtoMessage()               {}
func (*GetPersonListByIdRequestFilter) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *GetPersonListByIdRequestFilter) GetMeta() *requestmetafilter.RequestMetaFilter {
	if m != nil {
		return m.Meta
	}
	return nil
}

func (m *GetPersonListByIdRequestFilter) GetId() *idfilter.IdFilter {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *GetPersonListByIdRequestFilter) GetKind() *personskindfilter.PersonsKindFilter {
	if m != nil {
		return m.Kind
	}
	return nil
}

func (m *GetPersonListByIdRequestFilter) GetAnd() []*GetPersonListByIdRequestFilter {
	if m != nil {
		return m.And
	}
	return nil
}

func (m *GetPersonListByIdRequestFilter) GetOr() []*GetPersonListByIdRequestFilter {
	if m != nil {
		return m.Or
	}
	return nil
}

func (m *GetPersonListByIdRequestFilter) GetNot() []*GetPersonListByIdRequestFilter {
	if m != nil {
		return m.Not
	}
	return nil
}

func (m *GetPersonListByIdRequestFilter) GetHash() string {
	if m != nil && m.Hash != nil {
		return *m.Hash
	}
	return ""
}

func init() {
	proto.RegisterType((*GetPersonListByIdRequestFilter)(nil), "getpersonlistbyidrequestfilter.GetPersonListByIdRequestFilter")
}

func init() {
	proto.RegisterFile("getpersonlistbyidrequestfilter/getpersonlistbyidrequestfilter_.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 308 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x92, 0x31, 0x4f, 0xc3, 0x30,
	0x10, 0x85, 0xd5, 0xa4, 0x80, 0x70, 0x37, 0x4f, 0x56, 0x87, 0xaa, 0xaa, 0x18, 0xa2, 0x4a, 0x38,
	0x22, 0x13, 0x13, 0x43, 0x85, 0x40, 0x85, 0x82, 0xaa, 0x8c, 0x2c, 0xc8, 0x8d, 0x4d, 0x62, 0x91,
	0xd8, 0xc1, 0xbe, 0x0c, 0xf9, 0x77, 0xfc, 0x34, 0x14, 0xc7, 0x99, 0x22, 0x85, 0xa5, 0xdb, 0xf9,
	0xee, 0xbd, 0x77, 0x9f, 0x2d, 0xa3, 0xc7, 0x5c, 0x40, 0x2d, 0x8c, 0xd5, 0xaa, 0x94, 0x16, 0x4e,
	0xad, 0xe4, 0x46, 0xfc, 0x34, 0xc2, 0xc2, 0x97, 0x2c, 0x41, 0x98, 0x78, 0x7a, 0xfc, 0x49, 0x6b,
	0xa3, 0x41, 0xe3, 0xd5, 0xb4, 0x6c, 0xb9, 0xf5, 0xc7, 0x4a, 0x00, 0xf3, 0xc1, 0xa3, 0x8e, 0xcf,
	0x5a, 0x12, 0xc9, 0xbd, 0x64, 0x28, 0x86, 0xc9, 0xb6, 0x5f, 0x61, 0xbf, 0xa5, 0x1a, 0x24, 0xa3,
	0x8e, 0xd7, 0x6e, 0x7e, 0x43, 0xb4, 0x7a, 0x16, 0x70, 0x74, 0xf3, 0x83, 0xb4, 0xb0, 0x6b, 0xf7,
	0x3c, 0xed, 0x77, 0x3e, 0x39, 0x25, 0xbe, 0x47, 0xf3, 0x6e, 0x3b, 0x99, 0xad, 0x67, 0xd1, 0x22,
	0xb9, 0xa1, 0x23, 0x22, 0xea, 0xf5, 0x6f, 0x02, 0x58, 0xef, 0x49, 0x9d, 0x03, 0x6f, 0x50, 0x20,
	0x39, 0x09, 0x9c, 0x0f, 0xd3, 0x01, 0x93, 0xee, 0xb9, 0x57, 0x05, 0x92, 0x77, 0xe9, 0x1d, 0x15,
	0x09, 0x7d, 0xfa, 0x88, 0x94, 0xf6, 0x6c, 0xf6, 0x55, 0xaa, 0xc1, 0xe7, 0x1c, 0xf8, 0x88, 0x42,
	0xa6, 0x38, 0x99, 0xaf, 0xc3, 0x68, 0x91, 0x3c, 0xd0, 0xe9, 0xa7, 0xa5, 0xd3, 0x97, 0x4c, 0xbb,
	0x28, 0xfc, 0x8e, 0x02, 0x6d, 0xc8, 0xc5, 0x59, 0x02, 0x03, 0x6d, 0x3a, 0x42, 0xa5, 0x81, 0x5c,
	0x9e, 0x87, 0x50, 0x69, 0xc0, 0x18, 0xcd, 0x0b, 0x66, 0x0b, 0x72, 0xb5, 0x9e, 0x45, 0xd7, 0xa9,
	0xab, 0x77, 0x87, 0x8f, 0x97, 0x5c, 0x42, 0xd1, 0x9c, 0x68, 0xa6, 0xab, 0x38, 0xb9, 0xb3, 0x20,
	0x75, 0x9c, 0xeb, 0xdb, 0xba, 0x64, 0x6d, 0x6e, 0x74, 0xa3, 0x78, 0x9c, 0x9b, 0x3a, 0x8b, 0x6d,
	0x56, 0x88, 0x8a, 0xfd, 0xf3, 0x53, 0xff, 0x02, 0x00, 0x00, 0xff, 0xff, 0x4e, 0x1c, 0xe6, 0xa8,
	0xe9, 0x02, 0x00, 0x00,
}
