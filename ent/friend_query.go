// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/abc3354/CODEV-back/ent/friend"
	"github.com/abc3354/CODEV-back/ent/predicate"
	"github.com/abc3354/CODEV-back/ent/profile"
	"github.com/google/uuid"
)

// FriendQuery is the builder for querying Friend entities.
type FriendQuery struct {
	config
	limit       *int
	offset      *int
	unique      *bool
	order       []OrderFunc
	fields      []string
	inters      []Interceptor
	predicates  []predicate.Friend
	withProfile *ProfileQuery
	withFriend  *ProfileQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the FriendQuery builder.
func (fq *FriendQuery) Where(ps ...predicate.Friend) *FriendQuery {
	fq.predicates = append(fq.predicates, ps...)
	return fq
}

// Limit the number of records to be returned by this query.
func (fq *FriendQuery) Limit(limit int) *FriendQuery {
	fq.limit = &limit
	return fq
}

// Offset to start from.
func (fq *FriendQuery) Offset(offset int) *FriendQuery {
	fq.offset = &offset
	return fq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (fq *FriendQuery) Unique(unique bool) *FriendQuery {
	fq.unique = &unique
	return fq
}

// Order specifies how the records should be ordered.
func (fq *FriendQuery) Order(o ...OrderFunc) *FriendQuery {
	fq.order = append(fq.order, o...)
	return fq
}

// QueryProfile chains the current query on the "profile" edge.
func (fq *FriendQuery) QueryProfile() *ProfileQuery {
	query := (&ProfileClient{config: fq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := fq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := fq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(friend.Table, friend.ProfileColumn, selector),
			sqlgraph.To(profile.Table, profile.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, friend.ProfileTable, friend.ProfileColumn),
		)
		fromU = sqlgraph.SetNeighbors(fq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryFriend chains the current query on the "friend" edge.
func (fq *FriendQuery) QueryFriend() *ProfileQuery {
	query := (&ProfileClient{config: fq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := fq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := fq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(friend.Table, friend.FriendColumn, selector),
			sqlgraph.To(profile.Table, profile.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, friend.FriendTable, friend.FriendColumn),
		)
		fromU = sqlgraph.SetNeighbors(fq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Friend entity from the query.
// Returns a *NotFoundError when no Friend was found.
func (fq *FriendQuery) First(ctx context.Context) (*Friend, error) {
	nodes, err := fq.Limit(1).All(newQueryContext(ctx, TypeFriend, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{friend.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (fq *FriendQuery) FirstX(ctx context.Context) *Friend {
	node, err := fq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// Only returns a single Friend entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Friend entity is found.
// Returns a *NotFoundError when no Friend entities are found.
func (fq *FriendQuery) Only(ctx context.Context) (*Friend, error) {
	nodes, err := fq.Limit(2).All(newQueryContext(ctx, TypeFriend, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{friend.Label}
	default:
		return nil, &NotSingularError{friend.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (fq *FriendQuery) OnlyX(ctx context.Context) *Friend {
	node, err := fq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// All executes the query and returns a list of Friends.
func (fq *FriendQuery) All(ctx context.Context) ([]*Friend, error) {
	ctx = newQueryContext(ctx, TypeFriend, "All")
	if err := fq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Friend, *FriendQuery]()
	return withInterceptors[[]*Friend](ctx, fq, qr, fq.inters)
}

// AllX is like All, but panics if an error occurs.
func (fq *FriendQuery) AllX(ctx context.Context) []*Friend {
	nodes, err := fq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// Count returns the count of the given query.
func (fq *FriendQuery) Count(ctx context.Context) (int, error) {
	ctx = newQueryContext(ctx, TypeFriend, "Count")
	if err := fq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, fq, querierCount[*FriendQuery](), fq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (fq *FriendQuery) CountX(ctx context.Context) int {
	count, err := fq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (fq *FriendQuery) Exist(ctx context.Context) (bool, error) {
	ctx = newQueryContext(ctx, TypeFriend, "Exist")
	switch _, err := fq.First(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (fq *FriendQuery) ExistX(ctx context.Context) bool {
	exist, err := fq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the FriendQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (fq *FriendQuery) Clone() *FriendQuery {
	if fq == nil {
		return nil
	}
	return &FriendQuery{
		config:      fq.config,
		limit:       fq.limit,
		offset:      fq.offset,
		order:       append([]OrderFunc{}, fq.order...),
		inters:      append([]Interceptor{}, fq.inters...),
		predicates:  append([]predicate.Friend{}, fq.predicates...),
		withProfile: fq.withProfile.Clone(),
		withFriend:  fq.withFriend.Clone(),
		// clone intermediate query.
		sql:    fq.sql.Clone(),
		path:   fq.path,
		unique: fq.unique,
	}
}

// WithProfile tells the query-builder to eager-load the nodes that are connected to
// the "profile" edge. The optional arguments are used to configure the query builder of the edge.
func (fq *FriendQuery) WithProfile(opts ...func(*ProfileQuery)) *FriendQuery {
	query := (&ProfileClient{config: fq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	fq.withProfile = query
	return fq
}

// WithFriend tells the query-builder to eager-load the nodes that are connected to
// the "friend" edge. The optional arguments are used to configure the query builder of the edge.
func (fq *FriendQuery) WithFriend(opts ...func(*ProfileQuery)) *FriendQuery {
	query := (&ProfileClient{config: fq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	fq.withFriend = query
	return fq
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
//	client.Friend.Query().
//		GroupBy(friend.FieldProfileID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (fq *FriendQuery) GroupBy(field string, fields ...string) *FriendGroupBy {
	fq.fields = append([]string{field}, fields...)
	grbuild := &FriendGroupBy{build: fq}
	grbuild.flds = &fq.fields
	grbuild.label = friend.Label
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
//	client.Friend.Query().
//		Select(friend.FieldProfileID).
//		Scan(ctx, &v)
func (fq *FriendQuery) Select(fields ...string) *FriendSelect {
	fq.fields = append(fq.fields, fields...)
	sbuild := &FriendSelect{FriendQuery: fq}
	sbuild.label = friend.Label
	sbuild.flds, sbuild.scan = &fq.fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a FriendSelect configured with the given aggregations.
func (fq *FriendQuery) Aggregate(fns ...AggregateFunc) *FriendSelect {
	return fq.Select().Aggregate(fns...)
}

func (fq *FriendQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range fq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, fq); err != nil {
				return err
			}
		}
	}
	for _, f := range fq.fields {
		if !friend.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if fq.path != nil {
		prev, err := fq.path(ctx)
		if err != nil {
			return err
		}
		fq.sql = prev
	}
	return nil
}

func (fq *FriendQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Friend, error) {
	var (
		nodes       = []*Friend{}
		_spec       = fq.querySpec()
		loadedTypes = [2]bool{
			fq.withProfile != nil,
			fq.withFriend != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Friend).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Friend{config: fq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, fq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := fq.withProfile; query != nil {
		if err := fq.loadProfile(ctx, query, nodes, nil,
			func(n *Friend, e *Profile) { n.Edges.Profile = e }); err != nil {
			return nil, err
		}
	}
	if query := fq.withFriend; query != nil {
		if err := fq.loadFriend(ctx, query, nodes, nil,
			func(n *Friend, e *Profile) { n.Edges.Friend = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (fq *FriendQuery) loadProfile(ctx context.Context, query *ProfileQuery, nodes []*Friend, init func(*Friend), assign func(*Friend, *Profile)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*Friend)
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
func (fq *FriendQuery) loadFriend(ctx context.Context, query *ProfileQuery, nodes []*Friend, init func(*Friend), assign func(*Friend, *Profile)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*Friend)
	for i := range nodes {
		fk := nodes[i].FriendID
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
			return fmt.Errorf(`unexpected foreign-key "friend_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (fq *FriendQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := fq.querySpec()
	_spec.Unique = false
	_spec.Node.Columns = nil
	return sqlgraph.CountNodes(ctx, fq.driver, _spec)
}

func (fq *FriendQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   friend.Table,
			Columns: friend.Columns,
		},
		From:   fq.sql,
		Unique: true,
	}
	if unique := fq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := fq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		for i := range fields {
			_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
		}
	}
	if ps := fq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := fq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := fq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := fq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (fq *FriendQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(fq.driver.Dialect())
	t1 := builder.Table(friend.Table)
	columns := fq.fields
	if len(columns) == 0 {
		columns = friend.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if fq.sql != nil {
		selector = fq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if fq.unique != nil && *fq.unique {
		selector.Distinct()
	}
	for _, p := range fq.predicates {
		p(selector)
	}
	for _, p := range fq.order {
		p(selector)
	}
	if offset := fq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := fq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// FriendGroupBy is the group-by builder for Friend entities.
type FriendGroupBy struct {
	selector
	build *FriendQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (fgb *FriendGroupBy) Aggregate(fns ...AggregateFunc) *FriendGroupBy {
	fgb.fns = append(fgb.fns, fns...)
	return fgb
}

// Scan applies the selector query and scans the result into the given value.
func (fgb *FriendGroupBy) Scan(ctx context.Context, v any) error {
	ctx = newQueryContext(ctx, TypeFriend, "GroupBy")
	if err := fgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*FriendQuery, *FriendGroupBy](ctx, fgb.build, fgb, fgb.build.inters, v)
}

func (fgb *FriendGroupBy) sqlScan(ctx context.Context, root *FriendQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(fgb.fns))
	for _, fn := range fgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*fgb.flds)+len(fgb.fns))
		for _, f := range *fgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*fgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := fgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// FriendSelect is the builder for selecting fields of Friend entities.
type FriendSelect struct {
	*FriendQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (fs *FriendSelect) Aggregate(fns ...AggregateFunc) *FriendSelect {
	fs.fns = append(fs.fns, fns...)
	return fs
}

// Scan applies the selector query and scans the result into the given value.
func (fs *FriendSelect) Scan(ctx context.Context, v any) error {
	ctx = newQueryContext(ctx, TypeFriend, "Select")
	if err := fs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*FriendQuery, *FriendSelect](ctx, fs.FriendQuery, fs, fs.inters, v)
}

func (fs *FriendSelect) sqlScan(ctx context.Context, root *FriendQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(fs.fns))
	for _, fn := range fs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*fs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := fs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
