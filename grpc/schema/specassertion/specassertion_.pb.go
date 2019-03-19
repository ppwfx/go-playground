// Code generated by protoc-gen-go. DO NOT EDIT.
// source: specassertion/specassertion_.proto

/*
Package specassertion is a generated protocol buffer package.

It is generated from these files:
	specassertion/specassertion_.proto

It has these top-level messages:
	SpecAssertion
*/
package specassertion

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type SpecAssertion struct {
	Name             *string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Pending          *bool   `protobuf:"varint,2,opt,name=pending" json:"pending,omitempty"`
	Passed           *bool   `protobuf:"varint,3,opt,name=passed" json:"passed,omitempty"`
	Hash             *string `protobuf:"bytes,4,opt,name=hash" json:"hash,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *SpecAssertion) Reset()                    { *m = SpecAssertion{} }
func (m *SpecAssertion) String() string            { return proto.CompactTextString(m) }
func (*SpecAssertion) ProtoMessage()               {}
func (*SpecAssertion) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *SpecAssertion) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *SpecAssertion) GetPending() bool {
	if m != nil && m.Pending != nil {
		return *m.Pending
	}
	return false
}

func (m *SpecAssertion) GetPassed() bool {
	if m != nil && m.Passed != nil {
		return *m.Passed
	}
	return false
}

func (m *SpecAssertion) GetHash() string {
	if m != nil && m.Hash != nil {
		return *m.Hash
	}
	return ""
}

func init() {
	proto.RegisterType((*SpecAssertion)(nil), "specassertion.SpecAssertion")
}

func init() { proto.RegisterFile("specassertion/specassertion_.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 175 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x2a, 0x2e, 0x48, 0x4d,
	0x4e, 0x2c, 0x2e, 0x4e, 0x2d, 0x2a, 0xc9, 0xcc, 0xcf, 0xd3, 0x47, 0xe1, 0xc5, 0xeb, 0x15, 0x14,
	0xe5, 0x97, 0xe4, 0x0b, 0xf1, 0xa2, 0x88, 0x2a, 0x65, 0x72, 0xf1, 0x06, 0x17, 0xa4, 0x26, 0x3b,
	0xc2, 0x04, 0x84, 0x84, 0xb8, 0x58, 0xf2, 0x12, 0x73, 0x53, 0x25, 0x18, 0x15, 0x18, 0x35, 0x38,
	0x83, 0xc0, 0x6c, 0x21, 0x09, 0x2e, 0xf6, 0x82, 0xd4, 0xbc, 0x94, 0xcc, 0xbc, 0x74, 0x09, 0x26,
	0x05, 0x46, 0x0d, 0x8e, 0x20, 0x18, 0x57, 0x48, 0x8c, 0x8b, 0xad, 0x00, 0x64, 0x58, 0x8a, 0x04,
	0x33, 0x58, 0x02, 0xca, 0x03, 0x99, 0x92, 0x91, 0x58, 0x9c, 0x21, 0xc1, 0x02, 0x31, 0x05, 0xc4,
	0x76, 0xb2, 0x8e, 0xb2, 0x4c, 0xcf, 0x2c, 0xc9, 0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x37,
	0x32, 0x2c, 0x2e, 0xc9, 0xcc, 0xd7, 0x4f, 0xcf, 0xd7, 0x2d, 0xc8, 0x49, 0xac, 0x4c, 0x2f, 0xca,
	0x2f, 0xcd, 0x4b, 0xd1, 0x4f, 0x2f, 0x2a, 0x48, 0xd6, 0x2f, 0x4e, 0xce, 0x48, 0xcd, 0x4d, 0x44,
	0x75, 0x3d, 0x20, 0x00, 0x00, 0xff, 0xff, 0x7a, 0xaf, 0xd7, 0x7d, 0xdb, 0x00, 0x00, 0x00,
}
