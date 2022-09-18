// Code generated by ent, DO NOT EDIT.

package ent

import (
	"chat/ent/message"
	"chat/ent/predicate"
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// MessageDelete is the builder for deleting a Message entity.
type MessageDelete struct {
	config
	hooks    []Hook
	mutation *MessageMutation
}

// Where appends a list predicates to the MessageDelete builder.
func (md *MessageDelete) Where(ps ...predicate.Message) *MessageDelete {
	md.mutation.Where(ps...)
	return md
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (md *MessageDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(md.hooks) == 0 {
		affected, err = md.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*MessageMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			md.mutation = mutation
			affected, err = md.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(md.hooks) - 1; i >= 0; i-- {
			if md.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = md.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, md.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (md *MessageDelete) ExecX(ctx context.Context) int {
	n, err := md.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (md *MessageDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: message.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: message.FieldID,
			},
		},
	}
	if ps := md.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, md.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	return affected, err
}

// MessageDeleteOne is the builder for deleting a single Message entity.
type MessageDeleteOne struct {
	md *MessageDelete
}

// Exec executes the deletion query.
func (mdo *MessageDeleteOne) Exec(ctx context.Context) error {
	n, err := mdo.md.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{message.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (mdo *MessageDeleteOne) ExecX(ctx context.Context) {
	mdo.md.ExecX(ctx)
}
