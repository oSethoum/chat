// Code generated by ent, DO NOT EDIT.

package ent

import (
	"time"
)

// CreateGroupInput represents a mutation input for creating groups.
type CreateGroupInput struct {
	Name         string
	UserGroupIDs []int
	MessageIDs   []int
}

// Mutate applies the CreateGroupInput on the GroupMutation builder.
func (i *CreateGroupInput) Mutate(m *GroupMutation) {
	m.SetName(i.Name)
	if v := i.UserGroupIDs; len(v) > 0 {
		m.AddUserGroupIDs(v...)
	}
	if v := i.MessageIDs; len(v) > 0 {
		m.AddMessageIDs(v...)
	}
}

// SetInput applies the change-set in the CreateGroupInput on the GroupCreate builder.
func (c *GroupCreate) SetInput(i CreateGroupInput) *GroupCreate {
	i.Mutate(c.Mutation())
	return c
}

// UpdateGroupInput represents a mutation input for updating groups.
type UpdateGroupInput struct {
	Name               *string
	AddUserGroupIDs    []int
	RemoveUserGroupIDs []int
	AddMessageIDs      []int
	RemoveMessageIDs   []int
}

// Mutate applies the UpdateGroupInput on the GroupMutation builder.
func (i *UpdateGroupInput) Mutate(m *GroupMutation) {
	if v := i.Name; v != nil {
		m.SetName(*v)
	}
	if v := i.AddUserGroupIDs; len(v) > 0 {
		m.AddUserGroupIDs(v...)
	}
	if v := i.RemoveUserGroupIDs; len(v) > 0 {
		m.RemoveUserGroupIDs(v...)
	}
	if v := i.AddMessageIDs; len(v) > 0 {
		m.AddMessageIDs(v...)
	}
	if v := i.RemoveMessageIDs; len(v) > 0 {
		m.RemoveMessageIDs(v...)
	}
}

// SetInput applies the change-set in the UpdateGroupInput on the GroupUpdate builder.
func (c *GroupUpdate) SetInput(i UpdateGroupInput) *GroupUpdate {
	i.Mutate(c.Mutation())
	return c
}

// SetInput applies the change-set in the UpdateGroupInput on the GroupUpdateOne builder.
func (c *GroupUpdateOne) SetInput(i UpdateGroupInput) *GroupUpdateOne {
	i.Mutate(c.Mutation())
	return c
}

// CreateMessageInput represents a mutation input for creating messages.
type CreateMessageInput struct {
	Message string
	GroupID *int
	UserID  *int
}

// Mutate applies the CreateMessageInput on the MessageMutation builder.
func (i *CreateMessageInput) Mutate(m *MessageMutation) {
	m.SetMessage(i.Message)
	if v := i.GroupID; v != nil {
		m.SetGroupID(*v)
	}
	if v := i.UserID; v != nil {
		m.SetUserID(*v)
	}
}

// SetInput applies the change-set in the CreateMessageInput on the MessageCreate builder.
func (c *MessageCreate) SetInput(i CreateMessageInput) *MessageCreate {
	i.Mutate(c.Mutation())
	return c
}

// UpdateMessageInput represents a mutation input for updating messages.
type UpdateMessageInput struct {
	SentAt     *time.Time
	Message    *string
	ClearGroup bool
	GroupID    *int
	ClearUser  bool
	UserID     *int
}

// Mutate applies the UpdateMessageInput on the MessageMutation builder.
func (i *UpdateMessageInput) Mutate(m *MessageMutation) {
	if v := i.SentAt; v != nil {
		m.SetSentAt(*v)
	}
	if v := i.Message; v != nil {
		m.SetMessage(*v)
	}
	if i.ClearGroup {
		m.ClearGroup()
	}
	if v := i.GroupID; v != nil {
		m.SetGroupID(*v)
	}
	if i.ClearUser {
		m.ClearUser()
	}
	if v := i.UserID; v != nil {
		m.SetUserID(*v)
	}
}

// SetInput applies the change-set in the UpdateMessageInput on the MessageUpdate builder.
func (c *MessageUpdate) SetInput(i UpdateMessageInput) *MessageUpdate {
	i.Mutate(c.Mutation())
	return c
}

// SetInput applies the change-set in the UpdateMessageInput on the MessageUpdateOne builder.
func (c *MessageUpdateOne) SetInput(i UpdateMessageInput) *MessageUpdateOne {
	i.Mutate(c.Mutation())
	return c
}

