// Code generated by protoc-gen-go-pulsar. DO NOT EDIT.
package mint

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.0
// 	protoc        (unknown)
// source: layer/mint/query.proto

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_layer_mint_query_proto protoreflect.FileDescriptor

var file_layer_mint_query_proto_rawDesc = []byte{
	0x0a, 0x16, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x2f, 0x6d, 0x69, 0x6e, 0x74, 0x2f, 0x71, 0x75, 0x65,
	0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x2e,
	0x6d, 0x69, 0x6e, 0x74, 0x32, 0x07, 0x0a, 0x05, 0x51, 0x75, 0x65, 0x72, 0x79, 0x42, 0x82, 0x01,
	0x0a, 0x0e, 0x63, 0x6f, 0x6d, 0x2e, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x2e, 0x6d, 0x69, 0x6e, 0x74,
	0x42, 0x0a, 0x51, 0x75, 0x65, 0x72, 0x79, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x1b,
	0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x73, 0x64, 0x6b, 0x2e, 0x69, 0x6f, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x2f, 0x6d, 0x69, 0x6e, 0x74, 0xa2, 0x02, 0x03, 0x4c, 0x4d,
	0x58, 0xaa, 0x02, 0x0a, 0x4c, 0x61, 0x79, 0x65, 0x72, 0x2e, 0x4d, 0x69, 0x6e, 0x74, 0xca, 0x02,
	0x0a, 0x4c, 0x61, 0x79, 0x65, 0x72, 0x5c, 0x4d, 0x69, 0x6e, 0x74, 0xe2, 0x02, 0x16, 0x4c, 0x61,
	0x79, 0x65, 0x72, 0x5c, 0x4d, 0x69, 0x6e, 0x74, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0b, 0x4c, 0x61, 0x79, 0x65, 0x72, 0x3a, 0x3a, 0x4d, 0x69,
	0x6e, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_layer_mint_query_proto_goTypes = []interface{}{}
var file_layer_mint_query_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_layer_mint_query_proto_init() }
func file_layer_mint_query_proto_init() {
	if File_layer_mint_query_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_layer_mint_query_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_layer_mint_query_proto_goTypes,
		DependencyIndexes: file_layer_mint_query_proto_depIdxs,
	}.Build()
	File_layer_mint_query_proto = out.File
	file_layer_mint_query_proto_rawDesc = nil
	file_layer_mint_query_proto_goTypes = nil
	file_layer_mint_query_proto_depIdxs = nil
}
