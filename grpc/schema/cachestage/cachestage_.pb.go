// Code generated by protoc-gen-go. DO NOT EDIT.
// source: cachestage/cachestage_.proto

/*
Package cachestage is a generated protocol buffer package.

It is generated from these files:
	cachestage/cachestage_.proto

It has these top-level messages:
	CacheStage
*/
package cachestage

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import durationscalar "github.com/21stio/go-playground/grpc/schema/durationscalar"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type CacheStage struct {
	MaxAge           *durationscalar.DurationScalar `protobuf:"bytes,1,opt,name=maxAge" json:"maxAge,omitempty"`
	Hash             *string                        `protobuf:"bytes,2,opt,name=hash" json:"hash,omitempty"`
	XXX_unrecognized []byte                         `json:"-"`
}

func (m *CacheStage) Reset()                    { *m = CacheStage{} }
func (m *CacheStage) String() string            { return proto.CompactTextString(m) }
func (*CacheStage) ProtoMessage()               {}
func (*CacheStage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *CacheStage) GetMaxAge() *durationscalar.DurationScalar {
	if m != nil {
		return m.MaxAge
	}
	return nil
}

func (m *CacheStage) GetHash() string {
	if m != nil && m.Hash != nil {
		return *m.Hash
	}
	return ""
}

func init() {
	proto.RegisterType((*CacheStage)(nil), "cachestage.CacheStage")
}

func init() { proto.RegisterFile("cachestage/cachestage_.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 177 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x49, 0x4e, 0x4c, 0xce,
	0x48, 0x2d, 0x2e, 0x49, 0x4c, 0x4f, 0xd5, 0x47, 0x30, 0xe3, 0xf5, 0x0a, 0x8a, 0xf2, 0x4b, 0xf2,
	0x85, 0xb8, 0x10, 0x42, 0x52, 0x2a, 0x29, 0xa5, 0x45, 0x89, 0x25, 0x99, 0xf9, 0x79, 0xc5, 0xc9,
	0x89, 0x39, 0x89, 0x45, 0xfa, 0xa8, 0x5c, 0xa8, 0x0e, 0xa5, 0x08, 0x2e, 0x2e, 0x67, 0x90, 0x9e,
	0x60, 0x90, 0x1e, 0x21, 0x33, 0x2e, 0xb6, 0xdc, 0xc4, 0x0a, 0xc7, 0xf4, 0x54, 0x09, 0x46, 0x05,
	0x46, 0x0d, 0x6e, 0x23, 0x39, 0x3d, 0x54, 0x5d, 0x7a, 0x2e, 0x50, 0x6e, 0x30, 0x98, 0x1b, 0x04,
	0x55, 0x2d, 0x24, 0xc4, 0xc5, 0x92, 0x91, 0x58, 0x9c, 0x21, 0xc1, 0xa4, 0xc0, 0xa8, 0xc1, 0x19,
	0x04, 0x66, 0x3b, 0x59, 0x44, 0x99, 0xa5, 0x67, 0x96, 0x64, 0x94, 0x26, 0xe9, 0x25, 0xe7, 0xe7,
	0xea, 0x1b, 0x19, 0x16, 0x97, 0x64, 0xe6, 0xeb, 0xa7, 0xe7, 0xeb, 0x16, 0xe4, 0x24, 0x56, 0xa6,
	0x17, 0xe5, 0x97, 0xe6, 0xa5, 0xe8, 0xa7, 0x17, 0x15, 0x24, 0xeb, 0x17, 0x27, 0x67, 0xa4, 0xe6,
	0x26, 0x22, 0x79, 0x06, 0x10, 0x00, 0x00, 0xff, 0xff, 0x0d, 0x0c, 0x2c, 0xe2, 0xe4, 0x00, 0x00,
	0x00,
}