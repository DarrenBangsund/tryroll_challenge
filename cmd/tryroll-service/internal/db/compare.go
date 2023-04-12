package db

// compare comparable primitives and return if the given operation is true
func compare[T string | int | int64 | int32 | float32 | float64](a, b T, op FilterOp) (bool, error) {
	switch op {
	case FilterOpEQ:
		return a == b, nil
	case FilterOpGT:
		return a > b, nil
	case FilterOpGTE:
		return a >= b, nil
	case FilterOpLT:
		return a < b, nil
	case FilterOpLTE:
		return a <= b, nil
	}

	return false, ErrDBInvalidCompareType
}
