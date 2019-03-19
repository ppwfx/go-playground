// Code generated by protoc-gen-go. DO NOT EDIT.
// source: getpersonbyidresponse/getpersonbyidresponse_.proto

/*
Package getpersonbyidresponse is a generated protocol buffer package.

It is generated from these files:
	getpersonbyidresponse/getpersonbyidresponse_.proto

It has these top-level messages:
	GetPersonByIdResponse
*/
package getpersonbyidresponse

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

type GetPersonByIdResponse struct {
	Meta             *responsemeta.ResponseMeta `protobuf:"bytes,1,opt,name=meta" json:"meta,omitempty"`
	Person           *person.Person             `protobuf:"bytes,2,opt,name=person" json:"person,omitempty"`
	Hash             *string                    `protobuf:"bytes,3,opt,name=hash" json:"hash,omitempty"`
	XXX_unrecognized []byte                     `json:"-"`
}

func (m *GetPersonByIdResponse) Reset()                    { *m = GetPersonByIdResponse{} }
func (m *GetPersonByIdResponse) String() string            { return proto.CompactTextString(m) }
func (*GetPersonByIdResponse) ProtoMessage()               {}
func (*GetPersonByIdResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *GetPersonByIdResponse) GetMeta() *responsemeta.ResponseMeta {
	if m != nil {
		return m.Meta
	}
	return nil
}

func (m *GetPersonByIdResponse) GetPerson() *person.Person {
	if m != nil {
		return m.Person
	}
	return nil
}

func (m *GetPersonByIdResponse) GetHash() string {
	if m != nil && m.Hash != nil {
		return *m.Hash
	}
	return ""
}

func init() {
	proto.RegisterType((*GetPersonByIdResponse)(nil), "getpersonbyidresponse.GetPersonByIdResponse")
}

func init() { proto.RegisterFile("getpersonbyidresponse/getpersonbyidresponse_.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 213 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0xbf, 0x4a, 0xc5, 0x30,
	0x14, 0xc6, 0x89, 0x5e, 0x04, 0x23, 0x38, 0x04, 0x2f, 0x94, 0x4e, 0xc5, 0x41, 0xba, 0x98, 0x60,
	0xdf, 0xc0, 0x3a, 0x88, 0x83, 0x20, 0x19, 0x5d, 0x24, 0x6d, 0x0f, 0x49, 0xc1, 0xf6, 0x84, 0x24,
	0x1d, 0x3a, 0xfb, 0xe2, 0x72, 0x93, 0x14, 0x3a, 0x74, 0xca, 0xf9, 0xfe, 0xfc, 0xbe, 0x21, 0xb4,
	0xd1, 0x10, 0x2c, 0x38, 0x8f, 0x73, 0xb7, 0x8e, 0x83, 0x03, 0x6f, 0x71, 0xf6, 0x20, 0x0e, 0xdd,
	0x1f, 0x6e, 0x1d, 0x06, 0x64, 0xe7, 0xc3, 0xb4, 0xac, 0xb6, 0x6b, 0x82, 0xa0, 0xc4, 0x5e, 0x64,
	0xb0, 0x7c, 0x48, 0x94, 0x48, 0x4f, 0x76, 0x1f, 0xff, 0x08, 0x3d, 0xbf, 0x43, 0xf8, 0x8a, 0x66,
	0xbb, 0x7e, 0x0c, 0x32, 0xa3, 0x8c, 0xd3, 0xd3, 0x05, 0x2f, 0x48, 0x45, 0xea, 0xbb, 0xa6, 0xe4,
	0xfb, 0x4d, 0xbe, 0xb5, 0x3e, 0x21, 0x28, 0x19, 0x7b, 0xec, 0x89, 0xde, 0xa4, 0xe9, 0xe2, 0x2a,
	0x12, 0xf7, 0x3c, 0x49, 0x9e, 0xb6, 0x65, 0x4e, 0x19, 0xa3, 0x27, 0xa3, 0xbc, 0x29, 0xae, 0x2b,
	0x52, 0xdf, 0xca, 0x78, 0xb7, 0x6f, 0xdf, 0xaf, 0x7a, 0x0c, 0x66, 0xe9, 0x78, 0x8f, 0x93, 0x68,
	0x5e, 0x7c, 0x18, 0x51, 0x68, 0x7c, 0xb6, 0xbf, 0x6a, 0xd5, 0x0e, 0x97, 0x79, 0x10, 0xda, 0xd9,
	0x5e, 0xf8, 0xde, 0xc0, 0xa4, 0x8e, 0x3f, 0xe8, 0x3f, 0x00, 0x00, 0xff, 0xff, 0x32, 0xa1, 0xcd,
	0xbd, 0x4e, 0x01, 0x00, 0x00,
}