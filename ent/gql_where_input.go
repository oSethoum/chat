// Code generated by ent, DO NOT EDIT.

package ent

import (
	"chat/ent/group"
	"chat/ent/message"
	"chat/ent/predicate"
	"chat/ent/user"
	"chat/ent/usergroup"
	"errors"
	"fmt"
	"time"
)

// GroupWhereInput represents a where input for filtering Group queries.
type GroupWhereInput struct {
	Predicates []predicate.Group  `json:"-"`
	Not        *GroupWhereInput   `json:"not,omitempty"`
	Or         []*GroupWhereInput `json:"or,omitempty"`
	And        []*GroupWhereInput `json:"and,omitempty"`

	// "id" field predicates.
	ID      *int  `json:"id,omitempty"`
	IDNEQ   *int  `json:"idNEQ,omitempty"`
	IDIn    []int `json:"idIn,omitempty"`
	IDNotIn []int `json:"idNotIn,omitempty"`
	IDGT    *int  `json:"idGT,omitempty"`
	IDGTE   *int  `json:"idGTE,omitempty"`
	IDLT    *int  `json:"idLT,omitempty"`
	IDLTE   *int  `json:"idLTE,omitempty"`

	// "updated_at" field predicates.
	UpdatedAt      *time.Time  `json:"updatedAt,omitempty"`
	UpdatedAtNEQ   *time.Time  `json:"updatedAtNEQ,omitempty"`
	UpdatedAtIn    []time.Time `json:"updatedAtIn,omitempty"`
	UpdatedAtNotIn []time.Time `json:"updatedAtNotIn,omitempty"`
	UpdatedAtGT    *time.Time  `json:"updatedAtGT,omitempty"`
	UpdatedAtGTE   *time.Time  `json:"updatedAtGTE,omitempty"`
	UpdatedAtLT    *time.Time  `json:"updatedAtLT,omitempty"`
	UpdatedAtLTE   *time.Time  `json:"updatedAtLTE,omitempty"`

	// "created_at" field predicates.
	CreatedAt      *time.Time  `json:"createdAt,omitempty"`
	CreatedAtNEQ   *time.Time  `json:"createdAtNEQ,omitempty"`
	CreatedAtIn    []time.Time `json:"createdAtIn,omitempty"`
	CreatedAtNotIn []time.Time `json:"createdAtNotIn,omitempty"`
	CreatedAtGT    *time.Time  `json:"createdAtGT,omitempty"`
	CreatedAtGTE   *time.Time  `json:"createdAtGTE,omitempty"`
	CreatedAtLT    *time.Time  `json:"createdAtLT,omitempty"`
	CreatedAtLTE   *time.Time  `json:"createdAtLTE,omitempty"`

	// "name" field predicates.
	Name             *string  `json:"name,omitempty"`
	NameNEQ          *string  `json:"nameNEQ,omitempty"`
	NameIn           []string `json:"nameIn,omitempty"`
	NameNotIn        []string `json:"nameNotIn,omitempty"`
	NameGT           *string  `json:"nameGT,omitempty"`
	NameGTE          *string  `json:"nameGTE,omitempty"`
	NameLT           *string  `json:"nameLT,omitempty"`
	NameLTE          *string  `json:"nameLTE,omitempty"`
	NameContains     *string  `json:"nameContains,omitempty"`
	NameHasPrefix    *string  `json:"nameHasPrefix,omitempty"`
	NameHasSuffix    *string  `json:"nameHasSuffix,omitempty"`
	NameEqualFold    *string  `json:"nameEqualFold,omitempty"`
	NameContainsFold *string  `json:"nameContainsFold,omitempty"`

	// "userGroups" edge predicates.
	HasUserGroups     *bool                  `json:"hasUserGroups,omitempty"`
	HasUserGroupsWith []*UserGroupWhereInput `json:"hasUserGroupsWith,omitempty"`

	// "messages" edge predicates.
	HasMessages     *bool                `json:"hasMessages,omitempty"`
	HasMessagesWith []*MessageWhereInput `json:"hasMessagesWith,omitempty"`
}

// AddPredicates adds custom predicates to the where input to be used during the filtering phase.
func (i *GroupWhereInput) AddPredicates(predicates ...predicate.Group) {
	i.Predicates = append(i.Predicates, predicates...)
}

// Filter applies the GroupWhereInput filter on the GroupQuery builder.
func (i *GroupWhereInput) Filter(q *GroupQuery) (*GroupQuery, error) {
	if i == nil {
		return q, nil
	}
	p, err := i.P()
	if err != nil {
		if err == ErrEmptyGroupWhereInput {
			return q, nil
		}
		return nil, err
	}
	return q.Where(p), nil
}

// ErrEmptyGroupWhereInput is returned in case the GroupWhereInput is empty.
var ErrEmptyGroupWhereInput = errors.New("ent: empty predicate GroupWhereInput")

