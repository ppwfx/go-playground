// Code generated by protoc-gen-go. DO NOT EDIT.
// source: textfilter/textfilter_.proto

/*
Package textfilter is a generated protocol buffer package.

It is generated from these files:
	textfilter/textfilter_.proto

It has these top-level messages:
	TextFilter
*/
package textfilter

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import languagefilter "github.com/21stio/go-playground/grpc/schema/languagefilter"
import formattingkindfilter "github.com/21stio/go-playground/grpc/schema/formattingkindfilter"
import stringfilter "github.com/21stio/go-playground/grpc/schema/stringfilter"
import translationfilter "github.com/21stio/go-playground/grpc/schema/translationfilter"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type TextFilter struct {
	Language          *languagefilter.LanguageFilter             `protobuf:"bytes,1,opt,name=language" json:"language,omitempty"`
	Formatting        *formattingkindfilter.FormattingKindFilter `protobuf:"bytes,2,opt,name=formatting" json:"formatting,omitempty"`
	Value             *stringfilter.StringFilter                 `protobuf:"bytes,3,opt,name=value" json:"value,omitempty"`
	TranslationsSome  *translationfilter.TranslationFilter       `protobuf:"bytes,4,opt,name=translationsSome" json:"translationsSome,omitempty"`
	TranslationsEvery *translationfilter.TranslationFilter       `protobuf:"bytes,5,opt,name=translationsEvery" json:"translationsEvery,omitempty"`
	TranslationsNone  *translationfilter.TranslationFilter       `protobuf:"bytes,6,opt,name=translationsNone" json:"translationsNone,omitempty"`
	And               []*TextFilter                              `protobuf:"bytes,7,rep,name=and" json:"and,omitempty"`
	Or                []*TextFilter                              `protobuf:"bytes,8,rep,name=or" json:"or,omitempty"`
	Not               []*TextFilter                              `protobuf:"bytes,9,rep,name=not" json:"not,omitempty"`
	Hash              *string                                    `protobuf:"bytes,10,opt,name=hash" json:"hash,omitempty"`
	XXX_unrecognized  []byte                                     `json:"-"`
}

func (m *TextFilter) Reset()                    { *m = TextFilter{} }
func (m *TextFilter) String() string            { return proto.CompactTextString(m) }
func (*TextFilter) ProtoMessage()               {}
func (*TextFilter) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *TextFilter) GetLanguage() *languagefilter.LanguageFilter {
	if m != nil {
		return m.Language
	}
	return nil
}

func (m *TextFilter) GetFormatting() *formattingkindfilter.FormattingKindFilter {
	if m != nil {
		return m.Formatting
	}
	return nil
}

func (m *TextFilter) GetValue() *stringfilter.StringFilter {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *TextFilter) GetTranslationsSome() *translationfilter.TranslationFilter {
	if m != nil {
		return m.TranslationsSome
	}
	return nil
}

func (m *TextFilter) GetTranslationsEvery() *translationfilter.TranslationFilter {
	if m != nil {
		return m.TranslationsEvery
	}
	return nil
}

func (m *TextFilter) GetTranslationsNone() *translationfilter.TranslationFilter {
	if m != nil {
		return m.TranslationsNone
	}
	return nil
}

func (m *TextFilter) GetAnd() []*TextFilter {
	if m != nil {
		return m.And
	}
	return nil
}

func (m *TextFilter) GetOr() []*TextFilter {
	if m != nil {
		return m.Or
	}
	return nil
}

func (m *TextFilter) GetNot() []*TextFilter {
	if m != nil {
		return m.Not
	}
	return nil
}

func (m *TextFilter) GetHash() string {
	if m != nil && m.Hash != nil {
		return *m.Hash
	}
	return ""
}

func init() {
	proto.RegisterType((*TextFilter)(nil), "textfilter.TextFilter")
}

