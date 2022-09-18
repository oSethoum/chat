// Code generated by ent, DO NOT EDIT.

package ent

import (
	"chat/ent/group"
	"chat/ent/message"
	"chat/ent/user"
	"chat/ent/usergroup"
	"context"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"strconv"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/errcode"
	"github.com/vektah/gqlparser/v2/gqlerror"
	"github.com/vmihailenco/msgpack/v5"
)

// OrderDirection defines the directions in which to order a list of items.
type OrderDirection string

const (
	// OrderDirectionAsc specifies an ascending order.
	OrderDirectionAsc OrderDirection = "ASC"
	// OrderDirectionDesc specifies a descending order.
	OrderDirectionDesc OrderDirection = "DESC"
)

// Validate the order direction value.
func (o OrderDirection) Validate() error {
	if o != OrderDirectionAsc && o != OrderDirectionDesc {
		return fmt.Errorf("%s is not a valid OrderDirection", o)
	}
	return nil
}

// String implements fmt.Stringer interface.
func (o OrderDirection) String() string {
	return string(o)
}

// MarshalGQL implements graphql.Marshaler interface.
func (o OrderDirection) MarshalGQL(w io.Writer) {
	io.WriteString(w, strconv.Quote(o.String()))
}

// UnmarshalGQL implements graphql.Unmarshaler interface.
func (o *OrderDirection) UnmarshalGQL(val interface{}) error {
	str, ok := val.(string)
	if !ok {
		return fmt.Errorf("order direction %T must be a string", val)
	}
	*o = OrderDirection(str)
	return o.Validate()
}

func (o OrderDirection) reverse() OrderDirection {
	if o == OrderDirectionDesc {
		return OrderDirectionAsc
	}
	return OrderDirectionDesc
}

func (o OrderDirection) orderFunc(field string) OrderFunc {
	if o == OrderDirectionDesc {
		return Desc(field)
	}
	return Asc(field)
}

func cursorsToPredicates(direction OrderDirection, after, before *Cursor, field, idField string) []func(s *sql.Selector) {
	var predicates []func(s *sql.Selector)
	if after != nil {
		if after.Value != nil {
			var predicate func([]string, ...interface{}) *sql.Predicate
			if direction == OrderDirectionAsc {
				predicate = sql.CompositeGT
			} else {
				predicate = sql.CompositeLT
			}
			predicates = append(predicates, func(s *sql.Selector) {
				s.Where(predicate(
					s.Columns(field, idField),
					after.Value, after.ID,
				))
			})
		} else {
			var predicate func(string, interface{}) *sql.Predicate
			if direction == OrderDirectionAsc {
				predicate = sql.GT
			} else {
				predicate = sql.LT
			}
			predicates = append(predicates, func(s *sql.Selector) {
				s.Where(predicate(
					s.C(idField),
					after.ID,
				))
			})
		}
	}
	if before != nil {
		if before.Value != nil {
			var predicate func([]string, ...interface{}) *sql.Predicate
			if direction == OrderDirectionAsc {
				predicate = sql.CompositeLT
			} else {
				predicate = sql.CompositeGT
			}
			predicates = append(predicates, func(s *sql.Selector) {
				s.Where(predicate(
					s.Columns(field, idField),
					before.Value, before.ID,
				))
			})
		} else {
			var predicate func(string, interface{}) *sql.Predicate
			if direction == OrderDirectionAsc {
				predicate = sql.LT
			} else {
				predicate = sql.GT
			}
			predicates = append(predicates, func(s *sql.Selector) {
				s.Where(predicate(
					s.C(idField),
					before.ID,
				))
			})
		}
	}
	return predicates
}

// PageInfo of a connection type.
type PageInfo struct {
	HasNextPage     bool    `json:"hasNextPage"`
	HasPreviousPage bool    `json:"hasPreviousPage"`
	StartCursor     *Cursor `json:"startCursor"`
	EndCursor       *Cursor `json:"endCursor"`
}

// Cursor of an edge type.
type Cursor struct {
	ID    int   `msgpack:"i"`
	Value Value `msgpack:"v,omitempty"`
}

// MarshalGQL implements graphql.Marshaler interface.
func (c Cursor) MarshalGQL(w io.Writer) {
	quote := []byte{'"'}
	w.Write(quote)
	defer w.Write(quote)
	wc := base64.NewEncoder(base64.RawStdEncoding, w)
	defer wc.Close()
	_ = msgpack.NewEncoder(wc).Encode(c)
}

// UnmarshalGQL implements graphql.Unmarshaler interface.
func (c *Cursor) UnmarshalGQL(v interface{}) error {
	s, ok := v.(string)
	if !ok {
		return fmt.Errorf("%T is not a string", v)
	}
	if err := msgpack.NewDecoder(
		base64.NewDecoder(
			base64.RawStdEncoding,
			strings.NewReader(s),
		),
	).Decode(c); err != nil {
		return fmt.Errorf("cannot decode cursor: %w", err)
	}
	return nil
}

const errInvalidPagination = "INVALID_PAGINATION"

func validateFirstLast(first, last *int) (err *gqlerror.Error) {
	switch {
	case first != nil && last != nil:
		err = &gqlerror.Error{
			Message: "Passing both `first` and `last` to paginate a connection is not supported.",
		}
	case first != nil && *first < 0:
		err = &gqlerror.Error{
			Message: "`first` on a connection cannot be less than zero.",
		}
		errcode.Set(err, errInvalidPagination)
	case last != nil && *last < 0:
		err = &gqlerror.Error{
			Message: "`last` on a connection cannot be less than zero.",
		}
		errcode.Set(err, errInvalidPagination)
	}
	return err
}