// P returns a predicate for filtering groups.
// An error is returned if the input is empty or invalid.
func (i *GroupWhereInput) P() (predicate.Group, error) {
	var predicates []predicate.Group
	if i.Not != nil {
		p, err := i.Not.P()
		if err != nil {
			return nil, fmt.Errorf("%w: field 'not'", err)
		}
		predicates = append(predicates, group.Not(p))
	}
	switch n := len(i.Or); {
	case n == 1:
		p, err := i.Or[0].P()
		if err != nil {
			return nil, fmt.Errorf("%w: field 'or'", err)
		}
		predicates = append(predicates, p)
	case n > 1:
		or := make([]predicate.Group, 0, n)
		for _, w := range i.Or {
			p, err := w.P()
			if err != nil {
				return nil, fmt.Errorf("%w: field 'or'", err)
			}
			or = append(or, p)
		}
		predicates = append(predicates, group.Or(or...))
	}
	switch n := len(i.And); {
	case n == 1:
		p, err := i.And[0].P()
		if err != nil {
			return nil, fmt.Errorf("%w: field 'and'", err)
		}
		predicates = append(predicates, p)
	case n > 1:
		and := make([]predicate.Group, 0, n)
		for _, w := range i.And {
			p, err := w.P()
			if err != nil {
				return nil, fmt.Errorf("%w: field 'and'", err)
			}
			and = append(and, p)
		}
		predicates = append(predicates, group.And(and...))
	}
	predicates = append(predicates, i.Predicates...)
	if i.ID != nil {
		predicates = append(predicates, group.IDEQ(*i.ID))
	}
	if i.IDNEQ != nil {
		predicates = append(predicates, group.IDNEQ(*i.IDNEQ))
	}
	if len(i.IDIn) > 0 {
		predicates = append(predicates, group.IDIn(i.IDIn...))
	}
	if len(i.IDNotIn) > 0 {
		predicates = append(predicates, group.IDNotIn(i.IDNotIn...))
	}
	if i.IDGT != nil {
		predicates = append(predicates, group.IDGT(*i.IDGT))
	}
	if i.IDGTE != nil {
		predicates = append(predicates, group.IDGTE(*i.IDGTE))
	}
	if i.IDLT != nil {
		predicates = append(predicates, group.IDLT(*i.IDLT))
	}
	if i.IDLTE != nil {
		predicates = append(predicates, group.IDLTE(*i.IDLTE))
	}
	if i.UpdatedAt != nil {
		predicates = append(predicates, group.UpdatedAtEQ(*i.UpdatedAt))
	}
	if i.UpdatedAtNEQ != nil {
		predicates = append(predicates, group.UpdatedAtNEQ(*i.UpdatedAtNEQ))
	}
	if len(i.UpdatedAtIn) > 0 {
		predicates = append(predicates, group.UpdatedAtIn(i.UpdatedAtIn...))
	}
	if len(i.UpdatedAtNotIn) > 0 {
		predicates = append(predicates, group.UpdatedAtNotIn(i.UpdatedAtNotIn...))
	}
	if i.UpdatedAtGT != nil {
		predicates = append(predicates, group.UpdatedAtGT(*i.UpdatedAtGT))
	}
	if i.UpdatedAtGTE != nil {
		predicates = append(predicates, group.UpdatedAtGTE(*i.UpdatedAtGTE))
	}
	if i.UpdatedAtLT != nil {
		predicates = append(predicates, group.UpdatedAtLT(*i.UpdatedAtLT))
	}
	if i.UpdatedAtLTE != nil {
		predicates = append(predicates, group.UpdatedAtLTE(*i.UpdatedAtLTE))
	}
	if i.CreatedAt != nil {
		predicates = append(predicates, group.CreatedAtEQ(*i.CreatedAt))
	}
	if i.CreatedAtNEQ != nil {
		predicates = append(predicates, group.CreatedAtNEQ(*i.CreatedAtNEQ))
	}
	if len(i.CreatedAtIn) > 0 {
		predicates = append(predicates, group.CreatedAtIn(i.CreatedAtIn...))
	}
	if len(i.CreatedAtNotIn) > 0 {
		predicates = append(predicates, group.CreatedAtNotIn(i.CreatedAtNotIn...))
	}
	if i.CreatedAtGT != nil {
		predicates = append(predicates, group.CreatedAtGT(*i.CreatedAtGT))
	}
	if i.CreatedAtGTE != nil {
		predicates = append(predicates, group.CreatedAtGTE(*i.CreatedAtGTE))
	}
	if i.CreatedAtLT != nil {
		predicates = append(predicates, group.CreatedAtLT(*i.CreatedAtLT))
	}
	if i.CreatedAtLTE != nil {
		predicates = append(predicates, group.CreatedAtLTE(*i.CreatedAtLTE))
	}
	if i.Name != nil {
		predicates = append(predicates, group.NameEQ(*i.Name))
	}
	if i.NameNEQ != nil {
		predicates = append(predicates, group.NameNEQ(*i.NameNEQ))
	}
	if len(i.NameIn) > 0 {
		predicates = append(predicates, group.NameIn(i.NameIn...))
	}
	if len(i.NameNotIn) > 0 {
		predicates = append(predicates, group.NameNotIn(i.NameNotIn...))
	}
	if i.NameGT != nil {
		predicates = append(predicates, group.NameGT(*i.NameGT))
	}
	if i.NameGTE != nil {
		predicates = append(predicates, group.NameGTE(*i.NameGTE))
	}
	if i.NameLT != nil {
		predicates = append(predicates, group.NameLT(*i.NameLT))
	}
	if i.NameLTE != nil {
		predicates = append(predicates, group.NameLTE(*i.NameLTE))
	}
	if i.NameContains != nil {
		predicates = append(predicates, group.NameContains(*i.NameContains))
	}
	if i.NameHasPrefix != nil {
		predicates = append(predicates, group.NameHasPrefix(*i.NameHasPrefix))
	}
	if i.NameHasSuffix != nil {
		predicates = append(predicates, group.NameHasSuffix(*i.NameHasSuffix))
	}
	if i.NameEqualFold != nil {
		predicates = append(predicates, group.NameEqualFold(*i.NameEqualFold))
	}
	if i.NameContainsFold != nil {
		predicates = append(predicates, group.NameContainsFold(*i.NameContainsFold))
	}

	if i.HasUserGroups != nil {
		p := group.HasUserGroups()
		if !*i.HasUserGroups {
			p = group.Not(p)
		}
		predicates = append(predicates, p)
	}
	if len(i.HasUserGroupsWith) > 0 {
		with := make([]predicate.UserGroup, 0, len(i.HasUserGroupsWith))
		for _, w := range i.HasUserGroupsWith {
			p, err := w.P()
			if err != nil {
				return nil, fmt.Errorf("%w: field 'HasUserGroupsWith'", err)
			}
			with = append(with, p)
		}
		predicates = append(predicates, group.HasUserGroupsWith(with...))
	}
	if i.HasMessages != nil {
		p := group.HasMessages()
		if !*i.HasMessages {
			p = group.Not(p)
		}
		predicates = append(predicates, p)
	}
	if len(i.HasMessagesWith) > 0 {
		with := make([]predicate.Message, 0, len(i.HasMessagesWith))
		for _, w := range i.HasMessagesWith {
			p, err := w.P()
			if err != nil {
				return nil, fmt.Errorf("%w: field 'HasMessagesWith'", err)
			}
			with = append(with, p)
		}
		predicates = append(predicates, group.HasMessagesWith(with...))
	}
	switch len(predicates) {
	case 0:
		return nil, ErrEmptyGroupWhereInput
	case 1:
		return predicates[0], nil
	default:
		return group.And(predicates...), nil
	}
}

