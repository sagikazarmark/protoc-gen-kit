// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: test.proto

package testv1

import (
	subtest "github.com/sagikazarmark/protoc-gen-kit/test/subtest"
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

type TestProcedureRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *TestProcedureRequest) Reset() {
	*x = TestProcedureRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_test_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TestProcedureRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestProcedureRequest) ProtoMessage() {}

func (x *TestProcedureRequest) ProtoReflect() protoreflect.Message {
	mi := &file_test_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TestProcedureRequest.ProtoReflect.Descriptor instead.
func (*TestProcedureRequest) Descriptor() ([]byte, []int) {
	return file_test_proto_rawDescGZIP(), []int{0}
}

type TestProcedureResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *TestProcedureResponse) Reset() {
	*x = TestProcedureResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_test_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TestProcedureResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestProcedureResponse) ProtoMessage() {}

func (x *TestProcedureResponse) ProtoReflect() protoreflect.Message {
	mi := &file_test_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TestProcedureResponse.ProtoReflect.Descriptor instead.
func (*TestProcedureResponse) Descriptor() ([]byte, []int) {
	return file_test_proto_rawDescGZIP(), []int{1}
}

var File_test_proto protoreflect.FileDescriptor

var file_test_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x74, 0x65,
	0x73, 0x74, 0x2e, 0x76, 0x31, 0x1a, 0x15, 0x73, 0x75, 0x62, 0x74, 0x65, 0x73, 0x74, 0x2f, 0x73,
	0x75, 0x62, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x16, 0x0a, 0x14,
	0x54, 0x65, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x64, 0x75, 0x72, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x22, 0x17, 0x0a, 0x15, 0x54, 0x65, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x63,
	0x65, 0x64, 0x75, 0x72, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0xc5, 0x01,
	0x0a, 0x0b, 0x54, 0x65, 0x73, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4e, 0x0a,
	0x0d, 0x54, 0x65, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x64, 0x75, 0x72, 0x65, 0x12, 0x1d,
	0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x50, 0x72, 0x6f,
	0x63, 0x65, 0x64, 0x75, 0x72, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e,
	0x74, 0x65, 0x73, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x63,
	0x65, 0x64, 0x75, 0x72, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x66, 0x0a,
	0x0f, 0x53, 0x75, 0x62, 0x65, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x64, 0x75, 0x72, 0x65,
	0x12, 0x28, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x73, 0x75, 0x62, 0x74, 0x65,
	0x73, 0x74, 0x2e, 0x53, 0x75, 0x62, 0x74, 0x65, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x64,
	0x75, 0x72, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x29, 0x2e, 0x74, 0x65, 0x73,
	0x74, 0x2e, 0x76, 0x31, 0x2e, 0x73, 0x75, 0x62, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x53, 0x75, 0x62,
	0x74, 0x65, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x64, 0x75, 0x72, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x6d, 0x0a, 0x0b, 0x63, 0x6f, 0x6d, 0x2e, 0x74, 0x65, 0x73,
	0x74, 0x2e, 0x76, 0x31, 0x42, 0x0d, 0x54, 0x65, 0x73, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x33, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x73, 0x61, 0x67, 0x69, 0x6b, 0x61, 0x7a, 0x61, 0x72, 0x6d, 0x61, 0x72, 0x6b, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6b, 0x69, 0x74, 0x2f, 0x74,
	0x65, 0x73, 0x74, 0x3b, 0x74, 0x65, 0x73, 0x74, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x54, 0x58, 0x58,
	0xaa, 0x02, 0x07, 0x54, 0x65, 0x73, 0x74, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x07, 0x54, 0x65, 0x73,
	0x74, 0x5c, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_test_proto_rawDescOnce sync.Once
	file_test_proto_rawDescData = file_test_proto_rawDesc
)

func file_test_proto_rawDescGZIP() []byte {
	file_test_proto_rawDescOnce.Do(func() {
		file_test_proto_rawDescData = protoimpl.X.CompressGZIP(file_test_proto_rawDescData)
	})
	return file_test_proto_rawDescData
}

var file_test_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_test_proto_goTypes = []interface{}{
	(*TestProcedureRequest)(nil),             // 0: test.v1.TestProcedureRequest
	(*TestProcedureResponse)(nil),            // 1: test.v1.TestProcedureResponse
	(*subtest.SubtestProcedureRequest)(nil),  // 2: test.v1.subtest.SubtestProcedureRequest
	(*subtest.SubtestProcedureResponse)(nil), // 3: test.v1.subtest.SubtestProcedureResponse
}
var file_test_proto_depIdxs = []int32{
	0, // 0: test.v1.TestService.TestProcedure:input_type -> test.v1.TestProcedureRequest
	2, // 1: test.v1.TestService.SubestProcedure:input_type -> test.v1.subtest.SubtestProcedureRequest
	1, // 2: test.v1.TestService.TestProcedure:output_type -> test.v1.TestProcedureResponse
	3, // 3: test.v1.TestService.SubestProcedure:output_type -> test.v1.subtest.SubtestProcedureResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_test_proto_init() }
func file_test_proto_init() {
	if File_test_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_test_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TestProcedureRequest); i {
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
		file_test_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TestProcedureResponse); i {
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
			RawDescriptor: file_test_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_test_proto_goTypes,
		DependencyIndexes: file_test_proto_depIdxs,
		MessageInfos:      file_test_proto_msgTypes,
	}.Build()
	File_test_proto = out.File
	file_test_proto_rawDesc = nil
	file_test_proto_goTypes = nil
	file_test_proto_depIdxs = nil
}
