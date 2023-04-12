package db

type FilterOp string

const (
	// FilterOpGTE compares if 'a' is greater than or equal to 'b'
	FilterOpGTE FilterOp = "GTE"
	// FilterOpGT compares if 'a' is greater than 'b'
	FilterOpGT FilterOp = "GT"
	// FilterOpLTE compares if 'a' is less than or equal to 'b'
	FilterOpLTE FilterOp = "LTE"
	// FilterOpLT compares if 'a' is less than 'b'
	FilterOpLT FilterOp = "LT"
	// FilterOpEQ compares if 'a' is equal to 'b'
	FilterOpEQ FilterOp = "EQ"
)

// MatcherFunc is a helper type that allows the caller to make custom filters for complex data types
// return true to include the record in the query
type MatcherFunc func(a interface{}) (bool, error)

// Filter defines the filter type
type Filter struct {
	Key   string      // the key to compare when making a query
	Value interface{} // the value to compare the field to
	Op    FilterOp    // the operation to execute when making a comparison
	Func  MatcherFunc // optional field to include a custom filter (Op is ignored if this value is present)
}

// MultiFilterOp allows for multiple filters to be applied in a query
type MultiFilterOp string

const (
	// MultiFilterOpAND defines the logical operation to apply when making a multifilter query
	// given filters x and y, the record is included in the query if x && y
	MultiFilterOpAND = "AND"
	// MultiFilterOpOR defines the logical operation to apply when making a multifilter query
	// given filters x and y, the record is included in the query if x || y
	MultiFilterOpOR = "OR"
)

// MultiFilter is the data type for multifilters
type MultiFilter struct {
	Filters []Filter      // the filters to apply to the collection when executing a query
	Op      MultiFilterOp // the logical operation to apply when executing a query
}

func filterMatch(filter Filter, doc *document) (bool, error) {
	if filter.Func != nil {
		return filter.Func(doc.value)
	}

	field, ok := doc.fields[filter.Key]
	if !ok {
		return false, ErrDBInvalidFilter
	}

	if err := validateTypeMatch(filter.Value, field.Value); err != nil {
		return false, err
	}

	switch filter.Value.(type) {
	case int:
		return compare(field.Value.(int), filter.Value.(int), filter.Op)
	case int32:
		return compare(field.Value.(int32), filter.Value.(int32), filter.Op)
	case int64:
		return compare(field.Value.(int64), filter.Value.(int64), filter.Op)
	case string:
		return compare(field.Value.(string), filter.Value.(string), filter.Op)
	case float32:
		return compare(field.Value.(float32), filter.Value.(float32), filter.Op)
	case float64:
		return compare(field.Value.(float64), filter.Value.(float64), filter.Op)
	default:
		return false, ErrDBInvalidFilterValue
	}
}

func multiFilterMatch(multiFilter MultiFilter, doc *document) (bool, error) {
	if len(multiFilter.Filters) < 1 {
		return true, nil
	}

	for _, filter := range multiFilter.Filters {
		if match, err := filterMatch(filter, doc); err == nil {
			switch multiFilter.Op {
			case MultiFilterOpAND:
				if !match {
					return false, nil
				}
			case MultiFilterOpOR:
				if match {
					return true, nil
				}
			default:
				return false, ErrDBInvalidMultiFilterOp
			}
		} else {
			return false, err
		}
	}

	switch multiFilter.Op {
	case MultiFilterOpAND:
		return true, nil
	case MultiFilterOpOR:
		return false, nil
	}

	return false, nil
}
