// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        (unknown)
// source: user/service/v1/tenant.proto

package servicev1

import (
	_ "github.com/google/gnostic/openapiv3"
	v1 "github.com/tx7do/kratos-bootstrap/api/gen/go/pagination/v1"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	fieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
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

// 租户
type Tenant struct {
	state          protoimpl.MessageState `protogen:"open.v1"`
	Id             *uint32                `protobuf:"varint,1,opt,name=id,proto3,oneof" json:"id,omitempty"`                                      // 租户ID
	Name           *string                `protobuf:"bytes,2,opt,name=name,proto3,oneof" json:"name,omitempty"`                                   // 租户名称
	Code           *string                `protobuf:"bytes,3,opt,name=code,proto3,oneof" json:"code,omitempty"`                                   // 租户编码
	MemberCount    *int32                 `protobuf:"varint,4,opt,name=member_count,json=memberCount,proto3,oneof" json:"member_count,omitempty"` // 成员数量
	Status         *string                `protobuf:"bytes,5,opt,name=status,proto3,oneof" json:"status,omitempty"`
	Remark         *string                `protobuf:"bytes,6,opt,name=remark,proto3,oneof" json:"remark,omitempty"`                      // 备注
	CreateBy       *uint32                `protobuf:"varint,7,opt,name=create_by,json=createBy,proto3,oneof" json:"create_by,omitempty"` // 创建者ID
	UpdateBy       *uint32                `protobuf:"varint,8,opt,name=update_by,json=updateBy,proto3,oneof" json:"update_by,omitempty"` // 更新者ID
	SubscriptionAt *timestamppb.Timestamp `protobuf:"bytes,10,opt,name=subscription_at,json=subscriptionAt,proto3,oneof" json:"subscription_at,omitempty"`
	UnsubscribeAt  *timestamppb.Timestamp `protobuf:"bytes,11,opt,name=unsubscribe_at,json=unsubscribeAt,proto3,oneof" json:"unsubscribe_at,omitempty"`
	CreateTime     *timestamppb.Timestamp `protobuf:"bytes,200,opt,name=create_time,json=createTime,proto3,oneof" json:"create_time,omitempty"` // 创建时间
	UpdateTime     *timestamppb.Timestamp `protobuf:"bytes,201,opt,name=update_time,json=updateTime,proto3,oneof" json:"update_time,omitempty"` // 更新时间
	DeleteTime     *timestamppb.Timestamp `protobuf:"bytes,202,opt,name=delete_time,json=deleteTime,proto3,oneof" json:"delete_time,omitempty"` // 删除时间
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *Tenant) Reset() {
	*x = Tenant{}
	mi := &file_user_service_v1_tenant_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Tenant) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Tenant) ProtoMessage() {}

func (x *Tenant) ProtoReflect() protoreflect.Message {
	mi := &file_user_service_v1_tenant_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Tenant.ProtoReflect.Descriptor instead.
func (*Tenant) Descriptor() ([]byte, []int) {
	return file_user_service_v1_tenant_proto_rawDescGZIP(), []int{0}
}

func (x *Tenant) GetId() uint32 {
	if x != nil && x.Id != nil {
		return *x.Id
	}
	return 0
}

func (x *Tenant) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

func (x *Tenant) GetCode() string {
	if x != nil && x.Code != nil {
		return *x.Code
	}
	return ""
}

func (x *Tenant) GetMemberCount() int32 {
	if x != nil && x.MemberCount != nil {
		return *x.MemberCount
	}
	return 0
}

func (x *Tenant) GetStatus() string {
	if x != nil && x.Status != nil {
		return *x.Status
	}
	return ""
}

func (x *Tenant) GetRemark() string {
	if x != nil && x.Remark != nil {
		return *x.Remark
	}
	return ""
}

func (x *Tenant) GetCreateBy() uint32 {
	if x != nil && x.CreateBy != nil {
		return *x.CreateBy
	}
	return 0
}

func (x *Tenant) GetUpdateBy() uint32 {
	if x != nil && x.UpdateBy != nil {
		return *x.UpdateBy
	}
	return 0
}

func (x *Tenant) GetSubscriptionAt() *timestamppb.Timestamp {
	if x != nil {
		return x.SubscriptionAt
	}
	return nil
}

func (x *Tenant) GetUnsubscribeAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UnsubscribeAt
	}
	return nil
}

