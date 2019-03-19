// Code generated by protoc-gen-go. DO NOT EDIT.
// source: passurlendpoint/passurlendpoint_.proto

/*
Package passurlendpoint is a generated protocol buffer package.

It is generated from these files:
	passurlendpoint/passurlendpoint_.proto

It has these top-level messages:
	PassUrlEndpoint
*/
package passurlendpoint

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import passurlrequestfilter "github.com/21stio/go-playground/grpc/schema/passurlrequestfilter"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type PassUrlEndpoint struct {
	Filter           *passurlrequestfilter.PassUrlRequestFilter `protobuf:"bytes,1,opt,name=filter" json:"filter,omitempty"`
	Hash             *string                                    `protobuf:"bytes,2,opt,name=hash" json:"hash,omitempty"`
	XXX_unrecognized []byte                                     `json:"-"`
}

func (m *PassUrlEndpoint) Reset()                    { *m = PassUrlEndpoint{} }
func (m *PassUrlEndpoint) String() string            { return proto.CompactTextString(m) }
func (*PassUrlEndpoint) ProtoMessage()               {}
func (*PassUrlEndpoint) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *PassUrlEndpoint) GetFilter() *passurlrequestfilter.PassUrlRequestFilter {
	if m != nil {
		return m.Filter
	}
	return nil
}

func (m *PassUrlEndpoint) GetHash() string {
	if m != nil && m.Hash != nil {
		return *m.Hash
	}
	return ""
}

func init() {
	proto.RegisterType((*PassUrlEndpoint)(nil), "passurlendpoint.PassUrlEndpoint")
}

func init() { proto.RegisterFile("passurlendpoint/passurlendpoint_.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 186 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x2b, 0x48, 0x2c, 0x2e,
	0x2e, 0x2d, 0xca, 0x49, 0xcd, 0x4b, 0x29, 0xc8, 0xcf, 0xcc, 0x2b, 0xd1, 0x47, 0xe3, 0xc7, 0xeb,
	0x15, 0x14, 0xe5, 0x97, 0xe4, 0x0b, 0xf1, 0xa3, 0x89, 0x4b, 0x19, 0x40, 0x05, 0x8a, 0x52, 0x0b,
	0x4b, 0x53, 0x8b, 0x4b, 0xd2, 0x32, 0x73, 0x4a, 0x52, 0x8b, 0xf4, 0xb1, 0x09, 0x42, 0x8d, 0x50,
	0xca, 0xe4, 0xe2, 0x0f, 0x48, 0x2c, 0x2e, 0x0e, 0x2d, 0xca, 0x71, 0x85, 0x1a, 0x22, 0xe4, 0xc4,
	0xc5, 0x06, 0x51, 0x23, 0xc1, 0xa8, 0xc0, 0xa8, 0xc1, 0x6d, 0xa4, 0xa5, 0x87, 0xcd, 0x00, 0x3d,
	0xa8, 0xb6, 0x20, 0x88, 0xa0, 0x1b, 0x58, 0x30, 0x08, 0xaa, 0x53, 0x48, 0x88, 0x8b, 0x25, 0x23,
	0xb1, 0x38, 0x43, 0x82, 0x49, 0x81, 0x51, 0x83, 0x33, 0x08, 0xcc, 0x76, 0xb2, 0x8d, 0xb2, 0x4e,
	0xcf, 0x2c, 0xc9, 0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x37, 0x32, 0x2c, 0x2e, 0xc9, 0xcc,
	0xd7, 0x4f, 0xcf, 0xd7, 0x2d, 0xc8, 0x49, 0xac, 0x4c, 0x2f, 0xca, 0x2f, 0xcd, 0x4b, 0xd1, 0x4f,
	0x2f, 0x2a, 0x48, 0xd6, 0x2f, 0x4e, 0xce, 0x48, 0xcd, 0x4d, 0x44, 0xf7, 0x33, 0x20, 0x00, 0x00,
	0xff, 0xff, 0x30, 0x1a, 0x03, 0xcf, 0x15, 0x01, 0x00, 0x00,
}
