// Code generated by protoc-gen-go. DO NOT EDIT.
// source: passdeletefeedsrequestresponse/passdeletefeedsrequestresponse_.proto

/*
Package passdeletefeedsrequestresponse is a generated protocol buffer package.

It is generated from these files:
	passdeletefeedsrequestresponse/passdeletefeedsrequestresponse_.proto

It has these top-level messages:
	PassDeleteFeedsRequestResponse
*/
package passdeletefeedsrequestresponse

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import responsemeta "github.com/21stio/go-playground/grpc/schema/responsemeta"
import deletefeedsrequest "github.com/21stio/go-playground/grpc/schema/deletefeedsrequest"
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

type PassDeleteFeedsRequestResponse struct {
	Meta               *responsemeta.ResponseMeta             `protobuf:"bytes,1,opt,name=meta" json:"meta,omitempty"`
	DeleteFeedsRequest *deletefeedsrequest.DeleteFeedsRequest `protobuf:"bytes,2,opt,name=deleteFeedsRequest" json:"deleteFeedsRequest,omitempty"`
	Errors             []*error.Error                         `protobuf:"bytes,3,rep,name=errors" json:"errors,omitempty"`
	Hash               *string                                `protobuf:"bytes,4,opt,name=hash" json:"hash,omitempty"`
	XXX_unrecognized   []byte                                 `json:"-"`
}

func (m *PassDeleteFeedsRequestResponse) Reset()                    { *m = PassDeleteFeedsRequestResponse{} }
func (m *PassDeleteFeedsRequestResponse) String() string            { return proto.CompactTextString(m) }
func (*PassDeleteFeedsRequestResponse) ProtoMessage()               {}
func (*PassDeleteFeedsRequestResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *PassDeleteFeedsRequestResponse) GetMeta() *responsemeta.ResponseMeta {
	if m != nil {
		return m.Meta
	}
	return nil
}

func (m *PassDeleteFeedsRequestResponse) GetDeleteFeedsRequest() *deletefeedsrequest.DeleteFeedsRequest {
	if m != nil {
		return m.DeleteFeedsRequest
	}
	return nil
}

func (m *PassDeleteFeedsRequestResponse) GetErrors() []*error.Error {
	if m != nil {
		return m.Errors
	}
	return nil
}

func (m *PassDeleteFeedsRequestResponse) GetHash() string {
	if m != nil && m.Hash != nil {
		return *m.Hash
	}
	return ""
}

func init() {
	proto.RegisterType((*PassDeleteFeedsRequestResponse)(nil), "passdeletefeedsrequestresponse.PassDeleteFeedsRequestResponse")
}

func init() {
	proto.RegisterFile("passdeletefeedsrequestresponse/passdeletefeedsrequestresponse_.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 255 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x51, 0xbd, 0x4a, 0xc4, 0x40,
	0x10, 0x26, 0x5e, 0x10, 0xdc, 0xb3, 0xda, 0x2a, 0xa4, 0x38, 0x82, 0x88, 0xa4, 0xd0, 0x59, 0xcc,
	0x23, 0xc8, 0x69, 0x21, 0x0a, 0xb2, 0x85, 0x85, 0x8d, 0xac, 0xc9, 0x98, 0x1c, 0x5c, 0x6e, 0xd7,
	0x99, 0x4d, 0xe1, 0xeb, 0xfa, 0x24, 0x92, 0x35, 0x81, 0x83, 0x3d, 0xce, 0x66, 0x98, 0x9f, 0xef,
	0x0f, 0x46, 0xac, 0x9d, 0x61, 0x6e, 0x70, 0x8b, 0x1e, 0x3f, 0x11, 0x1b, 0x26, 0xfc, 0x1a, 0x90,
	0x3d, 0x21, 0x3b, 0xbb, 0x63, 0x54, 0xc7, 0xcf, 0xef, 0xe0, 0xc8, 0x7a, 0x2b, 0x57, 0xc7, 0x61,
	0x79, 0x31, 0x77, 0x3d, 0x7a, 0xa3, 0xf6, 0x87, 0x49, 0x21, 0xbf, 0x8e, 0xd9, 0x2a, 0x5e, 0xcd,
	0x68, 0x89, 0x44, 0x96, 0x54, 0xa8, 0xd3, 0xee, 0xe2, 0x27, 0x11, 0xab, 0x17, 0xc3, 0xbc, 0x0e,
	0xac, 0x87, 0x91, 0xa5, 0xff, 0x58, 0x7a, 0xf2, 0x93, 0x20, 0xd2, 0xd1, 0x33, 0x4b, 0x8a, 0xa4,
	0x5c, 0x56, 0x39, 0xec, 0x07, 0x81, 0x19, 0xf5, 0x8c, 0xde, 0xe8, 0x80, 0x93, 0xaf, 0x42, 0x36,
	0x91, 0x5a, 0x76, 0x12, 0xd8, 0x57, 0x10, 0xc7, 0x83, 0x03, 0xde, 0x07, 0x14, 0xe4, 0xa5, 0x38,
	0x0d, 0xd1, 0x39, 0x5b, 0x14, 0x8b, 0x72, 0x59, 0x9d, 0x43, 0x18, 0xe1, 0x7e, 0xac, 0x7a, 0xba,
	0x49, 0x29, 0xd2, 0xce, 0x70, 0x97, 0xa5, 0x45, 0x52, 0x9e, 0xe9, 0xd0, 0xdf, 0x3d, 0xbd, 0x3d,
	0xb6, 0x1b, 0xdf, 0x0d, 0x1f, 0x50, 0xdb, 0x5e, 0x55, 0xb7, 0xec, 0x37, 0x56, 0xb5, 0xf6, 0xc6,
	0x6d, 0xcd, 0x77, 0x4b, 0x76, 0xd8, 0x35, 0xaa, 0x25, 0x57, 0x2b, 0xae, 0x3b, 0xec, 0xcd, 0x3f,
	0xdf, 0xfb, 0x0d, 0x00, 0x00, 0xff, 0xff, 0xec, 0x5d, 0xdb, 0xeb, 0xfd, 0x01, 0x00, 0x00,
}
