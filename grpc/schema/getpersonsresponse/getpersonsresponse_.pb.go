// Code generated by protoc-gen-go. DO NOT EDIT.
// source: getpersonsresponse/getpersonsresponse_.proto

/*
Package getpersonsresponse is a generated protocol buffer package.

It is generated from these files:
	getpersonsresponse/getpersonsresponse_.proto

It has these top-level messages:
	GetPersonsResponse
*/
package getpersonsresponse

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import responsemeta "github.com/21stio/go-playground/grpc/schema/responsemeta"
import person "github.com/21stio/go-playground/grpc/schema/person"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type GetPersonsResponse struct {
	Meta             *responsemeta.ResponseMeta `protobuf:"bytes,1,opt,name=meta" json:"meta,omitempty"`
	TotalCount       *int32                     `protobuf:"varint,2,opt,name=totalCount" json:"totalCount,omitempty"`
	Persons          []*person.Person           `protobuf:"bytes,3,rep,name=persons" json:"persons,omitempty"`
	Hash             *string                    `protobuf:"bytes,4,opt,name=hash" json:"hash,omitempty"`
	XXX_unrecognized []byte                     `json:"-"`
}

func (m *GetPersonsResponse) Reset()                    { *m = GetPersonsResponse{} }
func (m *GetPersonsResponse) String() string            { return proto.CompactTextString(m) }
func (*GetPersonsResponse) ProtoMessage()               {}
func (*GetPersonsResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *GetPersonsResponse) GetMeta() *responsemeta.ResponseMeta {
	if m != nil {
		return m.Meta
	}
	return nil
}

func (m *GetPersonsResponse) GetTotalCount() int32 {
	if m != nil && m.TotalCount != nil {
		return *m.TotalCount
	}
	return 0
}

func (m *GetPersonsResponse) GetPersons() []*person.Person {
	if m != nil {
		return m.Persons
	}
	return nil
}

func (m *GetPersonsResponse) GetHash() string {
	if m != nil && m.Hash != nil {
		return *m.Hash
	}
	return ""
}

func init() {
	proto.RegisterType((*GetPersonsResponse)(nil), "getpersonsresponse.GetPersonsResponse")
}

func init() { proto.RegisterFile("getpersonsresponse/getpersonsresponse_.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 230 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0x41, 0x4b, 0xc4, 0x30,
	0x10, 0x85, 0x89, 0x5b, 0x11, 0xb3, 0xe0, 0x61, 0xf0, 0x50, 0x7a, 0x90, 0xe0, 0x29, 0x07, 0x4d,
	0xb0, 0x3f, 0x40, 0x44, 0x0f, 0x9e, 0x04, 0xc9, 0xd1, 0x8b, 0xc4, 0x3a, 0xa4, 0x0b, 0xdb, 0x4e,
	0x48, 0xa6, 0x07, 0xff, 0x8e, 0xbf, 0x54, 0xdc, 0xb4, 0x50, 0xd8, 0x3d, 0x65, 0xe6, 0xbd, 0xf7,
	0xbd, 0x84, 0xc8, 0xbb, 0x80, 0x1c, 0x31, 0x65, 0x1a, 0x73, 0xc2, 0x1c, 0x69, 0xcc, 0x68, 0x8f,
	0xa5, 0x4f, 0x13, 0x13, 0x31, 0x01, 0x1c, 0x5b, 0x8d, 0x5a, 0xa6, 0x01, 0xd9, 0xdb, 0xf5, 0x32,
	0x53, 0xcd, 0x75, 0x41, 0x6c, 0x39, 0x66, 0xf5, 0xf6, 0x57, 0x48, 0x78, 0x45, 0x7e, 0x2f, 0x75,
	0x6e, 0xe6, 0xc0, 0xc8, 0xea, 0x9f, 0xad, 0x85, 0x12, 0x7a, 0xdb, 0x36, 0x66, 0x5d, 0x68, 0x96,
	0xd4, 0x1b, 0xb2, 0x77, 0x87, 0x1c, 0xdc, 0x48, 0xc9, 0xc4, 0x7e, 0xff, 0x42, 0xd3, 0xc8, 0xf5,
	0x99, 0x12, 0xfa, 0xdc, 0xad, 0x14, 0xd0, 0xf2, 0x62, 0x7e, 0x71, 0xbd, 0x51, 0x1b, 0xbd, 0x6d,
	0xaf, 0x4c, 0xd9, 0x4d, 0xb9, 0xd9, 0x2d, 0x36, 0x80, 0xac, 0x7a, 0x9f, 0xfb, 0xba, 0x52, 0x42,
	0x5f, 0xba, 0xc3, 0xfc, 0xfc, 0xf4, 0xf1, 0x18, 0x76, 0xdc, 0x4f, 0x5f, 0xa6, 0xa3, 0xc1, 0xb6,
	0x0f, 0x99, 0x77, 0x64, 0x03, 0xdd, 0xc7, 0xbd, 0xff, 0x09, 0x89, 0xa6, 0xf1, 0xdb, 0x86, 0x14,
	0x3b, 0x9b, 0xbb, 0x1e, 0x07, 0x7f, 0xe2, 0xe7, 0xfe, 0x02, 0x00, 0x00, 0xff, 0xff, 0xa7, 0xed,
	0x40, 0x2c, 0x61, 0x01, 0x00, 0x00,
}