// MessageWhereInput represents a where input for filtering Message queries.
type MessageWhereInput struct {
	Predicates []predicate.Message  `json:"-"`
	Not        *MessageWhereInput   `json:"not,omitempty"`
	Or         []*MessageWhereInput `json:"or,omitempty"`
	And        []*MessageWhereInput `json:"and,omitempty"`

	// "id" field predicates.
	ID      *int  `json:"id,omitempty"`
	IDNEQ   *int  `json:"idNEQ,omitempty"`
	IDIn    []int `json:"idIn,omitempty"`
	IDNotIn []int `json:"idNotIn,omitempty"`
	IDGT    *int  `json:"idGT,omitempty"`
	IDGTE   *int  `json:"idGTE,omitempty"`
	IDLT    *int  `json:"idLT,omitempty"`
	IDLTE   *int  `json:"idLTE,omitempty"`

	// "sent_at" field predicates.
	SentAt      *time.Time  `json:"sentAt,omitempty"`
	SentAtNEQ   *time.Time  `json:"sentAtNEQ,omitempty"`
	SentAtIn    []time.Time `json:"sentAtIn,omitempty"`
	SentAtNotIn []time.Time `json:"sentAtNotIn,omitempty"`
	SentAtGT    *time.Time  `json:"sentAtGT,omitempty"`
	SentAtGTE   *time.Time  `json:"sentAtGTE,omitempty"`
	SentAtLT    *time.Time  `json:"sentAtLT,omitempty"`
	SentAtLTE   *time.Time  `json:"sentAtLTE,omitempty"`

	// "message" field predicates.
	Message             *string  `json:"message,omitempty"`
	MessageNEQ          *string  `json:"messageNEQ,omitempty"`
	MessageIn           []string `json:"messageIn,omitempty"`
	MessageNotIn        []string `json:"messageNotIn,omitempty"`
	MessageGT           *string  `json:"messageGT,omitempty"`
	MessageGTE          *string  `json:"messageGTE,omitempty"`
	MessageLT           *string  `json:"messageLT,omitempty"`
	MessageLTE          *string  `json:"messageLTE,omitempty"`
	MessageContains     *string  `json:"messageContains,omitempty"`
	MessageHasPrefix    *string  `json:"messageHasPrefix,omitempty"`
	MessageHasSuffix    *string  `json:"messageHasSuffix,omitempty"`
	MessageEqualFold    *string  `json:"messageEqualFold,omitempty"`
	MessageContainsFold *string  `json:"messageContainsFold,omitempty"`

	// "group" edge predicates.
	HasGroup     *bool              `json:"hasGroup,omitempty"`
	HasGroupWith []*GroupWhereInput `json:"hasGroupWith,omitempty"`

	// "user" edge predicates.
	HasUser     *bool             `json:"hasUser,omitempty"`
	HasUserWith []*UserWhereInput `json:"hasUserWith,omitempty"`
}

// AddPredicates adds custom predicates to the where input to be used during the filtering phase.
func (i *MessageWhereInput) AddPredicates(predicates ...predicate.Message) {
	i.Predicates = append(i.Predicates, predicates...)
}

// Filter applies the MessageWhereInput filter on the MessageQuery builder.
func (i *MessageWhereInput) Filter(q *MessageQuery) (*MessageQuery, error) {
	if i == nil {
		return q, nil
	}
	p, err := i.P()
	if err != nil {
		if err == ErrEmptyMessageWhereInput {
			return q, nil
		}
		return nil, err
	}
	return q.Where(p), nil
}

