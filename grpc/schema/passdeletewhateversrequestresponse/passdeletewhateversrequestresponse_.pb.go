// Code generated by protoc-gen-go. DO NOT EDIT.
// source: passdeletewhateversrequestresponse/passdeletewhateversrequestresponse_.proto

/*
Package passdeletewhateversrequestresponse is a generated protocol buffer package.

It is generated from these files:
	passdeletewhateversrequestresponse/passdeletewhateversrequestresponse_.proto

It has these top-level messages:
	PassDeleteWhateversRequestResponse
*/
package passdeletewhateversrequestresponse

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import responsemeta "github.com/21stio/go-playground/grpc/schema/responsemeta"
import deletewhateversrequest "github.com/21stio/go-playground/grpc/schema/deletewhateversrequest"
import error "github.com/21stio/go-playground/grpc/schema/error"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type PassDeleteWhateversRequestResponse struct {
	Meta                   *responsemeta.ResponseMeta                     `protobuf:"bytes,1,opt,name=meta" json:"meta,omitempty"`
	DeleteWhateversRequest *deletewhateversrequest.DeleteWhateversRequest `protobuf:"bytes,2,opt,name=deleteWhateversRequest" json:"deleteWhateversRequest,omitempty"`
	Errors                 []*error.Error                                 `protobuf:"bytes,3,rep,name=errors" json:"errors,omitempty"`
	Hash                   *string                                        `protobuf:"bytes,4,opt,name=hash" json:"hash,omitempty"`
	XXX_unrecognized       []byte                                         `json:"-"`
}

func (m *PassDeleteWhateversRequestResponse) Reset()         { *m = PassDeleteWhateversRequestResponse{} }
func (m *PassDeleteWhateversRequestResponse) String() string { return proto.CompactTextString(m) }
func (*PassDeleteWhateversRequestResponse) ProtoMessage()    {}
func (*PassDeleteWhateversRequestResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{0}
}

func (m *PassDeleteWhateversRequestResponse) GetMeta() *responsemeta.ResponseMeta {
	if m != nil {
		return m.Meta
	}
	return nil
}

func (m *PassDeleteWhateversRequestResponse) GetDeleteWhateversRequest() *deletewhateversrequest.DeleteWhateversRequest {
	if m != nil {
		return m.DeleteWhateversRequest
	}
	return nil
}

func (m *PassDeleteWhateversRequestResponse) GetErrors() []*error.Error {
	if m != nil {
		return m.Errors
	}
	return nil
}

func (m *PassDeleteWhateversRequestResponse) GetHash() string {
	if m != nil && m.Hash != nil {
		return *m.Hash
	}
	return ""
}

func init() {
	proto.RegisterType((*PassDeleteWhateversRequestResponse)(nil), "passdeletewhateversrequestresponse.PassDeleteWhateversRequestResponse")
}

func init() {
	proto.RegisterFile("passdeletewhateversrequestresponse/passdeletewhateversrequestresponse_.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 260 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0x31, 0x4b, 0xc4, 0x40,
	0x10, 0x85, 0x89, 0x17, 0x04, 0xf7, 0xac, 0xb6, 0x90, 0x90, 0x2a, 0x04, 0x8b, 0x34, 0xce, 0x62,
	0xf0, 0x17, 0x88, 0x76, 0x2a, 0xc7, 0x36, 0x82, 0x8d, 0xac, 0xc9, 0x98, 0x3d, 0xb8, 0xdc, 0xae,
	0x33, 0x1b, 0xc5, 0x3f, 0x6e, 0x2d, 0xb7, 0x26, 0x70, 0x45, 0x8e, 0xb3, 0x19, 0x66, 0x26, 0xef,
	0x7d, 0xf3, 0xc2, 0x8a, 0x07, 0x6f, 0x98, 0x5b, 0xdc, 0x60, 0xc0, 0x2f, 0x6b, 0x02, 0x7e, 0x22,
	0x31, 0xe1, 0xc7, 0x80, 0x1c, 0x08, 0xd9, 0xbb, 0x2d, 0xa3, 0x3a, 0x2e, 0x79, 0x05, 0x4f, 0x2e,
	0x38, 0x59, 0x1e, 0x97, 0xe6, 0xc5, 0xd4, 0xf5, 0x18, 0x8c, 0xda, 0x1f, 0x46, 0x4a, 0x7e, 0x33,
	0x4f, 0x50, 0xf3, 0xeb, 0xc9, 0x25, 0x91, 0xc8, 0x91, 0x8a, 0x75, 0xdc, 0x95, 0x3f, 0x89, 0x28,
	0x57, 0x86, 0xf9, 0x2e, 0x3a, 0x9f, 0x27, 0xa7, 0xfe, 0x73, 0xea, 0xf1, 0xb6, 0x04, 0x91, 0xee,
	0xee, 0x67, 0x49, 0x91, 0x54, 0xcb, 0x3a, 0x87, 0xfd, 0x50, 0x30, 0xa9, 0x1e, 0x31, 0x18, 0x1d,
	0x75, 0xf2, 0x5d, 0x5c, 0xb4, 0xb3, 0xc4, 0xec, 0x24, 0x12, 0x00, 0xe6, 0xa3, 0xc2, 0x81, 0x1c,
	0x07, 0x68, 0xf2, 0x52, 0x9c, 0xc6, 0xdf, 0xe1, 0x6c, 0x51, 0x2c, 0xaa, 0x65, 0x7d, 0x0e, 0x71,
	0x84, 0xfb, 0x5d, 0xd5, 0xe3, 0x37, 0x29, 0x45, 0x6a, 0x0d, 0xdb, 0x2c, 0x2d, 0x92, 0xea, 0x4c,
	0xc7, 0xfe, 0x76, 0xf5, 0xf2, 0xd4, 0xad, 0x83, 0x1d, 0xde, 0xa0, 0x71, 0xbd, 0xaa, 0xaf, 0x39,
	0xac, 0x9d, 0xea, 0xdc, 0x95, 0xdf, 0x98, 0xef, 0x8e, 0xdc, 0xb0, 0x6d, 0x55, 0x47, 0xbe, 0x51,
	0xdc, 0x58, 0xec, 0xcd, 0x3f, 0x5e, 0xf8, 0x37, 0x00, 0x00, 0xff, 0xff, 0x21, 0xc7, 0xf6, 0x79,
	0x29, 0x02, 0x00, 0x00,
}
