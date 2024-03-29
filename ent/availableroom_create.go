// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/abc3354/CODEV-back/ent/availableroom"
	"github.com/abc3354/CODEV-back/ent/room"
)

// AvailableRoomCreate is the builder for creating a AvailableRoom entity.
type AvailableRoomCreate struct {
	config
	mutation *AvailableRoomMutation
	hooks    []Hook
}

// SetStart sets the "start" field.
func (arc *AvailableRoomCreate) SetStart(t time.Time) *AvailableRoomCreate {
	arc.mutation.SetStart(t)
	return arc
}

// SetEnd sets the "end" field.
func (arc *AvailableRoomCreate) SetEnd(t time.Time) *AvailableRoomCreate {
	arc.mutation.SetEnd(t)
	return arc
}

// SetRoomsID sets the "rooms" edge to the Room entity by ID.
func (arc *AvailableRoomCreate) SetRoomsID(id int) *AvailableRoomCreate {
	arc.mutation.SetRoomsID(id)
	return arc
}

// SetNillableRoomsID sets the "rooms" edge to the Room entity by ID if the given value is not nil.
func (arc *AvailableRoomCreate) SetNillableRoomsID(id *int) *AvailableRoomCreate {
	if id != nil {
		arc = arc.SetRoomsID(*id)
	}
	return arc
}

// SetRooms sets the "rooms" edge to the Room entity.
func (arc *AvailableRoomCreate) SetRooms(r *Room) *AvailableRoomCreate {
	return arc.SetRoomsID(r.ID)
}

// Mutation returns the AvailableRoomMutation object of the builder.
func (arc *AvailableRoomCreate) Mutation() *AvailableRoomMutation {
	return arc.mutation
}

// Save creates the AvailableRoom in the database.
func (arc *AvailableRoomCreate) Save(ctx context.Context) (*AvailableRoom, error) {
	return withHooks[*AvailableRoom, AvailableRoomMutation](ctx, arc.sqlSave, arc.mutation, arc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (arc *AvailableRoomCreate) SaveX(ctx context.Context) *AvailableRoom {
	v, err := arc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (arc *AvailableRoomCreate) Exec(ctx context.Context) error {
	_, err := arc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (arc *AvailableRoomCreate) ExecX(ctx context.Context) {
	if err := arc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (arc *AvailableRoomCreate) check() error {
	if _, ok := arc.mutation.Start(); !ok {
		return &ValidationError{Name: "start", err: errors.New(`ent: missing required field "AvailableRoom.start"`)}
	}
	if _, ok := arc.mutation.End(); !ok {
		return &ValidationError{Name: "end", err: errors.New(`ent: missing required field "AvailableRoom.end"`)}
	}
	return nil
}

func (arc *AvailableRoomCreate) sqlSave(ctx context.Context) (*AvailableRoom, error) {
	if err := arc.check(); err != nil {
		return nil, err
	}
	_node, _spec := arc.createSpec()
	if err := sqlgraph.CreateNode(ctx, arc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	arc.mutation.id = &_node.ID
	arc.mutation.done = true
	return _node, nil
}

func (arc *AvailableRoomCreate) createSpec() (*AvailableRoom, *sqlgraph.CreateSpec) {
	var (
		_node = &AvailableRoom{config: arc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: availableroom.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: availableroom.FieldID,
			},
		}
	)
	if value, ok := arc.mutation.Start(); ok {
		_spec.SetField(availableroom.FieldStart, field.TypeTime, value)
		_node.Start = value
	}
	if value, ok := arc.mutation.End(); ok {
		_spec.SetField(availableroom.FieldEnd, field.TypeTime, value)
		_node.End = value
	}
	if nodes := arc.mutation.RoomsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   availableroom.RoomsTable,
			Columns: []string{availableroom.RoomsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: room.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.room_availability = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// AvailableRoomCreateBulk is the builder for creating many AvailableRoom entities in bulk.
type AvailableRoomCreateBulk struct {
	config
	builders []*AvailableRoomCreate
}

// Save creates the AvailableRoom entities in the database.
func (arcb *AvailableRoomCreateBulk) Save(ctx context.Context) ([]*AvailableRoom, error) {
	specs := make([]*sqlgraph.CreateSpec, len(arcb.builders))
	nodes := make([]*AvailableRoom, len(arcb.builders))
	mutators := make([]Mutator, len(arcb.builders))
	for i := range arcb.builders {
		func(i int, root context.Context) {
			builder := arcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*AvailableRoomMutation)
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
					_, err = mutators[i+1].Mutate(root, arcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, arcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
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
		if _, err := mutators[0].Mutate(ctx, arcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (arcb *AvailableRoomCreateBulk) SaveX(ctx context.Context) []*AvailableRoom {
	v, err := arcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (arcb *AvailableRoomCreateBulk) Exec(ctx context.Context) error {
	_, err := arcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (arcb *AvailableRoomCreateBulk) ExecX(ctx context.Context) {
	if err := arcb.Exec(ctx); err != nil {
		panic(err)
	}
}
