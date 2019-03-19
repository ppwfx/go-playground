// Code generated by protoc-gen-go. DO NOT EDIT.
// source: service/service_.proto

/*
Package service is a generated protocol buffer package.

It is generated from these files:
	service/service_.proto

It has these top-level messages:
	TypeMeta
	Service
*/
package service

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import rating "github.com/21stio/go-playground/grpc/schema/rating"
import durationvalue "github.com/21stio/go-playground/grpc/schema/durationvalue"
import timestamp "github.com/21stio/go-playground/grpc/schema/timestamp"
import endpoint "github.com/21stio/go-playground/grpc/schema/endpoint"
import serviceports "github.com/21stio/go-playground/grpc/schema/serviceports"
import id "github.com/21stio/go-playground/grpc/schema/id"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type TypeMeta struct {
	Service          *Service                     `protobuf:"bytes,1,opt,name=service" json:"service,omitempty"`
	Relevance        *float64                     `protobuf:"fixed64,2,opt,name=relevance" json:"relevance,omitempty"`
	Rating           *rating.Rating               `protobuf:"bytes,3,opt,name=rating" json:"rating,omitempty"`
	Age              *durationvalue.DurationValue `protobuf:"bytes,4,opt,name=age" json:"age,omitempty"`
	CreatedAt        *timestamp.Timestamp         `protobuf:"bytes,5,opt,name=createdAt" json:"createdAt,omitempty"`
	UpdatedAt        *timestamp.Timestamp         `protobuf:"bytes,6,opt,name=updatedAt" json:"updatedAt,omitempty"`
	Hash             *string                      `protobuf:"bytes,7,opt,name=hash" json:"hash,omitempty"`
	XXX_unrecognized []byte                       `json:"-"`
}

func (m *TypeMeta) Reset()                    { *m = TypeMeta{} }
func (m *TypeMeta) String() string            { return proto.CompactTextString(m) }
func (*TypeMeta) ProtoMessage()               {}
func (*TypeMeta) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *TypeMeta) GetService() *Service {
	if m != nil {
		return m.Service
	}
	return nil
}

func (m *TypeMeta) GetRelevance() float64 {
	if m != nil && m.Relevance != nil {
		return *m.Relevance
	}
	return 0
}

func (m *TypeMeta) GetRating() *rating.Rating {
	if m != nil {
		return m.Rating
	}
	return nil
}

func (m *TypeMeta) GetAge() *durationvalue.DurationValue {
	if m != nil {
		return m.Age
	}
	return nil
}

func (m *TypeMeta) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *TypeMeta) GetUpdatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.UpdatedAt
	}
	return nil
}

func (m *TypeMeta) GetHash() string {
	if m != nil && m.Hash != nil {
		return *m.Hash
	}
	return ""
}

