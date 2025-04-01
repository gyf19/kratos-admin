// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"kratos-admin/app/admin/service/internal/data/ent/notificationmessagerecipient"
	"kratos-admin/app/admin/service/internal/data/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// NotificationMessageRecipientDelete is the builder for deleting a NotificationMessageRecipient entity.
type NotificationMessageRecipientDelete struct {
	config
	hooks    []Hook
	mutation *NotificationMessageRecipientMutation
}

// Where appends a list predicates to the NotificationMessageRecipientDelete builder.
func (nmrd *NotificationMessageRecipientDelete) Where(ps ...predicate.NotificationMessageRecipient) *NotificationMessageRecipientDelete {
	nmrd.mutation.Where(ps...)
	return nmrd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (nmrd *NotificationMessageRecipientDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, nmrd.sqlExec, nmrd.mutation, nmrd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (nmrd *NotificationMessageRecipientDelete) ExecX(ctx context.Context) int {
	n, err := nmrd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (nmrd *NotificationMessageRecipientDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(notificationmessagerecipient.Table, sqlgraph.NewFieldSpec(notificationmessagerecipient.FieldID, field.TypeUint32))
	if ps := nmrd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, nmrd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	nmrd.mutation.done = true
	return affected, err
}

// NotificationMessageRecipientDeleteOne is the builder for deleting a single NotificationMessageRecipient entity.
type NotificationMessageRecipientDeleteOne struct {
	nmrd *NotificationMessageRecipientDelete
}

// Where appends a list predicates to the NotificationMessageRecipientDelete builder.
func (nmrdo *NotificationMessageRecipientDeleteOne) Where(ps ...predicate.NotificationMessageRecipient) *NotificationMessageRecipientDeleteOne {
	nmrdo.nmrd.mutation.Where(ps...)
	return nmrdo
}

// Exec executes the deletion query.
func (nmrdo *NotificationMessageRecipientDeleteOne) Exec(ctx context.Context) error {
	n, err := nmrdo.nmrd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{notificationmessagerecipient.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (nmrdo *NotificationMessageRecipientDeleteOne) ExecX(ctx context.Context) {
	if err := nmrdo.Exec(ctx); err != nil {
		panic(err)
	}
}
