// Code generated by protoc-gen-go. DO NOT EDIT.
// source: modifyproductlistresponse/modifyproductlistresponse_.proto

/*
Package modifyproductlistresponse is a generated protocol buffer package.

It is generated from these files:
	modifyproductlistresponse/modifyproductlistresponse_.proto

It has these top-level messages:
	ModifyProductListResponse
*/
package modifyproductlistresponse

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

type ModifyProductListResponse struct {
	Meta             *responsemeta.ResponseMeta `protobuf:"bytes,1,opt,name=meta" json:"meta,omitempty"`
	Hash             *string                    `protobuf:"bytes,2,opt,name=hash" json:"hash,omitempty"`
	XXX_unrecognized []byte                     `json:"-"`
}

func (m *ModifyProductListResponse) Reset()                    { *m = ModifyProductListResponse{} }
func (m *ModifyProductListResponse) String() string            { return proto.CompactTextString(m) }
func (*ModifyProductListResponse) ProtoMessage()               {}
func (*ModifyProductListResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *ModifyProductListResponse) GetMeta() *responsemeta.ResponseMeta {
	if m != nil {
		return m.Meta
	}
	return nil
}

func (m *ModifyProductListResponse) GetHash() string {
	if m != nil && m.Hash != nil {
		return *m.Hash
	}
	return ""
}

func init() {
	proto.RegisterType((*ModifyProductListResponse)(nil), "modifyproductlistresponse.ModifyProductListResponse")
}

func init() {
	proto.RegisterFile("modifyproductlistresponse/modifyproductlistresponse_.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 189 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xb2, 0xca, 0xcd, 0x4f, 0xc9,
	0x4c, 0xab, 0x2c, 0x28, 0xca, 0x4f, 0x29, 0x4d, 0x2e, 0xc9, 0xc9, 0x2c, 0x2e, 0x29, 0x4a, 0x2d,
	0x2e, 0xc8, 0xcf, 0x2b, 0x4e, 0xd5, 0xc7, 0x29, 0x13, 0xaf, 0x57, 0x50, 0x94, 0x5f, 0x92, 0x2f,
	0x24, 0x89, 0x53, 0x85, 0x94, 0x02, 0x8c, 0x95, 0x9b, 0x5a, 0x92, 0xa8, 0x8f, 0xcc, 0x81, 0x6a,
	0x56, 0x8a, 0xe7, 0x92, 0xf4, 0x05, 0x6b, 0x0f, 0x80, 0x68, 0xf7, 0xc9, 0x2c, 0x2e, 0x09, 0x82,
	0xaa, 0x13, 0xd2, 0xe3, 0x62, 0x01, 0xa9, 0x95, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x36, 0x92, 0xd2,
	0x43, 0x36, 0x40, 0x0f, 0xa6, 0xca, 0x37, 0xb5, 0x24, 0x31, 0x08, 0xac, 0x4e, 0x48, 0x88, 0x8b,
	0x25, 0x23, 0xb1, 0x38, 0x43, 0x82, 0x49, 0x81, 0x51, 0x83, 0x33, 0x08, 0xcc, 0x76, 0x72, 0x8f,
	0x72, 0x4d, 0xcf, 0x2c, 0xc9, 0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x37, 0x32, 0x2c, 0x2e,
	0xc9, 0xcc, 0xd7, 0x4f, 0xcf, 0xd7, 0x2d, 0xc8, 0x49, 0xac, 0x4c, 0x2f, 0xca, 0x2f, 0xcd, 0x4b,
	0xd1, 0x4f, 0x2f, 0x2a, 0x48, 0xd6, 0x2f, 0x4e, 0xce, 0x48, 0xcd, 0x4d, 0xc4, 0xed, 0x5b, 0x40,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x8e, 0xdc, 0x2d, 0x8b, 0x23, 0x01, 0x00, 0x00,
}
