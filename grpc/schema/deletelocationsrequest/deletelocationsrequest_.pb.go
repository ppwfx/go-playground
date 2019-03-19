// Code generated by protoc-gen-go. DO NOT EDIT.
// source: deletelocationsrequest/deletelocationsrequest_.proto

/*
Package deletelocationsrequest is a generated protocol buffer package.

It is generated from these files:
	deletelocationsrequest/deletelocationsrequest_.proto

It has these top-level messages:
	DeleteLocationsRequest
*/
package deletelocationsrequest

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import requestmeta "github.com/21stio/go-playground/grpc/schema/requestmeta"
import servicefilter "github.com/21stio/go-playground/grpc/schema/servicefilter"
import id "github.com/21stio/go-playground/grpc/schema/id"
import deletelocationsresponseselect "github.com/21stio/go-playground/grpc/schema/deletelocationsresponseselect"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type DeleteLocationsRequest struct {
	Meta             *requestmeta.RequestMeta                                     `protobuf:"bytes,1,opt,name=meta" json:"meta,omitempty"`
	ServiceFilter    *servicefilter.ServiceFilter                                 `protobuf:"bytes,2,opt,name=serviceFilter" json:"serviceFilter,omitempty"`
	Ids              []*id.Id                                                     `protobuf:"bytes,3,rep,name=ids" json:"ids,omitempty"`
	Select           *deletelocationsresponseselect.DeleteLocationsResponseSelect `protobuf:"bytes,4,opt,name=select" json:"select,omitempty"`
	Hash             *string                                                      `protobuf:"bytes,5,opt,name=hash" json:"hash,omitempty"`
	XXX_unrecognized []byte                                                       `json:"-"`
}

func (m *DeleteLocationsRequest) Reset()                    { *m = DeleteLocationsRequest{} }
func (m *DeleteLocationsRequest) String() string            { return proto.CompactTextString(m) }
func (*DeleteLocationsRequest) ProtoMessage()               {}
func (*DeleteLocationsRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *DeleteLocationsRequest) GetMeta() *requestmeta.RequestMeta {
	if m != nil {
		return m.Meta
	}
	return nil
}

func (m *DeleteLocationsRequest) GetServiceFilter() *servicefilter.ServiceFilter {
	if m != nil {
		return m.ServiceFilter
	}
	return nil
}

func (m *DeleteLocationsRequest) GetIds() []*id.Id {
	if m != nil {
		return m.Ids
	}
	return nil
}

func (m *DeleteLocationsRequest) GetSelect() *deletelocationsresponseselect.DeleteLocationsResponseSelect {
	if m != nil {
		return m.Select
	}
	return nil
}

func (m *DeleteLocationsRequest) GetHash() string {
	if m != nil && m.Hash != nil {
		return *m.Hash
	}
	return ""
}

func init() {
	proto.RegisterType((*DeleteLocationsRequest)(nil), "deletelocationsrequest.DeleteLocationsRequest")
}

func init() {
	proto.RegisterFile("deletelocationsrequest/deletelocationsrequest_.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 295 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0x4f, 0x4b, 0xf4, 0x30,
	0x10, 0xc6, 0xd9, 0x3f, 0xef, 0xc2, 0x9b, 0xd5, 0x4b, 0x0e, 0x4b, 0x58, 0x44, 0x96, 0x3d, 0xed,
	0x41, 0x13, 0x2c, 0x1e, 0x3d, 0x95, 0x45, 0x10, 0xf4, 0x92, 0xf5, 0xe4, 0x45, 0x62, 0x32, 0xb6,
	0x81, 0xb6, 0xa9, 0x49, 0x2a, 0xf8, 0x25, 0xfc, 0xcc, 0xb2, 0x69, 0x0a, 0x2d, 0x5b, 0xbd, 0xcd,
	0xcc, 0xf3, 0xcc, 0xe4, 0xf7, 0x40, 0xd0, 0xad, 0x82, 0x02, 0x3c, 0x14, 0x46, 0x0a, 0xaf, 0x4d,
	0xe5, 0x2c, 0x7c, 0x34, 0xe0, 0x3c, 0x1b, 0x1f, 0xbf, 0xd2, 0xda, 0x1a, 0x6f, 0xf0, 0x6a, 0x5c,
	0x5e, 0x5f, 0xc6, 0xa2, 0x04, 0x2f, 0x58, 0xaf, 0x8e, 0x7b, 0xeb, 0xad, 0x03, 0xfb, 0xa9, 0x25,
	0xbc, 0xeb, 0xc2, 0x83, 0x65, 0x83, 0xae, 0xf3, 0x9c, 0x69, 0xc5, 0xb4, 0xea, 0xba, 0xf4, 0xe4,
	0x25, 0x57, 0x9b, 0xca, 0x81, 0x83, 0x02, 0xe4, 0x08, 0x66, 0x5f, 0x8d, 0x37, 0xb6, 0xdf, 0x53,
	0xb4, 0xda, 0x07, 0xe3, 0x63, 0x67, 0xe4, 0x2d, 0x1b, 0xbe, 0x42, 0xf3, 0x23, 0x1f, 0x99, 0x6c,
	0x26, 0xbb, 0x65, 0x42, 0x68, 0x8f, 0x99, 0x46, 0xcf, 0x13, 0x78, 0xc1, 0x83, 0x0b, 0xa7, 0xe8,
	0x3c, 0x22, 0xdf, 0x07, 0x64, 0x32, 0x0d, 0x6b, 0x17, 0x74, 0x10, 0x84, 0x1e, 0xfa, 0x1e, 0x3e,
	0x5c, 0xc1, 0x04, 0xcd, 0xb4, 0x72, 0x64, 0xb6, 0x99, 0xed, 0x96, 0xc9, 0x82, 0x6a, 0x45, 0x1f,
	0x14, 0x3f, 0x8e, 0xf0, 0x33, 0x5a, 0xb4, 0xdc, 0x64, 0x1e, 0xce, 0xde, 0xd1, 0x3f, 0xd3, 0xd1,
	0x93, 0x48, 0xad, 0x7a, 0x08, 0x2a, 0x8f, 0xb7, 0x30, 0x46, 0xf3, 0x5c, 0xb8, 0x9c, 0xfc, 0xdb,
	0x4c, 0x76, 0xff, 0x79, 0xa8, 0xd3, 0xfd, 0x4b, 0x9a, 0x69, 0x9f, 0x37, 0x6f, 0x54, 0x9a, 0x92,
	0x25, 0x37, 0xce, 0x6b, 0xc3, 0x32, 0x73, 0x5d, 0x17, 0xe2, 0x2b, 0xb3, 0xa6, 0xa9, 0x14, 0xcb,
	0x6c, 0x2d, 0x99, 0x93, 0x39, 0x94, 0xe2, 0x97, 0xbf, 0xf0, 0x13, 0x00, 0x00, 0xff, 0xff, 0x9f,
	0x36, 0xdc, 0x6b, 0x3b, 0x02, 0x00, 0x00,
}