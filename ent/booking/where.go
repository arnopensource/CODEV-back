// Code generated by ent, DO NOT EDIT.

package booking

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/abc3354/CODEV-back/ent/predicate"
	"github.com/google/uuid"
)

// ProfileID applies equality check predicate on the "profile_id" field. It's identical to ProfileIDEQ.
func ProfileID(v uuid.UUID) predicate.Booking {
	return predicate.Booking(sql.FieldEQ(FieldProfileID, v))
}

// RoomID applies equality check predicate on the "room_id" field. It's identical to RoomIDEQ.
func RoomID(v int) predicate.Booking {
	return predicate.Booking(sql.FieldEQ(FieldRoomID, v))
}

// NumberOfPeople applies equality check predicate on the "number_of_people" field. It's identical to NumberOfPeopleEQ.
func NumberOfPeople(v int) predicate.Booking {
	return predicate.Booking(sql.FieldEQ(FieldNumberOfPeople, v))
}

// Start applies equality check predicate on the "start" field. It's identical to StartEQ.
func Start(v time.Time) predicate.Booking {
	return predicate.Booking(sql.FieldEQ(FieldStart, v))
}

// End applies equality check predicate on the "end" field. It's identical to EndEQ.
func End(v time.Time) predicate.Booking {
	return predicate.Booking(sql.FieldEQ(FieldEnd, v))
}

// ProfileIDEQ applies the EQ predicate on the "profile_id" field.
func ProfileIDEQ(v uuid.UUID) predicate.Booking {
	return predicate.Booking(sql.FieldEQ(FieldProfileID, v))
}

// ProfileIDNEQ applies the NEQ predicate on the "profile_id" field.
func ProfileIDNEQ(v uuid.UUID) predicate.Booking {
	return predicate.Booking(sql.FieldNEQ(FieldProfileID, v))
}

// ProfileIDIn applies the In predicate on the "profile_id" field.
func ProfileIDIn(vs ...uuid.UUID) predicate.Booking {
	return predicate.Booking(sql.FieldIn(FieldProfileID, vs...))
}

// ProfileIDNotIn applies the NotIn predicate on the "profile_id" field.
func ProfileIDNotIn(vs ...uuid.UUID) predicate.Booking {
	return predicate.Booking(sql.FieldNotIn(FieldProfileID, vs...))
}

// RoomIDEQ applies the EQ predicate on the "room_id" field.
func RoomIDEQ(v int) predicate.Booking {
	return predicate.Booking(sql.FieldEQ(FieldRoomID, v))
}

// RoomIDNEQ applies the NEQ predicate on the "room_id" field.
func RoomIDNEQ(v int) predicate.Booking {
	return predicate.Booking(sql.FieldNEQ(FieldRoomID, v))
}

// RoomIDIn applies the In predicate on the "room_id" field.
func RoomIDIn(vs ...int) predicate.Booking {
	return predicate.Booking(sql.FieldIn(FieldRoomID, vs...))
}

// RoomIDNotIn applies the NotIn predicate on the "room_id" field.
func RoomIDNotIn(vs ...int) predicate.Booking {
	return predicate.Booking(sql.FieldNotIn(FieldRoomID, vs...))
}

// NumberOfPeopleEQ applies the EQ predicate on the "number_of_people" field.
func NumberOfPeopleEQ(v int) predicate.Booking {
	return predicate.Booking(sql.FieldEQ(FieldNumberOfPeople, v))
}

// NumberOfPeopleNEQ applies the NEQ predicate on the "number_of_people" field.
func NumberOfPeopleNEQ(v int) predicate.Booking {
	return predicate.Booking(sql.FieldNEQ(FieldNumberOfPeople, v))
}

// NumberOfPeopleIn applies the In predicate on the "number_of_people" field.
func NumberOfPeopleIn(vs ...int) predicate.Booking {
	return predicate.Booking(sql.FieldIn(FieldNumberOfPeople, vs...))
}

// NumberOfPeopleNotIn applies the NotIn predicate on the "number_of_people" field.
func NumberOfPeopleNotIn(vs ...int) predicate.Booking {
	return predicate.Booking(sql.FieldNotIn(FieldNumberOfPeople, vs...))
}

// NumberOfPeopleGT applies the GT predicate on the "number_of_people" field.
func NumberOfPeopleGT(v int) predicate.Booking {
	return predicate.Booking(sql.FieldGT(FieldNumberOfPeople, v))
}

// NumberOfPeopleGTE applies the GTE predicate on the "number_of_people" field.
func NumberOfPeopleGTE(v int) predicate.Booking {
	return predicate.Booking(sql.FieldGTE(FieldNumberOfPeople, v))
}

