// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/abc3354/CODEV-back/ent/profile"
	"github.com/google/uuid"
)

// Profile is the model entity for the Profile schema.
type Profile struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// Firstname holds the value of the "firstname" field.
	Firstname string `json:"firstname,omitempty"`
	// Lastname holds the value of the "lastname" field.
	Lastname string `json:"lastname,omitempty"`
	// Phone holds the value of the "phone" field.
	Phone string `json:"phone,omitempty"`
	// Email holds the value of the "email" field.
	Email string `json:"email,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ProfileQuery when eager-loading is set.
	Edges ProfileEdges `json:"-"`
}

// ProfileEdges holds the relations/edges for other nodes in the graph.
type ProfileEdges struct {
	// Friends holds the value of the friends edge.
	Friends []*Profile `json:"friends,omitempty"`
	// Bookings holds the value of the bookings edge.
	Bookings []*Room `json:"bookings,omitempty"`
	// Events holds the value of the events edge.
	Events []*Event `json:"events,omitempty"`
	// InvitedTo holds the value of the invitedTo edge.
	InvitedTo []*Event `json:"invitedTo,omitempty"`
	// FriendsData holds the value of the friends_data edge.
	FriendsData []*Friend `json:"friends_data,omitempty"`
	// BookingsData holds the value of the bookings_data edge.
	BookingsData []*Booking `json:"bookings_data,omitempty"`
	// EventsData holds the value of the events_data edge.
	EventsData []*Member `json:"events_data,omitempty"`
	// InvitesData holds the value of the invites_data edge.
	InvitesData []*EventInvite `json:"invites_data,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [8]bool
}

// FriendsOrErr returns the Friends value or an error if the edge
// was not loaded in eager-loading.
func (e ProfileEdges) FriendsOrErr() ([]*Profile, error) {
	if e.loadedTypes[0] {
		return e.Friends, nil
	}
	return nil, &NotLoadedError{edge: "friends"}
}

// BookingsOrErr returns the Bookings value or an error if the edge
// was not loaded in eager-loading.
func (e ProfileEdges) BookingsOrErr() ([]*Room, error) {
	if e.loadedTypes[1] {
		return e.Bookings, nil
	}
	return nil, &NotLoadedError{edge: "bookings"}
}

// EventsOrErr returns the Events value or an error if the edge
// was not loaded in eager-loading.
func (e ProfileEdges) EventsOrErr() ([]*Event, error) {
	if e.loadedTypes[2] {
		return e.Events, nil
	}
	return nil, &NotLoadedError{edge: "events"}
}

// InvitedToOrErr returns the InvitedTo value or an error if the edge
// was not loaded in eager-loading.
func (e ProfileEdges) InvitedToOrErr() ([]*Event, error) {
	if e.loadedTypes[3] {
		return e.InvitedTo, nil
	}
	return nil, &NotLoadedError{edge: "invitedTo"}
}

// FriendsDataOrErr returns the FriendsData value or an error if the edge
// was not loaded in eager-loading.
func (e ProfileEdges) FriendsDataOrErr() ([]*Friend, error) {
	if e.loadedTypes[4] {
		return e.FriendsData, nil
	}
	return nil, &NotLoadedError{edge: "friends_data"}
}

// BookingsDataOrErr returns the BookingsData value or an error if the edge
// was not loaded in eager-loading.
func (e ProfileEdges) BookingsDataOrErr() ([]*Booking, error) {
	if e.loadedTypes[5] {
		return e.BookingsData, nil
	}
	return nil, &NotLoadedError{edge: "bookings_data"}
}

// EventsDataOrErr returns the EventsData value or an error if the edge
// was not loaded in eager-loading.
func (e ProfileEdges) EventsDataOrErr() ([]*Member, error) {
	if e.loadedTypes[6] {
		return e.EventsData, nil
	}
	return nil, &NotLoadedError{edge: "events_data"}
}

