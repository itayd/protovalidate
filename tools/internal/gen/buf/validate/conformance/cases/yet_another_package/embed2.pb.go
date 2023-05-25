// Copyright 2023 Buf Technologies, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        (unknown)
// source: buf/validate/conformance/cases/yet_another_package/embed2.proto

package yet_another_package

import (
	_ "github.com/bufbuild/protovalidate/tools/internal/gen/buf/validate"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Embed_Enumerated int32

const (
	Embed_ENUMERATED_UNSPECIFIED Embed_Enumerated = 0
	Embed_ENUMERATED_VALUE       Embed_Enumerated = 1
)

// Enum value maps for Embed_Enumerated.
var (
	Embed_Enumerated_name = map[int32]string{
		0: "ENUMERATED_UNSPECIFIED",
		1: "ENUMERATED_VALUE",
	}
	Embed_Enumerated_value = map[string]int32{
		"ENUMERATED_UNSPECIFIED": 0,
		"ENUMERATED_VALUE":       1,
	}
)

func (x Embed_Enumerated) Enum() *Embed_Enumerated {
	p := new(Embed_Enumerated)
	*p = x
	return p
}

func (x Embed_Enumerated) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Embed_Enumerated) Descriptor() protoreflect.EnumDescriptor {
	return file_buf_validate_conformance_cases_yet_another_package_embed2_proto_enumTypes[0].Descriptor()
}

func (Embed_Enumerated) Type() protoreflect.EnumType {
	return &file_buf_validate_conformance_cases_yet_another_package_embed2_proto_enumTypes[0]
}

func (x Embed_Enumerated) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Embed_Enumerated.Descriptor instead.
func (Embed_Enumerated) EnumDescriptor() ([]byte, []int) {
	return file_buf_validate_conformance_cases_yet_another_package_embed2_proto_rawDescGZIP(), []int{0, 0}
}

// Validate message embedding across packages.
type Embed struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Val int64 `protobuf:"varint,1,opt,name=val,proto3" json:"val,omitempty"`
}

func (x *Embed) Reset() {
	*x = Embed{}
	if protoimpl.UnsafeEnabled {
		mi := &file_buf_validate_conformance_cases_yet_another_package_embed2_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Embed) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Embed) ProtoMessage() {}

