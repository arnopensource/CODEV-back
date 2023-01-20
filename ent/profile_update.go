// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/abc3354/CODEV-back/ent/predicate"
	"github.com/abc3354/CODEV-back/ent/profile"
	"github.com/abc3354/CODEV-back/ent/salle"
	"github.com/google/uuid"
)

// ProfileUpdate is the builder for updating Profile entities.
type ProfileUpdate struct {
	config
	hooks    []Hook
	mutation *ProfileMutation
}

// Where appends a list predicates to the ProfileUpdate builder.
func (pu *ProfileUpdate) Where(ps ...predicate.Profile) *ProfileUpdate {
	pu.mutation.Where(ps...)
	return pu
}

// SetFirstname sets the "firstname" field.
func (pu *ProfileUpdate) SetFirstname(s string) *ProfileUpdate {
	pu.mutation.SetFirstname(s)
	return pu
}

// SetLastname sets the "lastname" field.
func (pu *ProfileUpdate) SetLastname(s string) *ProfileUpdate {
	pu.mutation.SetLastname(s)
	return pu
}

// SetTelephone sets the "telephone" field.
func (pu *ProfileUpdate) SetTelephone(s string) *ProfileUpdate {
	pu.mutation.SetTelephone(s)
	return pu
}

// AddFriendIDs adds the "friends" edge to the Profile entity by IDs.
func (pu *ProfileUpdate) AddFriendIDs(ids ...uuid.UUID) *ProfileUpdate {
	pu.mutation.AddFriendIDs(ids...)
	return pu
}

