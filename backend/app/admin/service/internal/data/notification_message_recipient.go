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
	"kratos-admin/app/admin/service/internal/data/ent/notificationmessagerecipient"

	internalMessageV1 "kratos-admin/api/gen/go/internal_message/service/v1"

	"kratos-admin/pkg/middleware/auth"
)

type NotificationMessageRecipientRepo struct {
	data *Data
	log  *log.Helper
}

func NewNotificationMessageRecipientRepo(data *Data, logger log.Logger) *NotificationMessageRecipientRepo {
	l := log.NewHelper(log.With(logger, "module", "notification-message-recipient/repo/admin-service"))
	return &NotificationMessageRecipientRepo{
		data: data,
		log:  l,
	}
}

func (r *NotificationMessageRecipientRepo) toProtoStatus(in *notificationmessagerecipient.Status) *internalMessageV1.NotificationMessageRecipientStatus {
	if in == nil {
		return nil
	}

	switch *in {
	case notificationmessagerecipient.StatusUnknown:
		return trans.Ptr(internalMessageV1.NotificationMessageRecipientStatus_NotificationMessageRecipientStatus_Unknown)

	case notificationmessagerecipient.StatusReceived:
		return trans.Ptr(internalMessageV1.NotificationMessageRecipientStatus_NotificationMessageRecipientStatus_Received)

	case notificationmessagerecipient.StatusRead:
		return trans.Ptr(internalMessageV1.NotificationMessageRecipientStatus_NotificationMessageRecipientStatus_Read)

	case notificationmessagerecipient.StatusDeleted:
		return trans.Ptr(internalMessageV1.NotificationMessageRecipientStatus_NotificationMessageRecipientStatus_Deleted)

	case notificationmessagerecipient.StatusArchived:
		return trans.Ptr(internalMessageV1.NotificationMessageRecipientStatus_NotificationMessageRecipientStatus_Archived)

	default:
		return nil
	}
}

func (r *NotificationMessageRecipientRepo) toEntStatus(in *internalMessageV1.NotificationMessageRecipientStatus) *notificationmessagerecipient.Status {
	if in == nil {
		return nil
	}

	switch *in {
	case internalMessageV1.NotificationMessageRecipientStatus_NotificationMessageRecipientStatus_Unknown:
		return trans.Ptr(notificationmessagerecipient.StatusUnknown)

	case internalMessageV1.NotificationMessageRecipientStatus_NotificationMessageRecipientStatus_Received:
		return trans.Ptr(notificationmessagerecipient.StatusReceived)

	case internalMessageV1.NotificationMessageRecipientStatus_NotificationMessageRecipientStatus_Read:
		return trans.Ptr(notificationmessagerecipient.StatusRead)

	case internalMessageV1.NotificationMessageRecipientStatus_NotificationMessageRecipientStatus_Deleted:
		return trans.Ptr(notificationmessagerecipient.StatusDeleted)

	case internalMessageV1.NotificationMessageRecipientStatus_NotificationMessageRecipientStatus_Archived:
		return trans.Ptr(notificationmessagerecipient.StatusArchived)

	default:
		return nil
	}
}

func (r *NotificationMessageRecipientRepo) convertEntToProto(in *ent.NotificationMessageRecipient) *internalMessageV1.NotificationMessageRecipient {
	if in == nil {
		return nil
	}
	return &internalMessageV1.NotificationMessageRecipient{
		Id:          trans.Ptr(in.ID),
		MessageId:   in.MessageID,
		RecipientId: in.RecipientID,
		Status:      r.toProtoStatus(in.Status),
		CreateTime:  timeutil.TimeToTimestamppb(in.CreateTime),
		UpdateTime:  timeutil.TimeToTimestamppb(in.UpdateTime),
		DeleteTime:  timeutil.TimeToTimestamppb(in.DeleteTime),
	}
}

