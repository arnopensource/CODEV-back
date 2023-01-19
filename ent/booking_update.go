// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/abc3354/CODEV-back/ent/booking"
	"github.com/abc3354/CODEV-back/ent/predicate"
)

// BookingUpdate is the builder for updating Booking entities.
type BookingUpdate struct {
	config
	hooks    []Hook
	mutation *BookingMutation
}

// Where appends a list predicates to the BookingUpdate builder.
func (bu *BookingUpdate) Where(ps ...predicate.Booking) *BookingUpdate {
	bu.mutation.Where(ps...)
	return bu
}

// SetName sets the "name" field.
func (bu *BookingUpdate) SetName(s string) *BookingUpdate {
	bu.mutation.SetName(s)
	return bu
}

// SetStart sets the "start" field.
func (bu *BookingUpdate) SetStart(t time.Time) *BookingUpdate {
	bu.mutation.SetStart(t)
	return bu
}

// SetEnd sets the "end" field.
func (bu *BookingUpdate) SetEnd(t time.Time) *BookingUpdate {
	bu.mutation.SetEnd(t)
	return bu
}

// Mutation returns the BookingMutation object of the builder.
func (bu *BookingUpdate) Mutation() *BookingMutation {
	return bu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (bu *BookingUpdate) Save(ctx context.Context) (int, error) {
	return withHooks[int, BookingMutation](ctx, bu.sqlSave, bu.mutation, bu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (bu *BookingUpdate) SaveX(ctx context.Context) int {
	affected, err := bu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (bu *BookingUpdate) Exec(ctx context.Context) error {
	_, err := bu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bu *BookingUpdate) ExecX(ctx context.Context) {
	if err := bu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (bu *BookingUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   booking.Table,
			Columns: booking.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: booking.FieldID,
			},
		},
	}
	if ps := bu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := bu.mutation.Name(); ok {
		_spec.SetField(booking.FieldName, field.TypeString, value)
	}
	if value, ok := bu.mutation.Start(); ok {
		_spec.SetField(booking.FieldStart, field.TypeTime, value)
	}
	if value, ok := bu.mutation.End(); ok {
		_spec.SetField(booking.FieldEnd, field.TypeTime, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, bu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{booking.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	bu.mutation.done = true
	return n, nil
}

// BookingUpdateOne is the builder for updating a single Booking entity.
type BookingUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *BookingMutation
}

// SetName sets the "name" field.
func (buo *BookingUpdateOne) SetName(s string) *BookingUpdateOne {
	buo.mutation.SetName(s)
	return buo
}

// SetStart sets the "start" field.
func (buo *BookingUpdateOne) SetStart(t time.Time) *BookingUpdateOne {
	buo.mutation.SetStart(t)
	return buo
}

// SetEnd sets the "end" field.
func (buo *BookingUpdateOne) SetEnd(t time.Time) *BookingUpdateOne {
	buo.mutation.SetEnd(t)
	return buo
}

// Mutation returns the BookingMutation object of the builder.
func (buo *BookingUpdateOne) Mutation() *BookingMutation {
	return buo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (buo *BookingUpdateOne) Select(field string, fields ...string) *BookingUpdateOne {
	buo.fields = append([]string{field}, fields...)
	return buo
}

// Save executes the query and returns the updated Booking entity.
func (buo *BookingUpdateOne) Save(ctx context.Context) (*Booking, error) {
	return withHooks[*Booking, BookingMutation](ctx, buo.sqlSave, buo.mutation, buo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (buo *BookingUpdateOne) SaveX(ctx context.Context) *Booking {
	node, err := buo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (buo *BookingUpdateOne) Exec(ctx context.Context) error {
	_, err := buo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (buo *BookingUpdateOne) ExecX(ctx context.Context) {
	if err := buo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (buo *BookingUpdateOne) sqlSave(ctx context.Context) (_node *Booking, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   booking.Table,
			Columns: booking.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: booking.FieldID,
			},
		},
	}
	id, ok := buo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Booking.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := buo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, booking.FieldID)
		for _, f := range fields {
			if !booking.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != booking.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := buo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := buo.mutation.Name(); ok {
		_spec.SetField(booking.FieldName, field.TypeString, value)
	}
	if value, ok := buo.mutation.Start(); ok {
		_spec.SetField(booking.FieldStart, field.TypeTime, value)
	}
	if value, ok := buo.mutation.End(); ok {
		_spec.SetField(booking.FieldEnd, field.TypeTime, value)
	}
	_node = &Booking{config: buo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, buo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{booking.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	buo.mutation.done = true
	return _node, nil
}