func collectedField(ctx context.Context, path ...string) *graphql.CollectedField {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return nil
	}
	field := fc.Field
	oc := graphql.GetOperationContext(ctx)
walk:
	for _, name := range path {
		for _, f := range graphql.CollectFields(oc, field.Selections, nil) {
			if f.Alias == name {
				field = f
				continue walk
			}
		}
		return nil
	}
	return &field
}

func hasCollectedField(ctx context.Context, path ...string) bool {
	if graphql.GetFieldContext(ctx) == nil {
		return true
	}
	return collectedField(ctx, path...) != nil
}

const (
	edgesField      = "edges"
	nodeField       = "node"
	pageInfoField   = "pageInfo"
	totalCountField = "totalCount"
)

func paginateLimit(first, last *int) int {
	var limit int
	if first != nil {
		limit = *first + 1
	} else if last != nil {
		limit = *last + 1
	}
	return limit
}

// GroupEdge is the edge representation of Group.
type GroupEdge struct {
	Node   *Group `json:"node"`
	Cursor Cursor `json:"cursor"`
}

// GroupConnection is the connection containing edges to Group.
type GroupConnection struct {
	Edges      []*GroupEdge `json:"edges"`
	PageInfo   PageInfo     `json:"pageInfo"`
	TotalCount int          `json:"totalCount"`
}

func (c *GroupConnection) build(nodes []*Group, pager *groupPager, after *Cursor, first *int, before *Cursor, last *int) {
	c.PageInfo.HasNextPage = before != nil
	c.PageInfo.HasPreviousPage = after != nil
	if first != nil && *first+1 == len(nodes) {
		c.PageInfo.HasNextPage = true
		nodes = nodes[:len(nodes)-1]
	} else if last != nil && *last+1 == len(nodes) {
		c.PageInfo.HasPreviousPage = true
		nodes = nodes[:len(nodes)-1]
	}
	var nodeAt func(int) *Group
	if last != nil {
		n := len(nodes) - 1
		nodeAt = func(i int) *Group {
			return nodes[n-i]
		}
	} else {
		nodeAt = func(i int) *Group {
			return nodes[i]
		}
	}
	c.Edges = make([]*GroupEdge, len(nodes))
	for i := range nodes {
		node := nodeAt(i)
		c.Edges[i] = &GroupEdge{
			Node:   node,
			Cursor: pager.toCursor(node),
		}
	}
	if l := len(c.Edges); l > 0 {
		c.PageInfo.StartCursor = &c.Edges[0].Cursor
		c.PageInfo.EndCursor = &c.Edges[l-1].Cursor
	}
	if c.TotalCount == 0 {
		c.TotalCount = len(nodes)
	}
}

// GroupPaginateOption enables pagination customization.
type GroupPaginateOption func(*groupPager) error

// WithGroupOrder configures pagination ordering.
func WithGroupOrder(order *GroupOrder) GroupPaginateOption {
	if order == nil {
		order = DefaultGroupOrder
	}
	o := *order
	return func(pager *groupPager) error {
		if err := o.Direction.Validate(); err != nil {
			return err
		}
		if o.Field == nil {
			o.Field = DefaultGroupOrder.Field
		}
		pager.order = &o
		return nil
	}
}

// WithGroupFilter configures pagination filter.
func WithGroupFilter(filter func(*GroupQuery) (*GroupQuery, error)) GroupPaginateOption {
	return func(pager *groupPager) error {
		if filter == nil {
			return errors.New("GroupQuery filter cannot be nil")
		}
		pager.filter = filter
		return nil
	}
}

type groupPager struct {
	order  *GroupOrder
	filter func(*GroupQuery) (*GroupQuery, error)
}

func newGroupPager(opts []GroupPaginateOption) (*groupPager, error) {
	pager := &groupPager{}
	for _, opt := range opts {
		if err := opt(pager); err != nil {
			return nil, err
		}
	}
	if pager.order == nil {
		pager.order = DefaultGroupOrder
	}
	return pager, nil
}

func (p *groupPager) applyFilter(query *GroupQuery) (*GroupQuery, error) {
	if p.filter != nil {
		return p.filter(query)
	}
	return query, nil
}

func (p *groupPager) toCursor(gr *Group) Cursor {
	return p.order.Field.toCursor(gr)
}

func (p *groupPager) applyCursors(query *GroupQuery, after, before *Cursor) *GroupQuery {
	for _, predicate := range cursorsToPredicates(
		p.order.Direction, after, before,
		p.order.Field.field, DefaultGroupOrder.Field.field,
	) {
		query = query.Where(predicate)
	}
	return query
}

func (p *groupPager) applyOrder(query *GroupQuery, reverse bool) *GroupQuery {
	direction := p.order.Direction
	if reverse {
		direction = direction.reverse()
	}
	query = query.Order(direction.orderFunc(p.order.Field.field))
	if p.order.Field != DefaultGroupOrder.Field {
		query = query.Order(direction.orderFunc(DefaultGroupOrder.Field.field))
	}
	return query
}

