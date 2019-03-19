// Code generated by protoc-gen-go. DO NOT EDIT.
// source: updatewhateversrequest/updatewhateversrequest_.proto

/*
Package updatewhateversrequest is a generated protocol buffer package.

It is generated from these files:
	updatewhateversrequest/updatewhateversrequest_.proto

It has these top-level messages:
	UpdateWhateversRequest
*/
package updatewhateversrequest

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import requestmeta "github.com/21stio/go-playground/grpc/schema/requestmeta"
import servicefilter "github.com/21stio/go-playground/grpc/schema/servicefilter"
import whatever "github.com/21stio/go-playground/grpc/schema/whatever"
import updatewhateversresponseselect "github.com/21stio/go-playground/grpc/schema/updatewhateversresponseselect"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type UpdateWhateversRequest struct {
	Meta             *requestmeta.RequestMeta                                     `protobuf:"bytes,1,opt,name=meta" json:"meta,omitempty"`
	ServiceFilter    *servicefilter.ServiceFilter                                 `protobuf:"bytes,2,opt,name=serviceFilter" json:"serviceFilter,omitempty"`
	Whatevers        []*whatever.Whatever                                         `protobuf:"bytes,3,rep,name=whatevers" json:"whatevers,omitempty"`
	Select           *updatewhateversresponseselect.UpdateWhateversResponseSelect `protobuf:"bytes,4,opt,name=select" json:"select,omitempty"`
	Hash             *string                                                      `protobuf:"bytes,5,opt,name=hash" json:"hash,omitempty"`
	XXX_unrecognized []byte                                                       `json:"-"`
}

func (m *UpdateWhateversRequest) Reset()                    { *m = UpdateWhateversRequest{} }
func (m *UpdateWhateversRequest) String() string            { return proto.CompactTextString(m) }
func (*UpdateWhateversRequest) ProtoMessage()               {}
func (*UpdateWhateversRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *UpdateWhateversRequest) GetMeta() *requestmeta.RequestMeta {
	if m != nil {
		return m.Meta
	}
	return nil
}

func (m *UpdateWhateversRequest) GetServiceFilter() *servicefilter.ServiceFilter {
	if m != nil {
		return m.ServiceFilter
	}
	return nil
}

func (m *UpdateWhateversRequest) GetWhatevers() []*whatever.Whatever {
	if m != nil {
		return m.Whatevers
	}
	return nil
}

func (m *UpdateWhateversRequest) GetSelect() *updatewhateversresponseselect.UpdateWhateversResponseSelect {
	if m != nil {
		return m.Select
	}
	return nil
}

func (m *UpdateWhateversRequest) GetHash() string {
	if m != nil && m.Hash != nil {
		return *m.Hash
	}
	return ""
}

func init() {
	proto.RegisterType((*UpdateWhateversRequest)(nil), "updatewhateversrequest.UpdateWhateversRequest")
}

func init() {
	proto.RegisterFile("updatewhateversrequest/updatewhateversrequest_.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 295 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0xc1, 0x4b, 0xc3, 0x30,
	0x14, 0xc6, 0xe9, 0x56, 0x85, 0x65, 0x78, 0xc9, 0x61, 0x84, 0x21, 0x52, 0x76, 0xea, 0x41, 0x13,
	0x2d, 0x1e, 0x3d, 0x15, 0xf1, 0xe6, 0x25, 0x53, 0x04, 0x2f, 0x12, 0xbb, 0x67, 0x5b, 0x68, 0x97,
	0x9a, 0xa4, 0x13, 0xff, 0x21, 0xff, 0x4e, 0x31, 0x4d, 0xb4, 0x65, 0x75, 0xb7, 0xaf, 0xf9, 0xbe,
	0xf7, 0xde, 0xef, 0x83, 0xa2, 0xeb, 0xb6, 0xd9, 0x08, 0x03, 0x1f, 0x85, 0x30, 0xb0, 0x03, 0xa5,
	0x15, 0xbc, 0xb7, 0xa0, 0x0d, 0x1b, 0x7f, 0x7e, 0xa1, 0x8d, 0x92, 0x46, 0xe2, 0xc5, 0xb8, 0xbd,
	0x3c, 0x73, 0xa2, 0x06, 0x23, 0x58, 0x4f, 0xbb, 0xb9, 0xe5, 0x4a, 0x83, 0xda, 0x95, 0x19, 0xbc,
	0x95, 0x95, 0x01, 0xc5, 0x06, 0x5f, 0x3e, 0x43, 0xfc, 0x56, 0xe6, 0x85, 0x77, 0xd2, 0xbd, 0xab,
	0xba, 0x91, 0x5b, 0x0d, 0x1a, 0x2a, 0xc8, 0x46, 0x90, 0xfb, 0xae, 0xdb, 0xb1, 0xfa, 0x9a, 0xa0,
	0xc5, 0xa3, 0x0d, 0x3e, 0xf9, 0x20, 0xef, 0x38, 0xf1, 0x39, 0x0a, 0x7f, 0x58, 0x49, 0x10, 0x05,
	0xf1, 0x3c, 0x21, 0xb4, 0xc7, 0x4f, 0x5d, 0xe6, 0x1e, 0x8c, 0xe0, 0x36, 0x85, 0x53, 0x74, 0xe2,
	0xf0, 0xef, 0x2c, 0x3e, 0x99, 0xd8, 0xb1, 0x53, 0x3a, 0x28, 0x45, 0xd7, 0xfd, 0x0c, 0x1f, 0x8e,
	0xe0, 0x4b, 0x34, 0xfb, 0xc5, 0x25, 0xd3, 0x68, 0x1a, 0xcf, 0x13, 0x4c, 0xfd, 0x0b, 0xf5, 0x80,
	0xfc, 0x2f, 0x84, 0x1f, 0xd0, 0x71, 0xd7, 0x87, 0x84, 0xf6, 0xdc, 0x0d, 0x3d, 0xd8, 0x9a, 0xee,
	0x55, 0xed, 0xdc, 0xb5, 0x75, 0xb9, 0xdb, 0x85, 0x31, 0x0a, 0x0b, 0xa1, 0x0b, 0x72, 0x14, 0x05,
	0xf1, 0x8c, 0x5b, 0x9d, 0xde, 0x3e, 0xa7, 0x79, 0x69, 0x8a, 0xf6, 0x95, 0x66, 0xb2, 0x66, 0xc9,
	0x95, 0x36, 0xa5, 0x64, 0xb9, 0xbc, 0x68, 0x2a, 0xf1, 0x99, 0x2b, 0xd9, 0x6e, 0x37, 0x2c, 0x57,
	0x4d, 0xc6, 0x74, 0x56, 0x40, 0x2d, 0xfe, 0xf9, 0x5f, 0xbe, 0x03, 0x00, 0x00, 0xff, 0xff, 0x90,
	0x7f, 0x7c, 0x65, 0x5f, 0x02, 0x00, 0x00,
}
