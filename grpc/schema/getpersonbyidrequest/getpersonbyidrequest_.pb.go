// Code generated by protoc-gen-go. DO NOT EDIT.
// source: getpersonbyidrequest/getpersonbyidrequest_.proto

/*
Package getpersonbyidrequest is a generated protocol buffer package.

It is generated from these files:
	getpersonbyidrequest/getpersonbyidrequest_.proto

It has these top-level messages:
	GetPersonByIdRequest
*/
package getpersonbyidrequest

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import requestmeta "github.com/21stio/go-playground/grpc/schema/requestmeta"
import cache "github.com/21stio/go-playground/grpc/schema/cache"
import id "github.com/21stio/go-playground/grpc/schema/id"
import getpersonbyidresponseselect "github.com/21stio/go-playground/grpc/schema/getpersonbyidresponseselect"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type GetPersonByIdRequest struct {
	Meta             *requestmeta.RequestMeta                                 `protobuf:"bytes,1,opt,name=meta" json:"meta,omitempty"`
	Cache            *cache.Cache                                             `protobuf:"bytes,2,opt,name=cache" json:"cache,omitempty"`
	Id               *id.Id                                                   `protobuf:"bytes,3,opt,name=id" json:"id,omitempty"`
	Select           *getpersonbyidresponseselect.GetPersonByIdResponseSelect `protobuf:"bytes,4,opt,name=select" json:"select,omitempty"`
	Hash             *string                                                  `protobuf:"bytes,5,opt,name=hash" json:"hash,omitempty"`
	XXX_unrecognized []byte                                                   `json:"-"`
}

func (m *GetPersonByIdRequest) Reset()                    { *m = GetPersonByIdRequest{} }
func (m *GetPersonByIdRequest) String() string            { return proto.CompactTextString(m) }
func (*GetPersonByIdRequest) ProtoMessage()               {}
func (*GetPersonByIdRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *GetPersonByIdRequest) GetMeta() *requestmeta.RequestMeta {
	if m != nil {
		return m.Meta
	}
	return nil
}

func (m *GetPersonByIdRequest) GetCache() *cache.Cache {
	if m != nil {
		return m.Cache
	}
	return nil
}

func (m *GetPersonByIdRequest) GetId() *id.Id {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *GetPersonByIdRequest) GetSelect() *getpersonbyidresponseselect.GetPersonByIdResponseSelect {
	if m != nil {
		return m.Select
	}
	return nil
}

func (m *GetPersonByIdRequest) GetHash() string {
	if m != nil && m.Hash != nil {
		return *m.Hash
	}
	return ""
}

func init() {
	proto.RegisterType((*GetPersonByIdRequest)(nil), "getpersonbyidrequest.GetPersonByIdRequest")
}

func init() { proto.RegisterFile("getpersonbyidrequest/getpersonbyidrequest_.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 281 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x91, 0xb1, 0x6a, 0xc3, 0x30,
	0x10, 0x86, 0xb1, 0xeb, 0x04, 0xaa, 0x66, 0x12, 0xa1, 0x88, 0x0c, 0x25, 0x64, 0xca, 0xd0, 0x4a,
	0x6d, 0xa6, 0x4e, 0xa5, 0xb8, 0x43, 0xc9, 0x50, 0x08, 0xea, 0xd6, 0xa5, 0x28, 0xd2, 0x61, 0x0b,
	0x62, 0xcb, 0xb5, 0xe4, 0xc1, 0x2f, 0xdc, 0xe7, 0x28, 0x39, 0x3b, 0x60, 0x52, 0x93, 0x45, 0xdc,
	0xdd, 0xf7, 0xdf, 0xf1, 0x81, 0xc8, 0x63, 0x06, 0xa1, 0x82, 0xda, 0xbb, 0x72, 0xdf, 0x5a, 0x53,
	0xc3, 0x4f, 0x03, 0x3e, 0x88, 0xb1, 0xe1, 0x37, 0xaf, 0x6a, 0x17, 0x1c, 0x9d, 0x8f, 0xc1, 0xc5,
	0x5d, 0x5f, 0x14, 0x10, 0x94, 0x18, 0xd4, 0xfd, 0xd6, 0x82, 0x6a, 0xa5, 0x73, 0x10, 0xf8, 0x9e,
	0x66, 0x33, 0x6b, 0x84, 0x35, 0xa7, 0xee, 0xe5, 0xec, 0xae, 0xaf, 0x5c, 0xe9, 0xc1, 0xc3, 0x01,
	0xf4, 0x3f, 0xa1, 0x21, 0xeb, 0xf7, 0x57, 0xbf, 0x11, 0x99, 0xbf, 0x43, 0xd8, 0x61, 0x2c, 0x6d,
	0xb7, 0x46, 0x76, 0x16, 0xf4, 0x9e, 0x24, 0x47, 0x13, 0x16, 0x2d, 0xa3, 0xf5, 0xcd, 0x86, 0xf1,
	0x81, 0x1d, 0xef, 0x33, 0x1f, 0x10, 0x94, 0xc4, 0x14, 0x5d, 0x91, 0x09, 0x4a, 0xb2, 0x18, 0xe3,
	0x33, 0x8e, 0x1d, 0x7f, 0x3b, 0xbe, 0xb2, 0x43, 0xf4, 0x96, 0xc4, 0xd6, 0xb0, 0x2b, 0x0c, 0x4c,
	0xb9, 0x35, 0x7c, 0x6b, 0x64, 0x6c, 0x0d, 0xdd, 0x91, 0x69, 0xe7, 0xc4, 0x12, 0x64, 0xcf, 0xfc,
	0x82, 0x37, 0x3f, 0x93, 0xed, 0xd8, 0x27, 0x32, 0xd9, 0xdf, 0xa1, 0x94, 0x24, 0xb9, 0xf2, 0x39,
	0x9b, 0x2c, 0xa3, 0xf5, 0xb5, 0xc4, 0x3a, 0x4d, 0xbf, 0x5e, 0x33, 0x1b, 0xf2, 0x66, 0xcf, 0xb5,
	0x2b, 0xc4, 0xe6, 0xc9, 0x07, 0xeb, 0x44, 0xe6, 0x1e, 0xaa, 0x83, 0x6a, 0xb3, 0xda, 0x35, 0xa5,
	0x11, 0x59, 0x5d, 0x69, 0xe1, 0x75, 0x0e, 0x85, 0x1a, 0xfd, 0xcb, 0xbf, 0x00, 0x00, 0x00, 0xff,
	0xff, 0xf7, 0xd2, 0x78, 0x8c, 0xf7, 0x01, 0x00, 0x00,
}