// InvitesDataOrErr returns the InvitesData value or an error if the edge
// was not loaded in eager-loading.
func (e ProfileEdges) InvitesDataOrErr() ([]*EventInvite, error) {
	if e.loadedTypes[7] {
		return e.InvitesData, nil
	}
	return nil, &NotLoadedError{edge: "invites_data"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Profile) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case profile.FieldFirstname, profile.FieldLastname, profile.FieldPhone, profile.FieldEmail:
			values[i] = new(sql.NullString)
		case profile.FieldID:
			values[i] = new(uuid.UUID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Profile", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Profile fields.
func (pr *Profile) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case profile.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				pr.ID = *value
			}
		case profile.FieldFirstname:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field firstname", values[i])
			} else if value.Valid {
				pr.Firstname = value.String
			}
		case profile.FieldLastname:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field lastname", values[i])
			} else if value.Valid {
				pr.Lastname = value.String
			}
		case profile.FieldPhone:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field phone", values[i])
			} else if value.Valid {
				pr.Phone = value.String
			}
		case profile.FieldEmail:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field email", values[i])
			} else if value.Valid {
				pr.Email = value.String
			}
		}
	}
	return nil
}

// QueryFriends queries the "friends" edge of the Profile entity.
func (pr *Profile) QueryFriends() *ProfileQuery {
	return (&ProfileClient{config: pr.config}).QueryFriends(pr)
}

// QueryBookings queries the "bookings" edge of the Profile entity.
func (pr *Profile) QueryBookings() *RoomQuery {
	return (&ProfileClient{config: pr.config}).QueryBookings(pr)
}

// QueryEvents queries the "events" edge of the Profile entity.
func (pr *Profile) QueryEvents() *EventQuery {
	return (&ProfileClient{config: pr.config}).QueryEvents(pr)
}

// QueryInvitedTo queries the "invitedTo" edge of the Profile entity.
func (pr *Profile) QueryInvitedTo() *EventQuery {
	return (&ProfileClient{config: pr.config}).QueryInvitedTo(pr)
}

// QueryFriendsData queries the "friends_data" edge of the Profile entity.
func (pr *Profile) QueryFriendsData() *FriendQuery {
	return (&ProfileClient{config: pr.config}).QueryFriendsData(pr)
}

// QueryBookingsData queries the "bookings_data" edge of the Profile entity.
func (pr *Profile) QueryBookingsData() *BookingQuery {
	return (&ProfileClient{config: pr.config}).QueryBookingsData(pr)
}

// QueryEventsData queries the "events_data" edge of the Profile entity.
func (pr *Profile) QueryEventsData() *MemberQuery {
	return (&ProfileClient{config: pr.config}).QueryEventsData(pr)
}

// QueryInvitesData queries the "invites_data" edge of the Profile entity.
func (pr *Profile) QueryInvitesData() *EventInviteQuery {
	return (&ProfileClient{config: pr.config}).QueryInvitesData(pr)
}

// Update returns a builder for updating this Profile.
// Note that you need to call Profile.Unwrap() before calling this method if this Profile
// was returned from a transaction, and the transaction was committed or rolled back.
func (pr *Profile) Update() *ProfileUpdateOne {
	return (&ProfileClient{config: pr.config}).UpdateOne(pr)
}

// Unwrap unwraps the Profile entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (pr *Profile) Unwrap() *Profile {
	_tx, ok := pr.config.driver.(*txDriver)
	if !ok {
		panic("ent: Profile is not a transactional entity")
	}
	pr.config.driver = _tx.drv
	return pr
}

// String implements the fmt.Stringer.
func (pr *Profile) String() string {
	var builder strings.Builder
	builder.WriteString("Profile(")
	builder.WriteString(fmt.Sprintf("id=%v, ", pr.ID))
	builder.WriteString("firstname=")
	builder.WriteString(pr.Firstname)
	builder.WriteString(", ")
	builder.WriteString("lastname=")
	builder.WriteString(pr.Lastname)
	builder.WriteString(", ")
	builder.WriteString("phone=")
	builder.WriteString(pr.Phone)
	builder.WriteString(", ")
	builder.WriteString("email=")
	builder.WriteString(pr.Email)
	builder.WriteByte(')')
	return builder.String()
}

// Profiles is a parsable slice of Profile.
type Profiles []*Profile

func (pr Profiles) config(cfg config) {
	for _i := range pr {
		pr[_i].config = cfg
	}
}