func (p *groupPager) orderExpr(reverse bool) sql.Querier {
	direction := p.order.Direction
	if reverse {
		direction = direction.reverse()
	}
	return sql.ExprFunc(func(b *sql.Builder) {
		b.Ident(p.order.Field.field).Pad().WriteString(string(direction))
		if p.order.Field != DefaultGroupOrder.Field {
			b.Comma().Ident(DefaultGroupOrder.Field.field).Pad().WriteString(string(direction))
		}
	})
}

// Paginate executes the query and returns a relay based cursor connection to Group.
func (gr *GroupQuery) Paginate(
	ctx context.Context, after *Cursor, first *int,
	before *Cursor, last *int, opts ...GroupPaginateOption,
) (*GroupConnection, error) {
	if err := validateFirstLast(first, last); err != nil {
		return nil, err
	}
	pager, err := newGroupPager(opts)
	if err != nil {
		return nil, err
	}
	if gr, err = pager.applyFilter(gr); err != nil {
		return nil, err
	}
	conn := &GroupConnection{Edges: []*GroupEdge{}}
	ignoredEdges := !hasCollectedField(ctx, edgesField)
	if hasCollectedField(ctx, totalCountField) || hasCollectedField(ctx, pageInfoField) {
		hasPagination := after != nil || first != nil || before != nil || last != nil
		if hasPagination || ignoredEdges {
			if conn.TotalCount, err = gr.Count(ctx); err != nil {
				return nil, err
			}
			conn.PageInfo.HasNextPage = first != nil && conn.TotalCount > 0
			conn.PageInfo.HasPreviousPage = last != nil && conn.TotalCount > 0
		}
	}
	if ignoredEdges || (first != nil && *first == 0) || (last != nil && *last == 0) {
		return conn, nil
	}

	gr = pager.applyCursors(gr, after, before)
	gr = pager.applyOrder(gr, last != nil)
	if limit := paginateLimit(first, last); limit != 0 {
		gr.Limit(limit)
	}
	if field := collectedField(ctx, edgesField, nodeField); field != nil {
		if err := gr.collectField(ctx, graphql.GetOperationContext(ctx), *field, []string{edgesField, nodeField}); err != nil {
			return nil, err
		}
	}

	nodes, err := gr.All(ctx)
	if err != nil {
		return nil, err
	}
	conn.build(nodes, pager, after, first, before, last)
	return conn, nil
}

var (
	// GroupOrderFieldUpdatedAt orders Group by updated_at.
	GroupOrderFieldUpdatedAt = &GroupOrderField{
		field: group.FieldUpdatedAt,
		toCursor: func(gr *Group) Cursor {
			return Cursor{
				ID:    gr.ID,
				Value: gr.UpdatedAt,
			}
		},
	}
	// GroupOrderFieldCreatedAt orders Group by created_at.
	GroupOrderFieldCreatedAt = &GroupOrderField{
		field: group.FieldCreatedAt,
		toCursor: func(gr *Group) Cursor {
			return Cursor{
				ID:    gr.ID,
				Value: gr.CreatedAt,
			}
		},
	}
	// GroupOrderFieldName orders Group by name.
	GroupOrderFieldName = &GroupOrderField{
		field: group.FieldName,
		toCursor: func(gr *Group) Cursor {
			return Cursor{
				ID:    gr.ID,
				Value: gr.Name,
			}
		},
	}
)

// String implement fmt.Stringer interface.
func (f GroupOrderField) String() string {
	var str string
	switch f.field {
	case group.FieldUpdatedAt:
		str = "UPDATED_AT"
	case group.FieldCreatedAt:
		str = "CREATED_AT"
	case group.FieldName:
		str = "NAME"
	}
	return str
}

// MarshalGQL implements graphql.Marshaler interface.
func (f GroupOrderField) MarshalGQL(w io.Writer) {
	io.WriteString(w, strconv.Quote(f.String()))
}

// UnmarshalGQL implements graphql.Unmarshaler interface.
func (f *GroupOrderField) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("GroupOrderField %T must be a string", v)
	}
	switch str {
	case "UPDATED_AT":
		*f = *GroupOrderFieldUpdatedAt
	case "CREATED_AT":
		*f = *GroupOrderFieldCreatedAt
	case "NAME":
		*f = *GroupOrderFieldName
	default:
		return fmt.Errorf("%s is not a valid GroupOrderField", str)
	}
	return nil
}

// GroupOrderField defines the ordering field of Group.
type GroupOrderField struct {
	field    string
	toCursor func(*Group) Cursor
}

// GroupOrder defines the ordering of Group.
type GroupOrder struct {
	Direction OrderDirection   `json:"direction"`
	Field     *GroupOrderField `json:"field"`
}

// DefaultGroupOrder is the default ordering of Group.
var DefaultGroupOrder = &GroupOrder{
	Direction: OrderDirectionAsc,
	Field: &GroupOrderField{
		field: group.FieldID,
		toCursor: func(gr *Group) Cursor {
			return Cursor{ID: gr.ID}
		},
	},
}

// ToEdge converts Group into GroupEdge.
func (gr *Group) ToEdge(order *GroupOrder) *GroupEdge {
	if order == nil {
		order = DefaultGroupOrder
	}
	return &GroupEdge{
		Node:   gr,
		Cursor: order.Field.toCursor(gr),
	}
}

// MessageEdge is the edge representation of Message.
type MessageEdge struct {
	Node   *Message `json:"node"`
	Cursor Cursor   `json:"cursor"`
}

