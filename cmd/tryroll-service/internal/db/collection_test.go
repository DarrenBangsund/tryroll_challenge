package db

import (
	"fmt"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testCase struct {
	input    []interface{}
	expected interface{}
	err      error
}

func Test_validateTypeMatch(t *testing.T) {
	tests := []testCase{
		{
			input: []interface{}{"string", "string"},
			err:   nil,
		},
		{
			input: []interface{}{"string", 1234},
			err:   ErrDBInvalidCompareTypeMismatch,
		},
	}

	for _, test := range tests {
		err := validateTypeMatch(test.input[0], test.input[1])
		assert.ErrorIs(t, test.err, err, test.input)
	}
}

func Test_filterMatch(t *testing.T) {
	tests := []testCase{
		{
			expected: true,
			input: []interface{}{
				Filter{
					Key:   "testKey",
					Value: "testValue",
					Op:    FilterOpEQ,
				},
				newDocument(map[string]interface{}{
					"testKey": "testValue",
				}),
			},
		},
		{
			expected: true,
			input: []interface{}{
				Filter{
					Key:   "testKey",
					Value: "testValue",
					Op:    FilterOpEQ,
				},
				newDocument(struct {
					TestKey string `db:"testKey"`
				}{
					TestKey: "testValue",
				}),
			},
		},
		{
			expected: false,
			input: []interface{}{
				Filter{
					Key:   "testKey",
					Value: 1234,
					Op:    FilterOpEQ,
				},
				newDocument(struct {
					TestKey string `db:"testKey"`
				}{
					TestKey: "testValue",
				}),
			},
			err: ErrDBInvalidCompareTypeMismatch,
		},
	}

	for _, test := range tests {
		match, err := filterMatch(test.input[0].(Filter), test.input[1].(*document))
		assert.ErrorIs(t, test.err, err, test.input)
		assert.Equal(t, match, test.expected, test.input)
	}
}

type testModel struct {
	TestKey    string `db:"testKey"`
	TestKeyInt int    `db:"testKeyInt"`
}

func setupTest() *Collection {
	c := NewCollection("testcollection")

	for i := 0; i < 10; i++ {
		c.Insert(&testModel{
			TestKey:    fmt.Sprintf("testValue_%d", i),
			TestKeyInt: i,
		})
	}

	return c
}

func Test_Find(t *testing.T) {
	tests := []testCase{
		{
			input: []interface{}{
				Filter{
					Key:   "testKey",
					Value: "testValue_1",
					Op:    FilterOpEQ,
				},
			},
			expected: func(values []*testModel) bool {
				return len(values) == 1
			},
		},
		{
			input: []interface{}{
				Filter{
					Key:   "testKeyInt",
					Value: 2,
					Op:    FilterOpLT,
				},
			},
			expected: func(values []*testModel) bool {
				log.Println(values, len(values))
				return len(values) == 2
			},
		},
		{
			input: []interface{}{
				Filter{
					Key:   "testKeyInt",
					Value: 2,
					Op:    FilterOpLTE,
				},
			},
			expected: func(values []*testModel) bool {
				return len(values) == 3
			},
		},
	}

	c := setupTest()

	// test basic queries
	for _, test := range tests {
		q := NewQuery()

		for _, filter := range test.input {
			q.WithFilter(filter.(Filter))
		}

		values, err := c.Find(q)
		assert.ErrorIs(t, test.err, err, test.input)

		vals := make([]*testModel, 0)
		for _, v := range values {
			vals = append(vals, v.(*testModel))
		}

		assert.True(t, test.expected.(func(values []*testModel) bool)(vals), test.input)
	}

	// test basic queries with limit and offset
	{
		q := NewQuery()

		q.WithFilter(tests[0].input[0].(Filter))
		q.WithLimit(1)
		r, err := c.Find(q)
		assert.NoError(t, err, r)
		assert.Len(t, r, 1, r)
	}

	{
		// we know this filter yields 3 results
		f := tests[2].input[0].(Filter)
		q := NewQuery()

		// test with an offset of 1, should yield 2
		q.WithFilter(f)
		q.WithOffset(1)
		r, err := c.Find(q)
		assert.NoError(t, err, r)
		assert.Len(t, r, 2, r)

		// test with a limit of 1, should yield 1
		q.WithoutOffset().WithLimit(1)
		r, err = c.Find(q)
		assert.NoError(t, err, r)
		assert.Len(t, r, 1, r)

		// test with a limit of 1, should yield 1
		q.WithOffset(2).WithLimit(1)
		r, err = c.Find(q)
		assert.NoError(t, err, r)
		assert.Len(t, r, 1, r)
		// with an offset of 2, we should receive the 3rd element which has an int of 2
		assert.Equal(t, 2, r[0].(*testModel).TestKeyInt, r)

		// test with huge numbers, should yield 0
		q.WithOffset(10000).WithLimit(10000)
		r, err = c.Find(q)
		assert.NoError(t, err, r)
		assert.Len(t, r, 0, r)
	}
}
