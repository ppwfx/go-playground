// Code generated by protoc-gen-go. DO NOT EDIT.
// source: specassertionsselect/specassertionsselect_.proto

/*
Package specassertionsselect is a generated protocol buffer package.

It is generated from these files:
	specassertionsselect/specassertionsselect_.proto

It has these top-level messages:
	SpecAssertionsSelect
*/
package specassertionsselect

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import specassertionfilter "github.com/21stio/go-playground/grpc/schema/specassertionfilter"
import specassertionsort "github.com/21stio/go-playground/grpc/schema/specassertionsort"
import page "github.com/21stio/go-playground/grpc/schema/page"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type SpecAssertionsSelect struct {
	Filter           *specassertionfilter.SpecAssertionFilter `protobuf:"bytes,1,opt,name=filter" json:"filter,omitempty"`
	Sort             *specassertionsort.SpecAssertionSort     `protobuf:"bytes,2,opt,name=sort" json:"sort,omitempty"`
	Page             *page.Page                               `protobuf:"bytes,3,opt,name=page" json:"page,omitempty"`
	SelectAll        *bool                                    `protobuf:"varint,4,opt,name=selectAll" json:"selectAll,omitempty"`
	Name             *bool                                    `protobuf:"varint,5,opt,name=name" json:"name,omitempty"`
	Pending          *bool                                    `protobuf:"varint,6,opt,name=pending" json:"pending,omitempty"`
	Passed           *bool                                    `protobuf:"varint,7,opt,name=passed" json:"passed,omitempty"`
	SelectHash       *bool                                    `protobuf:"varint,8,opt,name=selectHash" json:"selectHash,omitempty"`
	Hash             *string                                  `protobuf:"bytes,9,opt,name=hash" json:"hash,omitempty"`
	XXX_unrecognized []byte                                   `json:"-"`
}

func (m *SpecAssertionsSelect) Reset()                    { *m = SpecAssertionsSelect{} }
func (m *SpecAssertionsSelect) String() string            { return proto.CompactTextString(m) }
func (*SpecAssertionsSelect) ProtoMessage()               {}
func (*SpecAssertionsSelect) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *SpecAssertionsSelect) GetFilter() *specassertionfilter.SpecAssertionFilter {
	if m != nil {
		return m.Filter
	}
	return nil
}

func (m *SpecAssertionsSelect) GetSort() *specassertionsort.SpecAssertionSort {
	if m != nil {
		return m.Sort
	}
	return nil
}

func (m *SpecAssertionsSelect) GetPage() *page.Page {
	if m != nil {
		return m.Page
	}
	return nil
}

func (m *SpecAssertionsSelect) GetSelectAll() bool {
	if m != nil && m.SelectAll != nil {
		return *m.SelectAll
	}
	return false
}

func (m *SpecAssertionsSelect) GetName() bool {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return false
}

func (m *SpecAssertionsSelect) GetPending() bool {
	if m != nil && m.Pending != nil {
		return *m.Pending
	}
	return false
}

func (m *SpecAssertionsSelect) GetPassed() bool {
	if m != nil && m.Passed != nil {
		return *m.Passed
	}
	return false
}

func (m *SpecAssertionsSelect) GetSelectHash() bool {
	if m != nil && m.SelectHash != nil {
		return *m.SelectHash
	}
	return false
}

func (m *SpecAssertionsSelect) GetHash() string {
	if m != nil && m.Hash != nil {
		return *m.Hash
	}
	return ""
}

func init() {
	proto.RegisterType((*SpecAssertionsSelect)(nil), "specassertionsselect.SpecAssertionsSelect")
}