func (r *NotificationMessageRecipientRepo) Count(ctx context.Context, whereCond []func(s *sql.Selector)) (int, error) {
	builder := r.data.db.Client().NotificationMessageRecipient.Query()
	if len(whereCond) != 0 {
		builder.Modify(whereCond...)
	}

	count, err := builder.Count(ctx)
	if err != nil {
		r.log.Errorf("query count failed: %s", err.Error())
	}

	return count, err
}

func (r *NotificationMessageRecipientRepo) List(ctx context.Context, req *pagination.PagingRequest) (*internalMessageV1.ListNotificationMessageRecipientResponse, error) {
	builder := r.data.db.Client().NotificationMessageRecipient.Query()

	err, whereSelectors, querySelectors := entgo.BuildQuerySelector(
		req.GetQuery(), req.GetOrQuery(),
		req.GetPage(), req.GetPageSize(), req.GetNoPaging(),
		req.GetOrderBy(), notificationmessagerecipient.FieldCreateTime,
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

	items := make([]*internalMessageV1.NotificationMessageRecipient, 0, len(results))
	for _, res := range results {
		item := r.convertEntToProto(res)
		items = append(items, item)
	}

	count, err := r.Count(ctx, whereSelectors)
	if err != nil {
		return nil, err
	}

	return &internalMessageV1.ListNotificationMessageRecipientResponse{
		Total: uint32(count),
		Items: items,
	}, err
}

func (r *NotificationMessageRecipientRepo) IsExist(ctx context.Context, id uint32) (bool, error) {
	return r.data.db.Client().NotificationMessageRecipient.Query().
		Where(notificationmessagerecipient.IDEQ(id)).
		Exist(ctx)
}

func (r *NotificationMessageRecipientRepo) Get(ctx context.Context, req *internalMessageV1.GetNotificationMessageRecipientRequest) (*internalMessageV1.NotificationMessageRecipient, error) {
	ret, err := r.data.db.Client().NotificationMessageRecipient.Get(ctx, req.GetId())
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, internalMessageV1.ErrorResourceNotFound("message not found")
		}
		return nil, err
	}

	return r.convertEntToProto(ret), err
}

func (r *NotificationMessageRecipientRepo) Create(ctx context.Context, req *internalMessageV1.CreateNotificationMessageRecipientRequest, operator *auth.UserTokenPayload) error {
	if req.Data == nil {
		return errors.New("invalid request")
	}

	builder := r.data.db.Client().NotificationMessageRecipient.Create().
		SetNillableMessageID(req.Data.MessageId).
		SetNillableRecipientID(req.Data.RecipientId).
		SetNillableStatus(r.toEntStatus(req.Data.Status)).
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

func (r *NotificationMessageRecipientRepo) Update(ctx context.Context, req *internalMessageV1.UpdateNotificationMessageRecipientRequest, operator *auth.UserTokenPayload) error {
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
			return r.Create(ctx, &internalMessageV1.CreateNotificationMessageRecipientRequest{Data: req.Data}, operator)
		}
	}

	if req.UpdateMask != nil {
		req.UpdateMask.Normalize()
		if !req.UpdateMask.IsValid(req.Data) {
			return errors.New("invalid field mask")
		}
		fieldmaskutil.Filter(req.GetData(), req.UpdateMask.GetPaths())
	}

	builder := r.data.db.Client().NotificationMessageRecipient.UpdateOneID(req.Data.GetId()).
		SetNillableMessageID(req.Data.MessageId).
		SetNillableRecipientID(req.Data.RecipientId).
		SetNillableStatus(r.toEntStatus(req.Data.Status)).
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

func (r *NotificationMessageRecipientRepo) Delete(ctx context.Context, req *internalMessageV1.DeleteNotificationMessageRecipientRequest) (bool, error) {
	err := r.data.db.Client().NotificationMessageRecipient.
		DeleteOneID(req.GetId()).
		Exec(ctx)
	return err != nil, err
}
