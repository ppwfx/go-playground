// Code generated by protoc-gen-go. DO NOT EDIT.
// source: servicesort/servicesort_.proto

/*
Package servicesort is a generated protocol buffer package.

It is generated from these files:
	servicesort/servicesort_.proto

It has these top-level messages:
	TypeMetaSort
	ServiceSort
*/
package servicesort

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import sortkind "github.com/21stio/go-playground/grpc/schema/sortkind"
import ratingsort "github.com/21stio/go-playground/grpc/schema/ratingsort"
import durationvaluesort "github.com/21stio/go-playground/grpc/schema/durationvaluesort"
import timestampsort "github.com/21stio/go-playground/grpc/schema/timestampsort"
import serviceportssort "github.com/21stio/go-playground/grpc/schema/serviceportssort"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type TypeMetaSort struct {
	Service          *ServiceSort                         `protobuf:"bytes,1,opt,name=service" json:"service,omitempty"`
	Relevance        *sortkind.SortKind                   `protobuf:"varint,2,opt,name=relevance,enum=sortkind.SortKind" json:"relevance,omitempty"`
	Rating           *ratingsort.RatingSort               `protobuf:"bytes,3,opt,name=rating" json:"rating,omitempty"`
	Age              *durationvaluesort.DurationValueSort `protobuf:"bytes,4,opt,name=age" json:"age,omitempty"`
	CreatedAt        *timestampsort.TimestampSort         `protobuf:"bytes,5,opt,name=createdAt" json:"createdAt,omitempty"`
	UpdatedAt        *timestampsort.TimestampSort         `protobuf:"bytes,6,opt,name=updatedAt" json:"updatedAt,omitempty"`
	Hash             *string                              `protobuf:"bytes,7,opt,name=hash" json:"hash,omitempty"`
	XXX_unrecognized []byte                               `json:"-"`
}

func (m *TypeMetaSort) Reset()                    { *m = TypeMetaSort{} }
func (m *TypeMetaSort) String() string            { return proto.CompactTextString(m) }
func (*TypeMetaSort) ProtoMessage()               {}
func (*TypeMetaSort) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *TypeMetaSort) GetService() *ServiceSort {
	if m != nil {
		return m.Service
	}
	return nil
}

func (m *TypeMetaSort) GetRelevance() sortkind.SortKind {
	if m != nil && m.Relevance != nil {
		return *m.Relevance
	}
	return sortkind.SortKind_ASC
}

func (m *TypeMetaSort) GetRating() *ratingsort.RatingSort {
	if m != nil {
		return m.Rating
	}
	return nil
}

func (m *TypeMetaSort) GetAge() *durationvaluesort.DurationValueSort {
	if m != nil {
		return m.Age
	}
	return nil
}

func (m *TypeMetaSort) GetCreatedAt() *timestampsort.TimestampSort {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *TypeMetaSort) GetUpdatedAt() *timestampsort.TimestampSort {
	if m != nil {
		return m.UpdatedAt
	}
	return nil
}

func (m *TypeMetaSort) GetHash() string {
	if m != nil && m.Hash != nil {
		return *m.Hash
	}
	return ""
}

type ServiceSort struct {
	Name             *sortkind.SortKind                 `protobuf:"varint,1,opt,name=name,enum=sortkind.SortKind" json:"name,omitempty"`
	Host             *sortkind.SortKind                 `protobuf:"varint,2,opt,name=host,enum=sortkind.SortKind" json:"host,omitempty"`
	Ports            *serviceportssort.ServicePortsSort `protobuf:"bytes,3,opt,name=ports" json:"ports,omitempty"`
	Meta             *TypeMetaSort                      `protobuf:"bytes,4,opt,name=meta" json:"meta,omitempty"`
	Hash             *string                            `protobuf:"bytes,5,opt,name=hash" json:"hash,omitempty"`
	XXX_unrecognized []byte                             `json:"-"`
}

func (m *ServiceSort) Reset()                    { *m = ServiceSort{} }
func (m *ServiceSort) String() string            { return proto.CompactTextString(m) }
func (*ServiceSort) ProtoMessage()               {}
func (*ServiceSort) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *ServiceSort) GetName() sortkind.SortKind {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return sortkind.SortKind_ASC
}

func (m *ServiceSort) GetHost() sortkind.SortKind {
	if m != nil && m.Host != nil {
		return *m.Host
	}
	return sortkind.SortKind_ASC
}

func (m *ServiceSort) GetPorts() *serviceportssort.ServicePortsSort {
	if m != nil {
		return m.Ports
	}
	return nil
}

func (m *ServiceSort) GetMeta() *TypeMetaSort {
	if m != nil {
		return m.Meta
	}
	return nil
}

func (m *ServiceSort) GetHash() string {
	if m != nil && m.Hash != nil {
		return *m.Hash
	}
	return ""
}

func init() {
	proto.RegisterType((*TypeMetaSort)(nil), "servicesort.TypeMetaSort")
	proto.RegisterType((*ServiceSort)(nil), "servicesort.ServiceSort")
}