// ErrEmptyMessageWhereInput is returned in case the MessageWhereInput is empty.
var ErrEmptyMessageWhereInput = errors.New("ent: empty predicate MessageWhereInput")

// P returns a predicate for filtering messages.
// An error is returned if the input is empty or invalid.
func (i *MessageWhereInput) P() (predicate.Message, error) {
	var predicates []predicate.Message
	if i.Not != nil {
		p, err := i.Not.P()
		if err != nil {
			return nil, fmt.Errorf("%w: field 'not'", err)
		}
		predicates = append(predicates, message.Not(p))
	}
	switch n := len(i.Or); {
	case n == 1:
		p, err := i.Or[0].P()
		if err != nil {
			return nil, fmt.Errorf("%w: field 'or'", err)
		}
		predicates = append(predicates, p)
	case n > 1:
		or := make([]predicate.Message, 0, n)
		for _, w := range i.Or {
			p, err := w.P()
			if err != nil {
				return nil, fmt.Errorf("%w: field 'or'", err)
			}
			or = append(or, p)
		}
		predicates = append(predicates, message.Or(or...))
	}
	switch n := len(i.And); {
	case n == 1:
		p, err := i.And[0].P()
		if err != nil {
			return nil, fmt.Errorf("%w: field 'and'", err)
		}
		predicates = append(predicates, p)
	case n > 1:
		and := make([]predicate.Message, 0, n)
		for _, w := range i.And {
			p, err := w.P()
			if err != nil {
				return nil, fmt.Errorf("%w: field 'and'", err)
			}
			and = append(and, p)
		}
		predicates = append(predicates, message.And(and...))
	}
	predicates = append(predicates, i.Predicates...)
	if i.ID != nil {
		predicates = append(predicates, message.IDEQ(*i.ID))
	}
	if i.IDNEQ != nil {
		predicates = append(predicates, message.IDNEQ(*i.IDNEQ))
	}
	if len(i.IDIn) > 0 {
		predicates = append(predicates, message.IDIn(i.IDIn...))
	}
	if len(i.IDNotIn) > 0 {
		predicates = append(predicates, message.IDNotIn(i.IDNotIn...))
	}
	if i.IDGT != nil {
		predicates = append(predicates, message.IDGT(*i.IDGT))
	}
	if i.IDGTE != nil {
		predicates = append(predicates, message.IDGTE(*i.IDGTE))
	}
	if i.IDLT != nil {
		predicates = append(predicates, message.IDLT(*i.IDLT))
	}
	if i.IDLTE != nil {
		predicates = append(predicates, message.IDLTE(*i.IDLTE))
	}
	if i.SentAt != nil {
		predicates = append(predicates, message.SentAtEQ(*i.SentAt))
	}
	if i.SentAtNEQ != nil {
		predicates = append(predicates, message.SentAtNEQ(*i.SentAtNEQ))
	}
	if len(i.SentAtIn) > 0 {
		predicates = append(predicates, message.SentAtIn(i.SentAtIn...))
	}
	if len(i.SentAtNotIn) > 0 {
		predicates = append(predicates, message.SentAtNotIn(i.SentAtNotIn...))
	}
	if i.SentAtGT != nil {
		predicates = append(predicates, message.SentAtGT(*i.SentAtGT))
	}
	if i.SentAtGTE != nil {
		predicates = append(predicates, message.SentAtGTE(*i.SentAtGTE))
	}
	if i.SentAtLT != nil {
		predicates = append(predicates, message.SentAtLT(*i.SentAtLT))
	}
	if i.SentAtLTE != nil {
		predicates = append(predicates, message.SentAtLTE(*i.SentAtLTE))
	}
	if i.Message != nil {
		predicates = append(predicates, message.MessageEQ(*i.Message))
	}
	if i.MessageNEQ != nil {
		predicates = append(predicates, message.MessageNEQ(*i.MessageNEQ))
	}
	if len(i.MessageIn) > 0 {
		predicates = append(predicates, message.MessageIn(i.MessageIn...))
	}
	if len(i.MessageNotIn) > 0 {
		predicates = append(predicates, message.MessageNotIn(i.MessageNotIn...))
	}
	if i.MessageGT != nil {
		predicates = append(predicates, message.MessageGT(*i.MessageGT))
	}
	if i.MessageGTE != nil {
		predicates = append(predicates, message.MessageGTE(*i.MessageGTE))
	}
	if i.MessageLT != nil {
		predicates = append(predicates, message.MessageLT(*i.MessageLT))
	}
	if i.MessageLTE != nil {
		predicates = append(predicates, message.MessageLTE(*i.MessageLTE))
	}
	if i.MessageContains != nil {
		predicates = append(predicates, message.MessageContains(*i.MessageContains))
	}
	if i.MessageHasPrefix != nil {
		predicates = append(predicates, message.MessageHasPrefix(*i.MessageHasPrefix))
	}
	if i.MessageHasSuffix != nil {
		predicates = append(predicates, message.MessageHasSuffix(*i.MessageHasSuffix))
	}
	if i.MessageEqualFold != nil {
		predicates = append(predicates, message.MessageEqualFold(*i.MessageEqualFold))
	}
	if i.MessageContainsFold != nil {
		predicates = append(predicates, message.MessageContainsFold(*i.MessageContainsFold))
	}

	if i.HasGroup != nil {
		p := message.HasGroup()
		if !*i.HasGroup {
			p = message.Not(p)
		}
		predicates = append(predicates, p)
	}
	if len(i.HasGroupWith) > 0 {
		with := make([]predicate.Group, 0, len(i.HasGroupWith))
		for _, w := range i.HasGroupWith {
			p, err := w.P()
			if err != nil {
				return nil, fmt.Errorf("%w: field 'HasGroupWith'", err)
			}
			with = append(with, p)
		}
		predicates = append(predicates, message.HasGroupWith(with...))
	}
	if i.HasUser != nil {
		p := message.HasUser()
		if !*i.HasUser {
			p = message.Not(p)
		}
		predicates = append(predicates, p)
	}
	if len(i.HasUserWith) > 0 {
		with := make([]predicate.User, 0, len(i.HasUserWith))
		for _, w := range i.HasUserWith {
			p, err := w.P()
			if err != nil {
				return nil, fmt.Errorf("%w: field 'HasUserWith'", err)
			}
			with = append(with, p)
		}
		predicates = append(predicates, message.HasUserWith(with...))
	}
	switch len(predicates) {
	case 0:
		return nil, ErrEmptyMessageWhereInput
	case 1:
		return predicates[0], nil
	default:
		return message.And(predicates...), nil
	}
}

