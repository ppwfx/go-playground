// Code generated by protoc-gen-go. DO NOT EDIT.
// source: specdescriptionselect/specdescriptionselect_.proto

/*
Package specdescriptionselect is a generated protocol buffer package.

It is generated from these files:
	specdescriptionselect/specdescriptionselect_.proto

It has these top-level messages:
	SpecDescriptionSelect
*/
package specdescriptionselect

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import speccontextsselect "github.com/21stio/go-playground/grpc/schema/speccontextsselect"
import idsselect "github.com/21stio/go-playground/grpc/schema/idsselect"
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

type SpecDescriptionSelect struct {
	SelectAll        *bool                                  `protobuf:"varint,1,opt,name=selectAll" json:"selectAll,omitempty"`
	Name             *bool                                  `protobuf:"varint,2,opt,name=name" json:"name,omitempty"`
	Contexts         *speccontextsselect.SpecContextsSelect `protobuf:"bytes,3,opt,name=contexts" json:"contexts,omitempty"`
	SelectHash       *bool                                  `protobuf:"varint,4,opt,name=selectHash" json:"selectHash,omitempty"`
	Ids              *idsselect.IdsSelect                   `protobuf:"bytes,5,opt,name=ids" json:"ids,omitempty"`
	Meta             *serviceselect.TypeMetaSelect          `protobuf:"bytes,6,opt,name=meta" json:"meta,omitempty"`
	Hash             *string                                `protobuf:"bytes,7,opt,name=hash" json:"hash,omitempty"`
	XXX_unrecognized []byte                                 `json:"-"`
}

func (m *SpecDescriptionSelect) Reset()                    { *m = SpecDescriptionSelect{} }
func (m *SpecDescriptionSelect) String() string            { return proto.CompactTextString(m) }
func (*SpecDescriptionSelect) ProtoMessage()               {}
func (*SpecDescriptionSelect) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *SpecDescriptionSelect) GetSelectAll() bool {
	if m != nil && m.SelectAll != nil {
		return *m.SelectAll
	}
	return false
}

func (m *SpecDescriptionSelect) GetName() bool {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return false
}

func (m *SpecDescriptionSelect) GetContexts() *speccontextsselect.SpecContextsSelect {
	if m != nil {
		return m.Contexts
	}
	return nil
}

func (m *SpecDescriptionSelect) GetSelectHash() bool {
	if m != nil && m.SelectHash != nil {
		return *m.SelectHash
	}
	return false
}

func (m *SpecDescriptionSelect) GetIds() *idsselect.IdsSelect {
	if m != nil {
		return m.Ids
	}
	return nil
}

func (m *SpecDescriptionSelect) GetMeta() *serviceselect.TypeMetaSelect {
	if m != nil {
		return m.Meta
	}
	return nil
}

func (m *SpecDescriptionSelect) GetHash() string {
	if m != nil && m.Hash != nil {
		return *m.Hash
	}
	return ""
}

func init() {
	proto.RegisterType((*SpecDescriptionSelect)(nil), "specdescriptionselect.SpecDescriptionSelect")
}

func init() { proto.RegisterFile("specdescriptionselect/specdescriptionselect_.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 297 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x91, 0xbf, 0x4e, 0xc3, 0x30,
	0x10, 0xc6, 0x95, 0x36, 0x40, 0x6b, 0x36, 0x8b, 0x4a, 0x56, 0x04, 0x28, 0xea, 0x50, 0x65, 0x00,
	0x47, 0xcd, 0x1b, 0xb4, 0x65, 0x80, 0x81, 0x25, 0x65, 0x62, 0x41, 0xc6, 0x39, 0x25, 0x96, 0x92,
	0xd8, 0x8a, 0x5d, 0x44, 0xdf, 0x85, 0x87, 0x45, 0x71, 0xfe, 0x34, 0x55, 0xb3, 0x9d, 0xbf, 0xfb,
	0xdd, 0x77, 0xf7, 0xc9, 0x28, 0xd2, 0x0a, 0x78, 0x02, 0x9a, 0x57, 0x42, 0x19, 0x21, 0x4b, 0x0d,
	0x39, 0x70, 0x13, 0x8e, 0xaa, 0x5f, 0x54, 0x55, 0xd2, 0x48, 0xbc, 0x18, 0xed, 0x7a, 0x4f, 0xb5,
	0xcc, 0x65, 0x69, 0xe0, 0xd7, 0xe8, 0x81, 0xcf, 0xb9, 0xd4, 0x9a, 0x78, 0x9e, 0x48, 0x3a, 0xa8,
	0xaf, 0xba, 0xde, 0x52, 0x43, 0xf5, 0x23, 0x38, 0x74, 0x26, 0xc3, 0x57, 0xcb, 0x2c, 0xff, 0x26,
	0x68, 0xb1, 0x57, 0xc0, 0x5f, 0x4e, 0x77, 0xec, 0x2d, 0x80, 0xef, 0xd1, 0xbc, 0x41, 0x37, 0x79,
	0x4e, 0x1c, 0xdf, 0x09, 0x66, 0xf1, 0x49, 0xc0, 0x18, 0xb9, 0x25, 0x2b, 0x80, 0x4c, 0x6c, 0xc3,
	0xd6, 0x78, 0x8b, 0x66, 0xdd, 0x91, 0x64, 0xea, 0x3b, 0xc1, 0x6d, 0xb4, 0xa2, 0x97, 0x97, 0xd3,
	0x7a, 0xdd, 0xae, 0x95, 0x9a, 0x5d, 0x71, 0x3f, 0x87, 0x1f, 0x11, 0x6a, 0xb0, 0x57, 0xa6, 0x33,
	0xe2, 0x5a, 0xf7, 0x81, 0x82, 0x57, 0x68, 0x2a, 0x12, 0x4d, 0xae, 0xac, 0xfd, 0x1d, 0xed, 0x33,
	0xd3, 0xb7, 0xa4, 0x33, 0xab, 0x01, 0xbc, 0x46, 0x6e, 0x01, 0x86, 0x91, 0x6b, 0x0b, 0x3e, 0xd0,
	0xb3, 0xf0, 0xf4, 0xe3, 0xa8, 0xe0, 0x1d, 0x0c, 0x6b, 0x27, 0x2c, 0x5a, 0x47, 0xca, 0xea, 0xa5,
	0x37, 0xbe, 0x13, 0xcc, 0x63, 0x5b, 0x6f, 0x77, 0x9f, 0x9b, 0x54, 0x98, 0xec, 0xf0, 0x4d, 0xb9,
	0x2c, 0xc2, 0x68, 0xad, 0x8d, 0x90, 0x61, 0x2a, 0x9f, 0x55, 0xce, 0x8e, 0x69, 0x25, 0x0f, 0x65,
	0x12, 0xa6, 0x95, 0xe2, 0xa1, 0xe6, 0x19, 0x14, 0x6c, 0xfc, 0xbf, 0xff, 0x03, 0x00, 0x00, 0xff,
	0xff, 0x6d, 0x52, 0xd9, 0x45, 0x1d, 0x02, 0x00, 0x00,
}
