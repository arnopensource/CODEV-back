// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/abc3354/CODEV-back/ent/room"
)

// Room is the model entity for the Room schema.
type Room struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Floor holds the value of the "floor" field.
	Floor string `json:"floor,omitempty"`
	// Building holds the value of the "building" field.
	Building string `json:"building,omitempty"`
	// Capacity holds the value of the "capacity" field.
	Capacity int `json:"capacity,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the RoomQuery when eager-loading is set.
	Edges RoomEdges `json:"edges"`
}

// RoomEdges holds the relations/edges for other nodes in the graph.
type RoomEdges struct {
	// Bookings holds the value of the bookings edge.
	Bookings []*Profile `json:"bookings,omitempty"`
	// Availability holds the value of the availability edge.
	Availability []*AvailableRoom `json:"availability,omitempty"`
	// BookingsData holds the value of the bookings_data edge.
	BookingsData []*Booking `json:"bookings_data,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
}

// BookingsOrErr returns the Bookings value or an error if the edge
// was not loaded in eager-loading.
func (e RoomEdges) BookingsOrErr() ([]*Profile, error) {
	if e.loadedTypes[0] {
		return e.Bookings, nil
	}
	return nil, &NotLoadedError{edge: "bookings"}
}

// AvailabilityOrErr returns the Availability value or an error if the edge
// was not loaded in eager-loading.
func (e RoomEdges) AvailabilityOrErr() ([]*AvailableRoom, error) {
	if e.loadedTypes[1] {
		return e.Availability, nil
	}
	return nil, &NotLoadedError{edge: "availability"}
}

// BookingsDataOrErr returns the BookingsData value or an error if the edge
// was not loaded in eager-loading.
func (e RoomEdges) BookingsDataOrErr() ([]*Booking, error) {
	if e.loadedTypes[2] {
		return e.BookingsData, nil
	}
	return nil, &NotLoadedError{edge: "bookings_data"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Room) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case room.FieldID, room.FieldCapacity:
			values[i] = new(sql.NullInt64)
		case room.FieldName, room.FieldFloor, room.FieldBuilding:
			values[i] = new(sql.NullString)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Room", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Room fields.
func (r *Room) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case room.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			r.ID = int(value.Int64)
		case room.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				r.Name = value.String
			}
		case room.FieldFloor:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field floor", values[i])
			} else if value.Valid {
				r.Floor = value.String
			}
		case room.FieldBuilding:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field building", values[i])
			} else if value.Valid {
				r.Building = value.String
			}
		case room.FieldCapacity:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field capacity", values[i])
			} else if value.Valid {
				r.Capacity = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryBookings queries the "bookings" edge of the Room entity.
func (r *Room) QueryBookings() *ProfileQuery {
	return (&RoomClient{config: r.config}).QueryBookings(r)
}

// QueryAvailability queries the "availability" edge of the Room entity.
func (r *Room) QueryAvailability() *AvailableRoomQuery {
	return (&RoomClient{config: r.config}).QueryAvailability(r)
}

// QueryBookingsData queries the "bookings_data" edge of the Room entity.
func (r *Room) QueryBookingsData() *BookingQuery {
	return (&RoomClient{config: r.config}).QueryBookingsData(r)
}

// Update returns a builder for updating this Room.
// Note that you need to call Room.Unwrap() before calling this method if this Room
// was returned from a transaction, and the transaction was committed or rolled back.
func (r *Room) Update() *RoomUpdateOne {
	return (&RoomClient{config: r.config}).UpdateOne(r)
}

// Unwrap unwraps the Room entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (r *Room) Unwrap() *Room {
	_tx, ok := r.config.driver.(*txDriver)
	if !ok {
		panic("ent: Room is not a transactional entity")
	}
	r.config.driver = _tx.drv
	return r
}

// String implements the fmt.Stringer.
func (r *Room) String() string {
	var builder strings.Builder
	builder.WriteString("Room(")
	builder.WriteString(fmt.Sprintf("id=%v, ", r.ID))
	builder.WriteString("name=")
	builder.WriteString(r.Name)
	builder.WriteString(", ")
	builder.WriteString("floor=")
	builder.WriteString(r.Floor)
	builder.WriteString(", ")
	builder.WriteString("building=")
	builder.WriteString(r.Building)
	builder.WriteString(", ")
	builder.WriteString("capacity=")
	builder.WriteString(fmt.Sprintf("%v", r.Capacity))
	builder.WriteByte(')')
	return builder.String()
}

// Rooms is a parsable slice of Room.
type Rooms []*Room

func (r Rooms) config(cfg config) {
	for _i := range r {
		r[_i].config = cfg
	}
}
