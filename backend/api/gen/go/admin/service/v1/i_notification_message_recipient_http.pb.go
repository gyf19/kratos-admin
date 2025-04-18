// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.8.3
// - protoc             (unknown)
// source: admin/service/v1/i_notification_message_recipient.proto

package servicev1

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
	v1 "github.com/tx7do/kratos-bootstrap/api/gen/go/pagination/v1"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	v11 "kratos-admin/api/gen/go/internal_message/service/v1"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

const OperationNotificationMessageRecipientServiceCreateNotificationMessageRecipient = "/admin.service.v1.NotificationMessageRecipientService/CreateNotificationMessageRecipient"
const OperationNotificationMessageRecipientServiceDeleteNotificationMessageRecipient = "/admin.service.v1.NotificationMessageRecipientService/DeleteNotificationMessageRecipient"
const OperationNotificationMessageRecipientServiceGetNotificationMessageRecipient = "/admin.service.v1.NotificationMessageRecipientService/GetNotificationMessageRecipient"
const OperationNotificationMessageRecipientServiceListNotificationMessageRecipient = "/admin.service.v1.NotificationMessageRecipientService/ListNotificationMessageRecipient"
const OperationNotificationMessageRecipientServiceUpdateNotificationMessageRecipient = "/admin.service.v1.NotificationMessageRecipientService/UpdateNotificationMessageRecipient"

type NotificationMessageRecipientServiceHTTPServer interface {
	// CreateNotificationMessageRecipient 创建通知消息接收者
	CreateNotificationMessageRecipient(context.Context, *v11.CreateNotificationMessageRecipientRequest) (*emptypb.Empty, error)
	// DeleteNotificationMessageRecipient 删除通知消息接收者
	DeleteNotificationMessageRecipient(context.Context, *v11.DeleteNotificationMessageRecipientRequest) (*emptypb.Empty, error)
	// GetNotificationMessageRecipient 查询通知消息接收者详情
	GetNotificationMessageRecipient(context.Context, *v11.GetNotificationMessageRecipientRequest) (*v11.NotificationMessageRecipient, error)
	// ListNotificationMessageRecipient 查询通知消息接收者列表
	ListNotificationMessageRecipient(context.Context, *v1.PagingRequest) (*v11.ListNotificationMessageRecipientResponse, error)
	// UpdateNotificationMessageRecipient 更新通知消息接收者
	UpdateNotificationMessageRecipient(context.Context, *v11.UpdateNotificationMessageRecipientRequest) (*emptypb.Empty, error)
}

func RegisterNotificationMessageRecipientServiceHTTPServer(s *http.Server, srv NotificationMessageRecipientServiceHTTPServer) {
	r := s.Route("/")
	r.GET("/admin/v1/notifications:recipients", _NotificationMessageRecipientService_ListNotificationMessageRecipient0_HTTP_Handler(srv))
	r.GET("/admin/v1/notifications:recipients/{id}", _NotificationMessageRecipientService_GetNotificationMessageRecipient0_HTTP_Handler(srv))
	r.POST("/admin/v1/notifications:recipients", _NotificationMessageRecipientService_CreateNotificationMessageRecipient0_HTTP_Handler(srv))
	r.PUT("/admin/v1/notifications:recipients/{data.id}", _NotificationMessageRecipientService_UpdateNotificationMessageRecipient0_HTTP_Handler(srv))
	r.DELETE("/admin/v1/notifications:recipients/{id}", _NotificationMessageRecipientService_DeleteNotificationMessageRecipient0_HTTP_Handler(srv))
}

func _NotificationMessageRecipientService_ListNotificationMessageRecipient0_HTTP_Handler(srv NotificationMessageRecipientServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in v1.PagingRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationNotificationMessageRecipientServiceListNotificationMessageRecipient)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListNotificationMessageRecipient(ctx, req.(*v1.PagingRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*v11.ListNotificationMessageRecipientResponse)
		return ctx.Result(200, reply)
	}
}

func _NotificationMessageRecipientService_GetNotificationMessageRecipient0_HTTP_Handler(srv NotificationMessageRecipientServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in v11.GetNotificationMessageRecipientRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationNotificationMessageRecipientServiceGetNotificationMessageRecipient)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetNotificationMessageRecipient(ctx, req.(*v11.GetNotificationMessageRecipientRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*v11.NotificationMessageRecipient)
		return ctx.Result(200, reply)
	}
}

func _NotificationMessageRecipientService_CreateNotificationMessageRecipient0_HTTP_Handler(srv NotificationMessageRecipientServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in v11.CreateNotificationMessageRecipientRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationNotificationMessageRecipientServiceCreateNotificationMessageRecipient)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CreateNotificationMessageRecipient(ctx, req.(*v11.CreateNotificationMessageRecipientRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*emptypb.Empty)
		return ctx.Result(200, reply)
	}
}

