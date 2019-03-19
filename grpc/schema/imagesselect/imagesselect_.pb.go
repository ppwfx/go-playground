// Code generated by protoc-gen-go. DO NOT EDIT.
// source: imagesselect/imagesselect_.proto

/*
Package imagesselect is a generated protocol buffer package.

It is generated from these files:
	imagesselect/imagesselect_.proto

It has these top-level messages:
	ImagesSelect
*/
package imagesselect

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import imagefilter "github.com/21stio/go-playground/grpc/schema/imagefilter"
import imagesort "github.com/21stio/go-playground/grpc/schema/imagesort"
import page "github.com/21stio/go-playground/grpc/schema/page"
import urlselect "github.com/21stio/go-playground/grpc/schema/urlselect"
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

type ImagesSelect struct {
	Filter           *imagefilter.ImageFilter `protobuf:"bytes,1,opt,name=filter" json:"filter,omitempty"`
	Sort             *imagesort.ImageSort     `protobuf:"bytes,2,opt,name=sort" json:"sort,omitempty"`
	Page             *page.Page               `protobuf:"bytes,3,opt,name=page" json:"page,omitempty"`
	SelectAll        *bool                    `protobuf:"varint,4,opt,name=selectAll" json:"selectAll,omitempty"`
	Kind             *bool                    `protobuf:"varint,5,opt,name=kind" json:"kind,omitempty"`
	Url              *urlselect.UrlSelect     `protobuf:"bytes,6,opt,name=url" json:"url,omitempty"`
	Width            *bool                    `protobuf:"varint,7,opt,name=width" json:"width,omitempty"`
	Height           *bool                    `protobuf:"varint,8,opt,name=height" json:"height,omitempty"`
	Description      *textselect.TextSelect   `protobuf:"bytes,9,opt,name=description" json:"description,omitempty"`
	SelectHash       *bool                    `protobuf:"varint,10,opt,name=selectHash" json:"selectHash,omitempty"`
	Hash             *string                  `protobuf:"bytes,11,opt,name=hash" json:"hash,omitempty"`
	XXX_unrecognized []byte                   `json:"-"`
}

func (m *ImagesSelect) Reset()                    { *m = ImagesSelect{} }
func (m *ImagesSelect) String() string            { return proto.CompactTextString(m) }
func (*ImagesSelect) ProtoMessage()               {}
func (*ImagesSelect) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *ImagesSelect) GetFilter() *imagefilter.ImageFilter {
	if m != nil {
		return m.Filter
	}
	return nil
}

func (m *ImagesSelect) GetSort() *imagesort.ImageSort {
	if m != nil {
		return m.Sort
	}
	return nil
}

func (m *ImagesSelect) GetPage() *page.Page {
	if m != nil {
		return m.Page
	}
	return nil
}

func (m *ImagesSelect) GetSelectAll() bool {
	if m != nil && m.SelectAll != nil {
		return *m.SelectAll
	}
	return false
}

func (m *ImagesSelect) GetKind() bool {
	if m != nil && m.Kind != nil {
		return *m.Kind
	}
	return false
}

func (m *ImagesSelect) GetUrl() *urlselect.UrlSelect {
	if m != nil {
		return m.Url
	}
	return nil
}

func (m *ImagesSelect) GetWidth() bool {
	if m != nil && m.Width != nil {
		return *m.Width
	}
	return false
}

func (m *ImagesSelect) GetHeight() bool {
	if m != nil && m.Height != nil {
		return *m.Height
	}
	return false
}

func (m *ImagesSelect) GetDescription() *textselect.TextSelect {
	if m != nil {
		return m.Description
	}
	return nil
}

func (m *ImagesSelect) GetSelectHash() bool {
	if m != nil && m.SelectHash != nil {
		return *m.SelectHash
	}
	return false
}

func (m *ImagesSelect) GetHash() string {
	if m != nil && m.Hash != nil {
		return *m.Hash
	}
	return ""
}

func init() {
	proto.RegisterType((*ImagesSelect)(nil), "imagesselect.ImagesSelect")
}