func (x *Tenant) GetCreateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

func (x *Tenant) GetUpdateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdateTime
	}
	return nil
}

func (x *Tenant) GetDeleteTime() *timestamppb.Timestamp {
	if x != nil {
		return x.DeleteTime
	}
	return nil
}

// 租户列表 - 答复
type ListTenantResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Items         []*Tenant              `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
	Total         uint32                 `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListTenantResponse) Reset() {
	*x = ListTenantResponse{}
	mi := &file_user_service_v1_tenant_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListTenantResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListTenantResponse) ProtoMessage() {}

func (x *ListTenantResponse) ProtoReflect() protoreflect.Message {
	mi := &file_user_service_v1_tenant_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListTenantResponse.ProtoReflect.Descriptor instead.
func (*ListTenantResponse) Descriptor() ([]byte, []int) {
	return file_user_service_v1_tenant_proto_rawDescGZIP(), []int{1}
}

func (x *ListTenantResponse) GetItems() []*Tenant {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *ListTenantResponse) GetTotal() uint32 {
	if x != nil {
		return x.Total
	}
	return 0
}

// 租户数据 - 请求
type GetTenantRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            uint32                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetTenantRequest) Reset() {
	*x = GetTenantRequest{}
	mi := &file_user_service_v1_tenant_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetTenantRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTenantRequest) ProtoMessage() {}

func (x *GetTenantRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_service_v1_tenant_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTenantRequest.ProtoReflect.Descriptor instead.
func (*GetTenantRequest) Descriptor() ([]byte, []int) {
	return file_user_service_v1_tenant_proto_rawDescGZIP(), []int{2}
}

func (x *GetTenantRequest) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

// 创建租户 - 请求
type CreateTenantRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Data          *Tenant                `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	OperatorId    *uint32                `protobuf:"varint,2,opt,name=operator_id,json=operatorId,proto3,oneof" json:"operator_id,omitempty"` // 操作用户ID
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateTenantRequest) Reset() {
	*x = CreateTenantRequest{}
	mi := &file_user_service_v1_tenant_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateTenantRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTenantRequest) ProtoMessage() {}

func (x *CreateTenantRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_service_v1_tenant_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTenantRequest.ProtoReflect.Descriptor instead.
func (*CreateTenantRequest) Descriptor() ([]byte, []int) {
	return file_user_service_v1_tenant_proto_rawDescGZIP(), []int{3}
}

func (x *CreateTenantRequest) GetData() *Tenant {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *CreateTenantRequest) GetOperatorId() uint32 {
	if x != nil && x.OperatorId != nil {
		return *x.OperatorId
	}
	return 0
}

// 更新租户 -请求
type UpdateTenantRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	OperatorId    *uint32                `protobuf:"varint,1,opt,name=operator_id,json=operatorId,proto3,oneof" json:"operator_id,omitempty"` // 操作用户ID
	Data          *Tenant                `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	UpdateMask    *fieldmaskpb.FieldMask `protobuf:"bytes,3,opt,name=update_mask,json=updateMask,proto3" json:"update_mask,omitempty"`              // 要更新的字段列表
	AllowMissing  *bool                  `protobuf:"varint,4,opt,name=allow_missing,json=allowMissing,proto3,oneof" json:"allow_missing,omitempty"` // 如果设置为true的时候，资源不存在则会新增(插入)，并且在这种情况下`updateMask`字段将会被忽略。
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateTenantRequest) Reset() {
	*x = UpdateTenantRequest{}
	mi := &file_user_service_v1_tenant_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateTenantRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateTenantRequest) ProtoMessage() {}

