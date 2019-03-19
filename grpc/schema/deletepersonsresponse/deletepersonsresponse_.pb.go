// Code generated by protoc-gen-go. DO NOT EDIT.
// source: deletepersonsresponse/deletepersonsresponse_.proto

/*
Package deletepersonsresponse is a generated protocol buffer package.

It is generated from these files:
	deletepersonsresponse/deletepersonsresponse_.proto

It has these top-level messages:
	DeletePersonsResponse
*/
package deletepersonsresponse

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

type DeletePersonsResponse struct {
	Meta             *responsemeta.ResponseMeta `protobuf:"bytes,1,opt,name=meta" json:"meta,omitempty"`
	Hash             *string                    `protobuf:"bytes,2,opt,name=hash" json:"hash,omitempty"`
	XXX_unrecognized []byte                     `json:"-"`
}

func (m *DeletePersonsResponse) Reset()                    { *m = DeletePersonsResponse{} }
func (m *DeletePersonsResponse) String() string            { return proto.CompactTextString(m) }
func (*DeletePersonsResponse) ProtoMessage()               {}
func (*DeletePersonsResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *DeletePersonsResponse) GetMeta() *responsemeta.ResponseMeta {
	if m != nil {
		return m.Meta
	}
	return nil
}

func (m *DeletePersonsResponse) GetHash() string {
	if m != nil && m.Hash != nil {
		return *m.Hash
	}
	return ""
}

func init() {
	proto.RegisterType((*DeletePersonsResponse)(nil), "deletepersonsresponse.DeletePersonsResponse")
}

func init() { proto.RegisterFile("deletepersonsresponse/deletepersonsresponse_.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 179 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x8f, 0x3d, 0x0b, 0x83, 0x30,
	0x10, 0x40, 0xb1, 0xb8, 0x34, 0xdd, 0x02, 0x82, 0x38, 0x49, 0x27, 0x97, 0x26, 0xd4, 0x7f, 0xd0,
	0x8f, 0xb5, 0x50, 0x1c, 0xdb, 0xa1, 0xa4, 0x7a, 0x24, 0x82, 0x7a, 0x21, 0x17, 0x87, 0xfe, 0xfb,
	0x52, 0xab, 0xe0, 0x90, 0xed, 0xdd, 0xdd, 0x7b, 0xc3, 0xb1, 0xb2, 0x81, 0x0e, 0x3c, 0x58, 0x70,
	0x84, 0x03, 0x39, 0x20, 0x8b, 0x03, 0x81, 0x0c, 0x6e, 0x5f, 0xc2, 0x3a, 0xf4, 0xc8, 0x93, 0xe0,
	0x35, 0xcb, 0x17, 0xea, 0xc1, 0x2b, 0xb9, 0x1e, 0xe6, 0x70, 0xff, 0x64, 0xc9, 0x75, 0x4a, 0xef,
	0xff, 0xb4, 0x9a, 0x1d, 0x2e, 0x58, 0xfc, 0xf3, 0xd2, 0x28, 0x8f, 0x8a, 0x5d, 0x99, 0x89, 0x75,
	0x2c, 0x16, 0xeb, 0x06, 0x5e, 0x55, 0x93, 0xc7, 0x39, 0x8b, 0x8d, 0x22, 0x93, 0x6e, 0xf2, 0xa8,
	0xd8, 0x56, 0x13, 0x9f, 0x2f, 0x8f, 0x93, 0x6e, 0xbd, 0x19, 0xdf, 0xa2, 0xc6, 0x5e, 0x96, 0x47,
	0xf2, 0x2d, 0x4a, 0x8d, 0x07, 0xdb, 0xa9, 0x8f, 0x76, 0x38, 0x0e, 0x8d, 0xd4, 0xce, 0xd6, 0x92,
	0x6a, 0x03, 0xbd, 0x0a, 0x7f, 0xf8, 0x0d, 0x00, 0x00, 0xff, 0xff, 0x32, 0x10, 0x89, 0x21, 0x0f,
	0x01, 0x00, 0x00,
}