func _NotificationMessageRecipientService_UpdateNotificationMessageRecipient0_HTTP_Handler(srv NotificationMessageRecipientServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in v11.UpdateNotificationMessageRecipientRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationNotificationMessageRecipientServiceUpdateNotificationMessageRecipient)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UpdateNotificationMessageRecipient(ctx, req.(*v11.UpdateNotificationMessageRecipientRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*emptypb.Empty)
		return ctx.Result(200, reply)
	}
}

func _NotificationMessageRecipientService_DeleteNotificationMessageRecipient0_HTTP_Handler(srv NotificationMessageRecipientServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in v11.DeleteNotificationMessageRecipientRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationNotificationMessageRecipientServiceDeleteNotificationMessageRecipient)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.DeleteNotificationMessageRecipient(ctx, req.(*v11.DeleteNotificationMessageRecipientRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*emptypb.Empty)
		return ctx.Result(200, reply)
	}
}

type NotificationMessageRecipientServiceHTTPClient interface {
	CreateNotificationMessageRecipient(ctx context.Context, req *v11.CreateNotificationMessageRecipientRequest, opts ...http.CallOption) (rsp *emptypb.Empty, err error)
	DeleteNotificationMessageRecipient(ctx context.Context, req *v11.DeleteNotificationMessageRecipientRequest, opts ...http.CallOption) (rsp *emptypb.Empty, err error)
	GetNotificationMessageRecipient(ctx context.Context, req *v11.GetNotificationMessageRecipientRequest, opts ...http.CallOption) (rsp *v11.NotificationMessageRecipient, err error)
	ListNotificationMessageRecipient(ctx context.Context, req *v1.PagingRequest, opts ...http.CallOption) (rsp *v11.ListNotificationMessageRecipientResponse, err error)
	UpdateNotificationMessageRecipient(ctx context.Context, req *v11.UpdateNotificationMessageRecipientRequest, opts ...http.CallOption) (rsp *emptypb.Empty, err error)
}

type NotificationMessageRecipientServiceHTTPClientImpl struct {
	cc *http.Client
}

func NewNotificationMessageRecipientServiceHTTPClient(client *http.Client) NotificationMessageRecipientServiceHTTPClient {
	return &NotificationMessageRecipientServiceHTTPClientImpl{client}
}

func (c *NotificationMessageRecipientServiceHTTPClientImpl) CreateNotificationMessageRecipient(ctx context.Context, in *v11.CreateNotificationMessageRecipientRequest, opts ...http.CallOption) (*emptypb.Empty, error) {
	var out emptypb.Empty
	pattern := "/admin/v1/notifications:recipients"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationNotificationMessageRecipientServiceCreateNotificationMessageRecipient))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *NotificationMessageRecipientServiceHTTPClientImpl) DeleteNotificationMessageRecipient(ctx context.Context, in *v11.DeleteNotificationMessageRecipientRequest, opts ...http.CallOption) (*emptypb.Empty, error) {
	var out emptypb.Empty
	pattern := "/admin/v1/notifications:recipients/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationNotificationMessageRecipientServiceDeleteNotificationMessageRecipient))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *NotificationMessageRecipientServiceHTTPClientImpl) GetNotificationMessageRecipient(ctx context.Context, in *v11.GetNotificationMessageRecipientRequest, opts ...http.CallOption) (*v11.NotificationMessageRecipient, error) {
	var out v11.NotificationMessageRecipient
	pattern := "/admin/v1/notifications:recipients/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationNotificationMessageRecipientServiceGetNotificationMessageRecipient))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *NotificationMessageRecipientServiceHTTPClientImpl) ListNotificationMessageRecipient(ctx context.Context, in *v1.PagingRequest, opts ...http.CallOption) (*v11.ListNotificationMessageRecipientResponse, error) {
	var out v11.ListNotificationMessageRecipientResponse
	pattern := "/admin/v1/notifications:recipients"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationNotificationMessageRecipientServiceListNotificationMessageRecipient))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *NotificationMessageRecipientServiceHTTPClientImpl) UpdateNotificationMessageRecipient(ctx context.Context, in *v11.UpdateNotificationMessageRecipientRequest, opts ...http.CallOption) (*emptypb.Empty, error) {
	var out emptypb.Empty
	pattern := "/admin/v1/notifications:recipients/{data.id}"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationNotificationMessageRecipientServiceUpdateNotificationMessageRecipient))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}
