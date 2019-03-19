// Code generated by protoc-gen-go. DO NOT EDIT.
// source: verifytokenresponse/verifytokenresponse_.proto

/*
Package verifytokenresponse is a generated protocol buffer package.

It is generated from these files:
	verifytokenresponse/verifytokenresponse_.proto

It has these top-level messages:
	VerifyTokenResponse
*/
package verifytokenresponse

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import responsemeta "github.com/21stio/go-playground/grpc/schema/responsemeta"
import token "github.com/21stio/go-playground/grpc/schema/token"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type VerifyTokenResponse struct {
	Meta             *responsemeta.ResponseMeta `protobuf:"bytes,1,opt,name=meta" json:"meta,omitempty"`
	Token            *token.Token               `protobuf:"bytes,2,opt,name=token" json:"token,omitempty"`
	Hash             *string                    `protobuf:"bytes,3,opt,name=hash" json:"hash,omitempty"`
	XXX_unrecognized []byte                     `json:"-"`
}

func (m *VerifyTokenResponse) Reset()                    { *m = VerifyTokenResponse{} }
func (m *VerifyTokenResponse) String() string            { return proto.CompactTextString(m) }
func (*VerifyTokenResponse) ProtoMessage()               {}
func (*VerifyTokenResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *VerifyTokenResponse) GetMeta() *responsemeta.ResponseMeta {
	if m != nil {
		return m.Meta
	}
	return nil
}

func (m *VerifyTokenResponse) GetToken() *token.Token {
	if m != nil {
		return m.Token
	}
	return nil
}

func (m *VerifyTokenResponse) GetHash() string {
	if m != nil && m.Hash != nil {
		return *m.Hash
	}
	return ""
}

func init() {
	proto.RegisterType((*VerifyTokenResponse)(nil), "verifytokenresponse.VerifyTokenResponse")
}

func init() { proto.RegisterFile("verifytokenresponse/verifytokenresponse_.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 206 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0x2b, 0x4b, 0x2d, 0xca,
	0x4c, 0xab, 0x2c, 0xc9, 0xcf, 0x4e, 0xcd, 0x2b, 0x4a, 0x2d, 0x2e, 0xc8, 0xcf, 0x2b, 0x4e, 0xd5,
	0xc7, 0x22, 0x16, 0xaf, 0x57, 0x50, 0x94, 0x5f, 0x92, 0x2f, 0x24, 0x8c, 0x45, 0x4e, 0x4a, 0x01,
	0xc6, 0xca, 0x4d, 0x2d, 0x49, 0xd4, 0x47, 0xe6, 0x40, 0xb5, 0x49, 0x09, 0x81, 0x35, 0xe8, 0x83,
	0x49, 0xa8, 0x98, 0x52, 0x2d, 0x97, 0x70, 0x18, 0xd8, 0xb0, 0x10, 0x90, 0x68, 0x10, 0x54, 0x97,
	0x90, 0x1e, 0x17, 0x0b, 0x48, 0xa7, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0xb7, 0x91, 0x94, 0x1e, 0xb2,
	0x71, 0x7a, 0x30, 0x55, 0xbe, 0xa9, 0x25, 0x89, 0x41, 0x60, 0x75, 0x42, 0x4a, 0x5c, 0xac, 0x60,
	0x63, 0x25, 0x98, 0xc0, 0x1a, 0x78, 0xf4, 0xc0, 0x3c, 0x3d, 0x88, 0xa1, 0x10, 0x29, 0x21, 0x21,
	0x2e, 0x96, 0x8c, 0xc4, 0xe2, 0x0c, 0x09, 0x66, 0x05, 0x46, 0x0d, 0xce, 0x20, 0x30, 0xdb, 0xc9,
	0x31, 0xca, 0x3e, 0x3d, 0xb3, 0x24, 0xa3, 0x34, 0x49, 0x2f, 0x39, 0x3f, 0x57, 0xdf, 0xc8, 0xb0,
	0xb8, 0x24, 0x33, 0x5f, 0x3f, 0x3d, 0x5f, 0xb7, 0x20, 0x27, 0xb1, 0x32, 0xbd, 0x28, 0xbf, 0x34,
	0x2f, 0x45, 0x3f, 0xbd, 0xa8, 0x20, 0x59, 0xbf, 0x38, 0x39, 0x23, 0x35, 0x37, 0x11, 0x5b, 0x98,
	0x00, 0x02, 0x00, 0x00, 0xff, 0xff, 0x9d, 0x10, 0x34, 0x48, 0x3d, 0x01, 0x00, 0x00,
}
