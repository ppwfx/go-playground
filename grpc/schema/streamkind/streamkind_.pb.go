// Code generated by protoc-gen-go. DO NOT EDIT.
// source: streamkind/streamkind_.proto

/*
Package streamkind is a generated protocol buffer package.

It is generated from these files:
	streamkind/streamkind_.proto

It has these top-level messages:
*/
package streamkind

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

type StreamKind int32

const (
	StreamKind_WATCH StreamKind = 0
)

var StreamKind_name = map[int32]string{
	0: "WATCH",
}
var StreamKind_value = map[string]int32{
	"WATCH": 0,
}

func (x StreamKind) Enum() *StreamKind {
	p := new(StreamKind)
	*p = x
	return p
}
func (x StreamKind) String() string {
	return proto.EnumName(StreamKind_name, int32(x))
}
func (x *StreamKind) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(StreamKind_value, data, "StreamKind")
	if err != nil {
		return err
	}
	*x = StreamKind(value)
	return nil
}
func (StreamKind) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func init() {
	proto.RegisterEnum("streamkind.StreamKind", StreamKind_name, StreamKind_value)
}

func init() { proto.RegisterFile("streamkind/streamkind_.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 121 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x29, 0x2e, 0x29, 0x4a,
	0x4d, 0xcc, 0xcd, 0xce, 0xcc, 0x4b, 0xd1, 0x47, 0x30, 0xe3, 0xf5, 0x0a, 0x8a, 0xf2, 0x4b, 0xf2,
	0x85, 0xb8, 0x10, 0x42, 0x5a, 0xe2, 0x5c, 0x5c, 0xc1, 0x60, 0x9e, 0x77, 0x66, 0x5e, 0x8a, 0x10,
	0x27, 0x17, 0x6b, 0xb8, 0x63, 0x88, 0xb3, 0x87, 0x00, 0x83, 0x93, 0x45, 0x94, 0x59, 0x7a, 0x66,
	0x49, 0x46, 0x69, 0x92, 0x5e, 0x72, 0x7e, 0xae, 0xbe, 0x91, 0x61, 0x71, 0x49, 0x66, 0xbe, 0x7e,
	0x7a, 0xbe, 0x6e, 0x41, 0x4e, 0x62, 0x65, 0x7a, 0x51, 0x7e, 0x69, 0x5e, 0x8a, 0x7e, 0x7a, 0x51,
	0x41, 0xb2, 0x7e, 0x71, 0x72, 0x46, 0x6a, 0x6e, 0x22, 0x92, 0x2d, 0x80, 0x00, 0x00, 0x00, 0xff,
	0xff, 0x9b, 0x8e, 0xcc, 0x0e, 0x7d, 0x00, 0x00, 0x00,
}
