// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.0
// 	protoc        (unknown)
// source: file/service/v1/file.proto

package servicev1

import (
	_ "github.com/google/gnostic/openapiv3"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// 前端上传文件所用的HTTP方法
type UploadMethod int32

const (
	UploadMethod_Put  UploadMethod = 0
	UploadMethod_Post UploadMethod = 1
)

// Enum value maps for UploadMethod.
var (
	UploadMethod_name = map[int32]string{
		0: "Put",
		1: "Post",
	}
	UploadMethod_value = map[string]int32{
		"Put":  0,
		"Post": 1,
	}
)

func (x UploadMethod) Enum() *UploadMethod {
	p := new(UploadMethod)
	*p = x
	return p
}

func (x UploadMethod) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (UploadMethod) Descriptor() protoreflect.EnumDescriptor {
	return file_file_service_v1_file_proto_enumTypes[0].Descriptor()
}

func (UploadMethod) Type() protoreflect.EnumType {
	return &file_file_service_v1_file_proto_enumTypes[0]
}

func (x UploadMethod) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use UploadMethod.Descriptor instead.
func (UploadMethod) EnumDescriptor() ([]byte, []int) {
	return file_file_service_v1_file_proto_rawDescGZIP(), []int{0}
}

// 获取对象存储上传链接 - 请求
type OssUploadUrlRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Method        UploadMethod           `protobuf:"varint,1,opt,name=method,proto3,enum=file.service.v1.UploadMethod" json:"method,omitempty"` // 上传文件所用的HTTP方法
	ContentType   *string                `protobuf:"bytes,2,opt,name=contentType,proto3,oneof" json:"contentType,omitempty"`                    // 文件的MIME类型
	BucketName    *string                `protobuf:"bytes,3,opt,name=bucketName,proto3,oneof" json:"bucketName,omitempty"`                      // 文件桶名称，如果不填写，将会根据文件名或者MIME类型进行自动解析。
	FilePath      *string                `protobuf:"bytes,4,opt,name=filePath,proto3,oneof" json:"filePath,omitempty"`                          // 远端的文件路径，可以不填写。
	FileName      *string                `protobuf:"bytes,5,opt,name=fileName,proto3,oneof" json:"fileName,omitempty"`                          // 文件名，如果不填写，则会生成UUID，有同名文件也会改为UUID。
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *OssUploadUrlRequest) Reset() {
	*x = OssUploadUrlRequest{}
	mi := &file_file_service_v1_file_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *OssUploadUrlRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OssUploadUrlRequest) ProtoMessage() {}

