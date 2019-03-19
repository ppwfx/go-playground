// Code generated by protoc-gen-go. DO NOT EDIT.
// source: serviceportsfilter/serviceportsfilter_.proto

/*
Package serviceportsfilter is a generated protocol buffer package.

It is generated from these files:
	serviceportsfilter/serviceportsfilter_.proto

It has these top-level messages:
	ServicePortsFilter
*/
package serviceportsfilter

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import intfilter "github.com/21stio/go-playground/grpc/schema/intfilter"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type ServicePortsFilter struct {
	Grpc             *intfilter.IntFilter  `protobuf:"bytes,1,opt,name=grpc" json:"grpc,omitempty"`
	Graphql          *intfilter.IntFilter  `protobuf:"bytes,2,opt,name=graphql" json:"graphql,omitempty"`
	And              []*ServicePortsFilter `protobuf:"bytes,3,rep,name=and" json:"and,omitempty"`
	Or               []*ServicePortsFilter `protobuf:"bytes,4,rep,name=or" json:"or,omitempty"`
	Not              []*ServicePortsFilter `protobuf:"bytes,5,rep,name=not" json:"not,omitempty"`
	Hash             *string               `protobuf:"bytes,6,opt,name=hash" json:"hash,omitempty"`
	XXX_unrecognized []byte                `json:"-"`
}

func (m *ServicePortsFilter) Reset()                    { *m = ServicePortsFilter{} }
func (m *ServicePortsFilter) String() string            { return proto.CompactTextString(m) }
func (*ServicePortsFilter) ProtoMessage()               {}
func (*ServicePortsFilter) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *ServicePortsFilter) GetGrpc() *intfilter.IntFilter {
	if m != nil {
		return m.Grpc
	}
	return nil
}

func (m *ServicePortsFilter) GetGraphql() *intfilter.IntFilter {
	if m != nil {
		return m.Graphql
	}
	return nil
}

func (m *ServicePortsFilter) GetAnd() []*ServicePortsFilter {
	if m != nil {
		return m.And
	}
	return nil
}

func (m *ServicePortsFilter) GetOr() []*ServicePortsFilter {
	if m != nil {
		return m.Or
	}
	return nil
}

func (m *ServicePortsFilter) GetNot() []*ServicePortsFilter {
	if m != nil {
		return m.Not
	}
	return nil
}

func (m *ServicePortsFilter) GetHash() string {
	if m != nil && m.Hash != nil {
		return *m.Hash
	}
	return ""
}

func init() {
	proto.RegisterType((*ServicePortsFilter)(nil), "serviceportsfilter.ServicePortsFilter")
}

func init() { proto.RegisterFile("serviceportsfilter/serviceportsfilter_.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 243 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x91, 0x3f, 0x4b, 0x34, 0x31,
	0x10, 0xc6, 0xd9, 0x3f, 0xef, 0x2b, 0xe6, 0xba, 0xc1, 0x22, 0x5c, 0xb5, 0x58, 0xc8, 0x16, 0x9a,
	0xe0, 0x16, 0x62, 0x25, 0x62, 0x21, 0xd8, 0xc9, 0xda, 0xd9, 0x48, 0xdc, 0x5b, 0x93, 0xc0, 0x5e,
	0x26, 0x4e, 0xe6, 0x04, 0xbf, 0x94, 0x9f, 0x51, 0x76, 0xef, 0xbc, 0x66, 0x45, 0xb8, 0x6e, 0xc8,
	0xfc, 0x7e, 0x3c, 0x0f, 0x19, 0x71, 0x9e, 0x7a, 0xfa, 0xf0, 0x5d, 0x1f, 0x91, 0x38, 0xbd, 0xf9,
	0x81, 0x7b, 0xd2, 0xf3, 0xa7, 0x17, 0x15, 0x09, 0x19, 0x01, 0xe6, 0xab, 0xe5, 0xd2, 0x07, 0xde,
	0x89, 0xfb, 0x69, 0xc7, 0x9f, 0x7e, 0xe5, 0x02, 0x9e, 0xb6, 0xca, 0xe3, 0xa8, 0xdc, 0x4f, 0x5b,
	0xa8, 0x45, 0x69, 0x29, 0x76, 0x32, 0xab, 0xb2, 0x7a, 0xd1, 0x9c, 0xa8, 0xbd, 0xa7, 0x1e, 0x02,
	0x6f, 0x99, 0x76, 0x22, 0x40, 0x89, 0x23, 0x4b, 0x26, 0xba, 0xf7, 0x41, 0xe6, 0x7f, 0xc0, 0x3f,
	0x10, 0x5c, 0x8b, 0xc2, 0x84, 0x95, 0x2c, 0xaa, 0xa2, 0x5e, 0x34, 0x67, 0x6a, 0x5e, 0x57, 0xcd,
	0xeb, 0xb4, 0xa3, 0x02, 0x57, 0x22, 0x47, 0x92, 0xe5, 0x41, 0x62, 0x8e, 0x34, 0x26, 0x06, 0x64,
	0xf9, 0xef, 0xb0, 0xc4, 0x80, 0x0c, 0x20, 0x4a, 0x67, 0x92, 0x93, 0xff, 0xab, 0xac, 0x3e, 0x6e,
	0xa7, 0xf9, 0xee, 0xf6, 0xf9, 0xc6, 0x7a, 0x76, 0x9b, 0x57, 0xd5, 0xe1, 0x5a, 0x37, 0x97, 0x89,
	0x3d, 0x6a, 0x8b, 0x17, 0x71, 0x30, 0x9f, 0x96, 0x70, 0x13, 0x56, 0x7a, 0xfc, 0x17, 0x9d, 0x3a,
	0xd7, 0xaf, 0xcd, 0x2f, 0x97, 0xfa, 0x0e, 0x00, 0x00, 0xff, 0xff, 0x5a, 0x9c, 0x73, 0xfc, 0xd1,
	0x01, 0x00, 0x00,
}