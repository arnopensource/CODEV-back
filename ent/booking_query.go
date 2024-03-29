// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/abc3354/CODEV-back/ent/booking"
	"github.com/abc3354/CODEV-back/ent/predicate"
	"github.com/abc3354/CODEV-back/ent/profile"
	"github.com/abc3354/CODEV-back/ent/room"
	"github.com/google/uuid"
)

// BookingQuery is the builder for querying Booking entities.
type BookingQuery struct {
	config
	limit       *int
	offset      *int
	unique      *bool
	order       []OrderFunc
	fields      []string
	inters      []Interceptor
	predicates  []predicate.Booking
	withProfile *ProfileQuery
	withRoom    *RoomQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the BookingQuery builder.
func (bq *BookingQuery) Where(ps ...predicate.Booking) *BookingQuery {
	bq.predicates = append(bq.predicates, ps...)
	return bq
}

// Limit the number of records to be returned by this query.
func (bq *BookingQuery) Limit(limit int) *BookingQuery {
	bq.limit = &limit
	return bq
}

// Offset to start from.
func (bq *BookingQuery) Offset(offset int) *BookingQuery {
	bq.offset = &offset
	return bq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (bq *BookingQuery) Unique(unique bool) *BookingQuery {
	bq.unique = &unique
	return bq
}

// Order specifies how the records should be ordered.
func (bq *BookingQuery) Order(o ...OrderFunc) *BookingQuery {
	bq.order = append(bq.order, o...)
	return bq
}

// QueryProfile chains the current query on the "profile" edge.
func (bq *BookingQuery) QueryProfile() *ProfileQuery {
	query := (&ProfileClient{config: bq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := bq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := bq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(booking.Table, booking.ProfileColumn, selector),
			sqlgraph.To(profile.Table, profile.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, booking.ProfileTable, booking.ProfileColumn),
		)
		fromU = sqlgraph.SetNeighbors(bq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryRoom chains the current query on the "room" edge.
func (bq *BookingQuery) QueryRoom() *RoomQuery {
	query := (&RoomClient{config: bq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := bq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := bq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(booking.Table, booking.RoomColumn, selector),
			sqlgraph.To(room.Table, room.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, booking.RoomTable, booking.RoomColumn),
		)
		fromU = sqlgraph.SetNeighbors(bq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Booking entity from the query.
// Returns a *NotFoundError when no Booking was found.
func (bq *BookingQuery) First(ctx context.Context) (*Booking, error) {
	nodes, err := bq.Limit(1).All(newQueryContext(ctx, TypeBooking, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{booking.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (bq *BookingQuery) FirstX(ctx context.Context) *Booking {
	node, err := bq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// Only returns a single Booking entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Booking entity is found.
// Returns a *NotFoundError when no Booking entities are found.
func (bq *BookingQuery) Only(ctx context.Context) (*Booking, error) {
	nodes, err := bq.Limit(2).All(newQueryContext(ctx, TypeBooking, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{booking.Label}
	default:
		return nil, &NotSingularError{booking.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (bq *BookingQuery) OnlyX(ctx context.Context) *Booking {
	node, err := bq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// All executes the query and returns a list of Bookings.
func (bq *BookingQuery) All(ctx context.Context) ([]*Booking, error) {
	ctx = newQueryContext(ctx, TypeBooking, "All")
	if err := bq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Booking, *BookingQuery]()
	return withInterceptors[[]*Booking](ctx, bq, qr, bq.inters)
}

// AllX is like All, but panics if an error occurs.
func (bq *BookingQuery) AllX(ctx context.Context) []*Booking {
	nodes, err := bq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// Count returns the count of the given query.
func (bq *BookingQuery) Count(ctx context.Context) (int, error) {
	ctx = newQueryContext(ctx, TypeBooking, "Count")
	if err := bq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, bq, querierCount[*BookingQuery](), bq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (bq *BookingQuery) CountX(ctx context.Context) int {
	count, err := bq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (bq *BookingQuery) Exist(ctx context.Context) (bool, error) {
	ctx = newQueryContext(ctx, TypeBooking, "Exist")
	switch _, err := bq.First(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (bq *BookingQuery) ExistX(ctx context.Context) bool {
	exist, err := bq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the BookingQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (bq *BookingQuery) Clone() *BookingQuery {
	if bq == nil {
		return nil
	}
	return &BookingQuery{
		config:      bq.config,
		limit:       bq.limit,
		offset:      bq.offset,
		order:       append([]OrderFunc{}, bq.order...),
		inters:      append([]Interceptor{}, bq.inters...),
		predicates:  append([]predicate.Booking{}, bq.predicates...),
		withProfile: bq.withProfile.Clone(),
		withRoom:    bq.withRoom.Clone(),
		// clone intermediate query.
		sql:    bq.sql.Clone(),
		path:   bq.path,
		unique: bq.unique,
	}
}

// WithProfile tells the query-builder to eager-load the nodes that are connected to
// the "profile" edge. The optional arguments are used to configure the query builder of the edge.
func (bq *BookingQuery) WithProfile(opts ...func(*ProfileQuery)) *BookingQuery {
	query := (&ProfileClient{config: bq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	bq.withProfile = query
	return bq
}

// WithRoom tells the query-builder to eager-load the nodes that are connected to
// the "room" edge. The optional arguments are used to configure the query builder of the edge.
func (bq *BookingQuery) WithRoom(opts ...func(*RoomQuery)) *BookingQuery {
	query := (&RoomClient{config: bq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	bq.withRoom = query
	return bq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		ProfileID uuid.UUID `json:"profile_id,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Booking.Query().
//		GroupBy(booking.FieldProfileID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (bq *BookingQuery) GroupBy(field string, fields ...string) *BookingGroupBy {
	bq.fields = append([]string{field}, fields...)
	grbuild := &BookingGroupBy{build: bq}
	grbuild.flds = &bq.fields
	grbuild.label = booking.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		ProfileID uuid.UUID `json:"profile_id,omitempty"`
//	}
//
//	client.Booking.Query().
//		Select(booking.FieldProfileID).
//		Scan(ctx, &v)
func (bq *BookingQuery) Select(fields ...string) *BookingSelect {
	bq.fields = append(bq.fields, fields...)
	sbuild := &BookingSelect{BookingQuery: bq}
	sbuild.label = booking.Label
	sbuild.flds, sbuild.scan = &bq.fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a BookingSelect configured with the given aggregations.
func (bq *BookingQuery) Aggregate(fns ...AggregateFunc) *BookingSelect {
	return bq.Select().Aggregate(fns...)
}

func (bq *BookingQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range bq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, bq); err != nil {
				return err
			}
		}
	}
	for _, f := range bq.fields {
		if !booking.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if bq.path != nil {
		prev, err := bq.path(ctx)
		if err != nil {
			return err
		}
		bq.sql = prev
	}
	return nil
}

func (bq *BookingQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Booking, error) {
	var (
		nodes       = []*Booking{}
		_spec       = bq.querySpec()
		loadedTypes = [2]bool{
			bq.withProfile != nil,
			bq.withRoom != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Booking).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Booking{config: bq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, bq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := bq.withProfile; query != nil {
		if err := bq.loadProfile(ctx, query, nodes, nil,
			func(n *Booking, e *Profile) { n.Edges.Profile = e }); err != nil {
			return nil, err
		}
	}
	if query := bq.withRoom; query != nil {
		if err := bq.loadRoom(ctx, query, nodes, nil,
			func(n *Booking, e *Room) { n.Edges.Room = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (bq *BookingQuery) loadProfile(ctx context.Context, query *ProfileQuery, nodes []*Booking, init func(*Booking), assign func(*Booking, *Profile)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*Booking)
	for i := range nodes {
		fk := nodes[i].ProfileID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	query.Where(profile.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "profile_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (bq *BookingQuery) loadRoom(ctx context.Context, query *RoomQuery, nodes []*Booking, init func(*Booking), assign func(*Booking, *Room)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*Booking)
	for i := range nodes {
		fk := nodes[i].RoomID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	query.Where(room.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "room_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (bq *BookingQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := bq.querySpec()
	_spec.Unique = false
	_spec.Node.Columns = nil
	return sqlgraph.CountNodes(ctx, bq.driver, _spec)
}

func (bq *BookingQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   booking.Table,
			Columns: booking.Columns,
		},
		From:   bq.sql,
		Unique: true,
	}
	if unique := bq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := bq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		for i := range fields {
			_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
		}
	}
	if ps := bq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := bq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := bq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := bq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (bq *BookingQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(bq.driver.Dialect())
	t1 := builder.Table(booking.Table)
	columns := bq.fields
	if len(columns) == 0 {
		columns = booking.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if bq.sql != nil {
		selector = bq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if bq.unique != nil && *bq.unique {
		selector.Distinct()
	}
	for _, p := range bq.predicates {
		p(selector)
	}
	for _, p := range bq.order {
		p(selector)
	}
	if offset := bq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := bq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// BookingGroupBy is the group-by builder for Booking entities.
type BookingGroupBy struct {
	selector
	build *BookingQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (bgb *BookingGroupBy) Aggregate(fns ...AggregateFunc) *BookingGroupBy {
	bgb.fns = append(bgb.fns, fns...)
	return bgb
}

// Scan applies the selector query and scans the result into the given value.
func (bgb *BookingGroupBy) Scan(ctx context.Context, v any) error {
	ctx = newQueryContext(ctx, TypeBooking, "GroupBy")
	if err := bgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*BookingQuery, *BookingGroupBy](ctx, bgb.build, bgb, bgb.build.inters, v)
}

func (bgb *BookingGroupBy) sqlScan(ctx context.Context, root *BookingQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(bgb.fns))
	for _, fn := range bgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*bgb.flds)+len(bgb.fns))
		for _, f := range *bgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*bgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := bgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// BookingSelect is the builder for selecting fields of Booking entities.
type BookingSelect struct {
	*BookingQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (bs *BookingSelect) Aggregate(fns ...AggregateFunc) *BookingSelect {
	bs.fns = append(bs.fns, fns...)
	return bs
}

// Scan applies the selector query and scans the result into the given value.
func (bs *BookingSelect) Scan(ctx context.Context, v any) error {
	ctx = newQueryContext(ctx, TypeBooking, "Select")
	if err := bs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*BookingQuery, *BookingSelect](ctx, bs.BookingQuery, bs, bs.inters, v)
}

func (bs *BookingSelect) sqlScan(ctx context.Context, root *BookingQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(bs.fns))
	for _, fn := range bs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*bs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := bs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
