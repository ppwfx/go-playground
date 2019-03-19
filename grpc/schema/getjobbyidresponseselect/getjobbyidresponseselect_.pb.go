// Code generated by protoc-gen-go. DO NOT EDIT.
// source: getjobbyidresponseselect/getjobbyidresponseselect_.proto

/*
Package getjobbyidresponseselect is a generated protocol buffer package.

It is generated from these files:
	getjobbyidresponseselect/getjobbyidresponseselect_.proto

It has these top-level messages:
	GetJobByIdResponseSelect
*/
package getjobbyidresponseselect

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import responsemetaselect "github.com/21stio/go-playground/grpc/schema/responsemetaselect"
import jobselect "github.com/21stio/go-playground/grpc/schema/jobselect"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type GetJobByIdResponseSelect struct {
	SelectAll        *bool                                  `protobuf:"varint,1,opt,name=selectAll" json:"selectAll,omitempty"`
	Meta             *responsemetaselect.ResponseMetaSelect `protobuf:"bytes,2,opt,name=meta" json:"meta,omitempty"`
	Job              *jobselect.JobSelect                   `protobuf:"bytes,3,opt,name=job" json:"job,omitempty"`
	SelectHash       *bool                                  `protobuf:"varint,4,opt,name=selectHash" json:"selectHash,omitempty"`
	Hash             *string                                `protobuf:"bytes,5,opt,name=hash" json:"hash,omitempty"`
	XXX_unrecognized []byte                                 `json:"-"`
}

func (m *GetJobByIdResponseSelect) Reset()                    { *m = GetJobByIdResponseSelect{} }
func (m *GetJobByIdResponseSelect) String() string            { return proto.CompactTextString(m) }
func (*GetJobByIdResponseSelect) ProtoMessage()               {}
func (*GetJobByIdResponseSelect) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *GetJobByIdResponseSelect) GetSelectAll() bool {
	if m != nil && m.SelectAll != nil {
		return *m.SelectAll
	}
	return false
}

func (m *GetJobByIdResponseSelect) GetMeta() *responsemetaselect.ResponseMetaSelect {
	if m != nil {
		return m.Meta
	}
	return nil
}

func (m *GetJobByIdResponseSelect) GetJob() *jobselect.JobSelect {
	if m != nil {
		return m.Job
	}
	return nil
}

func (m *GetJobByIdResponseSelect) GetSelectHash() bool {
	if m != nil && m.SelectHash != nil {
		return *m.SelectHash
	}
	return false
}

func (m *GetJobByIdResponseSelect) GetHash() string {
	if m != nil && m.Hash != nil {
		return *m.Hash
	}
	return ""
}

func init() {
	proto.RegisterType((*GetJobByIdResponseSelect)(nil), "getjobbyidresponseselect.GetJobByIdResponseSelect")
}

func init() {
	proto.RegisterFile("getjobbyidresponseselect/getjobbyidresponseselect_.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 260 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x91, 0xcf, 0x4a, 0xf4, 0x30,
	0x14, 0xc5, 0xc9, 0x37, 0xfd, 0xc0, 0x89, 0xbb, 0xe0, 0x22, 0x14, 0x91, 0xe2, 0x62, 0xe8, 0x42,
	0x13, 0x9c, 0x95, 0xb8, 0x73, 0x10, 0xff, 0x0c, 0xb8, 0xa9, 0x3b, 0x37, 0x92, 0xb4, 0x97, 0xb4,
	0xa5, 0x9d, 0x5b, 0x9a, 0xcc, 0xa2, 0x6f, 0xe9, 0x23, 0x49, 0xd3, 0xcc, 0x28, 0x0c, 0xdd, 0xdd,
	0x9c, 0xf3, 0xbb, 0x27, 0x07, 0x2e, 0xbd, 0x37, 0xe0, 0x6a, 0xd4, 0x7a, 0xa8, 0x8a, 0x1e, 0x6c,
	0x87, 0x3b, 0x0b, 0x16, 0x1a, 0xc8, 0x9d, 0x9c, 0x33, 0xbe, 0x44, 0xd7, 0xa3, 0x43, 0xc6, 0xe7,
	0x80, 0xf8, 0xe6, 0xf0, 0x6e, 0xc1, 0xa9, 0x90, 0x76, 0x2a, 0x85, 0x9c, 0x38, 0xae, 0x51, 0x07,
	0xe8, 0x38, 0x05, 0xef, 0xfa, 0x9b, 0x50, 0xfe, 0x02, 0x6e, 0x8b, 0x7a, 0x33, 0xbc, 0x15, 0x59,
	0xc8, 0xf8, 0xf0, 0x0c, 0xbb, 0xa4, 0xcb, 0x89, 0x7e, 0x6c, 0x1a, 0x4e, 0x12, 0x92, 0x9e, 0x65,
	0xbf, 0x02, 0x7b, 0xa0, 0xd1, 0xf8, 0x17, 0xff, 0x97, 0x90, 0xf4, 0x7c, 0xbd, 0x12, 0xa7, 0x05,
	0xc4, 0x21, 0xef, 0x1d, 0x9c, 0x9a, 0x32, 0x33, 0xbf, 0xc3, 0x56, 0x74, 0x51, 0xa3, 0xe6, 0x0b,
	0xbf, 0x7a, 0x21, 0x8e, 0xb5, 0xc4, 0x16, 0x75, 0x00, 0x47, 0x80, 0x5d, 0x51, 0x3a, 0x19, 0xaf,
	0xca, 0x96, 0x3c, 0xf2, 0x15, 0xfe, 0x28, 0x8c, 0xd1, 0xa8, 0x1c, 0x9d, 0xff, 0x09, 0x49, 0x97,
	0x99, 0x9f, 0x37, 0xcf, 0x9f, 0x4f, 0xa6, 0x72, 0xe5, 0x5e, 0x8b, 0x1c, 0x5b, 0xb9, 0xbe, 0xb3,
	0xae, 0x42, 0x69, 0xf0, 0xb6, 0x6b, 0xd4, 0x60, 0x7a, 0xdc, 0xef, 0x0a, 0x69, 0xfa, 0x2e, 0x97,
	0x36, 0x2f, 0xa1, 0x55, 0xb3, 0x57, 0xf8, 0x09, 0x00, 0x00, 0xff, 0xff, 0xe2, 0xbe, 0x96, 0x48,
	0xb9, 0x01, 0x00, 0x00,
}
