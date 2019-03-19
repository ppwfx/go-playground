// Code generated by protoc-gen-go. DO NOT EDIT.
// source: getjoblistbyidrequestfilter/getjoblistbyidrequestfilter_.proto

/*
Package getjoblistbyidrequestfilter is a generated protocol buffer package.

It is generated from these files:
	getjoblistbyidrequestfilter/getjoblistbyidrequestfilter_.proto

It has these top-level messages:
	GetJobListByIdRequestFilter
*/
package getjoblistbyidrequestfilter

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import requestmetafilter "github.com/21stio/go-playground/grpc/schema/requestmetafilter"
import idfilter "github.com/21stio/go-playground/grpc/schema/idfilter"
import jobskindfilter "github.com/21stio/go-playground/grpc/schema/jobskindfilter"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type GetJobListByIdRequestFilter struct {
	Meta             *requestmetafilter.RequestMetaFilter `protobuf:"bytes,1,opt,name=meta" json:"meta,omitempty"`
	Id               *idfilter.IdFilter                   `protobuf:"bytes,2,opt,name=id" json:"id,omitempty"`
	Kind             *jobskindfilter.JobsKindFilter       `protobuf:"bytes,3,opt,name=kind" json:"kind,omitempty"`
	And              []*GetJobListByIdRequestFilter       `protobuf:"bytes,4,rep,name=and" json:"and,omitempty"`
	Or               []*GetJobListByIdRequestFilter       `protobuf:"bytes,5,rep,name=or" json:"or,omitempty"`
	Not              []*GetJobListByIdRequestFilter       `protobuf:"bytes,6,rep,name=not" json:"not,omitempty"`
	Hash             *string                              `protobuf:"bytes,7,opt,name=hash" json:"hash,omitempty"`
	XXX_unrecognized []byte                               `json:"-"`
}

func (m *GetJobListByIdRequestFilter) Reset()                    { *m = GetJobListByIdRequestFilter{} }
func (m *GetJobListByIdRequestFilter) String() string            { return proto.CompactTextString(m) }
func (*GetJobListByIdRequestFilter) ProtoMessage()               {}
func (*GetJobListByIdRequestFilter) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *GetJobListByIdRequestFilter) GetMeta() *requestmetafilter.RequestMetaFilter {
	if m != nil {
		return m.Meta
	}
	return nil
}

func (m *GetJobListByIdRequestFilter) GetId() *idfilter.IdFilter {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *GetJobListByIdRequestFilter) GetKind() *jobskindfilter.JobsKindFilter {
	if m != nil {
		return m.Kind
	}
	return nil
}

func (m *GetJobListByIdRequestFilter) GetAnd() []*GetJobListByIdRequestFilter {
	if m != nil {
		return m.And
	}
	return nil
}

func (m *GetJobListByIdRequestFilter) GetOr() []*GetJobListByIdRequestFilter {
	if m != nil {
		return m.Or
	}
	return nil
}

func (m *GetJobListByIdRequestFilter) GetNot() []*GetJobListByIdRequestFilter {
	if m != nil {
		return m.Not
	}
	return nil
}

func (m *GetJobListByIdRequestFilter) GetHash() string {
	if m != nil && m.Hash != nil {
		return *m.Hash
	}
	return ""
}

func init() {
	proto.RegisterType((*GetJobListByIdRequestFilter)(nil), "getjoblistbyidrequestfilter.GetJobListByIdRequestFilter")
}

func init() {
	proto.RegisterFile("getjoblistbyidrequestfilter/getjoblistbyidrequestfilter_.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 310 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x92, 0x31, 0x6f, 0xfa, 0x30,
	0x10, 0xc5, 0x45, 0x92, 0xff, 0xbf, 0xaa, 0xd9, 0x3c, 0x59, 0x20, 0x55, 0x08, 0x31, 0xa0, 0x4a,
	0x75, 0xd4, 0x4c, 0x4c, 0x1d, 0x18, 0x4a, 0xa1, 0xed, 0x92, 0xb1, 0x4b, 0x65, 0x63, 0x37, 0x31,
	0x05, 0x1f, 0xb5, 0x8f, 0x81, 0x0f, 0xd4, 0xef, 0x59, 0xc5, 0x38, 0x03, 0x42, 0xca, 0xc2, 0x76,
	0xe7, 0xfb, 0xbd, 0x77, 0xcf, 0x96, 0xc9, 0x53, 0xa5, 0x71, 0x03, 0x72, 0x6b, 0x3c, 0xca, 0xa3,
	0x51, 0x4e, 0xff, 0x1c, 0xb4, 0xc7, 0x2f, 0xb3, 0x45, 0xed, 0xf2, 0x8e, 0xd9, 0x27, 0xdf, 0x3b,
	0x40, 0xa0, 0xc3, 0x0e, 0x66, 0x70, 0x1f, 0xdb, 0x9d, 0x46, 0x11, 0x2d, 0x2f, 0x4e, 0xa2, 0xd1,
	0x80, 0x19, 0x15, 0x91, 0xb6, 0x68, 0x27, 0x93, 0x0d, 0x48, 0xff, 0x6d, 0x6c, 0x3b, 0x3f, 0x6f,
	0x23, 0x35, 0xfe, 0x4d, 0xc9, 0x70, 0xa1, 0x71, 0x05, 0xf2, 0xcd, 0x78, 0x9c, 0x1f, 0x97, 0xaa,
	0x3c, 0xad, 0x7a, 0x0e, 0x18, 0x9d, 0x91, 0xac, 0x59, 0xca, 0x7a, 0xa3, 0xde, 0xb4, 0x5f, 0x4c,
	0xf8, 0x45, 0x10, 0x1e, 0xf9, 0x77, 0x8d, 0xe2, 0xa4, 0x29, 0x83, 0x82, 0x8e, 0x49, 0x62, 0x14,
	0x4b, 0x82, 0x8e, 0xf2, 0x36, 0x1d, 0x5f, 0xaa, 0x48, 0x25, 0x46, 0xd1, 0x82, 0x64, 0x4d, 0x24,
	0x96, 0x06, 0xea, 0x8e, 0x9f, 0x67, 0xe4, 0x2b, 0x90, 0xfe, 0xd5, 0xd8, 0x56, 0x11, 0x58, 0xba,
	0x22, 0xa9, 0xb0, 0x8a, 0x65, 0xa3, 0x74, 0xda, 0x2f, 0x66, 0xbc, 0xe3, 0x21, 0x79, 0xc7, 0xc5,
	0xca, 0xc6, 0x84, 0xbe, 0x90, 0x04, 0x1c, 0xfb, 0x77, 0xa5, 0x55, 0x02, 0xae, 0x49, 0x65, 0x01,
	0xd9, 0xff, 0x6b, 0x53, 0x59, 0x40, 0x4a, 0x49, 0x56, 0x0b, 0x5f, 0xb3, 0x9b, 0x51, 0x6f, 0x7a,
	0x5b, 0x86, 0x7a, 0xbe, 0xfc, 0x58, 0x54, 0x06, 0xeb, 0x83, 0xe4, 0x6b, 0xd8, 0xe5, 0xc5, 0xa3,
	0x47, 0x03, 0x79, 0x05, 0x0f, 0xfb, 0xad, 0x38, 0x56, 0x0e, 0x0e, 0x56, 0xe5, 0x95, 0xdb, 0xaf,
	0x73, 0xbf, 0xae, 0xf5, 0x4e, 0x74, 0x7d, 0xc1, 0xbf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x15, 0x90,
	0xdb, 0x9a, 0xbc, 0x02, 0x00, 0x00,
}