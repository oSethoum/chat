// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
)

func (gr *Group) UserGroups(ctx context.Context) ([]*UserGroup, error) {
	result, err := gr.NamedUserGroups(graphql.GetFieldContext(ctx).Field.Alias)
	if IsNotLoaded(err) {
		result, err = gr.QueryUserGroups().All(ctx)
	}
	return result, err
}

func (gr *Group) Messages(ctx context.Context) ([]*Message, error) {
	result, err := gr.NamedMessages(graphql.GetFieldContext(ctx).Field.Alias)
	if IsNotLoaded(err) {
		result, err = gr.QueryMessages().All(ctx)
	}
	return result, err
}

func (m *Message) Group(ctx context.Context) (*Group, error) {
	result, err := m.Edges.GroupOrErr()
	if IsNotLoaded(err) {
		result, err = m.QueryGroup().Only(ctx)
	}
	return result, MaskNotFound(err)
}

func (m *Message) User(ctx context.Context) (*User, error) {
	result, err := m.Edges.UserOrErr()
	if IsNotLoaded(err) {
		result, err = m.QueryUser().Only(ctx)
	}
	return result, MaskNotFound(err)
}

func (u *User) UserGroups(ctx context.Context) ([]*UserGroup, error) {
	result, err := u.NamedUserGroups(graphql.GetFieldContext(ctx).Field.Alias)
	if IsNotLoaded(err) {
		result, err = u.QueryUserGroups().All(ctx)
	}
	return result, err
}

func (u *User) Messages(ctx context.Context) ([]*Message, error) {
	result, err := u.NamedMessages(graphql.GetFieldContext(ctx).Field.Alias)
	if IsNotLoaded(err) {
		result, err = u.QueryMessages().All(ctx)
	}
	return result, err
}

func (ug *UserGroup) Group(ctx context.Context) (*Group, error) {
	result, err := ug.Edges.GroupOrErr()
	if IsNotLoaded(err) {
		result, err = ug.QueryGroup().Only(ctx)
	}
	return result, MaskNotFound(err)
}

func (ug *UserGroup) User(ctx context.Context) (*User, error) {
	result, err := ug.Edges.UserOrErr()
	if IsNotLoaded(err) {
		result, err = ug.QueryUser().Only(ctx)
	}
	return result, MaskNotFound(err)
}