// Code generated by protoc-gen-go. DO NOT EDIT.
// source: passdeletepersonsrequestresponse/passdeletepersonsrequestresponse_.proto

/*
Package passdeletepersonsrequestresponse is a generated protocol buffer package.

It is generated from these files:
	passdeletepersonsrequestresponse/passdeletepersonsrequestresponse_.proto

It has these top-level messages:
	PassDeletePersonsRequestResponse
*/
package passdeletepersonsrequestresponse

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import responsemeta "github.com/21stio/go-playground/grpc/schema/responsemeta"
import deletepersonsrequest "github.com/21stio/go-playground/grpc/schema/deletepersonsrequest"
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

type PassDeletePersonsRequestResponse struct {
	Meta                 *responsemeta.ResponseMeta                 `protobuf:"bytes,1,opt,name=meta" json:"meta,omitempty"`
	DeletePersonsRequest *deletepersonsrequest.DeletePersonsRequest `protobuf:"bytes,2,opt,name=deletePersonsRequest" json:"deletePersonsRequest,omitempty"`
	Errors               []*error.Error                             `protobuf:"bytes,3,rep,name=errors" json:"errors,omitempty"`
	Hash                 *string                                    `protobuf:"bytes,4,opt,name=hash" json:"hash,omitempty"`
	XXX_unrecognized     []byte                                     `json:"-"`
}

func (m *PassDeletePersonsRequestResponse) Reset()         { *m = PassDeletePersonsRequestResponse{} }
func (m *PassDeletePersonsRequestResponse) String() string { return proto.CompactTextString(m) }
func (*PassDeletePersonsRequestResponse) ProtoMessage()    {}
func (*PassDeletePersonsRequestResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{0}
}

func (m *PassDeletePersonsRequestResponse) GetMeta() *responsemeta.ResponseMeta {
	if m != nil {
		return m.Meta
	}
	return nil
}

func (m *PassDeletePersonsRequestResponse) GetDeletePersonsRequest() *deletepersonsrequest.DeletePersonsRequest {
	if m != nil {
		return m.DeletePersonsRequest
	}
	return nil
}

func (m *PassDeletePersonsRequestResponse) GetErrors() []*error.Error {
	if m != nil {
		return m.Errors
	}
	return nil
}

func (m *PassDeletePersonsRequestResponse) GetHash() string {
	if m != nil && m.Hash != nil {
		return *m.Hash
	}
	return ""
}

func init() {
	proto.RegisterType((*PassDeletePersonsRequestResponse)(nil), "passdeletepersonsrequestresponse.PassDeletePersonsRequestResponse")
}

func init() {
	proto.RegisterFile("passdeletepersonsrequestresponse/passdeletepersonsrequestresponse_.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 257 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x91, 0xc1, 0x4a, 0xc4, 0x30,
	0x10, 0x86, 0xa9, 0xbb, 0x08, 0x66, 0x3d, 0x05, 0x0f, 0xa5, 0xa7, 0x20, 0x1e, 0x8a, 0x60, 0xa2,
	0x7d, 0x04, 0x51, 0xf0, 0xa0, 0xb2, 0xe4, 0xe8, 0x41, 0x89, 0xed, 0xd0, 0x2e, 0x6c, 0x3b, 0x71,
	0x26, 0x3d, 0xf8, 0xce, 0x3e, 0x84, 0x6c, 0xb6, 0x85, 0x3d, 0x04, 0x7a, 0x19, 0x66, 0xa6, 0xff,
	0xf7, 0xcf, 0x5f, 0x22, 0x5e, 0xbc, 0x63, 0x6e, 0x60, 0x0f, 0x01, 0x3c, 0x10, 0xe3, 0xc0, 0x04,
	0x3f, 0x23, 0x70, 0x20, 0x60, 0x8f, 0x03, 0x83, 0x59, 0x12, 0x7c, 0x69, 0x4f, 0x18, 0x50, 0xaa,
	0x25, 0x61, 0xa1, 0xe6, 0xae, 0x87, 0xe0, 0xcc, 0xe9, 0x30, 0x79, 0x14, 0xf7, 0x29, 0xde, 0xa4,
	0x96, 0x33, 0x21, 0x81, 0x08, 0xc9, 0xc4, 0x3a, 0xed, 0xae, 0xff, 0x32, 0xa1, 0xb6, 0x8e, 0xf9,
	0x29, 0x72, 0xdb, 0x23, 0x67, 0x8f, 0x9c, 0x9d, 0xae, 0x4a, 0x2d, 0xd6, 0x87, 0xcb, 0x79, 0xa6,
	0xb2, 0x72, 0x53, 0x15, 0xfa, 0x34, 0x8e, 0x9e, 0x55, 0x6f, 0x10, 0x9c, 0x8d, 0x3a, 0xf9, 0x29,
	0xae, 0x9a, 0x84, 0x5f, 0x7e, 0x16, 0xf9, 0x5b, 0x9d, 0x0a, 0xa9, 0x93, 0x09, 0x92, 0x3e, 0xf2,
	0x46, 0x9c, 0xc7, 0x9f, 0xe0, 0x7c, 0xa5, 0x56, 0xe5, 0xa6, 0xba, 0xd4, 0x71, 0xd4, 0xcf, 0x87,
	0x6a, 0xa7, 0x6f, 0x52, 0x8a, 0x75, 0xe7, 0xb8, 0xcb, 0xd7, 0x2a, 0x2b, 0x2f, 0x6c, 0xec, 0x1f,
	0xdf, 0x3f, 0x5e, 0xdb, 0x5d, 0xe8, 0xc6, 0x6f, 0x5d, 0x63, 0x6f, 0xaa, 0x07, 0x0e, 0x3b, 0x34,
	0x2d, 0xde, 0xf9, 0xbd, 0xfb, 0x6d, 0x09, 0xc7, 0xa1, 0x31, 0x2d, 0xf9, 0xda, 0x70, 0xdd, 0x41,
	0xef, 0x16, 0xdf, 0xf3, 0x3f, 0x00, 0x00, 0xff, 0xff, 0xb6, 0x84, 0x1c, 0xda, 0x13, 0x02, 0x00,
	0x00,
}