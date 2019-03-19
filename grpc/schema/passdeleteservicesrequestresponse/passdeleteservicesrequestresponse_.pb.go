// Code generated by protoc-gen-go. DO NOT EDIT.
// source: passdeleteservicesrequestresponse/passdeleteservicesrequestresponse_.proto

/*
Package passdeleteservicesrequestresponse is a generated protocol buffer package.

It is generated from these files:
	passdeleteservicesrequestresponse/passdeleteservicesrequestresponse_.proto

It has these top-level messages:
	PassDeleteServicesRequestResponse
*/
package passdeleteservicesrequestresponse

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import responsemeta "github.com/21stio/go-playground/grpc/schema/responsemeta"
import deleteservicesrequest "github.com/21stio/go-playground/grpc/schema/deleteservicesrequest"
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

type PassDeleteServicesRequestResponse struct {
	Meta                  *responsemeta.ResponseMeta                   `protobuf:"bytes,1,opt,name=meta" json:"meta,omitempty"`
	DeleteServicesRequest *deleteservicesrequest.DeleteServicesRequest `protobuf:"bytes,2,opt,name=deleteServicesRequest" json:"deleteServicesRequest,omitempty"`
	Errors                []*error.Error                               `protobuf:"bytes,3,rep,name=errors" json:"errors,omitempty"`
	Hash                  *string                                      `protobuf:"bytes,4,opt,name=hash" json:"hash,omitempty"`
	XXX_unrecognized      []byte                                       `json:"-"`
}

func (m *PassDeleteServicesRequestResponse) Reset()         { *m = PassDeleteServicesRequestResponse{} }
func (m *PassDeleteServicesRequestResponse) String() string { return proto.CompactTextString(m) }
func (*PassDeleteServicesRequestResponse) ProtoMessage()    {}
func (*PassDeleteServicesRequestResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{0}
}

func (m *PassDeleteServicesRequestResponse) GetMeta() *responsemeta.ResponseMeta {
	if m != nil {
		return m.Meta
	}
	return nil
}

func (m *PassDeleteServicesRequestResponse) GetDeleteServicesRequest() *deleteservicesrequest.DeleteServicesRequest {
	if m != nil {
		return m.DeleteServicesRequest
	}
	return nil
}

func (m *PassDeleteServicesRequestResponse) GetErrors() []*error.Error {
	if m != nil {
		return m.Errors
	}
	return nil
}

func (m *PassDeleteServicesRequestResponse) GetHash() string {
	if m != nil && m.Hash != nil {
		return *m.Hash
	}
	return ""
}

func init() {
	proto.RegisterType((*PassDeleteServicesRequestResponse)(nil), "passdeleteservicesrequestresponse.PassDeleteServicesRequestResponse")
}

func init() {
	proto.RegisterFile("passdeleteservicesrequestresponse/passdeleteservicesrequestresponse_.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 259 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x91, 0xc1, 0x4a, 0xc4, 0x30,
	0x10, 0x86, 0xa9, 0x5b, 0x04, 0xb3, 0x9e, 0x02, 0x42, 0xe9, 0xa9, 0x2b, 0x1e, 0x7a, 0xd0, 0x09,
	0xf6, 0x11, 0x44, 0x2f, 0xc2, 0xa2, 0xd4, 0x9b, 0x17, 0xc9, 0xb6, 0x43, 0x5b, 0xd8, 0x6e, 0xe2,
	0x4c, 0x2a, 0xf8, 0xd8, 0xbe, 0x81, 0x6c, 0x36, 0x85, 0x3d, 0x04, 0x7a, 0x19, 0x66, 0xa6, 0xff,
	0xff, 0xcd, 0x5f, 0x22, 0x5e, 0xad, 0x66, 0x6e, 0x71, 0x8f, 0x0e, 0x19, 0xe9, 0x67, 0x68, 0x90,
	0x09, 0xbf, 0x27, 0x64, 0x47, 0xc8, 0xd6, 0x1c, 0x18, 0xd5, 0xa2, 0xe2, 0x0b, 0x2c, 0x19, 0x67,
	0xe4, 0x66, 0x51, 0x99, 0x17, 0x73, 0x37, 0xa2, 0xd3, 0xea, 0x7c, 0x08, 0x90, 0xbc, 0x8a, 0x02,
	0x54, 0x74, 0x3b, 0x7b, 0x24, 0x12, 0x19, 0x52, 0xbe, 0x86, 0xdd, 0xed, 0x5f, 0x22, 0x36, 0xef,
	0x9a, 0xf9, 0xd9, 0x1b, 0x3f, 0x82, 0xb1, 0x3e, 0x19, 0xeb, 0x70, 0x58, 0x82, 0x48, 0x8f, 0xc7,
	0xb3, 0xa4, 0x48, 0xca, 0x75, 0x95, 0xc3, 0x79, 0x22, 0x98, 0x55, 0x5b, 0x74, 0xba, 0xf6, 0x3a,
	0xb9, 0x13, 0x37, 0x6d, 0x0c, 0x98, 0x5d, 0x78, 0xc0, 0x3d, 0x44, 0x73, 0x42, 0x3c, 0x44, 0x1c,
	0x25, 0xef, 0xc4, 0xa5, 0xff, 0x13, 0xce, 0x56, 0xc5, 0xaa, 0x5c, 0x57, 0xd7, 0xe0, 0x47, 0x78,
	0x39, 0xd6, 0x3a, 0x7c, 0x93, 0x52, 0xa4, 0xbd, 0xe6, 0x3e, 0x4b, 0x8b, 0xa4, 0xbc, 0xaa, 0x7d,
	0xff, 0xf4, 0xf6, 0xb9, 0xed, 0x06, 0xd7, 0x4f, 0x3b, 0x68, 0xcc, 0xa8, 0xaa, 0x47, 0x76, 0x83,
	0x51, 0x9d, 0x79, 0xb0, 0x7b, 0xfd, 0xdb, 0x91, 0x99, 0x0e, 0xad, 0xea, 0xc8, 0x36, 0x8a, 0x9b,
	0x1e, 0x47, 0xbd, 0xfc, 0xb0, 0xff, 0x01, 0x00, 0x00, 0xff, 0xff, 0x4f, 0xad, 0x38, 0xea, 0x1e,
	0x02, 0x00, 0x00,
}