// MessageConnection is the connection containing edges to Message.
type MessageConnection struct {
	Edges      []*MessageEdge `json:"edges"`
	PageInfo   PageInfo       `json:"pageInfo"`
	TotalCount int            `json:"totalCount"`
}

func (c *MessageConnection) build(nodes []*Message, pager *messagePager, after *Cursor, first *int, before *Cursor, last *int) {
	c.PageInfo.HasNextPage = before != nil
	c.PageInfo.HasPreviousPage = after != nil
	if first != nil && *first+1 == len(nodes) {
		c.PageInfo.HasNextPage = true
		nodes = nodes[:len(nodes)-1]
	} else if last != nil && *last+1 == len(nodes) {
		c.PageInfo.HasPreviousPage = true
		nodes = nodes[:len(nodes)-1]
	}
	var nodeAt func(int) *Message
	if last != nil {
		n := len(nodes) - 1
		nodeAt = func(i int) *Message {
			return nodes[n-i]
		}
	} else {
		nodeAt = func(i int) *Message {
			return nodes[i]
		}
	}
	c.Edges = make([]*MessageEdge, len(nodes))
	for i := range nodes {
		node := nodeAt(i)
		c.Edges[i] = &MessageEdge{
			Node:   node,
			Cursor: pager.toCursor(node),
		}
	}
	if l := len(c.Edges); l > 0 {
		c.PageInfo.StartCursor = &c.Edges[0].Cursor
		c.PageInfo.EndCursor = &c.Edges[l-1].Cursor
	}
	if c.TotalCount == 0 {
		c.TotalCount = len(nodes)
	}
}

// MessagePaginateOption enables pagination customization.
type MessagePaginateOption func(*messagePager) error

// WithMessageOrder configures pagination ordering.
func WithMessageOrder(order *MessageOrder) MessagePaginateOption {
	if order == nil {
		order = DefaultMessageOrder
	}
	o := *order
	return func(pager *messagePager) error {
		if err := o.Direction.Validate(); err != nil {
			return err
		}
		if o.Field == nil {
			o.Field = DefaultMessageOrder.Field
		}
		pager.order = &o
		return nil
	}
}

// WithMessageFilter configures pagination filter.
func WithMessageFilter(filter func(*MessageQuery) (*MessageQuery, error)) MessagePaginateOption {
	return func(pager *messagePager) error {
		if filter == nil {
			return errors.New("MessageQuery filter cannot be nil")
		}
		pager.filter = filter
		return nil
	}
}

type messagePager struct {
	order  *MessageOrder
	filter func(*MessageQuery) (*MessageQuery, error)
}

func newMessagePager(opts []MessagePaginateOption) (*messagePager, error) {
	pager := &messagePager{}
	for _, opt := range opts {
		if err := opt(pager); err != nil {
			return nil, err
		}
	}
	if pager.order == nil {
		pager.order = DefaultMessageOrder
	}
	return pager, nil
}

func (p *messagePager) applyFilter(query *MessageQuery) (*MessageQuery, error) {
	if p.filter != nil {
		return p.filter(query)
	}
	return query, nil
}

func (p *messagePager) toCursor(m *Message) Cursor {
	return p.order.Field.toCursor(m)
}

func (p *messagePager) applyCursors(query *MessageQuery, after, before *Cursor) *MessageQuery {
	for _, predicate := range cursorsToPredicates(
		p.order.Direction, after, before,
		p.order.Field.field, DefaultMessageOrder.Field.field,
	) {
		query = query.Where(predicate)
	}
	return query
}

func (p *messagePager) applyOrder(query *MessageQuery, reverse bool) *MessageQuery {
	direction := p.order.Direction
	if reverse {
		direction = direction.reverse()
	}
	query = query.Order(direction.orderFunc(p.order.Field.field))
	if p.order.Field != DefaultMessageOrder.Field {
		query = query.Order(direction.orderFunc(DefaultMessageOrder.Field.field))
	}
	return query
}

func (p *messagePager) orderExpr(reverse bool) sql.Querier {
	direction := p.order.Direction
	if reverse {
		direction = direction.reverse()
	}
	return sql.ExprFunc(func(b *sql.Builder) {
		b.Ident(p.order.Field.field).Pad().WriteString(string(direction))
		if p.order.Field != DefaultMessageOrder.Field {
			b.Comma().Ident(DefaultMessageOrder.Field.field).Pad().WriteString(string(direction))
		}
	})
}

// Paginate executes the query and returns a relay based cursor connection to Message.
func (m *MessageQuery) Paginate(
	ctx context.Context, after *Cursor, first *int,
	before *Cursor, last *int, opts ...MessagePaginateOption,
) (*MessageConnection, error) {
	if err := validateFirstLast(first, last); err != nil {
		return nil, err
	}
	pager, err := newMessagePager(opts)
	if err != nil {
		return nil, err
	}
	if m, err = pager.applyFilter(m); err != nil {
		return nil, err
	}
	conn := &MessageConnection{Edges: []*MessageEdge{}}
	ignoredEdges := !hasCollectedField(ctx, edgesField)
	if hasCollectedField(ctx, totalCountField) || hasCollectedField(ctx, pageInfoField) {
		hasPagination := after != nil || first != nil || before != nil || last != nil
		if hasPagination || ignoredEdges {
			if conn.TotalCount, err = m.Count(ctx); err != nil {
				return nil, err
			}
			conn.PageInfo.HasNextPage = first != nil && conn.TotalCount > 0
			conn.PageInfo.HasPreviousPage = last != nil && conn.TotalCount > 0
		}
	}
	if ignoredEdges || (first != nil && *first == 0) || (last != nil && *last == 0) {
		return conn, nil
	}

	m = pager.applyCursors(m, after, before)
	m = pager.applyOrder(m, last != nil)
	if limit := paginateLimit(first, last); limit != 0 {
		m.Limit(limit)
	}
	if field := collectedField(ctx, edgesField, nodeField); field != nil {
		if err := m.collectField(ctx, graphql.GetOperationContext(ctx), *field, []string{edgesField, nodeField}); err != nil {
			return nil, err
		}
	}

	nodes, err := m.All(ctx)
	if err != nil {
		return nil, err
	}
	conn.build(nodes, pager, after, first, before, last)
	return conn, nil
}

