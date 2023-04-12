package db

import (
	"fmt"
	"reflect"
)

// Collection is the list of documents that hold the persisted records
type Collection struct {
	name       string
	collection []*document
}

// NewCollection returns a new *Collection
func NewCollection(name string) *Collection {
	return &Collection{
		name:       name,
		collection: make([]*document, 0),
	}
}

// Insert handles converting the value into a queryable document
func (c *Collection) Insert(value interface{}) {
	c.collection = append(c.collection, newDocument(value))
}

func validateTypeMatch(a, b interface{}) error {
	if reflect.TypeOf(a).String() != reflect.TypeOf(b).String() {
		return ErrDBInvalidCompareTypeMismatch
	}

	return nil
}

// Find accepts a Query type and filters the collection based on the Query
func (c *Collection) Find(query Query) ([]interface{}, error) {
	out := []interface{}{}

	// iterate over the collection and filter
	for _, doc := range c.collection {
		match, err := multiFilterMatch(query.filters, doc)

		if err != nil {
			return nil, fmt.Errorf("error during match: %w", err)
		}

		if match {
			out = append(out, doc.value)
		}
	}

	start := query.offset
	end := query.offset + query.limit

	if start > uint(len(out)-1) {
		start = uint(len(out))
	}

	if query.limited {
		if start+end > uint(len(out)) {
			end = uint(len(out))
		}

		out = out[start:end]
	} else {
		out = out[start:]
	}

	return out, nil
}