// UserWhereInput represents a where input for filtering User queries.
type UserWhereInput struct {
	Predicates []predicate.User  `json:"-"`
	Not        *UserWhereInput   `json:"not,omitempty"`
	Or         []*UserWhereInput `json:"or,omitempty"`
	And        []*UserWhereInput `json:"and,omitempty"`

	// "id" field predicates.
	ID      *int  `json:"id,omitempty"`
	IDNEQ   *int  `json:"idNEQ,omitempty"`
	IDIn    []int `json:"idIn,omitempty"`
	IDNotIn []int `json:"idNotIn,omitempty"`
	IDGT    *int  `json:"idGT,omitempty"`
	IDGTE   *int  `json:"idGTE,omitempty"`
	IDLT    *int  `json:"idLT,omitempty"`
	IDLTE   *int  `json:"idLTE,omitempty"`

	// "created_at" field predicates.
	CreatedAt      *time.Time  `json:"createdAt,omitempty"`
	CreatedAtNEQ   *time.Time  `json:"createdAtNEQ,omitempty"`
	CreatedAtIn    []time.Time `json:"createdAtIn,omitempty"`
	CreatedAtNotIn []time.Time `json:"createdAtNotIn,omitempty"`
	CreatedAtGT    *time.Time  `json:"createdAtGT,omitempty"`
	CreatedAtGTE   *time.Time  `json:"createdAtGTE,omitempty"`
	CreatedAtLT    *time.Time  `json:"createdAtLT,omitempty"`
	CreatedAtLTE   *time.Time  `json:"createdAtLTE,omitempty"`

	// "updated_at" field predicates.
	UpdatedAt      *time.Time  `json:"updatedAt,omitempty"`
	UpdatedAtNEQ   *time.Time  `json:"updatedAtNEQ,omitempty"`
	UpdatedAtIn    []time.Time `json:"updatedAtIn,omitempty"`
	UpdatedAtNotIn []time.Time `json:"updatedAtNotIn,omitempty"`
	UpdatedAtGT    *time.Time  `json:"updatedAtGT,omitempty"`
	UpdatedAtGTE   *time.Time  `json:"updatedAtGTE,omitempty"`
	UpdatedAtLT    *time.Time  `json:"updatedAtLT,omitempty"`
	UpdatedAtLTE   *time.Time  `json:"updatedAtLTE,omitempty"`

	// "username" field predicates.
	Username             *string  `json:"username,omitempty"`
	UsernameNEQ          *string  `json:"usernameNEQ,omitempty"`
	UsernameIn           []string `json:"usernameIn,omitempty"`
	UsernameNotIn        []string `json:"usernameNotIn,omitempty"`
	UsernameGT           *string  `json:"usernameGT,omitempty"`
	UsernameGTE          *string  `json:"usernameGTE,omitempty"`
	UsernameLT           *string  `json:"usernameLT,omitempty"`
	UsernameLTE          *string  `json:"usernameLTE,omitempty"`
	UsernameContains     *string  `json:"usernameContains,omitempty"`
	UsernameHasPrefix    *string  `json:"usernameHasPrefix,omitempty"`
	UsernameHasSuffix    *string  `json:"usernameHasSuffix,omitempty"`
	UsernameEqualFold    *string  `json:"usernameEqualFold,omitempty"`
	UsernameContainsFold *string  `json:"usernameContainsFold,omitempty"`

	// "userGroups" edge predicates.
	HasUserGroups     *bool                  `json:"hasUserGroups,omitempty"`
	HasUserGroupsWith []*UserGroupWhereInput `json:"hasUserGroupsWith,omitempty"`

	// "messages" edge predicates.
	HasMessages     *bool                `json:"hasMessages,omitempty"`
	HasMessagesWith []*MessageWhereInput `json:"hasMessagesWith,omitempty"`
}

// AddPredicates adds custom predicates to the where input to be used during the filtering phase.
func (i *UserWhereInput) AddPredicates(predicates ...predicate.User) {
	i.Predicates = append(i.Predicates, predicates...)
}

// Filter applies the UserWhereInput filter on the UserQuery builder.
func (i *UserWhereInput) Filter(q *UserQuery) (*UserQuery, error) {
	if i == nil {
		return q, nil
	}
	p, err := i.P()
	if err != nil {
		if err == ErrEmptyUserWhereInput {
			return q, nil
		}
		return nil, err
	}
	return q.Where(p), nil
}

