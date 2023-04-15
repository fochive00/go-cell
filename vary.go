package cell

import (
	"encoding/json"
	"fmt"
)

// A Vary type that can be used to represent a value that may or may not be
// the default value for the type.
//
// Since nil value is the default value for pointer type, Vary` type is
// especially useful for wrapping pointer types, providing a safe way to
// inspect the value without worrying about panicking due to dereferencing
// a nil pointer by accident.
//
// A benefit of using `Vary` type over `Option` type is that, `Vary` type
// will not take any extra memory space, since it does not need to store a boolean flag.
type Vary[T comparable] struct {
	value T
}

// Make a default value for the given type.
func defaultval[T any]() (defaultval T) {
	return
}

// Make a Vary instance with a provided value.
func NewVary[T comparable](value T) Vary[T] {
	return Vary[T]{value: value}
}

// Returns `true` if the inner value is the default value.
func (v Vary[T]) IsNone() bool {
	return v.value == defaultval[T]()
}

// Returns `true` if the inner value is not the default value.
func (v Vary[T]) IsSome() bool {
	return v.value != defaultval[T]()
}

// Set the inner value to the provided value.
func (v *Vary[T]) Set(value T) {
	v.value = value
}

// Unwarp returns the inner value and a boolean indicating whether the value
// is the default value.
func (v Vary[T]) Unwrap() (value T, ok bool) {
	if v.IsNone() {
		return value, false
	}

	return v.value, true
}

// UnwrapOr Returns the contained value if it is not the default value,
// otherwise returns the provided value.
func (v Vary[T]) UnwrapOr(val T) T {
	if v.IsNone() {
		return val
	}

	return v.value
}

// Take returns the inner value and a boolean indicating whether the value has been taken.
func (v *Vary[T]) Take() (value T, ok bool) {
	if v.IsNone() {
		return value, false
	}

	value = v.value
	v.value = defaultval[T]()
	return value, true
}

// Debug returns a string representation of the Vary instance.
//
//	"None" is returned if the inner value is the default value.
//	"Some(%v)" is returned if the inner value is not the default value.
//
// The `%v` is replaced with the string representation of the inner value.
func (v Vary[T]) Debug() string {
	if v.IsNone() {
		return "None"
	}

	return fmt.Sprintf("Some(%v)", v.value)
}

// MarshalJSON implements the json.Marshaler interface.
func (v Vary[T]) MarshalJSON() ([]byte, error) {
	if v.IsSome() {
		return json.Marshal(v.value)
	}
	return json.Marshal(nil)
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (v *Vary[T]) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		v.value = defaultval[T]()
		return nil
	}

	var value T
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}

	v.value = value
	return nil
}
