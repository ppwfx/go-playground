// Code generated by protoc-gen-go. DO NOT EDIT.
// source: deletelocationsresponseselect/deletelocationsresponseselect_.proto

/*
Package deletelocationsresponseselect is a generated protocol buffer package.

It is generated from these files:
	deletelocationsresponseselect/deletelocationsresponseselect_.proto

It has these top-level messages:
	DeleteLocationsResponseSelect
*/
package deletelocationsresponseselect

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import responsemetaselect "github.com/21stio/go-playground/grpc/schema/responsemetaselect"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type DeleteLocationsResponseSelect struct {
	SelectAll        *bool                                  `protobuf:"varint,1,opt,name=selectAll" json:"selectAll,omitempty"`
	Meta             *responsemetaselect.ResponseMetaSelect `protobuf:"bytes,2,opt,name=meta" json:"meta,omitempty"`
	SelectHash       *bool                                  `protobuf:"varint,3,opt,name=selectHash" json:"selectHash,omitempty"`
	Hash             *string                                `protobuf:"bytes,4,opt,name=hash" json:"hash,omitempty"`
	XXX_unrecognized []byte                                 `json:"-"`
}

func (m *DeleteLocationsResponseSelect) Reset()                    { *m = DeleteLocationsResponseSelect{} }
func (m *DeleteLocationsResponseSelect) String() string            { return proto.CompactTextString(m) }
func (*DeleteLocationsResponseSelect) ProtoMessage()               {}
func (*DeleteLocationsResponseSelect) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *DeleteLocationsResponseSelect) GetSelectAll() bool {
	if m != nil && m.SelectAll != nil {
		return *m.SelectAll
	}
	return false
}

func (m *DeleteLocationsResponseSelect) GetMeta() *responsemetaselect.ResponseMetaSelect {
	if m != nil {
		return m.Meta
	}
	return nil
}

func (m *DeleteLocationsResponseSelect) GetSelectHash() bool {
	if m != nil && m.SelectHash != nil {
		return *m.SelectHash
	}
	return false
}

func (m *DeleteLocationsResponseSelect) GetHash() string {
	if m != nil && m.Hash != nil {
		return *m.Hash
	}
	return ""
}

func init() {
	proto.RegisterType((*DeleteLocationsResponseSelect)(nil), "deletelocationsresponseselect.DeleteLocationsResponseSelect")
}

func init() {
	proto.RegisterFile("deletelocationsresponseselect/deletelocationsresponseselect_.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 229 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0xd0, 0x31, 0x4b, 0xc4, 0x30,
	0x14, 0x07, 0x70, 0xa2, 0x1d, 0xbc, 0xb8, 0x65, 0x2a, 0xe2, 0x49, 0x71, 0x90, 0x0e, 0xda, 0xe0,
	0x8d, 0x6e, 0x1e, 0x0e, 0x8a, 0xba, 0xc4, 0xcd, 0x45, 0x62, 0xee, 0x91, 0x14, 0xd2, 0xbe, 0x90,
	0xbc, 0x1b, 0xfc, 0x50, 0x7e, 0x47, 0xb9, 0xb4, 0xc5, 0x83, 0x42, 0xb7, 0xe4, 0xff, 0xfe, 0xef,
	0x17, 0x08, 0xdf, 0xee, 0xc0, 0x03, 0x81, 0x47, 0xa3, 0xa9, 0xc5, 0x3e, 0x45, 0x48, 0x01, 0xfb,
	0x04, 0x09, 0x3c, 0x18, 0x92, 0x8b, 0xd3, 0xaf, 0x26, 0x44, 0x24, 0x14, 0xeb, 0xc5, 0xd6, 0xc5,
	0xed, 0x74, 0xef, 0x80, 0xf4, 0xe8, 0xce, 0xa3, 0x11, 0xbb, 0xfe, 0x65, 0x7c, 0xfd, 0x94, 0xbd,
	0xb7, 0xc9, 0x53, 0x63, 0xf9, 0x23, 0x17, 0xc5, 0x25, 0x5f, 0x0d, 0x2b, 0x8f, 0xde, 0x97, 0xac,
	0x62, 0xf5, 0x99, 0xfa, 0x0f, 0xc4, 0x03, 0x2f, 0x0e, 0x68, 0x79, 0x52, 0xb1, 0xfa, 0x7c, 0x73,
	0xd3, 0xcc, 0x5f, 0x6a, 0x26, 0xef, 0x1d, 0x48, 0x0f, 0xa6, 0xca, 0x3b, 0xe2, 0x8a, 0xf3, 0xa1,
	0xf2, 0xac, 0x93, 0x2b, 0x4f, 0x33, 0x7d, 0x94, 0x08, 0xc1, 0x0b, 0x77, 0x98, 0x14, 0x15, 0xab,
	0x57, 0x2a, 0x9f, 0xb7, 0xaf, 0x9f, 0x2f, 0xb6, 0x25, 0xb7, 0xff, 0x6e, 0x0c, 0x76, 0x72, 0x73,
	0x9f, 0xa8, 0x45, 0x69, 0xf1, 0x2e, 0x78, 0xfd, 0x63, 0x23, 0xee, 0xfb, 0x9d, 0xb4, 0x31, 0x18,
	0x99, 0x8c, 0x83, 0x4e, 0x2f, 0x7f, 0xe8, 0x5f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x79, 0x20, 0x0e,
	0x40, 0x8e, 0x01, 0x00, 0x00,
}
