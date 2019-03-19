// Code generated by protoc-gen-go. DO NOT EDIT.
// source: liketermkind/liketermkind_.proto

/*
Package liketermkind is a generated protocol buffer package.

It is generated from these files:
	liketermkind/liketermkind_.proto

It has these top-level messages:
*/
package liketermkind

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

type LikeTermKind int32

const (
	LikeTermKind_LIKE      LikeTermKind = 0
	LikeTermKind_FAVOURITE LikeTermKind = 1
)

var LikeTermKind_name = map[int32]string{
	0: "LIKE",
	1: "FAVOURITE",
}
var LikeTermKind_value = map[string]int32{
	"LIKE":      0,
	"FAVOURITE": 1,
}

func (x LikeTermKind) Enum() *LikeTermKind {
	p := new(LikeTermKind)
	*p = x
	return p
}
func (x LikeTermKind) String() string {
	return proto.EnumName(LikeTermKind_name, int32(x))
}
func (x *LikeTermKind) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(LikeTermKind_value, data, "LikeTermKind")
	if err != nil {
		return err
	}
	*x = LikeTermKind(value)
	return nil
}
func (LikeTermKind) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func init() {
	proto.RegisterEnum("liketermkind.LikeTermKind", LikeTermKind_name, LikeTermKind_value)
}

func init() { proto.RegisterFile("liketermkind/liketermkind_.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 142 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0xc8, 0xc9, 0xcc, 0x4e,
	0x2d, 0x49, 0x2d, 0xca, 0xcd, 0xce, 0xcc, 0x4b, 0xd1, 0x47, 0xe6, 0xc4, 0xeb, 0x15, 0x14, 0xe5,
	0x97, 0xe4, 0x0b, 0xf1, 0x20, 0x0b, 0x6a, 0xa9, 0x73, 0xf1, 0xf8, 0x64, 0x66, 0xa7, 0x86, 0xa4,
	0x16, 0xe5, 0x7a, 0x67, 0xe6, 0xa5, 0x08, 0x71, 0x70, 0xb1, 0xf8, 0x78, 0x7a, 0xbb, 0x0a, 0x30,
	0x08, 0xf1, 0x72, 0x71, 0xba, 0x39, 0x86, 0xf9, 0x87, 0x06, 0x79, 0x86, 0xb8, 0x0a, 0x30, 0x3a,
	0x59, 0x45, 0x59, 0xa4, 0x67, 0x96, 0x64, 0x94, 0x26, 0xe9, 0x25, 0xe7, 0xe7, 0xea, 0x1b, 0x19,
	0x16, 0x97, 0x64, 0xe6, 0xeb, 0xa7, 0xe7, 0xeb, 0x16, 0xe4, 0x24, 0x56, 0xa6, 0x17, 0xe5, 0x97,
	0xe6, 0xa5, 0xe8, 0xa7, 0x17, 0x15, 0x24, 0xeb, 0x17, 0x27, 0x67, 0xa4, 0xe6, 0x26, 0xa2, 0xd8,
	0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0xbb, 0x20, 0x74, 0x6a, 0x95, 0x00, 0x00, 0x00,
}
