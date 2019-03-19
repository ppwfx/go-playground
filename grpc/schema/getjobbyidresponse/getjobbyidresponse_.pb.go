// Code generated by protoc-gen-go. DO NOT EDIT.
// source: getjobbyidresponse/getjobbyidresponse_.proto

/*
Package getjobbyidresponse is a generated protocol buffer package.

It is generated from these files:
	getjobbyidresponse/getjobbyidresponse_.proto

It has these top-level messages:
	GetJobByIdResponse
*/
package getjobbyidresponse

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import responsemeta "github.com/21stio/go-playground/grpc/schema/responsemeta"
import company "github.com/21stio/go-playground/grpc/schema/company"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type GetJobByIdResponse struct {
	Meta             *responsemeta.ResponseMeta `protobuf:"bytes,1,opt,name=meta" json:"meta,omitempty"`
	Job              *company.Job               `protobuf:"bytes,2,opt,name=job" json:"job,omitempty"`
	Hash             *string                    `protobuf:"bytes,3,opt,name=hash" json:"hash,omitempty"`
	XXX_unrecognized []byte                     `json:"-"`
}

func (m *GetJobByIdResponse) Reset()                    { *m = GetJobByIdResponse{} }
func (m *GetJobByIdResponse) String() string            { return proto.CompactTextString(m) }
func (*GetJobByIdResponse) ProtoMessage()               {}
func (*GetJobByIdResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *GetJobByIdResponse) GetMeta() *responsemeta.ResponseMeta {
	if m != nil {
		return m.Meta
	}
	return nil
}

func (m *GetJobByIdResponse) GetJob() *company.Job {
	if m != nil {
		return m.Job
	}
	return nil
}

func (m *GetJobByIdResponse) GetHash() string {
	if m != nil && m.Hash != nil {
		return *m.Hash
	}
	return ""
}

func init() {
	proto.RegisterType((*GetJobByIdResponse)(nil), "getjobbyidresponse.GetJobByIdResponse")
}

func init() { proto.RegisterFile("getjobbyidresponse/getjobbyidresponse_.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 213 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0x31, 0x6b, 0x85, 0x30,
	0x14, 0x85, 0xb1, 0xba, 0x34, 0xed, 0x94, 0xa1, 0x88, 0x43, 0x91, 0x4e, 0x0e, 0x6d, 0x42, 0xfd,
	0x01, 0xa5, 0xb8, 0x94, 0x0a, 0x5d, 0x32, 0x76, 0x29, 0x89, 0x86, 0x44, 0xa9, 0xde, 0x90, 0x44,
	0x78, 0xf9, 0xf7, 0x0f, 0x25, 0x82, 0xe0, 0x9b, 0x72, 0x72, 0xce, 0x77, 0xee, 0x85, 0x8b, 0x5e,
	0x95, 0xf4, 0x23, 0x08, 0x11, 0x86, 0xde, 0x4a, 0x67, 0x60, 0x76, 0x92, 0x9e, 0xad, 0x3f, 0x62,
	0x2c, 0x78, 0xc0, 0xf8, 0x1c, 0x15, 0xe5, 0xae, 0x26, 0xe9, 0x39, 0x3d, 0x7e, 0x62, 0xab, 0x78,
	0xea, 0x60, 0x32, 0x7c, 0x0e, 0x34, 0xbe, 0xd1, 0x7f, 0xb9, 0x20, 0xfc, 0x25, 0x7d, 0x0b, 0xa2,
	0x09, 0xdf, 0x3d, 0x8b, 0x45, 0x4c, 0x50, 0xb6, 0x96, 0xf3, 0xa4, 0x4c, 0xaa, 0x87, 0xba, 0x20,
	0xc7, 0x89, 0x64, 0xa7, 0x7e, 0xa4, 0xe7, 0x6c, 0xe3, 0xf0, 0x33, 0x4a, 0x47, 0x10, 0xf9, 0xdd,
	0x86, 0x3f, 0x92, 0xb8, 0x83, 0xb4, 0x20, 0xd8, 0x1a, 0x60, 0x8c, 0x32, 0xcd, 0x9d, 0xce, 0xd3,
	0x32, 0xa9, 0xee, 0xd9, 0xa6, 0x9b, 0xcf, 0xdf, 0x0f, 0x35, 0x78, 0xbd, 0x88, 0x15, 0xa7, 0xf5,
	0xbb, 0xf3, 0x03, 0x50, 0x05, 0x6f, 0xe6, 0x9f, 0x07, 0x65, 0x61, 0x99, 0x7b, 0xaa, 0xac, 0xe9,
	0xa8, 0xeb, 0xb4, 0x9c, 0xf8, 0x8d, 0x83, 0x5c, 0x03, 0x00, 0x00, 0xff, 0xff, 0x38, 0x71, 0x55,
	0x42, 0x38, 0x01, 0x00, 0x00,
}
