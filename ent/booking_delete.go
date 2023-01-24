// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/abc3354/CODEV-back/ent/booking"
	"github.com/abc3354/CODEV-back/ent/predicate"
)

// BookingDelete is the builder for deleting a Booking entity.
type BookingDelete struct {
	config
	hooks    []Hook
	mutation *BookingMutation
}

// Where appends a list predicates to the BookingDelete builder.
func (bd *BookingDelete) Where(ps ...predicate.Booking) *BookingDelete {
	bd.mutation.Where(ps...)
	return bd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (bd *BookingDelete) Exec(ctx context.Context) (int, error) {
	return withHooks[int, BookingMutation](ctx, bd.sqlExec, bd.mutation, bd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (bd *BookingDelete) ExecX(ctx context.Context) int {
	n, err := bd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (bd *BookingDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: booking.Table,
		},
	}
	if ps := bd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, bd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	bd.mutation.done = true
	return affected, err
}

// BookingDeleteOne is the builder for deleting a single Booking entity.
type BookingDeleteOne struct {
	bd *BookingDelete
}

// Exec executes the deletion query.
func (bdo *BookingDeleteOne) Exec(ctx context.Context) error {
	n, err := bdo.bd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{booking.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (bdo *BookingDeleteOne) ExecX(ctx context.Context) {
	bdo.bd.ExecX(ctx)
}
