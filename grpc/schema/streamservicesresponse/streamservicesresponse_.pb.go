// Code generated by protoc-gen-go. DO NOT EDIT.
// source: streamservicesresponse/streamservicesresponse_.proto

/*
Package streamservicesresponse is a generated protocol buffer package.

It is generated from these files:
	streamservicesresponse/streamservicesresponse_.proto

It has these top-level messages:
	StreamServicesResponse
*/
package streamservicesresponse

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import responsemeta "github.com/21stio/go-playground/grpc/schema/responsemeta"
import service "github.com/21stio/go-playground/grpc/schema/service"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type StreamServicesResponse struct {
	Meta             *responsemeta.ResponseMeta `protobuf:"bytes,1,opt,name=meta" json:"meta,omitempty"`
	TotalCount       *int32                     `protobuf:"varint,2,opt,name=totalCount" json:"totalCount,omitempty"`
	Services         []*service.Service         `protobuf:"bytes,3,rep,name=services" json:"services,omitempty"`
	Hash             *string                    `protobuf:"bytes,4,opt,name=hash" json:"hash,omitempty"`
	XXX_unrecognized []byte                     `json:"-"`
}

func (m *StreamServicesResponse) Reset()                    { *m = StreamServicesResponse{} }
func (m *StreamServicesResponse) String() string            { return proto.CompactTextString(m) }
func (*StreamServicesResponse) ProtoMessage()               {}
func (*StreamServicesResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *StreamServicesResponse) GetMeta() *responsemeta.ResponseMeta {
	if m != nil {
		return m.Meta
	}
	return nil
}

func (m *StreamServicesResponse) GetTotalCount() int32 {
	if m != nil && m.TotalCount != nil {
		return *m.TotalCount
	}
	return 0
}

func (m *StreamServicesResponse) GetServices() []*service.Service {
	if m != nil {
		return m.Services
	}
	return nil
}

func (m *StreamServicesResponse) GetHash() string {
	if m != nil && m.Hash != nil {
		return *m.Hash
	}
	return ""
}

func init() {
	proto.RegisterType((*StreamServicesResponse)(nil), "streamservicesresponse.StreamServicesResponse")
}

func init() {
	proto.RegisterFile("streamservicesresponse/streamservicesresponse_.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 233 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x90, 0xb1, 0x4a, 0xc5, 0x30,
	0x14, 0x86, 0x89, 0xb7, 0x82, 0xe6, 0x2e, 0x92, 0xa1, 0x94, 0x0e, 0x12, 0x9c, 0x3a, 0x68, 0x82,
	0xc5, 0x27, 0xb8, 0xba, 0xba, 0xe4, 0x6e, 0x2e, 0x12, 0x6b, 0x68, 0x0b, 0xb7, 0x3d, 0x21, 0xe7,
	0x54, 0xf0, 0x95, 0x7c, 0x4a, 0xb1, 0x9c, 0x4a, 0x87, 0xde, 0x29, 0x27, 0x5f, 0xfe, 0xff, 0x4b,
	0x88, 0x7c, 0x42, 0x4a, 0xc1, 0x0f, 0x18, 0xd2, 0x57, 0xdf, 0x04, 0x4c, 0x01, 0x23, 0x8c, 0x18,
	0xec, 0x36, 0x7e, 0x37, 0x31, 0x01, 0x81, 0xca, 0xb7, 0x8f, 0x4b, 0xbd, 0x4c, 0x43, 0x20, 0x6f,
	0xd7, 0x1b, 0x6e, 0x96, 0x39, 0x77, 0x2c, 0xaf, 0xcc, 0xef, 0x7e, 0x84, 0xcc, 0x8f, 0xb3, 0xf4,
	0xc8, 0x52, 0xc7, 0x6d, 0x65, 0x64, 0xf6, 0x67, 0x28, 0x84, 0x16, 0xd5, 0xbe, 0x2e, 0xcd, 0x5a,
	0x6b, 0x96, 0xd4, 0x6b, 0x20, 0xef, 0xe6, 0x9c, 0xba, 0x95, 0x92, 0x80, 0xfc, 0xe9, 0x19, 0xa6,
	0x91, 0x8a, 0x0b, 0x2d, 0xaa, 0x4b, 0xb7, 0x22, 0xea, 0x5e, 0x5e, 0x2d, 0x0f, 0x2f, 0x76, 0x7a,
	0x57, 0xed, 0xeb, 0x1b, 0xc3, 0xc0, 0xf0, 0xe5, 0xee, 0x3f, 0xa1, 0x94, 0xcc, 0x3a, 0x8f, 0x5d,
	0x91, 0x69, 0x51, 0x5d, 0xbb, 0x79, 0x3e, 0xbc, 0xbc, 0x1d, 0xda, 0x9e, 0xba, 0xe9, 0xc3, 0x34,
	0x30, 0xd8, 0xfa, 0x11, 0xa9, 0x07, 0xdb, 0xc2, 0x43, 0x3c, 0xf9, 0xef, 0x36, 0xc1, 0x34, 0x7e,
	0xda, 0x36, 0xc5, 0xc6, 0x62, 0xd3, 0x85, 0xc1, 0x9f, 0xf9, 0xcb, 0xdf, 0x00, 0x00, 0x00, 0xff,
	0xff, 0x6b, 0x0a, 0x4c, 0x5d, 0x7b, 0x01, 0x00, 0x00,
}
