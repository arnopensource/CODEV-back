// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/abc3354/CODEV-back/ent/event"
	"github.com/abc3354/CODEV-back/ent/member"
	"github.com/abc3354/CODEV-back/ent/predicate"
	"github.com/abc3354/CODEV-back/ent/profile"
	"github.com/google/uuid"
)

// MemberUpdate is the builder for updating Member entities.
type MemberUpdate struct {
	config
	hooks    []Hook
	mutation *MemberMutation
}

// Where appends a list predicates to the MemberUpdate builder.
func (mu *MemberUpdate) Where(ps ...predicate.Member) *MemberUpdate {
	mu.mutation.Where(ps...)
	return mu
}

// SetProfileID sets the "profile_id" field.
func (mu *MemberUpdate) SetProfileID(u uuid.UUID) *MemberUpdate {
	mu.mutation.SetProfileID(u)
	return mu
}

// SetEventID sets the "event_id" field.
func (mu *MemberUpdate) SetEventID(i int) *MemberUpdate {
	mu.mutation.SetEventID(i)
	return mu
}

// SetIsAdmin sets the "is_admin" field.
func (mu *MemberUpdate) SetIsAdmin(b bool) *MemberUpdate {
	mu.mutation.SetIsAdmin(b)
	return mu
}

// SetNillableIsAdmin sets the "is_admin" field if the given value is not nil.
func (mu *MemberUpdate) SetNillableIsAdmin(b *bool) *MemberUpdate {
	if b != nil {
		mu.SetIsAdmin(*b)
	}
	return mu
}

// SetProfile sets the "profile" edge to the Profile entity.
func (mu *MemberUpdate) SetProfile(p *Profile) *MemberUpdate {
	return mu.SetProfileID(p.ID)
}

// SetEvent sets the "event" edge to the Event entity.
func (mu *MemberUpdate) SetEvent(e *Event) *MemberUpdate {
	return mu.SetEventID(e.ID)
}

// Mutation returns the MemberMutation object of the builder.
func (mu *MemberUpdate) Mutation() *MemberMutation {
	return mu.mutation
}

// ClearProfile clears the "profile" edge to the Profile entity.
func (mu *MemberUpdate) ClearProfile() *MemberUpdate {
	mu.mutation.ClearProfile()
	return mu
}