// ErrEmptyUserWhereInput is returned in case the UserWhereInput is empty.
var ErrEmptyUserWhereInput = errors.New("ent: empty predicate UserWhereInput")

// P returns a predicate for filtering users.
// An error is returned if the input is empty or invalid.
func (i *UserWhereInput) P() (predicate.User, error) {
	var predicates []predicate.User
	if i.Not != nil {
		p, err := i.Not.P()
		if err != nil {
			return nil, fmt.Errorf("%w: field 'not'", err)
		}
		predicates = append(predicates, user.Not(p))
	}
	switch n := len(i.Or); {
	case n == 1:
		p, err := i.Or[0].P()
		if err != nil {
			return nil, fmt.Errorf("%w: field 'or'", err)
		}
		predicates = append(predicates, p)
	case n > 1:
		or := make([]predicate.User, 0, n)
		for _, w := range i.Or {
			p, err := w.P()
			if err != nil {
				return nil, fmt.Errorf("%w: field 'or'", err)
			}
			or = append(or, p)
		}
		predicates = append(predicates, user.Or(or...))
	}
	switch n := len(i.And); {
	case n == 1:
		p, err := i.And[0].P()
		if err != nil {
			return nil, fmt.Errorf("%w: field 'and'", err)
		}
		predicates = append(predicates, p)
	case n > 1:
		and := make([]predicate.User, 0, n)
		for _, w := range i.And {
			p, err := w.P()
			if err != nil {
				return nil, fmt.Errorf("%w: field 'and'", err)
			}
			and = append(and, p)
		}
		predicates = append(predicates, user.And(and...))
	}
	predicates = append(predicates, i.Predicates...)
	if i.ID != nil {
		predicates = append(predicates, user.IDEQ(*i.ID))
	}
	if i.IDNEQ != nil {
		predicates = append(predicates, user.IDNEQ(*i.IDNEQ))
	}
	if len(i.IDIn) > 0 {
		predicates = append(predicates, user.IDIn(i.IDIn...))
	}
	if len(i.IDNotIn) > 0 {
		predicates = append(predicates, user.IDNotIn(i.IDNotIn...))
	}
	if i.IDGT != nil {
		predicates = append(predicates, user.IDGT(*i.IDGT))
	}
	if i.IDGTE != nil {
		predicates = append(predicates, user.IDGTE(*i.IDGTE))
	}
	if i.IDLT != nil {
		predicates = append(predicates, user.IDLT(*i.IDLT))
	}
	if i.IDLTE != nil {
		predicates = append(predicates, user.IDLTE(*i.IDLTE))
	}
	if i.CreatedAt != nil {
		predicates = append(predicates, user.CreatedAtEQ(*i.CreatedAt))
	}
	if i.CreatedAtNEQ != nil {
		predicates = append(predicates, user.CreatedAtNEQ(*i.CreatedAtNEQ))
	}
	if len(i.CreatedAtIn) > 0 {
		predicates = append(predicates, user.CreatedAtIn(i.CreatedAtIn...))
	}
	if len(i.CreatedAtNotIn) > 0 {
		predicates = append(predicates, user.CreatedAtNotIn(i.CreatedAtNotIn...))
	}
	if i.CreatedAtGT != nil {
		predicates = append(predicates, user.CreatedAtGT(*i.CreatedAtGT))
	}
	if i.CreatedAtGTE != nil {
		predicates = append(predicates, user.CreatedAtGTE(*i.CreatedAtGTE))
	}
	if i.CreatedAtLT != nil {
		predicates = append(predicates, user.CreatedAtLT(*i.CreatedAtLT))
	}
	if i.CreatedAtLTE != nil {
		predicates = append(predicates, user.CreatedAtLTE(*i.CreatedAtLTE))
	}
	if i.UpdatedAt != nil {
		predicates = append(predicates, user.UpdatedAtEQ(*i.UpdatedAt))
	}
	if i.UpdatedAtNEQ != nil {
		predicates = append(predicates, user.UpdatedAtNEQ(*i.UpdatedAtNEQ))
	}
	if len(i.UpdatedAtIn) > 0 {
		predicates = append(predicates, user.UpdatedAtIn(i.UpdatedAtIn...))
	}
	if len(i.UpdatedAtNotIn) > 0 {
		predicates = append(predicates, user.UpdatedAtNotIn(i.UpdatedAtNotIn...))
	}
	if i.UpdatedAtGT != nil {
		predicates = append(predicates, user.UpdatedAtGT(*i.UpdatedAtGT))
	}
	if i.UpdatedAtGTE != nil {
		predicates = append(predicates, user.UpdatedAtGTE(*i.UpdatedAtGTE))
	}
	if i.UpdatedAtLT != nil {
		predicates = append(predicates, user.UpdatedAtLT(*i.UpdatedAtLT))
	}
	if i.UpdatedAtLTE != nil {
		predicates = append(predicates, user.UpdatedAtLTE(*i.UpdatedAtLTE))
	}
	if i.Username != nil {
		predicates = append(predicates, user.UsernameEQ(*i.Username))
	}
	if i.UsernameNEQ != nil {
		predicates = append(predicates, user.UsernameNEQ(*i.UsernameNEQ))
	}
	if len(i.UsernameIn) > 0 {
		predicates = append(predicates, user.UsernameIn(i.UsernameIn...))
	}
	if len(i.UsernameNotIn) > 0 {
		predicates = append(predicates, user.UsernameNotIn(i.UsernameNotIn...))
	}
	if i.UsernameGT != nil {
		predicates = append(predicates, user.UsernameGT(*i.UsernameGT))
	}
	if i.UsernameGTE != nil {
		predicates = append(predicates, user.UsernameGTE(*i.UsernameGTE))
	}
	if i.UsernameLT != nil {
		predicates = append(predicates, user.UsernameLT(*i.UsernameLT))
	}
	if i.UsernameLTE != nil {
		predicates = append(predicates, user.UsernameLTE(*i.UsernameLTE))
	}
	if i.UsernameContains != nil {
		predicates = append(predicates, user.UsernameContains(*i.UsernameContains))
	}
	if i.UsernameHasPrefix != nil {
		predicates = append(predicates, user.UsernameHasPrefix(*i.UsernameHasPrefix))
	}
	if i.UsernameHasSuffix != nil {
		predicates = append(predicates, user.UsernameHasSuffix(*i.UsernameHasSuffix))
	}
	if i.UsernameEqualFold != nil {
		predicates = append(predicates, user.UsernameEqualFold(*i.UsernameEqualFold))
	}
	if i.UsernameContainsFold != nil {
		predicates = append(predicates, user.UsernameContainsFold(*i.UsernameContainsFold))
	}

	if i.HasUserGroups != nil {
		p := user.HasUserGroups()
		if !*i.HasUserGroups {
			p = user.Not(p)
		}
		predicates = append(predicates, p)
	}
	if len(i.HasUserGroupsWith) > 0 {
		with := make([]predicate.UserGroup, 0, len(i.HasUserGroupsWith))
		for _, w := range i.HasUserGroupsWith {
			p, err := w.P()
			if err != nil {
				return nil, fmt.Errorf("%w: field 'HasUserGroupsWith'", err)
			}
			with = append(with, p)
		}
		predicates = append(predicates, user.HasUserGroupsWith(with...))
	}
	if i.HasMessages != nil {
		p := user.HasMessages()
		if !*i.HasMessages {
			p = user.Not(p)
		}
		predicates = append(predicates, p)
	}
	if len(i.HasMessagesWith) > 0 {
		with := make([]predicate.Message, 0, len(i.HasMessagesWith))
		for _, w := range i.HasMessagesWith {
			p, err := w.P()
			if err != nil {
				return nil, fmt.Errorf("%w: field 'HasMessagesWith'", err)
			}
			with = append(with, p)
		}
		predicates = append(predicates, user.HasMessagesWith(with...))
	}
	switch len(predicates) {
	case 0:
		return nil, ErrEmptyUserWhereInput
	case 1:
		return predicates[0], nil
	default:
		return user.And(predicates...), nil
	}
}

