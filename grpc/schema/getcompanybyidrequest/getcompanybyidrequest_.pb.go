// Code generated by protoc-gen-go. DO NOT EDIT.
// source: getcompanybyidrequest/getcompanybyidrequest_.proto

/*
Package getcompanybyidrequest is a generated protocol buffer package.

It is generated from these files:
	getcompanybyidrequest/getcompanybyidrequest_.proto

It has these top-level messages:
	GetCompanyByIdRequest
*/
package getcompanybyidrequest

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import requestmeta "github.com/21stio/go-playground/grpc/schema/requestmeta"
import cache "github.com/21stio/go-playground/grpc/schema/cache"
import id "github.com/21stio/go-playground/grpc/schema/id"
import getcompanybyidresponseselect "github.com/21stio/go-playground/grpc/schema/getcompanybyidresponseselect"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type GetCompanyByIdRequest struct {
	Meta             *requestmeta.RequestMeta                                   `protobuf:"bytes,1,opt,name=meta" json:"meta,omitempty"`
	Cache            *cache.Cache                                               `protobuf:"bytes,2,opt,name=cache" json:"cache,omitempty"`
	Id               *id.Id                                                     `protobuf:"bytes,3,opt,name=id" json:"id,omitempty"`
	Select           *getcompanybyidresponseselect.GetCompanyByIdResponseSelect `protobuf:"bytes,4,opt,name=select" json:"select,omitempty"`
	Hash             *string                                                    `protobuf:"bytes,5,opt,name=hash" json:"hash,omitempty"`
	XXX_unrecognized []byte                                                     `json:"-"`
}

func (m *GetCompanyByIdRequest) Reset()                    { *m = GetCompanyByIdRequest{} }
func (m *GetCompanyByIdRequest) String() string            { return proto.CompactTextString(m) }
func (*GetCompanyByIdRequest) ProtoMessage()               {}
func (*GetCompanyByIdRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *GetCompanyByIdRequest) GetMeta() *requestmeta.RequestMeta {
	if m != nil {
		return m.Meta
	}
	return nil
}

func (m *GetCompanyByIdRequest) GetCache() *cache.Cache {
	if m != nil {
		return m.Cache
	}
	return nil
}

func (m *GetCompanyByIdRequest) GetId() *id.Id {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *GetCompanyByIdRequest) GetSelect() *getcompanybyidresponseselect.GetCompanyByIdResponseSelect {
	if m != nil {
		return m.Select
	}
	return nil
}

func (m *GetCompanyByIdRequest) GetHash() string {
	if m != nil && m.Hash != nil {
		return *m.Hash
	}
	return ""
}

func init() {
	proto.RegisterType((*GetCompanyByIdRequest)(nil), "getcompanybyidrequest.GetCompanyByIdRequest")
}

func init() { proto.RegisterFile("getcompanybyidrequest/getcompanybyidrequest_.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 280 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x91, 0xb1, 0x6a, 0xc3, 0x30,
	0x10, 0x86, 0xb1, 0xeb, 0x04, 0xaa, 0x66, 0x12, 0xa4, 0x88, 0x0c, 0x25, 0x64, 0xca, 0xd0, 0x9e,
	0xa8, 0xc7, 0x4e, 0x6d, 0x3c, 0x94, 0x0c, 0x5d, 0xd4, 0xad, 0x4b, 0x51, 0xa4, 0xc3, 0x16, 0xc4,
	0x96, 0x6b, 0xc9, 0x83, 0x1f, 0xb9, 0x6f, 0x51, 0x22, 0x3b, 0x60, 0x52, 0x93, 0xe5, 0xb8, 0xbb,
	0xef, 0xd7, 0xf1, 0x81, 0x48, 0x9a, 0xa3, 0x57, 0xb6, 0xac, 0x65, 0xd5, 0x1d, 0x3a, 0xa3, 0x1b,
	0xfc, 0x69, 0xd1, 0x79, 0x3e, 0xb9, 0xfd, 0x86, 0xba, 0xb1, 0xde, 0xd2, 0xe5, 0x24, 0x5d, 0x3d,
	0x0c, 0x4d, 0x89, 0x5e, 0xf2, 0x51, 0x3f, 0x3c, 0x5b, 0x51, 0x25, 0x55, 0x81, 0x3c, 0xd4, 0xf3,
	0x6e, 0x61, 0x34, 0x37, 0xfa, 0x3c, 0xbd, 0x5e, 0x1e, 0x76, 0xb5, 0xad, 0x1c, 0x3a, 0x3c, 0xa2,
	0xfa, 0xef, 0x34, 0x86, 0xc3, 0x85, 0xcd, 0x6f, 0x44, 0x96, 0xef, 0xe8, 0xb3, 0x3e, 0xb7, 0xeb,
	0xf6, 0x5a, 0xf4, 0x22, 0xf4, 0x91, 0x24, 0x27, 0x19, 0x16, 0xad, 0xa3, 0xed, 0x5d, 0xca, 0x60,
	0x24, 0x08, 0x43, 0xe6, 0x03, 0xbd, 0x14, 0x21, 0x45, 0x37, 0x64, 0x16, 0x3c, 0x59, 0x1c, 0xe2,
	0x0b, 0x08, 0x13, 0x64, 0xa7, 0x2a, 0x7a, 0x44, 0xef, 0x49, 0x6c, 0x34, 0xbb, 0x09, 0x81, 0x39,
	0x18, 0x0d, 0x7b, 0x2d, 0x62, 0xa3, 0xa9, 0x20, 0xf3, 0x5e, 0x8a, 0x25, 0x81, 0xbd, 0xc0, 0x35,
	0x73, 0xb8, 0xd4, 0xed, 0xe1, 0x67, 0x80, 0x62, 0xb8, 0x44, 0x29, 0x49, 0x0a, 0xe9, 0x0a, 0x36,
	0x5b, 0x47, 0xdb, 0x5b, 0x11, 0xfa, 0x5d, 0xf6, 0xf5, 0x96, 0x1b, 0x5f, 0xb4, 0x07, 0x50, 0xb6,
	0xe4, 0xe9, 0xb3, 0xf3, 0xc6, 0xf2, 0xdc, 0x3e, 0xd5, 0x47, 0xd9, 0xe5, 0x8d, 0x6d, 0x2b, 0xcd,
	0xf3, 0xa6, 0x56, 0xdc, 0xa9, 0x02, 0x4b, 0x39, 0xfd, 0xa5, 0x7f, 0x01, 0x00, 0x00, 0xff, 0xff,
	0xde, 0xab, 0x3c, 0x61, 0x00, 0x02, 0x00, 0x00,
}