// CreateUserInput represents a mutation input for creating users.
type CreateUserInput struct {
	Username     string
	Password     string
	UserGroupIDs []int
	MessageIDs   []int
}

// Mutate applies the CreateUserInput on the UserMutation builder.
func (i *CreateUserInput) Mutate(m *UserMutation) {
	m.SetUsername(i.Username)
	m.SetPassword(i.Password)
	if v := i.UserGroupIDs; len(v) > 0 {
		m.AddUserGroupIDs(v...)
	}
	if v := i.MessageIDs; len(v) > 0 {
		m.AddMessageIDs(v...)
	}
}

// SetInput applies the change-set in the CreateUserInput on the UserCreate builder.
func (c *UserCreate) SetInput(i CreateUserInput) *UserCreate {
	i.Mutate(c.Mutation())
	return c
}

// UpdateUserInput represents a mutation input for updating users.
type UpdateUserInput struct {
	Username           *string
	Password           *string
	AddUserGroupIDs    []int
	RemoveUserGroupIDs []int
	AddMessageIDs      []int
	RemoveMessageIDs   []int
}

// Mutate applies the UpdateUserInput on the UserMutation builder.
func (i *UpdateUserInput) Mutate(m *UserMutation) {
	if v := i.Username; v != nil {
		m.SetUsername(*v)
	}
	if v := i.Password; v != nil {
		m.SetPassword(*v)
	}
	if v := i.AddUserGroupIDs; len(v) > 0 {
		m.AddUserGroupIDs(v...)
	}
	if v := i.RemoveUserGroupIDs; len(v) > 0 {
		m.RemoveUserGroupIDs(v...)
	}
	if v := i.AddMessageIDs; len(v) > 0 {
		m.AddMessageIDs(v...)
	}
	if v := i.RemoveMessageIDs; len(v) > 0 {
		m.RemoveMessageIDs(v...)
	}
}

// SetInput applies the change-set in the UpdateUserInput on the UserUpdate builder.
func (c *UserUpdate) SetInput(i UpdateUserInput) *UserUpdate {
	i.Mutate(c.Mutation())
	return c
}

// SetInput applies the change-set in the UpdateUserInput on the UserUpdateOne builder.
func (c *UserUpdateOne) SetInput(i UpdateUserInput) *UserUpdateOne {
	i.Mutate(c.Mutation())
	return c
}

// CreateUserGroupInput represents a mutation input for creating usergroups.
type CreateUserGroupInput struct {
	GroupID *int
	UserID  *int
}

// Mutate applies the CreateUserGroupInput on the UserGroupMutation builder.
func (i *CreateUserGroupInput) Mutate(m *UserGroupMutation) {
	if v := i.GroupID; v != nil {
		m.SetGroupID(*v)
	}
	if v := i.UserID; v != nil {
		m.SetUserID(*v)
	}
}

// SetInput applies the change-set in the CreateUserGroupInput on the UserGroupCreate builder.
func (c *UserGroupCreate) SetInput(i CreateUserGroupInput) *UserGroupCreate {
	i.Mutate(c.Mutation())
	return c
}

// UpdateUserGroupInput represents a mutation input for updating usergroups.
type UpdateUserGroupInput struct {
	ClearGroup bool
	GroupID    *int
	ClearUser  bool
	UserID     *int
}

// Mutate applies the UpdateUserGroupInput on the UserGroupMutation builder.
func (i *UpdateUserGroupInput) Mutate(m *UserGroupMutation) {
	if i.ClearGroup {
		m.ClearGroup()
	}
	if v := i.GroupID; v != nil {
		m.SetGroupID(*v)
	}
	if i.ClearUser {
		m.ClearUser()
	}
	if v := i.UserID; v != nil {
		m.SetUserID(*v)
	}
}

// SetInput applies the change-set in the UpdateUserGroupInput on the UserGroupUpdate builder.
func (c *UserGroupUpdate) SetInput(i UpdateUserGroupInput) *UserGroupUpdate {
	i.Mutate(c.Mutation())
	return c
}

// SetInput applies the change-set in the UpdateUserGroupInput on the UserGroupUpdateOne builder.
func (c *UserGroupUpdateOne) SetInput(i UpdateUserGroupInput) *UserGroupUpdateOne {
	i.Mutate(c.Mutation())
	return c
}