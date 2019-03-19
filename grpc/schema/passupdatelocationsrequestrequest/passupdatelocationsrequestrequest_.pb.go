// Code generated by protoc-gen-go. DO NOT EDIT.
// source: passupdatelocationsrequestrequest/passupdatelocationsrequestrequest_.proto

/*
Package passupdatelocationsrequestrequest is a generated protocol buffer package.

It is generated from these files:
	passupdatelocationsrequestrequest/passupdatelocationsrequestrequest_.proto

It has these top-level messages:
	PassUpdateLocationsRequestRequest
*/
package passupdatelocationsrequestrequest

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import requestmeta "github.com/21stio/go-playground/grpc/schema/requestmeta"
import servicefilter "github.com/21stio/go-playground/grpc/schema/servicefilter"
import updatelocationsrequest "github.com/21stio/go-playground/grpc/schema/updatelocationsrequest"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type PassUpdateLocationsRequestRequest struct {
	Meta                   *requestmeta.RequestMeta                       `protobuf:"bytes,1,opt,name=meta" json:"meta,omitempty"`
	ServiceFilter          *servicefilter.ServiceFilter                   `protobuf:"bytes,2,opt,name=serviceFilter" json:"serviceFilter,omitempty"`
	UpdateLocationsRequest *updatelocationsrequest.UpdateLocationsRequest `protobuf:"bytes,3,opt,name=updateLocationsRequest" json:"updateLocationsRequest,omitempty"`
	Hash                   *string                                        `protobuf:"bytes,4,opt,name=hash" json:"hash,omitempty"`
	XXX_unrecognized       []byte                                         `json:"-"`
}

func (m *PassUpdateLocationsRequestRequest) Reset()         { *m = PassUpdateLocationsRequestRequest{} }
func (m *PassUpdateLocationsRequestRequest) String() string { return proto.CompactTextString(m) }
func (*PassUpdateLocationsRequestRequest) ProtoMessage()    {}
func (*PassUpdateLocationsRequestRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{0}
}

func (m *PassUpdateLocationsRequestRequest) GetMeta() *requestmeta.RequestMeta {
	if m != nil {
		return m.Meta
	}
	return nil
}

func (m *PassUpdateLocationsRequestRequest) GetServiceFilter() *servicefilter.ServiceFilter {
	if m != nil {
		return m.ServiceFilter
	}
	return nil
}

func (m *PassUpdateLocationsRequestRequest) GetUpdateLocationsRequest() *updatelocationsrequest.UpdateLocationsRequest {
	if m != nil {
		return m.UpdateLocationsRequest
	}
	return nil
}

func (m *PassUpdateLocationsRequestRequest) GetHash() string {
	if m != nil && m.Hash != nil {
		return *m.Hash
	}
	return ""
}

func init() {
	proto.RegisterType((*PassUpdateLocationsRequestRequest)(nil), "passupdatelocationsrequestrequest.PassUpdateLocationsRequestRequest")
}

func init() {
	proto.RegisterFile("passupdatelocationsrequestrequest/passupdatelocationsrequestrequest_.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 265 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0xc1, 0x4a, 0x03, 0x31,
	0x10, 0x86, 0xd9, 0xba, 0x17, 0x23, 0x5e, 0x72, 0x90, 0xa5, 0x88, 0xb4, 0x3d, 0xf5, 0xa0, 0x09,
	0x16, 0x9f, 0xa0, 0x07, 0x0f, 0x62, 0x51, 0x56, 0xbc, 0x78, 0x91, 0x98, 0x4e, 0x77, 0x03, 0xbb,
	0x4d, 0xcc, 0x4c, 0x04, 0x1f, 0xc3, 0x37, 0x16, 0x63, 0x0a, 0xbb, 0xb0, 0x65, 0x4f, 0x33, 0x09,
	0xdf, 0xfc, 0xf9, 0x06, 0xc2, 0x1e, 0x9c, 0x42, 0x0c, 0x6e, 0xab, 0x08, 0x1a, 0xab, 0x15, 0x19,
	0xbb, 0x47, 0x0f, 0x9f, 0x01, 0x90, 0x52, 0x91, 0xa3, 0xc4, 0xbb, 0x70, 0xde, 0x92, 0xe5, 0xf3,
	0x51, 0x72, 0x7a, 0x95, 0x9a, 0x16, 0x48, 0xc9, 0x4e, 0x9f, 0x22, 0xa6, 0x0b, 0x04, 0xff, 0x65,
	0x34, 0xec, 0x4c, 0x43, 0xe0, 0x65, 0xef, 0x74, 0x60, 0xee, 0x86, 0x9f, 0x90, 0xc3, 0xd7, 0x69,
	0x6a, 0xf1, 0x33, 0x61, 0xf3, 0x67, 0x85, 0xf8, 0x1a, 0xa9, 0xc7, 0x03, 0x55, 0xfe, 0x53, 0xa9,
	0xf0, 0x6b, 0x96, 0xff, 0xe9, 0x14, 0xd9, 0x2c, 0x5b, 0x9e, 0xad, 0x0a, 0xd1, 0x51, 0x14, 0x89,
	0xd9, 0x00, 0xa9, 0x32, 0x52, 0x7c, 0xcd, 0xce, 0x93, 0xe1, 0x7d, 0x34, 0x2c, 0x26, 0x71, 0xec,
	0x52, 0xf4, 0xbc, 0xc5, 0x4b, 0x97, 0x29, 0xfb, 0x23, 0x7c, 0xc7, 0x2e, 0xc2, 0xa0, 0x52, 0x71,
	0x12, 0xc3, 0x84, 0x18, 0xde, 0x4b, 0x1c, 0x59, 0xe4, 0x48, 0x1a, 0xe7, 0x2c, 0xaf, 0x15, 0xd6,
	0x45, 0x3e, 0xcb, 0x96, 0xa7, 0x65, 0xec, 0xd7, 0x4f, 0x6f, 0x9b, 0xca, 0x50, 0x1d, 0x3e, 0x84,
	0xb6, 0xad, 0x5c, 0xdd, 0x22, 0x19, 0x2b, 0x2b, 0x7b, 0xe3, 0x1a, 0xf5, 0x5d, 0x79, 0x1b, 0xf6,
	0x5b, 0x59, 0x79, 0xa7, 0x25, 0xea, 0x1a, 0x5a, 0x35, 0xfe, 0x11, 0x7e, 0x03, 0x00, 0x00, 0xff,
	0xff, 0xbf, 0x56, 0x60, 0xc2, 0x4e, 0x02, 0x00, 0x00,
}
