// Code generated by ent, DO NOT EDIT.

package friend

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/abc3354/CODEV-back/ent/predicate"
	"github.com/google/uuid"
)

// ProfileID applies equality check predicate on the "profile_id" field. It's identical to ProfileIDEQ.
func ProfileID(v uuid.UUID) predicate.Friend {
	return predicate.Friend(sql.FieldEQ(FieldProfileID, v))
}

// FriendID applies equality check predicate on the "friend_id" field. It's identical to FriendIDEQ.
func FriendID(v uuid.UUID) predicate.Friend {
	return predicate.Friend(sql.FieldEQ(FieldFriendID, v))
}

// Since applies equality check predicate on the "since" field. It's identical to SinceEQ.
func Since(v time.Time) predicate.Friend {
	return predicate.Friend(sql.FieldEQ(FieldSince, v))
}

// Accepted applies equality check predicate on the "accepted" field. It's identical to AcceptedEQ.
func Accepted(v bool) predicate.Friend {
	return predicate.Friend(sql.FieldEQ(FieldAccepted, v))
}

// ProfileIDEQ applies the EQ predicate on the "profile_id" field.
func ProfileIDEQ(v uuid.UUID) predicate.Friend {
	return predicate.Friend(sql.FieldEQ(FieldProfileID, v))
}

// ProfileIDNEQ applies the NEQ predicate on the "profile_id" field.
func ProfileIDNEQ(v uuid.UUID) predicate.Friend {
	return predicate.Friend(sql.FieldNEQ(FieldProfileID, v))
}

// ProfileIDIn applies the In predicate on the "profile_id" field.
func ProfileIDIn(vs ...uuid.UUID) predicate.Friend {
	return predicate.Friend(sql.FieldIn(FieldProfileID, vs...))
}

// ProfileIDNotIn applies the NotIn predicate on the "profile_id" field.
func ProfileIDNotIn(vs ...uuid.UUID) predicate.Friend {
	return predicate.Friend(sql.FieldNotIn(FieldProfileID, vs...))
}

// FriendIDEQ applies the EQ predicate on the "friend_id" field.
func FriendIDEQ(v uuid.UUID) predicate.Friend {
	return predicate.Friend(sql.FieldEQ(FieldFriendID, v))
}

// FriendIDNEQ applies the NEQ predicate on the "friend_id" field.
func FriendIDNEQ(v uuid.UUID) predicate.Friend {
	return predicate.Friend(sql.FieldNEQ(FieldFriendID, v))
}

// FriendIDIn applies the In predicate on the "friend_id" field.
func FriendIDIn(vs ...uuid.UUID) predicate.Friend {
	return predicate.Friend(sql.FieldIn(FieldFriendID, vs...))
}

// FriendIDNotIn applies the NotIn predicate on the "friend_id" field.
func FriendIDNotIn(vs ...uuid.UUID) predicate.Friend {
	return predicate.Friend(sql.FieldNotIn(FieldFriendID, vs...))
}

// SinceEQ applies the EQ predicate on the "since" field.
func SinceEQ(v time.Time) predicate.Friend {
	return predicate.Friend(sql.FieldEQ(FieldSince, v))
}

// SinceNEQ applies the NEQ predicate on the "since" field.
func SinceNEQ(v time.Time) predicate.Friend {
	return predicate.Friend(sql.FieldNEQ(FieldSince, v))
}

// SinceIn applies the In predicate on the "since" field.
func SinceIn(vs ...time.Time) predicate.Friend {
	return predicate.Friend(sql.FieldIn(FieldSince, vs...))
}

// SinceNotIn applies the NotIn predicate on the "since" field.
func SinceNotIn(vs ...time.Time) predicate.Friend {
	return predicate.Friend(sql.FieldNotIn(FieldSince, vs...))
}

// SinceGT applies the GT predicate on the "since" field.
func SinceGT(v time.Time) predicate.Friend {
	return predicate.Friend(sql.FieldGT(FieldSince, v))
}

// SinceGTE applies the GTE predicate on the "since" field.
func SinceGTE(v time.Time) predicate.Friend {
	return predicate.Friend(sql.FieldGTE(FieldSince, v))
}

// SinceLT applies the LT predicate on the "since" field.
func SinceLT(v time.Time) predicate.Friend {
	return predicate.Friend(sql.FieldLT(FieldSince, v))
}

// SinceLTE applies the LTE predicate on the "since" field.
func SinceLTE(v time.Time) predicate.Friend {
	return predicate.Friend(sql.FieldLTE(FieldSince, v))
}

// AcceptedEQ applies the EQ predicate on the "accepted" field.
func AcceptedEQ(v bool) predicate.Friend {
	return predicate.Friend(sql.FieldEQ(FieldAccepted, v))
}

// AcceptedNEQ applies the NEQ predicate on the "accepted" field.
func AcceptedNEQ(v bool) predicate.Friend {
	return predicate.Friend(sql.FieldNEQ(FieldAccepted, v))
}

// HasProfile applies the HasEdge predicate on the "profile" edge.
func HasProfile() predicate.Friend {
	return predicate.Friend(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, ProfileColumn),
			sqlgraph.Edge(sqlgraph.M2O, false, ProfileTable, ProfileColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasProfileWith applies the HasEdge predicate on the "profile" edge with a given conditions (other predicates).
func HasProfileWith(preds ...predicate.Profile) predicate.Friend {
	return predicate.Friend(func(s *sql.Selector) {
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

// HasFriend applies the HasEdge predicate on the "friend" edge.
func HasFriend() predicate.Friend {
	return predicate.Friend(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FriendColumn),
			sqlgraph.Edge(sqlgraph.M2O, false, FriendTable, FriendColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasFriendWith applies the HasEdge predicate on the "friend" edge with a given conditions (other predicates).
func HasFriendWith(preds ...predicate.Profile) predicate.Friend {
	return predicate.Friend(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FriendColumn),
			sqlgraph.To(FriendInverseTable, ProfileFieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, FriendTable, FriendColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Friend) predicate.Friend {
	return predicate.Friend(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Friend) predicate.Friend {
	return predicate.Friend(func(s *sql.Selector) {
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
func Not(p predicate.Friend) predicate.Friend {
	return predicate.Friend(func(s *sql.Selector) {
		p(s.Not())
	})
}
