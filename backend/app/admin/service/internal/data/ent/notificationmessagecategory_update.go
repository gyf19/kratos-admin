// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"kratos-admin/app/admin/service/internal/data/ent/notificationmessagecategory"
	"kratos-admin/app/admin/service/internal/data/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// NotificationMessageCategoryUpdate is the builder for updating NotificationMessageCategory entities.
type NotificationMessageCategoryUpdate struct {
	config
	hooks     []Hook
	mutation  *NotificationMessageCategoryMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the NotificationMessageCategoryUpdate builder.
func (nmcu *NotificationMessageCategoryUpdate) Where(ps ...predicate.NotificationMessageCategory) *NotificationMessageCategoryUpdate {
	nmcu.mutation.Where(ps...)
	return nmcu
}

// SetUpdateTime sets the "update_time" field.
func (nmcu *NotificationMessageCategoryUpdate) SetUpdateTime(t time.Time) *NotificationMessageCategoryUpdate {
	nmcu.mutation.SetUpdateTime(t)
	return nmcu
}

// SetNillableUpdateTime sets the "update_time" field if the given value is not nil.
func (nmcu *NotificationMessageCategoryUpdate) SetNillableUpdateTime(t *time.Time) *NotificationMessageCategoryUpdate {
	if t != nil {
		nmcu.SetUpdateTime(*t)
	}
	return nmcu
}

// ClearUpdateTime clears the value of the "update_time" field.
func (nmcu *NotificationMessageCategoryUpdate) ClearUpdateTime() *NotificationMessageCategoryUpdate {
	nmcu.mutation.ClearUpdateTime()
	return nmcu
}

// SetDeleteTime sets the "delete_time" field.
func (nmcu *NotificationMessageCategoryUpdate) SetDeleteTime(t time.Time) *NotificationMessageCategoryUpdate {
	nmcu.mutation.SetDeleteTime(t)
	return nmcu
}

// SetNillableDeleteTime sets the "delete_time" field if the given value is not nil.
func (nmcu *NotificationMessageCategoryUpdate) SetNillableDeleteTime(t *time.Time) *NotificationMessageCategoryUpdate {
	if t != nil {
		nmcu.SetDeleteTime(*t)
	}
	return nmcu
}

// ClearDeleteTime clears the value of the "delete_time" field.
func (nmcu *NotificationMessageCategoryUpdate) ClearDeleteTime() *NotificationMessageCategoryUpdate {
	nmcu.mutation.ClearDeleteTime()
	return nmcu
}

// SetCreateBy sets the "create_by" field.
func (nmcu *NotificationMessageCategoryUpdate) SetCreateBy(u uint32) *NotificationMessageCategoryUpdate {
	nmcu.mutation.ResetCreateBy()
	nmcu.mutation.SetCreateBy(u)
	return nmcu
}

// SetNillableCreateBy sets the "create_by" field if the given value is not nil.
func (nmcu *NotificationMessageCategoryUpdate) SetNillableCreateBy(u *uint32) *NotificationMessageCategoryUpdate {
	if u != nil {
		nmcu.SetCreateBy(*u)
	}
	return nmcu
}

// AddCreateBy adds u to the "create_by" field.
func (nmcu *NotificationMessageCategoryUpdate) AddCreateBy(u int32) *NotificationMessageCategoryUpdate {
	nmcu.mutation.AddCreateBy(u)
	return nmcu
}

// ClearCreateBy clears the value of the "create_by" field.
func (nmcu *NotificationMessageCategoryUpdate) ClearCreateBy() *NotificationMessageCategoryUpdate {
	nmcu.mutation.ClearCreateBy()
	return nmcu
}

// SetUpdateBy sets the "update_by" field.
func (nmcu *NotificationMessageCategoryUpdate) SetUpdateBy(u uint32) *NotificationMessageCategoryUpdate {
	nmcu.mutation.ResetUpdateBy()
	nmcu.mutation.SetUpdateBy(u)
	return nmcu
}

// SetNillableUpdateBy sets the "update_by" field if the given value is not nil.
func (nmcu *NotificationMessageCategoryUpdate) SetNillableUpdateBy(u *uint32) *NotificationMessageCategoryUpdate {
	if u != nil {
		nmcu.SetUpdateBy(*u)
	}
	return nmcu
}

// AddUpdateBy adds u to the "update_by" field.
func (nmcu *NotificationMessageCategoryUpdate) AddUpdateBy(u int32) *NotificationMessageCategoryUpdate {
	nmcu.mutation.AddUpdateBy(u)
	return nmcu
}

// ClearUpdateBy clears the value of the "update_by" field.
func (nmcu *NotificationMessageCategoryUpdate) ClearUpdateBy() *NotificationMessageCategoryUpdate {
	nmcu.mutation.ClearUpdateBy()
	return nmcu
}

// SetRemark sets the "remark" field.
func (nmcu *NotificationMessageCategoryUpdate) SetRemark(s string) *NotificationMessageCategoryUpdate {
	nmcu.mutation.SetRemark(s)
	return nmcu
}

// SetNillableRemark sets the "remark" field if the given value is not nil.
func (nmcu *NotificationMessageCategoryUpdate) SetNillableRemark(s *string) *NotificationMessageCategoryUpdate {
	if s != nil {
		nmcu.SetRemark(*s)
	}
	return nmcu
}

// ClearRemark clears the value of the "remark" field.
func (nmcu *NotificationMessageCategoryUpdate) ClearRemark() *NotificationMessageCategoryUpdate {
	nmcu.mutation.ClearRemark()
	return nmcu
}

// SetName sets the "name" field.
func (nmcu *NotificationMessageCategoryUpdate) SetName(s string) *NotificationMessageCategoryUpdate {
	nmcu.mutation.SetName(s)
	return nmcu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (nmcu *NotificationMessageCategoryUpdate) SetNillableName(s *string) *NotificationMessageCategoryUpdate {
	if s != nil {
		nmcu.SetName(*s)
	}
	return nmcu
}

// ClearName clears the value of the "name" field.
func (nmcu *NotificationMessageCategoryUpdate) ClearName() *NotificationMessageCategoryUpdate {
	nmcu.mutation.ClearName()
	return nmcu
}

// SetCode sets the "code" field.
func (nmcu *NotificationMessageCategoryUpdate) SetCode(s string) *NotificationMessageCategoryUpdate {
	nmcu.mutation.SetCode(s)
	return nmcu
}

// SetNillableCode sets the "code" field if the given value is not nil.
func (nmcu *NotificationMessageCategoryUpdate) SetNillableCode(s *string) *NotificationMessageCategoryUpdate {
	if s != nil {
		nmcu.SetCode(*s)
	}
	return nmcu
}

// ClearCode clears the value of the "code" field.
func (nmcu *NotificationMessageCategoryUpdate) ClearCode() *NotificationMessageCategoryUpdate {
	nmcu.mutation.ClearCode()
	return nmcu
}

// SetSortID sets the "sort_id" field.
func (nmcu *NotificationMessageCategoryUpdate) SetSortID(i int32) *NotificationMessageCategoryUpdate {
	nmcu.mutation.ResetSortID()
	nmcu.mutation.SetSortID(i)
	return nmcu
}

// SetNillableSortID sets the "sort_id" field if the given value is not nil.
func (nmcu *NotificationMessageCategoryUpdate) SetNillableSortID(i *int32) *NotificationMessageCategoryUpdate {
	if i != nil {
		nmcu.SetSortID(*i)
	}
	return nmcu
}

// AddSortID adds i to the "sort_id" field.
func (nmcu *NotificationMessageCategoryUpdate) AddSortID(i int32) *NotificationMessageCategoryUpdate {
	nmcu.mutation.AddSortID(i)
	return nmcu
}

// ClearSortID clears the value of the "sort_id" field.
func (nmcu *NotificationMessageCategoryUpdate) ClearSortID() *NotificationMessageCategoryUpdate {
	nmcu.mutation.ClearSortID()
	return nmcu
}

// SetEnable sets the "enable" field.
func (nmcu *NotificationMessageCategoryUpdate) SetEnable(b bool) *NotificationMessageCategoryUpdate {
	nmcu.mutation.SetEnable(b)
	return nmcu
}

// SetNillableEnable sets the "enable" field if the given value is not nil.
func (nmcu *NotificationMessageCategoryUpdate) SetNillableEnable(b *bool) *NotificationMessageCategoryUpdate {
	if b != nil {
		nmcu.SetEnable(*b)
	}
	return nmcu
}

// ClearEnable clears the value of the "enable" field.
func (nmcu *NotificationMessageCategoryUpdate) ClearEnable() *NotificationMessageCategoryUpdate {
	nmcu.mutation.ClearEnable()
	return nmcu
}

// SetParentID sets the "parent_id" field.
func (nmcu *NotificationMessageCategoryUpdate) SetParentID(u uint32) *NotificationMessageCategoryUpdate {
	nmcu.mutation.SetParentID(u)
	return nmcu
}

// SetNillableParentID sets the "parent_id" field if the given value is not nil.
func (nmcu *NotificationMessageCategoryUpdate) SetNillableParentID(u *uint32) *NotificationMessageCategoryUpdate {
	if u != nil {
		nmcu.SetParentID(*u)
	}
	return nmcu
}

// ClearParentID clears the value of the "parent_id" field.
func (nmcu *NotificationMessageCategoryUpdate) ClearParentID() *NotificationMessageCategoryUpdate {
	nmcu.mutation.ClearParentID()
	return nmcu
}

// SetParent sets the "parent" edge to the NotificationMessageCategory entity.
func (nmcu *NotificationMessageCategoryUpdate) SetParent(n *NotificationMessageCategory) *NotificationMessageCategoryUpdate {
	return nmcu.SetParentID(n.ID)
}

// AddChildIDs adds the "children" edge to the NotificationMessageCategory entity by IDs.
func (nmcu *NotificationMessageCategoryUpdate) AddChildIDs(ids ...uint32) *NotificationMessageCategoryUpdate {
	nmcu.mutation.AddChildIDs(ids...)
	return nmcu
}

// AddChildren adds the "children" edges to the NotificationMessageCategory entity.
func (nmcu *NotificationMessageCategoryUpdate) AddChildren(n ...*NotificationMessageCategory) *NotificationMessageCategoryUpdate {
	ids := make([]uint32, len(n))
	for i := range n {
		ids[i] = n[i].ID
	}
	return nmcu.AddChildIDs(ids...)
}

// Mutation returns the NotificationMessageCategoryMutation object of the builder.
func (nmcu *NotificationMessageCategoryUpdate) Mutation() *NotificationMessageCategoryMutation {
	return nmcu.mutation
}

// ClearParent clears the "parent" edge to the NotificationMessageCategory entity.
func (nmcu *NotificationMessageCategoryUpdate) ClearParent() *NotificationMessageCategoryUpdate {
	nmcu.mutation.ClearParent()
	return nmcu
}

// ClearChildren clears all "children" edges to the NotificationMessageCategory entity.
func (nmcu *NotificationMessageCategoryUpdate) ClearChildren() *NotificationMessageCategoryUpdate {
	nmcu.mutation.ClearChildren()
	return nmcu
}

// RemoveChildIDs removes the "children" edge to NotificationMessageCategory entities by IDs.
func (nmcu *NotificationMessageCategoryUpdate) RemoveChildIDs(ids ...uint32) *NotificationMessageCategoryUpdate {
	nmcu.mutation.RemoveChildIDs(ids...)
	return nmcu
}

// RemoveChildren removes "children" edges to NotificationMessageCategory entities.
func (nmcu *NotificationMessageCategoryUpdate) RemoveChildren(n ...*NotificationMessageCategory) *NotificationMessageCategoryUpdate {
	ids := make([]uint32, len(n))
	for i := range n {
		ids[i] = n[i].ID
	}
	return nmcu.RemoveChildIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (nmcu *NotificationMessageCategoryUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, nmcu.sqlSave, nmcu.mutation, nmcu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (nmcu *NotificationMessageCategoryUpdate) SaveX(ctx context.Context) int {
	affected, err := nmcu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (nmcu *NotificationMessageCategoryUpdate) Exec(ctx context.Context) error {
	_, err := nmcu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (nmcu *NotificationMessageCategoryUpdate) ExecX(ctx context.Context) {
	if err := nmcu.Exec(ctx); err != nil {
		panic(err)
	}
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (nmcu *NotificationMessageCategoryUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *NotificationMessageCategoryUpdate {
	nmcu.modifiers = append(nmcu.modifiers, modifiers...)
	return nmcu
}

func (nmcu *NotificationMessageCategoryUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(notificationmessagecategory.Table, notificationmessagecategory.Columns, sqlgraph.NewFieldSpec(notificationmessagecategory.FieldID, field.TypeUint32))
	if ps := nmcu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if nmcu.mutation.CreateTimeCleared() {
		_spec.ClearField(notificationmessagecategory.FieldCreateTime, field.TypeTime)
	}
	if value, ok := nmcu.mutation.UpdateTime(); ok {
		_spec.SetField(notificationmessagecategory.FieldUpdateTime, field.TypeTime, value)
	}
	if nmcu.mutation.UpdateTimeCleared() {
		_spec.ClearField(notificationmessagecategory.FieldUpdateTime, field.TypeTime)
	}
	if value, ok := nmcu.mutation.DeleteTime(); ok {
		_spec.SetField(notificationmessagecategory.FieldDeleteTime, field.TypeTime, value)
	}
	if nmcu.mutation.DeleteTimeCleared() {
		_spec.ClearField(notificationmessagecategory.FieldDeleteTime, field.TypeTime)
	}
	if value, ok := nmcu.mutation.CreateBy(); ok {
		_spec.SetField(notificationmessagecategory.FieldCreateBy, field.TypeUint32, value)
	}
	if value, ok := nmcu.mutation.AddedCreateBy(); ok {
		_spec.AddField(notificationmessagecategory.FieldCreateBy, field.TypeUint32, value)
	}
	if nmcu.mutation.CreateByCleared() {
		_spec.ClearField(notificationmessagecategory.FieldCreateBy, field.TypeUint32)
	}
	if value, ok := nmcu.mutation.UpdateBy(); ok {
		_spec.SetField(notificationmessagecategory.FieldUpdateBy, field.TypeUint32, value)
	}
	if value, ok := nmcu.mutation.AddedUpdateBy(); ok {
		_spec.AddField(notificationmessagecategory.FieldUpdateBy, field.TypeUint32, value)
	}
	if nmcu.mutation.UpdateByCleared() {
		_spec.ClearField(notificationmessagecategory.FieldUpdateBy, field.TypeUint32)
	}
	if value, ok := nmcu.mutation.Remark(); ok {
		_spec.SetField(notificationmessagecategory.FieldRemark, field.TypeString, value)
	}
	if nmcu.mutation.RemarkCleared() {
		_spec.ClearField(notificationmessagecategory.FieldRemark, field.TypeString)
	}
	if value, ok := nmcu.mutation.Name(); ok {
		_spec.SetField(notificationmessagecategory.FieldName, field.TypeString, value)
	}
	if nmcu.mutation.NameCleared() {
		_spec.ClearField(notificationmessagecategory.FieldName, field.TypeString)
	}
	if value, ok := nmcu.mutation.Code(); ok {
		_spec.SetField(notificationmessagecategory.FieldCode, field.TypeString, value)
	}
	if nmcu.mutation.CodeCleared() {
		_spec.ClearField(notificationmessagecategory.FieldCode, field.TypeString)
	}
	if value, ok := nmcu.mutation.SortID(); ok {
		_spec.SetField(notificationmessagecategory.FieldSortID, field.TypeInt32, value)
	}
	if value, ok := nmcu.mutation.AddedSortID(); ok {
		_spec.AddField(notificationmessagecategory.FieldSortID, field.TypeInt32, value)
	}
	if nmcu.mutation.SortIDCleared() {
		_spec.ClearField(notificationmessagecategory.FieldSortID, field.TypeInt32)
	}
	if value, ok := nmcu.mutation.Enable(); ok {
		_spec.SetField(notificationmessagecategory.FieldEnable, field.TypeBool, value)
	}
	if nmcu.mutation.EnableCleared() {
		_spec.ClearField(notificationmessagecategory.FieldEnable, field.TypeBool)
	}
	if nmcu.mutation.ParentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   notificationmessagecategory.ParentTable,
			Columns: []string{notificationmessagecategory.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(notificationmessagecategory.FieldID, field.TypeUint32),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := nmcu.mutation.ParentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   notificationmessagecategory.ParentTable,
			Columns: []string{notificationmessagecategory.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(notificationmessagecategory.FieldID, field.TypeUint32),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if nmcu.mutation.ChildrenCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   notificationmessagecategory.ChildrenTable,
			Columns: []string{notificationmessagecategory.ChildrenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(notificationmessagecategory.FieldID, field.TypeUint32),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := nmcu.mutation.RemovedChildrenIDs(); len(nodes) > 0 && !nmcu.mutation.ChildrenCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   notificationmessagecategory.ChildrenTable,
			Columns: []string{notificationmessagecategory.ChildrenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(notificationmessagecategory.FieldID, field.TypeUint32),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := nmcu.mutation.ChildrenIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   notificationmessagecategory.ChildrenTable,
			Columns: []string{notificationmessagecategory.ChildrenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(notificationmessagecategory.FieldID, field.TypeUint32),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.AddModifiers(nmcu.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, nmcu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{notificationmessagecategory.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	nmcu.mutation.done = true
	return n, nil
}

// NotificationMessageCategoryUpdateOne is the builder for updating a single NotificationMessageCategory entity.
type NotificationMessageCategoryUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *NotificationMessageCategoryMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetUpdateTime sets the "update_time" field.
func (nmcuo *NotificationMessageCategoryUpdateOne) SetUpdateTime(t time.Time) *NotificationMessageCategoryUpdateOne {
	nmcuo.mutation.SetUpdateTime(t)
	return nmcuo
}

// SetNillableUpdateTime sets the "update_time" field if the given value is not nil.
func (nmcuo *NotificationMessageCategoryUpdateOne) SetNillableUpdateTime(t *time.Time) *NotificationMessageCategoryUpdateOne {
	if t != nil {
		nmcuo.SetUpdateTime(*t)
	}
	return nmcuo
}

// ClearUpdateTime clears the value of the "update_time" field.
func (nmcuo *NotificationMessageCategoryUpdateOne) ClearUpdateTime() *NotificationMessageCategoryUpdateOne {
	nmcuo.mutation.ClearUpdateTime()
	return nmcuo
}

// SetDeleteTime sets the "delete_time" field.
func (nmcuo *NotificationMessageCategoryUpdateOne) SetDeleteTime(t time.Time) *NotificationMessageCategoryUpdateOne {
	nmcuo.mutation.SetDeleteTime(t)
	return nmcuo
}

// SetNillableDeleteTime sets the "delete_time" field if the given value is not nil.
func (nmcuo *NotificationMessageCategoryUpdateOne) SetNillableDeleteTime(t *time.Time) *NotificationMessageCategoryUpdateOne {
	if t != nil {
		nmcuo.SetDeleteTime(*t)
	}
	return nmcuo
}

// ClearDeleteTime clears the value of the "delete_time" field.
func (nmcuo *NotificationMessageCategoryUpdateOne) ClearDeleteTime() *NotificationMessageCategoryUpdateOne {
	nmcuo.mutation.ClearDeleteTime()
	return nmcuo
}

// SetCreateBy sets the "create_by" field.
func (nmcuo *NotificationMessageCategoryUpdateOne) SetCreateBy(u uint32) *NotificationMessageCategoryUpdateOne {
	nmcuo.mutation.ResetCreateBy()
	nmcuo.mutation.SetCreateBy(u)
	return nmcuo
}

// SetNillableCreateBy sets the "create_by" field if the given value is not nil.
func (nmcuo *NotificationMessageCategoryUpdateOne) SetNillableCreateBy(u *uint32) *NotificationMessageCategoryUpdateOne {
	if u != nil {
		nmcuo.SetCreateBy(*u)
	}
	return nmcuo
}

// AddCreateBy adds u to the "create_by" field.
func (nmcuo *NotificationMessageCategoryUpdateOne) AddCreateBy(u int32) *NotificationMessageCategoryUpdateOne {
	nmcuo.mutation.AddCreateBy(u)
	return nmcuo
}

// ClearCreateBy clears the value of the "create_by" field.
func (nmcuo *NotificationMessageCategoryUpdateOne) ClearCreateBy() *NotificationMessageCategoryUpdateOne {
	nmcuo.mutation.ClearCreateBy()
	return nmcuo
}

// SetUpdateBy sets the "update_by" field.
func (nmcuo *NotificationMessageCategoryUpdateOne) SetUpdateBy(u uint32) *NotificationMessageCategoryUpdateOne {
	nmcuo.mutation.ResetUpdateBy()
	nmcuo.mutation.SetUpdateBy(u)
	return nmcuo
}

// SetNillableUpdateBy sets the "update_by" field if the given value is not nil.
func (nmcuo *NotificationMessageCategoryUpdateOne) SetNillableUpdateBy(u *uint32) *NotificationMessageCategoryUpdateOne {
	if u != nil {
		nmcuo.SetUpdateBy(*u)
	}
	return nmcuo
}

// AddUpdateBy adds u to the "update_by" field.
func (nmcuo *NotificationMessageCategoryUpdateOne) AddUpdateBy(u int32) *NotificationMessageCategoryUpdateOne {
	nmcuo.mutation.AddUpdateBy(u)
	return nmcuo
}

// ClearUpdateBy clears the value of the "update_by" field.
func (nmcuo *NotificationMessageCategoryUpdateOne) ClearUpdateBy() *NotificationMessageCategoryUpdateOne {
	nmcuo.mutation.ClearUpdateBy()
	return nmcuo
}

// SetRemark sets the "remark" field.
func (nmcuo *NotificationMessageCategoryUpdateOne) SetRemark(s string) *NotificationMessageCategoryUpdateOne {
	nmcuo.mutation.SetRemark(s)
	return nmcuo
}

// SetNillableRemark sets the "remark" field if the given value is not nil.
func (nmcuo *NotificationMessageCategoryUpdateOne) SetNillableRemark(s *string) *NotificationMessageCategoryUpdateOne {
	if s != nil {
		nmcuo.SetRemark(*s)
	}
	return nmcuo
}

// ClearRemark clears the value of the "remark" field.
func (nmcuo *NotificationMessageCategoryUpdateOne) ClearRemark() *NotificationMessageCategoryUpdateOne {
	nmcuo.mutation.ClearRemark()
	return nmcuo
}

// SetName sets the "name" field.
func (nmcuo *NotificationMessageCategoryUpdateOne) SetName(s string) *NotificationMessageCategoryUpdateOne {
	nmcuo.mutation.SetName(s)
	return nmcuo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (nmcuo *NotificationMessageCategoryUpdateOne) SetNillableName(s *string) *NotificationMessageCategoryUpdateOne {
	if s != nil {
		nmcuo.SetName(*s)
	}
	return nmcuo
}

// ClearName clears the value of the "name" field.
func (nmcuo *NotificationMessageCategoryUpdateOne) ClearName() *NotificationMessageCategoryUpdateOne {
	nmcuo.mutation.ClearName()
	return nmcuo
}

// SetCode sets the "code" field.
func (nmcuo *NotificationMessageCategoryUpdateOne) SetCode(s string) *NotificationMessageCategoryUpdateOne {
	nmcuo.mutation.SetCode(s)
	return nmcuo
}

// SetNillableCode sets the "code" field if the given value is not nil.
func (nmcuo *NotificationMessageCategoryUpdateOne) SetNillableCode(s *string) *NotificationMessageCategoryUpdateOne {
	if s != nil {
		nmcuo.SetCode(*s)
	}
	return nmcuo
}

// ClearCode clears the value of the "code" field.
func (nmcuo *NotificationMessageCategoryUpdateOne) ClearCode() *NotificationMessageCategoryUpdateOne {
	nmcuo.mutation.ClearCode()
	return nmcuo
}

// SetSortID sets the "sort_id" field.
func (nmcuo *NotificationMessageCategoryUpdateOne) SetSortID(i int32) *NotificationMessageCategoryUpdateOne {
	nmcuo.mutation.ResetSortID()
	nmcuo.mutation.SetSortID(i)
	return nmcuo
}

// SetNillableSortID sets the "sort_id" field if the given value is not nil.
func (nmcuo *NotificationMessageCategoryUpdateOne) SetNillableSortID(i *int32) *NotificationMessageCategoryUpdateOne {
	if i != nil {
		nmcuo.SetSortID(*i)
	}
	return nmcuo
}

// AddSortID adds i to the "sort_id" field.
func (nmcuo *NotificationMessageCategoryUpdateOne) AddSortID(i int32) *NotificationMessageCategoryUpdateOne {
	nmcuo.mutation.AddSortID(i)
	return nmcuo
}

// ClearSortID clears the value of the "sort_id" field.
func (nmcuo *NotificationMessageCategoryUpdateOne) ClearSortID() *NotificationMessageCategoryUpdateOne {
	nmcuo.mutation.ClearSortID()
	return nmcuo
}

// SetEnable sets the "enable" field.
func (nmcuo *NotificationMessageCategoryUpdateOne) SetEnable(b bool) *NotificationMessageCategoryUpdateOne {
	nmcuo.mutation.SetEnable(b)
	return nmcuo
}

// SetNillableEnable sets the "enable" field if the given value is not nil.
func (nmcuo *NotificationMessageCategoryUpdateOne) SetNillableEnable(b *bool) *NotificationMessageCategoryUpdateOne {
	if b != nil {
		nmcuo.SetEnable(*b)
	}
	return nmcuo
}

// ClearEnable clears the value of the "enable" field.
func (nmcuo *NotificationMessageCategoryUpdateOne) ClearEnable() *NotificationMessageCategoryUpdateOne {
	nmcuo.mutation.ClearEnable()
	return nmcuo
}

// SetParentID sets the "parent_id" field.
func (nmcuo *NotificationMessageCategoryUpdateOne) SetParentID(u uint32) *NotificationMessageCategoryUpdateOne {
	nmcuo.mutation.SetParentID(u)
	return nmcuo
}

// SetNillableParentID sets the "parent_id" field if the given value is not nil.
func (nmcuo *NotificationMessageCategoryUpdateOne) SetNillableParentID(u *uint32) *NotificationMessageCategoryUpdateOne {
	if u != nil {
		nmcuo.SetParentID(*u)
	}
	return nmcuo
}

// ClearParentID clears the value of the "parent_id" field.
func (nmcuo *NotificationMessageCategoryUpdateOne) ClearParentID() *NotificationMessageCategoryUpdateOne {
	nmcuo.mutation.ClearParentID()
	return nmcuo
}

// SetParent sets the "parent" edge to the NotificationMessageCategory entity.
func (nmcuo *NotificationMessageCategoryUpdateOne) SetParent(n *NotificationMessageCategory) *NotificationMessageCategoryUpdateOne {
	return nmcuo.SetParentID(n.ID)
}

// AddChildIDs adds the "children" edge to the NotificationMessageCategory entity by IDs.
func (nmcuo *NotificationMessageCategoryUpdateOne) AddChildIDs(ids ...uint32) *NotificationMessageCategoryUpdateOne {
	nmcuo.mutation.AddChildIDs(ids...)
	return nmcuo
}

// AddChildren adds the "children" edges to the NotificationMessageCategory entity.
func (nmcuo *NotificationMessageCategoryUpdateOne) AddChildren(n ...*NotificationMessageCategory) *NotificationMessageCategoryUpdateOne {
	ids := make([]uint32, len(n))
	for i := range n {
		ids[i] = n[i].ID
	}
	return nmcuo.AddChildIDs(ids...)
}

// Mutation returns the NotificationMessageCategoryMutation object of the builder.
func (nmcuo *NotificationMessageCategoryUpdateOne) Mutation() *NotificationMessageCategoryMutation {
	return nmcuo.mutation
}

// ClearParent clears the "parent" edge to the NotificationMessageCategory entity.
func (nmcuo *NotificationMessageCategoryUpdateOne) ClearParent() *NotificationMessageCategoryUpdateOne {
	nmcuo.mutation.ClearParent()
	return nmcuo
}

// ClearChildren clears all "children" edges to the NotificationMessageCategory entity.
func (nmcuo *NotificationMessageCategoryUpdateOne) ClearChildren() *NotificationMessageCategoryUpdateOne {
	nmcuo.mutation.ClearChildren()
	return nmcuo
}

// RemoveChildIDs removes the "children" edge to NotificationMessageCategory entities by IDs.
func (nmcuo *NotificationMessageCategoryUpdateOne) RemoveChildIDs(ids ...uint32) *NotificationMessageCategoryUpdateOne {
	nmcuo.mutation.RemoveChildIDs(ids...)
	return nmcuo
}

// RemoveChildren removes "children" edges to NotificationMessageCategory entities.
func (nmcuo *NotificationMessageCategoryUpdateOne) RemoveChildren(n ...*NotificationMessageCategory) *NotificationMessageCategoryUpdateOne {
	ids := make([]uint32, len(n))
	for i := range n {
		ids[i] = n[i].ID
	}
	return nmcuo.RemoveChildIDs(ids...)
}

// Where appends a list predicates to the NotificationMessageCategoryUpdate builder.
func (nmcuo *NotificationMessageCategoryUpdateOne) Where(ps ...predicate.NotificationMessageCategory) *NotificationMessageCategoryUpdateOne {
	nmcuo.mutation.Where(ps...)
	return nmcuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (nmcuo *NotificationMessageCategoryUpdateOne) Select(field string, fields ...string) *NotificationMessageCategoryUpdateOne {
	nmcuo.fields = append([]string{field}, fields...)
	return nmcuo
}

// Save executes the query and returns the updated NotificationMessageCategory entity.
func (nmcuo *NotificationMessageCategoryUpdateOne) Save(ctx context.Context) (*NotificationMessageCategory, error) {
	return withHooks(ctx, nmcuo.sqlSave, nmcuo.mutation, nmcuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (nmcuo *NotificationMessageCategoryUpdateOne) SaveX(ctx context.Context) *NotificationMessageCategory {
	node, err := nmcuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (nmcuo *NotificationMessageCategoryUpdateOne) Exec(ctx context.Context) error {
	_, err := nmcuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (nmcuo *NotificationMessageCategoryUpdateOne) ExecX(ctx context.Context) {
	if err := nmcuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (nmcuo *NotificationMessageCategoryUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *NotificationMessageCategoryUpdateOne {
	nmcuo.modifiers = append(nmcuo.modifiers, modifiers...)
	return nmcuo
}

func (nmcuo *NotificationMessageCategoryUpdateOne) sqlSave(ctx context.Context) (_node *NotificationMessageCategory, err error) {
	_spec := sqlgraph.NewUpdateSpec(notificationmessagecategory.Table, notificationmessagecategory.Columns, sqlgraph.NewFieldSpec(notificationmessagecategory.FieldID, field.TypeUint32))
	id, ok := nmcuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "NotificationMessageCategory.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := nmcuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, notificationmessagecategory.FieldID)
		for _, f := range fields {
			if !notificationmessagecategory.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != notificationmessagecategory.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := nmcuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if nmcuo.mutation.CreateTimeCleared() {
		_spec.ClearField(notificationmessagecategory.FieldCreateTime, field.TypeTime)
	}
	if value, ok := nmcuo.mutation.UpdateTime(); ok {
		_spec.SetField(notificationmessagecategory.FieldUpdateTime, field.TypeTime, value)
	}
	if nmcuo.mutation.UpdateTimeCleared() {
		_spec.ClearField(notificationmessagecategory.FieldUpdateTime, field.TypeTime)
	}
	if value, ok := nmcuo.mutation.DeleteTime(); ok {
		_spec.SetField(notificationmessagecategory.FieldDeleteTime, field.TypeTime, value)
	}
	if nmcuo.mutation.DeleteTimeCleared() {
		_spec.ClearField(notificationmessagecategory.FieldDeleteTime, field.TypeTime)
	}
	if value, ok := nmcuo.mutation.CreateBy(); ok {
		_spec.SetField(notificationmessagecategory.FieldCreateBy, field.TypeUint32, value)
	}
	if value, ok := nmcuo.mutation.AddedCreateBy(); ok {
		_spec.AddField(notificationmessagecategory.FieldCreateBy, field.TypeUint32, value)
	}
	if nmcuo.mutation.CreateByCleared() {
		_spec.ClearField(notificationmessagecategory.FieldCreateBy, field.TypeUint32)
	}
	if value, ok := nmcuo.mutation.UpdateBy(); ok {
		_spec.SetField(notificationmessagecategory.FieldUpdateBy, field.TypeUint32, value)
	}
	if value, ok := nmcuo.mutation.AddedUpdateBy(); ok {
		_spec.AddField(notificationmessagecategory.FieldUpdateBy, field.TypeUint32, value)
	}
	if nmcuo.mutation.UpdateByCleared() {
		_spec.ClearField(notificationmessagecategory.FieldUpdateBy, field.TypeUint32)
	}
	if value, ok := nmcuo.mutation.Remark(); ok {
		_spec.SetField(notificationmessagecategory.FieldRemark, field.TypeString, value)
	}
	if nmcuo.mutation.RemarkCleared() {
		_spec.ClearField(notificationmessagecategory.FieldRemark, field.TypeString)
	}
	if value, ok := nmcuo.mutation.Name(); ok {
		_spec.SetField(notificationmessagecategory.FieldName, field.TypeString, value)
	}
	if nmcuo.mutation.NameCleared() {
		_spec.ClearField(notificationmessagecategory.FieldName, field.TypeString)
	}
	if value, ok := nmcuo.mutation.Code(); ok {
		_spec.SetField(notificationmessagecategory.FieldCode, field.TypeString, value)
	}
	if nmcuo.mutation.CodeCleared() {
		_spec.ClearField(notificationmessagecategory.FieldCode, field.TypeString)
	}
	if value, ok := nmcuo.mutation.SortID(); ok {
		_spec.SetField(notificationmessagecategory.FieldSortID, field.TypeInt32, value)
	}
	if value, ok := nmcuo.mutation.AddedSortID(); ok {
		_spec.AddField(notificationmessagecategory.FieldSortID, field.TypeInt32, value)
	}
	if nmcuo.mutation.SortIDCleared() {
		_spec.ClearField(notificationmessagecategory.FieldSortID, field.TypeInt32)
	}
	if value, ok := nmcuo.mutation.Enable(); ok {
		_spec.SetField(notificationmessagecategory.FieldEnable, field.TypeBool, value)
	}
	if nmcuo.mutation.EnableCleared() {
		_spec.ClearField(notificationmessagecategory.FieldEnable, field.TypeBool)
	}
	if nmcuo.mutation.ParentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   notificationmessagecategory.ParentTable,
			Columns: []string{notificationmessagecategory.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(notificationmessagecategory.FieldID, field.TypeUint32),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := nmcuo.mutation.ParentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   notificationmessagecategory.ParentTable,
			Columns: []string{notificationmessagecategory.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(notificationmessagecategory.FieldID, field.TypeUint32),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if nmcuo.mutation.ChildrenCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   notificationmessagecategory.ChildrenTable,
			Columns: []string{notificationmessagecategory.ChildrenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(notificationmessagecategory.FieldID, field.TypeUint32),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := nmcuo.mutation.RemovedChildrenIDs(); len(nodes) > 0 && !nmcuo.mutation.ChildrenCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   notificationmessagecategory.ChildrenTable,
			Columns: []string{notificationmessagecategory.ChildrenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(notificationmessagecategory.FieldID, field.TypeUint32),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := nmcuo.mutation.ChildrenIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   notificationmessagecategory.ChildrenTable,
			Columns: []string{notificationmessagecategory.ChildrenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(notificationmessagecategory.FieldID, field.TypeUint32),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.AddModifiers(nmcuo.modifiers...)
	_node = &NotificationMessageCategory{config: nmcuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, nmcuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{notificationmessagecategory.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	nmcuo.mutation.done = true
	return _node, nil
}