var (
	// MessageOrderFieldSentAt orders Message by sent_at.
	MessageOrderFieldSentAt = &MessageOrderField{
		field: message.FieldSentAt,
		toCursor: func(m *Message) Cursor {
			return Cursor{
				ID:    m.ID,
				Value: m.SentAt,
			}
		},
	}
	// MessageOrderFieldMessage orders Message by message.
	MessageOrderFieldMessage = &MessageOrderField{
		field: message.FieldMessage,
		toCursor: func(m *Message) Cursor {
			return Cursor{
				ID:    m.ID,
				Value: m.Message,
			}
		},
	}
)

// String implement fmt.Stringer interface.
func (f MessageOrderField) String() string {
	var str string
	switch f.field {
	case message.FieldSentAt:
		str = "SENT_AT"
	case message.FieldMessage:
		str = "MESSAGE"
	}
	return str
}

// MarshalGQL implements graphql.Marshaler interface.
func (f MessageOrderField) MarshalGQL(w io.Writer) {
	io.WriteString(w, strconv.Quote(f.String()))
}

// UnmarshalGQL implements graphql.Unmarshaler interface.
func (f *MessageOrderField) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("MessageOrderField %T must be a string", v)
	}
	switch str {
	case "SENT_AT":
		*f = *MessageOrderFieldSentAt
	case "MESSAGE":
		*f = *MessageOrderFieldMessage
	default:
		return fmt.Errorf("%s is not a valid MessageOrderField", str)
	}
	return nil
}

// MessageOrderField defines the ordering field of Message.
type MessageOrderField struct {
	field    string
	toCursor func(*Message) Cursor
}

// MessageOrder defines the ordering of Message.
type MessageOrder struct {
	Direction OrderDirection     `json:"direction"`
	Field     *MessageOrderField `json:"field"`
}

// DefaultMessageOrder is the default ordering of Message.
var DefaultMessageOrder = &MessageOrder{
	Direction: OrderDirectionAsc,
	Field: &MessageOrderField{
		field: message.FieldID,
		toCursor: func(m *Message) Cursor {
			return Cursor{ID: m.ID}
		},
	},
}

// ToEdge converts Message into MessageEdge.
func (m *Message) ToEdge(order *MessageOrder) *MessageEdge {
	if order == nil {
		order = DefaultMessageOrder
	}
	return &MessageEdge{
		Node:   m,
		Cursor: order.Field.toCursor(m),
	}
}

// UserEdge is the edge representation of User.
type UserEdge struct {
	Node   *User  `json:"node"`
	Cursor Cursor `json:"cursor"`
}

// UserConnection is the connection containing edges to User.
type UserConnection struct {
	Edges      []*UserEdge `json:"edges"`
	PageInfo   PageInfo    `json:"pageInfo"`
	TotalCount int         `json:"totalCount"`
}

func (c *UserConnection) build(nodes []*User, pager *userPager, after *Cursor, first *int, before *Cursor, last *int) {
	c.PageInfo.HasNextPage = before != nil
	c.PageInfo.HasPreviousPage = after != nil
	if first != nil && *first+1 == len(nodes) {
		c.PageInfo.HasNextPage = true
		nodes = nodes[:len(nodes)-1]
	} else if last != nil && *last+1 == len(nodes) {
		c.PageInfo.HasPreviousPage = true
		nodes = nodes[:len(nodes)-1]
	}
	var nodeAt func(int) *User
	if last != nil {
		n := len(nodes) - 1
		nodeAt = func(i int) *User {
			return nodes[n-i]
		}
	} else {
		nodeAt = func(i int) *User {
			return nodes[i]
		}
	}
	c.Edges = make([]*UserEdge, len(nodes))
	for i := range nodes {
		node := nodeAt(i)
		c.Edges[i] = &UserEdge{
			Node:   node,
			Cursor: pager.toCursor(node),
		}
	}
	if l := len(c.Edges); l > 0 {
		c.PageInfo.StartCursor = &c.Edges[0].Cursor
		c.PageInfo.EndCursor = &c.Edges[l-1].Cursor
	}
	if c.TotalCount == 0 {
		c.TotalCount = len(nodes)
	}
}

// UserPaginateOption enables pagination customization.
type UserPaginateOption func(*userPager) error

// WithUserOrder configures pagination ordering.
func WithUserOrder(order *UserOrder) UserPaginateOption {
	if order == nil {
		order = DefaultUserOrder
	}
	o := *order
	return func(pager *userPager) error {
		if err := o.Direction.Validate(); err != nil {
			return err
		}
		if o.Field == nil {
			o.Field = DefaultUserOrder.Field
		}
		pager.order = &o
		return nil
	}
}

