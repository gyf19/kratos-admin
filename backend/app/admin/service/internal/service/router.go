package service

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/tx7do/go-utils/trans"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/fieldmaskpb"

	"kratos-admin/app/admin/service/internal/data"
	"kratos-admin/app/admin/service/internal/middleware/auth"

	pagination "github.com/tx7do/kratos-bootstrap/api/gen/go/pagination/v1"
	adminV1 "kratos-admin/api/gen/go/admin/service/v1"
	systemV1 "kratos-admin/api/gen/go/system/service/v1"
)

type RouterService struct {
	adminV1.RouterServiceHTTPServer

	log *log.Helper

	menuRepo *data.MenuRepo
	roleRepo *data.RoleRepo
	userRepo *data.UserRepo
}

func NewRouterService(
	logger log.Logger,
	menuRepo *data.MenuRepo,
	roleRepo *data.RoleRepo,
	userRepo *data.UserRepo,
) *RouterService {
	l := log.NewHelper(log.With(logger, "module", "router/service/admin-service"))
	return &RouterService{
		log:      l,
		menuRepo: menuRepo,
		roleRepo: roleRepo,
		userRepo: userRepo,
	}
}

func (s *RouterService) menuListToQueryString(menus []uint32, onlyButton bool) string {
	var ids []string
	for _, menu := range menus {
		ids = append(ids, fmt.Sprintf("\"%d\"", menu))
	}
	idsStr := fmt.Sprintf("[%s]", strings.Join(ids, ", "))
	query := map[string]string{"id__in": idsStr}

	if onlyButton {
		query["type"] = systemV1.MenuType_BUTTON.String()
	} else {
		query["type__not"] = systemV1.MenuType_BUTTON.String()
	}

	query["status"] = "ON"

	queryStr, err := json.Marshal(query)
	if err != nil {
		return ""
	}

	return string(queryStr)
}

func (s *RouterService) queryRoleMenus(ctx context.Context, userId uint32) ([]uint32, error) {
	user, err := s.userRepo.GetUser(ctx, userId)
	if err != nil {
		s.log.Errorf("查询用户失败[%s]", err.Error())
		return nil, errors.New("查询用户失败")
	}

	role, err := s.roleRepo.GetRole(ctx, user.GetRoleId())
	if err != nil {
		s.log.Errorf("查询角色失败[%s]", err.Error())
		return nil, errors.New("查询角色失败")
	}

	return role.GetMenus(), nil
}

func (s *RouterService) ListPermissionCode(ctx context.Context, _ *emptypb.Empty) (*adminV1.ListPermissionCodeResponse, error) {
	authInfo, err := auth.FromContext(ctx)
	if err != nil {
		s.log.Errorf("用户认证失败[%s]", err.Error())
		return nil, adminV1.ErrorAccessForbidden("用户认证失败")
	}

	roleMenus, err := s.queryRoleMenus(ctx, authInfo.UserId)
	if err != nil {
		return nil, err
	}

	menus, err := s.menuRepo.List(ctx, &pagination.PagingRequest{
		NoPaging: trans.Ptr(true),
		Query:    trans.Ptr(s.menuListToQueryString(roleMenus, true)),
		FieldMask: &fieldmaskpb.FieldMask{
			Paths: []string{"id", "meta"},
		},
	}, false)
	if err != nil {
		s.log.Errorf("查询列表发生错误[%s]", err.Error())
		return nil, errors.New("读取列表发生错误")
	}

	var codes []string
	for menu := range menus.Items {
		if menus.Items[menu].GetMeta() == nil {
			continue
		}
		if len(menus.Items[menu].GetMeta().GetAuthority()) == 0 {
			continue
		}

		codes = append(codes, menus.Items[menu].GetMeta().GetAuthority()...)
	}

	return &adminV1.ListPermissionCodeResponse{
		Codes: codes,
	}, nil
}

func (s *RouterService) fillRouteItem(menus []*systemV1.Menu) []*systemV1.RouteItem {
	if len(menus) == 0 {
		return nil
	}

	var routers []*systemV1.RouteItem

	for _, v := range menus {
		if v.GetStatus() != systemV1.MenuStatus_ON {
			continue
		}
		if v.GetType() == systemV1.MenuType_BUTTON {
			continue
		}

		item := &systemV1.RouteItem{
			Path:      v.Path,
			Component: v.Component,
			Name:      v.Name,
			Redirect:  v.Redirect,
			Alias:     v.Alias,
			Meta:      v.Meta,
		}

		if len(v.Children) > 0 {
			item.Children = s.fillRouteItem(v.Children)
		}

		routers = append(routers, item)
	}

	return routers
}

func (s *RouterService) ListRoute(ctx context.Context, _ *emptypb.Empty) (*adminV1.ListRouteResponse, error) {
	authInfo, err := auth.FromContext(ctx)
	if err != nil {
		s.log.Errorf("用户认证失败[%s]", err.Error())
		return nil, adminV1.ErrorAccessForbidden("用户认证失败")
	}

	roleMenus, err := s.queryRoleMenus(ctx, authInfo.UserId)
	if err != nil {
		return nil, err
	}

	menuList, err := s.menuRepo.List(ctx, &pagination.PagingRequest{
		NoPaging: trans.Ptr(true),
		Query:    trans.Ptr(s.menuListToQueryString(roleMenus, false)),
	}, true)
	if err != nil {
		s.log.Errorf("查询列表发生错误[%s]", err.Error())
		return nil, errors.New("读取列表发生错误")
	}

	resp := &adminV1.ListRouteResponse{Items: s.fillRouteItem(menuList.Items)}

	return resp, nil
}
