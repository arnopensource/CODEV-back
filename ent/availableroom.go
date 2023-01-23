// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/abc3354/CODEV-back/ent/availableroom"
	"github.com/abc3354/CODEV-back/ent/room"
)

// AvailableRoom is the model entity for the AvailableRoom schema.
type AvailableRoom struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Start holds the value of the "start" field.
	Start time.Time `json:"start,omitempty"`
	// End holds the value of the "end" field.
	End time.Time `json:"end,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the AvailableRoomQuery when eager-loading is set.
	Edges             AvailableRoomEdges `json:"edges"`
	room_availability *int
}

// AvailableRoomEdges holds the relations/edges for other nodes in the graph.
type AvailableRoomEdges struct {
	// Rooms holds the value of the rooms edge.
	Rooms *Room `json:"rooms,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// RoomsOrErr returns the Rooms value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e AvailableRoomEdges) RoomsOrErr() (*Room, error) {
	if e.loadedTypes[0] {
		if e.Rooms == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: room.Label}
		}
		return e.Rooms, nil
	}
	return nil, &NotLoadedError{edge: "rooms"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*AvailableRoom) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case availableroom.FieldID:
			values[i] = new(sql.NullInt64)
		case availableroom.FieldStart, availableroom.FieldEnd:
			values[i] = new(sql.NullTime)
		case availableroom.ForeignKeys[0]: // room_availability
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type AvailableRoom", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the AvailableRoom fields.
func (ar *AvailableRoom) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case availableroom.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			ar.ID = int(value.Int64)
		case availableroom.FieldStart:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field start", values[i])
			} else if value.Valid {
				ar.Start = value.Time
			}
		case availableroom.FieldEnd:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field end", values[i])
			} else if value.Valid {
				ar.End = value.Time
			}
		case availableroom.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field room_availability", value)
			} else if value.Valid {
				ar.room_availability = new(int)
				*ar.room_availability = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryRooms queries the "rooms" edge of the AvailableRoom entity.
func (ar *AvailableRoom) QueryRooms() *RoomQuery {
	return (&AvailableRoomClient{config: ar.config}).QueryRooms(ar)
}

// Update returns a builder for updating this AvailableRoom.
// Note that you need to call AvailableRoom.Unwrap() before calling this method if this AvailableRoom
// was returned from a transaction, and the transaction was committed or rolled back.
func (ar *AvailableRoom) Update() *AvailableRoomUpdateOne {
	return (&AvailableRoomClient{config: ar.config}).UpdateOne(ar)
}

// Unwrap unwraps the AvailableRoom entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ar *AvailableRoom) Unwrap() *AvailableRoom {
	_tx, ok := ar.config.driver.(*txDriver)
	if !ok {
		panic("ent: AvailableRoom is not a transactional entity")
	}
	ar.config.driver = _tx.drv
	return ar
}

// String implements the fmt.Stringer.
func (ar *AvailableRoom) String() string {
	var builder strings.Builder
	builder.WriteString("AvailableRoom(")
	builder.WriteString(fmt.Sprintf("id=%v, ", ar.ID))
	builder.WriteString("start=")
	builder.WriteString(ar.Start.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("end=")
	builder.WriteString(ar.End.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// AvailableRooms is a parsable slice of AvailableRoom.
type AvailableRooms []*AvailableRoom

func (ar AvailableRooms) config(cfg config) {
	for _i := range ar {
		ar[_i].config = cfg
	}
}
