package data

import (
	"context"
	"errors"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/go-kratos/kratos/v2/log"
	entgo "github.com/tx7do/go-utils/entgo/query"
	entgoUpdate "github.com/tx7do/go-utils/entgo/update"
	"github.com/tx7do/go-utils/fieldmaskutil"
	"github.com/tx7do/go-utils/timeutil"
	"github.com/tx7do/go-utils/trans"
	pagination "github.com/tx7do/kratos-bootstrap/api/gen/go/pagination/v1"

	"kratos-admin/app/admin/service/internal/data/ent"
	"kratos-admin/app/admin/service/internal/data/ent/adminoperationlog"

	systemV1 "kratos-admin/api/gen/go/system/service/v1"
)

type AdminOperationLogRepo struct {
	data *Data
	log  *log.Helper
}

func NewAdminOperationLogRepo(data *Data, logger log.Logger) *AdminOperationLogRepo {
	l := log.NewHelper(log.With(logger, "module", "admin-operation-log/repo/admin-service"))
	return &AdminOperationLogRepo{
		data: data,
		log:  l,
	}
}

func (r *AdminOperationLogRepo) convertEntToProto(in *ent.AdminOperationLog) *systemV1.AdminOperationLog {
	if in == nil {
		return nil
	}
	return &systemV1.AdminOperationLog{
		Id:             trans.Ptr(in.ID),
		RequestId:      in.RequestID,
		Method:         in.Method,
		Operation:      in.Operation,
		Path:           in.Path,
		Referer:        in.Referer,
		RequestUri:     in.RequestURI,
		RequestBody:    in.RequestBody,
		RequestHeader:  in.RequestHeader,
		Response:       in.Response,
		CostTime:       timeutil.SecondToDurationpb(in.CostTime),
		UserId:         in.UserID,
		UserName:       in.UserName,
		ClientIp:       in.ClientIP,
		UserAgent:      in.UserAgent,
		BrowserName:    in.BrowserName,
		BrowserVersion: in.BrowserVersion,
		ClientId:       in.ClientID,
		ClientName:     in.ClientName,
		OsName:         in.OsName,
		OsVersion:      in.OsVersion,
		Success:        in.Success,
		StatusCode:     in.StatusCode,
		Reason:         in.Reason,
		Location:       in.Location,
		CreateTime:     timeutil.TimeToTimestamppb(in.CreateTime),
		UpdateTime:     timeutil.TimeToTimestamppb(in.UpdateTime),
		DeleteTime:     timeutil.TimeToTimestamppb(in.DeleteTime),
	}
}

func (r *AdminOperationLogRepo) Count(ctx context.Context, whereCond []func(s *sql.Selector)) (int, error) {
	builder := r.data.db.Client().AdminOperationLog.Query()
	if len(whereCond) != 0 {
		builder.Modify(whereCond...)
	}

	count, err := builder.Count(ctx)
	if err != nil {
		r.log.Errorf("query count failed: %s", err.Error())
	}

	return count, err
}

func (r *AdminOperationLogRepo) List(ctx context.Context, req *pagination.PagingRequest) (*systemV1.ListAdminOperationLogResponse, error) {
	builder := r.data.db.Client().AdminOperationLog.Query()

	err, whereSelectors, querySelectors := entgo.BuildQuerySelector(
		req.GetQuery(), req.GetOrQuery(),
		req.GetPage(), req.GetPageSize(), req.GetNoPaging(),
		req.GetOrderBy(), adminoperationlog.FieldCreateTime,
		req.GetFieldMask().GetPaths(),
	)
	if err != nil {
		r.log.Errorf("解析条件发生错误[%s]", err.Error())
		return nil, err
	}

	if querySelectors != nil {
		builder.Modify(querySelectors...)
	}

	results, err := builder.All(ctx)
	if err != nil {
		return nil, err
	}

	items := make([]*systemV1.AdminOperationLog, 0, len(results))
	for _, res := range results {
		item := r.convertEntToProto(res)
		items = append(items, item)
	}

	count, err := r.Count(ctx, whereSelectors)
	if err != nil {
		return nil, err
	}

	return &systemV1.ListAdminOperationLogResponse{
		Total: uint32(count),
		Items: items,
	}, err
}

func (r *AdminOperationLogRepo) IsExist(ctx context.Context, id uint32) (bool, error) {
	return r.data.db.Client().AdminOperationLog.Query().
		Where(adminoperationlog.IDEQ(id)).
		Exist(ctx)
}

func (r *AdminOperationLogRepo) Get(ctx context.Context, req *systemV1.GetAdminOperationLogRequest) (*systemV1.AdminOperationLog, error) {
	ret, err := r.data.db.Client().AdminOperationLog.Get(ctx, req.GetId())
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, systemV1.ErrorResourceNotFound("admin operation log not found")
		}
		return nil, err
	}

	return r.convertEntToProto(ret), err
}

