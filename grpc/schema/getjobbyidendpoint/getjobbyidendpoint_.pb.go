// Code generated by protoc-gen-go. DO NOT EDIT.
// source: getjobbyidendpoint/getjobbyidendpoint_.proto

/*
Package getjobbyidendpoint is a generated protocol buffer package.

It is generated from these files:
	getjobbyidendpoint/getjobbyidendpoint_.proto

It has these top-level messages:
	GetJobByIdEndpoint
*/
package getjobbyidendpoint

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import getjobbyidrequestfilter "github.com/21stio/go-playground/grpc/schema/getjobbyidrequestfilter"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type GetJobByIdEndpoint struct {
	Filter           *getjobbyidrequestfilter.GetJobByIdRequestFilter `protobuf:"bytes,1,opt,name=filter" json:"filter,omitempty"`
	Hash             *string                                          `protobuf:"bytes,2,opt,name=hash" json:"hash,omitempty"`
	XXX_unrecognized []byte                                           `json:"-"`
}

func (m *GetJobByIdEndpoint) Reset()                    { *m = GetJobByIdEndpoint{} }
func (m *GetJobByIdEndpoint) String() string            { return proto.CompactTextString(m) }
func (*GetJobByIdEndpoint) ProtoMessage()               {}
func (*GetJobByIdEndpoint) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *GetJobByIdEndpoint) GetFilter() *getjobbyidrequestfilter.GetJobByIdRequestFilter {
	if m != nil {
		return m.Filter
	}
	return nil
}

func (m *GetJobByIdEndpoint) GetHash() string {
	if m != nil && m.Hash != nil {
		return *m.Hash
	}
	return ""
}

func init() {
	proto.RegisterType((*GetJobByIdEndpoint)(nil), "getjobbyidendpoint.GetJobByIdEndpoint")
}

func init() { proto.RegisterFile("getjobbyidendpoint/getjobbyidendpoint_.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 193 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0x49, 0x4f, 0x2d, 0xc9,
	0xca, 0x4f, 0x4a, 0xaa, 0xcc, 0x4c, 0x49, 0xcd, 0x4b, 0x29, 0xc8, 0xcf, 0xcc, 0x2b, 0xd1, 0xc7,
	0x14, 0x8a, 0xd7, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0xc2, 0x94, 0x92, 0x32, 0x43, 0x88,
	0x15, 0xa5, 0x16, 0x96, 0xa6, 0x16, 0x97, 0xa4, 0x65, 0xe6, 0x94, 0xa4, 0x16, 0xe9, 0xe3, 0x10,
	0x87, 0x9a, 0xa5, 0x54, 0xc4, 0x25, 0xe4, 0x9e, 0x5a, 0xe2, 0x95, 0x9f, 0xe4, 0x54, 0xe9, 0x99,
	0xe2, 0x0a, 0x35, 0x4d, 0xc8, 0x83, 0x8b, 0x0d, 0xa2, 0x4c, 0x82, 0x51, 0x81, 0x51, 0x83, 0xdb,
	0xc8, 0x40, 0x0f, 0x87, 0x31, 0x7a, 0x08, 0xcd, 0x41, 0x10, 0x71, 0x37, 0xb0, 0x78, 0x10, 0x54,
	0xbf, 0x90, 0x10, 0x17, 0x4b, 0x46, 0x62, 0x71, 0x86, 0x04, 0x93, 0x02, 0xa3, 0x06, 0x67, 0x10,
	0x98, 0xed, 0xe4, 0x10, 0x65, 0x97, 0x9e, 0x59, 0x92, 0x51, 0x9a, 0xa4, 0x97, 0x9c, 0x9f, 0xab,
	0x6f, 0x64, 0x58, 0x5c, 0x92, 0x99, 0xaf, 0x9f, 0x9e, 0xaf, 0x5b, 0x90, 0x93, 0x58, 0x99, 0x5e,
	0x94, 0x5f, 0x9a, 0x97, 0xa2, 0x9f, 0x5e, 0x54, 0x90, 0xac, 0x5f, 0x9c, 0x9c, 0x91, 0x9a, 0x9b,
	0x88, 0x25, 0x20, 0x00, 0x01, 0x00, 0x00, 0xff, 0xff, 0xc4, 0xa1, 0xc4, 0x75, 0x30, 0x01, 0x00,
	0x00,
}