func (x *UpdateTenantRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_service_v1_tenant_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateTenantRequest.ProtoReflect.Descriptor instead.
func (*UpdateTenantRequest) Descriptor() ([]byte, []int) {
	return file_user_service_v1_tenant_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateTenantRequest) GetOperatorId() uint32 {
	if x != nil && x.OperatorId != nil {
		return *x.OperatorId
	}
	return 0
}

func (x *UpdateTenantRequest) GetData() *Tenant {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *UpdateTenantRequest) GetUpdateMask() *fieldmaskpb.FieldMask {
	if x != nil {
		return x.UpdateMask
	}
	return nil
}

func (x *UpdateTenantRequest) GetAllowMissing() bool {
	if x != nil && x.AllowMissing != nil {
		return *x.AllowMissing
	}
	return false
}

// 删除租户 - 请求
type DeleteTenantRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            uint32                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	OperatorId    *uint32                `protobuf:"varint,2,opt,name=operator_id,json=operatorId,proto3,oneof" json:"operator_id,omitempty"` // 操作用户ID
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteTenantRequest) Reset() {
	*x = DeleteTenantRequest{}
	mi := &file_user_service_v1_tenant_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteTenantRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteTenantRequest) ProtoMessage() {}

func (x *DeleteTenantRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_service_v1_tenant_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteTenantRequest.ProtoReflect.Descriptor instead.
func (*DeleteTenantRequest) Descriptor() ([]byte, []int) {
	return file_user_service_v1_tenant_proto_rawDescGZIP(), []int{5}
}

func (x *DeleteTenantRequest) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *DeleteTenantRequest) GetOperatorId() uint32 {
	if x != nil && x.OperatorId != nil {
		return *x.OperatorId
	}
	return 0
}

var File_user_service_v1_tenant_proto protoreflect.FileDescriptor

