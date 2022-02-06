// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"math"
	"oauth-az/ent/authrequest"
	"oauth-az/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// AuthRequestQuery is the builder for querying AuthRequest entities.
type AuthRequestQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.AuthRequest
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the AuthRequestQuery builder.
func (arq *AuthRequestQuery) Where(ps ...predicate.AuthRequest) *AuthRequestQuery {
	arq.predicates = append(arq.predicates, ps...)
	return arq
}

// Limit adds a limit step to the query.
func (arq *AuthRequestQuery) Limit(limit int) *AuthRequestQuery {
	arq.limit = &limit
	return arq
}

// Offset adds an offset step to the query.
func (arq *AuthRequestQuery) Offset(offset int) *AuthRequestQuery {
	arq.offset = &offset
	return arq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (arq *AuthRequestQuery) Unique(unique bool) *AuthRequestQuery {
	arq.unique = &unique
	return arq
}

// Order adds an order step to the query.
func (arq *AuthRequestQuery) Order(o ...OrderFunc) *AuthRequestQuery {
	arq.order = append(arq.order, o...)
	return arq
}

// First returns the first AuthRequest entity from the query.
// Returns a *NotFoundError when no AuthRequest was found.
func (arq *AuthRequestQuery) First(ctx context.Context) (*AuthRequest, error) {
	nodes, err := arq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{authrequest.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (arq *AuthRequestQuery) FirstX(ctx context.Context) *AuthRequest {
	node, err := arq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first AuthRequest ID from the query.
// Returns a *NotFoundError when no AuthRequest ID was found.
func (arq *AuthRequestQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = arq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{authrequest.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (arq *AuthRequestQuery) FirstIDX(ctx context.Context) int {
	id, err := arq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single AuthRequest entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when exactly one AuthRequest entity is not found.
// Returns a *NotFoundError when no AuthRequest entities are found.
func (arq *AuthRequestQuery) Only(ctx context.Context) (*AuthRequest, error) {
	nodes, err := arq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{authrequest.Label}
	default:
		return nil, &NotSingularError{authrequest.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (arq *AuthRequestQuery) OnlyX(ctx context.Context) *AuthRequest {
	node, err := arq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only AuthRequest ID in the query.
// Returns a *NotSingularError when exactly one AuthRequest ID is not found.
// Returns a *NotFoundError when no entities are found.
func (arq *AuthRequestQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = arq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{authrequest.Label}
	default:
		err = &NotSingularError{authrequest.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (arq *AuthRequestQuery) OnlyIDX(ctx context.Context) int {
	id, err := arq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of AuthRequests.
func (arq *AuthRequestQuery) All(ctx context.Context) ([]*AuthRequest, error) {
	if err := arq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return arq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (arq *AuthRequestQuery) AllX(ctx context.Context) []*AuthRequest {
	nodes, err := arq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of AuthRequest IDs.
func (arq *AuthRequestQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := arq.Select(authrequest.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (arq *AuthRequestQuery) IDsX(ctx context.Context) []int {
	ids, err := arq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (arq *AuthRequestQuery) Count(ctx context.Context) (int, error) {
	if err := arq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return arq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (arq *AuthRequestQuery) CountX(ctx context.Context) int {
	count, err := arq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (arq *AuthRequestQuery) Exist(ctx context.Context) (bool, error) {
	if err := arq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return arq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (arq *AuthRequestQuery) ExistX(ctx context.Context) bool {
	exist, err := arq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the AuthRequestQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (arq *AuthRequestQuery) Clone() *AuthRequestQuery {
	if arq == nil {
		return nil
	}
	return &AuthRequestQuery{
		config:     arq.config,
		limit:      arq.limit,
		offset:     arq.offset,
		order:      append([]OrderFunc{}, arq.order...),
		predicates: append([]predicate.AuthRequest{}, arq.predicates...),
		// clone intermediate query.
		sql:  arq.sql.Clone(),
		path: arq.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"created_at,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.AuthRequest.Query().
//		GroupBy(authrequest.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (arq *AuthRequestQuery) GroupBy(field string, fields ...string) *AuthRequestGroupBy {
	group := &AuthRequestGroupBy{config: arq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := arq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return arq.sqlQuery(ctx), nil
	}
	return group
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"created_at,omitempty"`
//	}
//
//	client.AuthRequest.Query().
//		Select(authrequest.FieldCreatedAt).
//		Scan(ctx, &v)
//
func (arq *AuthRequestQuery) Select(fields ...string) *AuthRequestSelect {
	arq.fields = append(arq.fields, fields...)
	return &AuthRequestSelect{AuthRequestQuery: arq}
}

func (arq *AuthRequestQuery) prepareQuery(ctx context.Context) error {
	for _, f := range arq.fields {
		if !authrequest.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if arq.path != nil {
		prev, err := arq.path(ctx)
		if err != nil {
			return err
		}
		arq.sql = prev
	}
	return nil
}

func (arq *AuthRequestQuery) sqlAll(ctx context.Context) ([]*AuthRequest, error) {
	var (
		nodes = []*AuthRequest{}
		_spec = arq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &AuthRequest{config: arq.config}
		nodes = append(nodes, node)
		return node.scanValues(columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		if len(nodes) == 0 {
			return fmt.Errorf("ent: Assign called without calling ScanValues")
		}
		node := nodes[len(nodes)-1]
		return node.assignValues(columns, values)
	}
	if err := sqlgraph.QueryNodes(ctx, arq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (arq *AuthRequestQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := arq.querySpec()
	return sqlgraph.CountNodes(ctx, arq.driver, _spec)
}

func (arq *AuthRequestQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := arq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (arq *AuthRequestQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   authrequest.Table,
			Columns: authrequest.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: authrequest.FieldID,
			},
		},
		From:   arq.sql,
		Unique: true,
	}
	if unique := arq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := arq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, authrequest.FieldID)
		for i := range fields {
			if fields[i] != authrequest.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := arq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := arq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := arq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := arq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (arq *AuthRequestQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(arq.driver.Dialect())
	t1 := builder.Table(authrequest.Table)
	columns := arq.fields
	if len(columns) == 0 {
		columns = authrequest.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if arq.sql != nil {
		selector = arq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	for _, p := range arq.predicates {
		p(selector)
	}
	for _, p := range arq.order {
		p(selector)
	}
	if offset := arq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := arq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// AuthRequestGroupBy is the group-by builder for AuthRequest entities.
type AuthRequestGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (argb *AuthRequestGroupBy) Aggregate(fns ...AggregateFunc) *AuthRequestGroupBy {
	argb.fns = append(argb.fns, fns...)
	return argb
}

// Scan applies the group-by query and scans the result into the given value.
func (argb *AuthRequestGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := argb.path(ctx)
	if err != nil {
		return err
	}
	argb.sql = query
	return argb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (argb *AuthRequestGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := argb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (argb *AuthRequestGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(argb.fields) > 1 {
		return nil, errors.New("ent: AuthRequestGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := argb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (argb *AuthRequestGroupBy) StringsX(ctx context.Context) []string {
	v, err := argb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (argb *AuthRequestGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = argb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{authrequest.Label}
	default:
		err = fmt.Errorf("ent: AuthRequestGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (argb *AuthRequestGroupBy) StringX(ctx context.Context) string {
	v, err := argb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (argb *AuthRequestGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(argb.fields) > 1 {
		return nil, errors.New("ent: AuthRequestGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := argb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (argb *AuthRequestGroupBy) IntsX(ctx context.Context) []int {
	v, err := argb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (argb *AuthRequestGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = argb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{authrequest.Label}
	default:
		err = fmt.Errorf("ent: AuthRequestGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (argb *AuthRequestGroupBy) IntX(ctx context.Context) int {
	v, err := argb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (argb *AuthRequestGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(argb.fields) > 1 {
		return nil, errors.New("ent: AuthRequestGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := argb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (argb *AuthRequestGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := argb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (argb *AuthRequestGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = argb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{authrequest.Label}
	default:
		err = fmt.Errorf("ent: AuthRequestGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (argb *AuthRequestGroupBy) Float64X(ctx context.Context) float64 {
	v, err := argb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (argb *AuthRequestGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(argb.fields) > 1 {
		return nil, errors.New("ent: AuthRequestGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := argb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (argb *AuthRequestGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := argb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (argb *AuthRequestGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = argb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{authrequest.Label}
	default:
		err = fmt.Errorf("ent: AuthRequestGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (argb *AuthRequestGroupBy) BoolX(ctx context.Context) bool {
	v, err := argb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (argb *AuthRequestGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range argb.fields {
		if !authrequest.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := argb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := argb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (argb *AuthRequestGroupBy) sqlQuery() *sql.Selector {
	selector := argb.sql.Select()
	aggregation := make([]string, 0, len(argb.fns))
	for _, fn := range argb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(argb.fields)+len(argb.fns))
		for _, f := range argb.fields {
			columns = append(columns, selector.C(f))
		}
		for _, c := range aggregation {
			columns = append(columns, c)
		}
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(argb.fields...)...)
}

// AuthRequestSelect is the builder for selecting fields of AuthRequest entities.
type AuthRequestSelect struct {
	*AuthRequestQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (ars *AuthRequestSelect) Scan(ctx context.Context, v interface{}) error {
	if err := ars.prepareQuery(ctx); err != nil {
		return err
	}
	ars.sql = ars.AuthRequestQuery.sqlQuery(ctx)
	return ars.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (ars *AuthRequestSelect) ScanX(ctx context.Context, v interface{}) {
	if err := ars.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (ars *AuthRequestSelect) Strings(ctx context.Context) ([]string, error) {
	if len(ars.fields) > 1 {
		return nil, errors.New("ent: AuthRequestSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := ars.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (ars *AuthRequestSelect) StringsX(ctx context.Context) []string {
	v, err := ars.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (ars *AuthRequestSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = ars.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{authrequest.Label}
	default:
		err = fmt.Errorf("ent: AuthRequestSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (ars *AuthRequestSelect) StringX(ctx context.Context) string {
	v, err := ars.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (ars *AuthRequestSelect) Ints(ctx context.Context) ([]int, error) {
	if len(ars.fields) > 1 {
		return nil, errors.New("ent: AuthRequestSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := ars.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (ars *AuthRequestSelect) IntsX(ctx context.Context) []int {
	v, err := ars.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (ars *AuthRequestSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = ars.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{authrequest.Label}
	default:
		err = fmt.Errorf("ent: AuthRequestSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (ars *AuthRequestSelect) IntX(ctx context.Context) int {
	v, err := ars.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (ars *AuthRequestSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(ars.fields) > 1 {
		return nil, errors.New("ent: AuthRequestSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := ars.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (ars *AuthRequestSelect) Float64sX(ctx context.Context) []float64 {
	v, err := ars.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (ars *AuthRequestSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = ars.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{authrequest.Label}
	default:
		err = fmt.Errorf("ent: AuthRequestSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (ars *AuthRequestSelect) Float64X(ctx context.Context) float64 {
	v, err := ars.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (ars *AuthRequestSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(ars.fields) > 1 {
		return nil, errors.New("ent: AuthRequestSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := ars.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (ars *AuthRequestSelect) BoolsX(ctx context.Context) []bool {
	v, err := ars.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (ars *AuthRequestSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = ars.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{authrequest.Label}
	default:
		err = fmt.Errorf("ent: AuthRequestSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (ars *AuthRequestSelect) BoolX(ctx context.Context) bool {
	v, err := ars.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (ars *AuthRequestSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := ars.sql.Query()
	if err := ars.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}