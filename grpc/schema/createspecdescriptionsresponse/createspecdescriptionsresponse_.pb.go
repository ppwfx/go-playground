// Code generated by protoc-gen-go. DO NOT EDIT.
// source: createspecdescriptionsresponse/createspecdescriptionsresponse_.proto

/*
Package createspecdescriptionsresponse is a generated protocol buffer package.

It is generated from these files:
	createspecdescriptionsresponse/createspecdescriptionsresponse_.proto

It has these top-level messages:
	CreateSpecDescriptionsResponse
*/
package createspecdescriptionsresponse

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import responsemeta "github.com/21stio/go-playground/grpc/schema/responsemeta"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type CreateSpecDescriptionsResponse struct {
	Meta             *responsemeta.ResponseMeta `protobuf:"bytes,1,opt,name=meta" json:"meta,omitempty"`
	Hash             *string                    `protobuf:"bytes,2,opt,name=hash" json:"hash,omitempty"`
	XXX_unrecognized []byte                     `json:"-"`
}

func (m *CreateSpecDescriptionsResponse) Reset()                    { *m = CreateSpecDescriptionsResponse{} }
func (m *CreateSpecDescriptionsResponse) String() string            { return proto.CompactTextString(m) }
func (*CreateSpecDescriptionsResponse) ProtoMessage()               {}
func (*CreateSpecDescriptionsResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *CreateSpecDescriptionsResponse) GetMeta() *responsemeta.ResponseMeta {
	if m != nil {
		return m.Meta
	}
	return nil
}

func (m *CreateSpecDescriptionsResponse) GetHash() string {
	if m != nil && m.Hash != nil {
		return *m.Hash
	}
	return ""
}

func init() {
	proto.RegisterType((*CreateSpecDescriptionsResponse)(nil), "createspecdescriptionsresponse.CreateSpecDescriptionsResponse")
}

func init() {
	proto.RegisterFile("createspecdescriptionsresponse/createspecdescriptionsresponse_.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 190 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x8f, 0x31, 0xab, 0x83, 0x30,
	0x10, 0x80, 0xf1, 0xe1, 0xf2, 0xf2, 0xb6, 0x4c, 0xe2, 0x20, 0xf2, 0x26, 0x97, 0x26, 0xd4, 0x9f,
	0xd0, 0x3a, 0x95, 0x76, 0xb1, 0x5b, 0x97, 0x92, 0xc6, 0x23, 0x09, 0x54, 0x2f, 0xe4, 0xe2, 0xd0,
	0x7f, 0x5f, 0x2a, 0x0a, 0x4e, 0x76, 0xfb, 0x8e, 0xfb, 0xee, 0x83, 0x63, 0x8d, 0x0e, 0xa0, 0x22,
	0x90, 0x07, 0xdd, 0x01, 0xe9, 0xe0, 0x7c, 0x74, 0x38, 0x50, 0x00, 0xf2, 0x38, 0x10, 0xc8, 0xed,
	0xf5, 0x5d, 0xf8, 0x80, 0x11, 0x79, 0xb1, 0xad, 0xe5, 0xe5, 0x42, 0x3d, 0x44, 0x25, 0xd7, 0xc3,
	0x5c, 0xf8, 0xef, 0x58, 0x71, 0x9c, 0x1a, 0x57, 0x0f, 0xba, 0x59, 0x35, 0xda, 0x59, 0xe6, 0x82,
	0xa5, 0x9f, 0x83, 0x2c, 0x29, 0x93, 0xea, 0xaf, 0xce, 0xc5, 0xba, 0x22, 0x16, 0xeb, 0x02, 0x51,
	0xb5, 0x93, 0xc7, 0x39, 0x4b, 0xad, 0x22, 0x9b, 0xfd, 0x94, 0x49, 0xf5, 0xdb, 0x4e, 0x7c, 0x38,
	0xdf, 0x4e, 0xc6, 0x45, 0x3b, 0x3e, 0x84, 0xc6, 0x5e, 0xd6, 0x7b, 0x8a, 0x0e, 0xa5, 0xc1, 0x9d,
	0x7f, 0xaa, 0x97, 0x09, 0x38, 0x0e, 0x9d, 0x34, 0xc1, 0x6b, 0x49, 0xda, 0x42, 0xaf, 0xbe, 0x3c,
	0xff, 0x0e, 0x00, 0x00, 0xff, 0xff, 0xae, 0xed, 0x86, 0x74, 0x3c, 0x01, 0x00, 0x00,
}