func (x *OssUploadUrlRequest) ProtoReflect() protoreflect.Message {
	mi := &file_file_service_v1_file_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OssUploadUrlRequest.ProtoReflect.Descriptor instead.
func (*OssUploadUrlRequest) Descriptor() ([]byte, []int) {
	return file_file_service_v1_file_proto_rawDescGZIP(), []int{0}
}

func (x *OssUploadUrlRequest) GetMethod() UploadMethod {
	if x != nil {
		return x.Method
	}
	return UploadMethod_Put
}

func (x *OssUploadUrlRequest) GetContentType() string {
	if x != nil && x.ContentType != nil {
		return *x.ContentType
	}
	return ""
}

func (x *OssUploadUrlRequest) GetBucketName() string {
	if x != nil && x.BucketName != nil {
		return *x.BucketName
	}
	return ""
}

func (x *OssUploadUrlRequest) GetFilePath() string {
	if x != nil && x.FilePath != nil {
		return *x.FilePath
	}
	return ""
}

func (x *OssUploadUrlRequest) GetFileName() string {
	if x != nil && x.FileName != nil {
		return *x.FileName
	}
	return ""
}

// 获取对象存储上传链接 - 回应
type OssUploadUrlResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UploadUrl     string                 `protobuf:"bytes,1,opt,name=uploadUrl,proto3" json:"uploadUrl,omitempty"`         // 文件的上传链接，默认1个小时的过期时间。
	DownloadUrl   string                 `protobuf:"bytes,2,opt,name=downloadUrl,proto3" json:"downloadUrl,omitempty"`     // 文件的下载链接
	BucketName    *string                `protobuf:"bytes,3,opt,name=bucketName,proto3,oneof" json:"bucketName,omitempty"` // 文件桶名称
	ObjectName    string                 `protobuf:"bytes,4,opt,name=objectName,proto3" json:"objectName,omitempty"`       // 文件名
	FormData      map[string]string      `protobuf:"bytes,5,rep,name=formData,proto3" json:"formData,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *OssUploadUrlResponse) Reset() {
	*x = OssUploadUrlResponse{}
	mi := &file_file_service_v1_file_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *OssUploadUrlResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OssUploadUrlResponse) ProtoMessage() {}

func (x *OssUploadUrlResponse) ProtoReflect() protoreflect.Message {
	mi := &file_file_service_v1_file_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OssUploadUrlResponse.ProtoReflect.Descriptor instead.
func (*OssUploadUrlResponse) Descriptor() ([]byte, []int) {
	return file_file_service_v1_file_proto_rawDescGZIP(), []int{1}
}

func (x *OssUploadUrlResponse) GetUploadUrl() string {
	if x != nil {
		return x.UploadUrl
	}
	return ""
}

func (x *OssUploadUrlResponse) GetDownloadUrl() string {
	if x != nil {
		return x.DownloadUrl
	}
	return ""
}

func (x *OssUploadUrlResponse) GetBucketName() string {
	if x != nil && x.BucketName != nil {
		return *x.BucketName
	}
	return ""
}

func (x *OssUploadUrlResponse) GetObjectName() string {
	if x != nil {
		return x.ObjectName
	}
	return ""
}

func (x *OssUploadUrlResponse) GetFormData() map[string]string {
	if x != nil {
		return x.FormData
	}
	return nil
}

type GetDownloadUrlRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetDownloadUrlRequest) Reset() {
	*x = GetDownloadUrlRequest{}
	mi := &file_file_service_v1_file_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetDownloadUrlRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDownloadUrlRequest) ProtoMessage() {}

func (x *GetDownloadUrlRequest) ProtoReflect() protoreflect.Message {
	mi := &file_file_service_v1_file_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDownloadUrlRequest.ProtoReflect.Descriptor instead.
func (*GetDownloadUrlRequest) Descriptor() ([]byte, []int) {
	return file_file_service_v1_file_proto_rawDescGZIP(), []int{2}
}

type GetDownloadUrlResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetDownloadUrlResponse) Reset() {
	*x = GetDownloadUrlResponse{}
	mi := &file_file_service_v1_file_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetDownloadUrlResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDownloadUrlResponse) ProtoMessage() {}

func (x *GetDownloadUrlResponse) ProtoReflect() protoreflect.Message {
	mi := &file_file_service_v1_file_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDownloadUrlResponse.ProtoReflect.Descriptor instead.
func (*GetDownloadUrlResponse) Descriptor() ([]byte, []int) {
	return file_file_service_v1_file_proto_rawDescGZIP(), []int{3}
}

type ListFileRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Bucket        *string                `protobuf:"bytes,1,opt,name=bucket,proto3,oneof" json:"bucket,omitempty"`        // 文件桶名称
	Folder        *string                `protobuf:"bytes,2,opt,name=folder,proto3,oneof" json:"folder,omitempty"`        // 文件夹名称
	Recursive     *bool                  `protobuf:"varint,3,opt,name=recursive,proto3,oneof" json:"recursive,omitempty"` // 是否递归文件夹
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListFileRequest) Reset() {
	*x = ListFileRequest{}
	mi := &file_file_service_v1_file_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListFileRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListFileRequest) ProtoMessage() {}

func (x *ListFileRequest) ProtoReflect() protoreflect.Message {
	mi := &file_file_service_v1_file_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListFileRequest.ProtoReflect.Descriptor instead.
func (*ListFileRequest) Descriptor() ([]byte, []int) {
	return file_file_service_v1_file_proto_rawDescGZIP(), []int{4}
}

func (x *ListFileRequest) GetBucket() string {
	if x != nil && x.Bucket != nil {
		return *x.Bucket
	}
	return ""
}

func (x *ListFileRequest) GetFolder() string {
	if x != nil && x.Folder != nil {
		return *x.Folder
	}
	return ""
}

func (x *ListFileRequest) GetRecursive() bool {
	if x != nil && x.Recursive != nil {
		return *x.Recursive
	}
	return false
}

type ListFileResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Files         []string               `protobuf:"bytes,1,rep,name=files,proto3" json:"files,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListFileResponse) Reset() {
	*x = ListFileResponse{}
	mi := &file_file_service_v1_file_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListFileResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListFileResponse) ProtoMessage() {}

