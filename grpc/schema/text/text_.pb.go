// Code generated by protoc-gen-go. DO NOT EDIT.
// source: text/text_.proto

/*
Package text is a generated protocol buffer package.

It is generated from these files:
	text/text_.proto

It has these top-level messages:
	Text
*/
package text

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import language "github.com/21stio/go-playground/grpc/schema/language"
import formattingkind "github.com/21stio/go-playground/grpc/schema/formattingkind"
import translation "github.com/21stio/go-playground/grpc/schema/translation"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Text struct {
	Language         *language.Language             `protobuf:"varint,1,opt,name=language,enum=language.Language" json:"language,omitempty"`
	Formatting       *formattingkind.FormattingKind `protobuf:"varint,2,opt,name=formatting,enum=formattingkind.FormattingKind" json:"formatting,omitempty"`
	Value            *string                        `protobuf:"bytes,3,opt,name=value" json:"value,omitempty"`
	Translations     []*translation.Translation     `protobuf:"bytes,4,rep,name=translations" json:"translations,omitempty"`
	Hash             *string                        `protobuf:"bytes,5,opt,name=hash" json:"hash,omitempty"`
	XXX_unrecognized []byte                         `json:"-"`
}

func (m *Text) Reset()                    { *m = Text{} }
func (m *Text) String() string            { return proto.CompactTextString(m) }
func (*Text) ProtoMessage()               {}
func (*Text) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Text) GetLanguage() language.Language {
	if m != nil && m.Language != nil {
		return *m.Language
	}
	return language.Language_DE
}

func (m *Text) GetFormatting() formattingkind.FormattingKind {
	if m != nil && m.Formatting != nil {
		return *m.Formatting
	}
	return formattingkind.FormattingKind_HTML
}

func (m *Text) GetValue() string {
	if m != nil && m.Value != nil {
		return *m.Value
	}
	return ""
}

func (m *Text) GetTranslations() []*translation.Translation {
	if m != nil {
		return m.Translations
	}
	return nil
}

func (m *Text) GetHash() string {
	if m != nil && m.Hash != nil {
		return *m.Hash
	}
	return ""
}

func init() {
	proto.RegisterType((*Text)(nil), "text.Text")
}

func init() { proto.RegisterFile("text/text_.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 255 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x91, 0x31, 0x4f, 0xc3, 0x30,
	0x10, 0x85, 0x15, 0x9a, 0x4a, 0x60, 0x10, 0x42, 0x16, 0x83, 0xd5, 0xa1, 0x8a, 0x10, 0x43, 0x16,
	0xce, 0x90, 0x19, 0x31, 0x30, 0xb0, 0xc0, 0x14, 0x75, 0x62, 0x41, 0x47, 0x62, 0x1c, 0x8b, 0xc4,
	0x8e, 0xec, 0x0b, 0x2a, 0x7f, 0x96, 0xdf, 0x82, 0x9a, 0x26, 0x6d, 0xb2, 0x58, 0xdf, 0xbb, 0x7b,
	0xe7, 0x7b, 0xd2, 0xb1, 0x2b, 0x52, 0x5b, 0x92, 0xbb, 0xe7, 0x03, 0x5a, 0xef, 0xc8, 0xf1, 0x78,
	0x27, 0x56, 0xa2, 0x46, 0xab, 0x3b, 0xd4, 0x4a, 0x8e, 0x30, 0xf4, 0x57, 0xb7, 0x5f, 0xce, 0x37,
	0x48, 0x64, 0xac, 0xfe, 0x36, 0xb6, 0x94, 0x73, 0x39, 0xba, 0xd6, 0xe4, 0xd1, 0x86, 0x1a, 0xc9,
	0x38, 0x2b, 0x27, 0x3c, 0xf4, 0x6f, 0xfe, 0x22, 0x16, 0x6f, 0xd4, 0x96, 0x38, 0xb0, 0xd3, 0x71,
	0x83, 0x88, 0x92, 0x28, 0xbd, 0xcc, 0x38, 0x8c, 0x05, 0x78, 0x1b, 0x20, 0x3f, 0x78, 0xf8, 0x13,
	0x63, 0xc7, 0x8d, 0xe2, 0xa4, 0x9f, 0x58, 0xc3, 0x3c, 0x04, 0xbc, 0x1c, 0xe4, 0xab, 0xb1, 0x65,
	0x3e, 0x99, 0xe0, 0xd7, 0x6c, 0xf9, 0x83, 0x75, 0xa7, 0xc4, 0x22, 0x89, 0xd2, 0xb3, 0x7c, 0x2f,
	0xf8, 0x23, 0xbb, 0x98, 0x84, 0x0c, 0x22, 0x4e, 0x16, 0xe9, 0x79, 0x26, 0x60, 0x52, 0x84, 0xcd,
	0x91, 0xf3, 0x99, 0x9b, 0x73, 0x16, 0x57, 0x18, 0x2a, 0xb1, 0xec, 0xbf, 0xec, 0xf9, 0x39, 0x7b,
	0xbf, 0xd7, 0x86, 0xaa, 0xee, 0x13, 0x0a, 0xd7, 0xc8, 0xec, 0x21, 0x90, 0x71, 0x52, 0xbb, 0xbb,
	0xb6, 0xc6, 0x5f, 0xed, 0x5d, 0x67, 0x4b, 0xa9, 0x7d, 0x5b, 0xc8, 0x50, 0x54, 0xaa, 0xc1, 0xfe,
	0x02, 0xff, 0x01, 0x00, 0x00, 0xff, 0xff, 0x47, 0x9d, 0x29, 0xc5, 0x8d, 0x01, 0x00, 0x00,
}