// WithUserFilter configures pagination filter.
func WithUserFilter(filter func(*UserQuery) (*UserQuery, error)) UserPaginateOption {
	return func(pager *userPager) error {
		if filter == nil {
			return errors.New("UserQuery filter cannot be nil")
		}
		pager.filter = filter
		return nil
	}
}

type userPager struct {
	order  *UserOrder
	filter func(*UserQuery) (*UserQuery, error)
}

func newUserPager(opts []UserPaginateOption) (*userPager, error) {
	pager := &userPager{}
	for _, opt := range opts {
		if err := opt(pager); err != nil {
			return nil, err
		}
	}
	if pager.order == nil {
		pager.order = DefaultUserOrder
	}
	return pager, nil
}

func (p *userPager) applyFilter(query *UserQuery) (*UserQuery, error) {
	if p.filter != nil {
		return p.filter(query)
	}
	return query, nil
}

func (p *userPager) toCursor(u *User) Cursor {
	return p.order.Field.toCursor(u)
}

func (p *userPager) applyCursors(query *UserQuery, after, before *Cursor) *UserQuery {
	for _, predicate := range cursorsToPredicates(
		p.order.Direction, after, before,
		p.order.Field.field, DefaultUserOrder.Field.field,
	) {
		query = query.Where(predicate)
	}
	return query
}

func (p *userPager) applyOrder(query *UserQuery, reverse bool) *UserQuery {
	direction := p.order.Direction
	if reverse {
		direction = direction.reverse()
	}
	query = query.Order(direction.orderFunc(p.order.Field.field))
	if p.order.Field != DefaultUserOrder.Field {
		query = query.Order(direction.orderFunc(DefaultUserOrder.Field.field))
	}
	return query
}

func (p *userPager) orderExpr(reverse bool) sql.Querier {
	direction := p.order.Direction
	if reverse {
		direction = direction.reverse()
	}
	return sql.ExprFunc(func(b *sql.Builder) {
		b.Ident(p.order.Field.field).Pad().WriteString(string(direction))
		if p.order.Field != DefaultUserOrder.Field {
			b.Comma().Ident(DefaultUserOrder.Field.field).Pad().WriteString(string(direction))
		}
	})
}

// Paginate executes the query and returns a relay based cursor connection to User.
func (u *UserQuery) Paginate(
	ctx context.Context, after *Cursor, first *int,
	before *Cursor, last *int, opts ...UserPaginateOption,
) (*UserConnection, error) {
	if err := validateFirstLast(first, last); err != nil {
		return nil, err
	}
	pager, err := newUserPager(opts)
	if err != nil {
		return nil, err
	}
	if u, err = pager.applyFilter(u); err != nil {
		return nil, err
	}
	conn := &UserConnection{Edges: []*UserEdge{}}
	ignoredEdges := !hasCollectedField(ctx, edgesField)
	if hasCollectedField(ctx, totalCountField) || hasCollectedField(ctx, pageInfoField) {
		hasPagination := after != nil || first != nil || before != nil || last != nil
		if hasPagination || ignoredEdges {
			if conn.TotalCount, err = u.Count(ctx); err != nil {
				return nil, err
			}
			conn.PageInfo.HasNextPage = first != nil && conn.TotalCount > 0
			conn.PageInfo.HasPreviousPage = last != nil && conn.TotalCount > 0
		}
	}
	if ignoredEdges || (first != nil && *first == 0) || (last != nil && *last == 0) {
		return conn, nil
	}

	u = pager.applyCursors(u, after, before)
	u = pager.applyOrder(u, last != nil)
	if limit := paginateLimit(first, last); limit != 0 {
		u.Limit(limit)
	}
	if field := collectedField(ctx, edgesField, nodeField); field != nil {
		if err := u.collectField(ctx, graphql.GetOperationContext(ctx), *field, []string{edgesField, nodeField}); err != nil {
			return nil, err
		}
	}

	nodes, err := u.All(ctx)
	if err != nil {
		return nil, err
	}
	conn.build(nodes, pager, after, first, before, last)
	return conn, nil
}

var (
	// UserOrderFieldCreatedAt orders User by created_at.
	UserOrderFieldCreatedAt = &UserOrderField{
		field: user.FieldCreatedAt,
		toCursor: func(u *User) Cursor {
			return Cursor{
				ID:    u.ID,
				Value: u.CreatedAt,
			}
		},
	}
	// UserOrderFieldUpdatedAt orders User by updated_at.
	UserOrderFieldUpdatedAt = &UserOrderField{
		field: user.FieldUpdatedAt,
		toCursor: func(u *User) Cursor {
			return Cursor{
				ID:    u.ID,
				Value: u.UpdatedAt,
			}
		},
	}
	// UserOrderFieldUsername orders User by username.
	UserOrderFieldUsername = &UserOrderField{
		field: user.FieldUsername,
		toCursor: func(u *User) Cursor {
			return Cursor{
				ID:    u.ID,
				Value: u.Username,
			}
		},
	}
)

// String implement fmt.Stringer interface.
func (f UserOrderField) String() string {
	var str string
	switch f.field {
	case user.FieldCreatedAt:
		str = "CREATED_AT"
	case user.FieldUpdatedAt:
		str = "UPDATED_AT"
	case user.FieldUsername:
		str = "USERNAME"
	}
	return str
}

