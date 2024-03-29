// Code generated by protoc-gen-go. DO NOT EDIT.
// source: getservicesresponseselect/getservicesresponseselect_.proto

/*
Package getservicesresponseselect is a generated protocol buffer package.

It is generated from these files:
	getservicesresponseselect/getservicesresponseselect_.proto

It has these top-level messages:
	GetServicesResponseSelect
*/
package getservicesresponseselect

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import responsemetaselect "github.com/21stio/go-playground/grpc/schema/responsemetaselect"
import servicesselect "github.com/21stio/go-playground/grpc/schema/servicesselect"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type GetServicesResponseSelect struct {
	SelectAll        *bool                                  `protobuf:"varint,1,opt,name=selectAll" json:"selectAll,omitempty"`
	Meta             *responsemetaselect.ResponseMetaSelect `protobuf:"bytes,2,opt,name=meta" json:"meta,omitempty"`
	TotalCount       *bool                                  `protobuf:"varint,3,opt,name=totalCount" json:"totalCount,omitempty"`
	Services         *servicesselect.ServicesSelect         `protobuf:"bytes,4,opt,name=services" json:"services,omitempty"`
	SelectHash       *bool                                  `protobuf:"varint,5,opt,name=selectHash" json:"selectHash,omitempty"`
	Hash             *string                                `protobuf:"bytes,6,opt,name=hash" json:"hash,omitempty"`
	XXX_unrecognized []byte                                 `json:"-"`
}

func (m *GetServicesResponseSelect) Reset()                    { *m = GetServicesResponseSelect{} }
func (m *GetServicesResponseSelect) String() string            { return proto.CompactTextString(m) }
func (*GetServicesResponseSelect) ProtoMessage()               {}
func (*GetServicesResponseSelect) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *GetServicesResponseSelect) GetSelectAll() bool {
	if m != nil && m.SelectAll != nil {
		return *m.SelectAll
	}
	return false
}

func (m *GetServicesResponseSelect) GetMeta() *responsemetaselect.ResponseMetaSelect {
	if m != nil {
		return m.Meta
	}
	return nil
}

func (m *GetServicesResponseSelect) GetTotalCount() bool {
	if m != nil && m.TotalCount != nil {
		return *m.TotalCount
	}
	return false
}

func (m *GetServicesResponseSelect) GetServices() *servicesselect.ServicesSelect {
	if m != nil {
		return m.Services
	}
	return nil
}

func (m *GetServicesResponseSelect) GetSelectHash() bool {
	if m != nil && m.SelectHash != nil {
		return *m.SelectHash
	}
	return false
}

func (m *GetServicesResponseSelect) GetHash() string {
	if m != nil && m.Hash != nil {
		return *m.Hash
	}
	return ""
}

func init() {
	proto.RegisterType((*GetServicesResponseSelect)(nil), "getservicesresponseselect.GetServicesResponseSelect")
}

func init() {
	proto.RegisterFile("getservicesresponseselect/getservicesresponseselect_.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 267 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x91, 0x41, 0x4b, 0xc3, 0x30,
	0x14, 0xc7, 0xe9, 0xac, 0xb2, 0xc5, 0x5b, 0x4e, 0xdd, 0x90, 0x51, 0x44, 0xa4, 0x07, 0x4d, 0x71,
	0xc7, 0xde, 0x54, 0x64, 0x5e, 0xbc, 0x74, 0x37, 0x2f, 0x12, 0xe3, 0x23, 0x2d, 0xa4, 0x4d, 0x49,
	0x5e, 0x05, 0x3f, 0x80, 0xdf, 0x5b, 0x9a, 0x26, 0x6e, 0x63, 0xf4, 0xd6, 0xf7, 0x7f, 0xbf, 0xfe,
	0xf2, 0x0f, 0x21, 0x85, 0x04, 0xb4, 0x60, 0xbe, 0x6b, 0x01, 0xd6, 0x80, 0xed, 0x74, 0x6b, 0xc1,
	0x82, 0x02, 0x81, 0xf9, 0xe4, 0xe6, 0x83, 0x75, 0x46, 0xa3, 0xa6, 0xcb, 0x49, 0x62, 0x75, 0x17,
	0xe6, 0x06, 0x90, 0x7b, 0xdf, 0x69, 0xe4, 0x45, 0xab, 0x9b, 0x60, 0xf1, 0xe4, 0xf1, 0xe8, 0xa9,
	0xeb, 0xdf, 0x19, 0x59, 0x6e, 0x01, 0x77, 0x7e, 0x59, 0x7a, 0xdd, 0xce, 0x41, 0xf4, 0x8a, 0x2c,
	0x46, 0xfc, 0x51, 0xa9, 0x24, 0x4a, 0xa3, 0x6c, 0x5e, 0xee, 0x03, 0x5a, 0x90, 0x78, 0x38, 0x36,
	0x99, 0xa5, 0x51, 0x76, 0xb9, 0xb9, 0x65, 0xa7, 0x5d, 0x58, 0xf0, 0xbd, 0x01, 0xf2, 0xd1, 0x59,
	0xba, 0x7f, 0xe8, 0x9a, 0x10, 0xd4, 0xc8, 0xd5, 0xb3, 0xee, 0x5b, 0x4c, 0xce, 0x9c, 0xfa, 0x20,
	0xa1, 0x05, 0x99, 0x87, 0xc2, 0x49, 0xec, 0xfc, 0x6b, 0x76, 0x7c, 0x03, 0x16, 0x3a, 0x7b, 0xef,
	0x3f, 0x3f, 0xb8, 0x47, 0xe4, 0x95, 0xdb, 0x2a, 0x39, 0x1f, 0xdd, 0xfb, 0x84, 0x52, 0x12, 0x57,
	0xc3, 0xe6, 0x22, 0x8d, 0xb2, 0x45, 0xe9, 0xbe, 0x9f, 0xb6, 0xef, 0x2f, 0xb2, 0xc6, 0xaa, 0xff,
	0x64, 0x42, 0x37, 0xf9, 0xe6, 0xc1, 0x62, 0xad, 0x73, 0xa9, 0xef, 0x3b, 0xc5, 0x7f, 0xa4, 0xd1,
	0x7d, 0xfb, 0x95, 0x4b, 0xd3, 0x89, 0xdc, 0x8a, 0x0a, 0x1a, 0x3e, 0xfd, 0x8c, 0x7f, 0x01, 0x00,
	0x00, 0xff, 0xff, 0xcb, 0xf2, 0x87, 0xe7, 0xfc, 0x01, 0x00, 0x00,
}