// AddFriends adds the "friends" edges to the Profile entity.
func (pu *ProfileUpdate) AddFriends(p ...*Profile) *ProfileUpdate {
	ids := make([]uuid.UUID, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return pu.AddFriendIDs(ids...)
}

// AddSalleReserveeIDs adds the "salle_reservee" edge to the Salle entity by IDs.
func (pu *ProfileUpdate) AddSalleReserveeIDs(ids ...int) *ProfileUpdate {
	pu.mutation.AddSalleReserveeIDs(ids...)
	return pu
}

// AddSalleReservee adds the "salle_reservee" edges to the Salle entity.
func (pu *ProfileUpdate) AddSalleReservee(s ...*Salle) *ProfileUpdate {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return pu.AddSalleReserveeIDs(ids...)
}

// Mutation returns the ProfileMutation object of the builder.
func (pu *ProfileUpdate) Mutation() *ProfileMutation {
	return pu.mutation
}

// ClearFriends clears all "friends" edges to the Profile entity.
func (pu *ProfileUpdate) ClearFriends() *ProfileUpdate {
	pu.mutation.ClearFriends()
	return pu
}

// RemoveFriendIDs removes the "friends" edge to Profile entities by IDs.
func (pu *ProfileUpdate) RemoveFriendIDs(ids ...uuid.UUID) *ProfileUpdate {
	pu.mutation.RemoveFriendIDs(ids...)
	return pu
}

// RemoveFriends removes "friends" edges to Profile entities.
func (pu *ProfileUpdate) RemoveFriends(p ...*Profile) *ProfileUpdate {
	ids := make([]uuid.UUID, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return pu.RemoveFriendIDs(ids...)
}

// ClearSalleReservee clears all "salle_reservee" edges to the Salle entity.
func (pu *ProfileUpdate) ClearSalleReservee() *ProfileUpdate {
	pu.mutation.ClearSalleReservee()
	return pu
}

// RemoveSalleReserveeIDs removes the "salle_reservee" edge to Salle entities by IDs.
func (pu *ProfileUpdate) RemoveSalleReserveeIDs(ids ...int) *ProfileUpdate {
	pu.mutation.RemoveSalleReserveeIDs(ids...)
	return pu
}

// RemoveSalleReservee removes "salle_reservee" edges to Salle entities.
func (pu *ProfileUpdate) RemoveSalleReservee(s ...*Salle) *ProfileUpdate {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return pu.RemoveSalleReserveeIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (pu *ProfileUpdate) Save(ctx context.Context) (int, error) {
	return withHooks[int, ProfileMutation](ctx, pu.sqlSave, pu.mutation, pu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (pu *ProfileUpdate) SaveX(ctx context.Context) int {
	affected, err := pu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pu *ProfileUpdate) Exec(ctx context.Context) error {
	_, err := pu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pu *ProfileUpdate) ExecX(ctx context.Context) {
	if err := pu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (pu *ProfileUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   profile.Table,
			Columns: profile.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: profile.FieldID,
			},
		},
	}
	if ps := pu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pu.mutation.Firstname(); ok {
		_spec.SetField(profile.FieldFirstname, field.TypeString, value)
	}
	if value, ok := pu.mutation.Lastname(); ok {
		_spec.SetField(profile.FieldLastname, field.TypeString, value)
	}
	if value, ok := pu.mutation.Telephone(); ok {
		_spec.SetField(profile.FieldTelephone, field.TypeString, value)
	}
	if pu.mutation.FriendsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   profile.FriendsTable,
			Columns: profile.FriendsPrimaryKey,
			Bidi:    true,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: profile.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.RemovedFriendsIDs(); len(nodes) > 0 && !pu.mutation.FriendsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   profile.FriendsTable,
			Columns: profile.FriendsPrimaryKey,
			Bidi:    true,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: profile.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.FriendsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   profile.FriendsTable,
			Columns: profile.FriendsPrimaryKey,
			Bidi:    true,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: profile.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if pu.mutation.SalleReserveeCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   profile.SalleReserveeTable,
			Columns: profile.SalleReserveePrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: salle.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.RemovedSalleReserveeIDs(); len(nodes) > 0 && !pu.mutation.SalleReserveeCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   profile.SalleReserveeTable,
			Columns: profile.SalleReserveePrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: salle.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.SalleReserveeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   profile.SalleReserveeTable,
			Columns: profile.SalleReserveePrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: salle.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, pu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{profile.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	pu.mutation.done = true
	return n, nil
}

// ProfileUpdateOne is the builder for updating a single Profile entity.
type ProfileUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ProfileMutation
}

// SetFirstname sets the "firstname" field.
func (puo *ProfileUpdateOne) SetFirstname(s string) *ProfileUpdateOne {
	puo.mutation.SetFirstname(s)
	return puo
}

// SetLastname sets the "lastname" field.
func (puo *ProfileUpdateOne) SetLastname(s string) *ProfileUpdateOne {
	puo.mutation.SetLastname(s)
	return puo
}

// SetTelephone sets the "telephone" field.
func (puo *ProfileUpdateOne) SetTelephone(s string) *ProfileUpdateOne {
	puo.mutation.SetTelephone(s)
	return puo
}

// AddFriendIDs adds the "friends" edge to the Profile entity by IDs.
func (puo *ProfileUpdateOne) AddFriendIDs(ids ...uuid.UUID) *ProfileUpdateOne {
	puo.mutation.AddFriendIDs(ids...)
	return puo
}

// AddFriends adds the "friends" edges to the Profile entity.
func (puo *ProfileUpdateOne) AddFriends(p ...*Profile) *ProfileUpdateOne {
	ids := make([]uuid.UUID, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return puo.AddFriendIDs(ids...)
}

// AddSalleReserveeIDs adds the "salle_reservee" edge to the Salle entity by IDs.
func (puo *ProfileUpdateOne) AddSalleReserveeIDs(ids ...int) *ProfileUpdateOne {
	puo.mutation.AddSalleReserveeIDs(ids...)
	return puo
}

// AddSalleReservee adds the "salle_reservee" edges to the Salle entity.
func (puo *ProfileUpdateOne) AddSalleReservee(s ...*Salle) *ProfileUpdateOne {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return puo.AddSalleReserveeIDs(ids...)
}

// Mutation returns the ProfileMutation object of the builder.
func (puo *ProfileUpdateOne) Mutation() *ProfileMutation {
	return puo.mutation
}

// ClearFriends clears all "friends" edges to the Profile entity.
func (puo *ProfileUpdateOne) ClearFriends() *ProfileUpdateOne {
	puo.mutation.ClearFriends()
	return puo
}

// RemoveFriendIDs removes the "friends" edge to Profile entities by IDs.
func (puo *ProfileUpdateOne) RemoveFriendIDs(ids ...uuid.UUID) *ProfileUpdateOne {
	puo.mutation.RemoveFriendIDs(ids...)
	return puo
}

// RemoveFriends removes "friends" edges to Profile entities.
func (puo *ProfileUpdateOne) RemoveFriends(p ...*Profile) *ProfileUpdateOne {
	ids := make([]uuid.UUID, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return puo.RemoveFriendIDs(ids...)
}

// ClearSalleReservee clears all "salle_reservee" edges to the Salle entity.
func (puo *ProfileUpdateOne) ClearSalleReservee() *ProfileUpdateOne {
	puo.mutation.ClearSalleReservee()
	return puo
}

// RemoveSalleReserveeIDs removes the "salle_reservee" edge to Salle entities by IDs.
func (puo *ProfileUpdateOne) RemoveSalleReserveeIDs(ids ...int) *ProfileUpdateOne {
	puo.mutation.RemoveSalleReserveeIDs(ids...)
	return puo
}

// RemoveSalleReservee removes "salle_reservee" edges to Salle entities.
func (puo *ProfileUpdateOne) RemoveSalleReservee(s ...*Salle) *ProfileUpdateOne {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return puo.RemoveSalleReserveeIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (puo *ProfileUpdateOne) Select(field string, fields ...string) *ProfileUpdateOne {
	puo.fields = append([]string{field}, fields...)
	return puo
}

// Save executes the query and returns the updated Profile entity.
func (puo *ProfileUpdateOne) Save(ctx context.Context) (*Profile, error) {
	return withHooks[*Profile, ProfileMutation](ctx, puo.sqlSave, puo.mutation, puo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (puo *ProfileUpdateOne) SaveX(ctx context.Context) *Profile {
	node, err := puo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (puo *ProfileUpdateOne) Exec(ctx context.Context) error {
	_, err := puo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (puo *ProfileUpdateOne) ExecX(ctx context.Context) {
	if err := puo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (puo *ProfileUpdateOne) sqlSave(ctx context.Context) (_node *Profile, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   profile.Table,
			Columns: profile.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: profile.FieldID,
			},
		},
	}
	id, ok := puo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Profile.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := puo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, profile.FieldID)
		for _, f := range fields {
			if !profile.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != profile.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := puo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := puo.mutation.Firstname(); ok {
		_spec.SetField(profile.FieldFirstname, field.TypeString, value)
	}
	if value, ok := puo.mutation.Lastname(); ok {
		_spec.SetField(profile.FieldLastname, field.TypeString, value)
	}
	if value, ok := puo.mutation.Telephone(); ok {
		_spec.SetField(profile.FieldTelephone, field.TypeString, value)
	}
	if puo.mutation.FriendsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   profile.FriendsTable,
			Columns: profile.FriendsPrimaryKey,
			Bidi:    true,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: profile.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.RemovedFriendsIDs(); len(nodes) > 0 && !puo.mutation.FriendsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   profile.FriendsTable,
			Columns: profile.FriendsPrimaryKey,
			Bidi:    true,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: profile.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.FriendsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   profile.FriendsTable,
			Columns: profile.FriendsPrimaryKey,
			Bidi:    true,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: profile.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if puo.mutation.SalleReserveeCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   profile.SalleReserveeTable,
			Columns: profile.SalleReserveePrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: salle.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.RemovedSalleReserveeIDs(); len(nodes) > 0 && !puo.mutation.SalleReserveeCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   profile.SalleReserveeTable,
			Columns: profile.SalleReserveePrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: salle.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.SalleReserveeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   profile.SalleReserveeTable,
			Columns: profile.SalleReserveePrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: salle.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Profile{config: puo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, puo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{profile.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	puo.mutation.done = true
	return _node, nil
}
