// Code generated by protoc-gen-go. DO NOT EDIT.
// source: videosort/videosort_.proto

/*
Package videosort is a generated protocol buffer package.

It is generated from these files:
	videosort/videosort_.proto

It has these top-level messages:
	VideoSort
*/
package videosort

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import urlsort "github.com/21stio/go-playground/grpc/schema/urlsort"
import sortkind "github.com/21stio/go-playground/grpc/schema/sortkind"
import durationvaluesort "github.com/21stio/go-playground/grpc/schema/durationvaluesort"
import textsort "github.com/21stio/go-playground/grpc/schema/textsort"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type VideoSort struct {
	Url              *urlsort.UrlSort                     `protobuf:"bytes,1,opt,name=url" json:"url,omitempty"`
	Width            *sortkind.SortKind                   `protobuf:"varint,2,opt,name=width,enum=sortkind.SortKind" json:"width,omitempty"`
	Height           *sortkind.SortKind                   `protobuf:"varint,3,opt,name=height,enum=sortkind.SortKind" json:"height,omitempty"`
	FrameRate        *sortkind.SortKind                   `protobuf:"varint,4,opt,name=frameRate,enum=sortkind.SortKind" json:"frameRate,omitempty"`
	Duration         *durationvaluesort.DurationValueSort `protobuf:"bytes,5,opt,name=duration" json:"duration,omitempty"`
	Description      *textsort.TextSort                   `protobuf:"bytes,6,opt,name=description" json:"description,omitempty"`
	Hash             *string                              `protobuf:"bytes,7,opt,name=hash" json:"hash,omitempty"`
	XXX_unrecognized []byte                               `json:"-"`
}

func (m *VideoSort) Reset()                    { *m = VideoSort{} }
func (m *VideoSort) String() string            { return proto.CompactTextString(m) }
func (*VideoSort) ProtoMessage()               {}
func (*VideoSort) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *VideoSort) GetUrl() *urlsort.UrlSort {
	if m != nil {
		return m.Url
	}
	return nil
}

func (m *VideoSort) GetWidth() sortkind.SortKind {
	if m != nil && m.Width != nil {
		return *m.Width
	}
	return sortkind.SortKind_ASC
}

func (m *VideoSort) GetHeight() sortkind.SortKind {
	if m != nil && m.Height != nil {
		return *m.Height
	}
	return sortkind.SortKind_ASC
}

func (m *VideoSort) GetFrameRate() sortkind.SortKind {
	if m != nil && m.FrameRate != nil {
		return *m.FrameRate
	}
	return sortkind.SortKind_ASC
}

func (m *VideoSort) GetDuration() *durationvaluesort.DurationValueSort {
	if m != nil {
		return m.Duration
	}
	return nil
}

func (m *VideoSort) GetDescription() *textsort.TextSort {
	if m != nil {
		return m.Description
	}
	return nil
}

func (m *VideoSort) GetHash() string {
	if m != nil && m.Hash != nil {
		return *m.Hash
	}
	return ""
}

func init() {
	proto.RegisterType((*VideoSort)(nil), "videosort.VideoSort")
}

func init() { proto.RegisterFile("videosort/videosort_.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 310 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x92, 0x4f, 0x4b, 0xc3, 0x30,
	0x18, 0xc6, 0xe9, 0xfe, 0x69, 0x33, 0x10, 0xc9, 0x41, 0xc2, 0x4e, 0x63, 0x78, 0x28, 0x03, 0x53,
	0x1d, 0x8a, 0x57, 0x11, 0x6f, 0xde, 0xaa, 0xee, 0xe0, 0x45, 0x62, 0x13, 0x9b, 0x60, 0xdb, 0x94,
	0xf4, 0xed, 0x9c, 0x9f, 0xcf, 0x2f, 0x26, 0x89, 0x49, 0x27, 0x8c, 0x9d, 0xfa, 0xf4, 0x79, 0x7e,
	0xef, 0xcb, 0x93, 0x10, 0x34, 0xdb, 0x28, 0x2e, 0x74, 0xab, 0x0d, 0xa4, 0xbd, 0x7a, 0xa3, 0x8d,
	0xd1, 0xa0, 0x71, 0xdc, 0x3b, 0xb3, 0xb3, 0xce, 0x94, 0x0e, 0xf2, 0x5f, 0x8f, 0xcc, 0x88, 0xfd,
	0xf9, 0x54, 0x35, 0x4f, 0x83, 0x08, 0xc9, 0x92, 0x77, 0x86, 0x81, 0xd2, 0xf5, 0x86, 0x95, 0x9d,
	0x70, 0xb3, 0x7b, 0x4e, 0xbf, 0x05, 0xc4, 0x16, 0x1c, 0x12, 0x84, 0x4f, 0x16, 0x3f, 0x03, 0x14,
	0xaf, 0x6d, 0x8b, 0x27, 0x6d, 0x00, 0x2f, 0xd0, 0xb0, 0x33, 0x25, 0x89, 0xe6, 0x51, 0x32, 0x5d,
	0x9d, 0x52, 0xdf, 0x85, 0xbe, 0x98, 0xd2, 0xc6, 0x99, 0x0d, 0x71, 0x82, 0xc6, 0x5f, 0x8a, 0x83,
	0x24, 0x83, 0x79, 0x94, 0x9c, 0xac, 0x30, 0x0d, 0xc5, 0xa8, 0x65, 0x1e, 0x55, 0xcd, 0xb3, 0x3f,
	0x00, 0x2f, 0xd1, 0x44, 0x0a, 0x55, 0x48, 0x20, 0xc3, 0x83, 0xa8, 0x27, 0xf0, 0x25, 0x8a, 0x3f,
	0x0c, 0xab, 0x44, 0xc6, 0x40, 0x90, 0xd1, 0x41, 0x7c, 0x07, 0xe1, 0x3b, 0x74, 0x1c, 0xce, 0x4b,
	0xc6, 0xae, 0xf0, 0x39, 0xdd, 0xbb, 0x00, 0xfa, 0xe0, 0x9d, 0xb5, 0x75, 0xdc, 0x21, 0xfa, 0x29,
	0x7c, 0x8d, 0xa6, 0x5c, 0xb4, 0xb9, 0x51, 0x8d, 0x5b, 0x32, 0x71, 0x4b, 0x30, 0x0d, 0x57, 0x44,
	0x9f, 0xc5, 0x16, 0xdc, 0xc8, 0x7f, 0x0c, 0x63, 0x34, 0x92, 0xac, 0x95, 0xe4, 0x68, 0x1e, 0x25,
	0x71, 0xe6, 0xf4, 0xfd, 0xed, 0xeb, 0x4d, 0xa1, 0x40, 0x76, 0xef, 0x34, 0xd7, 0x55, 0xba, 0xba,
	0x6a, 0x41, 0xe9, 0xb4, 0xd0, 0x17, 0x4d, 0xc9, 0xbe, 0x0b, 0xa3, 0xbb, 0x9a, 0xa7, 0x85, 0x69,
	0xf2, 0xb4, 0xcd, 0xa5, 0xa8, 0xd8, 0xee, 0x21, 0xfc, 0x06, 0x00, 0x00, 0xff, 0xff, 0x77, 0x9f,
	0xc8, 0xd1, 0x1e, 0x02, 0x00, 0x00,
}
