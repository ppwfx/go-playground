// Code generated by protoc-gen-go. DO NOT EDIT.
// source: getjobsrequestfilter/getjobsrequestfilter_.proto

/*
Package getjobsrequestfilter is a generated protocol buffer package.

It is generated from these files:
	getjobsrequestfilter/getjobsrequestfilter_.proto

It has these top-level messages:
	GetJobsRequestFilter
*/
package getjobsrequestfilter

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import requestmetafilter "github.com/21stio/go-playground/grpc/schema/requestmetafilter"
import locationqueryfilter "github.com/21stio/go-playground/grpc/schema/locationqueryfilter"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type GetJobsRequestFilter struct {
	Meta             *requestmetafilter.RequestMetaFilter     `protobuf:"bytes,1,opt,name=meta" json:"meta,omitempty"`
	LocationQuery    *locationqueryfilter.LocationQueryFilter `protobuf:"bytes,2,opt,name=locationQuery" json:"locationQuery,omitempty"`
	And              []*GetJobsRequestFilter                  `protobuf:"bytes,3,rep,name=and" json:"and,omitempty"`
	Or               []*GetJobsRequestFilter                  `protobuf:"bytes,4,rep,name=or" json:"or,omitempty"`
	Not              []*GetJobsRequestFilter                  `protobuf:"bytes,5,rep,name=not" json:"not,omitempty"`
	Hash             *string                                  `protobuf:"bytes,6,opt,name=hash" json:"hash,omitempty"`
	XXX_unrecognized []byte                                   `json:"-"`
}

func (m *GetJobsRequestFilter) Reset()                    { *m = GetJobsRequestFilter{} }
func (m *GetJobsRequestFilter) String() string            { return proto.CompactTextString(m) }
func (*GetJobsRequestFilter) ProtoMessage()               {}
func (*GetJobsRequestFilter) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *GetJobsRequestFilter) GetMeta() *requestmetafilter.RequestMetaFilter {
	if m != nil {
		return m.Meta
	}
	return nil
}

func (m *GetJobsRequestFilter) GetLocationQuery() *locationqueryfilter.LocationQueryFilter {
	if m != nil {
		return m.LocationQuery
	}
	return nil
}

func (m *GetJobsRequestFilter) GetAnd() []*GetJobsRequestFilter {
	if m != nil {
		return m.And
	}
	return nil
}

func (m *GetJobsRequestFilter) GetOr() []*GetJobsRequestFilter {
	if m != nil {
		return m.Or
	}
	return nil
}

func (m *GetJobsRequestFilter) GetNot() []*GetJobsRequestFilter {
	if m != nil {
		return m.Not
	}
	return nil
}

func (m *GetJobsRequestFilter) GetHash() string {
	if m != nil && m.Hash != nil {
		return *m.Hash
	}
	return ""
}

func init() {
	proto.RegisterType((*GetJobsRequestFilter)(nil), "getjobsrequestfilter.GetJobsRequestFilter")
}

func init() { proto.RegisterFile("getjobsrequestfilter/getjobsrequestfilter_.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 279 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0x3f, 0x4b, 0xc4, 0x30,
	0x18, 0xc6, 0xe9, 0x1f, 0x05, 0x73, 0xb8, 0x84, 0x1b, 0xc2, 0x4d, 0x45, 0x1c, 0xca, 0x81, 0xa9,
	0x76, 0x12, 0x71, 0x90, 0x1b, 0x14, 0x44, 0x05, 0x3b, 0xba, 0x48, 0xda, 0x8b, 0x69, 0xa5, 0xcd,
	0xdb, 0x4b, 0xde, 0x0e, 0xf7, 0x9d, 0xfd, 0x10, 0xd2, 0x5e, 0x04, 0x8f, 0xcb, 0x72, 0x5b, 0xf2,
	0xbc, 0xcf, 0xf3, 0x7b, 0x92, 0x10, 0x72, 0xad, 0x24, 0x7e, 0x43, 0x69, 0x8d, 0xdc, 0x0c, 0xd2,
	0xe2, 0x57, 0xd3, 0xa2, 0x34, 0x99, 0x4f, 0xfc, 0xe4, 0xbd, 0x01, 0x04, 0x3a, 0xf7, 0x0d, 0x17,
	0x4b, 0xb7, 0xed, 0x24, 0x0a, 0x07, 0x39, 0x50, 0x1c, 0x61, 0xc1, 0x5b, 0xa8, 0x04, 0x36, 0xa0,
	0x37, 0x83, 0x34, 0x5b, 0xe7, 0xf6, 0x68, 0xce, 0x7f, 0xf1, 0x13, 0x92, 0xf9, 0x93, 0xc4, 0x67,
	0x28, 0x6d, 0xb1, 0x63, 0x3e, 0x4e, 0x73, 0x7a, 0x4b, 0xe2, 0x91, 0xce, 0x82, 0x24, 0x48, 0x67,
	0xf9, 0x25, 0x3f, 0x68, 0xe4, 0xce, 0xff, 0x2a, 0x51, 0xec, 0x32, 0xc5, 0x94, 0xa0, 0x6f, 0xe4,
	0xfc, 0xaf, 0xf0, 0x7d, 0x2c, 0x64, 0xe1, 0x84, 0x48, 0x7d, 0x47, 0xe3, 0x2f, 0xff, 0x9d, 0x0e,
	0xb3, 0x1f, 0xa7, 0xf7, 0x24, 0x12, 0x7a, 0xcd, 0xa2, 0x24, 0x4a, 0x67, 0xf9, 0x92, 0xfb, 0x9e,
	0x88, 0xfb, 0xae, 0x50, 0x8c, 0x31, 0x7a, 0x47, 0x42, 0x30, 0x2c, 0x3e, 0x3a, 0x1c, 0x82, 0x19,
	0x9b, 0x35, 0x20, 0x3b, 0x39, 0xbe, 0x59, 0x03, 0x52, 0x4a, 0xe2, 0x5a, 0xd8, 0x9a, 0x9d, 0x26,
	0x41, 0x7a, 0x56, 0x4c, 0xeb, 0xd5, 0xea, 0xe3, 0x41, 0x35, 0x58, 0x0f, 0x25, 0xaf, 0xa0, 0xcb,
	0xf2, 0x1b, 0x8b, 0x0d, 0x64, 0x0a, 0xae, 0xfa, 0x56, 0x6c, 0x95, 0x81, 0x41, 0xaf, 0x33, 0x65,
	0xfa, 0x2a, 0xb3, 0x55, 0x2d, 0x3b, 0xe1, 0xfd, 0x2b, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x6b,
	0x2b, 0xde, 0xa7, 0x57, 0x02, 0x00, 0x00,
}
