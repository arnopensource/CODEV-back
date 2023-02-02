// Code generated by ent, DO NOT EDIT.

package member

const (
	// Label holds the string label denoting the member type in the database.
	Label = "member"
	// FieldProfileID holds the string denoting the profile_id field in the database.
	FieldProfileID = "profile_id"
	// FieldEventID holds the string denoting the event_id field in the database.
	FieldEventID = "event_id"
	// FieldIsAdmin holds the string denoting the is_admin field in the database.
	FieldIsAdmin = "is_admin"
	// EdgeProfile holds the string denoting the profile edge name in mutations.
	EdgeProfile = "profile"
	// EdgeEvent holds the string denoting the event edge name in mutations.
	EdgeEvent = "event"
	// ProfileFieldID holds the string denoting the ID field of the Profile.
	ProfileFieldID = "id"
	// EventFieldID holds the string denoting the ID field of the Event.
	EventFieldID = "id"
	// Table holds the table name of the member in the database.
	Table = "members"
	// ProfileTable is the table that holds the profile relation/edge.
	ProfileTable = "members"
	// ProfileInverseTable is the table name for the Profile entity.
	// It exists in this package in order to avoid circular dependency with the "profile" package.
	ProfileInverseTable = "profiles"
	// ProfileColumn is the table column denoting the profile relation/edge.
	ProfileColumn = "profile_id"
	// EventTable is the table that holds the event relation/edge.
	EventTable = "members"
	// EventInverseTable is the table name for the Event entity.
	// It exists in this package in order to avoid circular dependency with the "event" package.
	EventInverseTable = "events"
	// EventColumn is the table column denoting the event relation/edge.
	EventColumn = "event_id"
)

// Columns holds all SQL columns for member fields.
var Columns = []string{
	FieldProfileID,
	FieldEventID,
	FieldIsAdmin,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultIsAdmin holds the default value on creation for the "is_admin" field.
	DefaultIsAdmin bool
)
