// Code generated by ent, DO NOT EDIT.

package event

const (
	// Label holds the string label denoting the event type in the database.
	Label = "event"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldActivity holds the string denoting the activity field in the database.
	FieldActivity = "activity"
	// FieldStart holds the string denoting the start field in the database.
	FieldStart = "start"
	// FieldEnd holds the string denoting the end field in the database.
	FieldEnd = "end"
	// EdgeProfiles holds the string denoting the profiles edge name in mutations.
	EdgeProfiles = "profiles"
	// EdgeRoom holds the string denoting the room edge name in mutations.
	EdgeRoom = "room"
	// EdgeInvited holds the string denoting the invited edge name in mutations.
	EdgeInvited = "invited"
	// EdgeMembers holds the string denoting the members edge name in mutations.
	EdgeMembers = "members"
	// EdgeInvites holds the string denoting the invites edge name in mutations.
	EdgeInvites = "invites"
	// Table holds the table name of the event in the database.
	Table = "events"
	// ProfilesTable is the table that holds the profiles relation/edge. The primary key declared below.
	ProfilesTable = "members"
	// ProfilesInverseTable is the table name for the Profile entity.
	// It exists in this package in order to avoid circular dependency with the "profile" package.
	ProfilesInverseTable = "profiles"
	// RoomTable is the table that holds the room relation/edge.
	RoomTable = "events"
	// RoomInverseTable is the table name for the Room entity.
	// It exists in this package in order to avoid circular dependency with the "room" package.
	RoomInverseTable = "rooms"
	// RoomColumn is the table column denoting the room relation/edge.
	RoomColumn = "room_events"
	// InvitedTable is the table that holds the invited relation/edge. The primary key declared below.
	InvitedTable = "event_invites"
	// InvitedInverseTable is the table name for the Profile entity.
	// It exists in this package in order to avoid circular dependency with the "profile" package.
	InvitedInverseTable = "profiles"
	// MembersTable is the table that holds the members relation/edge.
	MembersTable = "members"
	// MembersInverseTable is the table name for the Member entity.
	// It exists in this package in order to avoid circular dependency with the "member" package.
	MembersInverseTable = "members"
	// MembersColumn is the table column denoting the members relation/edge.
	MembersColumn = "event_id"
	// InvitesTable is the table that holds the invites relation/edge.
	InvitesTable = "event_invites"
	// InvitesInverseTable is the table name for the EventInvite entity.
	// It exists in this package in order to avoid circular dependency with the "eventinvite" package.
	InvitesInverseTable = "event_invites"
	// InvitesColumn is the table column denoting the invites relation/edge.
	InvitesColumn = "event_id"
)

// Columns holds all SQL columns for event fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldActivity,
	FieldStart,
	FieldEnd,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "events"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"room_events",
}

var (
	// ProfilesPrimaryKey and ProfilesColumn2 are the table columns denoting the
	// primary key for the profiles relation (M2M).
	ProfilesPrimaryKey = []string{"event_id", "profile_id"}
	// InvitedPrimaryKey and InvitedColumn2 are the table columns denoting the
	// primary key for the invited relation (M2M).
	InvitedPrimaryKey = []string{"event_id", "profile_id"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}