func (x *ListFileResponse) ProtoReflect() protoreflect.Message {
	mi := &file_file_service_v1_file_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListFileResponse.ProtoReflect.Descriptor instead.
func (*ListFileResponse) Descriptor() ([]byte, []int) {
	return file_file_service_v1_file_proto_rawDescGZIP(), []int{5}
}

func (x *ListFileResponse) GetFiles() []string {
	if x != nil {
		return x.Files
	}
	return nil
}

type DeleteFileRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Bucket        *string                `protobuf:"bytes,1,opt,name=bucket,proto3,oneof" json:"bucket,omitempty"` // 文件桶名称
	Object        *string                `protobuf:"bytes,2,opt,name=object,proto3,oneof" json:"object,omitempty"` // 文件名
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteFileRequest) Reset() {
	*x = DeleteFileRequest{}
	mi := &file_file_service_v1_file_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteFileRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteFileRequest) ProtoMessage() {}

func (x *DeleteFileRequest) ProtoReflect() protoreflect.Message {
	mi := &file_file_service_v1_file_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteFileRequest.ProtoReflect.Descriptor instead.
func (*DeleteFileRequest) Descriptor() ([]byte, []int) {
	return file_file_service_v1_file_proto_rawDescGZIP(), []int{6}
}

func (x *DeleteFileRequest) GetBucket() string {
	if x != nil && x.Bucket != nil {
		return *x.Bucket
	}
	return ""
}

func (x *DeleteFileRequest) GetObject() string {
	if x != nil && x.Object != nil {
		return *x.Object
	}
	return ""
}

type DeleteFileResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteFileResponse) Reset() {
	*x = DeleteFileResponse{}
	mi := &file_file_service_v1_file_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteFileResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteFileResponse) ProtoMessage() {}

