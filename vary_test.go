package evil_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/fochive00/evil"
)

type test struct {
	Value     int
	TestVary  evil.Vary[*int]
	TestVary2 evil.Vary[*int]
}

func TestVary(t *testing.T) {
	var testRef evil.Vary[*test]

	// empty ref
	assert.True(t, testRef.IsNone())
	assert.False(t, testRef.IsSome())

	// Set value to Some 1
	testRef.Set(&test{Value: 1})
	assert.False(t, testRef.IsNone())
	assert.True(t, testRef.IsSome())

	ref, ok := testRef.Unwrap()
	assert.True(t, ok)
	assert.Equal(t, ref.Value, 1)

	// Set value to 5
	ref.Value = 5

	ref, ok = testRef.Unwrap()
	assert.True(t, ok)
	assert.Equal(t, 5, ref.Value)

	// Take the value
	refTaken, ok := testRef.Take()
	assert.True(t, ok)
	assert.NotNil(t, refTaken)
	assert.Equal(t, 5, refTaken.Value)

	ref, ok = testRef.Unwrap()
	assert.False(t, ok)
	assert.Nil(t, ref)

	// Set value to some 10
	testRef.Set(&test{Value: 10})

	ref, ok = testRef.Unwrap()
	assert.True(t, ok)
	assert.NotNil(t, ref)
	assert.Equal(t, 10, ref.Value)

	// The taken value still alive
	assert.NotNil(t, refTaken)
	assert.Equal(t, 5, refTaken.Value)
}

func TestVaryMarshalJSON(t *testing.T) {
	// var testRef evil.Vary[*test]

}

func TestVaryUnmarshalJSON(t *testing.T) {
	var testRef evil.Vary[*test]

	// Unmarshal empty
	err := testRef.UnmarshalJSON([]byte(`{}`))
	assert.Nil(t, err)

	val, ok := testRef.Unwrap()
	assert.True(t, ok)
	_, ok = val.TestVary.Unwrap()
	assert.False(t, ok)

	// Unmarshal null
	err = testRef.UnmarshalJSON([]byte(`{"TestVary": null}`))
	assert.Nil(t, err)

	val, ok = testRef.Unwrap()
	assert.True(t, ok)
	_, ok = val.TestVary.Unwrap()
	assert.False(t, ok)

	// Unmarshal value
	err = testRef.UnmarshalJSON([]byte(`{"TestVary": 5}`))
	assert.Nil(t, err)

	val, ok = testRef.Unwrap()
	assert.True(t, ok)
	a, ok := val.TestVary.Unwrap()
	assert.True(t, ok)
	assert.Equal(t, 5, *a)
}