func init() { proto.RegisterFile("servicesort/servicesort_.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 411 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x93, 0xcd, 0x6a, 0xe3, 0x30,
	0x14, 0x85, 0x71, 0xe2, 0x24, 0x44, 0x19, 0x66, 0xa1, 0xc5, 0xa0, 0x09, 0x61, 0x08, 0x66, 0x28,
	0xa6, 0x10, 0xb9, 0xf5, 0xa2, 0x7f, 0xbb, 0x96, 0xee, 0x4a, 0xa1, 0x38, 0xa1, 0x8b, 0x6e, 0x8a,
	0x6a, 0x0b, 0xdb, 0x34, 0xb6, 0x8c, 0x24, 0x07, 0xf2, 0xaa, 0x7d, 0x82, 0x3e, 0x46, 0x91, 0x2c,
	0xc5, 0x4e, 0x53, 0x4a, 0x37, 0xe6, 0xea, 0xe8, 0x3b, 0x96, 0xee, 0xb9, 0x36, 0xf8, 0x27, 0x28,
	0xdf, 0xe4, 0x31, 0x15, 0x8c, 0xcb, 0xa0, 0x53, 0x3f, 0xe3, 0x8a, 0x33, 0xc9, 0xe0, 0xa4, 0xa3,
	0x4d, 0x91, 0x7a, 0xbe, 0xe6, 0x65, 0x12, 0xd8, 0xc2, 0x60, 0xd3, 0x19, 0x27, 0x32, 0x2f, 0x53,
	0xfd, 0x96, 0xb6, 0xb4, 0xbb, 0xc7, 0x49, 0xad, 0x44, 0x56, 0x6e, 0xc8, 0xba, 0x6e, 0x8e, 0x3a,
	0x50, 0x2c, 0xeb, 0xc9, 0xbc, 0xa0, 0x42, 0x92, 0xa2, 0xd2, 0xdc, 0xde, 0xca, 0x32, 0xbe, 0xb9,
	0x54, 0xc5, 0xb8, 0x14, 0xdd, 0x9b, 0xef, 0x04, 0x43, 0x7a, 0xef, 0x3d, 0xf0, 0x6b, 0xb5, 0xad,
	0xe8, 0x3d, 0x95, 0x64, 0xc9, 0xb8, 0x84, 0x21, 0x18, 0x19, 0x16, 0x39, 0x73, 0xc7, 0x9f, 0x84,
	0x08, 0x77, 0x3a, 0xc4, 0xcb, 0xa6, 0x56, 0x68, 0x64, 0x41, 0x78, 0x02, 0xc6, 0x9c, 0xae, 0xe9,
	0x86, 0x94, 0x31, 0x45, 0xbd, 0xb9, 0xe3, 0xff, 0x0e, 0x21, 0xb6, 0x09, 0x60, 0xc5, 0xde, 0xe5,
	0x65, 0x12, 0xb5, 0x10, 0xc4, 0x60, 0xd8, 0xa4, 0x80, 0xfa, 0xfa, 0x90, 0x3f, 0xb8, 0x0d, 0x05,
	0x47, 0xba, 0xd4, 0x47, 0x18, 0x0a, 0x9e, 0x81, 0x3e, 0x49, 0x29, 0x72, 0x35, 0xfc, 0x1f, 0x1f,
	0x84, 0x83, 0x6f, 0x8d, 0xf2, 0xa8, 0x14, 0x6d, 0x55, 0x06, 0x78, 0x05, 0xc6, 0x31, 0xa7, 0x44,
	0xd2, 0xe4, 0x5a, 0xa2, 0x81, 0x76, 0xcf, 0xf0, 0x5e, 0x64, 0x78, 0x65, 0x57, 0xda, 0xd5, 0xe2,
	0xca, 0x5b, 0x57, 0x89, 0xf1, 0x0e, 0x7f, 0xe2, 0xdd, 0xe1, 0x10, 0x02, 0x37, 0x23, 0x22, 0x43,
	0xa3, 0xb9, 0xe3, 0x8f, 0x23, 0x5d, 0x7b, 0x6f, 0x0e, 0x98, 0x74, 0xe2, 0x83, 0x47, 0xc0, 0x2d,
	0x49, 0xd1, 0xc4, 0xfc, 0x75, 0x60, 0x7a, 0x5f, 0x71, 0x19, 0x13, 0xf2, 0x9b, 0x60, 0xf5, 0x3e,
	0xbc, 0x00, 0x03, 0x3d, 0x5e, 0x13, 0xa9, 0x87, 0x3f, 0xcf, 0xdc, 0x0e, 0xef, 0x41, 0x09, 0xfa,
	0xc6, 0x8d, 0x01, 0x2e, 0x80, 0x5b, 0x50, 0x49, 0x4c, 0xbc, 0x7f, 0xf7, 0x06, 0xde, 0xfd, 0x38,
	0x22, 0x8d, 0xed, 0x9a, 0x1b, 0xb4, 0xcd, 0xdd, 0x5c, 0x3e, 0x9d, 0xa7, 0xb9, 0xcc, 0xea, 0x17,
	0x1c, 0xb3, 0x22, 0x08, 0x4f, 0x85, 0xcc, 0x59, 0x90, 0xb2, 0x45, 0xb5, 0x26, 0xdb, 0x94, 0xb3,
	0xba, 0x4c, 0x82, 0x94, 0x57, 0x71, 0x20, 0xe2, 0x8c, 0x16, 0xa4, 0xfb, 0x23, 0x7d, 0x04, 0x00,
	0x00, 0xff, 0xff, 0x75, 0x32, 0xb7, 0x22, 0x62, 0x03, 0x00, 0x00,
}