func (x *DeleteFileResponse) ProtoReflect() protoreflect.Message {
	mi := &file_file_service_v1_file_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteFileResponse.ProtoReflect.Descriptor instead.
func (*DeleteFileResponse) Descriptor() ([]byte, []int) {
	return file_file_service_v1_file_proto_rawDescGZIP(), []int{7}
}

var File_file_service_v1_file_proto protoreflect.FileDescriptor

var file_file_service_v1_file_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x66, 0x69, 0x6c, 0x65, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x76,
	0x31, 0x2f, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x66, 0x69,
	0x6c, 0x65, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x24, 0x67,
	0x6e, 0x6f, 0x73, 0x74, 0x69, 0x63, 0x2f, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x2f, 0x76,
	0x33, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xd3, 0x04, 0x0a, 0x13, 0x4f, 0x73, 0x73, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x55,
	0x72, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x6f, 0x0a, 0x06, 0x6d, 0x65, 0x74,
	0x68, 0x6f, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1d, 0x2e, 0x66, 0x69, 0x6c, 0x65,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x6c, 0x6f,
	0x61, 0x64, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x42, 0x38, 0xba, 0x47, 0x35, 0x92, 0x02, 0x32,
	0xe4, 0xb8, 0x8a, 0xe4, 0xbc, 0xa0, 0xe6, 0x96, 0x87, 0xe4, 0xbb, 0xb6, 0xe6, 0x89, 0x80, 0xe7,
	0x94, 0xa8, 0xe7, 0x9a, 0x84, 0x48, 0x54, 0x54, 0x50, 0xe6, 0x96, 0xb9, 0xe6, 0xb3, 0x95, 0xef,
	0xbc, 0x8c, 0xe6, 0x94, 0xaf, 0xe6, 0x8c, 0x81, 0x50, 0x4f, 0x53, 0x54, 0xe5, 0x92, 0x8c, 0x50,
	0x55, 0x54, 0x52, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x40, 0x0a, 0x0b, 0x63, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x19, 0xba, 0x47, 0x16, 0x92, 0x02, 0x13, 0xe6, 0x96, 0x87, 0xe4, 0xbb, 0xb6, 0xe7, 0x9a, 0x84,
	0x4d, 0x49, 0x4d, 0x45, 0xe7, 0xb1, 0xbb, 0xe5, 0x9e, 0x8b, 0x48, 0x00, 0x52, 0x0b, 0x63, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x88, 0x01, 0x01, 0x12, 0x86, 0x01, 0x0a,
	0x0a, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x61, 0xba, 0x47, 0x5e, 0x92, 0x02, 0x5b, 0xe6, 0x96, 0x87, 0xe4, 0xbb, 0xb6, 0xe6,
	0xa1, 0xb6, 0xe5, 0x90, 0x8d, 0xe7, 0xa7, 0xb0, 0xef, 0xbc, 0x8c, 0xe5, 0xa6, 0x82, 0xe6, 0x9e,
	0x9c, 0xe4, 0xb8, 0x8d, 0xe5, 0xa1, 0xab, 0xe5, 0x86, 0x99, 0xef, 0xbc, 0x8c, 0xe5, 0xb0, 0x86,
	0xe4, 0xbc, 0x9a, 0xe6, 0xa0, 0xb9, 0xe6, 0x8d, 0xae, 0xe6, 0x96, 0x87, 0xe4, 0xbb, 0xb6, 0xe5,
	0x90, 0x8d, 0xe6, 0x88, 0x96, 0xe8, 0x80, 0x85, 0x4d, 0x49, 0x4d, 0x45, 0xe7, 0xb1, 0xbb, 0xe5,
	0x9e, 0x8b, 0xe8, 0xbf, 0x9b, 0xe8, 0xa1, 0x8c, 0xe8, 0x87, 0xaa, 0xe5, 0x8a, 0xa8, 0xe8, 0xa7,
	0xa3, 0xe6, 0x9e, 0x90, 0x48, 0x01, 0x52, 0x0a, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x4e, 0x61,
	0x6d, 0x65, 0x88, 0x01, 0x01, 0x12, 0x4e, 0x0a, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x50, 0x61, 0x74,
	0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x42, 0x2d, 0xba, 0x47, 0x2a, 0x92, 0x02, 0x27, 0xe8,
	0xbf, 0x9c, 0xe7, 0xab, 0xaf, 0xe7, 0x9a, 0x84, 0xe6, 0x96, 0x87, 0xe4, 0xbb, 0xb6, 0xe8, 0xb7,
	0xaf, 0xe5, 0xbe, 0x84, 0xef, 0xbc, 0x8c, 0xe5, 0x8f, 0xaf, 0xe4, 0xbb, 0xa5, 0xe4, 0xb8, 0x8d,
	0xe5, 0xa1, 0xab, 0xe5, 0x86, 0x99, 0x48, 0x02, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x50, 0x61,
	0x74, 0x68, 0x88, 0x01, 0x01, 0x12, 0x77, 0x0a, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x4e, 0x61, 0x6d,
	0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x42, 0x56, 0xba, 0x47, 0x53, 0x92, 0x02, 0x50, 0xe6,
	0x96, 0x87, 0xe4, 0xbb, 0xb6, 0xe5, 0x90, 0x8d, 0xef, 0xbc, 0x8c, 0xe5, 0xa6, 0x82, 0xe6, 0x9e,
	0x9c, 0xe4, 0xb8, 0x8d, 0xe5, 0xa1, 0xab, 0xe5, 0x86, 0x99, 0xef, 0xbc, 0x8c, 0xe5, 0x88, 0x99,
	0xe4, 0xbc, 0x9a, 0xe7, 0x94, 0x9f, 0xe6, 0x88, 0x90, 0x55, 0x55, 0x49, 0x44, 0xef, 0xbc, 0x8c,
	0xe6, 0x9c, 0x89, 0xe5, 0x90, 0x8c, 0xe5, 0x90, 0x8d, 0xe6, 0x96, 0x87, 0xe4, 0xbb, 0xb6, 0xe4,
	0xb9, 0x9f, 0xe4, 0xbc, 0x9a, 0xe6, 0x94, 0xb9, 0xe4, 0xb8, 0xba, 0x55, 0x55, 0x49, 0x44, 0x48,
	0x03, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x42, 0x0e,
	0x0a, 0x0c, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x42, 0x0d,
	0x0a, 0x0b, 0x5f, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x42, 0x0b, 0x0a,
	0x09, 0x5f, 0x66, 0x69, 0x6c, 0x65, 0x50, 0x61, 0x74, 0x68, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x66,
	0x69, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0xec, 0x03, 0x0a, 0x14, 0x4f, 0x73, 0x73, 0x55,
	0x70, 0x6c, 0x6f, 0x61, 0x64, 0x55, 0x72, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x5b, 0x0a, 0x09, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x55, 0x72, 0x6c, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x3d, 0xba, 0x47, 0x3a, 0x92, 0x02, 0x37, 0xe6, 0x96, 0x87, 0xe4, 0xbb,
	0xb6, 0xe7, 0x9a, 0x84, 0xe4, 0xb8, 0x8a, 0xe4, 0xbc, 0xa0, 0xe9, 0x93, 0xbe, 0xe6, 0x8e, 0xa5,
	0xef, 0xbc, 0x8c, 0xe9, 0xbb, 0x98, 0xe8, 0xae, 0xa4, 0x31, 0xe4, 0xb8, 0xaa, 0xe5, 0xb0, 0x8f,
	0xe6, 0x97, 0xb6, 0xe7, 0x9a, 0x84, 0xe8, 0xbf, 0x87, 0xe6, 0x9c, 0x9f, 0xe6, 0x97, 0xb6, 0xe9,
	0x97, 0xb4, 0x52, 0x09, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x55, 0x72, 0x6c, 0x12, 0x3d, 0x0a,
	0x0b, 0x64, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x55, 0x72, 0x6c, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x1b, 0xba, 0x47, 0x18, 0x92, 0x02, 0x15, 0xe6, 0x96, 0x87, 0xe4, 0xbb, 0xb6,
	0xe7, 0x9a, 0x84, 0xe4, 0xb8, 0x8b, 0xe8, 0xbd, 0xbd, 0xe9, 0x93, 0xbe, 0xe6, 0x8e, 0xa5, 0x52,
	0x0b, 0x64, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x55, 0x72, 0x6c, 0x12, 0x3a, 0x0a, 0x0a,
	0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x15, 0xba, 0x47, 0x12, 0x92, 0x02, 0x0f, 0xe6, 0x96, 0x87, 0xe4, 0xbb, 0xb6, 0xe6, 0xa1,
	0xb6, 0xe5, 0x90, 0x8d, 0xe7, 0xa7, 0xb0, 0x48, 0x00, 0x52, 0x0a, 0x62, 0x75, 0x63, 0x6b, 0x65,
	0x74, 0x4e, 0x61, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x12, 0x2f, 0x0a, 0x0a, 0x6f, 0x62, 0x6a, 0x65,
	0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0f, 0xba, 0x47,
	0x0c, 0x92, 0x02, 0x09, 0xe6, 0x96, 0x87, 0xe4, 0xbb, 0xb6, 0xe5, 0x90, 0x8d, 0x52, 0x0a, 0x6f,
	0x62, 0x6a, 0x65, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x7f, 0x0a, 0x08, 0x66, 0x6f, 0x72,
	0x6d, 0x44, 0x61, 0x74, 0x61, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x33, 0x2e, 0x66, 0x69,
	0x6c, 0x65, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x73,
	0x73, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x55, 0x72, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x2e, 0x46, 0x6f, 0x72, 0x6d, 0x44, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x42, 0x2e, 0xba, 0x47, 0x2b, 0x92, 0x02, 0x28, 0xe8, 0xa1, 0xa8, 0xe5, 0x8d, 0x95, 0xe6, 0x95,
	0xb0, 0xe6, 0x8d, 0xae, 0xef, 0xbc, 0x8c, 0xe4, 0xbd, 0xbf, 0xe7, 0x94, 0xa8, 0x50, 0x4f, 0x53,
	0x54, 0xe6, 0x96, 0xb9, 0xe6, 0xb3, 0x95, 0xe6, 0x97, 0xb6, 0xe5, 0xa1, 0xab, 0xe5, 0x86, 0x99,
	0x52, 0x08, 0x66, 0x6f, 0x72, 0x6d, 0x44, 0x61, 0x74, 0x61, 0x1a, 0x3b, 0x0a, 0x0d, 0x46, 0x6f,
	0x72, 0x6d, 0x44, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b,
	0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x62, 0x75, 0x63, 0x6b,
	0x65, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x17, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x44, 0x6f, 0x77,
	0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x55, 0x72, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22,
	0x18, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x55, 0x72,
	0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0xdd, 0x01, 0x0a, 0x0f, 0x4c, 0x69,
	0x73, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x32, 0x0a,
	0x06, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x15, 0xba,
	0x47, 0x12, 0x92, 0x02, 0x0f, 0xe6, 0x96, 0x87, 0xe4, 0xbb, 0xb6, 0xe6, 0xa1, 0xb6, 0xe5, 0x90,
	0x8d, 0xe7, 0xa7, 0xb0, 0x48, 0x00, 0x52, 0x06, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x88, 0x01,
	0x01, 0x12, 0x32, 0x0a, 0x06, 0x66, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x15, 0xba, 0x47, 0x12, 0x92, 0x02, 0x0f, 0xe6, 0x96, 0x87, 0xe4, 0xbb, 0xb6, 0xe5,
	0xa4, 0xb9, 0xe5, 0x90, 0x8d, 0xe7, 0xa7, 0xb0, 0x48, 0x01, 0x52, 0x06, 0x66, 0x6f, 0x6c, 0x64,
	0x65, 0x72, 0x88, 0x01, 0x01, 0x12, 0x3e, 0x0a, 0x09, 0x72, 0x65, 0x63, 0x75, 0x72, 0x73, 0x69,
	0x76, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x42, 0x1b, 0xba, 0x47, 0x18, 0x92, 0x02, 0x15,
	0xe6, 0x98, 0xaf, 0xe5, 0x90, 0xa6, 0xe9, 0x80, 0x92, 0xe5, 0xbd, 0x92, 0xe6, 0x96, 0x87, 0xe4,
	0xbb, 0xb6, 0xe5, 0xa4, 0xb9, 0x48, 0x02, 0x52, 0x09, 0x72, 0x65, 0x63, 0x75, 0x72, 0x73, 0x69,
	0x76, 0x65, 0x88, 0x01, 0x01, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74,
	0x42, 0x09, 0x0a, 0x07, 0x5f, 0x66, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x42, 0x0c, 0x0a, 0x0a, 0x5f,
	0x72, 0x65, 0x63, 0x75, 0x72, 0x73, 0x69, 0x76, 0x65, 0x22, 0x3c, 0x0a, 0x10, 0x4c, 0x69, 0x73,
	0x74, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x28, 0x0a,
	0x05, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x42, 0x12, 0xba, 0x47,
	0x0f, 0x92, 0x02, 0x0c, 0xe6, 0x96, 0x87, 0xe4, 0xbb, 0xb6, 0xe5, 0x88, 0x97, 0xe8, 0xa1, 0xa8,
	0x52, 0x05, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x22, 0x8b, 0x01, 0x0a, 0x11, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x32, 0x0a,
	0x06, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x15, 0xba,
	0x47, 0x12, 0x92, 0x02, 0x0f, 0xe6, 0x96, 0x87, 0xe4, 0xbb, 0xb6, 0xe6, 0xa1, 0xb6, 0xe5, 0x90,
	0x8d, 0xe7, 0xa7, 0xb0, 0x48, 0x00, 0x52, 0x06, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x88, 0x01,
	0x01, 0x12, 0x2c, 0x0a, 0x06, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x0f, 0xba, 0x47, 0x0c, 0x92, 0x02, 0x09, 0xe6, 0x96, 0x87, 0xe4, 0xbb, 0xb6, 0xe5,
	0x90, 0x8d, 0x48, 0x01, 0x52, 0x06, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x88, 0x01, 0x01, 0x42,
	0x09, 0x0a, 0x07, 0x5f, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x6f,
	0x62, 0x6a, 0x65, 0x63, 0x74, 0x22, 0x14, 0x0a, 0x12, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x46,
	0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2a, 0x21, 0x0a, 0x0c, 0x55,
	0x70, 0x6c, 0x6f, 0x61, 0x64, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x07, 0x0a, 0x03, 0x50,
	0x75, 0x74, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x50, 0x6f, 0x73, 0x74, 0x10, 0x01, 0x32, 0xfd,
	0x02, 0x0a, 0x0b, 0x46, 0x69, 0x6c, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x5d,
	0x0a, 0x0c, 0x4f, 0x73, 0x73, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x55, 0x72, 0x6c, 0x12, 0x24,
	0x2e, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x4f, 0x73, 0x73, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x55, 0x72, 0x6c, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x73, 0x73, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64,
	0x55, 0x72, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x63, 0x0a,
	0x0e, 0x47, 0x65, 0x74, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x55, 0x72, 0x6c, 0x12,
	0x26, 0x2e, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x47, 0x65, 0x74, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x55, 0x72, 0x6c,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x44, 0x6f, 0x77,
	0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x55, 0x72, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x51, 0x0a, 0x08, 0x4c, 0x69, 0x73, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x20,
	0x2e, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x4c, 0x69, 0x73, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x21, 0x2e, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x57, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x46,
	0x69, 0x6c, 0x65, 0x12, 0x22, 0x2e, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x46, 0x69, 0x6c, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0xb1,
	0x01, 0x0a, 0x13, 0x63, 0x6f, 0x6d, 0x2e, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x42, 0x09, 0x46, 0x69, 0x6c, 0x65, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x50, 0x01, 0x5a, 0x31, 0x6b, 0x72, 0x61, 0x74, 0x6f, 0x73, 0x2d, 0x61, 0x64, 0x6d, 0x69,
	0x6e, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x66, 0x69, 0x6c,
	0x65, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x76, 0x31, 0x3b, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x46, 0x53, 0x58, 0xaa, 0x02, 0x0f, 0x46,
	0x69, 0x6c, 0x65, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x56, 0x31, 0xca, 0x02,
	0x0f, 0x46, 0x69, 0x6c, 0x65, 0x5c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5c, 0x56, 0x31,
	0xe2, 0x02, 0x1b, 0x46, 0x69, 0x6c, 0x65, 0x5c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5c,
	0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02,
	0x11, 0x46, 0x69, 0x6c, 0x65, 0x3a, 0x3a, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x3a, 0x3a,
	0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_file_service_v1_file_proto_rawDescOnce sync.Once
	file_file_service_v1_file_proto_rawDescData = file_file_service_v1_file_proto_rawDesc
)

