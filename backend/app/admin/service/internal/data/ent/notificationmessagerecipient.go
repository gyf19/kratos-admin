// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"kratos-admin/app/admin/service/internal/data/ent/notificationmessagerecipient"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// NotificationMessageRecipient is the model entity for the NotificationMessageRecipient schema.
type NotificationMessageRecipient struct {
	config `json:"-"`
	// ID of the ent.
	// id
	ID uint32 `json:"id,omitempty"`
	// 创建时间
	CreateTime *time.Time `json:"create_time,omitempty"`
	// 更新时间
	UpdateTime *time.Time `json:"update_time,omitempty"`
	// 删除时间
	DeleteTime *time.Time `json:"delete_time,omitempty"`
	// 群发消息ID
	MessageID *uint32 `json:"message_id,omitempty"`
	// 接收者用户ID
	RecipientID *uint32 `json:"recipient_id,omitempty"`
	// 消息状态
	Status       *notificationmessagerecipient.Status `json:"status,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*NotificationMessageRecipient) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case notificationmessagerecipient.FieldID, notificationmessagerecipient.FieldMessageID, notificationmessagerecipient.FieldRecipientID:
			values[i] = new(sql.NullInt64)
		case notificationmessagerecipient.FieldStatus:
			values[i] = new(sql.NullString)
		case notificationmessagerecipient.FieldCreateTime, notificationmessagerecipient.FieldUpdateTime, notificationmessagerecipient.FieldDeleteTime:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the NotificationMessageRecipient fields.
func (nmr *NotificationMessageRecipient) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case notificationmessagerecipient.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			nmr.ID = uint32(value.Int64)
		case notificationmessagerecipient.FieldCreateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field create_time", values[i])
			} else if value.Valid {
				nmr.CreateTime = new(time.Time)
				*nmr.CreateTime = value.Time
			}
		case notificationmessagerecipient.FieldUpdateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field update_time", values[i])
			} else if value.Valid {
				nmr.UpdateTime = new(time.Time)
				*nmr.UpdateTime = value.Time
			}
		case notificationmessagerecipient.FieldDeleteTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field delete_time", values[i])
			} else if value.Valid {
				nmr.DeleteTime = new(time.Time)
				*nmr.DeleteTime = value.Time
			}
		case notificationmessagerecipient.FieldMessageID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field message_id", values[i])
			} else if value.Valid {
				nmr.MessageID = new(uint32)
				*nmr.MessageID = uint32(value.Int64)
			}
		case notificationmessagerecipient.FieldRecipientID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field recipient_id", values[i])
			} else if value.Valid {
				nmr.RecipientID = new(uint32)
				*nmr.RecipientID = uint32(value.Int64)
			}
		case notificationmessagerecipient.FieldStatus:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				nmr.Status = new(notificationmessagerecipient.Status)
				*nmr.Status = notificationmessagerecipient.Status(value.String)
			}
		default:
			nmr.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the NotificationMessageRecipient.
// This includes values selected through modifiers, order, etc.
func (nmr *NotificationMessageRecipient) Value(name string) (ent.Value, error) {
	return nmr.selectValues.Get(name)
}

// Update returns a builder for updating this NotificationMessageRecipient.
// Note that you need to call NotificationMessageRecipient.Unwrap() before calling this method if this NotificationMessageRecipient
// was returned from a transaction, and the transaction was committed or rolled back.
func (nmr *NotificationMessageRecipient) Update() *NotificationMessageRecipientUpdateOne {
	return NewNotificationMessageRecipientClient(nmr.config).UpdateOne(nmr)
}

// Unwrap unwraps the NotificationMessageRecipient entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (nmr *NotificationMessageRecipient) Unwrap() *NotificationMessageRecipient {
	_tx, ok := nmr.config.driver.(*txDriver)
	if !ok {
		panic("ent: NotificationMessageRecipient is not a transactional entity")
	}
	nmr.config.driver = _tx.drv
	return nmr
}

// String implements the fmt.Stringer.
func (nmr *NotificationMessageRecipient) String() string {
	var builder strings.Builder
	builder.WriteString("NotificationMessageRecipient(")
	builder.WriteString(fmt.Sprintf("id=%v, ", nmr.ID))
	if v := nmr.CreateTime; v != nil {
		builder.WriteString("create_time=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", ")
	if v := nmr.UpdateTime; v != nil {
		builder.WriteString("update_time=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", ")
	if v := nmr.DeleteTime; v != nil {
		builder.WriteString("delete_time=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", ")
	if v := nmr.MessageID; v != nil {
		builder.WriteString("message_id=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	if v := nmr.RecipientID; v != nil {
		builder.WriteString("recipient_id=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	if v := nmr.Status; v != nil {
		builder.WriteString("status=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteByte(')')
	return builder.String()
}

// NotificationMessageRecipients is a parsable slice of NotificationMessageRecipient.
type NotificationMessageRecipients []*NotificationMessageRecipient
