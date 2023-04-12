package db

// Query is the datatype for a query object
type Query struct {
	filters MultiFilter
	limited bool
	limit   uint
	offset  uint
}

// NewQuery creates a Query and returns with default values
func NewQuery() Query {
	return Query{
		filters: MultiFilter{
			Filters: []Filter{},
			Op:      MultiFilterOpAND,
		},
		limited: false,
		limit:   0,
		offset:  0,
	}
}

// WithFilter includes a filter in the query before execution
func (q *Query) WithFilter(filter Filter) *Query {
	q.filters.Filters = append(q.filters.Filters, filter)
	return q
}

// WithLimit includes a limit to the query, input must be uint as limits cannot be negative values
func (q *Query) WithLimit(limit uint) *Query {
	q.limited = true
	q.limit = limit
	return q
}

// WithoutLimit removes the limit from the query
func (q *Query) WithoutLimit() *Query {
	q.limited = false
	return q
}

// WithOffset includes an offset to the query, input must be uint as offset cannot be negative value
func (q *Query) WithOffset(offset uint) *Query {
	q.offset = offset
	return q
}

// WithoutOffset removes the offset from the query
func (q *Query) WithoutOffset() *Query {
	q.offset = 0
	return q
}
