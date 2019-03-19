// Code generated by protoc-gen-go. DO NOT EDIT.
// source: tokenfilter/tokenfilter_.proto

/*
Package tokenfilter is a generated protocol buffer package.

It is generated from these files:
	tokenfilter/tokenfilter_.proto

It has these top-level messages:
	TokenFilter
*/
package tokenfilter

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import tokenkindfilter "github.com/21stio/go-playground/grpc/schema/tokenkindfilter"
import jwttokenfilter "github.com/21stio/go-playground/grpc/schema/jwttokenfilter"
import bearertokenfilter "github.com/21stio/go-playground/grpc/schema/bearertokenfilter"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type TokenFilter struct {
	Kind             *tokenkindfilter.TokenKindFilter     `protobuf:"bytes,1,opt,name=kind" json:"kind,omitempty"`
	JwtToken         *jwttokenfilter.JwtTokenFilter       `protobuf:"bytes,2,opt,name=jwtToken" json:"jwtToken,omitempty"`
	BearerToken      *bearertokenfilter.BearerTokenFilter `protobuf:"bytes,3,opt,name=bearerToken" json:"bearerToken,omitempty"`
	And              []*TokenFilter                       `protobuf:"bytes,4,rep,name=and" json:"and,omitempty"`
	Or               []*TokenFilter                       `protobuf:"bytes,5,rep,name=or" json:"or,omitempty"`
	Not              []*TokenFilter                       `protobuf:"bytes,6,rep,name=not" json:"not,omitempty"`
	Hash             *string                              `protobuf:"bytes,7,opt,name=hash" json:"hash,omitempty"`
	XXX_unrecognized []byte                               `json:"-"`
}

func (m *TokenFilter) Reset()                    { *m = TokenFilter{} }
func (m *TokenFilter) String() string            { return proto.CompactTextString(m) }
func (*TokenFilter) ProtoMessage()               {}
func (*TokenFilter) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *TokenFilter) GetKind() *tokenkindfilter.TokenKindFilter {
	if m != nil {
		return m.Kind
	}
	return nil
}

func (m *TokenFilter) GetJwtToken() *jwttokenfilter.JwtTokenFilter {
	if m != nil {
		return m.JwtToken
	}
	return nil
}

func (m *TokenFilter) GetBearerToken() *bearertokenfilter.BearerTokenFilter {
	if m != nil {
		return m.BearerToken
	}
	return nil
}

func (m *TokenFilter) GetAnd() []*TokenFilter {
	if m != nil {
		return m.And
	}
	return nil
}

func (m *TokenFilter) GetOr() []*TokenFilter {
	if m != nil {
		return m.Or
	}
	return nil
}

func (m *TokenFilter) GetNot() []*TokenFilter {
	if m != nil {
		return m.Not
	}
	return nil
}

func (m *TokenFilter) GetHash() string {
	if m != nil && m.Hash != nil {
		return *m.Hash
	}
	return ""
}

func init() {
	proto.RegisterType((*TokenFilter)(nil), "tokenfilter.TokenFilter")
}

func init() { proto.RegisterFile("tokenfilter/tokenfilter_.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 285 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0xbd, 0x4e, 0xf3, 0x30,
	0x14, 0x86, 0x95, 0x9f, 0xef, 0x03, 0xec, 0xcd, 0x93, 0xd5, 0xa1, 0x8a, 0x50, 0x85, 0xa2, 0x4a,
	0xd8, 0x22, 0x42, 0x42, 0x30, 0x76, 0xe8, 0x00, 0x5b, 0xc5, 0xc4, 0x82, 0xdc, 0xc4, 0x24, 0xee,
	0x8f, 0x4f, 0xe4, 0xb8, 0xaa, 0xb8, 0x58, 0xee, 0x05, 0xc5, 0x31, 0xd4, 0xa5, 0x43, 0x37, 0x9f,
	0x73, 0x9e, 0xe7, 0xcd, 0x3b, 0x04, 0x8d, 0x2d, 0xac, 0xa5, 0xfe, 0x50, 0x1b, 0x2b, 0x0d, 0x0f,
	0xde, 0xef, 0xac, 0x35, 0x60, 0x81, 0xe0, 0x60, 0x37, 0xba, 0x71, 0xc3, 0x5a, 0xe9, 0x2a, 0x14,
	0x0e, 0xb3, 0x97, 0x46, 0x93, 0xd5, 0xde, 0x86, 0xb9, 0xc7, 0xe3, 0x0f, 0x35, 0x5d, 0x4a, 0x61,
	0xa4, 0x09, 0xc1, 0x93, 0x8d, 0x67, 0xaf, 0xbf, 0x62, 0x84, 0x5f, 0xfb, 0xf5, 0xdc, 0xad, 0xc9,
	0x3d, 0x4a, 0xfb, 0xcf, 0xd2, 0x28, 0x8b, 0x72, 0x5c, 0x64, 0xec, 0x4f, 0x11, 0xe6, 0xd8, 0x17,
	0xa5, 0xab, 0x81, 0x5f, 0x38, 0x9a, 0x3c, 0xa1, 0xcb, 0xd5, 0xde, 0xba, 0x1b, 0x8d, 0x9d, 0x39,
	0x66, 0xc7, 0xdd, 0xd8, 0xb3, 0xbf, 0x7b, 0xef, 0x97, 0x27, 0x73, 0x84, 0x87, 0x76, 0x83, 0x9e,
	0x38, 0x7d, 0xc2, 0x4e, 0x1a, 0xb3, 0xd9, 0x81, 0xf2, 0x21, 0xa1, 0x48, 0xa6, 0x28, 0x11, 0xba,
	0xa2, 0x69, 0x96, 0xe4, 0xb8, 0xa0, 0x2c, 0x34, 0x43, 0xa7, 0x87, 0x48, 0x8e, 0x62, 0x30, 0xf4,
	0xdf, 0x19, 0x34, 0x06, 0xd3, 0xa7, 0x6a, 0xb0, 0xf4, 0xff, 0xb9, 0x54, 0x0d, 0x96, 0x10, 0x94,
	0x36, 0xa2, 0x6b, 0xe8, 0x45, 0x16, 0xe5, 0x57, 0x0b, 0xf7, 0x9e, 0x3d, 0xbe, 0x3d, 0xd4, 0xca,
	0x36, 0xbb, 0x25, 0x2b, 0x61, 0xcb, 0x8b, 0xbb, 0xce, 0x2a, 0xe0, 0x35, 0xdc, 0xb6, 0x1b, 0xf1,
	0x59, 0x1b, 0xd8, 0xe9, 0x8a, 0xd7, 0xa6, 0x2d, 0x79, 0x57, 0x36, 0x72, 0x2b, 0xc2, 0x1f, 0xe5,
	0x3b, 0x00, 0x00, 0xff, 0xff, 0xb9, 0x2f, 0x6d, 0xf5, 0x42, 0x02, 0x00, 0x00,
}