var file_user_service_v1_tenant_proto_rawDesc = string([]byte{
	0x0a, 0x1c, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x76,
	0x31, 0x2f, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f,
	0x75, 0x73, 0x65, 0x72, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x1a,
	0x24, 0x67, 0x6e, 0x6f, 0x73, 0x74, 0x69, 0x63, 0x2f, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69,
	0x2f, 0x76, 0x33, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8b, 0x08, 0x0a, 0x06, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74,
	0x12, 0x23, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x42, 0x0e, 0xba, 0x47,
	0x0b, 0x92, 0x02, 0x08, 0xe7, 0xa7, 0x9f, 0xe6, 0x88, 0xb7, 0x49, 0x44, 0x48, 0x00, 0x52, 0x02,
	0x69, 0x64, 0x88, 0x01, 0x01, 0x12, 0x2b, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x12, 0xba, 0x47, 0x0f, 0x92, 0x02, 0x0c, 0xe7, 0xa7, 0x9f, 0xe6, 0x88,
	0xb7, 0xe5, 0x90, 0x8d, 0xe7, 0xa7, 0xb0, 0x48, 0x01, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x88,
	0x01, 0x01, 0x12, 0x2b, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x12, 0xba, 0x47, 0x0f, 0x92, 0x02, 0x0c, 0xe7, 0xa7, 0x9f, 0xe6, 0x88, 0xb7, 0xe7, 0xbc,
	0x96, 0xe7, 0xa0, 0x81, 0x48, 0x02, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x88, 0x01, 0x01, 0x12,
	0x3a, 0x0a, 0x0c, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x05, 0x42, 0x12, 0xba, 0x47, 0x0f, 0x92, 0x02, 0x0c, 0xe6, 0x88, 0x90,
	0xe5, 0x91, 0x98, 0xe6, 0x95, 0xb0, 0xe9, 0x87, 0x8f, 0x48, 0x03, 0x52, 0x0b, 0x6d, 0x65, 0x6d,
	0x62, 0x65, 0x72, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x88, 0x01, 0x01, 0x12, 0x3f, 0x0a, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x42, 0x22, 0xba, 0x47, 0x1f,
	0xc2, 0x01, 0x04, 0x12, 0x02, 0x4f, 0x4e, 0xc2, 0x01, 0x05, 0x12, 0x03, 0x4f, 0x46, 0x46, 0x8a,
	0x02, 0x04, 0x1a, 0x02, 0x4f, 0x4e, 0x92, 0x02, 0x06, 0xe7, 0x8a, 0xb6, 0xe6, 0x80, 0x81, 0x48,
	0x04, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x88, 0x01, 0x01, 0x12, 0x29, 0x0a, 0x06,
	0x72, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0c, 0xba, 0x47,
	0x09, 0x92, 0x02, 0x06, 0xe5, 0xa4, 0x87, 0xe6, 0xb3, 0xa8, 0x48, 0x05, 0x52, 0x06, 0x72, 0x65,
	0x6d, 0x61, 0x72, 0x6b, 0x88, 0x01, 0x01, 0x12, 0x33, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x5f, 0x62, 0x79, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0d, 0x42, 0x11, 0xba, 0x47, 0x0e, 0x92,
	0x02, 0x0b, 0xe5, 0x88, 0x9b, 0xe5, 0xbb, 0xba, 0xe8, 0x80, 0x85, 0x49, 0x44, 0x48, 0x06, 0x52,
	0x08, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x79, 0x88, 0x01, 0x01, 0x12, 0x33, 0x0a, 0x09,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x62, 0x79, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0d, 0x42,
	0x11, 0xba, 0x47, 0x0e, 0x92, 0x02, 0x0b, 0xe6, 0x9b, 0xb4, 0xe6, 0x96, 0xb0, 0xe8, 0x80, 0x85,
	0x49, 0x44, 0x48, 0x07, 0x52, 0x08, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x79, 0x88, 0x01,
	0x01, 0x12, 0x5c, 0x0a, 0x0f, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x61, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x12, 0xba, 0x47, 0x0f, 0x92, 0x02, 0x0c, 0xe8, 0xae,
	0xa2, 0xe9, 0x98, 0x85, 0xe6, 0x97, 0xb6, 0xe9, 0x97, 0xb4, 0x48, 0x08, 0x52, 0x0e, 0x73, 0x75,
	0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x74, 0x88, 0x01, 0x01, 0x12,
	0x5a, 0x0a, 0x0e, 0x75, 0x6e, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x5f, 0x61,
	0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x42, 0x12, 0xba, 0x47, 0x0f, 0x92, 0x02, 0x0c, 0xe9, 0x80, 0x80, 0xe8, 0xae,
	0xa2, 0xe6, 0x97, 0xb6, 0xe9, 0x97, 0xb4, 0x48, 0x09, 0x52, 0x0d, 0x75, 0x6e, 0x73, 0x75, 0x62,
	0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x41, 0x74, 0x88, 0x01, 0x01, 0x12, 0x55, 0x0a, 0x0b, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0xc8, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x12, 0xba,
	0x47, 0x0f, 0x92, 0x02, 0x0c, 0xe5, 0x88, 0x9b, 0xe5, 0xbb, 0xba, 0xe6, 0x97, 0xb6, 0xe9, 0x97,
	0xb4, 0x48, 0x0a, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x88,
	0x01, 0x01, 0x12, 0x55, 0x0a, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d,
	0x65, 0x18, 0xc9, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x42, 0x12, 0xba, 0x47, 0x0f, 0x92, 0x02, 0x0c, 0xe6, 0x9b, 0xb4, 0xe6,
	0x96, 0xb0, 0xe6, 0x97, 0xb6, 0xe9, 0x97, 0xb4, 0x48, 0x0b, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x12, 0x55, 0x0a, 0x0b, 0x64, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0xca, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x12, 0xba, 0x47, 0x0f,
	0x92, 0x02, 0x0c, 0xe5, 0x88, 0xa0, 0xe9, 0x99, 0xa4, 0xe6, 0x97, 0xb6, 0xe9, 0x97, 0xb4, 0x48,
	0x0c, 0x52, 0x0a, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x88, 0x01, 0x01,
	0x42, 0x05, 0x0a, 0x03, 0x5f, 0x69, 0x64, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x42, 0x07, 0x0a, 0x05, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x42, 0x0f, 0x0a, 0x0d, 0x5f, 0x6d, 0x65,
	0x6d, 0x62, 0x65, 0x72, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x72, 0x65, 0x6d, 0x61, 0x72, 0x6b,
	0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x62, 0x79, 0x42, 0x0c,
	0x0a, 0x0a, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x62, 0x79, 0x42, 0x12, 0x0a, 0x10,
	0x5f, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x61, 0x74,
	0x42, 0x11, 0x0a, 0x0f, 0x5f, 0x75, 0x6e, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65,
	0x5f, 0x61, 0x74, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x74,
	0x69, 0x6d, 0x65, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x74,
	0x69, 0x6d, 0x65, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x5f, 0x74,
	0x69, 0x6d, 0x65, 0x22, 0x59, 0x0a, 0x12, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x65, 0x6e, 0x61, 0x6e,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2d, 0x0a, 0x05, 0x69, 0x74, 0x65,
	0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x65, 0x6e, 0x61, 0x6e,
	0x74, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61,
	0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x22, 0x22,
	0x0a, 0x10, 0x47, 0x65, 0x74, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02,
	0x69, 0x64, 0x22, 0x90, 0x01, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x65, 0x6e,
	0x61, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2b, 0x0a, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x65, 0x6e, 0x61, 0x6e,
	0x74, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x3c, 0x0a, 0x0b, 0x6f, 0x70, 0x65, 0x72, 0x61,
	0x74, 0x6f, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x42, 0x16, 0xba, 0x47,
	0x13, 0x18, 0x01, 0x92, 0x02, 0x0e, 0xe6, 0x93, 0x8d, 0xe4, 0xbd, 0x9c, 0xe7, 0x94, 0xa8, 0xe6,
	0x88, 0xb7, 0x49, 0x44, 0x48, 0x00, 0x52, 0x0a, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72,
	0x49, 0x64, 0x88, 0x01, 0x01, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74,
	0x6f, 0x72, 0x5f, 0x69, 0x64, 0x22, 0xce, 0x03, 0x0a, 0x13, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3c, 0x0a,
	0x0b, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0d, 0x42, 0x16, 0xba, 0x47, 0x13, 0x18, 0x01, 0x92, 0x02, 0x0e, 0xe6, 0x93, 0x8d, 0xe4,
	0xbd, 0x9c, 0xe7, 0x94, 0xa8, 0xe6, 0x88, 0xb7, 0x49, 0x44, 0x48, 0x00, 0x52, 0x0a, 0x6f, 0x70,
	0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x49, 0x64, 0x88, 0x01, 0x01, 0x12, 0x2b, 0x0a, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x75, 0x73, 0x65, 0x72,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x65, 0x6e, 0x61,
	0x6e, 0x74, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x73, 0x0a, 0x0b, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x46, 0x69, 0x65, 0x6c, 0x64, 0x4d, 0x61, 0x73, 0x6b, 0x42, 0x36, 0xba, 0x47, 0x33, 0x3a, 0x16,
	0x12, 0x14, 0x69, 0x64, 0x2c, 0x72, 0x65, 0x61, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x2c, 0x75, 0x73,
	0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x92, 0x02, 0x18, 0xe8, 0xa6, 0x81, 0xe6, 0x9b, 0xb4, 0xe6,
	0x96, 0xb0, 0xe7, 0x9a, 0x84, 0xe5, 0xad, 0x97, 0xe6, 0xae, 0xb5, 0xe5, 0x88, 0x97, 0xe8, 0xa1,
	0xa8, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x61, 0x73, 0x6b, 0x12, 0xb4, 0x01,
	0x0a, 0x0d, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x5f, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x08, 0x42, 0x89, 0x01, 0xba, 0x47, 0x85, 0x01, 0x92, 0x02, 0x81, 0x01,
	0xe5, 0xa6, 0x82, 0xe6, 0x9e, 0x9c, 0xe8, 0xae, 0xbe, 0xe7, 0xbd, 0xae, 0xe4, 0xb8, 0xba, 0x74,
	0x72, 0x75, 0x65, 0xe7, 0x9a, 0x84, 0xe6, 0x97, 0xb6, 0xe5, 0x80, 0x99, 0xef, 0xbc, 0x8c, 0xe8,
	0xb5, 0x84, 0xe6, 0xba, 0x90, 0xe4, 0xb8, 0x8d, 0xe5, 0xad, 0x98, 0xe5, 0x9c, 0xa8, 0xe5, 0x88,
	0x99, 0xe4, 0xbc, 0x9a, 0xe6, 0x96, 0xb0, 0xe5, 0xa2, 0x9e, 0x28, 0xe6, 0x8f, 0x92, 0xe5, 0x85,
	0xa5, 0x29, 0xef, 0xbc, 0x8c, 0xe5, 0xb9, 0xb6, 0xe4, 0xb8, 0x94, 0xe5, 0x9c, 0xa8, 0xe8, 0xbf,
	0x99, 0xe7, 0xa7, 0x8d, 0xe6, 0x83, 0x85, 0xe5, 0x86, 0xb5, 0xe4, 0xb8, 0x8b, 0x60, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x4d, 0x61, 0x73, 0x6b, 0x60, 0xe5, 0xad, 0x97, 0xe6, 0xae, 0xb5, 0xe5,
	0xb0, 0x86, 0xe4, 0xbc, 0x9a, 0xe8, 0xa2, 0xab, 0xe5, 0xbf, 0xbd, 0xe7, 0x95, 0xa5, 0xe3, 0x80,
	0x82, 0x48, 0x01, 0x52, 0x0c, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x4d, 0x69, 0x73, 0x73, 0x69, 0x6e,
	0x67, 0x88, 0x01, 0x01, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f,
	0x72, 0x5f, 0x69, 0x64, 0x42, 0x10, 0x0a, 0x0e, 0x5f, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x5f, 0x6d,
	0x69, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x22, 0x73, 0x0a, 0x13, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x12, 0x3c, 0x0a,
	0x0b, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0d, 0x42, 0x16, 0xba, 0x47, 0x13, 0x18, 0x01, 0x92, 0x02, 0x0e, 0xe6, 0x93, 0x8d, 0xe4,
	0xbd, 0x9c, 0xe7, 0x94, 0xa8, 0xe6, 0x88, 0xb7, 0x49, 0x44, 0x48, 0x00, 0x52, 0x0a, 0x6f, 0x70,
	0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x49, 0x64, 0x88, 0x01, 0x01, 0x42, 0x0e, 0x0a, 0x0c, 0x5f,
	0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x5f, 0x69, 0x64, 0x32, 0x9a, 0x03, 0x0a, 0x0d,
	0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4e, 0x0a,
	0x0a, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x12, 0x19, 0x2e, 0x70, 0x61,
	0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x65, 0x6e,
	0x61, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x49, 0x0a,
	0x09, 0x47, 0x65, 0x74, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x12, 0x21, 0x2e, 0x75, 0x73, 0x65,
	0x72, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74,
	0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e,
	0x75, 0x73, 0x65, 0x72, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x22, 0x00, 0x12, 0x4e, 0x0a, 0x0c, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x12, 0x24, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x4e, 0x0a, 0x0c, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x12, 0x24, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x4e, 0x0a, 0x0c, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x12, 0x24, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x42, 0xb3, 0x01, 0x0a, 0x13, 0x63, 0x6f, 0x6d,
	0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31,
	0x42, 0x0b, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a,
	0x31, 0x6b, 0x72, 0x61, 0x74, 0x6f, 0x73, 0x2d, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x76, 0x31, 0x3b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x76, 0x31, 0xa2, 0x02, 0x03, 0x55, 0x53, 0x58, 0xaa, 0x02, 0x0f, 0x55, 0x73, 0x65, 0x72, 0x2e,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x0f, 0x55, 0x73, 0x65,
	0x72, 0x5c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x1b, 0x55,
	0x73, 0x65, 0x72, 0x5c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5c, 0x56, 0x31, 0x5c, 0x47,
	0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x11, 0x55, 0x73, 0x65,
	0x72, 0x3a, 0x3a, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_user_service_v1_tenant_proto_rawDescOnce sync.Once
	file_user_service_v1_tenant_proto_rawDescData []byte
)

func file_user_service_v1_tenant_proto_rawDescGZIP() []byte {
	file_user_service_v1_tenant_proto_rawDescOnce.Do(func() {
		file_user_service_v1_tenant_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_user_service_v1_tenant_proto_rawDesc), len(file_user_service_v1_tenant_proto_rawDesc)))
	})
	return file_user_service_v1_tenant_proto_rawDescData
}