func (x *Embed) ProtoReflect() protoreflect.Message {
	mi := &file_buf_validate_conformance_cases_yet_another_package_embed2_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Embed.ProtoReflect.Descriptor instead.
func (*Embed) Descriptor() ([]byte, []int) {
	return file_buf_validate_conformance_cases_yet_another_package_embed2_proto_rawDescGZIP(), []int{0}
}

func (x *Embed) GetVal() int64 {
	if x != nil {
		return x.Val
	}
	return 0
}

var File_buf_validate_conformance_cases_yet_another_package_embed2_proto protoreflect.FileDescriptor

var file_buf_validate_conformance_cases_yet_another_package_embed2_proto_rawDesc = []byte{
	0x0a, 0x3f, 0x62, 0x75, 0x66, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x63,
	0x6f, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x6e, 0x63, 0x65, 0x2f, 0x63, 0x61, 0x73, 0x65, 0x73,
	0x2f, 0x79, 0x65, 0x74, 0x5f, 0x61, 0x6e, 0x6f, 0x74, 0x68, 0x65, 0x72, 0x5f, 0x70, 0x61, 0x63,
	0x6b, 0x61, 0x67, 0x65, 0x2f, 0x65, 0x6d, 0x62, 0x65, 0x64, 0x32, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x32, 0x62, 0x75, 0x66, 0x2e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e,
	0x63, 0x6f, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x6e, 0x63, 0x65, 0x2e, 0x63, 0x61, 0x73, 0x65,
	0x73, 0x2e, 0x79, 0x65, 0x74, 0x5f, 0x61, 0x6e, 0x6f, 0x74, 0x68, 0x65, 0x72, 0x5f, 0x70, 0x61,
	0x63, 0x6b, 0x61, 0x67, 0x65, 0x1a, 0x1b, 0x62, 0x75, 0x66, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x63, 0x0a, 0x05, 0x45, 0x6d, 0x62, 0x65, 0x64, 0x12, 0x1a, 0x0a, 0x03, 0x76,
	0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x42, 0x08, 0xfa, 0xf7, 0x18, 0x04, 0x22, 0x02,
	0x20, 0x00, 0x52, 0x03, 0x76, 0x61, 0x6c, 0x22, 0x3e, 0x0a, 0x0a, 0x45, 0x6e, 0x75, 0x6d, 0x65,
	0x72, 0x61, 0x74, 0x65, 0x64, 0x12, 0x1a, 0x0a, 0x16, 0x45, 0x4e, 0x55, 0x4d, 0x45, 0x52, 0x41,
	0x54, 0x45, 0x44, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10,
	0x00, 0x12, 0x14, 0x0a, 0x10, 0x45, 0x4e, 0x55, 0x4d, 0x45, 0x52, 0x41, 0x54, 0x45, 0x44, 0x5f,
	0x56, 0x41, 0x4c, 0x55, 0x45, 0x10, 0x01, 0x42, 0x94, 0x03, 0x0a, 0x36, 0x63, 0x6f, 0x6d, 0x2e,
	0x62, 0x75, 0x66, 0x2e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x63, 0x6f, 0x6e,
	0x66, 0x6f, 0x72, 0x6d, 0x61, 0x6e, 0x63, 0x65, 0x2e, 0x63, 0x61, 0x73, 0x65, 0x73, 0x2e, 0x79,
	0x65, 0x74, 0x5f, 0x61, 0x6e, 0x6f, 0x74, 0x68, 0x65, 0x72, 0x5f, 0x70, 0x61, 0x63, 0x6b, 0x61,
	0x67, 0x65, 0x42, 0x0b, 0x45, 0x6d, 0x62, 0x65, 0x64, 0x32, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50,
	0x01, 0x5a, 0x67, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x62, 0x75,
	0x66, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x65, 0x2f, 0x74, 0x6f, 0x6f, 0x6c, 0x73, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72,
	0x6e, 0x61, 0x6c, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x62, 0x75, 0x66, 0x2f, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x65, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x6e, 0x63, 0x65,
	0x2f, 0x63, 0x61, 0x73, 0x65, 0x73, 0x2f, 0x79, 0x65, 0x74, 0x5f, 0x61, 0x6e, 0x6f, 0x74, 0x68,
	0x65, 0x72, 0x5f, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0xa2, 0x02, 0x05, 0x42, 0x56, 0x43,
	0x43, 0x59, 0xaa, 0x02, 0x30, 0x42, 0x75, 0x66, 0x2e, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x65, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x6e, 0x63, 0x65, 0x2e, 0x43, 0x61,
	0x73, 0x65, 0x73, 0x2e, 0x59, 0x65, 0x74, 0x41, 0x6e, 0x6f, 0x74, 0x68, 0x65, 0x72, 0x50, 0x61,
	0x63, 0x6b, 0x61, 0x67, 0x65, 0xca, 0x02, 0x30, 0x42, 0x75, 0x66, 0x5c, 0x56, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x65, 0x5c, 0x43, 0x6f, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x6e, 0x63, 0x65,
	0x5c, 0x43, 0x61, 0x73, 0x65, 0x73, 0x5c, 0x59, 0x65, 0x74, 0x41, 0x6e, 0x6f, 0x74, 0x68, 0x65,
	0x72, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0xe2, 0x02, 0x3c, 0x42, 0x75, 0x66, 0x5c, 0x56,
	0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x5c, 0x43, 0x6f, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61,
	0x6e, 0x63, 0x65, 0x5c, 0x43, 0x61, 0x73, 0x65, 0x73, 0x5c, 0x59, 0x65, 0x74, 0x41, 0x6e, 0x6f,
	0x74, 0x68, 0x65, 0x72, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x5c, 0x47, 0x50, 0x42, 0x4d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x34, 0x42, 0x75, 0x66, 0x3a, 0x3a, 0x56,
	0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x3a, 0x3a, 0x43, 0x6f, 0x6e, 0x66, 0x6f, 0x72, 0x6d,
	0x61, 0x6e, 0x63, 0x65, 0x3a, 0x3a, 0x43, 0x61, 0x73, 0x65, 0x73, 0x3a, 0x3a, 0x59, 0x65, 0x74,
	0x41, 0x6e, 0x6f, 0x74, 0x68, 0x65, 0x72, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_buf_validate_conformance_cases_yet_another_package_embed2_proto_rawDescOnce sync.Once
	file_buf_validate_conformance_cases_yet_another_package_embed2_proto_rawDescData = file_buf_validate_conformance_cases_yet_another_package_embed2_proto_rawDesc
)

func file_buf_validate_conformance_cases_yet_another_package_embed2_proto_rawDescGZIP() []byte {
	file_buf_validate_conformance_cases_yet_another_package_embed2_proto_rawDescOnce.Do(func() {
		file_buf_validate_conformance_cases_yet_another_package_embed2_proto_rawDescData = protoimpl.X.CompressGZIP(file_buf_validate_conformance_cases_yet_another_package_embed2_proto_rawDescData)
	})
	return file_buf_validate_conformance_cases_yet_another_package_embed2_proto_rawDescData
}

var file_buf_validate_conformance_cases_yet_another_package_embed2_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_buf_validate_conformance_cases_yet_another_package_embed2_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_buf_validate_conformance_cases_yet_another_package_embed2_proto_goTypes = []interface{}{
	(Embed_Enumerated)(0), // 0: buf.validate.conformance.cases.yet_another_package.Embed.Enumerated
	(*Embed)(nil),         // 1: buf.validate.conformance.cases.yet_another_package.Embed
}
var file_buf_validate_conformance_cases_yet_another_package_embed2_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_buf_validate_conformance_cases_yet_another_package_embed2_proto_init() }
func file_buf_validate_conformance_cases_yet_another_package_embed2_proto_init() {
	if File_buf_validate_conformance_cases_yet_another_package_embed2_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_buf_validate_conformance_cases_yet_another_package_embed2_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Embed); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_buf_validate_conformance_cases_yet_another_package_embed2_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_buf_validate_conformance_cases_yet_another_package_embed2_proto_goTypes,
		DependencyIndexes: file_buf_validate_conformance_cases_yet_another_package_embed2_proto_depIdxs,
		EnumInfos:         file_buf_validate_conformance_cases_yet_another_package_embed2_proto_enumTypes,
		MessageInfos:      file_buf_validate_conformance_cases_yet_another_package_embed2_proto_msgTypes,
	}.Build()
	File_buf_validate_conformance_cases_yet_another_package_embed2_proto = out.File
	file_buf_validate_conformance_cases_yet_another_package_embed2_proto_rawDesc = nil
	file_buf_validate_conformance_cases_yet_another_package_embed2_proto_goTypes = nil
	file_buf_validate_conformance_cases_yet_another_package_embed2_proto_depIdxs = nil
}
