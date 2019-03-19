// Code generated by protoc-gen-go. DO NOT EDIT.
// source: updatecompaniesresponseselect/updatecompaniesresponseselect_.proto

/*
Package updatecompaniesresponseselect is a generated protocol buffer package.

It is generated from these files:
	updatecompaniesresponseselect/updatecompaniesresponseselect_.proto

It has these top-level messages:
	UpdateCompaniesResponseSelect
*/
package updatecompaniesresponseselect

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import responsemetaselect "github.com/21stio/go-playground/grpc/schema/responsemetaselect"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type UpdateCompaniesResponseSelect struct {
	SelectAll        *bool                                  `protobuf:"varint,1,opt,name=selectAll" json:"selectAll,omitempty"`
	Meta             *responsemetaselect.ResponseMetaSelect `protobuf:"bytes,2,opt,name=meta" json:"meta,omitempty"`
	SelectHash       *bool                                  `protobuf:"varint,3,opt,name=selectHash" json:"selectHash,omitempty"`
	Hash             *string                                `protobuf:"bytes,4,opt,name=hash" json:"hash,omitempty"`
	XXX_unrecognized []byte                                 `json:"-"`
}

func (m *UpdateCompaniesResponseSelect) Reset()                    { *m = UpdateCompaniesResponseSelect{} }
func (m *UpdateCompaniesResponseSelect) String() string            { return proto.CompactTextString(m) }
func (*UpdateCompaniesResponseSelect) ProtoMessage()               {}
func (*UpdateCompaniesResponseSelect) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *UpdateCompaniesResponseSelect) GetSelectAll() bool {
	if m != nil && m.SelectAll != nil {
		return *m.SelectAll
	}
	return false
}

func (m *UpdateCompaniesResponseSelect) GetMeta() *responsemetaselect.ResponseMetaSelect {
	if m != nil {
		return m.Meta
	}
	return nil
}

func (m *UpdateCompaniesResponseSelect) GetSelectHash() bool {
	if m != nil && m.SelectHash != nil {
		return *m.SelectHash
	}
	return false
}

func (m *UpdateCompaniesResponseSelect) GetHash() string {
	if m != nil && m.Hash != nil {
		return *m.Hash
	}
	return ""
}

func init() {
	proto.RegisterType((*UpdateCompaniesResponseSelect)(nil), "updatecompaniesresponseselect.UpdateCompaniesResponseSelect")
}

func init() {
	proto.RegisterFile("updatecompaniesresponseselect/updatecompaniesresponseselect_.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 229 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x72, 0x2a, 0x2d, 0x48, 0x49,
	0x2c, 0x49, 0x4d, 0xce, 0xcf, 0x2d, 0x48, 0xcc, 0xcb, 0x4c, 0x2d, 0x2e, 0x4a, 0x2d, 0x2e, 0xc8,
	0xcf, 0x2b, 0x4e, 0x2d, 0x4e, 0xcd, 0x49, 0x4d, 0x2e, 0xd1, 0xc7, 0x2b, 0x1b, 0xaf, 0x57, 0x50,
	0x94, 0x5f, 0x92, 0x2f, 0x24, 0x8b, 0x57, 0x95, 0x94, 0x0e, 0x8c, 0x9f, 0x9b, 0x5a, 0x92, 0x08,
	0x35, 0x17, 0x53, 0x08, 0x6a, 0x98, 0xd2, 0x5a, 0x46, 0x2e, 0xd9, 0x50, 0xb0, 0x79, 0xce, 0x30,
	0xf3, 0x82, 0xa0, 0x8a, 0x83, 0xc1, 0x0a, 0x85, 0x64, 0xb8, 0x38, 0x21, 0x5a, 0x1c, 0x73, 0x72,
	0x24, 0x18, 0x15, 0x18, 0x35, 0x38, 0x82, 0x10, 0x02, 0x42, 0x56, 0x5c, 0x2c, 0x20, 0x43, 0x25,
	0x98, 0x14, 0x18, 0x35, 0xb8, 0x8d, 0xd4, 0xf4, 0x30, 0x6d, 0xd2, 0x83, 0x99, 0xe7, 0x9b, 0x5a,
	0x92, 0x08, 0x31, 0x33, 0x08, 0xac, 0x47, 0x48, 0x8e, 0x8b, 0x0b, 0xa2, 0xc4, 0x23, 0xb1, 0x38,
	0x43, 0x82, 0x19, 0x6c, 0x34, 0x92, 0x88, 0x90, 0x10, 0x17, 0x4b, 0x06, 0x48, 0x86, 0x45, 0x81,
	0x51, 0x83, 0x33, 0x08, 0xcc, 0x76, 0xf2, 0x8e, 0xf2, 0x4c, 0xcf, 0x2c, 0xc9, 0x28, 0x4d, 0xd2,
	0x4b, 0xce, 0xcf, 0xd5, 0x37, 0x32, 0x2c, 0x2e, 0xc9, 0xcc, 0xd7, 0x4f, 0xcf, 0xd7, 0x2d, 0xc8,
	0x49, 0xac, 0x4c, 0x2f, 0xca, 0x2f, 0xcd, 0x4b, 0xd1, 0x4f, 0x2f, 0x2a, 0x48, 0xd6, 0x2f, 0x4e,
	0xce, 0x48, 0xcd, 0x4d, 0xc4, 0x1f, 0xa0, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0xf9, 0x04, 0x37,
	0x36, 0x8e, 0x01, 0x00, 0x00,
}
