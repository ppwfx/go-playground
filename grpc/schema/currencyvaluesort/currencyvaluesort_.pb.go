// Code generated by protoc-gen-go. DO NOT EDIT.
// source: currencyvaluesort/currencyvaluesort_.proto

/*
Package currencyvaluesort is a generated protocol buffer package.

It is generated from these files:
	currencyvaluesort/currencyvaluesort_.proto

It has these top-level messages:
	CurrencyValueSort
*/
package currencyvaluesort

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import floatrangesort "github.com/21stio/go-playground/grpc/schema/floatrangesort"
import sortkind "github.com/21stio/go-playground/grpc/schema/sortkind"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type CurrencyValueSort struct {
	Range            *floatrangesort.FloatRangeSort `protobuf:"bytes,1,opt,name=range" json:"range,omitempty"`
	Value            *sortkind.SortKind             `protobuf:"varint,2,opt,name=value,enum=sortkind.SortKind" json:"value,omitempty"`
	IsEstimate       *sortkind.SortKind             `protobuf:"varint,3,opt,name=isEstimate,enum=sortkind.SortKind" json:"isEstimate,omitempty"`
	Hash             *string                        `protobuf:"bytes,4,opt,name=hash" json:"hash,omitempty"`
	XXX_unrecognized []byte                         `json:"-"`
}

func (m *CurrencyValueSort) Reset()                    { *m = CurrencyValueSort{} }
func (m *CurrencyValueSort) String() string            { return proto.CompactTextString(m) }
func (*CurrencyValueSort) ProtoMessage()               {}
func (*CurrencyValueSort) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *CurrencyValueSort) GetRange() *floatrangesort.FloatRangeSort {
	if m != nil {
		return m.Range
	}
	return nil
}

func (m *CurrencyValueSort) GetValue() sortkind.SortKind {
	if m != nil && m.Value != nil {
		return *m.Value
	}
	return sortkind.SortKind_ASC
}

func (m *CurrencyValueSort) GetIsEstimate() sortkind.SortKind {
	if m != nil && m.IsEstimate != nil {
		return *m.IsEstimate
	}
	return sortkind.SortKind_ASC
}

func (m *CurrencyValueSort) GetHash() string {
	if m != nil && m.Hash != nil {
		return *m.Hash
	}
	return ""
}

func init() {
	proto.RegisterType((*CurrencyValueSort)(nil), "currencyvaluesort.CurrencyValueSort")
}

func init() { proto.RegisterFile("currencyvaluesort/currencyvaluesort_.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 242 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x90, 0x31, 0x4b, 0xc4, 0x30,
	0x14, 0xc7, 0x89, 0xde, 0x0d, 0x46, 0x10, 0x2e, 0x53, 0xb8, 0x41, 0x8a, 0x38, 0x14, 0xc1, 0x04,
	0x8b, 0xab, 0x08, 0x8a, 0x2e, 0x6e, 0x15, 0x1c, 0x5c, 0x24, 0xa6, 0x35, 0x0d, 0xb6, 0x79, 0x25,
	0x79, 0x11, 0xee, 0xb3, 0xf9, 0xe5, 0x24, 0x6d, 0x0f, 0x2c, 0xe5, 0xb6, 0xf7, 0xcf, 0xfb, 0xfd,
	0x7f, 0x09, 0xa1, 0x57, 0x3a, 0x7a, 0x5f, 0x3b, 0xbd, 0xfb, 0x51, 0x6d, 0xac, 0x03, 0x78, 0x94,
	0x8b, 0x93, 0x0f, 0xd1, 0x7b, 0x40, 0x60, 0x9b, 0xc5, 0x66, 0x7b, 0xf9, 0xd5, 0x82, 0x42, 0xaf,
	0x9c, 0x19, 0xbb, 0xf3, 0x38, 0x15, 0xb7, 0x3c, 0x85, 0x6f, 0xeb, 0x2a, 0xb9, 0x1f, 0xa6, 0xcd,
	0xc5, 0x2f, 0xa1, 0x9b, 0xc7, 0xc9, 0xfa, 0x96, 0xac, 0xaf, 0xe0, 0x91, 0xdd, 0xd2, 0xf5, 0xe0,
	0xe0, 0x24, 0x23, 0xf9, 0x69, 0x71, 0x2e, 0xe6, 0x5a, 0xf1, 0x9c, 0x62, 0x99, 0x62, 0xc2, 0xcb,
	0x11, 0x66, 0x39, 0x5d, 0x0f, 0x0f, 0xe3, 0x47, 0x19, 0xc9, 0xcf, 0x0a, 0x26, 0xf6, 0x97, 0x89,
	0x44, 0xbd, 0x58, 0x57, 0x95, 0x23, 0xc0, 0x0a, 0x4a, 0x6d, 0x78, 0x0a, 0x68, 0x3b, 0x85, 0x35,
	0x3f, 0x3e, 0x88, 0xff, 0xa3, 0x18, 0xa3, 0xab, 0x46, 0x85, 0x86, 0xaf, 0x32, 0x92, 0x9f, 0x94,
	0xc3, 0xfc, 0x70, 0xff, 0x7e, 0x67, 0x2c, 0x36, 0xf1, 0x53, 0x68, 0xe8, 0x64, 0x71, 0x13, 0xd0,
	0x82, 0x34, 0x70, 0xdd, 0xb7, 0x6a, 0x67, 0x3c, 0x44, 0x57, 0x49, 0xe3, 0x7b, 0x2d, 0x83, 0x6e,
	0xea, 0x4e, 0x2d, 0x3f, 0xf6, 0x2f, 0x00, 0x00, 0xff, 0xff, 0xc8, 0x9a, 0x3e, 0x3d, 0x7e, 0x01,
	0x00, 0x00,
}