type Service struct {
	Name             *string                    `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Host             *string                    `protobuf:"bytes,2,opt,name=host" json:"host,omitempty"`
	Endpoints        []*endpoint.Endpoint       `protobuf:"bytes,3,rep,name=endpoints" json:"endpoints,omitempty"`
	Ports            *serviceports.ServicePorts `protobuf:"bytes,4,opt,name=ports" json:"ports,omitempty"`
	Hash             *string                    `protobuf:"bytes,5,opt,name=hash" json:"hash,omitempty"`
	Ids              []*id.Id                   `protobuf:"bytes,6,rep,name=ids" json:"ids,omitempty"`
	Meta             *TypeMeta                  `protobuf:"bytes,7,opt,name=meta" json:"meta,omitempty"`
	XXX_unrecognized []byte                     `json:"-"`
}

func (m *Service) Reset()                    { *m = Service{} }
func (m *Service) String() string            { return proto.CompactTextString(m) }
func (*Service) ProtoMessage()               {}
func (*Service) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Service) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *Service) GetHost() string {
	if m != nil && m.Host != nil {
		return *m.Host
	}
	return ""
}

func (m *Service) GetEndpoints() []*endpoint.Endpoint {
	if m != nil {
		return m.Endpoints
	}
	return nil
}

func (m *Service) GetPorts() *serviceports.ServicePorts {
	if m != nil {
		return m.Ports
	}
	return nil
}

func (m *Service) GetHash() string {
	if m != nil && m.Hash != nil {
		return *m.Hash
	}
	return ""
}

func (m *Service) GetIds() []*id.Id {
	if m != nil {
		return m.Ids
	}
	return nil
}

func (m *Service) GetMeta() *TypeMeta {
	if m != nil {
		return m.Meta
	}
	return nil
}

func init() {
	proto.RegisterType((*TypeMeta)(nil), "service.TypeMeta")
	proto.RegisterType((*Service)(nil), "service.Service")
}

func init() { proto.RegisterFile("service/service_.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 419 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x52, 0x5d, 0x8b, 0xd4, 0x30,
	0x14, 0xa5, 0xdb, 0xf9, 0xb0, 0x19, 0x11, 0x0d, 0x8b, 0x84, 0x61, 0x1f, 0xca, 0x80, 0x32, 0x08,
	0xde, 0xac, 0x15, 0x7f, 0x80, 0xa2, 0x0f, 0x3e, 0x08, 0x12, 0x17, 0x1f, 0x7c, 0x91, 0xd8, 0x84,
	0x4e, 0x60, 0xda, 0x84, 0x24, 0x1d, 0xd8, 0x7f, 0xe4, 0x6f, 0xf3, 0x57, 0x48, 0xd2, 0xa4, 0xb3,
	0xf3, 0xb2, 0x4f, 0xf7, 0x9e, 0x7b, 0xce, 0x69, 0x93, 0x73, 0x83, 0x5e, 0x3a, 0x69, 0x4f, 0xaa,
	0x95, 0x34, 0xd5, 0xdf, 0x60, 0xac, 0xf6, 0x1a, 0xaf, 0x13, 0xde, 0x5e, 0x5b, 0xee, 0xd5, 0xd0,
	0xd1, 0xa9, 0x24, 0x7a, 0xbb, 0x13, 0x63, 0x18, 0xe8, 0xe1, 0xc4, 0x8f, 0xa3, 0xa4, 0x17, 0x28,
	0x6b, 0xb6, 0x5e, 0xf5, 0xd2, 0x79, 0xde, 0x1b, 0x3a, 0x77, 0x99, 0x23, 0x72, 0x10, 0x46, 0xab,
	0xc1, 0xd3, 0xdc, 0x64, 0xa6, 0x4e, 0x3f, 0x36, 0xda, 0x7a, 0x47, 0x1f, 0x82, 0xac, 0x78, 0xaa,
	0x04, 0x55, 0x22, 0xa1, 0xdd, 0xdf, 0x2b, 0xf4, 0xe4, 0xee, 0xde, 0xc8, 0x6f, 0xd2, 0x73, 0xfc,
	0x06, 0xe5, 0x73, 0x93, 0xa2, 0x2e, 0xf6, 0x9b, 0xe6, 0x39, 0x24, 0x0c, 0x3f, 0xa6, 0xca, 0xb2,
	0x00, 0xdf, 0xa0, 0xca, 0xca, 0xa3, 0x3c, 0xf1, 0xa1, 0x95, 0xe4, 0xaa, 0x2e, 0xf6, 0x05, 0x3b,
	0x0f, 0xf0, 0x6b, 0xb4, 0x9a, 0x6e, 0x4c, 0xca, 0xf8, 0xa1, 0x67, 0x30, 0x41, 0x60, 0xb1, 0xb0,
	0xc4, 0x62, 0x40, 0x25, 0xef, 0x24, 0x59, 0x44, 0xd1, 0x0d, 0x5c, 0x04, 0x01, 0x9f, 0x13, 0xfa,
	0x19, 0x10, 0x0b, 0x42, 0xdc, 0xa0, 0xaa, 0xb5, 0x92, 0x7b, 0x29, 0x3e, 0x7a, 0xb2, 0x8c, 0xae,
	0x6b, 0x98, 0xe3, 0x81, 0xbb, 0xdc, 0xb1, 0xb3, 0x2c, 0x78, 0x46, 0x23, 0x92, 0x67, 0xf5, 0x98,
	0x67, 0x96, 0x61, 0x8c, 0x16, 0x07, 0xee, 0x0e, 0x64, 0x5d, 0x17, 0xfb, 0x8a, 0xc5, 0x7e, 0xf7,
	0xaf, 0x40, 0xeb, 0x14, 0x43, 0xe0, 0x07, 0xde, 0x4f, 0x31, 0x55, 0x2c, 0xf6, 0xd1, 0xa3, 0x9d,
	0x8f, 0x61, 0x04, 0x8f, 0x76, 0x1e, 0xdf, 0xa2, 0x2a, 0x6f, 0xc8, 0x91, 0xb2, 0x2e, 0xf7, 0x9b,
	0x06, 0x43, 0x9e, 0xc0, 0x97, 0xd4, 0xb0, 0xb3, 0x08, 0xdf, 0xa2, 0x65, 0x5c, 0x57, 0xca, 0x64,
	0x0b, 0x0f, 0x77, 0x98, 0xd7, 0xf0, 0x3d, 0x00, 0x36, 0x09, 0xe7, 0xb3, 0x2e, 0xcf, 0x67, 0xc5,
	0x04, 0x95, 0x4a, 0x38, 0xb2, 0x8a, 0x7f, 0x5c, 0x81, 0x12, 0xf0, 0x55, 0xb0, 0x30, 0xc2, 0xaf,
	0xd0, 0xa2, 0x97, 0x9e, 0xc7, 0x9b, 0x6d, 0x9a, 0x17, 0xf3, 0x82, 0xf3, 0x23, 0x60, 0x91, 0xfe,
	0xf4, 0xe1, 0xd7, 0xfb, 0x4e, 0xf9, 0xc3, 0xf8, 0x07, 0x5a, 0xdd, 0xd3, 0xe6, 0x9d, 0xf3, 0x4a,
	0xd3, 0x4e, 0xbf, 0x35, 0x47, 0x7e, 0xdf, 0x59, 0x3d, 0x0e, 0x82, 0x76, 0xd6, 0xb4, 0xd4, 0xb5,
	0x07, 0xd9, 0xf3, 0xfc, 0xd0, 0xfe, 0x07, 0x00, 0x00, 0xff, 0xff, 0x12, 0xbb, 0x38, 0x16, 0x10,
	0x03, 0x00, 0x00,
}
