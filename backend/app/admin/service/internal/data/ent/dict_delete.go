// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"kratos-admin/app/admin/service/internal/data/ent/dict"
	"kratos-admin/app/admin/service/internal/data/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// DictDelete is the builder for deleting a Dict entity.
type DictDelete struct {
	config
	hooks    []Hook
	mutation *DictMutation
}

// Where appends a list predicates to the DictDelete builder.
func (dd *DictDelete) Where(ps ...predicate.Dict) *DictDelete {
	dd.mutation.Where(ps...)
	return dd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (dd *DictDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, dd.sqlExec, dd.mutation, dd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (dd *DictDelete) ExecX(ctx context.Context) int {
	n, err := dd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (dd *DictDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(dict.Table, sqlgraph.NewFieldSpec(dict.FieldID, field.TypeUint32))
	if ps := dd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, dd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	dd.mutation.done = true
	return affected, err
}

// DictDeleteOne is the builder for deleting a single Dict entity.
type DictDeleteOne struct {
	dd *DictDelete
}

// Where appends a list predicates to the DictDelete builder.
func (ddo *DictDeleteOne) Where(ps ...predicate.Dict) *DictDeleteOne {
	ddo.dd.mutation.Where(ps...)
	return ddo
}

// Exec executes the deletion query.
func (ddo *DictDeleteOne) Exec(ctx context.Context) error {
	n, err := ddo.dd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{dict.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (ddo *DictDeleteOne) ExecX(ctx context.Context) {
	if err := ddo.Exec(ctx); err != nil {
		panic(err)
	}
}
