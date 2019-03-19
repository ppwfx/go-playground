// Code generated by protoc-gen-go. DO NOT EDIT.
// source: passupdatewhateversrequestendpoint/passupdatewhateversrequestendpoint_.proto

/*
Package passupdatewhateversrequestendpoint is a generated protocol buffer package.

It is generated from these files:
	passupdatewhateversrequestendpoint/passupdatewhateversrequestendpoint_.proto

It has these top-level messages:
	PassUpdateWhateversRequestEndpoint
*/
package passupdatewhateversrequestendpoint

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import passupdatewhateversrequestrequestfilter "github.com/21stio/go-playground/grpc/schema/passupdatewhateversrequestrequestfilter"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type PassUpdateWhateversRequestEndpoint struct {
	Filter           *passupdatewhateversrequestrequestfilter.PassUpdateWhateversRequestRequestFilter `protobuf:"bytes,1,opt,name=filter" json:"filter,omitempty"`
	Hash             *string                                                                          `protobuf:"bytes,2,opt,name=hash" json:"hash,omitempty"`
	XXX_unrecognized []byte                                                                           `json:"-"`
}

func (m *PassUpdateWhateversRequestEndpoint) Reset()         { *m = PassUpdateWhateversRequestEndpoint{} }
func (m *PassUpdateWhateversRequestEndpoint) String() string { return proto.CompactTextString(m) }
func (*PassUpdateWhateversRequestEndpoint) ProtoMessage()    {}
func (*PassUpdateWhateversRequestEndpoint) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{0}
}

func (m *PassUpdateWhateversRequestEndpoint) GetFilter() *passupdatewhateversrequestrequestfilter.PassUpdateWhateversRequestRequestFilter {
	if m != nil {
		return m.Filter
	}
	return nil
}

func (m *PassUpdateWhateversRequestEndpoint) GetHash() string {
	if m != nil && m.Hash != nil {
		return *m.Hash
	}
	return ""
}

func init() {
	proto.RegisterType((*PassUpdateWhateversRequestEndpoint)(nil), "passupdatewhateversrequestendpoint.PassUpdateWhateversRequestEndpoint")
}

func init() {
	proto.RegisterFile("passupdatewhateversrequestendpoint/passupdatewhateversrequestendpoint_.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 207 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x90, 0xb1, 0x4a, 0xc6, 0x30,
	0x14, 0x85, 0x89, 0x88, 0x60, 0xdc, 0x32, 0x15, 0xa7, 0xd2, 0xa9, 0x8b, 0x09, 0xf6, 0x11, 0x04,
	0x9d, 0x44, 0x4a, 0x41, 0x05, 0x17, 0x89, 0xed, 0x35, 0x09, 0xb4, 0x4d, 0xcc, 0xbd, 0x51, 0x7c,
	0x1d, 0x9f, 0x54, 0x68, 0xf3, 0x8f, 0xa5, 0x9d, 0xce, 0x1d, 0xbe, 0x73, 0x3e, 0xb8, 0xfc, 0x31,
	0x68, 0xc4, 0x14, 0x06, 0x4d, 0xf0, 0x63, 0x35, 0xc1, 0x37, 0x44, 0x8c, 0xf0, 0x95, 0x00, 0x09,
	0xe6, 0x21, 0x78, 0x37, 0x93, 0xda, 0x47, 0xde, 0x65, 0x88, 0x9e, 0xbc, 0xa8, 0xf6, 0xd1, 0xeb,
	0x97, 0x6d, 0x26, 0xc7, 0xa7, 0x1b, 0x09, 0xa2, 0x3a, 0xc8, 0x65, 0x77, 0xf5, 0xc7, 0x78, 0xd5,
	0x6a, 0xc4, 0xe7, 0xa5, 0xf2, 0x7a, 0xaa, 0x74, 0x2b, 0x7b, 0x9f, 0xf5, 0xc2, 0xf2, 0x8b, 0xb5,
	0x57, 0xb0, 0x92, 0xd5, 0x57, 0x4d, 0x2b, 0x0f, 0x7a, 0xe4, 0xf6, 0x78, 0x8e, 0x87, 0x85, 0xeb,
	0xf2, 0xbe, 0x10, 0xfc, 0xdc, 0x6a, 0xb4, 0xc5, 0x59, 0xc9, 0xea, 0xcb, 0x6e, 0xb9, 0xef, 0xda,
	0xb7, 0x27, 0xe3, 0xc8, 0xa6, 0x0f, 0xd9, 0xfb, 0x49, 0x35, 0xb7, 0x48, 0xce, 0x2b, 0xe3, 0x6f,
	0xc2, 0xa8, 0x7f, 0x4d, 0xf4, 0x69, 0x1e, 0x94, 0x89, 0xa1, 0x57, 0xd8, 0x5b, 0x98, 0xf4, 0x81,
	0xcf, 0xff, 0x07, 0x00, 0x00, 0xff, 0xff, 0xea, 0x0b, 0x83, 0x6c, 0xc1, 0x01, 0x00, 0x00,
}
