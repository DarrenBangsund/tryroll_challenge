package db

import "fmt"

type DBError error

var (
	// ErrDBInvalidCompareTypeMismatch occurs when the caller attempts to compare with a type that doesn't match the persisted type
	ErrDBInvalidCompareTypeMismatch DBError = fmt.Errorf("compare type mismatch")
	// ErrDBInvalidCompareType occurs when an invalid compare operation is provided to a filter
	ErrDBInvalidCompareType DBError = fmt.Errorf("compare type invalid")
	// ErrDBInvalidFilter occurs when a filter is invalid, eg. when the filter contains a key that doesn't exist on the persisted type
	ErrDBInvalidFilter DBError = fmt.Errorf("filter invalid")
	// ErrDBInvalidFilterValue occurs when the caller attempts to do a primitive filter on a field that isn't a primitive
	ErrDBInvalidFilterValue DBError = fmt.Errorf("filter match value invalid")
	// ErrDBInvalidMultiFilterOp occurs when a multifilter operation is invalid, eg. when the operation isn't either AND or OR
	ErrDBInvalidMultiFilterOp DBError = fmt.Errorf("multi filter Op invalid")
)
