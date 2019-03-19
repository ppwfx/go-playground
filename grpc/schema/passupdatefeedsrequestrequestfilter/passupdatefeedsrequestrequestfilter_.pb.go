// Code generated by protoc-gen-go. DO NOT EDIT.
// source: passupdatefeedsrequestrequestfilter/passupdatefeedsrequestrequestfilter_.proto

/*
Package passupdatefeedsrequestrequestfilter is a generated protocol buffer package.

It is generated from these files:
	passupdatefeedsrequestrequestfilter/passupdatefeedsrequestrequestfilter_.proto

It has these top-level messages:
	PassUpdateFeedsRequestRequestFilter
*/
package passupdatefeedsrequestrequestfilter

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import requestmetafilter "github.com/21stio/go-playground/grpc/schema/requestmetafilter"
import updatefeedsrequestfilter "github.com/21stio/go-playground/grpc/schema/updatefeedsrequestfilter"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type PassUpdateFeedsRequestRequestFilter struct {
	Meta               *requestmetafilter.RequestMetaFilter               `protobuf:"bytes,1,opt,name=meta" json:"meta,omitempty"`
	UpdateFeedsRequest *updatefeedsrequestfilter.UpdateFeedsRequestFilter `protobuf:"bytes,2,opt,name=updateFeedsRequest" json:"updateFeedsRequest,omitempty"`
	And                []*PassUpdateFeedsRequestRequestFilter             `protobuf:"bytes,3,rep,name=and" json:"and,omitempty"`
	Or                 []*PassUpdateFeedsRequestRequestFilter             `protobuf:"bytes,4,rep,name=or" json:"or,omitempty"`
	Not                []*PassUpdateFeedsRequestRequestFilter             `protobuf:"bytes,5,rep,name=not" json:"not,omitempty"`
	Hash               *string                                            `protobuf:"bytes,6,opt,name=hash" json:"hash,omitempty"`
	XXX_unrecognized   []byte                                             `json:"-"`
}

func (m *PassUpdateFeedsRequestRequestFilter) Reset()         { *m = PassUpdateFeedsRequestRequestFilter{} }
func (m *PassUpdateFeedsRequestRequestFilter) String() string { return proto.CompactTextString(m) }
func (*PassUpdateFeedsRequestRequestFilter) ProtoMessage()    {}
func (*PassUpdateFeedsRequestRequestFilter) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{0}
}

func (m *PassUpdateFeedsRequestRequestFilter) GetMeta() *requestmetafilter.RequestMetaFilter {
	if m != nil {
		return m.Meta
	}
	return nil
}

func (m *PassUpdateFeedsRequestRequestFilter) GetUpdateFeedsRequest() *updatefeedsrequestfilter.UpdateFeedsRequestFilter {
	if m != nil {
		return m.UpdateFeedsRequest
	}
	return nil
}

func (m *PassUpdateFeedsRequestRequestFilter) GetAnd() []*PassUpdateFeedsRequestRequestFilter {
	if m != nil {
		return m.And
	}
	return nil
}

func (m *PassUpdateFeedsRequestRequestFilter) GetOr() []*PassUpdateFeedsRequestRequestFilter {
	if m != nil {
		return m.Or
	}
	return nil
}

func (m *PassUpdateFeedsRequestRequestFilter) GetNot() []*PassUpdateFeedsRequestRequestFilter {
	if m != nil {
		return m.Not
	}
	return nil
}

func (m *PassUpdateFeedsRequestRequestFilter) GetHash() string {
	if m != nil && m.Hash != nil {
		return *m.Hash
	}
	return ""
}

func init() {
	proto.RegisterType((*PassUpdateFeedsRequestRequestFilter)(nil), "passupdatefeedsrequestrequestfilter.PassUpdateFeedsRequestRequestFilter")
}

func init() {
	proto.RegisterFile("passupdatefeedsrequestrequestfilter/passupdatefeedsrequestrequestfilter_.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 280 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x92, 0xbb, 0x6a, 0xc3, 0x30,
	0x14, 0x86, 0xf1, 0xa5, 0x85, 0x2a, 0x9b, 0x26, 0x91, 0xc9, 0x34, 0x1d, 0x4c, 0xa1, 0x32, 0xf5,
	0x94, 0xb9, 0x43, 0xe8, 0xd2, 0x9b, 0xa1, 0x50, 0xb2, 0x14, 0xc5, 0x56, 0x6c, 0x43, 0xec, 0xe3,
	0x4a, 0x47, 0x43, 0x1f, 0xb9, 0x6f, 0x51, 0xac, 0x68, 0x53, 0x02, 0x1a, 0x32, 0x1d, 0x5f, 0xbe,
	0xff, 0x3b, 0xbf, 0x8d, 0xc8, 0xeb, 0x24, 0xb4, 0x36, 0x53, 0x23, 0x50, 0xee, 0xa5, 0x6c, 0xb4,
	0x92, 0x3f, 0x46, 0x6a, 0x74, 0x63, 0xdf, 0x1f, 0x50, 0xaa, 0x22, 0x80, 0xf9, 0xe6, 0x93, 0x02,
	0x04, 0xba, 0x0a, 0x60, 0x97, 0xf7, 0xee, 0x76, 0x90, 0x28, 0xdc, 0x0a, 0xef, 0x89, 0x13, 0x2e,
	0xd7, 0xbe, 0xcc, 0x45, 0xce, 0xbd, 0x70, 0xc9, 0xdb, 0xbf, 0x84, 0xac, 0xde, 0x85, 0xd6, 0x9f,
	0x96, 0xdb, 0xcc, 0x5c, 0x75, 0xe4, 0xdc, 0xd8, 0x58, 0x9c, 0xae, 0x49, 0x3a, 0xaf, 0x65, 0x51,
	0x16, 0xe5, 0x8b, 0xf2, 0x8e, 0x7b, 0x55, 0xb8, 0xe3, 0x5f, 0x24, 0x8a, 0x63, 0xa6, 0xb2, 0x09,
	0xba, 0x23, 0xd4, 0x78, 0x72, 0x16, 0x5b, 0x4f, 0xc9, 0xcf, 0xf5, 0xe3, 0x7e, 0x21, 0x67, 0x3d,
	0x61, 0xa3, 0x5b, 0x92, 0x88, 0xb1, 0x61, 0x49, 0x96, 0xe4, 0x8b, 0xf2, 0x99, 0x07, 0xfc, 0x5e,
	0x1e, 0xf0, 0xd1, 0xd5, 0x2c, 0xa5, 0x5f, 0x24, 0x06, 0xc5, 0xd2, 0x0b, 0xab, 0x63, 0x50, 0x73,
	0xeb, 0x11, 0x90, 0x5d, 0x5d, 0xba, 0xf5, 0x08, 0x48, 0x29, 0x49, 0x3b, 0xa1, 0x3b, 0x76, 0x9d,
	0x45, 0xf9, 0x4d, 0x65, 0xaf, 0x9f, 0x3e, 0xb6, 0x6f, 0x6d, 0x8f, 0x9d, 0xd9, 0xf1, 0x1a, 0x86,
	0xa2, 0x7c, 0xd4, 0xd8, 0x43, 0xd1, 0xc2, 0xc3, 0x74, 0x10, 0xbf, 0xad, 0x02, 0x33, 0x36, 0x45,
	0xab, 0xa6, 0xba, 0xd0, 0x75, 0x27, 0x07, 0x11, 0x72, 0xa0, 0xff, 0x03, 0x00, 0x00, 0xff, 0xff,
	0x07, 0x68, 0x03, 0x1b, 0x1a, 0x03, 0x00, 0x00,
}
