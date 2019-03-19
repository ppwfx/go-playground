// Code generated by protoc-gen-go. DO NOT EDIT.
// source: physicals/physicals_.proto

/*
Package physicals is a generated protocol buffer package.

It is generated from these files:
	physicals/physicals_.proto

It has these top-level messages:
	Physicals
*/
package physicals

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import phasekind "github.com/21stio/go-playground/grpc/schema/phasekind"
import weightvalue "github.com/21stio/go-playground/grpc/schema/weightvalue"
import volumevalue "github.com/21stio/go-playground/grpc/schema/volumevalue"
import lengthvalue "github.com/21stio/go-playground/grpc/schema/lengthvalue"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Physicals struct {
	Phase            *phasekind.PhaseKind     `protobuf:"varint,1,opt,name=phase,enum=phasekind.PhaseKind" json:"phase,omitempty"`
	Weight           *weightvalue.WeightValue `protobuf:"bytes,2,opt,name=weight" json:"weight,omitempty"`
	Volume           *volumevalue.VolumeValue `protobuf:"bytes,3,opt,name=volume" json:"volume,omitempty"`
	Length           *lengthvalue.LengthValue `protobuf:"bytes,4,opt,name=length" json:"length,omitempty"`
	Height           *lengthvalue.LengthValue `protobuf:"bytes,5,opt,name=height" json:"height,omitempty"`
	Depth            *lengthvalue.LengthValue `protobuf:"bytes,6,opt,name=depth" json:"depth,omitempty"`
	Hash             *string                  `protobuf:"bytes,7,opt,name=hash" json:"hash,omitempty"`
	XXX_unrecognized []byte                   `json:"-"`
}

func (m *Physicals) Reset()                    { *m = Physicals{} }
func (m *Physicals) String() string            { return proto.CompactTextString(m) }
func (*Physicals) ProtoMessage()               {}
func (*Physicals) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Physicals) GetPhase() phasekind.PhaseKind {
	if m != nil && m.Phase != nil {
		return *m.Phase
	}
	return phasekind.PhaseKind_SOLID
}

func (m *Physicals) GetWeight() *weightvalue.WeightValue {
	if m != nil {
		return m.Weight
	}
	return nil
}

func (m *Physicals) GetVolume() *volumevalue.VolumeValue {
	if m != nil {
		return m.Volume
	}
	return nil
}

func (m *Physicals) GetLength() *lengthvalue.LengthValue {
	if m != nil {
		return m.Length
	}
	return nil
}

func (m *Physicals) GetHeight() *lengthvalue.LengthValue {
	if m != nil {
		return m.Height
	}
	return nil
}

func (m *Physicals) GetDepth() *lengthvalue.LengthValue {
	if m != nil {
		return m.Depth
	}
	return nil
}

func (m *Physicals) GetHash() string {
	if m != nil && m.Hash != nil {
		return *m.Hash
	}
	return ""
}

func init() {
	proto.RegisterType((*Physicals)(nil), "physicals.Physicals")
}

func init() { proto.RegisterFile("physicals/physicals_.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 281 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0x31, 0x4b, 0xc4, 0x30,
	0x1c, 0x47, 0xe9, 0x79, 0x77, 0x72, 0x11, 0x1c, 0x82, 0x43, 0xe8, 0x20, 0x87, 0x53, 0x11, 0x4c,
	0xb4, 0x20, 0xee, 0xae, 0x3a, 0x1c, 0x1d, 0x4e, 0x70, 0x91, 0xd8, 0x86, 0x24, 0xd8, 0x36, 0xa1,
	0x49, 0x4f, 0xee, 0xf3, 0xf9, 0xc5, 0xa4, 0xfd, 0xc7, 0x98, 0xed, 0xb6, 0x47, 0x7f, 0xff, 0x07,
	0xaf, 0x10, 0x94, 0x5b, 0x75, 0x74, 0xba, 0xe6, 0xad, 0x63, 0x91, 0x3e, 0xa8, 0x1d, 0x8c, 0x37,
	0x78, 0x13, 0xbf, 0xe4, 0xb9, 0x55, 0xdc, 0x89, 0x2f, 0xdd, 0x37, 0x2c, 0x52, 0x38, 0xcb, 0xaf,
	0xbf, 0x85, 0x96, 0xca, 0x1f, 0x78, 0x3b, 0x0a, 0x96, 0x70, 0xdc, 0x0f, 0xa6, 0x1d, 0x3b, 0x01,
	0x7b, 0xc2, 0x71, 0x6f, 0x45, 0x2f, 0xbd, 0x82, 0x3d, 0xe1, 0xb0, 0xdf, 0xfc, 0x2c, 0xd0, 0x66,
	0xf7, 0x57, 0x82, 0x6f, 0xd1, 0x6a, 0x2e, 0x20, 0xd9, 0x36, 0x2b, 0x2e, 0xcb, 0x2b, 0x1a, 0x7b,
	0xe8, 0x6e, 0xa2, 0x17, 0xdd, 0x37, 0x15, 0x9c, 0xe0, 0x7b, 0xb4, 0x86, 0x1e, 0xb2, 0xd8, 0x66,
	0xc5, 0x45, 0x49, 0x68, 0x92, 0x47, 0xdf, 0x66, 0xde, 0x4f, 0x5c, 0x85, 0xbb, 0xc9, 0x80, 0x42,
	0x72, 0x16, 0x8c, 0x24, 0x98, 0xee, 0x67, 0x0e, 0x06, 0x0c, 0x93, 0x01, 0xcd, 0x64, 0x19, 0x8c,
	0xe4, 0x17, 0xe8, 0xeb, 0xcc, 0xc1, 0x80, 0x61, 0x32, 0x14, 0x54, 0xad, 0x4e, 0x19, 0x70, 0x87,
	0x29, 0x5a, 0x35, 0xc2, 0x7a, 0x45, 0xd6, 0x27, 0x04, 0x38, 0xc3, 0x18, 0x2d, 0x15, 0x77, 0x8a,
	0x9c, 0x6f, 0xb3, 0x62, 0x53, 0xcd, 0xfc, 0xfc, 0xf4, 0xfe, 0x28, 0xb5, 0x57, 0xe3, 0x27, 0xad,
	0x4d, 0xc7, 0xca, 0x07, 0xe7, 0xb5, 0x61, 0xd2, 0xdc, 0xd9, 0x96, 0x1f, 0xe5, 0x60, 0xc6, 0xbe,
	0x61, 0x72, 0xb0, 0x35, 0x73, 0xb5, 0x12, 0x1d, 0xff, 0x7f, 0x0c, 0xbf, 0x01, 0x00, 0x00, 0xff,
	0xff, 0x9b, 0x59, 0x01, 0x1f, 0x22, 0x02, 0x00, 0x00,
}
