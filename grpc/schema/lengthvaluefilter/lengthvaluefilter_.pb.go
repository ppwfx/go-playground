// Code generated by protoc-gen-go. DO NOT EDIT.
// source: lengthvaluefilter/lengthvaluefilter_.proto

/*
Package lengthvaluefilter is a generated protocol buffer package.

It is generated from these files:
	lengthvaluefilter/lengthvaluefilter_.proto

It has these top-level messages:
	LengthValueFilter
*/
package lengthvaluefilter

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import valuekindfilter "github.com/21stio/go-playground/grpc/schema/valuekindfilter"
import lengthunitfilter "github.com/21stio/go-playground/grpc/schema/lengthunitfilter"
import floatrangefilter "github.com/21stio/go-playground/grpc/schema/floatrangefilter"
import floatfilter "github.com/21stio/go-playground/grpc/schema/floatfilter"
import boolfilter "github.com/21stio/go-playground/grpc/schema/boolfilter"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type LengthValueFilter struct {
	Kind             *valuekindfilter.ValueKindFilter   `protobuf:"bytes,1,opt,name=kind" json:"kind,omitempty"`
	Unit             *lengthunitfilter.LengthUnitFilter `protobuf:"bytes,2,opt,name=unit" json:"unit,omitempty"`
	Range            *floatrangefilter.FloatRangeFilter `protobuf:"bytes,3,opt,name=range" json:"range,omitempty"`
	Value            *floatfilter.FloatFilter           `protobuf:"bytes,4,opt,name=value" json:"value,omitempty"`
	IsEstimate       *boolfilter.BoolFilter             `protobuf:"bytes,5,opt,name=isEstimate" json:"isEstimate,omitempty"`
	And              []*LengthValueFilter               `protobuf:"bytes,6,rep,name=and" json:"and,omitempty"`
	Or               []*LengthValueFilter               `protobuf:"bytes,7,rep,name=or" json:"or,omitempty"`
	Not              []*LengthValueFilter               `protobuf:"bytes,8,rep,name=not" json:"not,omitempty"`
	Hash             *string                            `protobuf:"bytes,9,opt,name=hash" json:"hash,omitempty"`
	XXX_unrecognized []byte                             `json:"-"`
}

func (m *LengthValueFilter) Reset()                    { *m = LengthValueFilter{} }
func (m *LengthValueFilter) String() string            { return proto.CompactTextString(m) }
func (*LengthValueFilter) ProtoMessage()               {}
func (*LengthValueFilter) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *LengthValueFilter) GetKind() *valuekindfilter.ValueKindFilter {
	if m != nil {
		return m.Kind
	}
	return nil
}

func (m *LengthValueFilter) GetUnit() *lengthunitfilter.LengthUnitFilter {
	if m != nil {
		return m.Unit
	}
	return nil
}

func (m *LengthValueFilter) GetRange() *floatrangefilter.FloatRangeFilter {
	if m != nil {
		return m.Range
	}
	return nil
}

func (m *LengthValueFilter) GetValue() *floatfilter.FloatFilter {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *LengthValueFilter) GetIsEstimate() *boolfilter.BoolFilter {
	if m != nil {
		return m.IsEstimate
	}
	return nil
}

func (m *LengthValueFilter) GetAnd() []*LengthValueFilter {
	if m != nil {
		return m.And
	}
	return nil
}

func (m *LengthValueFilter) GetOr() []*LengthValueFilter {
	if m != nil {
		return m.Or
	}
	return nil
}

func (m *LengthValueFilter) GetNot() []*LengthValueFilter {
	if m != nil {
		return m.Not
	}
	return nil
}

func (m *LengthValueFilter) GetHash() string {
	if m != nil && m.Hash != nil {
		return *m.Hash
	}
	return ""
}

func init() {
	proto.RegisterType((*LengthValueFilter)(nil), "lengthvaluefilter.LengthValueFilter")
}

