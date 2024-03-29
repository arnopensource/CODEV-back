// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/abc3354/CODEV-back/ent/booking"
	"github.com/abc3354/CODEV-back/ent/event"
	"github.com/abc3354/CODEV-back/ent/eventinvite"
	"github.com/abc3354/CODEV-back/ent/friend"
	"github.com/abc3354/CODEV-back/ent/member"
	"github.com/abc3354/CODEV-back/ent/predicate"
	"github.com/abc3354/CODEV-back/ent/profile"
	"github.com/abc3354/CODEV-back/ent/room"
	"github.com/google/uuid"
)

// ProfileQuery is the builder for querying Profile entities.
type ProfileQuery struct {
	config
	limit            *int
	offset           *int
	unique           *bool
	order            []OrderFunc
	fields           []string
	inters           []Interceptor
	predicates       []predicate.Profile
	withFriends      *ProfileQuery
	withBookings     *RoomQuery
	withEvents       *EventQuery
	withInvitedTo    *EventQuery
	withFriendsData  *FriendQuery
	withBookingsData *BookingQuery
	withEventsData   *MemberQuery
	withInvitesData  *EventInviteQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the ProfileQuery builder.
func (pq *ProfileQuery) Where(ps ...predicate.Profile) *ProfileQuery {
	pq.predicates = append(pq.predicates, ps...)
	return pq
}

// Limit the number of records to be returned by this query.
func (pq *ProfileQuery) Limit(limit int) *ProfileQuery {
	pq.limit = &limit
	return pq
}

// Offset to start from.
func (pq *ProfileQuery) Offset(offset int) *ProfileQuery {
	pq.offset = &offset
	return pq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (pq *ProfileQuery) Unique(unique bool) *ProfileQuery {
	pq.unique = &unique
	return pq
}

// Order specifies how the records should be ordered.
func (pq *ProfileQuery) Order(o ...OrderFunc) *ProfileQuery {
	pq.order = append(pq.order, o...)
	return pq
}

// QueryFriends chains the current query on the "friends" edge.
func (pq *ProfileQuery) QueryFriends() *ProfileQuery {
	query := (&ProfileClient{config: pq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := pq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := pq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(profile.Table, profile.FieldID, selector),
			sqlgraph.To(profile.Table, profile.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, profile.FriendsTable, profile.FriendsPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(pq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryBookings chains the current query on the "bookings" edge.
func (pq *ProfileQuery) QueryBookings() *RoomQuery {
	query := (&RoomClient{config: pq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := pq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := pq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(profile.Table, profile.FieldID, selector),
			sqlgraph.To(room.Table, room.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, profile.BookingsTable, profile.BookingsPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(pq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryEvents chains the current query on the "events" edge.
func (pq *ProfileQuery) QueryEvents() *EventQuery {
	query := (&EventClient{config: pq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := pq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := pq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(profile.Table, profile.FieldID, selector),
			sqlgraph.To(event.Table, event.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, profile.EventsTable, profile.EventsPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(pq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryInvitedTo chains the current query on the "invitedTo" edge.
func (pq *ProfileQuery) QueryInvitedTo() *EventQuery {
	query := (&EventClient{config: pq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := pq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := pq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(profile.Table, profile.FieldID, selector),
			sqlgraph.To(event.Table, event.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, profile.InvitedToTable, profile.InvitedToPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(pq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryFriendsData chains the current query on the "friends_data" edge.
func (pq *ProfileQuery) QueryFriendsData() *FriendQuery {
	query := (&FriendClient{config: pq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := pq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := pq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(profile.Table, profile.FieldID, selector),
			sqlgraph.To(friend.Table, friend.ProfileColumn),
			sqlgraph.Edge(sqlgraph.O2M, true, profile.FriendsDataTable, profile.FriendsDataColumn),
		)
		fromU = sqlgraph.SetNeighbors(pq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryBookingsData chains the current query on the "bookings_data" edge.
func (pq *ProfileQuery) QueryBookingsData() *BookingQuery {
	query := (&BookingClient{config: pq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := pq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := pq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(profile.Table, profile.FieldID, selector),
			sqlgraph.To(booking.Table, booking.ProfileColumn),
			sqlgraph.Edge(sqlgraph.O2M, true, profile.BookingsDataTable, profile.BookingsDataColumn),
		)
		fromU = sqlgraph.SetNeighbors(pq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryEventsData chains the current query on the "events_data" edge.
func (pq *ProfileQuery) QueryEventsData() *MemberQuery {
	query := (&MemberClient{config: pq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := pq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := pq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(profile.Table, profile.FieldID, selector),
			sqlgraph.To(member.Table, member.ProfileColumn),
			sqlgraph.Edge(sqlgraph.O2M, true, profile.EventsDataTable, profile.EventsDataColumn),
		)
		fromU = sqlgraph.SetNeighbors(pq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryInvitesData chains the current query on the "invites_data" edge.
func (pq *ProfileQuery) QueryInvitesData() *EventInviteQuery {
	query := (&EventInviteClient{config: pq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := pq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := pq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(profile.Table, profile.FieldID, selector),
			sqlgraph.To(eventinvite.Table, eventinvite.ProfileColumn),
			sqlgraph.Edge(sqlgraph.O2M, true, profile.InvitesDataTable, profile.InvitesDataColumn),
		)
		fromU = sqlgraph.SetNeighbors(pq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Profile entity from the query.
// Returns a *NotFoundError when no Profile was found.
func (pq *ProfileQuery) First(ctx context.Context) (*Profile, error) {
	nodes, err := pq.Limit(1).All(newQueryContext(ctx, TypeProfile, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{profile.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (pq *ProfileQuery) FirstX(ctx context.Context) *Profile {
	node, err := pq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Profile ID from the query.
// Returns a *NotFoundError when no Profile ID was found.
func (pq *ProfileQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = pq.Limit(1).IDs(newQueryContext(ctx, TypeProfile, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{profile.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (pq *ProfileQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := pq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Profile entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Profile entity is found.
// Returns a *NotFoundError when no Profile entities are found.
func (pq *ProfileQuery) Only(ctx context.Context) (*Profile, error) {
	nodes, err := pq.Limit(2).All(newQueryContext(ctx, TypeProfile, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{profile.Label}
	default:
		return nil, &NotSingularError{profile.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (pq *ProfileQuery) OnlyX(ctx context.Context) *Profile {
	node, err := pq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Profile ID in the query.
// Returns a *NotSingularError when more than one Profile ID is found.
// Returns a *NotFoundError when no entities are found.
func (pq *ProfileQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = pq.Limit(2).IDs(newQueryContext(ctx, TypeProfile, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{profile.Label}
	default:
		err = &NotSingularError{profile.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (pq *ProfileQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := pq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Profiles.
func (pq *ProfileQuery) All(ctx context.Context) ([]*Profile, error) {
	ctx = newQueryContext(ctx, TypeProfile, "All")
	if err := pq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Profile, *ProfileQuery]()
	return withInterceptors[[]*Profile](ctx, pq, qr, pq.inters)
}

// AllX is like All, but panics if an error occurs.
func (pq *ProfileQuery) AllX(ctx context.Context) []*Profile {
	nodes, err := pq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Profile IDs.
func (pq *ProfileQuery) IDs(ctx context.Context) ([]uuid.UUID, error) {
	var ids []uuid.UUID
	ctx = newQueryContext(ctx, TypeProfile, "IDs")
	if err := pq.Select(profile.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (pq *ProfileQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := pq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (pq *ProfileQuery) Count(ctx context.Context) (int, error) {
	ctx = newQueryContext(ctx, TypeProfile, "Count")
	if err := pq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, pq, querierCount[*ProfileQuery](), pq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (pq *ProfileQuery) CountX(ctx context.Context) int {
	count, err := pq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (pq *ProfileQuery) Exist(ctx context.Context) (bool, error) {
	ctx = newQueryContext(ctx, TypeProfile, "Exist")
	switch _, err := pq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (pq *ProfileQuery) ExistX(ctx context.Context) bool {
	exist, err := pq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the ProfileQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (pq *ProfileQuery) Clone() *ProfileQuery {
	if pq == nil {
		return nil
	}
	return &ProfileQuery{
		config:           pq.config,
		limit:            pq.limit,
		offset:           pq.offset,
		order:            append([]OrderFunc{}, pq.order...),
		inters:           append([]Interceptor{}, pq.inters...),
		predicates:       append([]predicate.Profile{}, pq.predicates...),
		withFriends:      pq.withFriends.Clone(),
		withBookings:     pq.withBookings.Clone(),
		withEvents:       pq.withEvents.Clone(),
		withInvitedTo:    pq.withInvitedTo.Clone(),
		withFriendsData:  pq.withFriendsData.Clone(),
		withBookingsData: pq.withBookingsData.Clone(),
		withEventsData:   pq.withEventsData.Clone(),
		withInvitesData:  pq.withInvitesData.Clone(),
		// clone intermediate query.
		sql:    pq.sql.Clone(),
		path:   pq.path,
		unique: pq.unique,
	}
}

// WithFriends tells the query-builder to eager-load the nodes that are connected to
// the "friends" edge. The optional arguments are used to configure the query builder of the edge.
func (pq *ProfileQuery) WithFriends(opts ...func(*ProfileQuery)) *ProfileQuery {
	query := (&ProfileClient{config: pq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	pq.withFriends = query
	return pq
}

// WithBookings tells the query-builder to eager-load the nodes that are connected to
// the "bookings" edge. The optional arguments are used to configure the query builder of the edge.
func (pq *ProfileQuery) WithBookings(opts ...func(*RoomQuery)) *ProfileQuery {
	query := (&RoomClient{config: pq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	pq.withBookings = query
	return pq
}

// WithEvents tells the query-builder to eager-load the nodes that are connected to
// the "events" edge. The optional arguments are used to configure the query builder of the edge.
func (pq *ProfileQuery) WithEvents(opts ...func(*EventQuery)) *ProfileQuery {
	query := (&EventClient{config: pq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	pq.withEvents = query
	return pq
}

// WithInvitedTo tells the query-builder to eager-load the nodes that are connected to
// the "invitedTo" edge. The optional arguments are used to configure the query builder of the edge.
func (pq *ProfileQuery) WithInvitedTo(opts ...func(*EventQuery)) *ProfileQuery {
	query := (&EventClient{config: pq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	pq.withInvitedTo = query
	return pq
}

// WithFriendsData tells the query-builder to eager-load the nodes that are connected to
// the "friends_data" edge. The optional arguments are used to configure the query builder of the edge.
func (pq *ProfileQuery) WithFriendsData(opts ...func(*FriendQuery)) *ProfileQuery {
	query := (&FriendClient{config: pq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	pq.withFriendsData = query
	return pq
}

// WithBookingsData tells the query-builder to eager-load the nodes that are connected to
// the "bookings_data" edge. The optional arguments are used to configure the query builder of the edge.
func (pq *ProfileQuery) WithBookingsData(opts ...func(*BookingQuery)) *ProfileQuery {
	query := (&BookingClient{config: pq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	pq.withBookingsData = query
	return pq
}

// WithEventsData tells the query-builder to eager-load the nodes that are connected to
// the "events_data" edge. The optional arguments are used to configure the query builder of the edge.
func (pq *ProfileQuery) WithEventsData(opts ...func(*MemberQuery)) *ProfileQuery {
	query := (&MemberClient{config: pq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	pq.withEventsData = query
	return pq
}

// WithInvitesData tells the query-builder to eager-load the nodes that are connected to
// the "invites_data" edge. The optional arguments are used to configure the query builder of the edge.
func (pq *ProfileQuery) WithInvitesData(opts ...func(*EventInviteQuery)) *ProfileQuery {
	query := (&EventInviteClient{config: pq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	pq.withInvitesData = query
	return pq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Firstname string `json:"firstname,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Profile.Query().
//		GroupBy(profile.FieldFirstname).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (pq *ProfileQuery) GroupBy(field string, fields ...string) *ProfileGroupBy {
	pq.fields = append([]string{field}, fields...)
	grbuild := &ProfileGroupBy{build: pq}
	grbuild.flds = &pq.fields
	grbuild.label = profile.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Firstname string `json:"firstname,omitempty"`
//	}
//
//	client.Profile.Query().
//		Select(profile.FieldFirstname).
//		Scan(ctx, &v)
func (pq *ProfileQuery) Select(fields ...string) *ProfileSelect {
	pq.fields = append(pq.fields, fields...)
	sbuild := &ProfileSelect{ProfileQuery: pq}
	sbuild.label = profile.Label
	sbuild.flds, sbuild.scan = &pq.fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a ProfileSelect configured with the given aggregations.
func (pq *ProfileQuery) Aggregate(fns ...AggregateFunc) *ProfileSelect {
	return pq.Select().Aggregate(fns...)
}

func (pq *ProfileQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range pq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, pq); err != nil {
				return err
			}
		}
	}
	for _, f := range pq.fields {
		if !profile.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if pq.path != nil {
		prev, err := pq.path(ctx)
		if err != nil {
			return err
		}
		pq.sql = prev
	}
	return nil
}

func (pq *ProfileQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Profile, error) {
	var (
		nodes       = []*Profile{}
		_spec       = pq.querySpec()
		loadedTypes = [8]bool{
			pq.withFriends != nil,
			pq.withBookings != nil,
			pq.withEvents != nil,
			pq.withInvitedTo != nil,
			pq.withFriendsData != nil,
			pq.withBookingsData != nil,
			pq.withEventsData != nil,
			pq.withInvitesData != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Profile).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Profile{config: pq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, pq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := pq.withFriends; query != nil {
		if err := pq.loadFriends(ctx, query, nodes,
			func(n *Profile) { n.Edges.Friends = []*Profile{} },
			func(n *Profile, e *Profile) { n.Edges.Friends = append(n.Edges.Friends, e) }); err != nil {
			return nil, err
		}
	}
	if query := pq.withBookings; query != nil {
		if err := pq.loadBookings(ctx, query, nodes,
			func(n *Profile) { n.Edges.Bookings = []*Room{} },
			func(n *Profile, e *Room) { n.Edges.Bookings = append(n.Edges.Bookings, e) }); err != nil {
			return nil, err
		}
	}
	if query := pq.withEvents; query != nil {
		if err := pq.loadEvents(ctx, query, nodes,
			func(n *Profile) { n.Edges.Events = []*Event{} },
			func(n *Profile, e *Event) { n.Edges.Events = append(n.Edges.Events, e) }); err != nil {
			return nil, err
		}
	}
	if query := pq.withInvitedTo; query != nil {
		if err := pq.loadInvitedTo(ctx, query, nodes,
			func(n *Profile) { n.Edges.InvitedTo = []*Event{} },
			func(n *Profile, e *Event) { n.Edges.InvitedTo = append(n.Edges.InvitedTo, e) }); err != nil {
			return nil, err
		}
	}
	if query := pq.withFriendsData; query != nil {
		if err := pq.loadFriendsData(ctx, query, nodes,
			func(n *Profile) { n.Edges.FriendsData = []*Friend{} },
			func(n *Profile, e *Friend) { n.Edges.FriendsData = append(n.Edges.FriendsData, e) }); err != nil {
			return nil, err
		}
	}
	if query := pq.withBookingsData; query != nil {
		if err := pq.loadBookingsData(ctx, query, nodes,
			func(n *Profile) { n.Edges.BookingsData = []*Booking{} },
			func(n *Profile, e *Booking) { n.Edges.BookingsData = append(n.Edges.BookingsData, e) }); err != nil {
			return nil, err
		}
	}
	if query := pq.withEventsData; query != nil {
		if err := pq.loadEventsData(ctx, query, nodes,
			func(n *Profile) { n.Edges.EventsData = []*Member{} },
			func(n *Profile, e *Member) { n.Edges.EventsData = append(n.Edges.EventsData, e) }); err != nil {
			return nil, err
		}
	}
	if query := pq.withInvitesData; query != nil {
		if err := pq.loadInvitesData(ctx, query, nodes,
			func(n *Profile) { n.Edges.InvitesData = []*EventInvite{} },
			func(n *Profile, e *EventInvite) { n.Edges.InvitesData = append(n.Edges.InvitesData, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (pq *ProfileQuery) loadFriends(ctx context.Context, query *ProfileQuery, nodes []*Profile, init func(*Profile), assign func(*Profile, *Profile)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[uuid.UUID]*Profile)
	nids := make(map[uuid.UUID]map[*Profile]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(profile.FriendsTable)
		s.Join(joinT).On(s.C(profile.FieldID), joinT.C(profile.FriendsPrimaryKey[1]))
		s.Where(sql.InValues(joinT.C(profile.FriendsPrimaryKey[0]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(profile.FriendsPrimaryKey[0]))
		s.AppendSelect(columns...)
		s.SetDistinct(false)
	})
	if err := query.prepareQuery(ctx); err != nil {
		return err
	}
	neighbors, err := query.sqlAll(ctx, func(_ context.Context, spec *sqlgraph.QuerySpec) {
		assign := spec.Assign
		values := spec.ScanValues
		spec.ScanValues = func(columns []string) ([]any, error) {
			values, err := values(columns[1:])
			if err != nil {
				return nil, err
			}
			return append([]any{new(uuid.UUID)}, values...), nil
		}
		spec.Assign = func(columns []string, values []any) error {
			outValue := *values[0].(*uuid.UUID)
			inValue := *values[1].(*uuid.UUID)
			if nids[inValue] == nil {
				nids[inValue] = map[*Profile]struct{}{byID[outValue]: {}}
				return assign(columns[1:], values[1:])
			}
			nids[inValue][byID[outValue]] = struct{}{}
			return nil
		}
	})
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected "friends" node returned %v`, n.ID)
		}
		for kn := range nodes {
			assign(kn, n)
		}
	}
	return nil
}
func (pq *ProfileQuery) loadBookings(ctx context.Context, query *RoomQuery, nodes []*Profile, init func(*Profile), assign func(*Profile, *Room)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[uuid.UUID]*Profile)
	nids := make(map[int]map[*Profile]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(profile.BookingsTable)
		s.Join(joinT).On(s.C(room.FieldID), joinT.C(profile.BookingsPrimaryKey[1]))
		s.Where(sql.InValues(joinT.C(profile.BookingsPrimaryKey[0]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(profile.BookingsPrimaryKey[0]))
		s.AppendSelect(columns...)
		s.SetDistinct(false)
	})
	if err := query.prepareQuery(ctx); err != nil {
		return err
	}
	neighbors, err := query.sqlAll(ctx, func(_ context.Context, spec *sqlgraph.QuerySpec) {
		assign := spec.Assign
		values := spec.ScanValues
		spec.ScanValues = func(columns []string) ([]any, error) {
			values, err := values(columns[1:])
			if err != nil {
				return nil, err
			}
			return append([]any{new(uuid.UUID)}, values...), nil
		}
		spec.Assign = func(columns []string, values []any) error {
			outValue := *values[0].(*uuid.UUID)
			inValue := int(values[1].(*sql.NullInt64).Int64)
			if nids[inValue] == nil {
				nids[inValue] = map[*Profile]struct{}{byID[outValue]: {}}
				return assign(columns[1:], values[1:])
			}
			nids[inValue][byID[outValue]] = struct{}{}
			return nil
		}
	})
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected "bookings" node returned %v`, n.ID)
		}
		for kn := range nodes {
			assign(kn, n)
		}
	}
	return nil
}
func (pq *ProfileQuery) loadEvents(ctx context.Context, query *EventQuery, nodes []*Profile, init func(*Profile), assign func(*Profile, *Event)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[uuid.UUID]*Profile)
	nids := make(map[int]map[*Profile]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(profile.EventsTable)
		s.Join(joinT).On(s.C(event.FieldID), joinT.C(profile.EventsPrimaryKey[0]))
		s.Where(sql.InValues(joinT.C(profile.EventsPrimaryKey[1]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(profile.EventsPrimaryKey[1]))
		s.AppendSelect(columns...)
		s.SetDistinct(false)
	})
	if err := query.prepareQuery(ctx); err != nil {
		return err
	}
	neighbors, err := query.sqlAll(ctx, func(_ context.Context, spec *sqlgraph.QuerySpec) {
		assign := spec.Assign
		values := spec.ScanValues
		spec.ScanValues = func(columns []string) ([]any, error) {
			values, err := values(columns[1:])
			if err != nil {
				return nil, err
			}
			return append([]any{new(uuid.UUID)}, values...), nil
		}
		spec.Assign = func(columns []string, values []any) error {
			outValue := *values[0].(*uuid.UUID)
			inValue := int(values[1].(*sql.NullInt64).Int64)
			if nids[inValue] == nil {
				nids[inValue] = map[*Profile]struct{}{byID[outValue]: {}}
				return assign(columns[1:], values[1:])
			}
			nids[inValue][byID[outValue]] = struct{}{}
			return nil
		}
	})
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected "events" node returned %v`, n.ID)
		}
		for kn := range nodes {
			assign(kn, n)
		}
	}
	return nil
}
func (pq *ProfileQuery) loadInvitedTo(ctx context.Context, query *EventQuery, nodes []*Profile, init func(*Profile), assign func(*Profile, *Event)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[uuid.UUID]*Profile)
	nids := make(map[int]map[*Profile]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(profile.InvitedToTable)
		s.Join(joinT).On(s.C(event.FieldID), joinT.C(profile.InvitedToPrimaryKey[0]))
		s.Where(sql.InValues(joinT.C(profile.InvitedToPrimaryKey[1]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(profile.InvitedToPrimaryKey[1]))
		s.AppendSelect(columns...)
		s.SetDistinct(false)
	})
	if err := query.prepareQuery(ctx); err != nil {
		return err
	}
	neighbors, err := query.sqlAll(ctx, func(_ context.Context, spec *sqlgraph.QuerySpec) {
		assign := spec.Assign
		values := spec.ScanValues
		spec.ScanValues = func(columns []string) ([]any, error) {
			values, err := values(columns[1:])
			if err != nil {
				return nil, err
			}
			return append([]any{new(uuid.UUID)}, values...), nil
		}
		spec.Assign = func(columns []string, values []any) error {
			outValue := *values[0].(*uuid.UUID)
			inValue := int(values[1].(*sql.NullInt64).Int64)
			if nids[inValue] == nil {
				nids[inValue] = map[*Profile]struct{}{byID[outValue]: {}}
				return assign(columns[1:], values[1:])
			}
			nids[inValue][byID[outValue]] = struct{}{}
			return nil
		}
	})
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected "invitedTo" node returned %v`, n.ID)
		}
		for kn := range nodes {
			assign(kn, n)
		}
	}
	return nil
}
func (pq *ProfileQuery) loadFriendsData(ctx context.Context, query *FriendQuery, nodes []*Profile, init func(*Profile), assign func(*Profile, *Friend)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[uuid.UUID]*Profile)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.Where(predicate.Friend(func(s *sql.Selector) {
		s.Where(sql.InValues(profile.FriendsDataColumn, fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.ProfileID
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "profile_id" returned %v for node %v`, fk, n)
		}
		assign(node, n)
	}
	return nil
}
func (pq *ProfileQuery) loadBookingsData(ctx context.Context, query *BookingQuery, nodes []*Profile, init func(*Profile), assign func(*Profile, *Booking)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[uuid.UUID]*Profile)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.Where(predicate.Booking(func(s *sql.Selector) {
		s.Where(sql.InValues(profile.BookingsDataColumn, fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.ProfileID
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "profile_id" returned %v for node %v`, fk, n)
		}
		assign(node, n)
	}
	return nil
}
func (pq *ProfileQuery) loadEventsData(ctx context.Context, query *MemberQuery, nodes []*Profile, init func(*Profile), assign func(*Profile, *Member)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[uuid.UUID]*Profile)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.Where(predicate.Member(func(s *sql.Selector) {
		s.Where(sql.InValues(profile.EventsDataColumn, fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.ProfileID
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "profile_id" returned %v for node %v`, fk, n)
		}
		assign(node, n)
	}
	return nil
}
func (pq *ProfileQuery) loadInvitesData(ctx context.Context, query *EventInviteQuery, nodes []*Profile, init func(*Profile), assign func(*Profile, *EventInvite)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[uuid.UUID]*Profile)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.Where(predicate.EventInvite(func(s *sql.Selector) {
		s.Where(sql.InValues(profile.InvitesDataColumn, fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.ProfileID
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "profile_id" returned %v for node %v`, fk, n)
		}
		assign(node, n)
	}
	return nil
}

func (pq *ProfileQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := pq.querySpec()
	_spec.Node.Columns = pq.fields
	if len(pq.fields) > 0 {
		_spec.Unique = pq.unique != nil && *pq.unique
	}
	return sqlgraph.CountNodes(ctx, pq.driver, _spec)
}

func (pq *ProfileQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   profile.Table,
			Columns: profile.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: profile.FieldID,
			},
		},
		From:   pq.sql,
		Unique: true,
	}
	if unique := pq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := pq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, profile.FieldID)
		for i := range fields {
			if fields[i] != profile.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := pq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := pq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := pq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := pq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (pq *ProfileQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(pq.driver.Dialect())
	t1 := builder.Table(profile.Table)
	columns := pq.fields
	if len(columns) == 0 {
		columns = profile.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if pq.sql != nil {
		selector = pq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if pq.unique != nil && *pq.unique {
		selector.Distinct()
	}
	for _, p := range pq.predicates {
		p(selector)
	}
	for _, p := range pq.order {
		p(selector)
	}
	if offset := pq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := pq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ProfileGroupBy is the group-by builder for Profile entities.
type ProfileGroupBy struct {
	selector
	build *ProfileQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (pgb *ProfileGroupBy) Aggregate(fns ...AggregateFunc) *ProfileGroupBy {
	pgb.fns = append(pgb.fns, fns...)
	return pgb
}

// Scan applies the selector query and scans the result into the given value.
func (pgb *ProfileGroupBy) Scan(ctx context.Context, v any) error {
	ctx = newQueryContext(ctx, TypeProfile, "GroupBy")
	if err := pgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ProfileQuery, *ProfileGroupBy](ctx, pgb.build, pgb, pgb.build.inters, v)
}

func (pgb *ProfileGroupBy) sqlScan(ctx context.Context, root *ProfileQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(pgb.fns))
	for _, fn := range pgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*pgb.flds)+len(pgb.fns))
		for _, f := range *pgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*pgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := pgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// ProfileSelect is the builder for selecting fields of Profile entities.
type ProfileSelect struct {
	*ProfileQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ps *ProfileSelect) Aggregate(fns ...AggregateFunc) *ProfileSelect {
	ps.fns = append(ps.fns, fns...)
	return ps
}

// Scan applies the selector query and scans the result into the given value.
func (ps *ProfileSelect) Scan(ctx context.Context, v any) error {
	ctx = newQueryContext(ctx, TypeProfile, "Select")
	if err := ps.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ProfileQuery, *ProfileSelect](ctx, ps.ProfileQuery, ps, ps.inters, v)
}

func (ps *ProfileSelect) sqlScan(ctx context.Context, root *ProfileQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ps.fns))
	for _, fn := range ps.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ps.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ps.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
