// Code generated by protoc-gen-go. DO NOT EDIT.
// source: physicalssort/physicalssort_.proto

/*
Package physicalssort is a generated protocol buffer package.

It is generated from these files:
	physicalssort/physicalssort_.proto

It has these top-level messages:
	PhysicalsSort
*/
package physicalssort

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import weightvaluesort "github.com/21stio/go-playground/grpc/schema/weightvaluesort"
import volumevaluesort "github.com/21stio/go-playground/grpc/schema/volumevaluesort"
import lengthvaluesort "github.com/21stio/go-playground/grpc/schema/lengthvaluesort"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type PhysicalsSort struct {
	Weight           *weightvaluesort.WeightValueSort `protobuf:"bytes,1,opt,name=weight" json:"weight,omitempty"`
	Volume           *volumevaluesort.VolumeValueSort `protobuf:"bytes,2,opt,name=volume" json:"volume,omitempty"`
	Length           *lengthvaluesort.LengthValueSort `protobuf:"bytes,3,opt,name=length" json:"length,omitempty"`
	Height           *lengthvaluesort.LengthValueSort `protobuf:"bytes,4,opt,name=height" json:"height,omitempty"`
	Depth            *lengthvaluesort.LengthValueSort `protobuf:"bytes,5,opt,name=depth" json:"depth,omitempty"`
	Hash             *string                          `protobuf:"bytes,6,opt,name=hash" json:"hash,omitempty"`
	XXX_unrecognized []byte                           `json:"-"`
}

func (m *PhysicalsSort) Reset()                    { *m = PhysicalsSort{} }
func (m *PhysicalsSort) String() string            { return proto.CompactTextString(m) }
func (*PhysicalsSort) ProtoMessage()               {}
func (*PhysicalsSort) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *PhysicalsSort) GetWeight() *weightvaluesort.WeightValueSort {
	if m != nil {
		return m.Weight
	}
	return nil
}

func (m *PhysicalsSort) GetVolume() *volumevaluesort.VolumeValueSort {
	if m != nil {
		return m.Volume
	}
	return nil
}

func (m *PhysicalsSort) GetLength() *lengthvaluesort.LengthValueSort {
	if m != nil {
		return m.Length
	}
	return nil
}

func (m *PhysicalsSort) GetHeight() *lengthvaluesort.LengthValueSort {
	if m != nil {
		return m.Height
	}
	return nil
}

func (m *PhysicalsSort) GetDepth() *lengthvaluesort.LengthValueSort {
	if m != nil {
		return m.Depth
	}
	return nil
}

func (m *PhysicalsSort) GetHash() string {
	if m != nil && m.Hash != nil {
		return *m.Hash
	}
	return ""
}

func init() {
	proto.RegisterType((*PhysicalsSort)(nil), "physicalssort.PhysicalsSort")
}

func init() { proto.RegisterFile("physicalssort/physicalssort_.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 260 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xbf, 0x4b, 0x03, 0x31,
	0x1c, 0x47, 0xb9, 0xda, 0x16, 0x8c, 0x74, 0xc9, 0x14, 0x3a, 0x1d, 0x1d, 0x4a, 0x17, 0x13, 0xec,
	0x20, 0x8a, 0x9b, 0xb3, 0x83, 0x54, 0xa8, 0xe0, 0x22, 0xf1, 0x1a, 0x92, 0x83, 0x5c, 0xbf, 0x21,
	0xc9, 0x55, 0xfa, 0xf7, 0xf9, 0x8f, 0x49, 0x7e, 0x68, 0x7b, 0x99, 0xba, 0xe5, 0x85, 0xf7, 0xc8,
	0x67, 0x08, 0x5a, 0x18, 0x75, 0x74, 0x6d, 0xc3, 0xb5, 0x73, 0x60, 0x3d, 0x1b, 0xd0, 0x27, 0x35,
	0x16, 0x3c, 0xe0, 0xd9, 0xe0, 0x76, 0xbe, 0xfc, 0x16, 0xad, 0x54, 0xfe, 0xc0, 0x75, 0x2f, 0x62,
	0x54, 0x70, 0xce, 0xe6, 0xcb, 0x03, 0xe8, 0xbe, 0x13, 0x27, 0xaf, 0xe0, 0x7f, 0x4f, 0x8b, 0xbd,
	0xf4, 0xea, 0xe4, 0x15, 0x9c, 0xbd, 0xc5, 0xcf, 0x08, 0xcd, 0x5e, 0xff, 0x96, 0xbc, 0x81, 0xf5,
	0xf8, 0x01, 0x4d, 0xd3, 0xdb, 0xa4, 0xaa, 0xab, 0xd5, 0xcd, 0xba, 0xa6, 0xc5, 0x14, 0xfa, 0x1e,
	0x79, 0x1b, 0x38, 0x14, 0x9b, 0xec, 0x87, 0x32, 0xad, 0x21, 0xa3, 0x5c, 0x16, 0xe3, 0xe8, 0x36,
	0xf2, 0x59, 0x99, 0x84, 0x50, 0xa6, 0x7d, 0xe4, 0x2a, 0x97, 0xc5, 0x5c, 0xfa, 0x12, 0xf9, 0xac,
	0x4c, 0x42, 0x28, 0x55, 0x5a, 0x3b, 0xbe, 0xb4, 0x4c, 0x3e, 0xbe, 0x47, 0x93, 0x9d, 0x30, 0x5e,
	0x91, 0xc9, 0x85, 0x61, 0xd2, 0x31, 0x46, 0x63, 0xc5, 0x9d, 0x22, 0xd3, 0xba, 0x5a, 0x5d, 0x6f,
	0xe2, 0xf9, 0xf9, 0xe9, 0xe3, 0x51, 0xb6, 0x5e, 0xf5, 0x5f, 0xb4, 0x81, 0x8e, 0xad, 0xef, 0x9c,
	0x6f, 0x81, 0x49, 0xb8, 0x35, 0x9a, 0x1f, 0xa5, 0x85, 0x7e, 0xbf, 0x63, 0xd2, 0x9a, 0x86, 0xb9,
	0x46, 0x89, 0x8e, 0x0f, 0x3f, 0xc4, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x70, 0xfc, 0x10, 0x42,
	0x2e, 0x02, 0x00, 0x00,
}