// MarshalGQL implements graphql.Marshaler interface.
func (f UserOrderField) MarshalGQL(w io.Writer) {
	io.WriteString(w, strconv.Quote(f.String()))
}

// UnmarshalGQL implements graphql.Unmarshaler interface.
func (f *UserOrderField) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("UserOrderField %T must be a string", v)
	}
	switch str {
	case "CREATED_AT":
		*f = *UserOrderFieldCreatedAt
	case "UPDATED_AT":
		*f = *UserOrderFieldUpdatedAt
	case "USERNAME":
		*f = *UserOrderFieldUsername
	default:
		return fmt.Errorf("%s is not a valid UserOrderField", str)
	}
	return nil
}

// UserOrderField defines the ordering field of User.
type UserOrderField struct {
	field    string
	toCursor func(*User) Cursor
}

// UserOrder defines the ordering of User.
type UserOrder struct {
	Direction OrderDirection  `json:"direction"`
	Field     *UserOrderField `json:"field"`
}

// DefaultUserOrder is the default ordering of User.
var DefaultUserOrder = &UserOrder{
	Direction: OrderDirectionAsc,
	Field: &UserOrderField{
		field: user.FieldID,
		toCursor: func(u *User) Cursor {
			return Cursor{ID: u.ID}
		},
	},
}

// ToEdge converts User into UserEdge.
func (u *User) ToEdge(order *UserOrder) *UserEdge {
	if order == nil {
		order = DefaultUserOrder
	}
	return &UserEdge{
		Node:   u,
		Cursor: order.Field.toCursor(u),
	}
}

// UserGroupEdge is the edge representation of UserGroup.
type UserGroupEdge struct {
	Node   *UserGroup `json:"node"`
	Cursor Cursor     `json:"cursor"`
}

// UserGroupConnection is the connection containing edges to UserGroup.
type UserGroupConnection struct {
	Edges      []*UserGroupEdge `json:"edges"`
	PageInfo   PageInfo         `json:"pageInfo"`
	TotalCount int              `json:"totalCount"`
}

func (c *UserGroupConnection) build(nodes []*UserGroup, pager *usergroupPager, after *Cursor, first *int, before *Cursor, last *int) {
	c.PageInfo.HasNextPage = before != nil
	c.PageInfo.HasPreviousPage = after != nil
	if first != nil && *first+1 == len(nodes) {
		c.PageInfo.HasNextPage = true
		nodes = nodes[:len(nodes)-1]
	} else if last != nil && *last+1 == len(nodes) {
		c.PageInfo.HasPreviousPage = true
		nodes = nodes[:len(nodes)-1]
	}
	var nodeAt func(int) *UserGroup
	if last != nil {
		n := len(nodes) - 1
		nodeAt = func(i int) *UserGroup {
			return nodes[n-i]
		}
	} else {
		nodeAt = func(i int) *UserGroup {
			return nodes[i]
		}
	}
	c.Edges = make([]*UserGroupEdge, len(nodes))
	for i := range nodes {
		node := nodeAt(i)
		c.Edges[i] = &UserGroupEdge{
			Node:   node,
			Cursor: pager.toCursor(node),
		}
	}
	if l := len(c.Edges); l > 0 {
		c.PageInfo.StartCursor = &c.Edges[0].Cursor
		c.PageInfo.EndCursor = &c.Edges[l-1].Cursor
	}
	if c.TotalCount == 0 {
		c.TotalCount = len(nodes)
	}
}

// UserGroupPaginateOption enables pagination customization.
type UserGroupPaginateOption func(*usergroupPager) error

// WithUserGroupOrder configures pagination ordering.
func WithUserGroupOrder(order *UserGroupOrder) UserGroupPaginateOption {
	if order == nil {
		order = DefaultUserGroupOrder
	}
	o := *order
	return func(pager *usergroupPager) error {
		if err := o.Direction.Validate(); err != nil {
			return err
		}
		if o.Field == nil {
			o.Field = DefaultUserGroupOrder.Field
		}
		pager.order = &o
		return nil
	}
}

// WithUserGroupFilter configures pagination filter.
func WithUserGroupFilter(filter func(*UserGroupQuery) (*UserGroupQuery, error)) UserGroupPaginateOption {
	return func(pager *usergroupPager) error {
		if filter == nil {
			return errors.New("UserGroupQuery filter cannot be nil")
		}
		pager.filter = filter
		return nil
	}
}

type usergroupPager struct {
	order  *UserGroupOrder
	filter func(*UserGroupQuery) (*UserGroupQuery, error)
}

func newUserGroupPager(opts []UserGroupPaginateOption) (*usergroupPager, error) {
	pager := &usergroupPager{}
	for _, opt := range opts {
		if err := opt(pager); err != nil {
			return nil, err
		}
	}
	if pager.order == nil {
		pager.order = DefaultUserGroupOrder
	}
	return pager, nil
}

func (p *usergroupPager) applyFilter(query *UserGroupQuery) (*UserGroupQuery, error) {
	if p.filter != nil {
		return p.filter(query)
	}
	return query, nil
}

func (p *usergroupPager) toCursor(ug *UserGroup) Cursor {
	return p.order.Field.toCursor(ug)
}

