// Code generated by protoc-gen-go. DO NOT EDIT.
// source: serviceskind/serviceskind_.proto

/*
Package serviceskind is a generated protocol buffer package.

It is generated from these files:
	serviceskind/serviceskind_.proto

It has these top-level messages:
*/
package serviceskind

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type ServicesKind int32

const (
	ServicesKind_RESPONSE_META_SERVICES ServicesKind = 0
)

var ServicesKind_name = map[int32]string{
	0: "RESPONSE_META_SERVICES",
}
var ServicesKind_value = map[string]int32{
	"RESPONSE_META_SERVICES": 0,
}

func (x ServicesKind) Enum() *ServicesKind {
	p := new(ServicesKind)
	*p = x
	return p
}
func (x ServicesKind) String() string {
	return proto.EnumName(ServicesKind_name, int32(x))
}
func (x *ServicesKind) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(ServicesKind_value, data, "ServicesKind")
	if err != nil {
		return err
	}
	*x = ServicesKind(value)
	return nil
}
func (ServicesKind) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func init() {
	proto.RegisterEnum("serviceskind.ServicesKind", ServicesKind_name, ServicesKind_value)
}

func init() { proto.RegisterFile("serviceskind/serviceskind_.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 140 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x28, 0x4e, 0x2d, 0x2a,
	0xcb, 0x4c, 0x4e, 0x2d, 0xce, 0xce, 0xcc, 0x4b, 0xd1, 0x47, 0xe6, 0xc4, 0xeb, 0x15, 0x14, 0xe5,
	0x97, 0xe4, 0x0b, 0xf1, 0x20, 0x0b, 0x6a, 0x69, 0x71, 0xf1, 0x04, 0x43, 0xf9, 0xde, 0x99, 0x79,
	0x29, 0x42, 0x52, 0x5c, 0x62, 0x41, 0xae, 0xc1, 0x01, 0xfe, 0x7e, 0xc1, 0xae, 0xf1, 0xbe, 0xae,
	0x21, 0x8e, 0xf1, 0xc1, 0xae, 0x41, 0x61, 0x9e, 0xce, 0xae, 0xc1, 0x02, 0x0c, 0x4e, 0x56, 0x51,
	0x16, 0xe9, 0x99, 0x25, 0x19, 0xa5, 0x49, 0x7a, 0xc9, 0xf9, 0xb9, 0xfa, 0x46, 0x86, 0xc5, 0x25,
	0x99, 0xf9, 0xfa, 0xe9, 0xf9, 0xba, 0x05, 0x39, 0x89, 0x95, 0xe9, 0x45, 0xf9, 0xa5, 0x79, 0x29,
	0xfa, 0xe9, 0x45, 0x05, 0xc9, 0xfa, 0xc5, 0xc9, 0x19, 0xa9, 0xb9, 0x89, 0x28, 0x96, 0x03, 0x02,
	0x00, 0x00, 0xff, 0xff, 0xa3, 0x49, 0x95, 0x98, 0x98, 0x00, 0x00, 0x00,
}
