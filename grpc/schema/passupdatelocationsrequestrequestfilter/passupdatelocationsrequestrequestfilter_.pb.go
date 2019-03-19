// Code generated by protoc-gen-go. DO NOT EDIT.
// source: passupdatelocationsrequestrequestfilter/passupdatelocationsrequestrequestfilter_.proto

/*
Package passupdatelocationsrequestrequestfilter is a generated protocol buffer package.

It is generated from these files:
	passupdatelocationsrequestrequestfilter/passupdatelocationsrequestrequestfilter_.proto

It has these top-level messages:
	PassUpdateLocationsRequestRequestFilter
*/
package passupdatelocationsrequestrequestfilter

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import requestmetafilter "github.com/21stio/go-playground/grpc/schema/requestmetafilter"
import updatelocationsrequestfilter "github.com/21stio/go-playground/grpc/schema/updatelocationsrequestfilter"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type PassUpdateLocationsRequestRequestFilter struct {
	Meta                   *requestmetafilter.RequestMetaFilter                       `protobuf:"bytes,1,opt,name=meta" json:"meta,omitempty"`
	UpdateLocationsRequest *updatelocationsrequestfilter.UpdateLocationsRequestFilter `protobuf:"bytes,2,opt,name=updateLocationsRequest" json:"updateLocationsRequest,omitempty"`
	And                    []*PassUpdateLocationsRequestRequestFilter                 `protobuf:"bytes,3,rep,name=and" json:"and,omitempty"`
	Or                     []*PassUpdateLocationsRequestRequestFilter                 `protobuf:"bytes,4,rep,name=or" json:"or,omitempty"`
	Not                    []*PassUpdateLocationsRequestRequestFilter                 `protobuf:"bytes,5,rep,name=not" json:"not,omitempty"`
	Hash                   *string                                                    `protobuf:"bytes,6,opt,name=hash" json:"hash,omitempty"`
	XXX_unrecognized       []byte                                                     `json:"-"`
}

func (m *PassUpdateLocationsRequestRequestFilter) Reset() {
	*m = PassUpdateLocationsRequestRequestFilter{}
}
func (m *PassUpdateLocationsRequestRequestFilter) String() string { return proto.CompactTextString(m) }
func (*PassUpdateLocationsRequestRequestFilter) ProtoMessage()    {}
func (*PassUpdateLocationsRequestRequestFilter) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{0}
}

func (m *PassUpdateLocationsRequestRequestFilter) GetMeta() *requestmetafilter.RequestMetaFilter {
	if m != nil {
		return m.Meta
	}
	return nil
}

func (m *PassUpdateLocationsRequestRequestFilter) GetUpdateLocationsRequest() *updatelocationsrequestfilter.UpdateLocationsRequestFilter {
	if m != nil {
		return m.UpdateLocationsRequest
	}
	return nil
}

func (m *PassUpdateLocationsRequestRequestFilter) GetAnd() []*PassUpdateLocationsRequestRequestFilter {
	if m != nil {
		return m.And
	}
	return nil
}

func (m *PassUpdateLocationsRequestRequestFilter) GetOr() []*PassUpdateLocationsRequestRequestFilter {
	if m != nil {
		return m.Or
	}
	return nil
}

func (m *PassUpdateLocationsRequestRequestFilter) GetNot() []*PassUpdateLocationsRequestRequestFilter {
	if m != nil {
		return m.Not
	}
	return nil
}

func (m *PassUpdateLocationsRequestRequestFilter) GetHash() string {
	if m != nil && m.Hash != nil {
		return *m.Hash
	}
	return ""
}

func init() {
	proto.RegisterType((*PassUpdateLocationsRequestRequestFilter)(nil), "passupdatelocationsrequestrequestfilter.PassUpdateLocationsRequestRequestFilter")
}

