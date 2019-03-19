// Code generated by protoc-gen-go. DO NOT EDIT.
// source: productsort/productsort_.proto

/*
Package productsort is a generated protocol buffer package.

It is generated from these files:
	productsort/productsort_.proto

It has these top-level messages:
	ProductSort
*/
package productsort

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import infosort "github.com/21stio/go-playground/grpc/schema/infosort"
import brandsort "github.com/21stio/go-playground/grpc/schema/brandsort"
import nutritionssort "github.com/21stio/go-playground/grpc/schema/nutritionssort"
import physicalssort "github.com/21stio/go-playground/grpc/schema/physicalssort"
import servicesort "github.com/21stio/go-playground/grpc/schema/servicesort"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type ProductSort struct {
	Info             *infosort.InfoSort             `protobuf:"bytes,1,opt,name=info" json:"info,omitempty"`
	Brand            *brandsort.BrandSort           `protobuf:"bytes,2,opt,name=brand" json:"brand,omitempty"`
	Nutritions       *nutritionssort.NutritionsSort `protobuf:"bytes,3,opt,name=nutritions" json:"nutritions,omitempty"`
	Physicals        *physicalssort.PhysicalsSort   `protobuf:"bytes,4,opt,name=physicals" json:"physicals,omitempty"`
	Meta             *servicesort.TypeMetaSort      `protobuf:"bytes,5,opt,name=meta" json:"meta,omitempty"`
	Hash             *string                        `protobuf:"bytes,6,opt,name=hash" json:"hash,omitempty"`
	XXX_unrecognized []byte                         `json:"-"`
}

func (m *ProductSort) Reset()                    { *m = ProductSort{} }
func (m *ProductSort) String() string            { return proto.CompactTextString(m) }
func (*ProductSort) ProtoMessage()               {}
func (*ProductSort) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *ProductSort) GetInfo() *infosort.InfoSort {
	if m != nil {
		return m.Info
	}
	return nil
}

func (m *ProductSort) GetBrand() *brandsort.BrandSort {
	if m != nil {
		return m.Brand
	}
	return nil
}

func (m *ProductSort) GetNutritions() *nutritionssort.NutritionsSort {
	if m != nil {
		return m.Nutritions
	}
	return nil
}

func (m *ProductSort) GetPhysicals() *physicalssort.PhysicalsSort {
	if m != nil {
		return m.Physicals
	}
	return nil
}

func (m *ProductSort) GetMeta() *servicesort.TypeMetaSort {
	if m != nil {
		return m.Meta
	}
	return nil
}

func (m *ProductSort) GetHash() string {
	if m != nil && m.Hash != nil {
		return *m.Hash
	}
	return ""
}

func init() {
	proto.RegisterType((*ProductSort)(nil), "productsort.ProductSort")
}

func init() { proto.RegisterFile("productsort/productsort_.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 306 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x92, 0x41, 0x6b, 0xfa, 0x40,
	0x14, 0xc4, 0xd1, 0x7f, 0xfc, 0x83, 0xeb, 0x6d, 0xe9, 0x61, 0x2b, 0x45, 0x44, 0x4a, 0x91, 0x82,
	0x2f, 0xd4, 0x4b, 0x69, 0x0f, 0x3d, 0x78, 0xeb, 0xa1, 0x45, 0x6c, 0x4f, 0xbd, 0x94, 0x35, 0xc6,
	0x64, 0x41, 0xf7, 0x2d, 0xbb, 0x9b, 0x82, 0xdf, 0xa6, 0x1f, 0xb5, 0xe4, 0xc5, 0x8d, 0x9b, 0xdb,
	0xbc, 0xcc, 0x6f, 0x60, 0x26, 0x2c, 0x9b, 0x18, 0x8b, 0xbb, 0x2a, 0xf3, 0x0e, 0xad, 0x4f, 0x23,
	0xfd, 0x0d, 0xc6, 0xa2, 0x47, 0x3e, 0x8a, 0xbe, 0x8d, 0x85, 0xd2, 0x7b, 0x24, 0x32, 0x88, 0x33,
	0x36, 0x1e, 0x6f, 0xad, 0xd4, 0x3b, 0xb2, 0x5a, 0x15, 0xbc, 0x5b, 0x5d, 0x79, 0xab, 0xbc, 0x42,
	0xed, 0x08, 0xe8, 0x9e, 0x81, 0x9a, 0x99, 0xf2, 0xe4, 0x54, 0x26, 0x0f, 0x0d, 0xd4, 0xb9, 0x02,
	0x33, 0x71, 0xb9, 0xfd, 0x51, 0x59, 0x4e, 0x44, 0xa4, 0xcf, 0xfe, 0xec, 0xb7, 0xcf, 0x46, 0xeb,
	0xa6, 0xef, 0x07, 0x5a, 0xcf, 0xef, 0x58, 0x52, 0x17, 0x15, 0xbd, 0x69, 0x6f, 0x3e, 0x5a, 0x72,
	0x08, 0xad, 0xe1, 0x55, 0xef, 0xb1, 0x26, 0x36, 0xe4, 0xf3, 0x7b, 0x36, 0xa0, 0xd6, 0xa2, 0x4f,
	0xe0, 0x15, 0xb4, 0x1b, 0x60, 0x55, 0x2b, 0x42, 0x1b, 0x84, 0xbf, 0x30, 0x76, 0x19, 0x20, 0xfe,
	0x51, 0x60, 0x02, 0xdd, 0x4d, 0xf0, 0xde, 0x9e, 0x14, 0x8d, 0x12, 0xfc, 0x99, 0x0d, 0xdb, 0x6d,
	0x22, 0xa1, 0xf8, 0x0d, 0x74, 0xd6, 0xc2, 0x3a, 0x5c, 0x14, 0xbe, 0xe0, 0x7c, 0xc1, 0x92, 0x63,
	0xee, 0xa5, 0x18, 0x50, 0xec, 0x1a, 0xa2, 0x5f, 0x00, 0x9f, 0x27, 0x93, 0xbf, 0xe5, 0x5e, 0x36,
	0xb3, 0x6a, 0x8c, 0x73, 0x96, 0x94, 0xd2, 0x95, 0xe2, 0xff, 0xb4, 0x37, 0x1f, 0x6e, 0x48, 0xaf,
	0x9e, 0xbe, 0x1e, 0x0b, 0xe5, 0xcb, 0x6a, 0x0b, 0x19, 0x1e, 0xd3, 0xe5, 0x83, 0xf3, 0x0a, 0xd3,
	0x02, 0x17, 0xe6, 0x20, 0x4f, 0x85, 0xc5, 0x4a, 0xef, 0xd2, 0xc2, 0x9a, 0x2c, 0x75, 0x59, 0x99,
	0x1f, 0x65, 0xfc, 0x22, 0xfe, 0x02, 0x00, 0x00, 0xff, 0xff, 0x26, 0xeb, 0xda, 0xe2, 0x2b, 0x02,
	0x00, 0x00,
}