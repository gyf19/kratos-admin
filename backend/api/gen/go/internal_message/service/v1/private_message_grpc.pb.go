// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             (unknown)
// source: internal_message/service/v1/private_message.proto

package servicev1

import (
	context "context"
	v1 "github.com/tx7do/kratos-bootstrap/api/gen/go/pagination/v1"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	PrivateMessageService_ListPrivateMessage_FullMethodName   = "/internal_message.service.v1.PrivateMessageService/ListPrivateMessage"
	PrivateMessageService_GetPrivateMessage_FullMethodName    = "/internal_message.service.v1.PrivateMessageService/GetPrivateMessage"
	PrivateMessageService_CreatePrivateMessage_FullMethodName = "/internal_message.service.v1.PrivateMessageService/CreatePrivateMessage"
	PrivateMessageService_UpdatePrivateMessage_FullMethodName = "/internal_message.service.v1.PrivateMessageService/UpdatePrivateMessage"
	PrivateMessageService_DeletePrivateMessage_FullMethodName = "/internal_message.service.v1.PrivateMessageService/DeletePrivateMessage"
)

// PrivateMessageServiceClient is the client API for PrivateMessageService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// 私信消息服务
type PrivateMessageServiceClient interface {
	// 查询私信消息列表
	ListPrivateMessage(ctx context.Context, in *v1.PagingRequest, opts ...grpc.CallOption) (*ListPrivateMessageResponse, error)
	// 查询私信消息详情
	GetPrivateMessage(ctx context.Context, in *GetPrivateMessageRequest, opts ...grpc.CallOption) (*PrivateMessage, error)
	// 创建私信消息
	CreatePrivateMessage(ctx context.Context, in *CreatePrivateMessageRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// 更新私信消息
	UpdatePrivateMessage(ctx context.Context, in *UpdatePrivateMessageRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// 删除私信消息
	DeletePrivateMessage(ctx context.Context, in *DeletePrivateMessageRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type privateMessageServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPrivateMessageServiceClient(cc grpc.ClientConnInterface) PrivateMessageServiceClient {
	return &privateMessageServiceClient{cc}
}

func (c *privateMessageServiceClient) ListPrivateMessage(ctx context.Context, in *v1.PagingRequest, opts ...grpc.CallOption) (*ListPrivateMessageResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListPrivateMessageResponse)
	err := c.cc.Invoke(ctx, PrivateMessageService_ListPrivateMessage_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *privateMessageServiceClient) GetPrivateMessage(ctx context.Context, in *GetPrivateMessageRequest, opts ...grpc.CallOption) (*PrivateMessage, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PrivateMessage)
	err := c.cc.Invoke(ctx, PrivateMessageService_GetPrivateMessage_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *privateMessageServiceClient) CreatePrivateMessage(ctx context.Context, in *CreatePrivateMessageRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, PrivateMessageService_CreatePrivateMessage_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *privateMessageServiceClient) UpdatePrivateMessage(ctx context.Context, in *UpdatePrivateMessageRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, PrivateMessageService_UpdatePrivateMessage_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *privateMessageServiceClient) DeletePrivateMessage(ctx context.Context, in *DeletePrivateMessageRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, PrivateMessageService_DeletePrivateMessage_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PrivateMessageServiceServer is the server API for PrivateMessageService service.
// All implementations must embed UnimplementedPrivateMessageServiceServer
// for forward compatibility.
//
// 私信消息服务
type PrivateMessageServiceServer interface {
	// 查询私信消息列表
	ListPrivateMessage(context.Context, *v1.PagingRequest) (*ListPrivateMessageResponse, error)
	// 查询私信消息详情
	GetPrivateMessage(context.Context, *GetPrivateMessageRequest) (*PrivateMessage, error)
	// 创建私信消息
	CreatePrivateMessage(context.Context, *CreatePrivateMessageRequest) (*emptypb.Empty, error)
	// 更新私信消息
	UpdatePrivateMessage(context.Context, *UpdatePrivateMessageRequest) (*emptypb.Empty, error)
	// 删除私信消息
	DeletePrivateMessage(context.Context, *DeletePrivateMessageRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedPrivateMessageServiceServer()
}

// UnimplementedPrivateMessageServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedPrivateMessageServiceServer struct{}

func (UnimplementedPrivateMessageServiceServer) ListPrivateMessage(context.Context, *v1.PagingRequest) (*ListPrivateMessageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListPrivateMessage not implemented")
}
func (UnimplementedPrivateMessageServiceServer) GetPrivateMessage(context.Context, *GetPrivateMessageRequest) (*PrivateMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPrivateMessage not implemented")
}
func (UnimplementedPrivateMessageServiceServer) CreatePrivateMessage(context.Context, *CreatePrivateMessageRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePrivateMessage not implemented")
}
func (UnimplementedPrivateMessageServiceServer) UpdatePrivateMessage(context.Context, *UpdatePrivateMessageRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePrivateMessage not implemented")
}
func (UnimplementedPrivateMessageServiceServer) DeletePrivateMessage(context.Context, *DeletePrivateMessageRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeletePrivateMessage not implemented")
}
func (UnimplementedPrivateMessageServiceServer) mustEmbedUnimplementedPrivateMessageServiceServer() {}
func (UnimplementedPrivateMessageServiceServer) testEmbeddedByValue()                               {}

// UnsafePrivateMessageServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PrivateMessageServiceServer will
// result in compilation errors.
type UnsafePrivateMessageServiceServer interface {
	mustEmbedUnimplementedPrivateMessageServiceServer()
}

func RegisterPrivateMessageServiceServer(s grpc.ServiceRegistrar, srv PrivateMessageServiceServer) {
	// If the following call pancis, it indicates UnimplementedPrivateMessageServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&PrivateMessageService_ServiceDesc, srv)
}

func _PrivateMessageService_ListPrivateMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v1.PagingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PrivateMessageServiceServer).ListPrivateMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PrivateMessageService_ListPrivateMessage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PrivateMessageServiceServer).ListPrivateMessage(ctx, req.(*v1.PagingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PrivateMessageService_GetPrivateMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPrivateMessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PrivateMessageServiceServer).GetPrivateMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PrivateMessageService_GetPrivateMessage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PrivateMessageServiceServer).GetPrivateMessage(ctx, req.(*GetPrivateMessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PrivateMessageService_CreatePrivateMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePrivateMessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PrivateMessageServiceServer).CreatePrivateMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PrivateMessageService_CreatePrivateMessage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PrivateMessageServiceServer).CreatePrivateMessage(ctx, req.(*CreatePrivateMessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PrivateMessageService_UpdatePrivateMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePrivateMessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PrivateMessageServiceServer).UpdatePrivateMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PrivateMessageService_UpdatePrivateMessage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PrivateMessageServiceServer).UpdatePrivateMessage(ctx, req.(*UpdatePrivateMessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PrivateMessageService_DeletePrivateMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeletePrivateMessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PrivateMessageServiceServer).DeletePrivateMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PrivateMessageService_DeletePrivateMessage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PrivateMessageServiceServer).DeletePrivateMessage(ctx, req.(*DeletePrivateMessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PrivateMessageService_ServiceDesc is the grpc.ServiceDesc for PrivateMessageService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PrivateMessageService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "internal_message.service.v1.PrivateMessageService",
	HandlerType: (*PrivateMessageServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListPrivateMessage",
			Handler:    _PrivateMessageService_ListPrivateMessage_Handler,
		},
		{
			MethodName: "GetPrivateMessage",
			Handler:    _PrivateMessageService_GetPrivateMessage_Handler,
		},
		{
			MethodName: "CreatePrivateMessage",
			Handler:    _PrivateMessageService_CreatePrivateMessage_Handler,
		},
		{
			MethodName: "UpdatePrivateMessage",
			Handler:    _PrivateMessageService_UpdatePrivateMessage_Handler,
		},
		{
			MethodName: "DeletePrivateMessage",
			Handler:    _PrivateMessageService_DeletePrivateMessage_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "internal_message/service/v1/private_message.proto",
}