// NumberOfPeopleLT applies the LT predicate on the "number_of_people" field.
func NumberOfPeopleLT(v int) predicate.Booking {
	return predicate.Booking(sql.FieldLT(FieldNumberOfPeople, v))
}

// NumberOfPeopleLTE applies the LTE predicate on the "number_of_people" field.
func NumberOfPeopleLTE(v int) predicate.Booking {
	return predicate.Booking(sql.FieldLTE(FieldNumberOfPeople, v))
}

// StartEQ applies the EQ predicate on the "start" field.
func StartEQ(v time.Time) predicate.Booking {
	return predicate.Booking(sql.FieldEQ(FieldStart, v))
}

// StartNEQ applies the NEQ predicate on the "start" field.
func StartNEQ(v time.Time) predicate.Booking {
	return predicate.Booking(sql.FieldNEQ(FieldStart, v))
}

// StartIn applies the In predicate on the "start" field.
func StartIn(vs ...time.Time) predicate.Booking {
	return predicate.Booking(sql.FieldIn(FieldStart, vs...))
}

// StartNotIn applies the NotIn predicate on the "start" field.
func StartNotIn(vs ...time.Time) predicate.Booking {
	return predicate.Booking(sql.FieldNotIn(FieldStart, vs...))
}

// StartGT applies the GT predicate on the "start" field.
func StartGT(v time.Time) predicate.Booking {
	return predicate.Booking(sql.FieldGT(FieldStart, v))
}

// StartGTE applies the GTE predicate on the "start" field.
func StartGTE(v time.Time) predicate.Booking {
	return predicate.Booking(sql.FieldGTE(FieldStart, v))
}

// StartLT applies the LT predicate on the "start" field.
func StartLT(v time.Time) predicate.Booking {
	return predicate.Booking(sql.FieldLT(FieldStart, v))
}

// StartLTE applies the LTE predicate on the "start" field.
func StartLTE(v time.Time) predicate.Booking {
	return predicate.Booking(sql.FieldLTE(FieldStart, v))
}

// EndEQ applies the EQ predicate on the "end" field.
func EndEQ(v time.Time) predicate.Booking {
	return predicate.Booking(sql.FieldEQ(FieldEnd, v))
}

// EndNEQ applies the NEQ predicate on the "end" field.
func EndNEQ(v time.Time) predicate.Booking {
	return predicate.Booking(sql.FieldNEQ(FieldEnd, v))
}

// EndIn applies the In predicate on the "end" field.
func EndIn(vs ...time.Time) predicate.Booking {
	return predicate.Booking(sql.FieldIn(FieldEnd, vs...))
}

// EndNotIn applies the NotIn predicate on the "end" field.
func EndNotIn(vs ...time.Time) predicate.Booking {
	return predicate.Booking(sql.FieldNotIn(FieldEnd, vs...))
}

// EndGT applies the GT predicate on the "end" field.
func EndGT(v time.Time) predicate.Booking {
	return predicate.Booking(sql.FieldGT(FieldEnd, v))
}

// EndGTE applies the GTE predicate on the "end" field.
func EndGTE(v time.Time) predicate.Booking {
	return predicate.Booking(sql.FieldGTE(FieldEnd, v))
}

// EndLT applies the LT predicate on the "end" field.
func EndLT(v time.Time) predicate.Booking {
	return predicate.Booking(sql.FieldLT(FieldEnd, v))
}

// EndLTE applies the LTE predicate on the "end" field.
func EndLTE(v time.Time) predicate.Booking {
	return predicate.Booking(sql.FieldLTE(FieldEnd, v))
}

// HasProfile applies the HasEdge predicate on the "profile" edge.
func HasProfile() predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, ProfileColumn),
			sqlgraph.Edge(sqlgraph.M2O, false, ProfileTable, ProfileColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasProfileWith applies the HasEdge predicate on the "profile" edge with a given conditions (other predicates).
func HasProfileWith(preds ...predicate.Profile) predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, ProfileColumn),
			sqlgraph.To(ProfileInverseTable, ProfileFieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, ProfileTable, ProfileColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasRoom applies the HasEdge predicate on the "room" edge.
func HasRoom() predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, RoomColumn),
			sqlgraph.Edge(sqlgraph.M2O, false, RoomTable, RoomColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasRoomWith applies the HasEdge predicate on the "room" edge with a given conditions (other predicates).
func HasRoomWith(preds ...predicate.Room) predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, RoomColumn),
			sqlgraph.To(RoomInverseTable, RoomFieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, RoomTable, RoomColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Booking) predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Booking) predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Booking) predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		p(s.Not())
	})
}