func init() { proto.RegisterFile("specassertionsselect/specassertionsselect_.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 308 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x92, 0x4f, 0x4b, 0xc3, 0x30,
	0x18, 0xc6, 0xe9, 0xac, 0xfb, 0xf3, 0x7a, 0x91, 0x30, 0x24, 0x0c, 0x19, 0x43, 0x3c, 0x14, 0xc1,
	0x44, 0x77, 0xf2, 0xb8, 0xed, 0x20, 0x1e, 0xa5, 0xbb, 0x79, 0x91, 0x98, 0xc5, 0xb4, 0xd0, 0x35,
	0x21, 0xc9, 0x0e, 0x7e, 0x45, 0x3f, 0x95, 0xe4, 0x6d, 0xa7, 0x56, 0x7b, 0x29, 0xef, 0xfb, 0x7b,
	0xfa, 0xbc, 0xcf, 0x43, 0x29, 0xdc, 0x79, 0xab, 0xa4, 0xf0, 0x5e, 0xb9, 0x50, 0x9a, 0xda, 0x7b,
	0x55, 0x29, 0x19, 0x78, 0x1f, 0x7c, 0x65, 0xd6, 0x99, 0x60, 0xc8, 0xb4, 0x4f, 0x9c, 0xb1, 0x0e,
	0x7d, 0x2f, 0xab, 0xa0, 0x1c, 0xef, 0x61, 0xed, 0x95, 0xd9, 0x4d, 0xf7, 0x8a, 0x71, 0x7f, 0x43,
	0x8d, 0x3b, 0x26, 0xce, 0xce, 0xad, 0xd0, 0x8a, 0xc7, 0x47, 0x4b, 0xae, 0x3e, 0x07, 0x30, 0xdd,
	0x5a, 0x25, 0xd7, 0xdf, 0xaf, 0x6f, 0xb1, 0x06, 0x59, 0xc1, 0xb0, 0xc9, 0xa1, 0xc9, 0x22, 0xc9,
	0xce, 0x96, 0x59, 0x5f, 0x2f, 0xd6, 0xb1, 0x3e, 0x22, 0xcb, 0x5b, 0x1f, 0x79, 0x80, 0x34, 0x66,
	0xd3, 0x01, 0xfa, 0xaf, 0xd9, 0xbf, 0x56, 0x5d, 0xf7, 0xd6, 0xb8, 0x90, 0xa3, 0x83, 0xcc, 0x21,
	0x8d, 0x1d, 0xe9, 0x09, 0x3a, 0x81, 0xc5, 0x85, 0x3d, 0x0b, 0xad, 0x72, 0xe4, 0xe4, 0x12, 0x26,
	0xcd, 0xc7, 0x5a, 0x57, 0x15, 0x4d, 0x17, 0x49, 0x36, 0xce, 0x7f, 0x00, 0x21, 0x90, 0xd6, 0x62,
	0xaf, 0xe8, 0x29, 0x0a, 0x38, 0x13, 0x0a, 0x23, 0xab, 0xea, 0x5d, 0x59, 0x6b, 0x3a, 0x44, 0x7c,
	0x5c, 0xc9, 0x05, 0x0c, 0x6d, 0x6c, 0xb5, 0xa3, 0x23, 0x14, 0xda, 0x8d, 0xcc, 0x01, 0x9a, 0x93,
	0x4f, 0xc2, 0x17, 0x74, 0x8c, 0xda, 0x2f, 0x12, 0x53, 0x8a, 0xa8, 0x4c, 0x16, 0x49, 0x36, 0xc9,
	0x71, 0xde, 0x6c, 0x5e, 0x56, 0xba, 0x0c, 0xc5, 0xe1, 0x8d, 0x49, 0xb3, 0xe7, 0xcb, 0x7b, 0x1f,
	0x4a, 0xc3, 0xb5, 0xb9, 0xb5, 0x95, 0xf8, 0xd0, 0xce, 0x1c, 0xea, 0x1d, 0xd7, 0xce, 0x4a, 0xee,
	0x65, 0xa1, 0xf6, 0xa2, 0xf7, 0xdf, 0xf8, 0x0a, 0x00, 0x00, 0xff, 0xff, 0x21, 0x8c, 0x7e, 0xdc,
	0x47, 0x02, 0x00, 0x00,
}
