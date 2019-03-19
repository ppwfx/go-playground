// Code generated by protoc-gen-go. DO NOT EDIT.
// source: mediasselect/mediasselect_.proto

/*
Package mediasselect is a generated protocol buffer package.

It is generated from these files:
	mediasselect/mediasselect_.proto

It has these top-level messages:
	MediasSelect
*/
package mediasselect

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import mediafilter "github.com/21stio/go-playground/grpc/schema/mediafilter"
import mediasort "github.com/21stio/go-playground/grpc/schema/mediasort"
import page "github.com/21stio/go-playground/grpc/schema/page"
import imageselect "github.com/21stio/go-playground/grpc/schema/imageselect"
import videoselect "github.com/21stio/go-playground/grpc/schema/videoselect"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type MediasSelect struct {
	Filter           *mediafilter.MediaFilter `protobuf:"bytes,1,opt,name=filter" json:"filter,omitempty"`
	Sort             *mediasort.MediaSort     `protobuf:"bytes,2,opt,name=sort" json:"sort,omitempty"`
	Page             *page.Page               `protobuf:"bytes,3,opt,name=page" json:"page,omitempty"`
	SelectAll        *bool                    `protobuf:"varint,4,opt,name=selectAll" json:"selectAll,omitempty"`
	Kind             *bool                    `protobuf:"varint,5,opt,name=kind" json:"kind,omitempty"`
	Image            *imageselect.ImageSelect `protobuf:"bytes,6,opt,name=image" json:"image,omitempty"`
	Video            *videoselect.VideoSelect `protobuf:"bytes,7,opt,name=video" json:"video,omitempty"`
	SelectHash       *bool                    `protobuf:"varint,8,opt,name=selectHash" json:"selectHash,omitempty"`
	Hash             *string                  `protobuf:"bytes,9,opt,name=hash" json:"hash,omitempty"`
	XXX_unrecognized []byte                   `json:"-"`
}

func (m *MediasSelect) Reset()                    { *m = MediasSelect{} }
func (m *MediasSelect) String() string            { return proto.CompactTextString(m) }
func (*MediasSelect) ProtoMessage()               {}
func (*MediasSelect) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *MediasSelect) GetFilter() *mediafilter.MediaFilter {
	if m != nil {
		return m.Filter
	}
	return nil
}

func (m *MediasSelect) GetSort() *mediasort.MediaSort {
	if m != nil {
		return m.Sort
	}
	return nil
}

func (m *MediasSelect) GetPage() *page.Page {
	if m != nil {
		return m.Page
	}
	return nil
}

func (m *MediasSelect) GetSelectAll() bool {
	if m != nil && m.SelectAll != nil {
		return *m.SelectAll
	}
	return false
}

func (m *MediasSelect) GetKind() bool {
	if m != nil && m.Kind != nil {
		return *m.Kind
	}
	return false
}

func (m *MediasSelect) GetImage() *imageselect.ImageSelect {
	if m != nil {
		return m.Image
	}
	return nil
}

func (m *MediasSelect) GetVideo() *videoselect.VideoSelect {
	if m != nil {
		return m.Video
	}
	return nil
}

func (m *MediasSelect) GetSelectHash() bool {
	if m != nil && m.SelectHash != nil {
		return *m.SelectHash
	}
	return false
}

func (m *MediasSelect) GetHash() string {
	if m != nil && m.Hash != nil {
		return *m.Hash
	}
	return ""
}

func init() {
	proto.RegisterType((*MediasSelect)(nil), "mediasselect.MediasSelect")
}

