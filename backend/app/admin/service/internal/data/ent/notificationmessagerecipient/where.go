// Code generated by ent, DO NOT EDIT.

package notificationmessagerecipient

import (
	"kratos-admin/app/admin/service/internal/data/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
)

// ID filters vertices based on their ID field.
func ID(id uint32) predicate.NotificationMessageRecipient {
	return predicate.NotificationMessageRecipient(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uint32) predicate.NotificationMessageRecipient {
	return predicate.NotificationMessageRecipient(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uint32) predicate.NotificationMessageRecipient {
	return predicate.NotificationMessageRecipient(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uint32) predicate.NotificationMessageRecipient {
	return predicate.NotificationMessageRecipient(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uint32) predicate.NotificationMessageRecipient {
	return predicate.NotificationMessageRecipient(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uint32) predicate.NotificationMessageRecipient {
	return predicate.NotificationMessageRecipient(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uint32) predicate.NotificationMessageRecipient {
	return predicate.NotificationMessageRecipient(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uint32) predicate.NotificationMessageRecipient {
	return predicate.NotificationMessageRecipient(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uint32) predicate.NotificationMessageRecipient {
	return predicate.NotificationMessageRecipient(sql.FieldLTE(FieldID, id))
}

// CreateTime applies equality check predicate on the "create_time" field. It's identical to CreateTimeEQ.
func CreateTime(v time.Time) predicate.NotificationMessageRecipient {
	return predicate.NotificationMessageRecipient(sql.FieldEQ(FieldCreateTime, v))
}

// UpdateTime applies equality check predicate on the "update_time" field. It's identical to UpdateTimeEQ.
func UpdateTime(v time.Time) predicate.NotificationMessageRecipient {
	return predicate.NotificationMessageRecipient(sql.FieldEQ(FieldUpdateTime, v))
}

// DeleteTime applies equality check predicate on the "delete_time" field. It's identical to DeleteTimeEQ.
func DeleteTime(v time.Time) predicate.NotificationMessageRecipient {
	return predicate.NotificationMessageRecipient(sql.FieldEQ(FieldDeleteTime, v))
}

// MessageID applies equality check predicate on the "message_id" field. It's identical to MessageIDEQ.
func MessageID(v uint32) predicate.NotificationMessageRecipient {
	return predicate.NotificationMessageRecipient(sql.FieldEQ(FieldMessageID, v))
}

// RecipientID applies equality check predicate on the "recipient_id" field. It's identical to RecipientIDEQ.
func RecipientID(v uint32) predicate.NotificationMessageRecipient {
	return predicate.NotificationMessageRecipient(sql.FieldEQ(FieldRecipientID, v))
}

// CreateTimeEQ applies the EQ predicate on the "create_time" field.
func CreateTimeEQ(v time.Time) predicate.NotificationMessageRecipient {
	return predicate.NotificationMessageRecipient(sql.FieldEQ(FieldCreateTime, v))
}

// CreateTimeNEQ applies the NEQ predicate on the "create_time" field.
func CreateTimeNEQ(v time.Time) predicate.NotificationMessageRecipient {
	return predicate.NotificationMessageRecipient(sql.FieldNEQ(FieldCreateTime, v))
}

// CreateTimeIn applies the In predicate on the "create_time" field.
func CreateTimeIn(vs ...time.Time) predicate.NotificationMessageRecipient {
	return predicate.NotificationMessageRecipient(sql.FieldIn(FieldCreateTime, vs...))
}

// CreateTimeNotIn applies the NotIn predicate on the "create_time" field.
func CreateTimeNotIn(vs ...time.Time) predicate.NotificationMessageRecipient {
	return predicate.NotificationMessageRecipient(sql.FieldNotIn(FieldCreateTime, vs...))
}

// CreateTimeGT applies the GT predicate on the "create_time" field.
func CreateTimeGT(v time.Time) predicate.NotificationMessageRecipient {
	return predicate.NotificationMessageRecipient(sql.FieldGT(FieldCreateTime, v))
}

// CreateTimeGTE applies the GTE predicate on the "create_time" field.
func CreateTimeGTE(v time.Time) predicate.NotificationMessageRecipient {
	return predicate.NotificationMessageRecipient(sql.FieldGTE(FieldCreateTime, v))
}

// CreateTimeLT applies the LT predicate on the "create_time" field.
func CreateTimeLT(v time.Time) predicate.NotificationMessageRecipient {
	return predicate.NotificationMessageRecipient(sql.FieldLT(FieldCreateTime, v))
}

// CreateTimeLTE applies the LTE predicate on the "create_time" field.
func CreateTimeLTE(v time.Time) predicate.NotificationMessageRecipient {
	return predicate.NotificationMessageRecipient(sql.FieldLTE(FieldCreateTime, v))
}

// CreateTimeIsNil applies the IsNil predicate on the "create_time" field.
func CreateTimeIsNil() predicate.NotificationMessageRecipient {
	return predicate.NotificationMessageRecipient(sql.FieldIsNull(FieldCreateTime))
}

// CreateTimeNotNil applies the NotNil predicate on the "create_time" field.
func CreateTimeNotNil() predicate.NotificationMessageRecipient {
	return predicate.NotificationMessageRecipient(sql.FieldNotNull(FieldCreateTime))
}

// UpdateTimeEQ applies the EQ predicate on the "update_time" field.
func UpdateTimeEQ(v time.Time) predicate.NotificationMessageRecipient {
	return predicate.NotificationMessageRecipient(sql.FieldEQ(FieldUpdateTime, v))
}

// UpdateTimeNEQ applies the NEQ predicate on the "update_time" field.
func UpdateTimeNEQ(v time.Time) predicate.NotificationMessageRecipient {
	return predicate.NotificationMessageRecipient(sql.FieldNEQ(FieldUpdateTime, v))
}

// UpdateTimeIn applies the In predicate on the "update_time" field.
func UpdateTimeIn(vs ...time.Time) predicate.NotificationMessageRecipient {
	return predicate.NotificationMessageRecipient(sql.FieldIn(FieldUpdateTime, vs...))
}

// UpdateTimeNotIn applies the NotIn predicate on the "update_time" field.
func UpdateTimeNotIn(vs ...time.Time) predicate.NotificationMessageRecipient {
	return predicate.NotificationMessageRecipient(sql.FieldNotIn(FieldUpdateTime, vs...))
}

// UpdateTimeGT applies the GT predicate on the "update_time" field.
func UpdateTimeGT(v time.Time) predicate.NotificationMessageRecipient {
	return predicate.NotificationMessageRecipient(sql.FieldGT(FieldUpdateTime, v))
}

// UpdateTimeGTE applies the GTE predicate on the "update_time" field.
func UpdateTimeGTE(v time.Time) predicate.NotificationMessageRecipient {
	return predicate.NotificationMessageRecipient(sql.FieldGTE(FieldUpdateTime, v))
}

// UpdateTimeLT applies the LT predicate on the "update_time" field.
func UpdateTimeLT(v time.Time) predicate.NotificationMessageRecipient {
	return predicate.NotificationMessageRecipient(sql.FieldLT(FieldUpdateTime, v))
}

// UpdateTimeLTE applies the LTE predicate on the "update_time" field.
func UpdateTimeLTE(v time.Time) predicate.NotificationMessageRecipient {
	return predicate.NotificationMessageRecipient(sql.FieldLTE(FieldUpdateTime, v))
}

// UpdateTimeIsNil applies the IsNil predicate on the "update_time" field.
func UpdateTimeIsNil() predicate.NotificationMessageRecipient {
	return predicate.NotificationMessageRecipient(sql.FieldIsNull(FieldUpdateTime))
}

// UpdateTimeNotNil applies the NotNil predicate on the "update_time" field.
func UpdateTimeNotNil() predicate.NotificationMessageRecipient {
	return predicate.NotificationMessageRecipient(sql.FieldNotNull(FieldUpdateTime))
}

// DeleteTimeEQ applies the EQ predicate on the "delete_time" field.
func DeleteTimeEQ(v time.Time) predicate.NotificationMessageRecipient {
	return predicate.NotificationMessageRecipient(sql.FieldEQ(FieldDeleteTime, v))
}

// DeleteTimeNEQ applies the NEQ predicate on the "delete_time" field.
func DeleteTimeNEQ(v time.Time) predicate.NotificationMessageRecipient {
	return predicate.NotificationMessageRecipient(sql.FieldNEQ(FieldDeleteTime, v))
}

// DeleteTimeIn applies the In predicate on the "delete_time" field.
func DeleteTimeIn(vs ...time.Time) predicate.NotificationMessageRecipient {
	return predicate.NotificationMessageRecipient(sql.FieldIn(FieldDeleteTime, vs...))
}

// DeleteTimeNotIn applies the NotIn predicate on the "delete_time" field.
func DeleteTimeNotIn(vs ...time.Time) predicate.NotificationMessageRecipient {
	return predicate.NotificationMessageRecipient(sql.FieldNotIn(FieldDeleteTime, vs...))
}

// DeleteTimeGT applies the GT predicate on the "delete_time" field.
func DeleteTimeGT(v time.Time) predicate.NotificationMessageRecipient {
	return predicate.NotificationMessageRecipient(sql.FieldGT(FieldDeleteTime, v))
}

// DeleteTimeGTE applies the GTE predicate on the "delete_time" field.
func DeleteTimeGTE(v time.Time) predicate.NotificationMessageRecipient {
	return predicate.NotificationMessageRecipient(sql.FieldGTE(FieldDeleteTime, v))
}

// DeleteTimeLT applies the LT predicate on the "delete_time" field.
func DeleteTimeLT(v time.Time) predicate.NotificationMessageRecipient {
	return predicate.NotificationMessageRecipient(sql.FieldLT(FieldDeleteTime, v))
}

// DeleteTimeLTE applies the LTE predicate on the "delete_time" field.
func DeleteTimeLTE(v time.Time) predicate.NotificationMessageRecipient {
	return predicate.NotificationMessageRecipient(sql.FieldLTE(FieldDeleteTime, v))
}

// DeleteTimeIsNil applies the IsNil predicate on the "delete_time" field.
func DeleteTimeIsNil() predicate.NotificationMessageRecipient {
	return predicate.NotificationMessageRecipient(sql.FieldIsNull(FieldDeleteTime))
}

// DeleteTimeNotNil applies the NotNil predicate on the "delete_time" field.
func DeleteTimeNotNil() predicate.NotificationMessageRecipient {
	return predicate.NotificationMessageRecipient(sql.FieldNotNull(FieldDeleteTime))
}

// MessageIDEQ applies the EQ predicate on the "message_id" field.
func MessageIDEQ(v uint32) predicate.NotificationMessageRecipient {
	return predicate.NotificationMessageRecipient(sql.FieldEQ(FieldMessageID, v))
}

// MessageIDNEQ applies the NEQ predicate on the "message_id" field.
func MessageIDNEQ(v uint32) predicate.NotificationMessageRecipient {
	return predicate.NotificationMessageRecipient(sql.FieldNEQ(FieldMessageID, v))
}

// MessageIDIn applies the In predicate on the "message_id" field.
func MessageIDIn(vs ...uint32) predicate.NotificationMessageRecipient {
	return predicate.NotificationMessageRecipient(sql.FieldIn(FieldMessageID, vs...))
}

// MessageIDNotIn applies the NotIn predicate on the "message_id" field.
func MessageIDNotIn(vs ...uint32) predicate.NotificationMessageRecipient {
	return predicate.NotificationMessageRecipient(sql.FieldNotIn(FieldMessageID, vs...))
}

// MessageIDGT applies the GT predicate on the "message_id" field.
func MessageIDGT(v uint32) predicate.NotificationMessageRecipient {
	return predicate.NotificationMessageRecipient(sql.FieldGT(FieldMessageID, v))
}

// MessageIDGTE applies the GTE predicate on the "message_id" field.
func MessageIDGTE(v uint32) predicate.NotificationMessageRecipient {
	return predicate.NotificationMessageRecipient(sql.FieldGTE(FieldMessageID, v))
}

// MessageIDLT applies the LT predicate on the "message_id" field.
func MessageIDLT(v uint32) predicate.NotificationMessageRecipient {
	return predicate.NotificationMessageRecipient(sql.FieldLT(FieldMessageID, v))
}

// MessageIDLTE applies the LTE predicate on the "message_id" field.
func MessageIDLTE(v uint32) predicate.NotificationMessageRecipient {
	return predicate.NotificationMessageRecipient(sql.FieldLTE(FieldMessageID, v))
}

// MessageIDIsNil applies the IsNil predicate on the "message_id" field.
func MessageIDIsNil() predicate.NotificationMessageRecipient {
	return predicate.NotificationMessageRecipient(sql.FieldIsNull(FieldMessageID))
}

// MessageIDNotNil applies the NotNil predicate on the "message_id" field.
func MessageIDNotNil() predicate.NotificationMessageRecipient {
	return predicate.NotificationMessageRecipient(sql.FieldNotNull(FieldMessageID))
}

// RecipientIDEQ applies the EQ predicate on the "recipient_id" field.
func RecipientIDEQ(v uint32) predicate.NotificationMessageRecipient {
	return predicate.NotificationMessageRecipient(sql.FieldEQ(FieldRecipientID, v))
}

// RecipientIDNEQ applies the NEQ predicate on the "recipient_id" field.
func RecipientIDNEQ(v uint32) predicate.NotificationMessageRecipient {
	return predicate.NotificationMessageRecipient(sql.FieldNEQ(FieldRecipientID, v))
}

// RecipientIDIn applies the In predicate on the "recipient_id" field.
func RecipientIDIn(vs ...uint32) predicate.NotificationMessageRecipient {
	return predicate.NotificationMessageRecipient(sql.FieldIn(FieldRecipientID, vs...))
}

// RecipientIDNotIn applies the NotIn predicate on the "recipient_id" field.
func RecipientIDNotIn(vs ...uint32) predicate.NotificationMessageRecipient {
	return predicate.NotificationMessageRecipient(sql.FieldNotIn(FieldRecipientID, vs...))
}

// RecipientIDGT applies the GT predicate on the "recipient_id" field.
func RecipientIDGT(v uint32) predicate.NotificationMessageRecipient {
	return predicate.NotificationMessageRecipient(sql.FieldGT(FieldRecipientID, v))
}

// RecipientIDGTE applies the GTE predicate on the "recipient_id" field.
func RecipientIDGTE(v uint32) predicate.NotificationMessageRecipient {
	return predicate.NotificationMessageRecipient(sql.FieldGTE(FieldRecipientID, v))
}

// RecipientIDLT applies the LT predicate on the "recipient_id" field.
func RecipientIDLT(v uint32) predicate.NotificationMessageRecipient {
	return predicate.NotificationMessageRecipient(sql.FieldLT(FieldRecipientID, v))
}

// RecipientIDLTE applies the LTE predicate on the "recipient_id" field.
func RecipientIDLTE(v uint32) predicate.NotificationMessageRecipient {
	return predicate.NotificationMessageRecipient(sql.FieldLTE(FieldRecipientID, v))
}

// RecipientIDIsNil applies the IsNil predicate on the "recipient_id" field.
func RecipientIDIsNil() predicate.NotificationMessageRecipient {
	return predicate.NotificationMessageRecipient(sql.FieldIsNull(FieldRecipientID))
}

// RecipientIDNotNil applies the NotNil predicate on the "recipient_id" field.
func RecipientIDNotNil() predicate.NotificationMessageRecipient {
	return predicate.NotificationMessageRecipient(sql.FieldNotNull(FieldRecipientID))
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v Status) predicate.NotificationMessageRecipient {
	return predicate.NotificationMessageRecipient(sql.FieldEQ(FieldStatus, v))
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v Status) predicate.NotificationMessageRecipient {
	return predicate.NotificationMessageRecipient(sql.FieldNEQ(FieldStatus, v))
}

// StatusIn applies the In predicate on the "status" field.
func StatusIn(vs ...Status) predicate.NotificationMessageRecipient {
	return predicate.NotificationMessageRecipient(sql.FieldIn(FieldStatus, vs...))
}

// StatusNotIn applies the NotIn predicate on the "status" field.
func StatusNotIn(vs ...Status) predicate.NotificationMessageRecipient {
	return predicate.NotificationMessageRecipient(sql.FieldNotIn(FieldStatus, vs...))
}

// StatusIsNil applies the IsNil predicate on the "status" field.
func StatusIsNil() predicate.NotificationMessageRecipient {
	return predicate.NotificationMessageRecipient(sql.FieldIsNull(FieldStatus))
}

// StatusNotNil applies the NotNil predicate on the "status" field.
func StatusNotNil() predicate.NotificationMessageRecipient {
	return predicate.NotificationMessageRecipient(sql.FieldNotNull(FieldStatus))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.NotificationMessageRecipient) predicate.NotificationMessageRecipient {
	return predicate.NotificationMessageRecipient(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.NotificationMessageRecipient) predicate.NotificationMessageRecipient {
	return predicate.NotificationMessageRecipient(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.NotificationMessageRecipient) predicate.NotificationMessageRecipient {
	return predicate.NotificationMessageRecipient(sql.NotPredicates(p))
}