// UserGroupWhereInput represents a where input for filtering UserGroup queries.
type UserGroupWhereInput struct {
	Predicates []predicate.UserGroup  `json:"-"`
	Not        *UserGroupWhereInput   `json:"not,omitempty"`
	Or         []*UserGroupWhereInput `json:"or,omitempty"`
	And        []*UserGroupWhereInput `json:"and,omitempty"`

	// "id" field predicates.
	ID      *int  `json:"id,omitempty"`
	IDNEQ   *int  `json:"idNEQ,omitempty"`
	IDIn    []int `json:"idIn,omitempty"`
	IDNotIn []int `json:"idNotIn,omitempty"`
	IDGT    *int  `json:"idGT,omitempty"`
	IDGTE   *int  `json:"idGTE,omitempty"`
	IDLT    *int  `json:"idLT,omitempty"`
	IDLTE   *int  `json:"idLTE,omitempty"`

	// "created_at" field predicates.
	CreatedAt      *time.Time  `json:"createdAt,omitempty"`
	CreatedAtNEQ   *time.Time  `json:"createdAtNEQ,omitempty"`
	CreatedAtIn    []time.Time `json:"createdAtIn,omitempty"`
	CreatedAtNotIn []time.Time `json:"createdAtNotIn,omitempty"`
	CreatedAtGT    *time.Time  `json:"createdAtGT,omitempty"`
	CreatedAtGTE   *time.Time  `json:"createdAtGTE,omitempty"`
	CreatedAtLT    *time.Time  `json:"createdAtLT,omitempty"`
	CreatedAtLTE   *time.Time  `json:"createdAtLTE,omitempty"`

	// "group" edge predicates.
	HasGroup     *bool              `json:"hasGroup,omitempty"`
	HasGroupWith []*GroupWhereInput `json:"hasGroupWith,omitempty"`

	// "user" edge predicates.
	HasUser     *bool             `json:"hasUser,omitempty"`
	HasUserWith []*UserWhereInput `json:"hasUserWith,omitempty"`
}

// AddPredicates adds custom predicates to the where input to be used during the filtering phase.
func (i *UserGroupWhereInput) AddPredicates(predicates ...predicate.UserGroup) {
	i.Predicates = append(i.Predicates, predicates...)
}

// Filter applies the UserGroupWhereInput filter on the UserGroupQuery builder.
func (i *UserGroupWhereInput) Filter(q *UserGroupQuery) (*UserGroupQuery, error) {
	if i == nil {
		return q, nil
	}
	p, err := i.P()
	if err != nil {
		if err == ErrEmptyUserGroupWhereInput {
			return q, nil
		}
		return nil, err
	}
	return q.Where(p), nil
}