func init() { proto.RegisterFile("mediasselect/mediasselect_.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 324 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x92, 0xc1, 0x4a, 0xf4, 0x30,
	0x14, 0x85, 0xe9, 0xfc, 0x9d, 0xf9, 0x67, 0xe2, 0x2c, 0x24, 0xb8, 0x08, 0x45, 0x4a, 0x71, 0xd5,
	0x8d, 0x89, 0xce, 0x4a, 0xdc, 0xe9, 0x42, 0x74, 0x21, 0x48, 0x07, 0x5c, 0xb8, 0x91, 0xd8, 0xc6,
	0x34, 0xd8, 0x4e, 0x4a, 0x9a, 0x11, 0x7c, 0x4b, 0x1f, 0x49, 0x72, 0x93, 0x8e, 0x71, 0x53, 0x4e,
	0xce, 0xf9, 0xc8, 0x3d, 0xb9, 0x14, 0x15, 0xbd, 0x68, 0x14, 0x1f, 0x47, 0xd1, 0x89, 0xda, 0xb2,
	0xf8, 0xf0, 0x4a, 0x07, 0xa3, 0xad, 0xc6, 0xeb, 0xd8, 0xcc, 0x72, 0x38, 0xbd, 0xab, 0xce, 0x0a,
	0xc3, 0x22, 0x1d, 0xe8, 0x2c, 0xf3, 0xb4, 0x36, 0xd3, 0x65, 0xda, 0x4c, 0x37, 0x65, 0xc7, 0x03,
	0x97, 0x82, 0xb9, 0xcf, 0xe4, 0xe4, 0xaa, 0xe7, 0x52, 0x84, 0xe1, 0x91, 0x3e, 0xe4, 0x9f, 0xaa,
	0x11, 0x3a, 0xe4, 0x91, 0x0e, 0xf9, 0xd9, 0xf7, 0x0c, 0xad, 0x1f, 0x61, 0xcc, 0x16, 0x7c, 0x7c,
	0x81, 0x16, 0xbe, 0x0f, 0x49, 0x8a, 0xa4, 0x3c, 0xda, 0x10, 0x1a, 0x75, 0xa4, 0x80, 0xde, 0x81,
	0xae, 0x02, 0x87, 0x4b, 0x94, 0xba, 0x8e, 0x64, 0x06, 0xfc, 0x09, 0x3d, 0xb4, 0xf6, 0xf4, 0x56,
	0x1b, 0x5b, 0x01, 0x81, 0x73, 0x94, 0xba, 0xee, 0xe4, 0x1f, 0x90, 0x88, 0xba, 0x03, 0x7d, 0xe2,
	0x52, 0x54, 0xe0, 0xe3, 0x53, 0xb4, 0xf2, 0xed, 0x6e, 0xba, 0x8e, 0xa4, 0x45, 0x52, 0x2e, 0xab,
	0x5f, 0x03, 0x63, 0x94, 0x7e, 0xa8, 0x5d, 0x43, 0xe6, 0x10, 0x80, 0xc6, 0x14, 0xcd, 0xe1, 0xd1,
	0x64, 0x11, 0xca, 0x46, 0x2b, 0xa0, 0x0f, 0x4e, 0xfb, 0x67, 0x55, 0x1e, 0x73, 0x3c, 0x2c, 0x81,
	0xfc, 0x0f, 0x7c, 0xb4, 0x12, 0xfa, 0xec, 0xf4, 0xc4, 0x43, 0x80, 0x73, 0x84, 0x7c, 0x78, 0xcf,
	0xc7, 0x96, 0x2c, 0x61, 0x72, 0xe4, 0xb8, 0x4e, 0xad, 0x4b, 0x56, 0x45, 0x52, 0xae, 0x2a, 0xd0,
	0xb7, 0xd7, 0x2f, 0x57, 0x52, 0xd9, 0x76, 0xff, 0x46, 0x6b, 0xdd, 0xb3, 0xcd, 0xe5, 0x68, 0x95,
	0x66, 0x52, 0x9f, 0x0f, 0x1d, 0xff, 0x92, 0x46, 0xef, 0x77, 0x0d, 0x93, 0x66, 0xa8, 0xd9, 0x58,
	0xb7, 0xa2, 0xe7, 0x7f, 0xfe, 0x98, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x36, 0x49, 0x08, 0x44,
	0x4d, 0x02, 0x00, 0x00,
}