// Code generated by protoc-gen-go. DO NOT EDIT.
// source: passdeletefeedsrequestrequestfilter/passdeletefeedsrequestrequestfilter_.proto

/*
Package passdeletefeedsrequestrequestfilter is a generated protocol buffer package.

It is generated from these files:
	passdeletefeedsrequestrequestfilter/passdeletefeedsrequestrequestfilter_.proto

It has these top-level messages:
	PassDeleteFeedsRequestRequestFilter
*/
package passdeletefeedsrequestrequestfilter

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import requestmetafilter "github.com/21stio/go-playground/grpc/schema/requestmetafilter"
import deletefeedsrequestfilter "github.com/21stio/go-playground/grpc/schema/deletefeedsrequestfilter"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type PassDeleteFeedsRequestRequestFilter struct {
	Meta               *requestmetafilter.RequestMetaFilter               `protobuf:"bytes,1,opt,name=meta" json:"meta,omitempty"`
	DeleteFeedsRequest *deletefeedsrequestfilter.DeleteFeedsRequestFilter `protobuf:"bytes,2,opt,name=deleteFeedsRequest" json:"deleteFeedsRequest,omitempty"`
	And                []*PassDeleteFeedsRequestRequestFilter             `protobuf:"bytes,3,rep,name=and" json:"and,omitempty"`
	Or                 []*PassDeleteFeedsRequestRequestFilter             `protobuf:"bytes,4,rep,name=or" json:"or,omitempty"`
	Not                []*PassDeleteFeedsRequestRequestFilter             `protobuf:"bytes,5,rep,name=not" json:"not,omitempty"`
	Hash               *string                                            `protobuf:"bytes,6,opt,name=hash" json:"hash,omitempty"`
	XXX_unrecognized   []byte                                             `json:"-"`
}

func (m *PassDeleteFeedsRequestRequestFilter) Reset()         { *m = PassDeleteFeedsRequestRequestFilter{} }
func (m *PassDeleteFeedsRequestRequestFilter) String() string { return proto.CompactTextString(m) }
func (*PassDeleteFeedsRequestRequestFilter) ProtoMessage()    {}
func (*PassDeleteFeedsRequestRequestFilter) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{0}
}

func (m *PassDeleteFeedsRequestRequestFilter) GetMeta() *requestmetafilter.RequestMetaFilter {
	if m != nil {
		return m.Meta
	}
	return nil
}

func (m *PassDeleteFeedsRequestRequestFilter) GetDeleteFeedsRequest() *deletefeedsrequestfilter.DeleteFeedsRequestFilter {
	if m != nil {
		return m.DeleteFeedsRequest
	}
	return nil
}

func (m *PassDeleteFeedsRequestRequestFilter) GetAnd() []*PassDeleteFeedsRequestRequestFilter {
	if m != nil {
		return m.And
	}
	return nil
}

func (m *PassDeleteFeedsRequestRequestFilter) GetOr() []*PassDeleteFeedsRequestRequestFilter {
	if m != nil {
		return m.Or
	}
	return nil
}

func (m *PassDeleteFeedsRequestRequestFilter) GetNot() []*PassDeleteFeedsRequestRequestFilter {
	if m != nil {
		return m.Not
	}
	return nil
}

func (m *PassDeleteFeedsRequestRequestFilter) GetHash() string {
	if m != nil && m.Hash != nil {
		return *m.Hash
	}
	return ""
}

func init() {
	proto.RegisterType((*PassDeleteFeedsRequestRequestFilter)(nil), "passdeletefeedsrequestrequestfilter.PassDeleteFeedsRequestRequestFilter")
}

func init() {
	proto.RegisterFile("passdeletefeedsrequestrequestfilter/passdeletefeedsrequestrequestfilter_.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 280 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x92, 0x3b, 0x6f, 0xb3, 0x30,
	0x14, 0x86, 0xc5, 0xe5, 0xfb, 0xa4, 0x3a, 0x9b, 0x27, 0x2b, 0x13, 0x6a, 0x3a, 0xa0, 0x4a, 0x35,
	0x2a, 0x53, 0xe6, 0xaa, 0x8a, 0xba, 0xf4, 0xc6, 0x54, 0x65, 0xa9, 0x1c, 0x38, 0x01, 0x24, 0xe0,
	0x50, 0xfb, 0x30, 0xf4, 0x27, 0xf7, 0x5f, 0x54, 0x10, 0x6f, 0x6e, 0x24, 0x0f, 0x99, 0x0e, 0x97,
	0xe7, 0x7d, 0xce, 0x0b, 0x32, 0x7b, 0x19, 0x95, 0x31, 0x15, 0x74, 0x40, 0x70, 0x04, 0xa8, 0x8c,
	0x86, 0xaf, 0x09, 0x0c, 0xd9, 0x71, 0x6c, 0x3b, 0x02, 0x9d, 0x79, 0x30, 0x9f, 0x72, 0xd4, 0x48,
	0xc8, 0x37, 0x1e, 0xec, 0xfa, 0xd6, 0xde, 0xf6, 0x40, 0xca, 0xae, 0x70, 0x9e, 0x58, 0xe1, 0x7a,
	0xeb, 0xca, 0x6c, 0xe4, 0xdc, 0x0b, 0x9b, 0xbc, 0xfe, 0x89, 0xd8, 0xe6, 0x4d, 0x19, 0xf3, 0xb8,
	0x70, 0xbb, 0x99, 0x2b, 0x4e, 0x9c, 0x1d, 0xbb, 0x05, 0xe7, 0x5b, 0x16, 0xcf, 0x6b, 0x45, 0x90,
	0x04, 0xe9, 0x2a, 0xbf, 0x91, 0x4e, 0x15, 0x69, 0xf9, 0x67, 0x20, 0x75, 0xca, 0x14, 0x4b, 0x82,
	0x1f, 0x18, 0xaf, 0x1c, 0xb9, 0x08, 0x17, 0x4f, 0x2e, 0xcf, 0xf5, 0x93, 0x6e, 0x21, 0x6b, 0xfd,
	0xc3, 0xc6, 0xf7, 0x2c, 0x52, 0x43, 0x25, 0xa2, 0x24, 0x4a, 0x57, 0xf9, 0x93, 0xf4, 0xf8, 0xbd,
	0xd2, 0xe3, 0xa3, 0x8b, 0x59, 0xca, 0x3f, 0x58, 0x88, 0x5a, 0xc4, 0x17, 0x56, 0x87, 0xa8, 0xe7,
	0xd6, 0x03, 0x92, 0xf8, 0x77, 0xe9, 0xd6, 0x03, 0x12, 0xe7, 0x2c, 0x6e, 0x94, 0x69, 0xc4, 0xff,
	0x24, 0x48, 0xaf, 0x8a, 0xe5, 0xfa, 0xe1, 0x7d, 0xff, 0x5a, 0xb7, 0xd4, 0x4c, 0x07, 0x59, 0x62,
	0x9f, 0xe5, 0xf7, 0x86, 0x5a, 0xcc, 0x6a, 0xbc, 0x1b, 0x3b, 0xf5, 0x5d, 0x6b, 0x9c, 0x86, 0x2a,
	0xab, 0xf5, 0x58, 0x66, 0xa6, 0x6c, 0xa0, 0x57, 0x3e, 0x07, 0xfa, 0x37, 0x00, 0x00, 0xff, 0xff,
	0xbb, 0xa3, 0x9c, 0x2c, 0x1a, 0x03, 0x00, 0x00,
}
