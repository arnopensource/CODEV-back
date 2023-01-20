// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/abc3354/CODEV-back/ent/predicate"
	"github.com/abc3354/CODEV-back/ent/salle"
	"github.com/abc3354/CODEV-back/ent/salledisponible"
)

// SalleDisponibleQuery is the builder for querying SalleDisponible entities.
type SalleDisponibleQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	inters     []Interceptor
	predicates []predicate.SalleDisponible
	withSalle  *SalleQuery
	withFKs    bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the SalleDisponibleQuery builder.
func (sdq *SalleDisponibleQuery) Where(ps ...predicate.SalleDisponible) *SalleDisponibleQuery {
	sdq.predicates = append(sdq.predicates, ps...)
	return sdq
}

// Limit the number of records to be returned by this query.
func (sdq *SalleDisponibleQuery) Limit(limit int) *SalleDisponibleQuery {
	sdq.limit = &limit
	return sdq
}

// Offset to start from.
func (sdq *SalleDisponibleQuery) Offset(offset int) *SalleDisponibleQuery {
	sdq.offset = &offset
	return sdq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (sdq *SalleDisponibleQuery) Unique(unique bool) *SalleDisponibleQuery {
	sdq.unique = &unique
	return sdq
}

// Order specifies how the records should be ordered.
func (sdq *SalleDisponibleQuery) Order(o ...OrderFunc) *SalleDisponibleQuery {
	sdq.order = append(sdq.order, o...)
	return sdq
}

// QuerySalle chains the current query on the "salle" edge.
func (sdq *SalleDisponibleQuery) QuerySalle() *SalleQuery {
	query := (&SalleClient{config: sdq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := sdq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := sdq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(salledisponible.Table, salledisponible.FieldID, selector),
			sqlgraph.To(salle.Table, salle.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, salledisponible.SalleTable, salledisponible.SalleColumn),
		)
		fromU = sqlgraph.SetNeighbors(sdq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first SalleDisponible entity from the query.
// Returns a *NotFoundError when no SalleDisponible was found.
func (sdq *SalleDisponibleQuery) First(ctx context.Context) (*SalleDisponible, error) {
	nodes, err := sdq.Limit(1).All(newQueryContext(ctx, TypeSalleDisponible, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{salledisponible.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (sdq *SalleDisponibleQuery) FirstX(ctx context.Context) *SalleDisponible {
	node, err := sdq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first SalleDisponible ID from the query.
// Returns a *NotFoundError when no SalleDisponible ID was found.
func (sdq *SalleDisponibleQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = sdq.Limit(1).IDs(newQueryContext(ctx, TypeSalleDisponible, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{salledisponible.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (sdq *SalleDisponibleQuery) FirstIDX(ctx context.Context) int {
	id, err := sdq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single SalleDisponible entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one SalleDisponible entity is found.
// Returns a *NotFoundError when no SalleDisponible entities are found.
func (sdq *SalleDisponibleQuery) Only(ctx context.Context) (*SalleDisponible, error) {
	nodes, err := sdq.Limit(2).All(newQueryContext(ctx, TypeSalleDisponible, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{salledisponible.Label}
	default:
		return nil, &NotSingularError{salledisponible.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (sdq *SalleDisponibleQuery) OnlyX(ctx context.Context) *SalleDisponible {
	node, err := sdq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only SalleDisponible ID in the query.
// Returns a *NotSingularError when more than one SalleDisponible ID is found.
// Returns a *NotFoundError when no entities are found.
func (sdq *SalleDisponibleQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = sdq.Limit(2).IDs(newQueryContext(ctx, TypeSalleDisponible, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{salledisponible.Label}
	default:
		err = &NotSingularError{salledisponible.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (sdq *SalleDisponibleQuery) OnlyIDX(ctx context.Context) int {
	id, err := sdq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of SalleDisponibles.
func (sdq *SalleDisponibleQuery) All(ctx context.Context) ([]*SalleDisponible, error) {
	ctx = newQueryContext(ctx, TypeSalleDisponible, "All")
	if err := sdq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*SalleDisponible, *SalleDisponibleQuery]()
	return withInterceptors[[]*SalleDisponible](ctx, sdq, qr, sdq.inters)
}

// AllX is like All, but panics if an error occurs.
func (sdq *SalleDisponibleQuery) AllX(ctx context.Context) []*SalleDisponible {
	nodes, err := sdq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of SalleDisponible IDs.
func (sdq *SalleDisponibleQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	ctx = newQueryContext(ctx, TypeSalleDisponible, "IDs")
	if err := sdq.Select(salledisponible.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (sdq *SalleDisponibleQuery) IDsX(ctx context.Context) []int {
	ids, err := sdq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (sdq *SalleDisponibleQuery) Count(ctx context.Context) (int, error) {
	ctx = newQueryContext(ctx, TypeSalleDisponible, "Count")
	if err := sdq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, sdq, querierCount[*SalleDisponibleQuery](), sdq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (sdq *SalleDisponibleQuery) CountX(ctx context.Context) int {
	count, err := sdq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (sdq *SalleDisponibleQuery) Exist(ctx context.Context) (bool, error) {
	ctx = newQueryContext(ctx, TypeSalleDisponible, "Exist")
	switch _, err := sdq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (sdq *SalleDisponibleQuery) ExistX(ctx context.Context) bool {
	exist, err := sdq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the SalleDisponibleQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (sdq *SalleDisponibleQuery) Clone() *SalleDisponibleQuery {
	if sdq == nil {
		return nil
	}
	return &SalleDisponibleQuery{
		config:     sdq.config,
		limit:      sdq.limit,
		offset:     sdq.offset,
		order:      append([]OrderFunc{}, sdq.order...),
		inters:     append([]Interceptor{}, sdq.inters...),
		predicates: append([]predicate.SalleDisponible{}, sdq.predicates...),
		withSalle:  sdq.withSalle.Clone(),
		// clone intermediate query.
		sql:    sdq.sql.Clone(),
		path:   sdq.path,
		unique: sdq.unique,
	}
}

// WithSalle tells the query-builder to eager-load the nodes that are connected to
// the "salle" edge. The optional arguments are used to configure the query builder of the edge.
func (sdq *SalleDisponibleQuery) WithSalle(opts ...func(*SalleQuery)) *SalleDisponibleQuery {
	query := (&SalleClient{config: sdq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	sdq.withSalle = query
	return sdq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		IDSalle string `json:"id_salle,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.SalleDisponible.Query().
//		GroupBy(salledisponible.FieldIDSalle).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (sdq *SalleDisponibleQuery) GroupBy(field string, fields ...string) *SalleDisponibleGroupBy {
	sdq.fields = append([]string{field}, fields...)
	grbuild := &SalleDisponibleGroupBy{build: sdq}
	grbuild.flds = &sdq.fields
	grbuild.label = salledisponible.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		IDSalle string `json:"id_salle,omitempty"`
//	}
//
//	client.SalleDisponible.Query().
//		Select(salledisponible.FieldIDSalle).
//		Scan(ctx, &v)
func (sdq *SalleDisponibleQuery) Select(fields ...string) *SalleDisponibleSelect {
	sdq.fields = append(sdq.fields, fields...)
	sbuild := &SalleDisponibleSelect{SalleDisponibleQuery: sdq}
	sbuild.label = salledisponible.Label
	sbuild.flds, sbuild.scan = &sdq.fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a SalleDisponibleSelect configured with the given aggregations.
func (sdq *SalleDisponibleQuery) Aggregate(fns ...AggregateFunc) *SalleDisponibleSelect {
	return sdq.Select().Aggregate(fns...)
}

func (sdq *SalleDisponibleQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range sdq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, sdq); err != nil {
				return err
			}
		}
	}
	for _, f := range sdq.fields {
		if !salledisponible.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if sdq.path != nil {
		prev, err := sdq.path(ctx)
		if err != nil {
			return err
		}
		sdq.sql = prev
	}
	return nil
}

func (sdq *SalleDisponibleQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*SalleDisponible, error) {
	var (
		nodes       = []*SalleDisponible{}
		withFKs     = sdq.withFKs
		_spec       = sdq.querySpec()
		loadedTypes = [1]bool{
			sdq.withSalle != nil,
		}
	)
	if sdq.withSalle != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, salledisponible.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*SalleDisponible).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &SalleDisponible{config: sdq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, sdq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := sdq.withSalle; query != nil {
		if err := sdq.loadSalle(ctx, query, nodes, nil,
			func(n *SalleDisponible, e *Salle) { n.Edges.Salle = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (sdq *SalleDisponibleQuery) loadSalle(ctx context.Context, query *SalleQuery, nodes []*SalleDisponible, init func(*SalleDisponible), assign func(*SalleDisponible, *Salle)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*SalleDisponible)
	for i := range nodes {
		if nodes[i].salle_disponibilite == nil {
			continue
		}
		fk := *nodes[i].salle_disponibilite
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
			return fmt.Errorf(`unexpected foreign-key "salle_disponibilite" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (sdq *SalleDisponibleQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := sdq.querySpec()
	_spec.Node.Columns = sdq.fields
	if len(sdq.fields) > 0 {
		_spec.Unique = sdq.unique != nil && *sdq.unique
	}
	return sqlgraph.CountNodes(ctx, sdq.driver, _spec)
}

func (sdq *SalleDisponibleQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   salledisponible.Table,
			Columns: salledisponible.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: salledisponible.FieldID,
			},
		},
		From:   sdq.sql,
		Unique: true,
	}
	if unique := sdq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := sdq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, salledisponible.FieldID)
		for i := range fields {
			if fields[i] != salledisponible.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := sdq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := sdq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := sdq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := sdq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (sdq *SalleDisponibleQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(sdq.driver.Dialect())
	t1 := builder.Table(salledisponible.Table)
	columns := sdq.fields
	if len(columns) == 0 {
		columns = salledisponible.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if sdq.sql != nil {
		selector = sdq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if sdq.unique != nil && *sdq.unique {
		selector.Distinct()
	}
	for _, p := range sdq.predicates {
		p(selector)
	}
	for _, p := range sdq.order {
		p(selector)
	}
	if offset := sdq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := sdq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// SalleDisponibleGroupBy is the group-by builder for SalleDisponible entities.
type SalleDisponibleGroupBy struct {
	selector
	build *SalleDisponibleQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (sdgb *SalleDisponibleGroupBy) Aggregate(fns ...AggregateFunc) *SalleDisponibleGroupBy {
	sdgb.fns = append(sdgb.fns, fns...)
	return sdgb
}

// Scan applies the selector query and scans the result into the given value.
func (sdgb *SalleDisponibleGroupBy) Scan(ctx context.Context, v any) error {
	ctx = newQueryContext(ctx, TypeSalleDisponible, "GroupBy")
	if err := sdgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*SalleDisponibleQuery, *SalleDisponibleGroupBy](ctx, sdgb.build, sdgb, sdgb.build.inters, v)
}

func (sdgb *SalleDisponibleGroupBy) sqlScan(ctx context.Context, root *SalleDisponibleQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(sdgb.fns))
	for _, fn := range sdgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*sdgb.flds)+len(sdgb.fns))
		for _, f := range *sdgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*sdgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := sdgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// SalleDisponibleSelect is the builder for selecting fields of SalleDisponible entities.
type SalleDisponibleSelect struct {
	*SalleDisponibleQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (sds *SalleDisponibleSelect) Aggregate(fns ...AggregateFunc) *SalleDisponibleSelect {
	sds.fns = append(sds.fns, fns...)
	return sds
}

// Scan applies the selector query and scans the result into the given value.
func (sds *SalleDisponibleSelect) Scan(ctx context.Context, v any) error {
	ctx = newQueryContext(ctx, TypeSalleDisponible, "Select")
	if err := sds.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*SalleDisponibleQuery, *SalleDisponibleSelect](ctx, sds.SalleDisponibleQuery, sds, sds.inters, v)
}

func (sds *SalleDisponibleSelect) sqlScan(ctx context.Context, root *SalleDisponibleQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(sds.fns))
	for _, fn := range sds.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*sds.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := sds.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}