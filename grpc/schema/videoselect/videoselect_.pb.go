// Code generated by protoc-gen-go. DO NOT EDIT.
// source: videoselect/videoselect_.proto

/*
Package videoselect is a generated protocol buffer package.

It is generated from these files:
	videoselect/videoselect_.proto

It has these top-level messages:
	VideoSelect
*/
package videoselect

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import urlselect "github.com/21stio/go-playground/grpc/schema/urlselect"
import durationvalueselect "github.com/21stio/go-playground/grpc/schema/durationvalueselect"
import textselect "github.com/21stio/go-playground/grpc/schema/textselect"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type VideoSelect struct {
	SelectAll        *bool                                    `protobuf:"varint,1,opt,name=selectAll" json:"selectAll,omitempty"`
	Url              *urlselect.UrlSelect                     `protobuf:"bytes,2,opt,name=url" json:"url,omitempty"`
	Width            *bool                                    `protobuf:"varint,3,opt,name=width" json:"width,omitempty"`
	Height           *bool                                    `protobuf:"varint,4,opt,name=height" json:"height,omitempty"`
	FrameRate        *bool                                    `protobuf:"varint,5,opt,name=frameRate" json:"frameRate,omitempty"`
	Duration         *durationvalueselect.DurationValueSelect `protobuf:"bytes,6,opt,name=duration" json:"duration,omitempty"`
	Description      *textselect.TextSelect                   `protobuf:"bytes,7,opt,name=description" json:"description,omitempty"`
	SelectHash       *bool                                    `protobuf:"varint,8,opt,name=selectHash" json:"selectHash,omitempty"`
	Hash             *string                                  `protobuf:"bytes,9,opt,name=hash" json:"hash,omitempty"`
	XXX_unrecognized []byte                                   `json:"-"`
}

func (m *VideoSelect) Reset()                    { *m = VideoSelect{} }
func (m *VideoSelect) String() string            { return proto.CompactTextString(m) }
func (*VideoSelect) ProtoMessage()               {}
func (*VideoSelect) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *VideoSelect) GetSelectAll() bool {
	if m != nil && m.SelectAll != nil {
		return *m.SelectAll
	}
	return false
}

func (m *VideoSelect) GetUrl() *urlselect.UrlSelect {
	if m != nil {
		return m.Url
	}
	return nil
}

func (m *VideoSelect) GetWidth() bool {
	if m != nil && m.Width != nil {
		return *m.Width
	}
	return false
}

func (m *VideoSelect) GetHeight() bool {
	if m != nil && m.Height != nil {
		return *m.Height
	}
	return false
}

func (m *VideoSelect) GetFrameRate() bool {
	if m != nil && m.FrameRate != nil {
		return *m.FrameRate
	}
	return false
}

func (m *VideoSelect) GetDuration() *durationvalueselect.DurationValueSelect {
	if m != nil {
		return m.Duration
	}
	return nil
}

func (m *VideoSelect) GetDescription() *textselect.TextSelect {
	if m != nil {
		return m.Description
	}
	return nil
}

func (m *VideoSelect) GetSelectHash() bool {
	if m != nil && m.SelectHash != nil {
		return *m.SelectHash
	}
	return false
}

func (m *VideoSelect) GetHash() string {
	if m != nil && m.Hash != nil {
		return *m.Hash
	}
	return ""
}

func init() {
	proto.RegisterType((*VideoSelect)(nil), "videoselect.VideoSelect")
}

