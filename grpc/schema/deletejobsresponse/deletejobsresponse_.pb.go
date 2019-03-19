// Code generated by protoc-gen-go. DO NOT EDIT.
// source: deletejobsresponse/deletejobsresponse_.proto

/*
Package deletejobsresponse is a generated protocol buffer package.

It is generated from these files:
	deletejobsresponse/deletejobsresponse_.proto

It has these top-level messages:
	DeleteJobsResponse
*/
package deletejobsresponse

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import responsemeta "github.com/21stio/go-playground/grpc/schema/responsemeta"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type DeleteJobsResponse struct {
	Meta             *responsemeta.ResponseMeta `protobuf:"bytes,1,opt,name=meta" json:"meta,omitempty"`
	Hash             *string                    `protobuf:"bytes,2,opt,name=hash" json:"hash,omitempty"`
	XXX_unrecognized []byte                     `json:"-"`
}

func (m *DeleteJobsResponse) Reset()                    { *m = DeleteJobsResponse{} }
func (m *DeleteJobsResponse) String() string            { return proto.CompactTextString(m) }
func (*DeleteJobsResponse) ProtoMessage()               {}
func (*DeleteJobsResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *DeleteJobsResponse) GetMeta() *responsemeta.ResponseMeta {
	if m != nil {
		return m.Meta
	}
	return nil
}

func (m *DeleteJobsResponse) GetHash() string {
	if m != nil && m.Hash != nil {
		return *m.Hash
	}
	return ""
}

func init() {
	proto.RegisterType((*DeleteJobsResponse)(nil), "deletejobsresponse.DeleteJobsResponse")
}

func init() { proto.RegisterFile("deletejobsresponse/deletejobsresponse_.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 178 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0x49, 0x49, 0xcd, 0x49,
	0x2d, 0x49, 0xcd, 0xca, 0x4f, 0x2a, 0x2e, 0x4a, 0x2d, 0x2e, 0xc8, 0xcf, 0x2b, 0x4e, 0xd5, 0xc7,
	0x14, 0x8a, 0xd7, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0xc2, 0x94, 0x92, 0x52, 0x80, 0xb1,
	0x72, 0x53, 0x4b, 0x12, 0xf5, 0x91, 0x39, 0x50, 0x5d, 0x4a, 0x11, 0x5c, 0x42, 0x2e, 0x60, 0x7d,
	0x5e, 0xf9, 0x49, 0xc5, 0x41, 0x50, 0x05, 0x42, 0x7a, 0x5c, 0x2c, 0x20, 0x45, 0x12, 0x8c, 0x0a,
	0x8c, 0x1a, 0xdc, 0x46, 0x52, 0x7a, 0xc8, 0x3a, 0xf5, 0x60, 0xaa, 0x7c, 0x53, 0x4b, 0x12, 0x83,
	0xc0, 0xea, 0x84, 0x84, 0xb8, 0x58, 0x32, 0x12, 0x8b, 0x33, 0x24, 0x98, 0x14, 0x18, 0x35, 0x38,
	0x83, 0xc0, 0x6c, 0x27, 0x87, 0x28, 0xbb, 0xf4, 0xcc, 0x92, 0x8c, 0xd2, 0x24, 0xbd, 0xe4, 0xfc,
	0x5c, 0x7d, 0x23, 0xc3, 0xe2, 0x92, 0xcc, 0x7c, 0xfd, 0xf4, 0x7c, 0xdd, 0x82, 0x9c, 0xc4, 0xca,
	0xf4, 0xa2, 0xfc, 0xd2, 0xbc, 0x14, 0xfd, 0xf4, 0xa2, 0x82, 0x64, 0xfd, 0xe2, 0xe4, 0x8c, 0xd4,
	0xdc, 0x44, 0x2c, 0x1e, 0x03, 0x04, 0x00, 0x00, 0xff, 0xff, 0x7b, 0x39, 0x8a, 0x5f, 0x00, 0x01,
	0x00, 0x00,
}
