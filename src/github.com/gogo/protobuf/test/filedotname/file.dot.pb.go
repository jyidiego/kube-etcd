// Code generated by protoc-gen-gogo.
// source: file.dot.proto
// DO NOT EDIT!

/*
Package filedotname is a generated protocol buffer package.

It is generated from these files:
	file.dot.proto

It has these top-level messages:
	M
*/
package filedotname

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"

import github_com_gogo_protobuf_protoc_gen_gogo_descriptor "github.com/gogo/protobuf/protoc-gen-gogo/descriptor"
import github_com_gogo_protobuf_proto "github.com/gogo/protobuf/proto"
import compress_gzip "compress/gzip"
import bytes "bytes"
import io_ioutil "io/ioutil"

import strings "strings"
import sort "sort"
import strconv "strconv"
import reflect "reflect"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
const _ = proto.GoGoProtoPackageIsVersion1

type M struct {
	A                *string `protobuf:"bytes,1,opt,name=a" json:"a,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *M) Reset()                    { *m = M{} }
func (*M) ProtoMessage()               {}
func (*M) Descriptor() ([]byte, []int) { return fileDescriptorFileDot, []int{0} }

func init() {
	proto.RegisterType((*M)(nil), "filedotname.M")
}
func (this *M) Description() (desc *github_com_gogo_protobuf_protoc_gen_gogo_descriptor.FileDescriptorSet) {
	return FileDotDescription()
}
func FileDotDescription() (desc *github_com_gogo_protobuf_protoc_gen_gogo_descriptor.FileDescriptorSet) {
	d := &github_com_gogo_protobuf_protoc_gen_gogo_descriptor.FileDescriptorSet{}
	var gzipped = []byte{
		// 3329 bytes of a gzipped FileDescriptorSet
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xcc, 0x5a, 0x5d, 0x6c, 0x23, 0xd5,
		0x15, 0xc6, 0xb1, 0x9d, 0xd8, 0xc7, 0x8e, 0xe3, 0x4c, 0xc2, 0xae, 0x37, 0xc0, 0x2e, 0x1b, 0xfe,
		0x96, 0xa5, 0x64, 0xe9, 0xb6, 0xc0, 0x62, 0x5a, 0x50, 0x7e, 0xbc, 0xc1, 0xab, 0x24, 0x76, 0xc7,
		0x31, 0x2c, 0xf4, 0x61, 0x34, 0x19, 0xdf, 0x38, 0xde, 0x1d, 0xcf, 0xb8, 0x9e, 0xf1, 0xb2, 0xe1,
		0x89, 0x8a, 0xfe, 0x08, 0x55, 0xfd, 0xaf, 0x54, 0xfe, 0x5b, 0x90, 0x5a, 0x28, 0x2d, 0x2d, 0xf4,
		0x4f, 0x55, 0x9f, 0x90, 0x2a, 0x5a, 0x9e, 0xaa, 0xb6, 0x4f, 0x7d, 0xe8, 0x03, 0x50, 0xa4, 0xd2,
		0x96, 0xb6, 0x54, 0x5a, 0xa9, 0x48, 0xfb, 0xd2, 0x73, 0xff, 0xc6, 0x33, 0xb6, 0x93, 0x71, 0x90,
		0x28, 0x5d, 0x69, 0x14, 0xdf, 0x73, 0xcf, 0xf7, 0xcd, 0xbd, 0xe7, 0x9e, 0x7b, 0xce, 0xb9, 0x77,
		0x16, 0x7e, 0xf5, 0x61, 0xb8, 0xb2, 0x6e, 0xdb, 0x75, 0x93, 0x1c, 0x6b, 0xb5, 0x6d, 0xd7, 0xde,
		0xe8, 0x6c, 0x1e, 0xab, 0x11, 0xc7, 0x68, 0x37, 0x5a, 0xae, 0xdd, 0x9e, 0x63, 0x32, 0x65, 0x82,
		0x6b, 0xcc, 0x49, 0x8d, 0xd9, 0x55, 0x98, 0x3c, 0xd9, 0x30, 0xc9, 0x92, 0xa7, 0x58, 0x21, 0xae,
		0x72, 0x02, 0x62, 0x9b, 0x28, 0xcc, 0x45, 0xae, 0x8c, 0x1e, 0x49, 0x1d, 0xbf, 0x7a, 0xae, 0x07,
		0x34, 0x17, 0x44, 0x94, 0xa9, 0x58, 0x65, 0x88, 0xd9, 0x37, 0x63, 0x30, 0x35, 0xa0, 0x57, 0x51,
		0x20, 0x66, 0xe9, 0x4d, 0xca, 0x18, 0x39, 0x92, 0x54, 0xd9, 0x6f, 0x25, 0x07, 0x63, 0x2d, 0xdd,
		0x38, 0xab, 0xd7, 0x49, 0x6e, 0x84, 0x89, 0x65, 0x53, 0x39, 0x08, 0x50, 0x23, 0x2d, 0x62, 0xd5,
		0x88, 0x65, 0x6c, 0xe7, 0xa2, 0x38, 0x8a, 0xa4, 0xea, 0x93, 0x28, 0x37, 0xc0, 0x64, 0xab, 0xb3,
		0x61, 0x36, 0x0c, 0xcd, 0xa7, 0x06, 0xa8, 0x16, 0x57, 0xb3, 0xbc, 0x63, 0xa9, 0xab, 0x7c, 0x1d,
		0x4c, 0xdc, 0x4f, 0xf4, 0xb3, 0x7e, 0xd5, 0x14, 0x53, 0xcd, 0x50, 0xb1, 0x4f, 0x71, 0x11, 0xd2,
		0x4d, 0xe2, 0x38, 0x38, 0x00, 0xcd, 0xdd, 0x6e, 0x91, 0x5c, 0x8c, 0xcd, 0xfe, 0xca, 0xbe, 0xd9,
		0xf7, 0xce, 0x3c, 0x25, 0x50, 0xeb, 0x08, 0x52, 0xe6, 0x21, 0x49, 0xac, 0x4e, 0x93, 0x33, 0xc4,
		0x77, 0xb0, 0x5f, 0x01, 0x35, 0x7a, 0x59, 0x12, 0x14, 0x26, 0x28, 0xc6, 0x1c, 0xd2, 0x3e, 0xd7,
		0x30, 0x48, 0x6e, 0x94, 0x11, 0x5c, 0xd7, 0x47, 0x50, 0xe1, 0xfd, 0xbd, 0x1c, 0x12, 0x87, 0x53,
		0x49, 0x92, 0xf3, 0x2e, 0xb1, 0x9c, 0x86, 0x6d, 0xe5, 0xc6, 0x18, 0xc9, 0x35, 0x03, 0x56, 0x91,
		0x98, 0xb5, 0x5e, 0x8a, 0x2e, 0x4e, 0xb9, 0x05, 0xc6, 0xec, 0x96, 0x8b, 0xbf, 0x9c, 0x5c, 0x02,
		0xd7, 0x27, 0x75, 0xfc, 0xf2, 0x81, 0x8e, 0x50, 0xe2, 0x3a, 0xaa, 0x54, 0x56, 0x8a, 0x90, 0x75,
		0xec, 0x4e, 0xdb, 0x20, 0x9a, 0x61, 0xd7, 0x88, 0xd6, 0xb0, 0x36, 0xed, 0x5c, 0x92, 0x11, 0x1c,
		0xea, 0x9f, 0x08, 0x53, 0x5c, 0x44, 0xbd, 0x22, 0xaa, 0xa9, 0x19, 0x27, 0xd0, 0x56, 0xf6, 0xc1,
		0xa8, 0xb3, 0x6d, 0xb9, 0xfa, 0xf9, 0x5c, 0x9a, 0x79, 0x88, 0x68, 0xcd, 0xfe, 0x27, 0x0e, 0x13,
		0xc3, 0xb8, 0xd8, 0xed, 0x10, 0xdf, 0xa4, 0xb3, 0x44, 0x07, 0xdb, 0x83, 0x0d, 0x38, 0x26, 0x68,
		0xc4, 0xd1, 0xf7, 0x68, 0xc4, 0x79, 0x48, 0x59, 0xc4, 0x71, 0x49, 0x8d, 0x7b, 0x44, 0x74, 0x48,
		0x9f, 0x02, 0x0e, 0xea, 0x77, 0xa9, 0xd8, 0x7b, 0x72, 0xa9, 0xd3, 0x30, 0xe1, 0x0d, 0x49, 0x6b,
		0xeb, 0x56, 0x5d, 0xfa, 0xe6, 0xb1, 0xb0, 0x91, 0xcc, 0x15, 0x24, 0x4e, 0xa5, 0x30, 0x35, 0x43,
		0x02, 0x6d, 0x65, 0x09, 0xc0, 0xb6, 0x88, 0xbd, 0x89, 0xdb, 0xcb, 0x30, 0xd1, 0x4f, 0x06, 0x5b,
		0xa9, 0x44, 0x55, 0xfa, 0xac, 0x64, 0x73, 0xa9, 0x61, 0x2a, 0xb7, 0x75, 0x5d, 0x6d, 0x6c, 0x07,
		0x4f, 0x59, 0xe5, 0x9b, 0xac, 0xcf, 0xdb, 0xaa, 0x90, 0x69, 0x13, 0xea, 0xf7, 0x68, 0x62, 0x3e,
		0xb3, 0x24, 0x1b, 0xc4, 0x5c, 0xe8, 0xcc, 0x54, 0x01, 0xe3, 0x13, 0x1b, 0x6f, 0xfb, 0x9b, 0xca,
		0x55, 0xe0, 0x09, 0x34, 0xe6, 0x56, 0xc0, 0xa2, 0x50, 0x5a, 0x0a, 0xd7, 0x50, 0x36, 0x73, 0x02,
		0x32, 0x41, 0xf3, 0x28, 0xd3, 0x10, 0x77, 0x5c, 0xbd, 0xed, 0x32, 0x2f, 0x8c, 0xab, 0xbc, 0xa1,
		0x64, 0x21, 0x8a, 0x41, 0x86, 0x45, 0xb9, 0xb8, 0x4a, 0x7f, 0xce, 0xdc, 0x0a, 0xe3, 0x81, 0xd7,
		0x0f, 0x0b, 0x9c, 0x7d, 0x64, 0x14, 0xa6, 0x07, 0xf9, 0xdc, 0x40, 0xf7, 0xc7, 0xed, 0x83, 0x1e,
		0xb0, 0x41, 0xda, 0xe8, 0x77, 0x94, 0x41, 0xb4, 0xd0, 0xa3, 0xe2, 0xa6, 0xbe, 0x41, 0x4c, 0xf4,
		0xa6, 0xc8, 0x91, 0xcc, 0xf1, 0x1b, 0x86, 0xf2, 0xea, 0xb9, 0x15, 0x0a, 0x51, 0x39, 0x52, 0xb9,
		0x03, 0x62, 0x22, 0xc4, 0x51, 0x86, 0xa3, 0xc3, 0x31, 0x50, 0x5f, 0x54, 0x19, 0x4e, 0xb9, 0x0c,
		0x92, 0xf4, 0x2f, 0xb7, 0xed, 0x28, 0x1b, 0x73, 0x82, 0x0a, 0xa8, 0x5d, 0x95, 0x19, 0x48, 0x30,
		0x37, 0xab, 0x11, 0x99, 0x1a, 0xbc, 0x36, 0x5d, 0x98, 0x1a, 0xd9, 0xd4, 0x3b, 0xa6, 0xab, 0x9d,
		0xd3, 0xcd, 0x0e, 0x61, 0x0e, 0x83, 0x0b, 0x23, 0x84, 0x77, 0x53, 0x99, 0x72, 0x08, 0x52, 0xdc,
		0x2b, 0x1b, 0x88, 0x39, 0xcf, 0xa2, 0x4f, 0x5c, 0xe5, 0x8e, 0x5a, 0xa4, 0x12, 0xfa, 0xfa, 0x33,
		0x0e, 0xee, 0x05, 0xb1, 0xb4, 0xec, 0x15, 0x54, 0xc0, 0x5e, 0x7f, 0x6b, 0x6f, 0xe0, 0xbb, 0x62,
		0xf0, 0xf4, 0x7a, 0x7d, 0x71, 0xf6, 0xe7, 0x23, 0x10, 0x63, 0xfb, 0x6d, 0x02, 0x52, 0xeb, 0xf7,
		0x96, 0x0b, 0xda, 0x52, 0xa9, 0xba, 0xb0, 0x52, 0xc8, 0x46, 0x94, 0x0c, 0x00, 0x13, 0x9c, 0x5c,
		0x29, 0xcd, 0xaf, 0x67, 0x47, 0xbc, 0x76, 0x71, 0x6d, 0xfd, 0x96, 0x8f, 0x66, 0xa3, 0x1e, 0xa0,
		0xca, 0x05, 0x31, 0xbf, 0xc2, 0x47, 0x8e, 0x67, 0xe3, 0xe8, 0x09, 0x69, 0x4e, 0x50, 0x3c, 0x5d,
		0x58, 0x42, 0x8d, 0xd1, 0xa0, 0x04, 0x75, 0xc6, 0x94, 0x71, 0x48, 0x32, 0xc9, 0x42, 0xa9, 0xb4,
		0x92, 0x4d, 0x78, 0x9c, 0x95, 0x75, 0xb5, 0xb8, 0xb6, 0x9c, 0x4d, 0x7a, 0x9c, 0xcb, 0x6a, 0xa9,
		0x5a, 0xce, 0x82, 0xc7, 0xb0, 0x5a, 0xa8, 0x54, 0xe6, 0x97, 0x0b, 0xd9, 0x94, 0xa7, 0xb1, 0x70,
		0xef, 0x7a, 0xa1, 0x92, 0x4d, 0x07, 0x86, 0x85, 0xaf, 0x18, 0xf7, 0x5e, 0x51, 0x58, 0xab, 0xae,
		0x66, 0x33, 0xca, 0x24, 0x8c, 0xf3, 0x57, 0xc8, 0x41, 0x4c, 0xf4, 0x88, 0x70, 0xa4, 0xd9, 0xee,
		0x40, 0x38, 0xcb, 0x64, 0x40, 0x80, 0x1a, 0xca, 0xec, 0x22, 0xc4, 0x99, 0x77, 0xa1, 0x17, 0x67,
		0x56, 0xe6, 0x17, 0x0a, 0x2b, 0x5a, 0xa9, 0xbc, 0x5e, 0x2c, 0xad, 0xcd, 0xaf, 0xa0, 0xed, 0x3c,
		0x99, 0x5a, 0xf8, 0x44, 0xb5, 0xa8, 0x16, 0x96, 0xd0, 0x7e, 0x3e, 0x59, 0xb9, 0x30, 0xbf, 0x8e,
		0xb2, 0xe8, 0xec, 0x51, 0x98, 0x1e, 0x14, 0x67, 0x06, 0xed, 0x8c, 0xd9, 0x67, 0x22, 0x30, 0x35,
		0x20, 0x64, 0x0e, 0xdc, 0x45, 0x77, 0x42, 0x9c, 0x7b, 0x1a, 0x4f, 0x22, 0xd7, 0x0f, 0x8c, 0xbd,
		0xcc, 0xef, 0xfa, 0x12, 0x09, 0xc3, 0xf9, 0x13, 0x69, 0x74, 0x87, 0x44, 0x4a, 0x29, 0xfa, 0xdc,
		0xe9, 0xa1, 0x08, 0xe4, 0x76, 0xe2, 0x0e, 0xd9, 0xef, 0x23, 0x81, 0xfd, 0x7e, 0x7b, 0xef, 0x00,
		0x0e, 0xef, 0x3c, 0x87, 0xbe, 0x51, 0x3c, 0x1b, 0x81, 0x7d, 0x83, 0xeb, 0x8d, 0x81, 0x63, 0xb8,
		0x03, 0x46, 0x9b, 0xc4, 0xdd, 0xb2, 0x65, 0xce, 0xbd, 0x76, 0x40, 0x24, 0xa7, 0xdd, 0xbd, 0xb6,
		0x12, 0x28, 0x7f, 0x2a, 0x88, 0xee, 0x54, 0x34, 0xf0, 0xd1, 0xf4, 0x8d, 0xf4, 0xe1, 0x11, 0xb8,
		0x74, 0x20, 0xf9, 0xc0, 0x81, 0x5e, 0x01, 0xd0, 0xb0, 0x5a, 0x1d, 0x97, 0xe7, 0x55, 0x1e, 0x66,
		0x92, 0x4c, 0xc2, 0xb6, 0x30, 0x0d, 0x21, 0x1d, 0xd7, 0xeb, 0x8f, 0xb2, 0x7e, 0xe0, 0x22, 0xa6,
		0x70, 0xa2, 0x3b, 0xd0, 0x18, 0x1b, 0xe8, 0xc1, 0x1d, 0x66, 0xda, 0x97, 0xb2, 0x6e, 0x82, 0xac,
		0x61, 0x36, 0x88, 0xe5, 0x6a, 0x8e, 0xdb, 0x26, 0x7a, 0xb3, 0x61, 0xd5, 0x59, 0x1c, 0x4d, 0xe4,
		0xe3, 0x9b, 0xba, 0xe9, 0x10, 0x75, 0x82, 0x77, 0x57, 0x64, 0x2f, 0x45, 0xb0, 0x64, 0xd1, 0xf6,
		0x21, 0x46, 0x03, 0x08, 0xde, 0xed, 0x21, 0x66, 0xff, 0x30, 0x06, 0x29, 0x5f, 0x75, 0xa6, 0x1c,
		0x86, 0xf4, 0x19, 0xfd, 0x9c, 0xae, 0xc9, 0x8a, 0x9b, 0x5b, 0x22, 0x45, 0x65, 0x65, 0x51, 0x75,
		0xdf, 0x04, 0xd3, 0x4c, 0x05, 0xe7, 0x88, 0x2f, 0x32, 0x4c, 0xdd, 0x71, 0x98, 0xd1, 0x12, 0x4c,
		0x55, 0xa1, 0x7d, 0x25, 0xda, 0xb5, 0x28, 0x7b, 0x94, 0x9b, 0x61, 0x8a, 0x21, 0x9a, 0x18, 0x78,
		0x1b, 0x2d, 0x93, 0x68, 0xf4, 0x0c, 0xe0, 0xb0, 0x78, 0xea, 0x8d, 0x6c, 0x92, 0x6a, 0xac, 0x0a,
		0x05, 0x3a, 0x22, 0x47, 0x59, 0x86, 0x2b, 0x18, 0xac, 0x4e, 0x2c, 0xd2, 0xd6, 0x5d, 0xa2, 0x91,
		0x4f, 0x75, 0x50, 0x57, 0xd3, 0xad, 0x9a, 0xb6, 0xa5, 0x3b, 0x5b, 0xb9, 0x69, 0x3f, 0xc1, 0x01,
		0xaa, 0xbb, 0x2c, 0x54, 0x0b, 0x4c, 0x73, 0xde, 0xaa, 0xdd, 0x85, 0x7a, 0x4a, 0x1e, 0xf6, 0x31,
		0x22, 0x34, 0x0a, 0xce, 0x59, 0x33, 0xb6, 0x88, 0x71, 0x56, 0xeb, 0xb8, 0x9b, 0x27, 0x72, 0x97,
		0xf9, 0x19, 0xd8, 0x20, 0x2b, 0x4c, 0x67, 0x91, 0xaa, 0x54, 0x51, 0x43, 0xa9, 0x40, 0x9a, 0xae,
		0x47, 0xb3, 0xf1, 0x00, 0x0e, 0xdb, 0x6e, 0xb3, 0x1c, 0x91, 0x19, 0xb0, 0xb9, 0x7d, 0x46, 0x9c,
		0x2b, 0x09, 0xc0, 0x2a, 0xd6, 0xa7, 0xf9, 0x78, 0xa5, 0x5c, 0x28, 0x2c, 0xa9, 0x29, 0xc9, 0x72,
		0xd2, 0x6e, 0x53, 0x9f, 0xaa, 0xdb, 0x9e, 0x8d, 0x53, 0xdc, 0xa7, 0xea, 0xb6, 0xb4, 0x30, 0xda,
		0xcb, 0x30, 0xf8, 0xb4, 0xf1, 0xec, 0x22, 0x8a, 0x75, 0x27, 0x97, 0x0d, 0xd8, 0xcb, 0x30, 0x96,
		0xb9, 0x82, 0x70, 0x73, 0x07, 0xb7, 0xc4, 0xa5, 0x5d, 0x7b, 0xf9, 0x81, 0x93, 0x7d, 0xb3, 0xec,
		0x85, 0xe2, 0x1b, 0x5b, 0xdb, 0xfd, 0x40, 0x25, 0xf0, 0xc6, 0xd6, 0x76, 0x2f, 0xec, 0x1a, 0x76,
		0x00, 0x6b, 0x13, 0x03, 0x4d, 0x5e, 0xcb, 0xed, 0xf7, 0x6b, 0xfb, 0x3a, 0x94, 0x63, 0xe8, 0xc8,
		0x86, 0x46, 0x2c, 0x7d, 0x03, 0xd7, 0x5e, 0x6f, 0xe3, 0x0f, 0x27, 0x77, 0xc8, 0xaf, 0x9c, 0x31,
		0x8c, 0x02, 0xeb, 0x9d, 0x67, 0x9d, 0xca, 0x51, 0x98, 0xb4, 0x37, 0xce, 0x18, 0xdc, 0xb9, 0x34,
		0xe4, 0xd9, 0x6c, 0x9c, 0xcf, 0x5d, 0xcd, 0xcc, 0x34, 0x41, 0x3b, 0x98, 0x6b, 0x95, 0x99, 0x58,
		0xb9, 0x1e, 0xc9, 0x9d, 0x2d, 0xbd, 0xdd, 0x62, 0x49, 0xda, 0x41, 0xa3, 0x92, 0xdc, 0x35, 0x5c,
		0x95, 0xcb, 0xd7, 0xa4, 0x58, 0x29, 0xc0, 0x21, 0x3a, 0x79, 0x4b, 0xb7, 0x6c, 0xad, 0xe3, 0x10,
		0xad, 0x3b, 0x44, 0x6f, 0x2d, 0xae, 0xa5, 0xc3, 0x52, 0x2f, 0x97, 0x6a, 0x55, 0x07, 0x83, 0x99,
		0x54, 0x92, 0xcb, 0x73, 0x1a, 0xa6, 0x3b, 0x56, 0xc3, 0x42, 0x17, 0xc7, 0x1e, 0x0a, 0xe6, 0x1b,
		0x36, 0xf7, 0x97, 0xb1, 0x1d, 0x8a, 0xee, 0xaa, 0x5f, 0x9b, 0x3b, 0x89, 0x3a, 0xd5, 0xe9, 0x17,
		0xce, 0xe6, 0x21, 0xed, 0xf7, 0x1d, 0x25, 0x09, 0xdc, 0x7b, 0x30, 0xbb, 0x61, 0x46, 0x5d, 0x2c,
		0x2d, 0xd1, 0x5c, 0x78, 0x5f, 0x01, 0x13, 0x1b, 0xe6, 0xe4, 0x95, 0xe2, 0x7a, 0x41, 0x53, 0xab,
		0x6b, 0xeb, 0xc5, 0xd5, 0x42, 0x36, 0x7a, 0x34, 0x99, 0x78, 0x6b, 0x2c, 0xfb, 0x20, 0xfe, 0x1b,
		0x99, 0x7d, 0x65, 0x04, 0x32, 0xc1, 0x3a, 0x58, 0xf9, 0x18, 0xec, 0x97, 0x87, 0x56, 0x87, 0xb8,
		0xda, 0xfd, 0x8d, 0x36, 0x73, 0xe7, 0xa6, 0xce, 0x2b, 0x49, 0x6f, 0x25, 0xa6, 0x85, 0x16, 0x1e,
		0xef, 0xef, 0x41, 0x9d, 0x93, 0x4c, 0x45, 0x59, 0x81, 0x43, 0x68, 0x32, 0xac, 0x35, 0xad, 0x9a,
		0xde, 0xae, 0x69, 0xdd, 0xeb, 0x02, 0x4d, 0x37, 0xd0, 0x0f, 0x1c, 0x9b, 0x67, 0x12, 0x8f, 0xe5,
		0x72, 0xcb, 0xae, 0x08, 0xe5, 0x6e, 0x88, 0x9d, 0x17, 0xaa, 0x3d, 0x5e, 0x13, 0xdd, 0xc9, 0x6b,
		0xb0, 0xf6, 0x6a, 0xea, 0x2d, 0x74, 0x1b, 0xb7, 0xbd, 0xcd, 0xaa, 0xb7, 0x84, 0x9a, 0x40, 0x41,
		0x81, 0xb6, 0xdf, 0xbf, 0x35, 0xf0, 0xdb, 0xf1, 0x4f, 0x51, 0x48, 0xfb, 0x2b, 0x38, 0x5a, 0x10,
		0x1b, 0x2c, 0xcc, 0x47, 0x58, 0x14, 0xb8, 0x6a, 0xd7, 0x7a, 0x6f, 0x6e, 0x91, 0xc6, 0xff, 0xfc,
		0x28, 0xaf, 0xab, 0x54, 0x8e, 0xa4, 0xb9, 0x97, 0xfa, 0x1a, 0xe1, 0xd5, 0x7a, 0x42, 0x15, 0x2d,
		0x0c, 0x76, 0xa3, 0x67, 0x1c, 0xc6, 0x3d, 0xca, 0xb8, 0xaf, 0xde, 0x9d, 0xfb, 0x54, 0x85, 0x91,
		0x27, 0x4f, 0x55, 0xb4, 0xb5, 0x92, 0xba, 0x3a, 0xbf, 0xa2, 0x0a, 0xb8, 0x72, 0x00, 0x62, 0xa6,
		0xfe, 0xc0, 0x76, 0x30, 0x53, 0x30, 0xd1, 0xb0, 0x86, 0x47, 0x06, 0x7a, 0xe5, 0x11, 0x8c, 0xcf,
		0x4c, 0xf4, 0x3e, 0xba, 0xfe, 0x31, 0x88, 0x33, 0x7b, 0x29, 0x00, 0xc2, 0x62, 0xd9, 0x4b, 0x94,
		0x04, 0xc4, 0x16, 0x4b, 0x2a, 0x75, 0x7f, 0xf4, 0x77, 0x2e, 0xd5, 0xca, 0xc5, 0xc2, 0x22, 0xee,
		0x80, 0xd9, 0x9b, 0x61, 0x94, 0x1b, 0x81, 0x6e, 0x0d, 0xcf, 0x0c, 0x08, 0xe2, 0x4d, 0xc1, 0x11,
		0x91, 0xbd, 0xd5, 0xd5, 0x85, 0x82, 0x9a, 0x1d, 0xf1, 0x2f, 0xef, 0x2f, 0x23, 0x90, 0xf2, 0x15,
		0x54, 0x34, 0x95, 0xeb, 0xa6, 0x69, 0xdf, 0xaf, 0xe9, 0x66, 0x03, 0x23, 0x14, 0x5f, 0x1f, 0x60,
		0xa2, 0x79, 0x2a, 0x19, 0xd6, 0x7e, 0xff, 0x13, 0xdf, 0x7c, 0x2a, 0x02, 0xd9, 0xde, 0x62, 0xac,
		0x67, 0x80, 0x91, 0x0f, 0x74, 0x80, 0x4f, 0x44, 0x20, 0x13, 0xac, 0xc0, 0x7a, 0x86, 0x77, 0xf8,
		0x03, 0x1d, 0xde, 0xe3, 0x11, 0x18, 0x0f, 0xd4, 0x5d, 0xff, 0x57, 0xa3, 0x7b, 0x2c, 0x0a, 0x53,
		0x03, 0x70, 0x18, 0x80, 0x78, 0x81, 0xca, 0x6b, 0xe6, 0x1b, 0x87, 0x79, 0xd7, 0x1c, 0xcd, 0x7f,
		0x65, 0xbd, 0xed, 0x8a, 0x7a, 0x16, 0xf3, 0x65, 0xa3, 0x86, 0x41, 0xb5, 0xb1, 0xd9, 0xc0, 0xf2,
		0x8d, 0x9f, 0x58, 0x78, 0xd5, 0x3a, 0xd1, 0x95, 0xf3, 0xe3, 0xf1, 0x87, 0x40, 0x69, 0xd9, 0x4e,
		0xc3, 0x6d, 0x9c, 0xa3, 0xd7, 0x73, 0xf2, 0x20, 0x4d, 0xab, 0xd8, 0x98, 0x9a, 0x95, 0x3d, 0x45,
		0xcb, 0xf5, 0xb4, 0x2d, 0x52, 0xd7, 0x7b, 0xb4, 0x69, 0x18, 0x8a, 0xaa, 0x59, 0xd9, 0xe3, 0x69,
		0x63, 0xa1, 0x59, 0xb3, 0x3b, 0xb4, 0x20, 0xe0, 0x7a, 0x34, 0xea, 0x45, 0xd4, 0x14, 0x97, 0x79,
		0x2a, 0xa2, 0x62, 0xeb, 0x9e, 0xe0, 0xd3, 0x6a, 0x8a, 0xcb, 0xb8, 0xca, 0x75, 0x30, 0xa1, 0xd7,
		0xeb, 0x6d, 0x4a, 0x2e, 0x89, 0x78, 0x19, 0x9a, 0xf1, 0xc4, 0x4c, 0x71, 0xe6, 0x14, 0x24, 0xa4,
		0x1d, 0x68, 0x62, 0xa1, 0x96, 0xc0, 0x9c, 0xcf, 0xee, 0x51, 0x46, 0xe8, 0xa1, 0xde, 0x92, 0x9d,
		0xf8, 0xd2, 0x86, 0xa3, 0x75, 0x2f, 0xf4, 0x46, 0xb0, 0x3f, 0xa1, 0xa6, 0x1a, 0x8e, 0x77, 0x83,
		0x33, 0xfb, 0x2c, 0xa6, 0xd7, 0xe0, 0x85, 0xa4, 0xb2, 0x04, 0x09, 0xd3, 0x46, 0xff, 0xa0, 0x08,
		0x7e, 0x1b, 0x7e, 0x24, 0xe4, 0x0e, 0x73, 0x6e, 0x45, 0xe8, 0xab, 0x1e, 0x72, 0xe6, 0xb7, 0x11,
		0x48, 0x48, 0x31, 0x26, 0x8a, 0x58, 0x4b, 0x77, 0xb7, 0x18, 0x5d, 0x7c, 0x61, 0x24, 0x1b, 0x51,
		0x59, 0x9b, 0xca, 0xb1, 0x9a, 0xb1, 0x98, 0x0b, 0x08, 0x39, 0x6d, 0xd3, 0x75, 0x35, 0x89, 0x5e,
		0x63, 0x05, 0xae, 0xdd, 0x6c, 0xe2, 0x4a, 0x3a, 0x72, 0x5d, 0x85, 0x7c, 0x51, 0x88, 0xe9, 0xbd,
		0xb8, 0xdb, 0xd6, 0x1b, 0x66, 0x40, 0x37, 0xc6, 0x74, 0xb3, 0xb2, 0xc3, 0x53, 0xce, 0xc3, 0x01,
		0xc9, 0x5b, 0x23, 0xae, 0x8e, 0xc5, 0x73, 0xad, 0x0b, 0x1a, 0x65, 0xb7, 0x5d, 0xfb, 0x85, 0xc2,
		0x92, 0xe8, 0x97, 0xd8, 0x85, 0xd3, 0x58, 0xc8, 0xda, 0xcd, 0x5e, 0x4b, 0x2c, 0x64, 0x7b, 0xce,
		0x5d, 0xce, 0x5d, 0x91, 0xfb, 0xa0, 0x5b, 0x54, 0x3c, 0x33, 0x12, 0x5d, 0x2e, 0x2f, 0x3c, 0x3f,
		0x32, 0xb3, 0xcc, 0x71, 0x65, 0x69, 0x41, 0x95, 0x6c, 0x9a, 0xc4, 0xa0, 0xd6, 0x81, 0xa7, 0xaf,
		0x82, 0x1b, 0xeb, 0x0d, 0x77, 0xab, 0xb3, 0x31, 0x87, 0x6f, 0x38, 0x56, 0xb7, 0xeb, 0x76, 0xf7,
		0x73, 0x06, 0x6d, 0xb1, 0x06, 0xfb, 0x25, 0x3e, 0x69, 0x24, 0x3d, 0xe9, 0x4c, 0xe8, 0xf7, 0x8f,
		0xfc, 0x1a, 0x4c, 0x09, 0x65, 0x8d, 0xdd, 0xa9, 0xf2, 0x12, 0x54, 0xd9, 0xf5, 0x40, 0x9e, 0x7b,
		0xe9, 0x4d, 0x96, 0x12, 0xd4, 0x49, 0x01, 0xa5, 0x7d, 0xbc, 0x48, 0xcd, 0xab, 0x70, 0x69, 0x80,
		0x8f, 0xfb, 0x30, 0x1e, 0xb9, 0x77, 0x67, 0x7c, 0x45, 0x30, 0x4e, 0xf9, 0x18, 0x2b, 0x02, 0x9a,
		0x5f, 0x84, 0xf1, 0xbd, 0x70, 0xfd, 0x5a, 0x70, 0xa5, 0x89, 0x9f, 0x64, 0x19, 0x26, 0x18, 0x89,
		0xd1, 0x71, 0x5c, 0xbb, 0xc9, 0x02, 0xc4, 0xee, 0x34, 0xbf, 0x79, 0x93, 0x3b, 0x55, 0x86, 0xc2,
		0x16, 0x3d, 0x54, 0xfe, 0x6e, 0x98, 0xa6, 0x12, 0xb6, 0x07, 0xfd, 0x6c, 0xe1, 0x57, 0x08, 0xb9,
		0xdf, 0x3f, 0xc4, 0x7d, 0x6f, 0xca, 0x23, 0xf0, 0xf1, 0xfa, 0x56, 0xa2, 0x4e, 0x5c, 0x8c, 0x6d,
		0x78, 0xfe, 0x33, 0x4d, 0x65, 0xd7, 0x6f, 0x0c, 0xb9, 0x47, 0xdf, 0x0e, 0xae, 0xc4, 0x32, 0x47,
		0xce, 0x9b, 0x66, 0xbe, 0x0a, 0xfb, 0x07, 0xac, 0xec, 0x10, 0x9c, 0x8f, 0x09, 0xce, 0xe9, 0xbe,
		0xd5, 0xa5, 0xb4, 0x65, 0x90, 0x72, 0x6f, 0x3d, 0x86, 0xe0, 0x7c, 0x5c, 0x70, 0x2a, 0x02, 0x2b,
		0x97, 0x85, 0x32, 0x9e, 0x82, 0x49, 0x3c, 0xa9, 0x6f, 0xd8, 0x8e, 0x38, 0xf7, 0x0e, 0x41, 0xf7,
		0x84, 0xa0, 0x9b, 0x10, 0x40, 0x76, 0x0a, 0xa6, 0x5c, 0xb7, 0x41, 0x62, 0x13, 0x0f, 0x40, 0x43,
		0x50, 0x3c, 0x29, 0x28, 0xc6, 0xa8, 0x3e, 0x85, 0xce, 0x43, 0xba, 0x6e, 0x8b, 0x30, 0x1c, 0x0e,
		0x7f, 0x4a, 0xc0, 0x53, 0x12, 0x23, 0x28, 0x5a, 0x76, 0xab, 0x63, 0xd2, 0x18, 0x1d, 0x4e, 0xf1,
		0x2d, 0x49, 0x21, 0x31, 0x82, 0x62, 0x0f, 0x66, 0xfd, 0xb6, 0xa4, 0x70, 0x7c, 0xf6, 0xbc, 0x93,
		0xde, 0xf5, 0x9a, 0xdb, 0xb6, 0x35, 0xcc, 0x20, 0x9e, 0x16, 0x0c, 0x20, 0x20, 0x94, 0xe0, 0x76,
		0x48, 0x0e, 0xbb, 0x10, 0xdf, 0x11, 0xf0, 0x04, 0x91, 0x2b, 0x80, 0xfb, 0x4c, 0x06, 0x19, 0xfa,
		0x6d, 0x25, 0x9c, 0xe2, 0xbb, 0x82, 0x22, 0xe3, 0x83, 0x89, 0x69, 0xb8, 0xc4, 0x71, 0xf1, 0xa8,
		0x3e, 0x04, 0xc9, 0xb3, 0x72, 0x1a, 0x02, 0x22, 0x4c, 0xb9, 0x41, 0x2c, 0x63, 0x6b, 0x38, 0x86,
		0xe7, 0xa4, 0x29, 0x25, 0x86, 0x52, 0x60, 0xe4, 0x69, 0xea, 0x6d, 0x3c, 0x5c, 0x9b, 0x43, 0x2d,
		0xc7, 0xf7, 0x04, 0x47, 0xda, 0x03, 0x09, 0x8b, 0x74, 0xac, 0xbd, 0xd0, 0x3c, 0x2f, 0x2d, 0xe2,
		0x83, 0x89, 0xad, 0x87, 0x27, 0x53, 0x5a, 0x49, 0xec, 0x85, 0xed, 0xfb, 0x72, 0xeb, 0x71, 0xec,
		0xaa, 0x9f, 0x11, 0x57, 0xda, 0xc1, 0x23, 0xf8, 0x30, 0x34, 0x3f, 0x90, 0x2b, 0xcd, 0x00, 0x14,
		0x7c, 0x2f, 0x1c, 0x18, 0x18, 0xea, 0x87, 0x20, 0x7b, 0x41, 0x90, 0xed, 0x1b, 0x10, 0xee, 0x45,
		0x48, 0xd8, 0x2b, 0xe5, 0x0f, 0x65, 0x48, 0x20, 0x3d, 0x5c, 0x65, 0x5a, 0xc6, 0x3a, 0xfa, 0xe6,
		0xde, 0xac, 0xf6, 0x23, 0x69, 0x35, 0x8e, 0x0d, 0x58, 0x6d, 0x1d, 0xf6, 0x09, 0xc6, 0xbd, 0xad,
		0xeb, 0x8b, 0x32, 0xb0, 0x72, 0x74, 0x35, 0xb8, 0xba, 0x9f, 0x84, 0x19, 0xcf, 0x9c, 0xb2, 0x02,
		0x73, 0x34, 0x7a, 0x31, 0x10, 0xce, 0xfc, 0x92, 0x60, 0x96, 0x11, 0xdf, 0x2b, 0xe1, 0x9c, 0x55,
		0xbd, 0x45, 0xc9, 0x4f, 0x43, 0x4e, 0x92, 0x77, 0x2c, 0x2c, 0xf0, 0xed, 0xba, 0x85, 0xcb, 0x58,
		0x1b, 0x82, 0xfa, 0xc7, 0x3d, 0x4b, 0x55, 0xf5, 0xc1, 0x29, 0x73, 0x11, 0xb2, 0x5e, 0xbd, 0xa1,
		0x35, 0x9a, 0x2d, 0x1b, 0x4b, 0xcb, 0xdd, 0x19, 0x7f, 0x22, 0x57, 0xca, 0xc3, 0x15, 0x19, 0x2c,
		0x5f, 0x80, 0x0c, 0x6b, 0x0e, 0xeb, 0x92, 0x3f, 0x15, 0x44, 0xe3, 0x5d, 0x94, 0x08, 0x1c, 0x58,
		0x29, 0x61, 0xcd, 0x3b, 0x4c, 0xfc, 0xfb, 0x99, 0x0c, 0x1c, 0x02, 0xc2, 0xbd, 0x6f, 0xa2, 0x27,
		0x13, 0x2b, 0x61, 0x9f, 0x5f, 0x73, 0x9f, 0xbe, 0x20, 0xf6, 0x6c, 0x30, 0x11, 0xe7, 0x57, 0xa8,
		0x79, 0x82, 0xe9, 0x32, 0x9c, 0xec, 0xa1, 0x0b, 0x9e, 0x85, 0x02, 0xd9, 0x32, 0x7f, 0x12, 0xc6,
		0x03, 0xa9, 0x32, 0x9c, 0xea, 0x33, 0x82, 0x2a, 0xed, 0xcf, 0x94, 0xf9, 0x9b, 0x21, 0x46, 0xd3,
		0x5e, 0x38, 0xfc, 0xb3, 0x02, 0xce, 0xd4, 0xf3, 0x1f, 0x87, 0x84, 0x4c, 0x77, 0xe1, 0xd0, 0xcf,
		0x09, 0xa8, 0x07, 0xa1, 0x70, 0x99, 0xea, 0xc2, 0xe1, 0x9f, 0x97, 0x70, 0x09, 0xa1, 0xf0, 0xe1,
		0x4d, 0xf8, 0xf2, 0x17, 0x62, 0x22, 0x5c, 0x49, 0xdb, 0xd1, 0x6f, 0x3e, 0x3c, 0xc7, 0x85, 0xa3,
		0x1f, 0x16, 0x2f, 0x97, 0x88, 0xfc, 0xad, 0x10, 0x1f, 0xd2, 0xe0, 0x5f, 0x14, 0x50, 0xae, 0x8f,
		0x19, 0x24, 0xe5, 0xcb, 0x6b, 0xe1, 0xf0, 0x2f, 0x09, 0xb8, 0x1f, 0x45, 0x87, 0x2e, 0xf2, 0x5a,
		0x38, 0xc1, 0x97, 0xe5, 0xd0, 0x05, 0x82, 0x9a, 0x4d, 0xa6, 0xb4, 0x70, 0xf4, 0x57, 0xa4, 0xd5,
		0x25, 0x04, 0x77, 0x53, 0xd2, 0x0b, 0x53, 0xe1, 0xf8, 0xaf, 0x0a, 0x7c, 0x17, 0x43, 0x2d, 0xe0,
		0x0b, 0x93, 0xe1, 0x14, 0x5f, 0x93, 0x16, 0xf0, 0xa1, 0xe8, 0x36, 0xea, 0x4d, 0x7d, 0xe1, 0x4c,
		0x5f, 0x97, 0xdb, 0xa8, 0x27, 0xf3, 0xd1, 0xd5, 0x64, 0xd1, 0x22, 0x9c, 0xe2, 0x1b, 0x72, 0x35,
		0x99, 0x3e, 0x1d, 0x46, 0x6f, 0x2e, 0x09, 0xe7, 0xf8, 0xa6, 0x1c, 0x46, 0x4f, 0x2a, 0xc1, 0xcc,
		0xa4, 0xf4, 0xe7, 0x91, 0x70, 0xbe, 0x47, 0x04, 0xdf, 0x64, 0x5f, 0x1a, 0xc9, 0xdf, 0x03, 0xfb,
		0x06, 0xe7, 0x90, 0x70, 0xd6, 0x47, 0x2f, 0xf4, 0x54, 0xfd, 0xfe, 0x14, 0x82, 0x29, 0x6f, 0x7a,
		0x50, 0xfe, 0x08, 0xa7, 0x7d, 0xec, 0x42, 0xf0, 0x60, 0xe7, 0x4f, 0x1f, 0x58, 0xa1, 0x41, 0x37,
		0x74, 0x87, 0x73, 0x3d, 0x21, 0xb8, 0x7c, 0x20, 0xba, 0x35, 0x44, 0xe4, 0x0e, 0xc7, 0x3f, 0x29,
		0xb7, 0x86, 0x40, 0x20, 0x38, 0x61, 0x75, 0x4c, 0x93, 0x3a, 0x87, 0xb2, 0xfb, 0x7f, 0x69, 0xc8,
		0xfd, 0xf5, 0xa2, 0xd8, 0x18, 0x12, 0x80, 0x31, 0x34, 0x4e, 0x9a, 0x1b, 0x68, 0x83, 0x10, 0xe4,
		0xdf, 0x2e, 0xca, 0x80, 0x40, 0xb5, 0x71, 0x3f, 0x01, 0x3f, 0x34, 0xb2, 0x3b, 0xec, 0x10, 0xec,
		0xdf, 0x2f, 0x8a, 0xcf, 0xac, 0x5d, 0x48, 0x97, 0x80, 0x7f, 0xb4, 0xdd, 0x9d, 0xe0, 0xed, 0x20,
		0x01, 0x3b, 0x68, 0xde, 0x06, 0x63, 0xf4, 0x7f, 0x76, 0xb8, 0x7a, 0x3d, 0x0c, 0xfd, 0x0f, 0x81,
		0x96, 0xfa, 0xd4, 0x60, 0x4d, 0xbb, 0x4d, 0xf0, 0xa7, 0x13, 0x86, 0xfd, 0xa7, 0xc0, 0x7a, 0x00,
		0x0a, 0x36, 0x74, 0xc7, 0x1d, 0x66, 0xde, 0xff, 0x92, 0x60, 0x09, 0xa0, 0x83, 0xa6, 0xbf, 0xcf,
		0x92, 0xed, 0x30, 0xec, 0x3b, 0x72, 0xd0, 0x42, 0x1f, 0x03, 0x60, 0x92, 0xfe, 0xe4, 0xff, 0xf5,
		0x20, 0x04, 0xfc, 0x6f, 0x01, 0xee, 0x22, 0x16, 0x0e, 0x0f, 0xbe, 0xda, 0x81, 0x65, 0x7b, 0xd9,
		0xe6, 0x97, 0x3a, 0xf0, 0x42, 0x04, 0x32, 0xf4, 0x4b, 0xef, 0x5c, 0xcd, 0x76, 0xc5, 0x25, 0x4c,
		0x8a, 0xb6, 0xb1, 0x49, 0x2d, 0x3e, 0xb3, 0xb7, 0x0b, 0x9c, 0xd9, 0x49, 0x88, 0xac, 0x2a, 0x69,
		0x88, 0xe8, 0xe2, 0xa3, 0x74, 0x44, 0x5f, 0x58, 0x79, 0xf5, 0xf5, 0x83, 0x97, 0xfc, 0x0e, 0x9f,
		0x3f, 0xe2, 0xf3, 0xda, 0xeb, 0x07, 0x23, 0x6f, 0xe1, 0xf3, 0x0e, 0x3e, 0xef, 0xe2, 0xf3, 0xe0,
		0x1b, 0x07, 0x23, 0xcf, 0xe1, 0xf3, 0x22, 0x3e, 0xbf, 0xc0, 0xe7, 0x65, 0x7c, 0x5e, 0x7d, 0x03,
		0xf5, 0xf1, 0x79, 0x0d, 0x7f, 0xbf, 0x85, 0x7f, 0xdf, 0xc1, 0xbf, 0xef, 0xe2, 0xdf, 0x07, 0xff,
		0x7c, 0xf0, 0x92, 0xff, 0x06, 0x00, 0x00, 0xff, 0xff, 0xe7, 0x43, 0x74, 0x16, 0x1a, 0x2b, 0x00,
		0x00,
	}
	r := bytes.NewReader(gzipped)
	gzipr, err := compress_gzip.NewReader(r)
	if err != nil {
		panic(err)
	}
	ungzipped, err := io_ioutil.ReadAll(gzipr)
	if err != nil {
		panic(err)
	}
	if err := github_com_gogo_protobuf_proto.Unmarshal(ungzipped, d); err != nil {
		panic(err)
	}
	return d
}
func (this *M) VerboseEqual(that interface{}) error {
	if that == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that == nil && this != nil")
	}

	that1, ok := that.(*M)
	if !ok {
		that2, ok := that.(M)
		if ok {
			that1 = &that2
		} else {
			return fmt.Errorf("that is not of type *M")
		}
	}
	if that1 == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that is type *M but is nil && this != nil")
	} else if this == nil {
		return fmt.Errorf("that is type *M but is not nil && this == nil")
	}
	if this.A != nil && that1.A != nil {
		if *this.A != *that1.A {
			return fmt.Errorf("A this(%v) Not Equal that(%v)", *this.A, *that1.A)
		}
	} else if this.A != nil {
		return fmt.Errorf("this.A == nil && that.A != nil")
	} else if that1.A != nil {
		return fmt.Errorf("A this(%v) Not Equal that(%v)", this.A, that1.A)
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return fmt.Errorf("XXX_unrecognized this(%v) Not Equal that(%v)", this.XXX_unrecognized, that1.XXX_unrecognized)
	}
	return nil
}
func (this *M) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*M)
	if !ok {
		that2, ok := that.(M)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	if this.A != nil && that1.A != nil {
		if *this.A != *that1.A {
			return false
		}
	} else if this.A != nil {
		return false
	} else if that1.A != nil {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}

type MFace interface {
	Proto() github_com_gogo_protobuf_proto.Message
	GetA() *string
}

func (this *M) Proto() github_com_gogo_protobuf_proto.Message {
	return this
}

func (this *M) TestProto() github_com_gogo_protobuf_proto.Message {
	return NewMFromFace(this)
}

func (this *M) GetA() *string {
	return this.A
}

func NewMFromFace(that MFace) *M {
	this := &M{}
	this.A = that.GetA()
	return this
}

func (this *M) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 5)
	s = append(s, "&filedotname.M{")
	if this.A != nil {
		s = append(s, "A: "+valueToGoStringFileDot(this.A, "string")+",\n")
	}
	if this.XXX_unrecognized != nil {
		s = append(s, "XXX_unrecognized:"+fmt.Sprintf("%#v", this.XXX_unrecognized)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringFileDot(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func extensionToGoStringFileDot(e map[int32]github_com_gogo_protobuf_proto.Extension) string {
	if e == nil {
		return "nil"
	}
	s := "map[int32]proto.Extension{"
	keys := make([]int, 0, len(e))
	for k := range e {
		keys = append(keys, int(k))
	}
	sort.Ints(keys)
	ss := []string{}
	for _, k := range keys {
		ss = append(ss, strconv.Itoa(k)+": "+e[int32(k)].GoString())
	}
	s += strings.Join(ss, ",") + "}"
	return s
}
func NewPopulatedM(r randyFileDot, easy bool) *M {
	this := &M{}
	if r.Intn(10) != 0 {
		v1 := randStringFileDot(r)
		this.A = &v1
	}
	if !easy && r.Intn(10) != 0 {
		this.XXX_unrecognized = randUnrecognizedFileDot(r, 2)
	}
	return this
}

type randyFileDot interface {
	Float32() float32
	Float64() float64
	Int63() int64
	Int31() int32
	Uint32() uint32
	Intn(n int) int
}

func randUTF8RuneFileDot(r randyFileDot) rune {
	ru := r.Intn(62)
	if ru < 10 {
		return rune(ru + 48)
	} else if ru < 36 {
		return rune(ru + 55)
	}
	return rune(ru + 61)
}
func randStringFileDot(r randyFileDot) string {
	v2 := r.Intn(100)
	tmps := make([]rune, v2)
	for i := 0; i < v2; i++ {
		tmps[i] = randUTF8RuneFileDot(r)
	}
	return string(tmps)
}
func randUnrecognizedFileDot(r randyFileDot, maxFieldNumber int) (data []byte) {
	l := r.Intn(5)
	for i := 0; i < l; i++ {
		wire := r.Intn(4)
		if wire == 3 {
			wire = 5
		}
		fieldNumber := maxFieldNumber + r.Intn(100)
		data = randFieldFileDot(data, r, fieldNumber, wire)
	}
	return data
}
func randFieldFileDot(data []byte, r randyFileDot, fieldNumber int, wire int) []byte {
	key := uint32(fieldNumber)<<3 | uint32(wire)
	switch wire {
	case 0:
		data = encodeVarintPopulateFileDot(data, uint64(key))
		v3 := r.Int63()
		if r.Intn(2) == 0 {
			v3 *= -1
		}
		data = encodeVarintPopulateFileDot(data, uint64(v3))
	case 1:
		data = encodeVarintPopulateFileDot(data, uint64(key))
		data = append(data, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	case 2:
		data = encodeVarintPopulateFileDot(data, uint64(key))
		ll := r.Intn(100)
		data = encodeVarintPopulateFileDot(data, uint64(ll))
		for j := 0; j < ll; j++ {
			data = append(data, byte(r.Intn(256)))
		}
	default:
		data = encodeVarintPopulateFileDot(data, uint64(key))
		data = append(data, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	}
	return data
}
func encodeVarintPopulateFileDot(data []byte, v uint64) []byte {
	for v >= 1<<7 {
		data = append(data, uint8(uint64(v)&0x7f|0x80))
		v >>= 7
	}
	data = append(data, uint8(v))
	return data
}
func (m *M) Size() (n int) {
	var l int
	_ = l
	if m.A != nil {
		l = len(*m.A)
		n += 1 + l + sovFileDot(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovFileDot(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozFileDot(x uint64) (n int) {
	return sovFileDot(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *M) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&M{`,
		`A:` + valueToStringFileDot(this.A) + `,`,
		`XXX_unrecognized:` + fmt.Sprintf("%v", this.XXX_unrecognized) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringFileDot(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}

func init() { proto.RegisterFile("file.dot.proto", fileDescriptorFileDot) }

var fileDescriptorFileDot = []byte{
	// 158 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0x4b, 0xcb, 0xcc, 0x49,
	0xd5, 0x4b, 0xc9, 0x2f, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x06, 0xf1, 0x81, 0xdc,
	0xbc, 0xc4, 0xdc, 0x54, 0x29, 0xdd, 0xf4, 0xcc, 0x92, 0x8c, 0xd2, 0x24, 0xbd, 0xe4, 0xfc, 0x5c,
	0xfd, 0xf4, 0xfc, 0xf4, 0x7c, 0x7d, 0xb0, 0x9a, 0xa4, 0xd2, 0x34, 0x30, 0x0f, 0xcc, 0x01, 0xb3,
	0x20, 0x7a, 0x95, 0x04, 0xb9, 0x18, 0x7d, 0x85, 0x78, 0xb8, 0x18, 0x13, 0x25, 0x18, 0x15, 0x18,
	0x35, 0x38, 0x83, 0x18, 0x13, 0x9d, 0x7c, 0x4e, 0x3c, 0x94, 0x63, 0xb8, 0x00, 0xc4, 0x37, 0x80,
	0xf8, 0xc1, 0x43, 0x39, 0xc6, 0x17, 0x40, 0xfc, 0x01, 0x88, 0x7f, 0x00, 0x71, 0xc3, 0x23, 0x39,
	0xc6, 0x15, 0x40, 0xbc, 0x01, 0x88, 0x77, 0x00, 0xf1, 0x01, 0x20, 0x3e, 0xf1, 0x08, 0xa8, 0x1e,
	0x88, 0x1f, 0x00, 0xd9, 0x2f, 0x80, 0xf4, 0x07, 0x20, 0xfd, 0x03, 0x48, 0x37, 0x3c, 0x96, 0x63,
	0x00, 0x04, 0x00, 0x00, 0xff, 0xff, 0x3f, 0x59, 0x32, 0x8a, 0xad, 0x00, 0x00, 0x00,
}
