// Code generated by protoc-gen-go. DO NOT EDIT.
// source: deleteservicesresponse/deleteservicesresponse_.proto

/*
Package deleteservicesresponse is a generated protocol buffer package.

It is generated from these files:
	deleteservicesresponse/deleteservicesresponse_.proto

It has these top-level messages:
	DeleteServicesResponse
*/
package deleteservicesresponse

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

type DeleteServicesResponse struct {
	Meta             *responsemeta.ResponseMeta `protobuf:"bytes,1,opt,name=meta" json:"meta,omitempty"`
	Hash             *string                    `protobuf:"bytes,2,opt,name=hash" json:"hash,omitempty"`
	XXX_unrecognized []byte                     `json:"-"`
}

func (m *DeleteServicesResponse) Reset()                    { *m = DeleteServicesResponse{} }
func (m *DeleteServicesResponse) String() string            { return proto.CompactTextString(m) }
func (*DeleteServicesResponse) ProtoMessage()               {}
func (*DeleteServicesResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *DeleteServicesResponse) GetMeta() *responsemeta.ResponseMeta {
	if m != nil {
		return m.Meta
	}
	return nil
}

func (m *DeleteServicesResponse) GetHash() string {
	if m != nil && m.Hash != nil {
		return *m.Hash
	}
	return ""
}

func init() {
	proto.RegisterType((*DeleteServicesResponse)(nil), "deleteservicesresponse.DeleteServicesResponse")
}

func init() {
	proto.RegisterFile("deleteservicesresponse/deleteservicesresponse_.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 181 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x32, 0x49, 0x49, 0xcd, 0x49,
	0x2d, 0x49, 0x2d, 0x4e, 0x2d, 0x2a, 0xcb, 0x4c, 0x4e, 0x2d, 0x2e, 0x4a, 0x2d, 0x2e, 0xc8, 0xcf,
	0x2b, 0x4e, 0xd5, 0xc7, 0x2e, 0x1c, 0xaf, 0x57, 0x50, 0x94, 0x5f, 0x92, 0x2f, 0x24, 0x86, 0x5d,
	0x5a, 0x4a, 0x01, 0xc6, 0xca, 0x4d, 0x2d, 0x49, 0xd4, 0x47, 0xe6, 0x40, 0x75, 0x2a, 0xc5, 0x70,
	0x89, 0xb9, 0x80, 0xf5, 0x06, 0x43, 0xf5, 0x06, 0x41, 0x15, 0x09, 0xe9, 0x71, 0xb1, 0x80, 0x14,
	0x4a, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x1b, 0x49, 0xe9, 0x21, 0xeb, 0xd6, 0x83, 0xa9, 0xf2, 0x4d,
	0x2d, 0x49, 0x0c, 0x02, 0xab, 0x13, 0x12, 0xe2, 0x62, 0xc9, 0x48, 0x2c, 0xce, 0x90, 0x60, 0x52,
	0x60, 0xd4, 0xe0, 0x0c, 0x02, 0xb3, 0x9d, 0x5c, 0xa2, 0x9c, 0xd2, 0x33, 0x4b, 0x32, 0x4a, 0x93,
	0xf4, 0x92, 0xf3, 0x73, 0xf5, 0x8d, 0x0c, 0x8b, 0x4b, 0x32, 0xf3, 0xf5, 0xd3, 0xf3, 0x75, 0x0b,
	0x72, 0x12, 0x2b, 0xd3, 0x8b, 0xf2, 0x4b, 0xf3, 0x52, 0xf4, 0xd3, 0x8b, 0x0a, 0x92, 0xf5, 0x8b,
	0x93, 0x33, 0x52, 0x73, 0x13, 0x71, 0x78, 0x12, 0x10, 0x00, 0x00, 0xff, 0xff, 0x1d, 0x10, 0x11,
	0x08, 0x14, 0x01, 0x00, 0x00,
}
