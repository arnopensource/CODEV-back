// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/abc3354/CODEV-back/ent/predicate"
	"github.com/abc3354/CODEV-back/ent/profile"
	"github.com/abc3354/CODEV-back/ent/reservation"
	"github.com/abc3354/CODEV-back/ent/salle"
	"github.com/google/uuid"
)

// ReservationQuery is the builder for querying Reservation entities.
type ReservationQuery struct {
	config
	limit       *int
	offset      *int
	unique      *bool
	order       []OrderFunc
	fields      []string
	inters      []Interceptor
	predicates  []predicate.Reservation
	withProfile *ProfileQuery
	withSalle   *SalleQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the ReservationQuery builder.
func (rq *ReservationQuery) Where(ps ...predicate.Reservation) *ReservationQuery {
	rq.predicates = append(rq.predicates, ps...)
	return rq
}

// Limit the number of records to be returned by this query.
func (rq *ReservationQuery) Limit(limit int) *ReservationQuery {
	rq.limit = &limit
	return rq
}

// Offset to start from.
func (rq *ReservationQuery) Offset(offset int) *ReservationQuery {
	rq.offset = &offset
	return rq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (rq *ReservationQuery) Unique(unique bool) *ReservationQuery {
	rq.unique = &unique
	return rq
}

// Order specifies how the records should be ordered.
func (rq *ReservationQuery) Order(o ...OrderFunc) *ReservationQuery {
	rq.order = append(rq.order, o...)
	return rq
}

// QueryProfile chains the current query on the "profile" edge.
func (rq *ReservationQuery) QueryProfile() *ProfileQuery {
	query := (&ProfileClient{config: rq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := rq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := rq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(reservation.Table, reservation.ProfileColumn, selector),
			sqlgraph.To(profile.Table, profile.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, reservation.ProfileTable, reservation.ProfileColumn),
		)
		fromU = sqlgraph.SetNeighbors(rq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QuerySalle chains the current query on the "salle" edge.
func (rq *ReservationQuery) QuerySalle() *SalleQuery {
	query := (&SalleClient{config: rq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := rq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := rq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(reservation.Table, reservation.SalleColumn, selector),
			sqlgraph.To(salle.Table, salle.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, reservation.SalleTable, reservation.SalleColumn),
		)
		fromU = sqlgraph.SetNeighbors(rq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Reservation entity from the query.
// Returns a *NotFoundError when no Reservation was found.
func (rq *ReservationQuery) First(ctx context.Context) (*Reservation, error) {
	nodes, err := rq.Limit(1).All(newQueryContext(ctx, TypeReservation, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{reservation.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (rq *ReservationQuery) FirstX(ctx context.Context) *Reservation {
	node, err := rq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// Only returns a single Reservation entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Reservation entity is found.
// Returns a *NotFoundError when no Reservation entities are found.
func (rq *ReservationQuery) Only(ctx context.Context) (*Reservation, error) {
	nodes, err := rq.Limit(2).All(newQueryContext(ctx, TypeReservation, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{reservation.Label}
	default:
		return nil, &NotSingularError{reservation.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (rq *ReservationQuery) OnlyX(ctx context.Context) *Reservation {
	node, err := rq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// All executes the query and returns a list of Reservations.
func (rq *ReservationQuery) All(ctx context.Context) ([]*Reservation, error) {
	ctx = newQueryContext(ctx, TypeReservation, "All")
	if err := rq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Reservation, *ReservationQuery]()
	return withInterceptors[[]*Reservation](ctx, rq, qr, rq.inters)
}

// AllX is like All, but panics if an error occurs.
func (rq *ReservationQuery) AllX(ctx context.Context) []*Reservation {
	nodes, err := rq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// Count returns the count of the given query.
func (rq *ReservationQuery) Count(ctx context.Context) (int, error) {
	ctx = newQueryContext(ctx, TypeReservation, "Count")
	if err := rq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, rq, querierCount[*ReservationQuery](), rq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (rq *ReservationQuery) CountX(ctx context.Context) int {
	count, err := rq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (rq *ReservationQuery) Exist(ctx context.Context) (bool, error) {
	ctx = newQueryContext(ctx, TypeReservation, "Exist")
	switch _, err := rq.First(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (rq *ReservationQuery) ExistX(ctx context.Context) bool {
	exist, err := rq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the ReservationQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (rq *ReservationQuery) Clone() *ReservationQuery {
	if rq == nil {
		return nil
	}
	return &ReservationQuery{
		config:      rq.config,
		limit:       rq.limit,
		offset:      rq.offset,
		order:       append([]OrderFunc{}, rq.order...),
		inters:      append([]Interceptor{}, rq.inters...),
		predicates:  append([]predicate.Reservation{}, rq.predicates...),
		withProfile: rq.withProfile.Clone(),
		withSalle:   rq.withSalle.Clone(),
		// clone intermediate query.
		sql:    rq.sql.Clone(),
		path:   rq.path,
		unique: rq.unique,
	}
}

// WithProfile tells the query-builder to eager-load the nodes that are connected to
// the "profile" edge. The optional arguments are used to configure the query builder of the edge.
func (rq *ReservationQuery) WithProfile(opts ...func(*ProfileQuery)) *ReservationQuery {
	query := (&ProfileClient{config: rq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	rq.withProfile = query
	return rq
}

// WithSalle tells the query-builder to eager-load the nodes that are connected to
// the "salle" edge. The optional arguments are used to configure the query builder of the edge.
func (rq *ReservationQuery) WithSalle(opts ...func(*SalleQuery)) *ReservationQuery {
	query := (&SalleClient{config: rq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	rq.withSalle = query
	return rq
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
//	client.Reservation.Query().
//		GroupBy(reservation.FieldProfileID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (rq *ReservationQuery) GroupBy(field string, fields ...string) *ReservationGroupBy {
	rq.fields = append([]string{field}, fields...)
	grbuild := &ReservationGroupBy{build: rq}
	grbuild.flds = &rq.fields
	grbuild.label = reservation.Label
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
//	client.Reservation.Query().
//		Select(reservation.FieldProfileID).
//		Scan(ctx, &v)
func (rq *ReservationQuery) Select(fields ...string) *ReservationSelect {
	rq.fields = append(rq.fields, fields...)
	sbuild := &ReservationSelect{ReservationQuery: rq}
	sbuild.label = reservation.Label
	sbuild.flds, sbuild.scan = &rq.fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a ReservationSelect configured with the given aggregations.
func (rq *ReservationQuery) Aggregate(fns ...AggregateFunc) *ReservationSelect {
	return rq.Select().Aggregate(fns...)
}

func (rq *ReservationQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range rq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, rq); err != nil {
				return err
			}
		}
	}
	for _, f := range rq.fields {
		if !reservation.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if rq.path != nil {
		prev, err := rq.path(ctx)
		if err != nil {
			return err
		}
		rq.sql = prev
	}
	return nil
}

func (rq *ReservationQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Reservation, error) {
	var (
		nodes       = []*Reservation{}
		_spec       = rq.querySpec()
		loadedTypes = [2]bool{
			rq.withProfile != nil,
			rq.withSalle != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Reservation).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Reservation{config: rq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, rq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := rq.withProfile; query != nil {
		if err := rq.loadProfile(ctx, query, nodes, nil,
			func(n *Reservation, e *Profile) { n.Edges.Profile = e }); err != nil {
			return nil, err
		}
	}
	if query := rq.withSalle; query != nil {
		if err := rq.loadSalle(ctx, query, nodes, nil,
			func(n *Reservation, e *Salle) { n.Edges.Salle = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (rq *ReservationQuery) loadProfile(ctx context.Context, query *ProfileQuery, nodes []*Reservation, init func(*Reservation), assign func(*Reservation, *Profile)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*Reservation)
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
func (rq *ReservationQuery) loadSalle(ctx context.Context, query *SalleQuery, nodes []*Reservation, init func(*Reservation), assign func(*Reservation, *Salle)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*Reservation)
	for i := range nodes {
		fk := nodes[i].SalleID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	query.Where(salle.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "salle_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (rq *ReservationQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := rq.querySpec()
	_spec.Unique = false
	_spec.Node.Columns = nil
	return sqlgraph.CountNodes(ctx, rq.driver, _spec)
}

func (rq *ReservationQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   reservation.Table,
			Columns: reservation.Columns,
		},
		From:   rq.sql,
		Unique: true,
	}
	if unique := rq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := rq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		for i := range fields {
			_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
		}
	}
	if ps := rq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := rq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := rq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := rq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (rq *ReservationQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(rq.driver.Dialect())
	t1 := builder.Table(reservation.Table)
	columns := rq.fields
	if len(columns) == 0 {
		columns = reservation.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if rq.sql != nil {
		selector = rq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if rq.unique != nil && *rq.unique {
		selector.Distinct()
	}
	for _, p := range rq.predicates {
		p(selector)
	}
	for _, p := range rq.order {
		p(selector)
	}
	if offset := rq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := rq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ReservationGroupBy is the group-by builder for Reservation entities.
type ReservationGroupBy struct {
	selector
	build *ReservationQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (rgb *ReservationGroupBy) Aggregate(fns ...AggregateFunc) *ReservationGroupBy {
	rgb.fns = append(rgb.fns, fns...)
	return rgb
}

// Scan applies the selector query and scans the result into the given value.
func (rgb *ReservationGroupBy) Scan(ctx context.Context, v any) error {
	ctx = newQueryContext(ctx, TypeReservation, "GroupBy")
	if err := rgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ReservationQuery, *ReservationGroupBy](ctx, rgb.build, rgb, rgb.build.inters, v)
}

func (rgb *ReservationGroupBy) sqlScan(ctx context.Context, root *ReservationQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(rgb.fns))
	for _, fn := range rgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*rgb.flds)+len(rgb.fns))
		for _, f := range *rgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*rgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := rgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// ReservationSelect is the builder for selecting fields of Reservation entities.
type ReservationSelect struct {
	*ReservationQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (rs *ReservationSelect) Aggregate(fns ...AggregateFunc) *ReservationSelect {
	rs.fns = append(rs.fns, fns...)
	return rs
}

// Scan applies the selector query and scans the result into the given value.
func (rs *ReservationSelect) Scan(ctx context.Context, v any) error {
	ctx = newQueryContext(ctx, TypeReservation, "Select")
	if err := rs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ReservationQuery, *ReservationSelect](ctx, rs.ReservationQuery, rs, rs.inters, v)
}

func (rs *ReservationSelect) sqlScan(ctx context.Context, root *ReservationQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(rs.fns))
	for _, fn := range rs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*rs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := rs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