func init() { proto.RegisterFile("videoselect/videoselect_.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 313 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x92, 0xcf, 0x4f, 0x83, 0x30,
	0x14, 0xc7, 0xc3, 0x7e, 0x09, 0xe5, 0xd6, 0x2c, 0x4b, 0x43, 0x96, 0x85, 0x78, 0x30, 0x5c, 0x2c,
	0x71, 0x17, 0xf5, 0xa8, 0xd9, 0xc1, 0x33, 0xea, 0x0e, 0x5e, 0x4c, 0x85, 0x4a, 0x9b, 0x94, 0x95,
	0x94, 0x76, 0xce, 0xbf, 0xd3, 0x7f, 0xc8, 0x50, 0x3a, 0xe0, 0xb0, 0xdb, 0xeb, 0xa7, 0x9f, 0xbe,
	0xef, 0xe3, 0x05, 0xb0, 0x39, 0xf2, 0x82, 0xca, 0x86, 0x0a, 0x9a, 0xeb, 0x74, 0x54, 0x7f, 0xe2,
	0x5a, 0x49, 0x2d, 0x61, 0x38, 0x62, 0x51, 0x64, 0x94, 0x70, 0x6a, 0x5f, 0x39, 0x31, 0xc2, 0x85,
	0x51, 0x44, 0x73, 0x79, 0x38, 0x12, 0x61, 0xa8, 0xb3, 0x2e, 0xb0, 0xb3, 0xbf, 0xd6, 0xf4, 0xa4,
	0x9d, 0x36, 0x94, 0xee, 0xf6, 0xfa, 0x6f, 0x02, 0xc2, 0x7d, 0x9b, 0xfc, 0x6a, 0x31, 0x5c, 0x83,
	0xa0, 0x13, 0x9e, 0x84, 0x40, 0x5e, 0xec, 0x25, 0x7e, 0x36, 0x00, 0x78, 0x03, 0xa6, 0x46, 0x09,
	0x34, 0x89, 0xbd, 0x24, 0xdc, 0x2e, 0x71, 0x3f, 0x1b, 0x7e, 0x57, 0xa2, 0x6b, 0x90, 0xb5, 0x02,
	0x5c, 0x82, 0xf9, 0x0f, 0x2f, 0x34, 0x43, 0x53, 0xdb, 0xa1, 0x3b, 0xc0, 0x15, 0x58, 0x30, 0xca,
	0x4b, 0xa6, 0xd1, 0xcc, 0x62, 0x77, 0x6a, 0x33, 0xbf, 0x15, 0xa9, 0x68, 0x46, 0x34, 0x45, 0xf3,
	0x2e, 0xb3, 0x07, 0x70, 0x07, 0xfc, 0xf3, 0xd7, 0xa1, 0x85, 0x0d, 0x4e, 0x2e, 0xad, 0x00, 0xef,
	0x1c, 0xdb, 0xb7, 0xcc, 0x0d, 0xd3, 0xbf, 0x84, 0x0f, 0x20, 0x2c, 0x68, 0x93, 0x2b, 0x5e, 0xdb,
	0x46, 0x57, 0xb6, 0xd1, 0x0a, 0x0f, 0x0b, 0xc1, 0x6f, 0xf4, 0xa4, 0xdd, 0xb3, 0xb1, 0x0a, 0x37,
	0x00, 0x74, 0xc6, 0x0b, 0x69, 0x18, 0xf2, 0xed, 0x78, 0x23, 0x02, 0x21, 0x98, 0xb1, 0xf6, 0x26,
	0x88, 0xbd, 0x24, 0xc8, 0x6c, 0xfd, 0xfc, 0xf8, 0x71, 0x5f, 0x72, 0xcd, 0xcc, 0x17, 0xce, 0x65,
	0x95, 0x6e, 0xef, 0x1a, 0xcd, 0x65, 0x5a, 0xca, 0xdb, 0x5a, 0x90, 0xdf, 0x52, 0x49, 0x73, 0x28,
	0xd2, 0x52, 0xd5, 0x79, 0xda, 0xe4, 0x8c, 0x56, 0x64, 0xfc, 0x3b, 0xfc, 0x07, 0x00, 0x00, 0xff,
	0xff, 0xe6, 0x57, 0x43, 0xb1, 0x28, 0x02, 0x00, 0x00,
}
