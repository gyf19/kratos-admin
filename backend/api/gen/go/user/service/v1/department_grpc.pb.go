// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             (unknown)
// source: user/service/v1/department.proto

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
	DepartmentService_ListDepartment_FullMethodName   = "/user.service.v1.DepartmentService/ListDepartment"
	DepartmentService_GetDepartment_FullMethodName    = "/user.service.v1.DepartmentService/GetDepartment"
	DepartmentService_CreateDepartment_FullMethodName = "/user.service.v1.DepartmentService/CreateDepartment"
	DepartmentService_UpdateDepartment_FullMethodName = "/user.service.v1.DepartmentService/UpdateDepartment"
	DepartmentService_DeleteDepartment_FullMethodName = "/user.service.v1.DepartmentService/DeleteDepartment"
)

// DepartmentServiceClient is the client API for DepartmentService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// 部门服务
type DepartmentServiceClient interface {
	// 查询部门列表
	ListDepartment(ctx context.Context, in *v1.PagingRequest, opts ...grpc.CallOption) (*ListDepartmentResponse, error)
	// 查询部门详情
	GetDepartment(ctx context.Context, in *GetDepartmentRequest, opts ...grpc.CallOption) (*Department, error)
	// 创建部门
	CreateDepartment(ctx context.Context, in *CreateDepartmentRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// 更新部门
	UpdateDepartment(ctx context.Context, in *UpdateDepartmentRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// 删除部门
	DeleteDepartment(ctx context.Context, in *DeleteDepartmentRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type departmentServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDepartmentServiceClient(cc grpc.ClientConnInterface) DepartmentServiceClient {
	return &departmentServiceClient{cc}
}

func (c *departmentServiceClient) ListDepartment(ctx context.Context, in *v1.PagingRequest, opts ...grpc.CallOption) (*ListDepartmentResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListDepartmentResponse)
	err := c.cc.Invoke(ctx, DepartmentService_ListDepartment_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *departmentServiceClient) GetDepartment(ctx context.Context, in *GetDepartmentRequest, opts ...grpc.CallOption) (*Department, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Department)
	err := c.cc.Invoke(ctx, DepartmentService_GetDepartment_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *departmentServiceClient) CreateDepartment(ctx context.Context, in *CreateDepartmentRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, DepartmentService_CreateDepartment_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *departmentServiceClient) UpdateDepartment(ctx context.Context, in *UpdateDepartmentRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, DepartmentService_UpdateDepartment_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *departmentServiceClient) DeleteDepartment(ctx context.Context, in *DeleteDepartmentRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, DepartmentService_DeleteDepartment_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DepartmentServiceServer is the server API for DepartmentService service.
// All implementations must embed UnimplementedDepartmentServiceServer
// for forward compatibility.
//
// 部门服务
type DepartmentServiceServer interface {
	// 查询部门列表
	ListDepartment(context.Context, *v1.PagingRequest) (*ListDepartmentResponse, error)
	// 查询部门详情
	GetDepartment(context.Context, *GetDepartmentRequest) (*Department, error)
	// 创建部门
	CreateDepartment(context.Context, *CreateDepartmentRequest) (*emptypb.Empty, error)
	// 更新部门
	UpdateDepartment(context.Context, *UpdateDepartmentRequest) (*emptypb.Empty, error)
	// 删除部门
	DeleteDepartment(context.Context, *DeleteDepartmentRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedDepartmentServiceServer()
}

// UnimplementedDepartmentServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedDepartmentServiceServer struct{}

func (UnimplementedDepartmentServiceServer) ListDepartment(context.Context, *v1.PagingRequest) (*ListDepartmentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListDepartment not implemented")
}
func (UnimplementedDepartmentServiceServer) GetDepartment(context.Context, *GetDepartmentRequest) (*Department, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDepartment not implemented")
}
func (UnimplementedDepartmentServiceServer) CreateDepartment(context.Context, *CreateDepartmentRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateDepartment not implemented")
}
func (UnimplementedDepartmentServiceServer) UpdateDepartment(context.Context, *UpdateDepartmentRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateDepartment not implemented")
}
func (UnimplementedDepartmentServiceServer) DeleteDepartment(context.Context, *DeleteDepartmentRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteDepartment not implemented")
}
func (UnimplementedDepartmentServiceServer) mustEmbedUnimplementedDepartmentServiceServer() {}
func (UnimplementedDepartmentServiceServer) testEmbeddedByValue()                           {}

// UnsafeDepartmentServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DepartmentServiceServer will
// result in compilation errors.
type UnsafeDepartmentServiceServer interface {
	mustEmbedUnimplementedDepartmentServiceServer()
}

func RegisterDepartmentServiceServer(s grpc.ServiceRegistrar, srv DepartmentServiceServer) {
	// If the following call pancis, it indicates UnimplementedDepartmentServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&DepartmentService_ServiceDesc, srv)
}

func _DepartmentService_ListDepartment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v1.PagingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DepartmentServiceServer).ListDepartment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DepartmentService_ListDepartment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DepartmentServiceServer).ListDepartment(ctx, req.(*v1.PagingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DepartmentService_GetDepartment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDepartmentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DepartmentServiceServer).GetDepartment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DepartmentService_GetDepartment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DepartmentServiceServer).GetDepartment(ctx, req.(*GetDepartmentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DepartmentService_CreateDepartment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateDepartmentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DepartmentServiceServer).CreateDepartment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DepartmentService_CreateDepartment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DepartmentServiceServer).CreateDepartment(ctx, req.(*CreateDepartmentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DepartmentService_UpdateDepartment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateDepartmentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DepartmentServiceServer).UpdateDepartment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DepartmentService_UpdateDepartment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DepartmentServiceServer).UpdateDepartment(ctx, req.(*UpdateDepartmentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DepartmentService_DeleteDepartment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteDepartmentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DepartmentServiceServer).DeleteDepartment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DepartmentService_DeleteDepartment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DepartmentServiceServer).DeleteDepartment(ctx, req.(*DeleteDepartmentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DepartmentService_ServiceDesc is the grpc.ServiceDesc for DepartmentService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DepartmentService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "user.service.v1.DepartmentService",
	HandlerType: (*DepartmentServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListDepartment",
			Handler:    _DepartmentService_ListDepartment_Handler,
		},
		{
			MethodName: "GetDepartment",
			Handler:    _DepartmentService_GetDepartment_Handler,
		},
		{
			MethodName: "CreateDepartment",
			Handler:    _DepartmentService_CreateDepartment_Handler,
		},
		{
			MethodName: "UpdateDepartment",
			Handler:    _DepartmentService_UpdateDepartment_Handler,
		},
		{
			MethodName: "DeleteDepartment",
			Handler:    _DepartmentService_DeleteDepartment_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user/service/v1/department.proto",
}
