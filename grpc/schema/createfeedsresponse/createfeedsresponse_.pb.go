// Code generated by protoc-gen-go. DO NOT EDIT.
// source: createfeedsresponse/createfeedsresponse_.proto

/*
Package createfeedsresponse is a generated protocol buffer package.

It is generated from these files:
	createfeedsresponse/createfeedsresponse_.proto

It has these top-level messages:
	CreateFeedsResponse
*/
package createfeedsresponse

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

type CreateFeedsResponse struct {
	Meta             *responsemeta.ResponseMeta `protobuf:"bytes,1,opt,name=meta" json:"meta,omitempty"`
	Hash             *string                    `protobuf:"bytes,2,opt,name=hash" json:"hash,omitempty"`
	XXX_unrecognized []byte                     `json:"-"`
}

func (m *CreateFeedsResponse) Reset()                    { *m = CreateFeedsResponse{} }
func (m *CreateFeedsResponse) String() string            { return proto.CompactTextString(m) }
func (*CreateFeedsResponse) ProtoMessage()               {}
func (*CreateFeedsResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *CreateFeedsResponse) GetMeta() *responsemeta.ResponseMeta {
	if m != nil {
		return m.Meta
	}
	return nil
}

func (m *CreateFeedsResponse) GetHash() string {
	if m != nil && m.Hash != nil {
		return *m.Hash
	}
	return ""
}

func init() {
	proto.RegisterType((*CreateFeedsResponse)(nil), "createfeedsresponse.CreateFeedsResponse")
}

func init() { proto.RegisterFile("createfeedsresponse/createfeedsresponse_.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 178 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x8f, 0xb1, 0x0e, 0x82, 0x30,
	0x10, 0x40, 0x83, 0x61, 0xb1, 0x6e, 0x65, 0x21, 0x4c, 0xc4, 0x89, 0xc5, 0x36, 0xf2, 0x03, 0x46,
	0x4d, 0xdc, 0x5c, 0xd8, 0x74, 0x31, 0xb5, 0x9c, 0x2d, 0x89, 0x70, 0x4d, 0x7b, 0x0c, 0xfe, 0xbd,
	0x81, 0x40, 0xc2, 0xc0, 0xf6, 0xee, 0xee, 0xbd, 0xe1, 0x98, 0xd0, 0x1e, 0x14, 0xc1, 0x07, 0xa0,
	0x0e, 0x1e, 0x82, 0xc3, 0x2e, 0x80, 0x5c, 0xd9, 0xbd, 0x84, 0xf3, 0x48, 0xc8, 0x93, 0x95, 0x5b,
	0x96, 0xcf, 0xd4, 0x02, 0x29, 0xb9, 0x1c, 0xa6, 0x6c, 0xff, 0x60, 0xc9, 0x75, 0x0c, 0x6f, 0x43,
	0x58, 0x4d, 0x06, 0x17, 0x2c, 0x1e, 0xac, 0x34, 0xca, 0xa3, 0x62, 0x57, 0x66, 0x62, 0x99, 0x8a,
	0xd9, 0xba, 0x03, 0xa9, 0x6a, 0xf4, 0x38, 0x67, 0xb1, 0x55, 0xc1, 0xa6, 0x9b, 0x3c, 0x2a, 0xb6,
	0xd5, 0xc8, 0x97, 0xf3, 0xf3, 0x64, 0x1a, 0xb2, 0xfd, 0x5b, 0x68, 0x6c, 0x65, 0x79, 0x0c, 0xd4,
	0xa0, 0x34, 0x78, 0x70, 0x5f, 0xf5, 0x33, 0x1e, 0xfb, 0xae, 0x96, 0xc6, 0x3b, 0x2d, 0x83, 0xb6,
	0xd0, 0xaa, 0xb5, 0xdf, 0xfe, 0x01, 0x00, 0x00, 0xff, 0xff, 0xa5, 0xaa, 0x51, 0x46, 0x05, 0x01,
	0x00, 0x00,
}