func file_file_service_v1_file_proto_rawDescGZIP() []byte {
	file_file_service_v1_file_proto_rawDescOnce.Do(func() {
		file_file_service_v1_file_proto_rawDescData = protoimpl.X.CompressGZIP(file_file_service_v1_file_proto_rawDescData)
	})
	return file_file_service_v1_file_proto_rawDescData
}

var file_file_service_v1_file_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_file_service_v1_file_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_file_service_v1_file_proto_goTypes = []any{
	(UploadMethod)(0),              // 0: file.service.v1.UploadMethod
	(*OssUploadUrlRequest)(nil),    // 1: file.service.v1.OssUploadUrlRequest
	(*OssUploadUrlResponse)(nil),   // 2: file.service.v1.OssUploadUrlResponse
	(*GetDownloadUrlRequest)(nil),  // 3: file.service.v1.GetDownloadUrlRequest
	(*GetDownloadUrlResponse)(nil), // 4: file.service.v1.GetDownloadUrlResponse
	(*ListFileRequest)(nil),        // 5: file.service.v1.ListFileRequest
	(*ListFileResponse)(nil),       // 6: file.service.v1.ListFileResponse
	(*DeleteFileRequest)(nil),      // 7: file.service.v1.DeleteFileRequest
	(*DeleteFileResponse)(nil),     // 8: file.service.v1.DeleteFileResponse
	nil,                            // 9: file.service.v1.OssUploadUrlResponse.FormDataEntry
}
var file_file_service_v1_file_proto_depIdxs = []int32{
	0, // 0: file.service.v1.OssUploadUrlRequest.method:type_name -> file.service.v1.UploadMethod
	9, // 1: file.service.v1.OssUploadUrlResponse.formData:type_name -> file.service.v1.OssUploadUrlResponse.FormDataEntry
	1, // 2: file.service.v1.FileService.OssUploadUrl:input_type -> file.service.v1.OssUploadUrlRequest
	3, // 3: file.service.v1.FileService.GetDownloadUrl:input_type -> file.service.v1.GetDownloadUrlRequest
	5, // 4: file.service.v1.FileService.ListFile:input_type -> file.service.v1.ListFileRequest
	7, // 5: file.service.v1.FileService.DeleteFile:input_type -> file.service.v1.DeleteFileRequest
	2, // 6: file.service.v1.FileService.OssUploadUrl:output_type -> file.service.v1.OssUploadUrlResponse
	4, // 7: file.service.v1.FileService.GetDownloadUrl:output_type -> file.service.v1.GetDownloadUrlResponse
	6, // 8: file.service.v1.FileService.ListFile:output_type -> file.service.v1.ListFileResponse
	8, // 9: file.service.v1.FileService.DeleteFile:output_type -> file.service.v1.DeleteFileResponse
	6, // [6:10] is the sub-list for method output_type
	2, // [2:6] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_file_service_v1_file_proto_init() }
func file_file_service_v1_file_proto_init() {
	if File_file_service_v1_file_proto != nil {
		return
	}
	file_file_service_v1_file_proto_msgTypes[0].OneofWrappers = []any{}
	file_file_service_v1_file_proto_msgTypes[1].OneofWrappers = []any{}
	file_file_service_v1_file_proto_msgTypes[4].OneofWrappers = []any{}
	file_file_service_v1_file_proto_msgTypes[6].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_file_service_v1_file_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_file_service_v1_file_proto_goTypes,
		DependencyIndexes: file_file_service_v1_file_proto_depIdxs,
		EnumInfos:         file_file_service_v1_file_proto_enumTypes,
		MessageInfos:      file_file_service_v1_file_proto_msgTypes,
	}.Build()
	File_file_service_v1_file_proto = out.File
	file_file_service_v1_file_proto_rawDesc = nil
	file_file_service_v1_file_proto_goTypes = nil
	file_file_service_v1_file_proto_depIdxs = nil
}