func (r *AdminOperationLogRepo) Create(ctx context.Context, req *systemV1.CreateAdminOperationLogRequest) error {
	if req.Data == nil {
		return errors.New("invalid request")
	}

	builder := r.data.db.Client().AdminOperationLog.
		Create().
		SetNillableRequestID(req.Data.RequestId).
		SetNillableMethod(req.Data.Method).
		SetNillableOperation(req.Data.Operation).
		SetNillablePath(req.Data.Path).
		SetNillableReferer(req.Data.Referer).
		SetNillableRequestURI(req.Data.RequestUri).
		SetNillableRequestBody(req.Data.RequestBody).
		SetNillableRequestHeader(req.Data.RequestHeader).
		SetNillableResponse(req.Data.Response).
		SetNillableCostTime(timeutil.DurationpbSecond(req.Data.CostTime)).
		SetNillableUserID(req.Data.UserId).
		SetNillableUserName(req.Data.UserName).
		SetNillableClientIP(req.Data.ClientIp).
		SetNillableUserAgent(req.Data.UserAgent).
		SetNillableBrowserName(req.Data.BrowserName).
		SetNillableBrowserVersion(req.Data.BrowserVersion).
		SetNillableClientID(req.Data.ClientId).
		SetNillableClientName(req.Data.ClientName).
		SetNillableOsName(req.Data.OsName).
		SetNillableOsVersion(req.Data.OsVersion).
		SetNillableStatusCode(req.Data.StatusCode).
		SetNillableSuccess(req.Data.Success).
		SetNillableReason(req.Data.Reason).
		SetNillableLocation(req.Data.Location).
		SetNillableCreateTime(timeutil.TimestamppbToTime(req.Data.CreateTime))

	if req.Data.CreateTime == nil {
		builder.SetCreateTime(time.Now())
	}

	err := builder.Exec(ctx)
	if err != nil {
		r.log.Errorf("insert one data failed: %s", err.Error())
		return err
	}

	return err
}

func (r *AdminOperationLogRepo) Update(ctx context.Context, req *systemV1.UpdateAdminOperationLogRequest) error {
	if req.Data == nil {
		return errors.New("invalid request")
	}

	// 如果不存在则创建
	if req.GetAllowMissing() {
		exist, err := r.IsExist(ctx, req.GetData().GetId())
		if err != nil {
			return err
		}
		if !exist {
			return r.Create(ctx, &systemV1.CreateAdminOperationLogRequest{Data: req.Data})
		}
	}

	if req.UpdateMask != nil {
		req.UpdateMask.Normalize()
		if !req.UpdateMask.IsValid(req.Data) {
			return errors.New("invalid field mask")
		}
		fieldmaskutil.Filter(req.GetData(), req.UpdateMask.GetPaths())
	}

	builder := r.data.db.Client().AdminOperationLog.
		UpdateOneID(req.Data.GetId()).
		SetNillableRequestID(req.Data.RequestId).
		SetNillableMethod(req.Data.Method).
		SetNillableOperation(req.Data.Operation).
		SetNillablePath(req.Data.Path).
		SetNillableReferer(req.Data.Referer).
		SetNillableRequestURI(req.Data.RequestUri).
		SetNillableRequestBody(req.Data.RequestBody).
		SetNillableRequestHeader(req.Data.RequestHeader).
		SetNillableResponse(req.Data.Response).
		SetNillableCostTime(timeutil.DurationpbSecond(req.Data.CostTime)).
		SetNillableUserID(req.Data.UserId).
		SetNillableUserName(req.Data.UserName).
		SetNillableClientIP(req.Data.ClientIp).
		SetNillableUserAgent(req.Data.UserAgent).
		SetNillableBrowserName(req.Data.BrowserName).
		SetNillableBrowserVersion(req.Data.BrowserVersion).
		SetNillableClientID(req.Data.ClientId).
		SetNillableClientName(req.Data.ClientName).
		SetNillableOsName(req.Data.OsName).
		SetNillableOsVersion(req.Data.OsVersion).
		SetNillableStatusCode(req.Data.StatusCode).
		SetNillableSuccess(req.Data.Success).
		SetNillableReason(req.Data.Reason).
		SetNillableLocation(req.Data.Location).
		SetNillableUpdateTime(timeutil.TimestamppbToTime(req.Data.UpdateTime))

	if req.Data.UpdateTime == nil {
		builder.SetUpdateTime(time.Now())
	}

	if req.UpdateMask != nil {
		nilPaths := fieldmaskutil.NilValuePaths(req.Data, req.GetUpdateMask().GetPaths())
		nilUpdater := entgoUpdate.BuildSetNullUpdater(nilPaths)
		if nilUpdater != nil {
			builder.Modify(nilUpdater)
		}
	}

	err := builder.Exec(ctx)
	if err != nil {
		r.log.Errorf("update one data failed: %s", err.Error())
		return err
	}

	return err
}

func (r *AdminOperationLogRepo) Delete(ctx context.Context, req *systemV1.DeleteAdminOperationLogRequest) (bool, error) {
	err := r.data.db.Client().AdminOperationLog.
		DeleteOneID(req.GetId()).
		Exec(ctx)
	return err != nil, err
}