// ErrEmptyUserGroupWhereInput is returned in case the UserGroupWhereInput is empty.
var ErrEmptyUserGroupWhereInput = errors.New("ent: empty predicate UserGroupWhereInput")

// P returns a predicate for filtering usergroups.
// An error is returned if the input is empty or invalid.
func (i *UserGroupWhereInput) P() (predicate.UserGroup, error) {
	var predicates []predicate.UserGroup
	if i.Not != nil {
		p, err := i.Not.P()
		if err != nil {
			return nil, fmt.Errorf("%w: field 'not'", err)
		}
		predicates = append(predicates, usergroup.Not(p))
	}
	switch n := len(i.Or); {
	case n == 1:
		p, err := i.Or[0].P()
		if err != nil {
			return nil, fmt.Errorf("%w: field 'or'", err)
		}
		predicates = append(predicates, p)
	case n > 1:
		or := make([]predicate.UserGroup, 0, n)
		for _, w := range i.Or {
			p, err := w.P()
			if err != nil {
				return nil, fmt.Errorf("%w: field 'or'", err)
			}
			or = append(or, p)
		}
		predicates = append(predicates, usergroup.Or(or...))
	}
	switch n := len(i.And); {
	case n == 1:
		p, err := i.And[0].P()
		if err != nil {
			return nil, fmt.Errorf("%w: field 'and'", err)
		}
		predicates = append(predicates, p)
	case n > 1:
		and := make([]predicate.UserGroup, 0, n)
		for _, w := range i.And {
			p, err := w.P()
			if err != nil {
				return nil, fmt.Errorf("%w: field 'and'", err)
			}
			and = append(and, p)
		}
		predicates = append(predicates, usergroup.And(and...))
	}
	predicates = append(predicates, i.Predicates...)
	if i.ID != nil {
		predicates = append(predicates, usergroup.IDEQ(*i.ID))
	}
	if i.IDNEQ != nil {
		predicates = append(predicates, usergroup.IDNEQ(*i.IDNEQ))
	}
	if len(i.IDIn) > 0 {
		predicates = append(predicates, usergroup.IDIn(i.IDIn...))
	}
	if len(i.IDNotIn) > 0 {
		predicates = append(predicates, usergroup.IDNotIn(i.IDNotIn...))
	}
	if i.IDGT != nil {
		predicates = append(predicates, usergroup.IDGT(*i.IDGT))
	}
	if i.IDGTE != nil {
		predicates = append(predicates, usergroup.IDGTE(*i.IDGTE))
	}
	if i.IDLT != nil {
		predicates = append(predicates, usergroup.IDLT(*i.IDLT))
	}
	if i.IDLTE != nil {
		predicates = append(predicates, usergroup.IDLTE(*i.IDLTE))
	}
	if i.CreatedAt != nil {
		predicates = append(predicates, usergroup.CreatedAtEQ(*i.CreatedAt))
	}
	if i.CreatedAtNEQ != nil {
		predicates = append(predicates, usergroup.CreatedAtNEQ(*i.CreatedAtNEQ))
	}
	if len(i.CreatedAtIn) > 0 {
		predicates = append(predicates, usergroup.CreatedAtIn(i.CreatedAtIn...))
	}
	if len(i.CreatedAtNotIn) > 0 {
		predicates = append(predicates, usergroup.CreatedAtNotIn(i.CreatedAtNotIn...))
	}
	if i.CreatedAtGT != nil {
		predicates = append(predicates, usergroup.CreatedAtGT(*i.CreatedAtGT))
	}
	if i.CreatedAtGTE != nil {
		predicates = append(predicates, usergroup.CreatedAtGTE(*i.CreatedAtGTE))
	}
	if i.CreatedAtLT != nil {
		predicates = append(predicates, usergroup.CreatedAtLT(*i.CreatedAtLT))
	}
	if i.CreatedAtLTE != nil {
		predicates = append(predicates, usergroup.CreatedAtLTE(*i.CreatedAtLTE))
	}

	if i.HasGroup != nil {
		p := usergroup.HasGroup()
		if !*i.HasGroup {
			p = usergroup.Not(p)
		}
		predicates = append(predicates, p)
	}
	if len(i.HasGroupWith) > 0 {
		with := make([]predicate.Group, 0, len(i.HasGroupWith))
		for _, w := range i.HasGroupWith {
			p, err := w.P()
			if err != nil {
				return nil, fmt.Errorf("%w: field 'HasGroupWith'", err)
			}
			with = append(with, p)
		}
		predicates = append(predicates, usergroup.HasGroupWith(with...))
	}
	if i.HasUser != nil {
		p := usergroup.HasUser()
		if !*i.HasUser {
			p = usergroup.Not(p)
		}
		predicates = append(predicates, p)
	}
	if len(i.HasUserWith) > 0 {
		with := make([]predicate.User, 0, len(i.HasUserWith))
		for _, w := range i.HasUserWith {
			p, err := w.P()
			if err != nil {
				return nil, fmt.Errorf("%w: field 'HasUserWith'", err)
			}
			with = append(with, p)
		}
		predicates = append(predicates, usergroup.HasUserWith(with...))
	}
	switch len(predicates) {
	case 0:
		return nil, ErrEmptyUserGroupWhereInput
	case 1:
		return predicates[0], nil
	default:
		return usergroup.And(predicates...), nil
	}
}
