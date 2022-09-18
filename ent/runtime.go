// Code generated by ent, DO NOT EDIT.

package ent

import (
	"chat/ent/group"
	"chat/ent/message"
	"chat/ent/schema"
	"chat/ent/user"
	"chat/ent/usergroup"
	"time"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	groupFields := schema.Group{}.Fields()
	_ = groupFields
	// groupDescUpdatedAt is the schema descriptor for updated_at field.
	groupDescUpdatedAt := groupFields[0].Descriptor()
	// group.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	group.DefaultUpdatedAt = groupDescUpdatedAt.Default.(func() time.Time)
	// group.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	group.UpdateDefaultUpdatedAt = groupDescUpdatedAt.UpdateDefault.(func() time.Time)
	// groupDescCreatedAt is the schema descriptor for created_at field.
	groupDescCreatedAt := groupFields[1].Descriptor()
	// group.DefaultCreatedAt holds the default value on creation for the created_at field.
	group.DefaultCreatedAt = groupDescCreatedAt.Default.(func() time.Time)
	messageFields := schema.Message{}.Fields()
	_ = messageFields
	// messageDescSentAt is the schema descriptor for sent_at field.
	messageDescSentAt := messageFields[0].Descriptor()
	// message.DefaultSentAt holds the default value on creation for the sent_at field.
	message.DefaultSentAt = messageDescSentAt.Default.(func() time.Time)
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescCreatedAt is the schema descriptor for created_at field.
	userDescCreatedAt := userFields[0].Descriptor()
	// user.DefaultCreatedAt holds the default value on creation for the created_at field.
	user.DefaultCreatedAt = userDescCreatedAt.Default.(func() time.Time)
	// userDescUpdatedAt is the schema descriptor for updated_at field.
	userDescUpdatedAt := userFields[1].Descriptor()
	// user.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	user.DefaultUpdatedAt = userDescUpdatedAt.Default.(func() time.Time)
	// user.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	user.UpdateDefaultUpdatedAt = userDescUpdatedAt.UpdateDefault.(func() time.Time)
	usergroupFields := schema.UserGroup{}.Fields()
	_ = usergroupFields
	// usergroupDescCreatedAt is the schema descriptor for created_at field.
	usergroupDescCreatedAt := usergroupFields[0].Descriptor()
	// usergroup.DefaultCreatedAt holds the default value on creation for the created_at field.
	usergroup.DefaultCreatedAt = usergroupDescCreatedAt.Default.(func() time.Time)
}