func init() { proto.RegisterFile("lengthvaluefilter/lengthvaluefilter_.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 357 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0x41, 0x4b, 0xfb, 0x30,
	0x18, 0xc6, 0xd9, 0xda, 0xfd, 0xff, 0x2e, 0x3b, 0x2d, 0x07, 0x09, 0x43, 0xa4, 0x0c, 0x91, 0x22,
	0x98, 0xe2, 0x18, 0xc3, 0x8b, 0x08, 0x03, 0x77, 0xd1, 0x53, 0x41, 0x0f, 0x5e, 0x24, 0x5b, 0xbb,
	0x36, 0x98, 0xe5, 0x1d, 0x69, 0x2a, 0xf8, 0x61, 0xfd, 0x2e, 0x92, 0x34, 0x75, 0x5d, 0x7b, 0xda,
	0x2d, 0x79, 0x9f, 0xe7, 0xc7, 0xfb, 0x3c, 0x21, 0xe8, 0x46, 0xa4, 0x32, 0xd3, 0xf9, 0x17, 0x13,
	0x65, 0xba, 0xe5, 0x42, 0xa7, 0x2a, 0xea, 0x4c, 0x3e, 0xe8, 0x5e, 0x81, 0x06, 0x3c, 0xee, 0x28,
	0x93, 0x6b, 0x7b, 0xf9, 0xe4, 0x32, 0x71, 0x70, 0xeb, 0xee, 0xd0, 0x49, 0x58, 0xa1, 0xa5, 0xe4,
	0xfa, 0x68, 0xcb, 0x61, 0xf0, 0xe7, 0xdc, 0x0a, 0x60, 0x5a, 0x31, 0x99, 0xd5, 0x79, 0xda, 0x83,
	0xda, 0x79, 0x69, 0x85, 0xa6, 0xe9, 0x58, 0xbf, 0x58, 0x03, 0x08, 0x27, 0x1f, 0x8e, 0x4e, 0x9d,
	0xfe, 0x78, 0x68, 0xfc, 0x62, 0x33, 0xbc, 0x99, 0xc8, 0x2b, 0x2b, 0xe2, 0x39, 0xf2, 0x4d, 0x78,
	0xd2, 0x0b, 0x7a, 0xe1, 0x68, 0x16, 0xd0, 0x56, 0x1d, 0x6a, 0xbd, 0xcf, 0x5c, 0x26, 0x95, 0x3f,
	0xb6, 0x6e, 0xbc, 0x40, 0xbe, 0x29, 0x42, 0xfa, 0x96, 0x9a, 0xd2, 0x76, 0x37, 0x5a, 0x2d, 0x7a,
	0x95, 0x5c, 0xd7, 0x9c, 0x11, 0xf1, 0x3d, 0x1a, 0xd8, 0x5e, 0xc4, 0x73, 0x60, 0xbb, 0x2a, 0x5d,
	0x99, 0x41, 0x6c, 0x06, 0x0e, 0xac, 0x00, 0x4c, 0xd1, 0xc0, 0x46, 0x23, 0xbe, 0x25, 0x09, 0x6d,
	0xf4, 0xaf, 0xa0, 0xda, 0x6f, 0x6d, 0x78, 0x81, 0x10, 0x2f, 0x9e, 0x0a, 0xcd, 0x77, 0x4c, 0xa7,
	0x64, 0x60, 0xa1, 0x73, 0x7a, 0x78, 0x15, 0xba, 0x04, 0x10, 0x0e, 0x69, 0x38, 0xf1, 0x02, 0x79,
	0x4c, 0x26, 0xe4, 0x5f, 0xe0, 0x85, 0xa3, 0xd9, 0x15, 0xed, 0x7c, 0x00, 0xda, 0x79, 0xc2, 0xd8,
	0x00, 0x78, 0x8e, 0xfa, 0xa0, 0xc8, 0xff, 0x13, 0xb0, 0x3e, 0x28, 0xb3, 0x4d, 0x82, 0x26, 0x67,
	0xa7, 0x6c, 0x93, 0xa0, 0x31, 0x46, 0x7e, 0xce, 0x8a, 0x9c, 0x0c, 0x83, 0x5e, 0x38, 0x8c, 0xed,
	0x79, 0xf9, 0xf8, 0xfe, 0x90, 0x71, 0x9d, 0x97, 0x6b, 0xba, 0x81, 0x5d, 0x34, 0xbb, 0x2b, 0x34,
	0x87, 0x28, 0x83, 0xdb, 0xbd, 0x60, 0xdf, 0x99, 0x82, 0x52, 0x26, 0x51, 0xa6, 0xf6, 0x9b, 0xa8,
	0xd8, 0xe4, 0xe9, 0x8e, 0x75, 0x3f, 0xfd, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x70, 0xc2, 0xf4,
	0x7e, 0x1a, 0x03, 0x00, 0x00,
}
