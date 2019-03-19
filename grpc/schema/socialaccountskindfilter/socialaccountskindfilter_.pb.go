// Code generated by protoc-gen-go. DO NOT EDIT.
// source: socialaccountskindfilter/socialaccountskindfilter_.proto

/*
Package socialaccountskindfilter is a generated protocol buffer package.

It is generated from these files:
	socialaccountskindfilter/socialaccountskindfilter_.proto

It has these top-level messages:
	SocialAccountsKindFilter
*/
package socialaccountskindfilter

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import socialaccountskind "github.com/21stio/go-playground/grpc/schema/socialaccountskind"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type SocialAccountsKindFilter struct {
	Is               *socialaccountskind.SocialAccountsKind  `protobuf:"varint,1,opt,name=is,enum=socialaccountskind.SocialAccountsKind" json:"is,omitempty"`
	Not              *socialaccountskind.SocialAccountsKind  `protobuf:"varint,2,opt,name=not,enum=socialaccountskind.SocialAccountsKind" json:"not,omitempty"`
	In               []socialaccountskind.SocialAccountsKind `protobuf:"varint,3,rep,name=in,enum=socialaccountskind.SocialAccountsKind" json:"in,omitempty"`
	NotIn            []socialaccountskind.SocialAccountsKind `protobuf:"varint,4,rep,name=notIn,enum=socialaccountskind.SocialAccountsKind" json:"notIn,omitempty"`
	Hash             *string                                 `protobuf:"bytes,5,opt,name=hash" json:"hash,omitempty"`
	XXX_unrecognized []byte                                  `json:"-"`
}

func (m *SocialAccountsKindFilter) Reset()                    { *m = SocialAccountsKindFilter{} }
func (m *SocialAccountsKindFilter) String() string            { return proto.CompactTextString(m) }
func (*SocialAccountsKindFilter) ProtoMessage()               {}
func (*SocialAccountsKindFilter) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *SocialAccountsKindFilter) GetIs() socialaccountskind.SocialAccountsKind {
	if m != nil && m.Is != nil {
		return *m.Is
	}
	return socialaccountskind.SocialAccountsKind_SOCIAL_ACCOUNT_FOLLOWERS
}

func (m *SocialAccountsKindFilter) GetNot() socialaccountskind.SocialAccountsKind {
	if m != nil && m.Not != nil {
		return *m.Not
	}
	return socialaccountskind.SocialAccountsKind_SOCIAL_ACCOUNT_FOLLOWERS
}

func (m *SocialAccountsKindFilter) GetIn() []socialaccountskind.SocialAccountsKind {
	if m != nil {
		return m.In
	}
	return nil
}

func (m *SocialAccountsKindFilter) GetNotIn() []socialaccountskind.SocialAccountsKind {
	if m != nil {
		return m.NotIn
	}
	return nil
}

func (m *SocialAccountsKindFilter) GetHash() string {
	if m != nil && m.Hash != nil {
		return *m.Hash
	}
	return ""
}

func init() {
	proto.RegisterType((*SocialAccountsKindFilter)(nil), "socialaccountskindfilter.SocialAccountsKindFilter")
}

func init() {
	proto.RegisterFile("socialaccountskindfilter/socialaccountskindfilter_.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 221 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xb2, 0x28, 0xce, 0x4f, 0xce,
	0x4c, 0xcc, 0x49, 0x4c, 0x4e, 0xce, 0x2f, 0xcd, 0x2b, 0x29, 0xce, 0xce, 0xcc, 0x4b, 0x49, 0xcb,
	0xcc, 0x29, 0x49, 0x2d, 0xd2, 0xc7, 0x25, 0x11, 0xaf, 0x57, 0x50, 0x94, 0x5f, 0x92, 0x2f, 0x24,
	0x81, 0x4b, 0x81, 0x94, 0x0e, 0xa6, 0x0c, 0x16, 0xd3, 0xa0, 0xe6, 0x28, 0x4d, 0x63, 0xe2, 0x92,
	0x08, 0x06, 0xcb, 0x3a, 0x42, 0x65, 0xbd, 0x33, 0xf3, 0x52, 0xdc, 0xc0, 0x46, 0x09, 0x99, 0x71,
	0x31, 0x65, 0x16, 0x4b, 0x30, 0x2a, 0x30, 0x6a, 0xf0, 0x19, 0xa9, 0xe9, 0x61, 0x1a, 0xa2, 0x87,
	0xa9, 0x33, 0x88, 0x29, 0xb3, 0x58, 0xc8, 0x82, 0x8b, 0x39, 0x2f, 0xbf, 0x44, 0x82, 0x89, 0x24,
	0x8d, 0x20, 0x2d, 0x60, 0x1b, 0xf3, 0x24, 0x98, 0x15, 0x98, 0x49, 0xb2, 0x31, 0x4f, 0xc8, 0x86,
	0x8b, 0x35, 0x2f, 0xbf, 0xc4, 0x33, 0x4f, 0x82, 0x85, 0x24, 0xad, 0x10, 0x4d, 0x42, 0x42, 0x5c,
	0x2c, 0x19, 0x89, 0xc5, 0x19, 0x12, 0xac, 0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0x60, 0xb6, 0x93, 0x5b,
	0x94, 0x4b, 0x7a, 0x66, 0x49, 0x46, 0x69, 0x92, 0x5e, 0x72, 0x7e, 0xae, 0xbe, 0x91, 0x61, 0x71,
	0x49, 0x66, 0xbe, 0x7e, 0x7a, 0xbe, 0x6e, 0x41, 0x4e, 0x62, 0x65, 0x7a, 0x51, 0x7e, 0x69, 0x5e,
	0x8a, 0x7e, 0x7a, 0x51, 0x41, 0xb2, 0x7e, 0x71, 0x72, 0x46, 0x6a, 0x6e, 0x22, 0xce, 0xf8, 0x02,
	0x04, 0x00, 0x00, 0xff, 0xff, 0xfe, 0xed, 0x34, 0x56, 0xe3, 0x01, 0x00, 0x00,
}