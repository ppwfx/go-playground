// Code generated by protoc-gen-go. DO NOT EDIT.
// source: serviceidfilter/serviceidfilter_.proto

/*
Package serviceidfilter is a generated protocol buffer package.

It is generated from these files:
	serviceidfilter/serviceidfilter_.proto

It has these top-level messages:
	ServiceIdFilter
*/
package serviceidfilter

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import stringfilter "github.com/21stio/go-playground/grpc/schema/stringfilter"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type ServiceIdFilter struct {
	Value            *stringfilter.StringFilter `protobuf:"bytes,1,opt,name=value" json:"value,omitempty"`
	ServiceName      *stringfilter.StringFilter `protobuf:"bytes,2,opt,name=serviceName" json:"serviceName,omitempty"`
	And              []*ServiceIdFilter         `protobuf:"bytes,3,rep,name=and" json:"and,omitempty"`
	Or               []*ServiceIdFilter         `protobuf:"bytes,4,rep,name=or" json:"or,omitempty"`
	Not              []*ServiceIdFilter         `protobuf:"bytes,5,rep,name=not" json:"not,omitempty"`
	Hash             *string                    `protobuf:"bytes,6,opt,name=hash" json:"hash,omitempty"`
	XXX_unrecognized []byte                     `json:"-"`
}

func (m *ServiceIdFilter) Reset()                    { *m = ServiceIdFilter{} }
func (m *ServiceIdFilter) String() string            { return proto.CompactTextString(m) }
func (*ServiceIdFilter) ProtoMessage()               {}
func (*ServiceIdFilter) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *ServiceIdFilter) GetValue() *stringfilter.StringFilter {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *ServiceIdFilter) GetServiceName() *stringfilter.StringFilter {
	if m != nil {
		return m.ServiceName
	}
	return nil
}

func (m *ServiceIdFilter) GetAnd() []*ServiceIdFilter {
	if m != nil {
		return m.And
	}
	return nil
}

func (m *ServiceIdFilter) GetOr() []*ServiceIdFilter {
	if m != nil {
		return m.Or
	}
	return nil
}

func (m *ServiceIdFilter) GetNot() []*ServiceIdFilter {
	if m != nil {
		return m.Not
	}
	return nil
}

func (m *ServiceIdFilter) GetHash() string {
	if m != nil && m.Hash != nil {
		return *m.Hash
	}
	return ""
}

func init() {
	proto.RegisterType((*ServiceIdFilter)(nil), "serviceidfilter.ServiceIdFilter")
}

func init() { proto.RegisterFile("serviceidfilter/serviceidfilter_.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 242 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0x31, 0x4b, 0xc4, 0x30,
	0x14, 0xc7, 0x69, 0x7b, 0x27, 0x98, 0x0e, 0x07, 0x99, 0xc2, 0x4d, 0xc1, 0x41, 0xba, 0x98, 0x9c,
	0x5d, 0xd5, 0xc5, 0x41, 0x70, 0x71, 0xe8, 0x6d, 0x2e, 0x12, 0xdb, 0x98, 0x06, 0xda, 0xbc, 0xf2,
	0x92, 0x1e, 0xf8, 0xa9, 0xfc, 0x8a, 0x72, 0xed, 0x09, 0x25, 0x8b, 0xdd, 0xf2, 0x4f, 0x7e, 0x3f,
	0xfe, 0x79, 0x3c, 0x72, 0xeb, 0x35, 0x9e, 0x6c, 0xad, 0x6d, 0xf3, 0x65, 0xbb, 0xa0, 0x51, 0x46,
	0xf9, 0x43, 0x0c, 0x08, 0x01, 0xe8, 0x2e, 0xba, 0xdf, 0x73, 0x1f, 0xd0, 0x3a, 0xf3, 0x67, 0x2d,
	0xc2, 0x45, 0xb9, 0xf9, 0x49, 0xc9, 0xee, 0x38, 0x5b, 0xaf, 0xcd, 0xcb, 0xf4, 0x44, 0x0f, 0x64,
	0x7b, 0x52, 0xdd, 0xa8, 0x59, 0xc2, 0x93, 0x22, 0x2f, 0xf7, 0x62, 0x29, 0x8a, 0xe3, 0x14, 0x66,
	0xb4, 0x9a, 0x41, 0xfa, 0x48, 0xf2, 0x4b, 0xf5, 0x9b, 0xea, 0x35, 0x4b, 0xff, 0xf5, 0x96, 0x38,
	0x2d, 0x49, 0xa6, 0x5c, 0xc3, 0x32, 0x9e, 0x15, 0x79, 0xc9, 0x45, 0x34, 0x84, 0x88, 0xbe, 0x57,
	0x9d, 0x61, 0x7a, 0x20, 0x29, 0x20, 0xdb, 0xac, 0x54, 0x52, 0xc0, 0x73, 0x8b, 0x83, 0xc0, 0xb6,
	0x6b, 0x5b, 0x1c, 0x04, 0x4a, 0xc9, 0xa6, 0x55, 0xbe, 0x65, 0x57, 0x3c, 0x29, 0xae, 0xab, 0xe9,
	0xfc, 0xfc, 0xf4, 0xfe, 0x60, 0x6c, 0x68, 0xc7, 0x4f, 0x51, 0x43, 0x2f, 0xcb, 0x7b, 0x1f, 0x2c,
	0x48, 0x03, 0x77, 0x43, 0xa7, 0xbe, 0x0d, 0xc2, 0xe8, 0x1a, 0x69, 0x70, 0xa8, 0xa5, 0xaf, 0x5b,
	0xdd, 0xab, 0x78, 0x55, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x19, 0x1c, 0xd8, 0x53, 0xcc, 0x01,
	0x00, 0x00,
}
