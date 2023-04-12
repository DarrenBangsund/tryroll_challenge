package db

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_compare_string(t *testing.T) {
	tests := []testCase{
		{
			input:    []interface{}{"test", "test", FilterOpEQ},
			expected: true,
			err:      nil,
		},
		{
			input:    []interface{}{"test", "test1", FilterOpEQ},
			expected: false,
			err:      nil,
		},
		{
			input:    []interface{}{"test1", "test2", FilterOpGT},
			expected: false,
			err:      nil,
		},
		{
			input:    []interface{}{"test2", "test1", FilterOpGT},
			expected: true,
			err:      nil,
		},
		{
			input:    []interface{}{"test2", "test1", FilterOpLT},
			expected: false,
			err:      nil,
		},
		{
			input:    []interface{}{"test1", "test2", FilterOpLT},
			expected: true,
			err:      nil,
		},
		{
			input:    []interface{}{"test2", "test1", FilterOpLTE},
			expected: false,
			err:      nil,
		},
		{
			input:    []interface{}{"test1", "test1", FilterOpLTE},
			expected: true,
			err:      nil,
		},
		{
			input:    []interface{}{"test1", "test2", FilterOpLTE},
			expected: true,
			err:      nil,
		},
		{
			input:    []interface{}{"test1", "test2", FilterOpGTE},
			expected: false,
			err:      nil,
		},
		{
			input:    []interface{}{"test1", "test1", FilterOpGTE},
			expected: true,
			err:      nil,
		},
		{
			input:    []interface{}{"test2", "test1", FilterOpGTE},
			expected: true,
			err:      nil,
		},
	}

	for _, test := range tests {
		match, err := compare(test.input[0].(string), test.input[1].(string), test.input[2].(FilterOp))
		assert.NoError(t, err)
		assert.Equal(t, test.expected, match, test.input)
	}
}

func Test_compare_int(t *testing.T) {
	tests := []testCase{
		{
			input:    []interface{}{int(1), int(1), FilterOpEQ},
			expected: true,
			err:      nil,
		},
		{
			input:    []interface{}{int(1), int(2), FilterOpEQ},
			expected: false,
			err:      nil,
		},
		{
			input:    []interface{}{int(11), int(12), FilterOpGT},
			expected: false,
			err:      nil,
		},
		{
			input:    []interface{}{int(12), int(11), FilterOpGT},
			expected: true,
			err:      nil,
		},
		{
			input:    []interface{}{int(12), int(11), FilterOpLT},
			expected: false,
			err:      nil,
		},
		{
			input:    []interface{}{int(11), int(12), FilterOpLT},
			expected: true,
			err:      nil,
		},
		{
			input:    []interface{}{int(12), int(11), FilterOpLTE},
			expected: false,
			err:      nil,
		},
		{
			input:    []interface{}{int(11), int(11), FilterOpLTE},
			expected: true,
			err:      nil,
		},
		{
			input:    []interface{}{int(11), int(12), FilterOpLTE},
			expected: true,
			err:      nil,
		},
		{
			input:    []interface{}{int(11), int(12), FilterOpGTE},
			expected: false,
			err:      nil,
		},
		{
			input:    []interface{}{int(11), int(11), FilterOpGTE},
			expected: true,
			err:      nil,
		},
		{
			input:    []interface{}{int(12), int(11), FilterOpGTE},
			expected: true,
			err:      nil,
		},
	}

	for _, test := range tests {
		match, err := compare(test.input[0].(int), test.input[1].(int), test.input[2].(FilterOp))
		assert.NoError(t, err, test.input)
		assert.Equal(t, test.expected, match, test.input)
	}
}

func Test_compare_int32(t *testing.T) {
	tests := []testCase{
		{
			input:    []interface{}{int32(1), int32(1), FilterOpEQ},
			expected: true,
			err:      nil,
		},
		{
			input:    []interface{}{int32(1), int32(2), FilterOpEQ},
			expected: false,
			err:      nil,
		},
		{
			input:    []interface{}{int32(11), int32(12), FilterOpGT},
			expected: false,
			err:      nil,
		},
		{
			input:    []interface{}{int32(12), int32(11), FilterOpGT},
			expected: true,
			err:      nil,
		},
		{
			input:    []interface{}{int32(12), int32(11), FilterOpLT},
			expected: false,
			err:      nil,
		},
		{
			input:    []interface{}{int32(11), int32(12), FilterOpLT},
			expected: true,
			err:      nil,
		},
		{
			input:    []interface{}{int32(12), int32(11), FilterOpLTE},
			expected: false,
			err:      nil,
		},
		{
			input:    []interface{}{int32(11), int32(11), FilterOpLTE},
			expected: true,
			err:      nil,
		},
		{
			input:    []interface{}{int32(11), int32(12), FilterOpLTE},
			expected: true,
			err:      nil,
		},
		{
			input:    []interface{}{int32(11), int32(12), FilterOpGTE},
			expected: false,
			err:      nil,
		},
		{
			input:    []interface{}{int32(11), int32(11), FilterOpGTE},
			expected: true,
			err:      nil,
		},
		{
			input:    []interface{}{int32(12), int32(11), FilterOpGTE},
			expected: true,
			err:      nil,
		},
	}

	for _, test := range tests {
		match, err := compare(test.input[0].(int32), test.input[1].(int32), test.input[2].(FilterOp))
		assert.NoError(t, err, test.input)
		assert.Equal(t, test.expected, match, test.input)
	}
}

func Test_compare_int64(t *testing.T) {
	tests := []testCase{
		{
			input:    []interface{}{int64(1), int64(1), FilterOpEQ},
			expected: true,
			err:      nil,
		},
		{
			input:    []interface{}{int64(1), int64(2), FilterOpEQ},
			expected: false,
			err:      nil,
		},
		{
			input:    []interface{}{int64(11), int64(12), FilterOpGT},
			expected: false,
			err:      nil,
		},
		{
			input:    []interface{}{int64(12), int64(11), FilterOpGT},
			expected: true,
			err:      nil,
		},
		{
			input:    []interface{}{int64(12), int64(11), FilterOpLT},
			expected: false,
			err:      nil,
		},
		{
			input:    []interface{}{int64(11), int64(12), FilterOpLT},
			expected: true,
			err:      nil,
		},
		{
			input:    []interface{}{int64(12), int64(11), FilterOpLTE},
			expected: false,
			err:      nil,
		},
		{
			input:    []interface{}{int64(11), int64(11), FilterOpLTE},
			expected: true,
			err:      nil,
		},
		{
			input:    []interface{}{int64(11), int64(12), FilterOpLTE},
			expected: true,
			err:      nil,
		},
		{
			input:    []interface{}{int64(11), int64(12), FilterOpGTE},
			expected: false,
			err:      nil,
		},
		{
			input:    []interface{}{int64(11), int64(11), FilterOpGTE},
			expected: true,
			err:      nil,
		},
		{
			input:    []interface{}{int64(12), int64(11), FilterOpGTE},
			expected: true,
			err:      nil,
		},
	}

	for _, test := range tests {
		match, err := compare(test.input[0].(int64), test.input[1].(int64), test.input[2].(FilterOp))
		assert.NoError(t, err, test.input)
		assert.Equal(t, test.expected, match, test.input)
	}
}