func init() { proto.RegisterFile("imagesselect/imagesselect_.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 357 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x92, 0xc1, 0x4f, 0xb3, 0x40,
	0x10, 0xc5, 0xc3, 0x57, 0xda, 0xaf, 0x9d, 0xf6, 0x60, 0x36, 0x4d, 0xb3, 0x21, 0x4d, 0x43, 0x3c,
	0x18, 0x2e, 0x2e, 0xda, 0x53, 0xe3, 0x4d, 0x0f, 0x46, 0x6f, 0x86, 0xea, 0xc5, 0x8b, 0x41, 0x58,
	0x97, 0x8d, 0xdb, 0x2e, 0x59, 0x96, 0x58, 0x6f, 0xfe, 0xe9, 0x86, 0x61, 0x4b, 0xf1, 0x42, 0xde,
	0xbc, 0xf9, 0x31, 0xf3, 0x86, 0x00, 0xa1, 0xdc, 0xa5, 0x82, 0x57, 0x15, 0x57, 0x3c, 0xb3, 0x71,
	0xbf, 0x78, 0x63, 0xa5, 0xd1, 0x56, 0x93, 0x59, 0xdf, 0x0c, 0x56, 0x58, 0x7d, 0x48, 0x65, 0xb9,
	0x89, 0x7b, 0xda, 0xd1, 0x41, 0xd0, 0xd2, 0xda, 0x1c, 0x87, 0x69, 0x73, 0x9c, 0x14, 0x9c, 0x95,
	0xa9, 0xe0, 0x71, 0xf3, 0xe8, 0xe8, 0xda, 0x28, 0xb7, 0xba, 0x53, 0xc7, 0xde, 0xd2, 0xf2, 0x83,
	0x75, 0xcd, 0x93, 0x74, 0xdd, 0xf3, 0x9f, 0x01, 0xcc, 0x1e, 0x71, 0xc1, 0x16, 0x7d, 0x72, 0x05,
	0xa3, 0x36, 0x09, 0xf5, 0x42, 0x2f, 0x9a, 0xae, 0x29, 0xeb, 0xa5, 0x63, 0x88, 0xde, 0xa3, 0x4e,
	0x1c, 0x47, 0x22, 0xf0, 0x9b, 0x74, 0xf4, 0x1f, 0xf2, 0x73, 0xd6, 0xe5, 0x6d, 0xe9, 0xad, 0x36,
	0x36, 0x41, 0x82, 0xac, 0xc0, 0x6f, 0x52, 0xd3, 0x01, 0x92, 0xc0, 0x9a, 0x82, 0x3d, 0xa5, 0x82,
	0x27, 0xe8, 0x93, 0x25, 0x4c, 0xda, 0x74, 0xb7, 0x4a, 0x51, 0x3f, 0xf4, 0xa2, 0x71, 0x72, 0x32,
	0x08, 0x01, 0xff, 0x53, 0xee, 0x73, 0x3a, 0xc4, 0x06, 0x6a, 0x72, 0x01, 0x83, 0xda, 0x28, 0x3a,
	0x72, 0xab, 0xbb, 0xe3, 0xd9, 0x8b, 0x51, 0xed, 0x41, 0x49, 0x03, 0x90, 0x39, 0x0c, 0xbf, 0x64,
	0x6e, 0x0b, 0xfa, 0x1f, 0x5f, 0x6e, 0x0b, 0xb2, 0x80, 0x51, 0xc1, 0xa5, 0x28, 0x2c, 0x1d, 0xa3,
	0xed, 0x2a, 0xb2, 0x81, 0x69, 0xce, 0xab, 0xcc, 0xc8, 0xd2, 0x4a, 0xbd, 0xa7, 0x13, 0x9c, 0xbe,
	0x60, 0xa7, 0xaf, 0xc7, 0x9e, 0xf9, 0xc1, 0xba, 0xf9, 0x7d, 0x94, 0xac, 0x00, 0x5a, 0xe2, 0x21,
	0xad, 0x0a, 0x0a, 0x38, 0xb5, 0xe7, 0x34, 0x37, 0x14, 0x4d, 0x67, 0x1a, 0x7a, 0xd1, 0x24, 0x41,
	0x7d, 0x77, 0xf3, 0xba, 0x11, 0xd2, 0x16, 0xf5, 0x3b, 0xcb, 0xf4, 0x2e, 0x5e, 0x5f, 0x57, 0x56,
	0xea, 0x58, 0xe8, 0xcb, 0x52, 0xa5, 0xdf, 0xc2, 0xe8, 0x7a, 0x9f, 0xc7, 0xc2, 0x94, 0x59, 0x5c,
	0x65, 0x05, 0xdf, 0xa5, 0x7f, 0xfe, 0xad, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xb8, 0x1f, 0x49,
	0xd6, 0x77, 0x02, 0x00, 0x00,
}