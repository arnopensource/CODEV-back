// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/abc3354/CODEV-back/ent/event"
	"github.com/abc3354/CODEV-back/ent/eventinvite"
	"github.com/abc3354/CODEV-back/ent/profile"
	"github.com/google/uuid"
)

// EventInviteCreate is the builder for creating a EventInvite entity.
type EventInviteCreate struct {
	config
	mutation *EventInviteMutation
	hooks    []Hook
}

// SetProfileID sets the "profile_id" field.
func (eic *EventInviteCreate) SetProfileID(u uuid.UUID) *EventInviteCreate {
	eic.mutation.SetProfileID(u)
	return eic
}

// SetEventID sets the "event_id" field.
func (eic *EventInviteCreate) SetEventID(i int) *EventInviteCreate {
	eic.mutation.SetEventID(i)
	return eic
}

// SetSince sets the "since" field.
func (eic *EventInviteCreate) SetSince(t time.Time) *EventInviteCreate {
	eic.mutation.SetSince(t)
	return eic
}

// SetNillableSince sets the "since" field if the given value is not nil.
func (eic *EventInviteCreate) SetNillableSince(t *time.Time) *EventInviteCreate {
	if t != nil {
		eic.SetSince(*t)
	}
	return eic
}

// SetProfile sets the "profile" edge to the Profile entity.
func (eic *EventInviteCreate) SetProfile(p *Profile) *EventInviteCreate {
	return eic.SetProfileID(p.ID)
}

// SetEvent sets the "event" edge to the Event entity.
func (eic *EventInviteCreate) SetEvent(e *Event) *EventInviteCreate {
	return eic.SetEventID(e.ID)
}

// Mutation returns the EventInviteMutation object of the builder.
func (eic *EventInviteCreate) Mutation() *EventInviteMutation {
	return eic.mutation
}

// Save creates the EventInvite in the database.
func (eic *EventInviteCreate) Save(ctx context.Context) (*EventInvite, error) {
	eic.defaults()
	return withHooks[*EventInvite, EventInviteMutation](ctx, eic.sqlSave, eic.mutation, eic.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (eic *EventInviteCreate) SaveX(ctx context.Context) *EventInvite {
	v, err := eic.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (eic *EventInviteCreate) Exec(ctx context.Context) error {
	_, err := eic.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (eic *EventInviteCreate) ExecX(ctx context.Context) {
	if err := eic.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (eic *EventInviteCreate) defaults() {
	if _, ok := eic.mutation.Since(); !ok {
		v := eventinvite.DefaultSince()
		eic.mutation.SetSince(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (eic *EventInviteCreate) check() error {
	if _, ok := eic.mutation.ProfileID(); !ok {
		return &ValidationError{Name: "profile_id", err: errors.New(`ent: missing required field "EventInvite.profile_id"`)}
	}
	if _, ok := eic.mutation.EventID(); !ok {
		return &ValidationError{Name: "event_id", err: errors.New(`ent: missing required field "EventInvite.event_id"`)}
	}
	if _, ok := eic.mutation.Since(); !ok {
		return &ValidationError{Name: "since", err: errors.New(`ent: missing required field "EventInvite.since"`)}
	}
	if _, ok := eic.mutation.ProfileID(); !ok {
		return &ValidationError{Name: "profile", err: errors.New(`ent: missing required edge "EventInvite.profile"`)}
	}
	if _, ok := eic.mutation.EventID(); !ok {
		return &ValidationError{Name: "event", err: errors.New(`ent: missing required edge "EventInvite.event"`)}
	}
	return nil
}

func (eic *EventInviteCreate) sqlSave(ctx context.Context) (*EventInvite, error) {
	if err := eic.check(); err != nil {
		return nil, err
	}
	_node, _spec := eic.createSpec()
	if err := sqlgraph.CreateNode(ctx, eic.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}

func (eic *EventInviteCreate) createSpec() (*EventInvite, *sqlgraph.CreateSpec) {
	var (
		_node = &EventInvite{config: eic.config}
		_spec = &sqlgraph.CreateSpec{
			Table: eventinvite.Table,
		}
	)
	if value, ok := eic.mutation.Since(); ok {
		_spec.SetField(eventinvite.FieldSince, field.TypeTime, value)
		_node.Since = value
	}
	if nodes := eic.mutation.ProfileIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   eventinvite.ProfileTable,
			Columns: []string{eventinvite.ProfileColumn},
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
		_node.ProfileID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := eic.mutation.EventIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   eventinvite.EventTable,
			Columns: []string{eventinvite.EventColumn},
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
		_node.EventID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// EventInviteCreateBulk is the builder for creating many EventInvite entities in bulk.
type EventInviteCreateBulk struct {
	config
	builders []*EventInviteCreate
}

// Save creates the EventInvite entities in the database.
func (eicb *EventInviteCreateBulk) Save(ctx context.Context) ([]*EventInvite, error) {
	specs := make([]*sqlgraph.CreateSpec, len(eicb.builders))
	nodes := make([]*EventInvite, len(eicb.builders))
	mutators := make([]Mutator, len(eicb.builders))
	for i := range eicb.builders {
		func(i int, root context.Context) {
			builder := eicb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*EventInviteMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, eicb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, eicb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, eicb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (eicb *EventInviteCreateBulk) SaveX(ctx context.Context) []*EventInvite {
	v, err := eicb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (eicb *EventInviteCreateBulk) Exec(ctx context.Context) error {
	_, err := eicb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (eicb *EventInviteCreateBulk) ExecX(ctx context.Context) {
	if err := eicb.Exec(ctx); err != nil {
		panic(err)
	}
}