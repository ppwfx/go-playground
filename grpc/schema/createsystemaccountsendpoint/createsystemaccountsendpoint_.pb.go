// Code generated by protoc-gen-go. DO NOT EDIT.
// source: createsystemaccountsendpoint/createsystemaccountsendpoint_.proto

/*
Package createsystemaccountsendpoint is a generated protocol buffer package.

It is generated from these files:
	createsystemaccountsendpoint/createsystemaccountsendpoint_.proto

It has these top-level messages:
	CreateSystemAccountsEndpoint
*/
package createsystemaccountsendpoint

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import createsystemaccountsrequestfilter "github.com/21stio/go-playground/grpc/schema/createsystemaccountsrequestfilter"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type CreateSystemAccountsEndpoint struct {
	Filter           *createsystemaccountsrequestfilter.CreateSystemAccountsRequestFilter `protobuf:"bytes,1,opt,name=filter" json:"filter,omitempty"`
	Hash             *string                                                              `protobuf:"bytes,2,opt,name=hash" json:"hash,omitempty"`
	XXX_unrecognized []byte                                                               `json:"-"`
}

func (m *CreateSystemAccountsEndpoint) Reset()                    { *m = CreateSystemAccountsEndpoint{} }
func (m *CreateSystemAccountsEndpoint) String() string            { return proto.CompactTextString(m) }
func (*CreateSystemAccountsEndpoint) ProtoMessage()               {}
func (*CreateSystemAccountsEndpoint) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *CreateSystemAccountsEndpoint) GetFilter() *createsystemaccountsrequestfilter.CreateSystemAccountsRequestFilter {
	if m != nil {
		return m.Filter
	}
	return nil
}

func (m *CreateSystemAccountsEndpoint) GetHash() string {
	if m != nil && m.Hash != nil {
		return *m.Hash
	}
	return ""
}

func init() {
	proto.RegisterType((*CreateSystemAccountsEndpoint)(nil), "createsystemaccountsendpoint.CreateSystemAccountsEndpoint")
}

func init() {
	proto.RegisterFile("createsystemaccountsendpoint/createsystemaccountsendpoint_.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 204 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x72, 0x48, 0x2e, 0x4a, 0x4d,
	0x2c, 0x49, 0x2d, 0xae, 0x2c, 0x2e, 0x49, 0xcd, 0x4d, 0x4c, 0x4e, 0xce, 0x2f, 0xcd, 0x2b, 0x29,
	0x4e, 0xcd, 0x4b, 0x29, 0xc8, 0xcf, 0xcc, 0x2b, 0xd1, 0xc7, 0x27, 0x19, 0xaf, 0x57, 0x50, 0x94,
	0x5f, 0x92, 0x2f, 0x24, 0x83, 0x4f, 0x91, 0x94, 0x17, 0x36, 0xd9, 0xa2, 0xd4, 0xc2, 0xd2, 0xd4,
	0xe2, 0x92, 0xb4, 0xcc, 0x9c, 0x92, 0xd4, 0x22, 0x7d, 0x82, 0x2a, 0xa0, 0x36, 0x29, 0x4d, 0x60,
	0xe4, 0x92, 0x71, 0x06, 0x2b, 0x0e, 0x06, 0x2b, 0x76, 0x84, 0x2a, 0x76, 0x85, 0x5a, 0x26, 0x14,
	0xc3, 0xc5, 0x06, 0xd1, 0x21, 0xc1, 0xa8, 0xc0, 0xa8, 0xc1, 0x6d, 0xe4, 0xa2, 0x47, 0xd0, 0x6c,
	0x3d, 0x6c, 0x06, 0x06, 0x41, 0x54, 0xb8, 0x81, 0x55, 0x04, 0x41, 0xcd, 0x14, 0x12, 0xe2, 0x62,
	0xc9, 0x48, 0x2c, 0xce, 0x90, 0x60, 0x52, 0x60, 0xd4, 0xe0, 0x0c, 0x02, 0xb3, 0x9d, 0xbc, 0xa2,
	0x3c, 0xd2, 0x33, 0x4b, 0x32, 0x4a, 0x93, 0xf4, 0x92, 0xf3, 0x73, 0xf5, 0x8d, 0x0c, 0x8b, 0x4b,
	0x32, 0xf3, 0xf5, 0xd3, 0xf3, 0x75, 0x0b, 0x72, 0x12, 0x2b, 0xd3, 0x8b, 0xf2, 0x4b, 0xf3, 0x52,
	0xf4, 0xd3, 0x8b, 0x0a, 0x92, 0xf5, 0x8b, 0x93, 0x33, 0x52, 0x73, 0x13, 0xf1, 0x86, 0x27, 0x20,
	0x00, 0x00, 0xff, 0xff, 0x9f, 0xd9, 0xc6, 0xf0, 0x8b, 0x01, 0x00, 0x00,
}
