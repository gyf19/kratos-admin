package service

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	pagination "github.com/tx7do/kratos-bootstrap/api/gen/go/pagination/v1"
	"google.golang.org/protobuf/types/known/emptypb"

	"kratos-admin/app/admin/service/internal/data"

	adminV1 "kratos-admin/api/gen/go/admin/service/v1"
	internalMessageV1 "kratos-admin/api/gen/go/internal_message/service/v1"

	"kratos-admin/pkg/middleware/auth"
)

type NotificationMessageCategoryService struct {
	adminV1.NotificationMessageCategoryServiceHTTPServer

	log *log.Helper

	uc *data.NotificationMessageCategoryRepo
}

func NewNotificationMessageCategoryService(uc *data.NotificationMessageCategoryRepo, logger log.Logger) *NotificationMessageCategoryService {
	l := log.NewHelper(log.With(logger, "module", "notification-message-category/service/admin-service"))
	return &NotificationMessageCategoryService{
		log: l,
		uc:  uc,
	}
}

func (s *NotificationMessageCategoryService) List(ctx context.Context, req *pagination.PagingRequest) (*internalMessageV1.ListNotificationMessageCategoryResponse, error) {
	return s.uc.List(ctx, req)
}

func (s *NotificationMessageCategoryService) Get(ctx context.Context, req *internalMessageV1.GetNotificationMessageCategoryRequest) (*internalMessageV1.NotificationMessageCategory, error) {
	return s.uc.Get(ctx, req)
}

func (s *NotificationMessageCategoryService) Create(ctx context.Context, req *internalMessageV1.CreateNotificationMessageCategoryRequest) (*emptypb.Empty, error) {
	if req.Data == nil {
		return nil, adminV1.ErrorBadRequest("错误的参数")
	}

	// 获取操作人信息
	operator, err := auth.FromContext(ctx)
	if err != nil {
		return &emptypb.Empty{}, err
	}

	if err = s.uc.Create(ctx, req, operator); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (s *NotificationMessageCategoryService) Update(ctx context.Context, req *internalMessageV1.UpdateNotificationMessageCategoryRequest) (*emptypb.Empty, error) {
	if req.Data == nil {
		return nil, adminV1.ErrorBadRequest("错误的参数")
	}

	// 获取操作人信息
	operator, err := auth.FromContext(ctx)
	if err != nil {
		return &emptypb.Empty{}, err
	}

	if err = s.uc.Update(ctx, req, operator); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (s *NotificationMessageCategoryService) Delete(ctx context.Context, req *internalMessageV1.DeleteNotificationMessageCategoryRequest) (*emptypb.Empty, error) {
	if _, err := s.uc.Delete(ctx, req); err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}