// ClearEvent clears the "event" edge to the Event entity.
func (mu *MemberUpdate) ClearEvent() *MemberUpdate {
	mu.mutation.ClearEvent()
	return mu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (mu *MemberUpdate) Save(ctx context.Context) (int, error) {
	return withHooks[int, MemberMutation](ctx, mu.sqlSave, mu.mutation, mu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (mu *MemberUpdate) SaveX(ctx context.Context) int {
	affected, err := mu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (mu *MemberUpdate) Exec(ctx context.Context) error {
	_, err := mu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mu *MemberUpdate) ExecX(ctx context.Context) {
	if err := mu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (mu *MemberUpdate) check() error {
	if _, ok := mu.mutation.ProfileID(); mu.mutation.ProfileCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Member.profile"`)
	}
	if _, ok := mu.mutation.EventID(); mu.mutation.EventCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Member.event"`)
	}
	return nil
}

func (mu *MemberUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := mu.check(); err != nil {
		return n, err
	}
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   member.Table,
			Columns: member.Columns,
			CompositeID: []*sqlgraph.FieldSpec{
				{
					Type:   field.TypeInt,
					Column: member.FieldEventID,
				},
				{
					Type:   field.TypeUUID,
					Column: member.FieldProfileID,
				},
			},
		},
	}
	if ps := mu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := mu.mutation.IsAdmin(); ok {
		_spec.SetField(member.FieldIsAdmin, field.TypeBool, value)
	}
	if mu.mutation.ProfileCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   member.ProfileTable,
			Columns: []string{member.ProfileColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: profile.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mu.mutation.ProfileIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   member.ProfileTable,
			Columns: []string{member.ProfileColumn},
			Bidi:    false,
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
	if mu.mutation.EventCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   member.EventTable,
			Columns: []string{member.EventColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: event.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mu.mutation.EventIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   member.EventTable,
			Columns: []string{member.EventColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: event.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, mu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{member.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	mu.mutation.done = true
	return n, nil
}

// MemberUpdateOne is the builder for updating a single Member entity.
type MemberUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *MemberMutation
}

// SetProfileID sets the "profile_id" field.
func (muo *MemberUpdateOne) SetProfileID(u uuid.UUID) *MemberUpdateOne {
	muo.mutation.SetProfileID(u)
	return muo
}

// SetEventID sets the "event_id" field.
func (muo *MemberUpdateOne) SetEventID(i int) *MemberUpdateOne {
	muo.mutation.SetEventID(i)
	return muo
}

// SetIsAdmin sets the "is_admin" field.
func (muo *MemberUpdateOne) SetIsAdmin(b bool) *MemberUpdateOne {
	muo.mutation.SetIsAdmin(b)
	return muo
}

// SetNillableIsAdmin sets the "is_admin" field if the given value is not nil.
func (muo *MemberUpdateOne) SetNillableIsAdmin(b *bool) *MemberUpdateOne {
	if b != nil {
		muo.SetIsAdmin(*b)
	}
	return muo
}

// SetProfile sets the "profile" edge to the Profile entity.
func (muo *MemberUpdateOne) SetProfile(p *Profile) *MemberUpdateOne {
	return muo.SetProfileID(p.ID)
}

// SetEvent sets the "event" edge to the Event entity.
func (muo *MemberUpdateOne) SetEvent(e *Event) *MemberUpdateOne {
	return muo.SetEventID(e.ID)
}

// Mutation returns the MemberMutation object of the builder.
func (muo *MemberUpdateOne) Mutation() *MemberMutation {
	return muo.mutation
}

// ClearProfile clears the "profile" edge to the Profile entity.
func (muo *MemberUpdateOne) ClearProfile() *MemberUpdateOne {
	muo.mutation.ClearProfile()
	return muo
}

// ClearEvent clears the "event" edge to the Event entity.
func (muo *MemberUpdateOne) ClearEvent() *MemberUpdateOne {
	muo.mutation.ClearEvent()
	return muo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (muo *MemberUpdateOne) Select(field string, fields ...string) *MemberUpdateOne {
	muo.fields = append([]string{field}, fields...)
	return muo
}

// Save executes the query and returns the updated Member entity.
func (muo *MemberUpdateOne) Save(ctx context.Context) (*Member, error) {
	return withHooks[*Member, MemberMutation](ctx, muo.sqlSave, muo.mutation, muo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (muo *MemberUpdateOne) SaveX(ctx context.Context) *Member {
	node, err := muo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (muo *MemberUpdateOne) Exec(ctx context.Context) error {
	_, err := muo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (muo *MemberUpdateOne) ExecX(ctx context.Context) {
	if err := muo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (muo *MemberUpdateOne) check() error {
	if _, ok := muo.mutation.ProfileID(); muo.mutation.ProfileCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Member.profile"`)
	}
	if _, ok := muo.mutation.EventID(); muo.mutation.EventCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Member.event"`)
	}
	return nil
}

func (muo *MemberUpdateOne) sqlSave(ctx context.Context) (_node *Member, err error) {
	if err := muo.check(); err != nil {
		return _node, err
	}
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   member.Table,
			Columns: member.Columns,
			CompositeID: []*sqlgraph.FieldSpec{
				{
					Type:   field.TypeInt,
					Column: member.FieldEventID,
				},
				{
					Type:   field.TypeUUID,
					Column: member.FieldProfileID,
				},
			},
		},
	}
	if id, ok := muo.mutation.EventID(); !ok {
		return nil, &ValidationError{Name: "event_id", err: errors.New(`ent: missing "Member.event_id" for update`)}
	} else {
		_spec.Node.CompositeID[0].Value = id
	}
	if id, ok := muo.mutation.ProfileID(); !ok {
		return nil, &ValidationError{Name: "profile_id", err: errors.New(`ent: missing "Member.profile_id" for update`)}
	} else {
		_spec.Node.CompositeID[1].Value = id
	}
	if fields := muo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, len(fields))
		for i, f := range fields {
			if !member.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			_spec.Node.Columns[i] = f
		}
	}
	if ps := muo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := muo.mutation.IsAdmin(); ok {
		_spec.SetField(member.FieldIsAdmin, field.TypeBool, value)
	}
	if muo.mutation.ProfileCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   member.ProfileTable,
			Columns: []string{member.ProfileColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: profile.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := muo.mutation.ProfileIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   member.ProfileTable,
			Columns: []string{member.ProfileColumn},
			Bidi:    false,
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
	if muo.mutation.EventCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   member.EventTable,
			Columns: []string{member.EventColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: event.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := muo.mutation.EventIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   member.EventTable,
			Columns: []string{member.EventColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: event.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Member{config: muo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, muo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{member.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	muo.mutation.done = true
	return _node, nil
}