// Code generated by protoc-gen-go. DO NOT EDIT.
// source: test_rpc_service.proto

package pb

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

func init() { proto.RegisterFile("test_rpc_service.proto", fileDescriptor_50b324180f7b6965) }

var fileDescriptor_50b324180f7b6965 = []byte{
	// 234 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x90, 0x41, 0x4a, 0x03, 0x41,
	0x10, 0x45, 0x99, 0x30, 0xd9, 0xd4, 0xce, 0x56, 0x5c, 0xe4, 0x0e, 0x96, 0x41, 0xf0, 0x00, 0x06,
	0xb2, 0x0f, 0xa3, 0x20, 0xb8, 0x09, 0x3d, 0xdd, 0x95, 0xe9, 0x59, 0x38, 0x55, 0x56, 0xf5, 0x08,
	0xde, 0x40, 0x3c, 0x85, 0x47, 0x95, 0x99, 0x36, 0x60, 0x40, 0xd0, 0x55, 0xff, 0xfe, 0xbc, 0x5f,
	0x1f, 0x3e, 0x5c, 0x66, 0xb2, 0xbc, 0x57, 0x09, 0x7b, 0x23, 0x7d, 0xed, 0x03, 0xa1, 0x28, 0x67,
	0x5e, 0xc1, 0xe4, 0x17, 0x7d, 0xf3, 0xb1, 0x80, 0x8b, 0x07, 0xb2, 0xbc, 0x9b, 0x7e, 0xed, 0x78,
	0x68, 0x24, 0xcc, 0xd2, 0x5d, 0x43, 0x2d, 0xfd, 0xd0, 0x39, 0x87, 0xdb, 0x67, 0xc9, 0x6f, 0x0d,
	0xbd, 0x8c, 0x47, 0x6c, 0x75, 0x7e, 0xf4, 0x4c, 0x78, 0x30, 0x2a, 0x81, 0x2b, 0xa8, 0x29, 0x24,
	0x76, 0x67, 0xb8, 0x0d, 0x89, 0x4f, 0x78, 0xf7, 0x6d, 0xfd, 0xc4, 0xd7, 0xb0, 0x24, 0x55, 0xd6,
	0xff, 0x17, 0xac, 0x61, 0x29, 0x7e, 0x34, 0x72, 0x0e, 0x77, 0xd3, 0xfb, 0x77, 0x02, 0xa1, 0xf6,
	0x31, 0xfe, 0x5e, 0xe1, 0xf0, 0x2e, 0x46, 0x3d, 0xe1, 0x37, 0x8f, 0x70, 0xcb, 0xda, 0xa1, 0x17,
	0x1f, 0x12, 0x61, 0xf2, 0x91, 0x59, 0x30, 0xb5, 0xde, 0x08, 0x2d, 0xf9, 0x48, 0x11, 0x7b, 0x09,
	0x65, 0xb8, 0x76, 0x3c, 0x60, 0x47, 0x03, 0xa9, 0xcf, 0x14, 0x37, 0xf3, 0x84, 0x8d, 0x84, 0xfb,
	0xb2, 0xf3, 0x7c, 0xcd, 0x9e, 0x16, 0xd2, 0xbe, 0x57, 0xd5, 0x67, 0x55, 0x7d, 0x05, 0x00, 0x00,
	0xff, 0xff, 0xfb, 0xc5, 0x09, 0xf7, 0x8a, 0x01, 0x00, 0x00,
}
