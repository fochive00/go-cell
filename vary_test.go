package cell_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	cell "github.com/fochive00/go-cell"
)

type test struct {
	Value     int
	TestVary  cell.Vary[*int]
	TestVary2 cell.Vary[*int]
}

// testing the basic functionality of the Vary type
func TestVary(t *testing.T) {
	var testvary cell.Vary[*test]

	// empty ref
	assert.True(t, testvary.IsNone())
	assert.False(t, testvary.IsSome())

	// Set value to Some 1
	testvary.Set(&test{Value: 1})
	assert.False(t, testvary.IsNone())
	assert.True(t, testvary.IsSome())

	ref, ok := testvary.Unwrap()
	assert.True(t, ok)
	assert.Equal(t, ref.Value, 1)

	// Set value to 5
	ref.Value = 5

	ref, ok = testvary.Unwrap()
	assert.True(t, ok)
	assert.Equal(t, 5, ref.Value)

	// Take the value
	refTaken, ok := testvary.Take()
	assert.True(t, ok)
	assert.NotNil(t, refTaken)
	assert.Equal(t, 5, refTaken.Value)

	ref, ok = testvary.Unwrap()
	assert.False(t, ok)
	assert.Nil(t, ref)

	// Set value to some 10
	testvary.Set(&test{Value: 10})

	ref, ok = testvary.Unwrap()
	assert.True(t, ok)
	assert.NotNil(t, ref)
	assert.Equal(t, 10, ref.Value)

	// The taken value still alive
	assert.NotNil(t, refTaken)
	assert.Equal(t, 5, refTaken.Value)
}

// testing json marshaling for Vary type.
func TestVaryMarshalJSON(t *testing.T) {
	// var testRef cell.Vary[*test]

}

// testing json unmarshaling for Vary type.
func TestVaryUnmarshalJSON(t *testing.T) {
	var testvary cell.Vary[*test]

	// Unmarshal empty
	err := testvary.UnmarshalJSON([]byte(`{}`))
	assert.Nil(t, err)

	val, ok := testvary.Unwrap()
	assert.True(t, ok)
	_, ok = val.TestVary.Unwrap()
	assert.False(t, ok)

	// Unmarshal null
	err = testvary.UnmarshalJSON([]byte(`{"TestVary": null}`))
	assert.Nil(t, err)

	val, ok = testvary.Unwrap()
	assert.True(t, ok)
	_, ok = val.TestVary.Unwrap()
	assert.False(t, ok)

	// Unmarshal value
	err = testvary.UnmarshalJSON([]byte(`{"TestVary": 5}`))
	assert.Nil(t, err)

	val, ok = testvary.Unwrap()
	assert.True(t, ok)
	a, ok := val.TestVary.Unwrap()
	assert.True(t, ok)
	assert.Equal(t, 5, *a)
}
