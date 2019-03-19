// Code generated by protoc-gen-go. DO NOT EDIT.
// source: timestampkind/timestampkind_.proto

/*
Package timestampkind is a generated protocol buffer package.

It is generated from these files:
	timestampkind/timestampkind_.proto

It has these top-level messages:
*/
package timestampkind

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

type TimestampKind int32

const (
	TimestampKind_UNIX      TimestampKind = 0
	TimestampKind_UNIX_NANO TimestampKind = 1
)

var TimestampKind_name = map[int32]string{
	0: "UNIX",
	1: "UNIX_NANO",
}
var TimestampKind_value = map[string]int32{
	"UNIX":      0,
	"UNIX_NANO": 1,
}

func (x TimestampKind) Enum() *TimestampKind {
	p := new(TimestampKind)
	*p = x
	return p
}
func (x TimestampKind) String() string {
	return proto.EnumName(TimestampKind_name, int32(x))
}
func (x *TimestampKind) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(TimestampKind_value, data, "TimestampKind")
	if err != nil {
		return err
	}
	*x = TimestampKind(value)
	return nil
}
func (TimestampKind) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func init() {
	proto.RegisterEnum("timestampkind.TimestampKind", TimestampKind_name, TimestampKind_value)
}

func init() { proto.RegisterFile("timestampkind/timestampkind_.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 136 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x2a, 0xc9, 0xcc, 0x4d,
	0x2d, 0x2e, 0x49, 0xcc, 0x2d, 0xc8, 0xce, 0xcc, 0x4b, 0xd1, 0x47, 0xe1, 0xc5, 0xeb, 0x15, 0x14,
	0xe5, 0x97, 0xe4, 0x0b, 0xf1, 0xa2, 0x88, 0x6a, 0x69, 0x70, 0xf1, 0x86, 0xc0, 0x04, 0xbc, 0x33,
	0xf3, 0x52, 0x84, 0x38, 0xb8, 0x58, 0x42, 0xfd, 0x3c, 0x23, 0x04, 0x18, 0x84, 0x78, 0xb9, 0x38,
	0x41, 0xac, 0x78, 0x3f, 0x47, 0x3f, 0x7f, 0x01, 0x46, 0x27, 0xeb, 0x28, 0xcb, 0xf4, 0xcc, 0x92,
	0x8c, 0xd2, 0x24, 0xbd, 0xe4, 0xfc, 0x5c, 0x7d, 0x23, 0xc3, 0xe2, 0x92, 0xcc, 0x7c, 0xfd, 0xf4,
	0x7c, 0xdd, 0x82, 0x9c, 0xc4, 0xca, 0xf4, 0xa2, 0xfc, 0xd2, 0xbc, 0x14, 0xfd, 0xf4, 0xa2, 0x82,
	0x64, 0xfd, 0xe2, 0xe4, 0x8c, 0xd4, 0xdc, 0x44, 0x54, 0xcb, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff,
	0x45, 0x57, 0x21, 0x5e, 0x9a, 0x00, 0x00, 0x00,
}
