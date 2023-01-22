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
	// EdgeFriends holds the string denoting the friends edge name in mutations.
	EdgeFriends = "friends"
	// EdgeBookings holds the string denoting the bookings edge name in mutations.
	EdgeBookings = "bookings"
	// EdgeFriendsData holds the string denoting the friends_data edge name in mutations.
	EdgeFriendsData = "friends_data"
	// EdgeBookingsData holds the string denoting the bookings_data edge name in mutations.
	EdgeBookingsData = "bookings_data"
	// Table holds the table name of the profile in the database.
	Table = "profiles"
	// FriendsTable is the table that holds the friends relation/edge. The primary key declared below.
	FriendsTable = "friends"
	// BookingsTable is the table that holds the bookings relation/edge. The primary key declared below.
	BookingsTable = "bookings"
	// BookingsInverseTable is the table name for the Room entity.
	// It exists in this package in order to avoid circular dependency with the "room" package.
	BookingsInverseTable = "rooms"
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
)

// Columns holds all SQL columns for profile fields.
var Columns = []string{
	FieldID,
	FieldFirstname,
	FieldLastname,
	FieldPhone,
}

var (
	// FriendsPrimaryKey and FriendsColumn2 are the table columns denoting the
	// primary key for the friends relation (M2M).
	FriendsPrimaryKey = []string{"profile_id", "friend_id"}
	// BookingsPrimaryKey and BookingsColumn2 are the table columns denoting the
	// primary key for the bookings relation (M2M).
	BookingsPrimaryKey = []string{"profile_id", "room_id"}
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
