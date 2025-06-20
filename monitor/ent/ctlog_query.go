// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/anonymous-author-00000/ra-webs/monitor/ent/ctlog"
	"github.com/anonymous-author-00000/ra-webs/monitor/ent/predicate"
	"github.com/anonymous-author-00000/ra-webs/monitor/ent/ta"
)

// CTLogQuery is the builder for querying CTLog entities.
type CTLogQuery struct {
	config
	ctx        *QueryContext
	order      []ctlog.OrderOption
	inters     []Interceptor
	predicates []predicate.CTLog
	withTa     *TAQuery
	withFKs    bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the CTLogQuery builder.
func (clq *CTLogQuery) Where(ps ...predicate.CTLog) *CTLogQuery {
	clq.predicates = append(clq.predicates, ps...)
	return clq
}

// Limit the number of records to be returned by this query.
func (clq *CTLogQuery) Limit(limit int) *CTLogQuery {
	clq.ctx.Limit = &limit
	return clq
}

// Offset to start from.
func (clq *CTLogQuery) Offset(offset int) *CTLogQuery {
	clq.ctx.Offset = &offset
	return clq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (clq *CTLogQuery) Unique(unique bool) *CTLogQuery {
	clq.ctx.Unique = &unique
	return clq
}

// Order specifies how the records should be ordered.
func (clq *CTLogQuery) Order(o ...ctlog.OrderOption) *CTLogQuery {
	clq.order = append(clq.order, o...)
	return clq
}

// QueryTa chains the current query on the "ta" edge.
func (clq *CTLogQuery) QueryTa() *TAQuery {
	query := (&TAClient{config: clq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := clq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := clq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(ctlog.Table, ctlog.FieldID, selector),
			sqlgraph.To(ta.Table, ta.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, ctlog.TaTable, ctlog.TaColumn),
		)
		fromU = sqlgraph.SetNeighbors(clq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first CTLog entity from the query.
// Returns a *NotFoundError when no CTLog was found.
func (clq *CTLogQuery) First(ctx context.Context) (*CTLog, error) {
	nodes, err := clq.Limit(1).All(setContextOp(ctx, clq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{ctlog.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (clq *CTLogQuery) FirstX(ctx context.Context) *CTLog {
	node, err := clq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first CTLog ID from the query.
// Returns a *NotFoundError when no CTLog ID was found.
func (clq *CTLogQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = clq.Limit(1).IDs(setContextOp(ctx, clq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{ctlog.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (clq *CTLogQuery) FirstIDX(ctx context.Context) int {
	id, err := clq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single CTLog entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one CTLog entity is found.
// Returns a *NotFoundError when no CTLog entities are found.
func (clq *CTLogQuery) Only(ctx context.Context) (*CTLog, error) {
	nodes, err := clq.Limit(2).All(setContextOp(ctx, clq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{ctlog.Label}
	default:
		return nil, &NotSingularError{ctlog.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (clq *CTLogQuery) OnlyX(ctx context.Context) *CTLog {
	node, err := clq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only CTLog ID in the query.
// Returns a *NotSingularError when more than one CTLog ID is found.
// Returns a *NotFoundError when no entities are found.
func (clq *CTLogQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = clq.Limit(2).IDs(setContextOp(ctx, clq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{ctlog.Label}
	default:
		err = &NotSingularError{ctlog.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (clq *CTLogQuery) OnlyIDX(ctx context.Context) int {
	id, err := clq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of CTLogs.
func (clq *CTLogQuery) All(ctx context.Context) ([]*CTLog, error) {
	ctx = setContextOp(ctx, clq.ctx, ent.OpQueryAll)
	if err := clq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*CTLog, *CTLogQuery]()
	return withInterceptors[[]*CTLog](ctx, clq, qr, clq.inters)
}

// AllX is like All, but panics if an error occurs.
func (clq *CTLogQuery) AllX(ctx context.Context) []*CTLog {
	nodes, err := clq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of CTLog IDs.
func (clq *CTLogQuery) IDs(ctx context.Context) (ids []int, err error) {
	if clq.ctx.Unique == nil && clq.path != nil {
		clq.Unique(true)
	}
	ctx = setContextOp(ctx, clq.ctx, ent.OpQueryIDs)
	if err = clq.Select(ctlog.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (clq *CTLogQuery) IDsX(ctx context.Context) []int {
	ids, err := clq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (clq *CTLogQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, clq.ctx, ent.OpQueryCount)
	if err := clq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, clq, querierCount[*CTLogQuery](), clq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (clq *CTLogQuery) CountX(ctx context.Context) int {
	count, err := clq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (clq *CTLogQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, clq.ctx, ent.OpQueryExist)
	switch _, err := clq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (clq *CTLogQuery) ExistX(ctx context.Context) bool {
	exist, err := clq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the CTLogQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (clq *CTLogQuery) Clone() *CTLogQuery {
	if clq == nil {
		return nil
	}
	return &CTLogQuery{
		config:     clq.config,
		ctx:        clq.ctx.Clone(),
		order:      append([]ctlog.OrderOption{}, clq.order...),
		inters:     append([]Interceptor{}, clq.inters...),
		predicates: append([]predicate.CTLog{}, clq.predicates...),
		withTa:     clq.withTa.Clone(),
		// clone intermediate query.
		sql:  clq.sql.Clone(),
		path: clq.path,
	}
}

// WithTa tells the query-builder to eager-load the nodes that are connected to
// the "ta" edge. The optional arguments are used to configure the query builder of the edge.
func (clq *CTLogQuery) WithTa(opts ...func(*TAQuery)) *CTLogQuery {
	query := (&TAClient{config: clq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	clq.withTa = query
	return clq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		MonitorLogID int `json:"monitor_log_id,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.CTLog.Query().
//		GroupBy(ctlog.FieldMonitorLogID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (clq *CTLogQuery) GroupBy(field string, fields ...string) *CTLogGroupBy {
	clq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &CTLogGroupBy{build: clq}
	grbuild.flds = &clq.ctx.Fields
	grbuild.label = ctlog.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		MonitorLogID int `json:"monitor_log_id,omitempty"`
//	}
//
//	client.CTLog.Query().
//		Select(ctlog.FieldMonitorLogID).
//		Scan(ctx, &v)
func (clq *CTLogQuery) Select(fields ...string) *CTLogSelect {
	clq.ctx.Fields = append(clq.ctx.Fields, fields...)
	sbuild := &CTLogSelect{CTLogQuery: clq}
	sbuild.label = ctlog.Label
	sbuild.flds, sbuild.scan = &clq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a CTLogSelect configured with the given aggregations.
func (clq *CTLogQuery) Aggregate(fns ...AggregateFunc) *CTLogSelect {
	return clq.Select().Aggregate(fns...)
}

func (clq *CTLogQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range clq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, clq); err != nil {
				return err
			}
		}
	}
	for _, f := range clq.ctx.Fields {
		if !ctlog.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if clq.path != nil {
		prev, err := clq.path(ctx)
		if err != nil {
			return err
		}
		clq.sql = prev
	}
	return nil
}

func (clq *CTLogQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*CTLog, error) {
	var (
		nodes       = []*CTLog{}
		withFKs     = clq.withFKs
		_spec       = clq.querySpec()
		loadedTypes = [1]bool{
			clq.withTa != nil,
		}
	)
	if clq.withTa != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, ctlog.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*CTLog).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &CTLog{config: clq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, clq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := clq.withTa; query != nil {
		if err := clq.loadTa(ctx, query, nodes, nil,
			func(n *CTLog, e *TA) { n.Edges.Ta = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (clq *CTLogQuery) loadTa(ctx context.Context, query *TAQuery, nodes []*CTLog, init func(*CTLog), assign func(*CTLog, *TA)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*CTLog)
	for i := range nodes {
		if nodes[i].ct_log_ta == nil {
			continue
		}
		fk := *nodes[i].ct_log_ta
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(ta.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "ct_log_ta" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (clq *CTLogQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := clq.querySpec()
	_spec.Node.Columns = clq.ctx.Fields
	if len(clq.ctx.Fields) > 0 {
		_spec.Unique = clq.ctx.Unique != nil && *clq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, clq.driver, _spec)
}

func (clq *CTLogQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(ctlog.Table, ctlog.Columns, sqlgraph.NewFieldSpec(ctlog.FieldID, field.TypeInt))
	_spec.From = clq.sql
	if unique := clq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if clq.path != nil {
		_spec.Unique = true
	}
	if fields := clq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, ctlog.FieldID)
		for i := range fields {
			if fields[i] != ctlog.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := clq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := clq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := clq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := clq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (clq *CTLogQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(clq.driver.Dialect())
	t1 := builder.Table(ctlog.Table)
	columns := clq.ctx.Fields
	if len(columns) == 0 {
		columns = ctlog.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if clq.sql != nil {
		selector = clq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if clq.ctx.Unique != nil && *clq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range clq.predicates {
		p(selector)
	}
	for _, p := range clq.order {
		p(selector)
	}
	if offset := clq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := clq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// CTLogGroupBy is the group-by builder for CTLog entities.
type CTLogGroupBy struct {
	selector
	build *CTLogQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (clgb *CTLogGroupBy) Aggregate(fns ...AggregateFunc) *CTLogGroupBy {
	clgb.fns = append(clgb.fns, fns...)
	return clgb
}

// Scan applies the selector query and scans the result into the given value.
func (clgb *CTLogGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, clgb.build.ctx, ent.OpQueryGroupBy)
	if err := clgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*CTLogQuery, *CTLogGroupBy](ctx, clgb.build, clgb, clgb.build.inters, v)
}

func (clgb *CTLogGroupBy) sqlScan(ctx context.Context, root *CTLogQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(clgb.fns))
	for _, fn := range clgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*clgb.flds)+len(clgb.fns))
		for _, f := range *clgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*clgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := clgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// CTLogSelect is the builder for selecting fields of CTLog entities.
type CTLogSelect struct {
	*CTLogQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (cls *CTLogSelect) Aggregate(fns ...AggregateFunc) *CTLogSelect {
	cls.fns = append(cls.fns, fns...)
	return cls
}

// Scan applies the selector query and scans the result into the given value.
func (cls *CTLogSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, cls.ctx, ent.OpQuerySelect)
	if err := cls.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*CTLogQuery, *CTLogSelect](ctx, cls.CTLogQuery, cls, cls.inters, v)
}

func (cls *CTLogSelect) sqlScan(ctx context.Context, root *CTLogQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(cls.fns))
	for _, fn := range cls.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*cls.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := cls.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
