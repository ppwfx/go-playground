// Code generated by protoc-gen-go. DO NOT EDIT.
// source: sortkind/sortkind_.proto

/*
Package sortkind is a generated protocol buffer package.

It is generated from these files:
	sortkind/sortkind_.proto

It has these top-level messages:
*/
package sortkind

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

type SortKind int32

const (
	SortKind_ASC  SortKind = 0
	SortKind_DESC SortKind = 1
)

var SortKind_name = map[int32]string{
	0: "ASC",
	1: "DESC",
}
var SortKind_value = map[string]int32{
	"ASC":  0,
	"DESC": 1,
}

func (x SortKind) Enum() *SortKind {
	p := new(SortKind)
	*p = x
	return p
}
func (x SortKind) String() string {
	return proto.EnumName(SortKind_name, int32(x))
}
func (x *SortKind) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(SortKind_value, data, "SortKind")
	if err != nil {
		return err
	}
	*x = SortKind(value)
	return nil
}
func (SortKind) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func init() {
	proto.RegisterEnum("sortkind.SortKind", SortKind_name, SortKind_value)
}

func init() { proto.RegisterFile("sortkind/sortkind_.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 128 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x28, 0xce, 0x2f, 0x2a,
	0xc9, 0xce, 0xcc, 0x4b, 0xd1, 0x87, 0x31, 0xe2, 0xf5, 0x0a, 0x8a, 0xf2, 0x4b, 0xf2, 0x85, 0x38,
	0x60, 0x02, 0x5a, 0xb2, 0x5c, 0x1c, 0xc1, 0xf9, 0x45, 0x25, 0xde, 0x99, 0x79, 0x29, 0x42, 0xec,
	0x5c, 0xcc, 0x8e, 0xc1, 0xce, 0x02, 0x0c, 0x42, 0x1c, 0x5c, 0x2c, 0x2e, 0xae, 0xc1, 0xce, 0x02,
	0x8c, 0x4e, 0x66, 0x51, 0x26, 0xe9, 0x99, 0x25, 0x19, 0xa5, 0x49, 0x7a, 0xc9, 0xf9, 0xb9, 0xfa,
	0x46, 0x86, 0xc5, 0x25, 0x99, 0xf9, 0xfa, 0xe9, 0xf9, 0xba, 0x05, 0x39, 0x89, 0x95, 0xe9, 0x45,
	0xf9, 0xa5, 0x79, 0x29, 0xfa, 0xe9, 0x45, 0x05, 0xc9, 0xfa, 0xc5, 0xc9, 0x19, 0xa9, 0xb9, 0x89,
	0x70, 0x7b, 0x00, 0x01, 0x00, 0x00, 0xff, 0xff, 0xed, 0x28, 0xfa, 0x4a, 0x7b, 0x00, 0x00, 0x00,
}
