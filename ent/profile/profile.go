// Code generated by ent, DO NOT EDIT.

package profile

const (
	// Label holds the string label denoting the profile type in the database.
	Label = "profile"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldFirstname holds the string denoting the firstname field in the database.
	FieldFirstname = "firstname"
	// FieldLastname holds the string denoting the lastname field in the database.
	FieldLastname = "lastname"
	// FieldPhone holds the string denoting the phone field in the database.
	FieldPhone = "phone"
	// FieldEmail holds the string denoting the email field in the database.
	FieldEmail = "email"
	// EdgeFriends holds the string denoting the friends edge name in mutations.
	EdgeFriends = "friends"
	// EdgeBookings holds the string denoting the bookings edge name in mutations.
	EdgeBookings = "bookings"
	// EdgeEvents holds the string denoting the events edge name in mutations.
	EdgeEvents = "events"
	// EdgeInvitedTo holds the string denoting the invitedto edge name in mutations.
	EdgeInvitedTo = "invitedTo"
	// EdgeFriendsData holds the string denoting the friends_data edge name in mutations.
	EdgeFriendsData = "friends_data"
	// EdgeBookingsData holds the string denoting the bookings_data edge name in mutations.
	EdgeBookingsData = "bookings_data"
	// EdgeEventsData holds the string denoting the events_data edge name in mutations.
	EdgeEventsData = "events_data"
	// EdgeInvitesData holds the string denoting the invites_data edge name in mutations.
	EdgeInvitesData = "invites_data"
	// Table holds the table name of the profile in the database.
	Table = "profiles"
	// FriendsTable is the table that holds the friends relation/edge. The primary key declared below.
	FriendsTable = "friends"
	// BookingsTable is the table that holds the bookings relation/edge. The primary key declared below.
	BookingsTable = "bookings"
	// BookingsInverseTable is the table name for the Room entity.
	// It exists in this package in order to avoid circular dependency with the "room" package.
	BookingsInverseTable = "rooms"
	// EventsTable is the table that holds the events relation/edge. The primary key declared below.
	EventsTable = "members"
	// EventsInverseTable is the table name for the Event entity.
	// It exists in this package in order to avoid circular dependency with the "event" package.
	EventsInverseTable = "events"
	// InvitedToTable is the table that holds the invitedTo relation/edge. The primary key declared below.
	InvitedToTable = "event_invites"
	// InvitedToInverseTable is the table name for the Event entity.
	// It exists in this package in order to avoid circular dependency with the "event" package.
	InvitedToInverseTable = "events"
	// FriendsDataTable is the table that holds the friends_data relation/edge.
	FriendsDataTable = "friends"
	// FriendsDataInverseTable is the table name for the Friend entity.
	// It exists in this package in order to avoid circular dependency with the "friend" package.
	FriendsDataInverseTable = "friends"
	// FriendsDataColumn is the table column denoting the friends_data relation/edge.
	FriendsDataColumn = "profile_id"
	// BookingsDataTable is the table that holds the bookings_data relation/edge.
	BookingsDataTable = "bookings"
	// BookingsDataInverseTable is the table name for the Booking entity.
	// It exists in this package in order to avoid circular dependency with the "booking" package.
	BookingsDataInverseTable = "bookings"
	// BookingsDataColumn is the table column denoting the bookings_data relation/edge.
	BookingsDataColumn = "profile_id"
	// EventsDataTable is the table that holds the events_data relation/edge.
	EventsDataTable = "members"
	// EventsDataInverseTable is the table name for the Member entity.
	// It exists in this package in order to avoid circular dependency with the "member" package.
	EventsDataInverseTable = "members"
	// EventsDataColumn is the table column denoting the events_data relation/edge.
	EventsDataColumn = "profile_id"
	// InvitesDataTable is the table that holds the invites_data relation/edge.
	InvitesDataTable = "event_invites"
	// InvitesDataInverseTable is the table name for the EventInvite entity.
	// It exists in this package in order to avoid circular dependency with the "eventinvite" package.
	InvitesDataInverseTable = "event_invites"
	// InvitesDataColumn is the table column denoting the invites_data relation/edge.
	InvitesDataColumn = "profile_id"
)

// Columns holds all SQL columns for profile fields.
var Columns = []string{
	FieldID,
	FieldFirstname,
	FieldLastname,
	FieldPhone,
	FieldEmail,
}

var (
	// FriendsPrimaryKey and FriendsColumn2 are the table columns denoting the
	// primary key for the friends relation (M2M).
	FriendsPrimaryKey = []string{"profile_id", "friend_id"}
	// BookingsPrimaryKey and BookingsColumn2 are the table columns denoting the
	// primary key for the bookings relation (M2M).
	BookingsPrimaryKey = []string{"profile_id", "room_id"}
	// EventsPrimaryKey and EventsColumn2 are the table columns denoting the
	// primary key for the events relation (M2M).
	EventsPrimaryKey = []string{"event_id", "profile_id"}
	// InvitedToPrimaryKey and InvitedToColumn2 are the table columns denoting the
	// primary key for the invitedTo relation (M2M).
	InvitedToPrimaryKey = []string{"event_id", "profile_id"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}
