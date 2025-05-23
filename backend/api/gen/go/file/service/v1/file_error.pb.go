// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        (unknown)
// source: file/service/v1/file_error.proto

package servicev1

import (
	_ "github.com/go-kratos/kratos/v2/errors"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type FileErrorReason int32

const (
	// 400
	FileErrorReason_BAD_REQUEST FileErrorReason = 0
	// 403
	FileErrorReason_ACCESS_FORBIDDEN FileErrorReason = 410
	// 404
	FileErrorReason_RESOURCE_NOT_FOUND FileErrorReason = 420
	FileErrorReason_FILE_NOT_FOUND     FileErrorReason = 421
	// 405
	FileErrorReason_METHOD_NOT_ALLOWED FileErrorReason = 430
	// 408
	FileErrorReason_REQUEST_TIMEOUT FileErrorReason = 440
	// 413
	FileErrorReason_FILE_TOO_LARGE FileErrorReason = 450
	// 415
	FileErrorReason_UNSUPPORTED_FILE_TYPE FileErrorReason = 460
	// 500
	FileErrorReason_INTERNAL_SERVER_ERROR FileErrorReason = 500
	FileErrorReason_UPLOAD_FAILED         FileErrorReason = 501
	FileErrorReason_DOWNLOAD_FAILED       FileErrorReason = 502
	// 501
	FileErrorReason_NOT_IMPLEMENTED FileErrorReason = 510
	// 502
	FileErrorReason_NETWORK_ERROR FileErrorReason = 511
	// 503
	FileErrorReason_SERVICE_UNAVAILABLE FileErrorReason = 512
	// 504
	FileErrorReason_NETWORK_TIMEOUT FileErrorReason = 513
	// 505
	FileErrorReason_REQUEST_NOT_SUPPORT FileErrorReason = 514
)

// Enum value maps for FileErrorReason.
var (
	FileErrorReason_name = map[int32]string{
		0:   "BAD_REQUEST",
		410: "ACCESS_FORBIDDEN",
		420: "RESOURCE_NOT_FOUND",
		421: "FILE_NOT_FOUND",
		430: "METHOD_NOT_ALLOWED",
		440: "REQUEST_TIMEOUT",
		450: "FILE_TOO_LARGE",
		460: "UNSUPPORTED_FILE_TYPE",
		500: "INTERNAL_SERVER_ERROR",
		501: "UPLOAD_FAILED",
		502: "DOWNLOAD_FAILED",
		510: "NOT_IMPLEMENTED",
		511: "NETWORK_ERROR",
		512: "SERVICE_UNAVAILABLE",
		513: "NETWORK_TIMEOUT",
		514: "REQUEST_NOT_SUPPORT",
	}
	FileErrorReason_value = map[string]int32{
		"BAD_REQUEST":           0,
		"ACCESS_FORBIDDEN":      410,
		"RESOURCE_NOT_FOUND":    420,
		"FILE_NOT_FOUND":        421,
		"METHOD_NOT_ALLOWED":    430,
		"REQUEST_TIMEOUT":       440,
		"FILE_TOO_LARGE":        450,
		"UNSUPPORTED_FILE_TYPE": 460,
		"INTERNAL_SERVER_ERROR": 500,
		"UPLOAD_FAILED":         501,
		"DOWNLOAD_FAILED":       502,
		"NOT_IMPLEMENTED":       510,
		"NETWORK_ERROR":         511,
		"SERVICE_UNAVAILABLE":   512,
		"NETWORK_TIMEOUT":       513,
		"REQUEST_NOT_SUPPORT":   514,
	}
)

func (x FileErrorReason) Enum() *FileErrorReason {
	p := new(FileErrorReason)
	*p = x
	return p
}

func (x FileErrorReason) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (FileErrorReason) Descriptor() protoreflect.EnumDescriptor {
	return file_file_service_v1_file_error_proto_enumTypes[0].Descriptor()
}

func (FileErrorReason) Type() protoreflect.EnumType {
	return &file_file_service_v1_file_error_proto_enumTypes[0]
}

func (x FileErrorReason) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use FileErrorReason.Descriptor instead.
func (FileErrorReason) EnumDescriptor() ([]byte, []int) {
	return file_file_service_v1_file_error_proto_rawDescGZIP(), []int{0}
}

var File_file_service_v1_file_error_proto protoreflect.FileDescriptor

var file_file_service_v1_file_error_proto_rawDesc = string([]byte{
	0x0a, 0x20, 0x66, 0x69, 0x6c, 0x65, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x76,
	0x31, 0x2f, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0f, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x76, 0x31, 0x1a, 0x13, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x2f, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2a, 0xe7, 0x03, 0x0a, 0x0f, 0x46, 0x69, 0x6c,
	0x65, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12, 0x15, 0x0a, 0x0b,
	0x42, 0x41, 0x44, 0x5f, 0x52, 0x45, 0x51, 0x55, 0x45, 0x53, 0x54, 0x10, 0x00, 0x1a, 0x04, 0xa8,
	0x45, 0x90, 0x03, 0x12, 0x1b, 0x0a, 0x10, 0x41, 0x43, 0x43, 0x45, 0x53, 0x53, 0x5f, 0x46, 0x4f,
	0x52, 0x42, 0x49, 0x44, 0x44, 0x45, 0x4e, 0x10, 0x9a, 0x03, 0x1a, 0x04, 0xa8, 0x45, 0x93, 0x03,
	0x12, 0x1d, 0x0a, 0x12, 0x52, 0x45, 0x53, 0x4f, 0x55, 0x52, 0x43, 0x45, 0x5f, 0x4e, 0x4f, 0x54,
	0x5f, 0x46, 0x4f, 0x55, 0x4e, 0x44, 0x10, 0xa4, 0x03, 0x1a, 0x04, 0xa8, 0x45, 0x94, 0x03, 0x12,
	0x19, 0x0a, 0x0e, 0x46, 0x49, 0x4c, 0x45, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x46, 0x4f, 0x55, 0x4e,
	0x44, 0x10, 0xa5, 0x03, 0x1a, 0x04, 0xa8, 0x45, 0x94, 0x03, 0x12, 0x1d, 0x0a, 0x12, 0x4d, 0x45,
	0x54, 0x48, 0x4f, 0x44, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x41, 0x4c, 0x4c, 0x4f, 0x57, 0x45, 0x44,
	0x10, 0xae, 0x03, 0x1a, 0x04, 0xa8, 0x45, 0x95, 0x03, 0x12, 0x1a, 0x0a, 0x0f, 0x52, 0x45, 0x51,
	0x55, 0x45, 0x53, 0x54, 0x5f, 0x54, 0x49, 0x4d, 0x45, 0x4f, 0x55, 0x54, 0x10, 0xb8, 0x03, 0x1a,
	0x04, 0xa8, 0x45, 0x98, 0x03, 0x12, 0x19, 0x0a, 0x0e, 0x46, 0x49, 0x4c, 0x45, 0x5f, 0x54, 0x4f,
	0x4f, 0x5f, 0x4c, 0x41, 0x52, 0x47, 0x45, 0x10, 0xc2, 0x03, 0x1a, 0x04, 0xa8, 0x45, 0x9d, 0x03,
	0x12, 0x20, 0x0a, 0x15, 0x55, 0x4e, 0x53, 0x55, 0x50, 0x50, 0x4f, 0x52, 0x54, 0x45, 0x44, 0x5f,
	0x46, 0x49, 0x4c, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x10, 0xcc, 0x03, 0x1a, 0x04, 0xa8, 0x45,
	0x9f, 0x03, 0x12, 0x20, 0x0a, 0x15, 0x49, 0x4e, 0x54, 0x45, 0x52, 0x4e, 0x41, 0x4c, 0x5f, 0x53,
	0x45, 0x52, 0x56, 0x45, 0x52, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0xf4, 0x03, 0x1a, 0x04,
	0xa8, 0x45, 0xf4, 0x03, 0x12, 0x18, 0x0a, 0x0d, 0x55, 0x50, 0x4c, 0x4f, 0x41, 0x44, 0x5f, 0x46,
	0x41, 0x49, 0x4c, 0x45, 0x44, 0x10, 0xf5, 0x03, 0x1a, 0x04, 0xa8, 0x45, 0xf4, 0x03, 0x12, 0x1a,
	0x0a, 0x0f, 0x44, 0x4f, 0x57, 0x4e, 0x4c, 0x4f, 0x41, 0x44, 0x5f, 0x46, 0x41, 0x49, 0x4c, 0x45,
	0x44, 0x10, 0xf6, 0x03, 0x1a, 0x04, 0xa8, 0x45, 0xf4, 0x03, 0x12, 0x1a, 0x0a, 0x0f, 0x4e, 0x4f,
	0x54, 0x5f, 0x49, 0x4d, 0x50, 0x4c, 0x45, 0x4d, 0x45, 0x4e, 0x54, 0x45, 0x44, 0x10, 0xfe, 0x03,
	0x1a, 0x04, 0xa8, 0x45, 0xf5, 0x03, 0x12, 0x18, 0x0a, 0x0d, 0x4e, 0x45, 0x54, 0x57, 0x4f, 0x52,
	0x4b, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0xff, 0x03, 0x1a, 0x04, 0xa8, 0x45, 0xf6, 0x03,
	0x12, 0x1e, 0x0a, 0x13, 0x53, 0x45, 0x52, 0x56, 0x49, 0x43, 0x45, 0x5f, 0x55, 0x4e, 0x41, 0x56,
	0x41, 0x49, 0x4c, 0x41, 0x42, 0x4c, 0x45, 0x10, 0x80, 0x04, 0x1a, 0x04, 0xa8, 0x45, 0xf7, 0x03,
	0x12, 0x1a, 0x0a, 0x0f, 0x4e, 0x45, 0x54, 0x57, 0x4f, 0x52, 0x4b, 0x5f, 0x54, 0x49, 0x4d, 0x45,
	0x4f, 0x55, 0x54, 0x10, 0x81, 0x04, 0x1a, 0x04, 0xa8, 0x45, 0xf8, 0x03, 0x12, 0x1e, 0x0a, 0x13,
	0x52, 0x45, 0x51, 0x55, 0x45, 0x53, 0x54, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x53, 0x55, 0x50, 0x50,
	0x4f, 0x52, 0x54, 0x10, 0x82, 0x04, 0x1a, 0x04, 0xa8, 0x45, 0xf9, 0x03, 0x1a, 0x04, 0xa0, 0x45,
	0xf4, 0x03, 0x42, 0xb6, 0x01, 0x0a, 0x13, 0x63, 0x6f, 0x6d, 0x2e, 0x66, 0x69, 0x6c, 0x65, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x42, 0x0e, 0x46, 0x69, 0x6c, 0x65,
	0x45, 0x72, 0x72, 0x6f, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x31, 0x6b, 0x72,
	0x61, 0x74, 0x6f, 0x73, 0x2d, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67,
	0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x66, 0x69, 0x6c, 0x65, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2f, 0x76, 0x31, 0x3b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x76, 0x31, 0xa2,
	0x02, 0x03, 0x46, 0x53, 0x58, 0xaa, 0x02, 0x0f, 0x46, 0x69, 0x6c, 0x65, 0x2e, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x0f, 0x46, 0x69, 0x6c, 0x65, 0x5c, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x1b, 0x46, 0x69, 0x6c, 0x65,
	0x5c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x11, 0x46, 0x69, 0x6c, 0x65, 0x3a, 0x3a,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
})

var (
	file_file_service_v1_file_error_proto_rawDescOnce sync.Once
	file_file_service_v1_file_error_proto_rawDescData []byte
)

func file_file_service_v1_file_error_proto_rawDescGZIP() []byte {
	file_file_service_v1_file_error_proto_rawDescOnce.Do(func() {
		file_file_service_v1_file_error_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_file_service_v1_file_error_proto_rawDesc), len(file_file_service_v1_file_error_proto_rawDesc)))
	})
	return file_file_service_v1_file_error_proto_rawDescData
}

var file_file_service_v1_file_error_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_file_service_v1_file_error_proto_goTypes = []any{
	(FileErrorReason)(0), // 0: file.service.v1.FileErrorReason
}
var file_file_service_v1_file_error_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_file_service_v1_file_error_proto_init() }
func file_file_service_v1_file_error_proto_init() {
	if File_file_service_v1_file_error_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_file_service_v1_file_error_proto_rawDesc), len(file_file_service_v1_file_error_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_file_service_v1_file_error_proto_goTypes,
		DependencyIndexes: file_file_service_v1_file_error_proto_depIdxs,
		EnumInfos:         file_file_service_v1_file_error_proto_enumTypes,
	}.Build()
	File_file_service_v1_file_error_proto = out.File
	file_file_service_v1_file_error_proto_goTypes = nil
	file_file_service_v1_file_error_proto_depIdxs = nil
}
