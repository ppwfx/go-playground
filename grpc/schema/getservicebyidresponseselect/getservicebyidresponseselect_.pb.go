// Code generated by protoc-gen-go. DO NOT EDIT.
// source: getservicebyidresponseselect/getservicebyidresponseselect_.proto

/*
Package getservicebyidresponseselect is a generated protocol buffer package.

It is generated from these files:
	getservicebyidresponseselect/getservicebyidresponseselect_.proto

It has these top-level messages:
	GetServiceByIdResponseSelect
*/
package getservicebyidresponseselect

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import responsemetaselect "github.com/21stio/go-playground/grpc/schema/responsemetaselect"
import serviceselect "github.com/21stio/go-playground/grpc/schema/serviceselect"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type GetServiceByIdResponseSelect struct {
	SelectAll        *bool                                  `protobuf:"varint,1,opt,name=selectAll" json:"selectAll,omitempty"`
	Meta             *responsemetaselect.ResponseMetaSelect `protobuf:"bytes,2,opt,name=meta" json:"meta,omitempty"`
	Service          *serviceselect.ServiceSelect           `protobuf:"bytes,3,opt,name=service" json:"service,omitempty"`
	SelectHash       *bool                                  `protobuf:"varint,4,opt,name=selectHash" json:"selectHash,omitempty"`
	Hash             *string                                `protobuf:"bytes,5,opt,name=hash" json:"hash,omitempty"`
	XXX_unrecognized []byte                                 `json:"-"`
}

func (m *GetServiceByIdResponseSelect) Reset()                    { *m = GetServiceByIdResponseSelect{} }
func (m *GetServiceByIdResponseSelect) String() string            { return proto.CompactTextString(m) }
func (*GetServiceByIdResponseSelect) ProtoMessage()               {}
func (*GetServiceByIdResponseSelect) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *GetServiceByIdResponseSelect) GetSelectAll() bool {
	if m != nil && m.SelectAll != nil {
		return *m.SelectAll
	}
	return false
}

func (m *GetServiceByIdResponseSelect) GetMeta() *responsemetaselect.ResponseMetaSelect {
	if m != nil {
		return m.Meta
	}
	return nil
}

func (m *GetServiceByIdResponseSelect) GetService() *serviceselect.ServiceSelect {
	if m != nil {
		return m.Service
	}
	return nil
}

func (m *GetServiceByIdResponseSelect) GetSelectHash() bool {
	if m != nil && m.SelectHash != nil {
		return *m.SelectHash
	}
	return false
}

func (m *GetServiceByIdResponseSelect) GetHash() string {
	if m != nil && m.Hash != nil {
		return *m.Hash
	}
	return ""
}

func init() {
	proto.RegisterType((*GetServiceByIdResponseSelect)(nil), "getservicebyidresponseselect.GetServiceByIdResponseSelect")
}

func init() {
	proto.RegisterFile("getservicebyidresponseselect/getservicebyidresponseselect_.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 261 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x91, 0x41, 0x4b, 0xc3, 0x30,
	0x14, 0xc7, 0x89, 0x56, 0x74, 0xf1, 0x96, 0x53, 0x19, 0x45, 0xca, 0x0e, 0xd2, 0x83, 0x26, 0xb8,
	0x83, 0x07, 0x4f, 0xba, 0x8b, 0x53, 0xf0, 0xd2, 0xdd, 0xbc, 0x48, 0x96, 0x3e, 0xd2, 0x42, 0xbb,
	0x94, 0x24, 0x13, 0xfa, 0x7d, 0xfd, 0x20, 0xb2, 0xe4, 0x15, 0x1d, 0x83, 0xde, 0xfa, 0xde, 0xfb,
	0xfd, 0x7f, 0xfd, 0x43, 0xe8, 0xb3, 0x06, 0xef, 0xc0, 0x7e, 0x37, 0x0a, 0xb6, 0x43, 0x53, 0x59,
	0x70, 0xbd, 0xd9, 0x39, 0x70, 0xd0, 0x82, 0xf2, 0x62, 0xea, 0xf8, 0xc5, 0x7b, 0x6b, 0xbc, 0x61,
	0xd9, 0x14, 0x34, 0xbf, 0x1b, 0xe7, 0x0e, 0xbc, 0x44, 0xeb, 0xe9, 0x0a, 0x5d, 0xf3, 0x05, 0x8a,
	0x10, 0x3c, 0x9a, 0x90, 0x59, 0xfc, 0x10, 0x9a, 0xbd, 0x82, 0xdf, 0xc4, 0xdb, 0x6a, 0x78, 0xab,
	0x4a, 0xf4, 0x6d, 0x02, 0xc7, 0x32, 0x3a, 0x8b, 0x89, 0x97, 0xb6, 0x4d, 0x49, 0x4e, 0x8a, 0xab,
	0xf2, 0x6f, 0xc1, 0x9e, 0x68, 0x72, 0xf8, 0x6f, 0x7a, 0x96, 0x93, 0xe2, 0x7a, 0x79, 0xcb, 0x4f,
	0xcb, 0xf0, 0xd1, 0xf7, 0x01, 0x5e, 0x46, 0x67, 0x19, 0x32, 0xec, 0x91, 0x5e, 0x62, 0xa5, 0xf4,
	0x3c, 0xc4, 0x33, 0x7e, 0x54, 0x91, 0x63, 0x29, 0x0c, 0x8d, 0x30, 0xbb, 0xa1, 0x34, 0x02, 0x6b,
	0xe9, 0xea, 0x34, 0x09, 0x95, 0xfe, 0x6d, 0x18, 0xa3, 0x49, 0x7d, 0xb8, 0x5c, 0xe4, 0xa4, 0x98,
	0x95, 0xe1, 0x7b, 0xf5, 0xfe, 0xb9, 0xd6, 0x8d, 0xaf, 0xf7, 0x5b, 0xae, 0x4c, 0x27, 0x96, 0x0f,
	0xce, 0x37, 0x46, 0x68, 0x73, 0xdf, 0xb7, 0x72, 0xd0, 0xd6, 0xec, 0x77, 0x95, 0xd0, 0xb6, 0x57,
	0xc2, 0xa9, 0x1a, 0x3a, 0x39, 0xf9, 0x52, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x70, 0xa6, 0xa2,
	0xbe, 0xe5, 0x01, 0x00, 0x00,
}