func (p *usergroupPager) applyCursors(query *UserGroupQuery, after, before *Cursor) *UserGroupQuery {
	for _, predicate := range cursorsToPredicates(
		p.order.Direction, after, before,
		p.order.Field.field, DefaultUserGroupOrder.Field.field,
	) {
		query = query.Where(predicate)
	}
	return query
}

func (p *usergroupPager) applyOrder(query *UserGroupQuery, reverse bool) *UserGroupQuery {
	direction := p.order.Direction
	if reverse {
		direction = direction.reverse()
	}
	query = query.Order(direction.orderFunc(p.order.Field.field))
	if p.order.Field != DefaultUserGroupOrder.Field {
		query = query.Order(direction.orderFunc(DefaultUserGroupOrder.Field.field))
	}
	return query
}

func (p *usergroupPager) orderExpr(reverse bool) sql.Querier {
	direction := p.order.Direction
	if reverse {
		direction = direction.reverse()
	}
	return sql.ExprFunc(func(b *sql.Builder) {
		b.Ident(p.order.Field.field).Pad().WriteString(string(direction))
		if p.order.Field != DefaultUserGroupOrder.Field {
			b.Comma().Ident(DefaultUserGroupOrder.Field.field).Pad().WriteString(string(direction))
		}
	})
}

// Paginate executes the query and returns a relay based cursor connection to UserGroup.
func (ug *UserGroupQuery) Paginate(
	ctx context.Context, after *Cursor, first *int,
	before *Cursor, last *int, opts ...UserGroupPaginateOption,
) (*UserGroupConnection, error) {
	if err := validateFirstLast(first, last); err != nil {
		return nil, err
	}
	pager, err := newUserGroupPager(opts)
	if err != nil {
		return nil, err
	}
	if ug, err = pager.applyFilter(ug); err != nil {
		return nil, err
	}
	conn := &UserGroupConnection{Edges: []*UserGroupEdge{}}
	ignoredEdges := !hasCollectedField(ctx, edgesField)
	if hasCollectedField(ctx, totalCountField) || hasCollectedField(ctx, pageInfoField) {
		hasPagination := after != nil || first != nil || before != nil || last != nil
		if hasPagination || ignoredEdges {
			if conn.TotalCount, err = ug.Count(ctx); err != nil {
				return nil, err
			}
			conn.PageInfo.HasNextPage = first != nil && conn.TotalCount > 0
			conn.PageInfo.HasPreviousPage = last != nil && conn.TotalCount > 0
		}
	}
	if ignoredEdges || (first != nil && *first == 0) || (last != nil && *last == 0) {
		return conn, nil
	}

	ug = pager.applyCursors(ug, after, before)
	ug = pager.applyOrder(ug, last != nil)
	if limit := paginateLimit(first, last); limit != 0 {
		ug.Limit(limit)
	}
	if field := collectedField(ctx, edgesField, nodeField); field != nil {
		if err := ug.collectField(ctx, graphql.GetOperationContext(ctx), *field, []string{edgesField, nodeField}); err != nil {
			return nil, err
		}
	}

	nodes, err := ug.All(ctx)
	if err != nil {
		return nil, err
	}
	conn.build(nodes, pager, after, first, before, last)
	return conn, nil
}

var (
	// UserGroupOrderFieldCreatedAt orders UserGroup by created_at.
	UserGroupOrderFieldCreatedAt = &UserGroupOrderField{
		field: usergroup.FieldCreatedAt,
		toCursor: func(ug *UserGroup) Cursor {
			return Cursor{
				ID:    ug.ID,
				Value: ug.CreatedAt,
			}
		},
	}
)

// String implement fmt.Stringer interface.
func (f UserGroupOrderField) String() string {
	var str string
	switch f.field {
	case usergroup.FieldCreatedAt:
		str = "CREATED_AT"
	}
	return str
}

// MarshalGQL implements graphql.Marshaler interface.
func (f UserGroupOrderField) MarshalGQL(w io.Writer) {
	io.WriteString(w, strconv.Quote(f.String()))
}

// UnmarshalGQL implements graphql.Unmarshaler interface.
func (f *UserGroupOrderField) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("UserGroupOrderField %T must be a string", v)
	}
	switch str {
	case "CREATED_AT":
		*f = *UserGroupOrderFieldCreatedAt
	default:
		return fmt.Errorf("%s is not a valid UserGroupOrderField", str)
	}
	return nil
}

// UserGroupOrderField defines the ordering field of UserGroup.
type UserGroupOrderField struct {
	field    string
	toCursor func(*UserGroup) Cursor
}

// UserGroupOrder defines the ordering of UserGroup.
type UserGroupOrder struct {
	Direction OrderDirection       `json:"direction"`
	Field     *UserGroupOrderField `json:"field"`
}

// DefaultUserGroupOrder is the default ordering of UserGroup.
var DefaultUserGroupOrder = &UserGroupOrder{
	Direction: OrderDirectionAsc,
	Field: &UserGroupOrderField{
		field: usergroup.FieldID,
		toCursor: func(ug *UserGroup) Cursor {
			return Cursor{ID: ug.ID}
		},
	},
}

// ToEdge converts UserGroup into UserGroupEdge.
func (ug *UserGroup) ToEdge(order *UserGroupOrder) *UserGroupEdge {
	if order == nil {
		order = DefaultUserGroupOrder
	}
	return &UserGroupEdge{
		Node:   ug,
		Cursor: order.Field.toCursor(ug),
	}
}