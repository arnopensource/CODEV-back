// Code generated by ent, DO NOT EDIT.

package networking

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/abc3354/CODEV-back/ent/predicate"
	"github.com/google/uuid"
)

// ProfileID applies equality check predicate on the "profile_id" field. It's identical to ProfileIDEQ.
func ProfileID(v uuid.UUID) predicate.Networking {
	return predicate.Networking(sql.FieldEQ(FieldProfileID, v))
}

// FriendID applies equality check predicate on the "friend_id" field. It's identical to FriendIDEQ.
func FriendID(v uuid.UUID) predicate.Networking {
	return predicate.Networking(sql.FieldEQ(FieldFriendID, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Networking {
	return predicate.Networking(sql.FieldEQ(FieldCreatedAt, v))
}

// Accepted applies equality check predicate on the "accepted" field. It's identical to AcceptedEQ.
func Accepted(v bool) predicate.Networking {
	return predicate.Networking(sql.FieldEQ(FieldAccepted, v))
}

// ProfileIDEQ applies the EQ predicate on the "profile_id" field.
func ProfileIDEQ(v uuid.UUID) predicate.Networking {
	return predicate.Networking(sql.FieldEQ(FieldProfileID, v))
}

// ProfileIDNEQ applies the NEQ predicate on the "profile_id" field.
func ProfileIDNEQ(v uuid.UUID) predicate.Networking {
	return predicate.Networking(sql.FieldNEQ(FieldProfileID, v))
}

// ProfileIDIn applies the In predicate on the "profile_id" field.
func ProfileIDIn(vs ...uuid.UUID) predicate.Networking {
	return predicate.Networking(sql.FieldIn(FieldProfileID, vs...))
}

// ProfileIDNotIn applies the NotIn predicate on the "profile_id" field.
func ProfileIDNotIn(vs ...uuid.UUID) predicate.Networking {
	return predicate.Networking(sql.FieldNotIn(FieldProfileID, vs...))
}

// FriendIDEQ applies the EQ predicate on the "friend_id" field.
func FriendIDEQ(v uuid.UUID) predicate.Networking {
	return predicate.Networking(sql.FieldEQ(FieldFriendID, v))
}

// FriendIDNEQ applies the NEQ predicate on the "friend_id" field.
func FriendIDNEQ(v uuid.UUID) predicate.Networking {
	return predicate.Networking(sql.FieldNEQ(FieldFriendID, v))
}

// FriendIDIn applies the In predicate on the "friend_id" field.
func FriendIDIn(vs ...uuid.UUID) predicate.Networking {
	return predicate.Networking(sql.FieldIn(FieldFriendID, vs...))
}

// FriendIDNotIn applies the NotIn predicate on the "friend_id" field.
func FriendIDNotIn(vs ...uuid.UUID) predicate.Networking {
	return predicate.Networking(sql.FieldNotIn(FieldFriendID, vs...))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Networking {
	return predicate.Networking(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Networking {
	return predicate.Networking(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Networking {
	return predicate.Networking(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Networking {
	return predicate.Networking(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Networking {
	return predicate.Networking(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Networking {
	return predicate.Networking(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Networking {
	return predicate.Networking(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Networking {
	return predicate.Networking(sql.FieldLTE(FieldCreatedAt, v))
}

// AcceptedEQ applies the EQ predicate on the "accepted" field.
func AcceptedEQ(v bool) predicate.Networking {
	return predicate.Networking(sql.FieldEQ(FieldAccepted, v))
}

// AcceptedNEQ applies the NEQ predicate on the "accepted" field.
func AcceptedNEQ(v bool) predicate.Networking {
	return predicate.Networking(sql.FieldNEQ(FieldAccepted, v))
}

// HasProfile applies the HasEdge predicate on the "profile" edge.
func HasProfile() predicate.Networking {
	return predicate.Networking(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, ProfileColumn),
			sqlgraph.Edge(sqlgraph.M2O, false, ProfileTable, ProfileColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasProfileWith applies the HasEdge predicate on the "profile" edge with a given conditions (other predicates).
func HasProfileWith(preds ...predicate.Profile) predicate.Networking {
	return predicate.Networking(func(s *sql.Selector) {
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
func HasFriend() predicate.Networking {
	return predicate.Networking(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FriendColumn),
			sqlgraph.Edge(sqlgraph.M2O, false, FriendTable, FriendColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasFriendWith applies the HasEdge predicate on the "friend" edge with a given conditions (other predicates).
func HasFriendWith(preds ...predicate.Profile) predicate.Networking {
	return predicate.Networking(func(s *sql.Selector) {
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
func And(predicates ...predicate.Networking) predicate.Networking {
	return predicate.Networking(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Networking) predicate.Networking {
	return predicate.Networking(func(s *sql.Selector) {
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
func Not(p predicate.Networking) predicate.Networking {
	return predicate.Networking(func(s *sql.Selector) {
		p(s.Not())
	})
}