var file_user_service_v1_tenant_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_user_service_v1_tenant_proto_goTypes = []any{
	(*Tenant)(nil),                // 0: user.service.v1.Tenant
	(*ListTenantResponse)(nil),    // 1: user.service.v1.ListTenantResponse
	(*GetTenantRequest)(nil),      // 2: user.service.v1.GetTenantRequest
	(*CreateTenantRequest)(nil),   // 3: user.service.v1.CreateTenantRequest
	(*UpdateTenantRequest)(nil),   // 4: user.service.v1.UpdateTenantRequest
	(*DeleteTenantRequest)(nil),   // 5: user.service.v1.DeleteTenantRequest
	(*timestamppb.Timestamp)(nil), // 6: google.protobuf.Timestamp
	(*fieldmaskpb.FieldMask)(nil), // 7: google.protobuf.FieldMask
	(*v1.PagingRequest)(nil),      // 8: pagination.PagingRequest
	(*emptypb.Empty)(nil),         // 9: google.protobuf.Empty
}
var file_user_service_v1_tenant_proto_depIdxs = []int32{
	6,  // 0: user.service.v1.Tenant.subscription_at:type_name -> google.protobuf.Timestamp
	6,  // 1: user.service.v1.Tenant.unsubscribe_at:type_name -> google.protobuf.Timestamp
	6,  // 2: user.service.v1.Tenant.create_time:type_name -> google.protobuf.Timestamp
	6,  // 3: user.service.v1.Tenant.update_time:type_name -> google.protobuf.Timestamp
	6,  // 4: user.service.v1.Tenant.delete_time:type_name -> google.protobuf.Timestamp
	0,  // 5: user.service.v1.ListTenantResponse.items:type_name -> user.service.v1.Tenant
	0,  // 6: user.service.v1.CreateTenantRequest.data:type_name -> user.service.v1.Tenant
	0,  // 7: user.service.v1.UpdateTenantRequest.data:type_name -> user.service.v1.Tenant
	7,  // 8: user.service.v1.UpdateTenantRequest.update_mask:type_name -> google.protobuf.FieldMask
	8,  // 9: user.service.v1.TenantService.ListTenant:input_type -> pagination.PagingRequest
	2,  // 10: user.service.v1.TenantService.GetTenant:input_type -> user.service.v1.GetTenantRequest
	3,  // 11: user.service.v1.TenantService.CreateTenant:input_type -> user.service.v1.CreateTenantRequest
	4,  // 12: user.service.v1.TenantService.UpdateTenant:input_type -> user.service.v1.UpdateTenantRequest
	5,  // 13: user.service.v1.TenantService.DeleteTenant:input_type -> user.service.v1.DeleteTenantRequest
	1,  // 14: user.service.v1.TenantService.ListTenant:output_type -> user.service.v1.ListTenantResponse
	0,  // 15: user.service.v1.TenantService.GetTenant:output_type -> user.service.v1.Tenant
	9,  // 16: user.service.v1.TenantService.CreateTenant:output_type -> google.protobuf.Empty
	9,  // 17: user.service.v1.TenantService.UpdateTenant:output_type -> google.protobuf.Empty
	9,  // 18: user.service.v1.TenantService.DeleteTenant:output_type -> google.protobuf.Empty
	14, // [14:19] is the sub-list for method output_type
	9,  // [9:14] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_user_service_v1_tenant_proto_init() }
func file_user_service_v1_tenant_proto_init() {
	if File_user_service_v1_tenant_proto != nil {
		return
	}
	file_user_service_v1_tenant_proto_msgTypes[0].OneofWrappers = []any{}
	file_user_service_v1_tenant_proto_msgTypes[3].OneofWrappers = []any{}
	file_user_service_v1_tenant_proto_msgTypes[4].OneofWrappers = []any{}
	file_user_service_v1_tenant_proto_msgTypes[5].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_user_service_v1_tenant_proto_rawDesc), len(file_user_service_v1_tenant_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_user_service_v1_tenant_proto_goTypes,
		DependencyIndexes: file_user_service_v1_tenant_proto_depIdxs,
		MessageInfos:      file_user_service_v1_tenant_proto_msgTypes,
	}.Build()
	File_user_service_v1_tenant_proto = out.File
	file_user_service_v1_tenant_proto_goTypes = nil
	file_user_service_v1_tenant_proto_depIdxs = nil
}
