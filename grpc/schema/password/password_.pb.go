// Code generated by protoc-gen-go. DO NOT EDIT.
// source: password/password_.proto

/*
Package password is a generated protocol buffer package.

It is generated from these files:
	password/password_.proto

It has these top-level messages:
	Password
*/
package password

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import hashfunctions "github.com/21stio/go-playground/grpc/schema/hashfunctions"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Password struct {
	IsHashed         *bool                        `protobuf:"varint,1,opt,name=isHashed" json:"isHashed,omitempty"`
	HashFunction     *hashfunctions.HashFunctions `protobuf:"varint,2,opt,name=hashFunction,enum=hashfunctions.HashFunctions" json:"hashFunction,omitempty"`
	Value            *string                      `protobuf:"bytes,3,opt,name=value" json:"value,omitempty"`
	Hash             *string                      `protobuf:"bytes,4,opt,name=hash" json:"hash,omitempty"`
	XXX_unrecognized []byte                       `json:"-"`
}

func (m *Password) Reset()                    { *m = Password{} }
func (m *Password) String() string            { return proto.CompactTextString(m) }
func (*Password) ProtoMessage()               {}
func (*Password) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Password) GetIsHashed() bool {
	if m != nil && m.IsHashed != nil {
		return *m.IsHashed
	}
	return false
}

func (m *Password) GetHashFunction() hashfunctions.HashFunctions {
	if m != nil && m.HashFunction != nil {
		return *m.HashFunction
	}
	return hashfunctions.HashFunctions_BLAKE_256
}

func (m *Password) GetValue() string {
	if m != nil && m.Value != nil {
		return *m.Value
	}
	return ""
}

func (m *Password) GetHash() string {
	if m != nil && m.Hash != nil {
		return *m.Hash
	}
	return ""
}

func init() {
	proto.RegisterType((*Password)(nil), "password.Password")
}

func init() { proto.RegisterFile("password/password_.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 204 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x28, 0x48, 0x2c, 0x2e,
	0x2e, 0xcf, 0x2f, 0x4a, 0xd1, 0x87, 0x31, 0xe2, 0xf5, 0x0a, 0x8a, 0xf2, 0x4b, 0xf2, 0x85, 0x38,
	0x60, 0x02, 0x52, 0x4a, 0x19, 0x89, 0xc5, 0x19, 0x69, 0xa5, 0x79, 0xc9, 0x25, 0x99, 0xf9, 0x79,
	0xc5, 0xfa, 0x28, 0x3c, 0xa8, 0x6a, 0xa5, 0x49, 0x8c, 0x5c, 0x1c, 0x01, 0x50, 0x0d, 0x42, 0x52,
	0x5c, 0x1c, 0x99, 0xc5, 0x1e, 0x89, 0xc5, 0x19, 0xa9, 0x29, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x1c,
	0x41, 0x70, 0xbe, 0x90, 0x03, 0x17, 0x0f, 0xc8, 0x00, 0x37, 0xa8, 0x01, 0x12, 0x4c, 0x0a, 0x8c,
	0x1a, 0x7c, 0x46, 0x32, 0x7a, 0x28, 0xa6, 0xea, 0x79, 0x20, 0x29, 0x29, 0x0e, 0x42, 0xd1, 0x21,
	0x24, 0xc2, 0xc5, 0x5a, 0x96, 0x98, 0x53, 0x9a, 0x2a, 0xc1, 0xac, 0xc0, 0xa8, 0xc1, 0x19, 0x04,
	0xe1, 0x08, 0x09, 0x71, 0xb1, 0x80, 0x54, 0x49, 0xb0, 0x80, 0x05, 0xc1, 0x6c, 0x27, 0xb3, 0x28,
	0x93, 0xf4, 0xcc, 0x92, 0x8c, 0xd2, 0x24, 0xbd, 0xe4, 0xfc, 0x5c, 0x7d, 0x23, 0xc3, 0xe2, 0x92,
	0xcc, 0x7c, 0xfd, 0xf4, 0x7c, 0xdd, 0x82, 0x9c, 0xc4, 0xca, 0xf4, 0xa2, 0xfc, 0xd2, 0xbc, 0x14,
	0xfd, 0xf4, 0xa2, 0x82, 0x64, 0xfd, 0xe2, 0xe4, 0x8c, 0xd4, 0xdc, 0x44, 0x78, 0x08, 0x00, 0x02,
	0x00, 0x00, 0xff, 0xff, 0xd9, 0x7a, 0xf5, 0xa1, 0x15, 0x01, 0x00, 0x00,
}