func init() {
	proto.RegisterFile("passupdatelocationsrequestrequestfilter/passupdatelocationsrequestrequestfilter_.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 285 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x53, 0x3b, 0x4b, 0xc4, 0x40,
	0x10, 0x26, 0x0f, 0x05, 0xf7, 0xba, 0x2d, 0x64, 0xb9, 0x2a, 0x88, 0x70, 0x41, 0x70, 0x83, 0xa9,
	0xc4, 0x4a, 0x2c, 0xac, 0x14, 0x8e, 0xc8, 0x59, 0xd8, 0xe8, 0x5c, 0xb2, 0x26, 0x81, 0x24, 0x13,
	0x77, 0x27, 0x85, 0x3f, 0xc4, 0xff, 0x2b, 0xc9, 0xad, 0xd5, 0xc5, 0x63, 0x9b, 0xab, 0x66, 0x1f,
	0xdf, 0x6b, 0x06, 0x86, 0xbd, 0xf6, 0x60, 0xcc, 0xd0, 0x17, 0x40, 0xaa, 0xc1, 0x1c, 0xa8, 0xc6,
	0xce, 0x68, 0xf5, 0x35, 0x28, 0x43, 0xb6, 0x7c, 0xd6, 0x0d, 0x29, 0x9d, 0x38, 0xe2, 0xde, 0x65,
	0xaf, 0x91, 0x90, 0xaf, 0x1c, 0xf1, 0xcb, 0x2b, 0x7b, 0x6d, 0x15, 0x81, 0xb5, 0xda, 0x7b, 0xb1,
	0xa2, 0xcb, 0xfb, 0x79, 0x41, 0x4b, 0x3b, 0xf4, 0x69, 0x15, 0x2e, 0x7e, 0x42, 0xb6, 0x5a, 0x83,
	0x31, 0x9b, 0x09, 0xfb, 0xf4, 0x87, 0xcd, 0x76, 0x58, 0x5b, 0x1e, 0x27, 0x0a, 0xbf, 0x65, 0xe1,
	0x18, 0x41, 0x78, 0x91, 0x17, 0x2f, 0xd2, 0x4b, 0xb9, 0x17, 0x4b, 0x5a, 0xfc, 0xb3, 0x22, 0xd8,
	0x71, 0xb2, 0x89, 0xc1, 0x35, 0x3b, 0x1f, 0x66, 0x0d, 0x84, 0x3f, 0x69, 0xdd, 0xc9, 0x43, 0x59,
	0xe5, 0x7c, 0x38, 0xeb, 0xf0, 0x8f, 0x32, 0xdf, 0xb2, 0x00, 0xba, 0x42, 0x04, 0x51, 0x10, 0x2f,
	0xd2, 0xb5, 0x74, 0x1c, 0xbf, 0x74, 0x1c, 0x46, 0x36, 0x8a, 0xf3, 0x0f, 0xe6, 0xa3, 0x16, 0xe1,
	0x91, 0x2c, 0x7c, 0xd4, 0x63, 0x17, 0x1d, 0x92, 0x38, 0x39, 0x56, 0x17, 0x1d, 0x12, 0xe7, 0x2c,
	0xac, 0xc0, 0x54, 0xe2, 0x34, 0xf2, 0xe2, 0xb3, 0x6c, 0x3a, 0x3f, 0x6c, 0xde, 0x5e, 0xca, 0x9a,
	0xaa, 0x61, 0x2b, 0x73, 0x6c, 0x93, 0xf4, 0xc6, 0x50, 0x8d, 0x49, 0x89, 0xd7, 0x7d, 0x03, 0xdf,
	0xa5, 0xc6, 0xa1, 0x2b, 0x92, 0x52, 0xf7, 0x79, 0x62, 0xf2, 0x4a, 0xb5, 0xe0, 0xba, 0x0c, 0xbf,
	0x01, 0x00, 0x00, 0xff, 0xff, 0x8e, 0x25, 0x7f, 0xc7, 0x5e, 0x03, 0x00, 0x00,
}