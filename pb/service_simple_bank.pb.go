// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.25.1
// source: service_simple_bank.proto

package pb

import (
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_service_simple_bank_proto protoreflect.FileDescriptor

var file_service_simple_bank_proto_rawDesc = []byte{
	0x0a, 0x19, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x73, 0x69, 0x6d, 0x70, 0x6c, 0x65,
	0x5f, 0x62, 0x61, 0x6e, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x1a,
	0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x72,
	0x70, 0x63, 0x5f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x72, 0x70, 0x63, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x5f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x14, 0x72, 0x70, 0x63,
	0x5f, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x16, 0x72, 0x70, 0x63, 0x5f, 0x76, 0x65, 0x72, 0x69, 0x66, 0x79, 0x5f, 0x65, 0x6d,
	0x61, 0x69, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x72, 0x70, 0x63, 0x5f, 0x72,
	0x65, 0x6e, 0x65, 0x77, 0x5f, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x18, 0x72, 0x70, 0x63, 0x5f, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x15, 0x72, 0x70, 0x63, 0x5f, 0x67, 0x65, 0x74, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x72, 0x70, 0x63, 0x5f, 0x6c, 0x69,
	0x73, 0x74, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x18, 0x72, 0x70, 0x63, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x61, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x18, 0x72, 0x70, 0x63,
	0x5f, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65,
	0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xf7, 0x0c, 0x0a, 0x0a, 0x53, 0x69, 0x6d, 0x70, 0x6c, 0x65,
	0x42, 0x61, 0x6e, 0x6b, 0x12, 0x8e, 0x01, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55,
	0x73, 0x65, 0x72, 0x12, 0x15, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55,
	0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x70, 0x62, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x51, 0x92, 0x41, 0x34, 0x12, 0x0f, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x20,
	0x6e, 0x65, 0x77, 0x20, 0x75, 0x73, 0x65, 0x72, 0x1a, 0x21, 0x55, 0x73, 0x65, 0x20, 0x74, 0x68,
	0x69, 0x73, 0x20, 0x41, 0x50, 0x49, 0x20, 0x74, 0x6f, 0x20, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x20, 0x61, 0x20, 0x6e, 0x65, 0x77, 0x20, 0x75, 0x73, 0x65, 0x72, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x14, 0x3a, 0x01, 0x2a, 0x22, 0x0f, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x5f, 0x75, 0x73, 0x65, 0x72, 0x12, 0x9c, 0x01, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x55, 0x73, 0x65, 0x72, 0x12, 0x15, 0x2e, 0x70, 0x62, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x70, 0x62,
	0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x5f, 0x92, 0x41, 0x42, 0x12, 0x17, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x20, 0x61, 0x6e, 0x20, 0x65, 0x78, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x20, 0x75, 0x73, 0x65,
	0x72, 0x1a, 0x27, 0x55, 0x73, 0x65, 0x20, 0x74, 0x68, 0x69, 0x73, 0x20, 0x41, 0x50, 0x49, 0x20,
	0x74, 0x6f, 0x20, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x20, 0x61, 0x6e, 0x20, 0x65, 0x78, 0x69,
	0x73, 0x74, 0x69, 0x6e, 0x67, 0x20, 0x75, 0x73, 0x65, 0x72, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x14,
	0x3a, 0x01, 0x2a, 0x32, 0x0f, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f,
	0x75, 0x73, 0x65, 0x72, 0x12, 0xa3, 0x01, 0x0a, 0x09, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x55, 0x73,
	0x65, 0x72, 0x12, 0x14, 0x2e, 0x70, 0x62, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x55, 0x73, 0x65,
	0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x70, 0x62, 0x2e, 0x4c, 0x6f,
	0x67, 0x69, 0x6e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x69, 0x92, 0x41, 0x4d, 0x12, 0x0a, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x20, 0x75, 0x73, 0x65, 0x72,
	0x1a, 0x3f, 0x55, 0x73, 0x65, 0x20, 0x74, 0x68, 0x69, 0x73, 0x20, 0x41, 0x50, 0x49, 0x20, 0x74,
	0x6f, 0x20, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x20, 0x75, 0x73, 0x65, 0x72, 0x20, 0x61, 0x6e, 0x64,
	0x20, 0x67, 0x65, 0x74, 0x20, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x20, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x20, 0x26, 0x20, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x20, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x13, 0x3a, 0x01, 0x2a, 0x22, 0x0e, 0x2f, 0x76, 0x31, 0x2f,
	0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x12, 0x96, 0x01, 0x0a, 0x0b, 0x56,
	0x65, 0x72, 0x69, 0x66, 0x79, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x16, 0x2e, 0x70, 0x62, 0x2e,
	0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x17, 0x2e, 0x70, 0x62, 0x2e, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x45, 0x6d,
	0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x56, 0x92, 0x41, 0x3b,
	0x12, 0x0c, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x20, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x1a, 0x2b,
	0x55, 0x73, 0x65, 0x20, 0x74, 0x68, 0x69, 0x73, 0x20, 0x41, 0x50, 0x49, 0x20, 0x74, 0x6f, 0x20,
	0x76, 0x65, 0x72, 0x69, 0x66, 0x79, 0x20, 0x75, 0x73, 0x65, 0x72, 0x27, 0x73, 0x20, 0x65, 0x6d,
	0x61, 0x69, 0x6c, 0x20, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x12, 0x12, 0x10, 0x2f, 0x76, 0x31, 0x2f, 0x76, 0x65, 0x72, 0x69, 0x66, 0x79, 0x5f, 0x65, 0x6d,
	0x61, 0x69, 0x6c, 0x12, 0xb5, 0x01, 0x0a, 0x10, 0x52, 0x65, 0x6e, 0x65, 0x77, 0x41, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1b, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x65,
	0x6e, 0x65, 0x77, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x6e, 0x65, 0x77,
	0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x66, 0x92, 0x41, 0x41, 0x12, 0x12, 0x52, 0x65, 0x6e, 0x65, 0x77, 0x20,
	0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x20, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x1a, 0x2b, 0x55, 0x73,
	0x65, 0x20, 0x74, 0x68, 0x69, 0x73, 0x20, 0x41, 0x50, 0x49, 0x20, 0x74, 0x6f, 0x20, 0x72, 0x65,
	0x6e, 0x65, 0x77, 0x20, 0x61, 0x20, 0x75, 0x73, 0x65, 0x72, 0x27, 0x73, 0x20, 0x61, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x20, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1c, 0x3a,
	0x01, 0x2a, 0x22, 0x17, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x2f, 0x72,
	0x65, 0x6e, 0x65, 0x77, 0x5f, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0xb8, 0x01, 0x0a, 0x0d,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x18, 0x2e,
	0x70, 0x62, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x72, 0x92, 0x41, 0x58, 0x12, 0x15, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x20,
	0x61, 0x20, 0x62, 0x61, 0x6e, 0x6b, 0x20, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x1a, 0x3f,
	0x55, 0x73, 0x65, 0x20, 0x74, 0x68, 0x69, 0x73, 0x20, 0x41, 0x50, 0x49, 0x20, 0x74, 0x6f, 0x20,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x20, 0x61, 0x20, 0x62, 0x61, 0x6e, 0x6b, 0x20, 0x61, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x20, 0x66, 0x6f, 0x72, 0x20, 0x61, 0x6e, 0x20, 0x61, 0x75, 0x74,
	0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x65, 0x64, 0x20, 0x75, 0x73, 0x65, 0x72, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x11, 0x3a, 0x01, 0x2a, 0x22, 0x0c, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x12, 0x92, 0x01, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x41, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x15, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x70,
	0x62, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x55, 0x92, 0x41, 0x39, 0x12, 0x11, 0x47, 0x65, 0x74, 0x20, 0x61,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x20, 0x62, 0x79, 0x20, 0x49, 0x44, 0x1a, 0x24, 0x55, 0x73,
	0x65, 0x20, 0x74, 0x68, 0x69, 0x73, 0x20, 0x41, 0x50, 0x49, 0x20, 0x74, 0x6f, 0x20, 0x67, 0x65,
	0x74, 0x20, 0x61, 0x6e, 0x20, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x20, 0x62, 0x79, 0x20,
	0x69, 0x64, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x13, 0x12, 0x11, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0xae, 0x01, 0x0a, 0x0c,
	0x4c, 0x69, 0x73, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x12, 0x17, 0x2e, 0x70,
	0x62, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x70, 0x62, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x41,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x6b, 0x92, 0x41, 0x54, 0x12, 0x1b, 0x4c, 0x69, 0x73, 0x74, 0x20, 0x61, 0x6c, 0x6c, 0x20, 0x6f,
	0x66, 0x20, 0x75, 0x73, 0x65, 0x72, 0x27, 0x73, 0x20, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x73, 0x1a, 0x35, 0x55, 0x73, 0x65, 0x20, 0x74, 0x68, 0x69, 0x73, 0x20, 0x41, 0x50, 0x49, 0x20,
	0x74, 0x6f, 0x20, 0x6c, 0x69, 0x73, 0x74, 0x20, 0x61, 0x6c, 0x6c, 0x20, 0x61, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x73, 0x20, 0x62, 0x65, 0x6c, 0x6f, 0x6e, 0x67, 0x69, 0x6e, 0x67, 0x20, 0x74,
	0x6f, 0x20, 0x61, 0x20, 0x75, 0x73, 0x65, 0x72, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0e, 0x12, 0x0c,
	0x2f, 0x76, 0x31, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x12, 0xa5, 0x01, 0x0a,
	0x0d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x18,
	0x2e, 0x70, 0x62, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x70, 0x62, 0x2e, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x5f, 0x92, 0x41, 0x40, 0x12, 0x16, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x20, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x20, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65,
	0x1a, 0x26, 0x55, 0x73, 0x65, 0x20, 0x74, 0x68, 0x69, 0x73, 0x20, 0x41, 0x50, 0x49, 0x20, 0x74,
	0x6f, 0x20, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x20, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x20, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x16, 0x3a, 0x01,
	0x2a, 0x32, 0x11, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x2f,
	0x7b, 0x69, 0x64, 0x7d, 0x12, 0x98, 0x01, 0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x18, 0x2e, 0x70, 0x62, 0x2e, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x19, 0x2e, 0x70, 0x62, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x52, 0x92, 0x41, 0x36,
	0x12, 0x11, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x20, 0x61, 0x6e, 0x20, 0x61, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x1a, 0x21, 0x55, 0x73, 0x65, 0x20, 0x74, 0x68, 0x69, 0x73, 0x20, 0x41, 0x50,
	0x49, 0x20, 0x74, 0x6f, 0x20, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x20, 0x61, 0x6e, 0x20, 0x61,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x13, 0x2a, 0x11, 0x2f, 0x76,
	0x31, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x42,
	0x7a, 0x92, 0x41, 0x52, 0x12, 0x50, 0x0a, 0x0f, 0x53, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x20, 0x42,
	0x61, 0x6e, 0x6b, 0x20, 0x41, 0x50, 0x49, 0x22, 0x38, 0x0a, 0x05, 0x62, 0x72, 0x75, 0x63, 0x65,
	0x12, 0x1c, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x62, 0x72, 0x75, 0x63, 0x65, 0x2d, 0x6d, 0x69, 0x67, 0x1a, 0x11,
	0x62, 0x6d, 0x69, 0x67, 0x65, 0x72, 0x69, 0x40, 0x67, 0x6d, 0x61, 0x69, 0x6c, 0x2e, 0x63, 0x6f,
	0x6d, 0x32, 0x03, 0x31, 0x2e, 0x31, 0x5a, 0x23, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x62, 0x72, 0x75, 0x63, 0x65, 0x2d, 0x6d, 0x69, 0x67, 0x2f, 0x73, 0x69, 0x6d,
	0x70, 0x6c, 0x65, 0x2d, 0x62, 0x61, 0x6e, 0x6b, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var file_service_simple_bank_proto_goTypes = []interface{}{
	(*CreateUserRequest)(nil),        // 0: pb.CreateUserRequest
	(*UpdateUserRequest)(nil),        // 1: pb.UpdateUserRequest
	(*LoginUserRequest)(nil),         // 2: pb.LoginUserRequest
	(*VerifyEmailRequest)(nil),       // 3: pb.VerifyEmailRequest
	(*RenewAccessTokenRequest)(nil),  // 4: pb.RenewAccessTokenRequest
	(*CreateAccountRequest)(nil),     // 5: pb.CreateAccountRequest
	(*GetAccountRequest)(nil),        // 6: pb.GetAccountRequest
	(*ListAccountsRequest)(nil),      // 7: pb.ListAccountsRequest
	(*UpdateAccountRequest)(nil),     // 8: pb.UpdateAccountRequest
	(*DeleteAccountRequest)(nil),     // 9: pb.DeleteAccountRequest
	(*CreateUserResponse)(nil),       // 10: pb.CreateUserResponse
	(*UpdateUserResponse)(nil),       // 11: pb.UpdateUserResponse
	(*LoginUserResponse)(nil),        // 12: pb.LoginUserResponse
	(*VerifyEmailResponse)(nil),      // 13: pb.VerifyEmailResponse
	(*RenewAccessTokenResponse)(nil), // 14: pb.RenewAccessTokenResponse
	(*CreateAccountResponse)(nil),    // 15: pb.CreateAccountResponse
	(*GetAccountResponse)(nil),       // 16: pb.GetAccountResponse
	(*ListAccountsResponse)(nil),     // 17: pb.ListAccountsResponse
	(*UpdateAccountResponse)(nil),    // 18: pb.UpdateAccountResponse
	(*DeleteAccountResponse)(nil),    // 19: pb.DeleteAccountResponse
}
var file_service_simple_bank_proto_depIdxs = []int32{
	0,  // 0: pb.SimpleBank.CreateUser:input_type -> pb.CreateUserRequest
	1,  // 1: pb.SimpleBank.UpdateUser:input_type -> pb.UpdateUserRequest
	2,  // 2: pb.SimpleBank.LoginUser:input_type -> pb.LoginUserRequest
	3,  // 3: pb.SimpleBank.VerifyEmail:input_type -> pb.VerifyEmailRequest
	4,  // 4: pb.SimpleBank.RenewAccessToken:input_type -> pb.RenewAccessTokenRequest
	5,  // 5: pb.SimpleBank.CreateAccount:input_type -> pb.CreateAccountRequest
	6,  // 6: pb.SimpleBank.GetAccount:input_type -> pb.GetAccountRequest
	7,  // 7: pb.SimpleBank.ListAccounts:input_type -> pb.ListAccountsRequest
	8,  // 8: pb.SimpleBank.UpdateAccount:input_type -> pb.UpdateAccountRequest
	9,  // 9: pb.SimpleBank.DeleteAccount:input_type -> pb.DeleteAccountRequest
	10, // 10: pb.SimpleBank.CreateUser:output_type -> pb.CreateUserResponse
	11, // 11: pb.SimpleBank.UpdateUser:output_type -> pb.UpdateUserResponse
	12, // 12: pb.SimpleBank.LoginUser:output_type -> pb.LoginUserResponse
	13, // 13: pb.SimpleBank.VerifyEmail:output_type -> pb.VerifyEmailResponse
	14, // 14: pb.SimpleBank.RenewAccessToken:output_type -> pb.RenewAccessTokenResponse
	15, // 15: pb.SimpleBank.CreateAccount:output_type -> pb.CreateAccountResponse
	16, // 16: pb.SimpleBank.GetAccount:output_type -> pb.GetAccountResponse
	17, // 17: pb.SimpleBank.ListAccounts:output_type -> pb.ListAccountsResponse
	18, // 18: pb.SimpleBank.UpdateAccount:output_type -> pb.UpdateAccountResponse
	19, // 19: pb.SimpleBank.DeleteAccount:output_type -> pb.DeleteAccountResponse
	10, // [10:20] is the sub-list for method output_type
	0,  // [0:10] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_service_simple_bank_proto_init() }
func file_service_simple_bank_proto_init() {
	if File_service_simple_bank_proto != nil {
		return
	}
	file_rpc_create_user_proto_init()
	file_rpc_update_user_proto_init()
	file_rpc_login_user_proto_init()
	file_rpc_verify_email_proto_init()
	file_rpc_renew_access_token_proto_init()
	file_rpc_create_account_proto_init()
	file_rpc_get_account_proto_init()
	file_rpc_list_accounts_proto_init()
	file_rpc_update_account_proto_init()
	file_rpc_delete_account_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_service_simple_bank_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_service_simple_bank_proto_goTypes,
		DependencyIndexes: file_service_simple_bank_proto_depIdxs,
	}.Build()
	File_service_simple_bank_proto = out.File
	file_service_simple_bank_proto_rawDesc = nil
	file_service_simple_bank_proto_goTypes = nil
	file_service_simple_bank_proto_depIdxs = nil
}
