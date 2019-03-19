// Code generated by protoc-gen-go. DO NOT EDIT.
// source: formattingkind/formattingkind_.proto

/*
Package formattingkind is a generated protocol buffer package.

It is generated from these files:
	formattingkind/formattingkind_.proto

It has these top-level messages:
*/
package formattingkind

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

type FormattingKind int32

const (
	FormattingKind_HTML     FormattingKind = 0
	FormattingKind_PLAIN    FormattingKind = 1
	FormattingKind_MARKDOWN FormattingKind = 2
)

var FormattingKind_name = map[int32]string{
	0: "HTML",
	1: "PLAIN",
	2: "MARKDOWN",
}
var FormattingKind_value = map[string]int32{
	"HTML":     0,
	"PLAIN":    1,
	"MARKDOWN": 2,
}

func (x FormattingKind) Enum() *FormattingKind {
	p := new(FormattingKind)
	*p = x
	return p
}
func (x FormattingKind) String() string {
	return proto.EnumName(FormattingKind_name, int32(x))
}
func (x *FormattingKind) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(FormattingKind_value, data, "FormattingKind")
	if err != nil {
		return err
	}
	*x = FormattingKind(value)
	return nil
}
func (FormattingKind) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func init() {
	proto.RegisterEnum("formattingkind.FormattingKind", FormattingKind_name, FormattingKind_value)
}

func init() { proto.RegisterFile("formattingkind/formattingkind_.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 149 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x49, 0xcb, 0x2f, 0xca,
	0x4d, 0x2c, 0x29, 0xc9, 0xcc, 0x4b, 0xcf, 0xce, 0xcc, 0x4b, 0xd1, 0x47, 0xe5, 0xc6, 0xeb, 0x15,
	0x14, 0xe5, 0x97, 0xe4, 0x0b, 0xf1, 0xa1, 0x0a, 0x6b, 0x19, 0x73, 0xf1, 0xb9, 0xc1, 0x45, 0xbc,
	0x33, 0xf3, 0x52, 0x84, 0x38, 0xb8, 0x58, 0x3c, 0x42, 0x7c, 0x7d, 0x04, 0x18, 0x84, 0x38, 0xb9,
	0x58, 0x03, 0x7c, 0x1c, 0x3d, 0xfd, 0x04, 0x18, 0x85, 0x78, 0xb8, 0x38, 0x7c, 0x1d, 0x83, 0xbc,
	0x5d, 0xfc, 0xc3, 0xfd, 0x04, 0x98, 0x9c, 0x6c, 0xa2, 0xac, 0xd2, 0x33, 0x4b, 0x32, 0x4a, 0x93,
	0xf4, 0x92, 0xf3, 0x73, 0xf5, 0x8d, 0x0c, 0x8b, 0x4b, 0x32, 0xf3, 0xf5, 0xd3, 0xf3, 0x75, 0x0b,
	0x72, 0x12, 0x2b, 0xd3, 0x8b, 0xf2, 0x4b, 0xf3, 0x52, 0xf4, 0xd3, 0x8b, 0x0a, 0x92, 0xf5, 0x8b,
	0x93, 0x33, 0x52, 0x73, 0x13, 0xd1, 0x5c, 0x02, 0x08, 0x00, 0x00, 0xff, 0xff, 0x30, 0x55, 0xea,
	0x54, 0xa9, 0x00, 0x00, 0x00,
}