func init() { proto.RegisterFile("textfilter/textfilter_.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 363 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x93, 0x4d, 0x4b, 0xf3, 0x40,
	0x14, 0x85, 0xe9, 0xd7, 0xfb, 0xb6, 0xd3, 0x8d, 0xce, 0x42, 0x86, 0x22, 0x12, 0xa4, 0x48, 0x28,
	0x38, 0xa9, 0x5d, 0x88, 0xb8, 0x14, 0xec, 0x42, 0x45, 0x24, 0xed, 0xca, 0x8d, 0x8c, 0xcd, 0x74,
	0x32, 0x98, 0xcc, 0x94, 0xc9, 0x4d, 0x69, 0xff, 0xa4, 0xbf, 0x49, 0x9a, 0x4e, 0x9a, 0xf4, 0x83,
	0x82, 0xee, 0xee, 0x3d, 0x9c, 0xe7, 0x70, 0x2e, 0xc9, 0xa0, 0x73, 0xe0, 0x0b, 0x98, 0xca, 0x08,
	0xb8, 0xf1, 0x8a, 0xf1, 0x83, 0xce, 0x8c, 0x06, 0x8d, 0x51, 0x21, 0x75, 0xba, 0x11, 0x53, 0x22,
	0x65, 0x82, 0x5b, 0xf7, 0xf6, 0x6a, 0x89, 0x4e, 0x7f, 0xaa, 0x4d, 0xcc, 0x00, 0xa4, 0x12, 0x5f,
	0x52, 0x05, 0xd6, 0x7b, 0x48, 0xcc, 0x09, 0x27, 0x01, 0x23, 0x95, 0xb0, 0xce, 0xf2, 0x92, 0x3b,
	0x7a, 0x60, 0x98, 0x4a, 0x22, 0x06, 0x52, 0xab, 0xbc, 0xea, 0xae, 0x62, 0xbd, 0x97, 0xdf, 0x75,
	0x84, 0xc6, 0x7c, 0x01, 0xc3, 0x4c, 0xc5, 0xf7, 0xa8, 0x99, 0xf7, 0x24, 0x15, 0xa7, 0xe2, 0xb6,
	0x07, 0x17, 0x74, 0xbb, 0x38, 0x7d, 0xb1, 0xeb, 0x9a, 0xf0, 0x37, 0x7e, 0xfc, 0x84, 0x50, 0xd1,
	0x9b, 0x54, 0x33, 0xba, 0x47, 0x0f, 0x9d, 0x42, 0x87, 0x1b, 0xf1, 0x59, 0xaa, 0xc0, 0x26, 0x95,
	0x68, 0xdc, 0x47, 0x8d, 0x39, 0x8b, 0x52, 0x4e, 0x6a, 0x59, 0x4c, 0x87, 0x96, 0xef, 0xa4, 0xa3,
	0x6c, 0xb1, 0xd8, 0xda, 0x88, 0xdf, 0xd0, 0x49, 0xe9, 0xc8, 0x64, 0xa4, 0x63, 0x4e, 0xea, 0x19,
	0xdc, 0xa5, 0x7b, 0xd7, 0xd3, 0x71, 0xa1, 0xd8, 0x98, 0x3d, 0x1a, 0xfb, 0xe8, 0xb4, 0xac, 0x3d,
	0xce, 0xb9, 0x59, 0x92, 0xc6, 0x2f, 0x22, 0xf7, 0xf1, 0xdd, 0x96, 0xaf, 0x5a, 0x71, 0xf2, 0xef,
	0xaf, 0x2d, 0x57, 0x34, 0x76, 0x51, 0x8d, 0xa9, 0x80, 0xfc, 0x77, 0x6a, 0x6e, 0x7b, 0x70, 0x46,
	0x8b, 0x1f, 0x90, 0x16, 0x9f, 0xd5, 0x5f, 0x59, 0xf0, 0x15, 0xaa, 0x6a, 0x43, 0x9a, 0x47, 0x8d,
	0x55, 0x6d, 0x56, 0x89, 0x4a, 0x03, 0x69, 0x1d, 0x4f, 0x54, 0x1a, 0x30, 0x46, 0xf5, 0x90, 0x25,
	0x21, 0x41, 0x4e, 0xc5, 0x6d, 0xf9, 0xd9, 0xfc, 0x70, 0xf7, 0x7e, 0x2b, 0x24, 0x84, 0xe9, 0x27,
	0x9d, 0xe8, 0xd8, 0x1b, 0xdc, 0x24, 0x20, 0xb5, 0x27, 0xf4, 0xf5, 0x2c, 0x62, 0x4b, 0x61, 0x74,
	0xaa, 0x02, 0x4f, 0x98, 0xd9, 0xc4, 0x4b, 0x26, 0x21, 0x8f, 0x59, 0xe9, 0x0d, 0xfd, 0x04, 0x00,
	0x00, 0xff, 0xff, 0x8c, 0x73, 0xff, 0x57, 0x5b, 0x03, 0x00, 0x00,
}