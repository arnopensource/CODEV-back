// Code generated by ent, DO NOT EDIT.

package ent

import (
	"time"

	"github.com/abc3354/CODEV-back/ent/eventinvite"
	"github.com/abc3354/CODEV-back/ent/member"
	"github.com/abc3354/CODEV-back/ent/schema"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	eventinviteFields := schema.EventInvite{}.Fields()
	_ = eventinviteFields
	// eventinviteDescSince is the schema descriptor for since field.
	eventinviteDescSince := eventinviteFields[2].Descriptor()
	// eventinvite.DefaultSince holds the default value on creation for the since field.
	eventinvite.DefaultSince = eventinviteDescSince.Default.(func() time.Time)
	memberFields := schema.Member{}.Fields()
	_ = memberFields
	// memberDescIsAdmin is the schema descriptor for is_admin field.
	memberDescIsAdmin := memberFields[2].Descriptor()
	// member.DefaultIsAdmin holds the default value on creation for the is_admin field.
	member.DefaultIsAdmin = memberDescIsAdmin.Default.(bool)
}
