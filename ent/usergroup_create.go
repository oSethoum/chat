// Code generated by ent, DO NOT EDIT.

package ent

import (
	"chat/ent/group"
	"chat/ent/user"
	"chat/ent/usergroup"
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// UserGroupCreate is the builder for creating a UserGroup entity.
type UserGroupCreate struct {
	config
	mutation *UserGroupMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (ugc *UserGroupCreate) SetCreatedAt(t time.Time) *UserGroupCreate {
	ugc.mutation.SetCreatedAt(t)
	return ugc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ugc *UserGroupCreate) SetNillableCreatedAt(t *time.Time) *UserGroupCreate {
	if t != nil {
		ugc.SetCreatedAt(*t)
	}
	return ugc
}

// SetGroupID sets the "group" edge to the Group entity by ID.
func (ugc *UserGroupCreate) SetGroupID(id int) *UserGroupCreate {
	ugc.mutation.SetGroupID(id)
	return ugc
}

// SetNillableGroupID sets the "group" edge to the Group entity by ID if the given value is not nil.
func (ugc *UserGroupCreate) SetNillableGroupID(id *int) *UserGroupCreate {
	if id != nil {
		ugc = ugc.SetGroupID(*id)
	}
	return ugc
}

// SetGroup sets the "group" edge to the Group entity.
func (ugc *UserGroupCreate) SetGroup(g *Group) *UserGroupCreate {
	return ugc.SetGroupID(g.ID)
}

// SetUserID sets the "user" edge to the User entity by ID.
func (ugc *UserGroupCreate) SetUserID(id int) *UserGroupCreate {
	ugc.mutation.SetUserID(id)
	return ugc
}

// SetNillableUserID sets the "user" edge to the User entity by ID if the given value is not nil.
func (ugc *UserGroupCreate) SetNillableUserID(id *int) *UserGroupCreate {
	if id != nil {
		ugc = ugc.SetUserID(*id)
	}
	return ugc
}

// SetUser sets the "user" edge to the User entity.
func (ugc *UserGroupCreate) SetUser(u *User) *UserGroupCreate {
	return ugc.SetUserID(u.ID)
}

// Mutation returns the UserGroupMutation object of the builder.
func (ugc *UserGroupCreate) Mutation() *UserGroupMutation {
	return ugc.mutation
}

// Save creates the UserGroup in the database.
func (ugc *UserGroupCreate) Save(ctx context.Context) (*UserGroup, error) {
	var (
		err  error
		node *UserGroup
	)
	ugc.defaults()
	if len(ugc.hooks) == 0 {
		if err = ugc.check(); err != nil {
			return nil, err
		}
		node, err = ugc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UserGroupMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ugc.check(); err != nil {
				return nil, err
			}
			ugc.mutation = mutation
			if node, err = ugc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(ugc.hooks) - 1; i >= 0; i-- {
			if ugc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ugc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, ugc.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*UserGroup)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from UserGroupMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (ugc *UserGroupCreate) SaveX(ctx context.Context) *UserGroup {
	v, err := ugc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ugc *UserGroupCreate) Exec(ctx context.Context) error {
	_, err := ugc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ugc *UserGroupCreate) ExecX(ctx context.Context) {
	if err := ugc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ugc *UserGroupCreate) defaults() {
	if _, ok := ugc.mutation.CreatedAt(); !ok {
		v := usergroup.DefaultCreatedAt()
		ugc.mutation.SetCreatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ugc *UserGroupCreate) check() error {
	if _, ok := ugc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "UserGroup.created_at"`)}
	}
	return nil
}

func (ugc *UserGroupCreate) sqlSave(ctx context.Context) (*UserGroup, error) {
	_node, _spec := ugc.createSpec()
	if err := sqlgraph.CreateNode(ctx, ugc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (ugc *UserGroupCreate) createSpec() (*UserGroup, *sqlgraph.CreateSpec) {
	var (
		_node = &UserGroup{config: ugc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: usergroup.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: usergroup.FieldID,
			},
		}
	)
	if value, ok := ugc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: usergroup.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if nodes := ugc.mutation.GroupIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   usergroup.GroupTable,
			Columns: []string{usergroup.GroupColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: group.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.group_user_groups = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ugc.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   usergroup.UserTable,
			Columns: []string{usergroup.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.user_user_groups = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// UserGroupCreateBulk is the builder for creating many UserGroup entities in bulk.
type UserGroupCreateBulk struct {
	config
	builders []*UserGroupCreate
}

// Save creates the UserGroup entities in the database.
func (ugcb *UserGroupCreateBulk) Save(ctx context.Context) ([]*UserGroup, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ugcb.builders))
	nodes := make([]*UserGroup, len(ugcb.builders))
	mutators := make([]Mutator, len(ugcb.builders))
	for i := range ugcb.builders {
		func(i int, root context.Context) {
			builder := ugcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*UserGroupMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, ugcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ugcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, ugcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ugcb *UserGroupCreateBulk) SaveX(ctx context.Context) []*UserGroup {
	v, err := ugcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ugcb *UserGroupCreateBulk) Exec(ctx context.Context) error {
	_, err := ugcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ugcb *UserGroupCreateBulk) ExecX(ctx context.Context) {
	if err := ugcb.Exec(ctx); err != nil {
		panic(err)
	}
}