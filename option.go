package evil

import (
	"encoding/json"
	"fmt"
)

type Option[T comparable] struct {
	val    T
	isSome bool
}

func (o Option[T]) String() string {
	if o.isSome {
		return fmt.Sprintf("Some(%v)", o.val)
	}

	return "None"
}

// Make a `None` value for the given option type.
func None[T comparable]() Option[T] {
	return Option[T]{}
}

// Make a `Some` value for the given option type
func Some[T comparable](val T) Option[T] {
	return Option[T]{
		val:    val,
		isSome: true,
	}
}

// Returns `true` if the option is a `Some` value.
func (o *Option[T]) IsSome() bool {
	return o.isSome
}

// Returns `true` if the option is a `None` value.
func (o *Option[T]) IsNone() bool {
	return !o.isSome
}

// Returns the contained `Some` value.
//
// Panics if the value is a `None` with a custom panic message provided by `msg`.
func (o Option[T]) Expect(msg string) T {
	if !o.isSome {
		panic(msg)
	}

	return o.val
}

// Returns the contained `Some` value.
func (o Option[T]) Unwrap() (value T, ok bool) {
	return o.val, o.isSome
}

// Returns the contained `Some` value or a provided default.
//
// Arguments passed to UnwrapOr are eagerly evaluated; if you are passing the result of a function call, it is recommended to use UnwrapOrElse, which is lazily evaluated.
func (o Option[T]) UnwrapOr(val T) T {
	if o.isSome {
		return o.val
	}

	return val
}

// Returns the contained `Some` value or computes it from a closure.
func (o Option[T]) UnwrapOrElse(f func() T) T {
	if o.isSome {
		return o.val
	}

	return f()
}

// Returns the contained `Some` value or a default.
//
// If `Some`, returns the contained value, otherwise if `None`, returns the `default value` for that type.
func (o Option[T]) UnwrapOrDefault() T {
	if o.IsSome() {
		return o.val
	}

	var defaultValue T
	return defaultValue
}

// Returns `None` if the option is `None`, otherwise returns `optb`.
//
// Arguments passed to `And` are eagerly evaluated; if you are passing the result of a function call, it is recommended to use `AndThen`, which is lazily evaluated.
func (o Option[T]) And(optb Option[T]) Option[T] {
	if o.IsNone() {
		return o
	}

	return optb
}

// Returns `None` if the option is `None`, otherwise calls `predicate` with the wrapped value and returns:
//
//   - `Some(val)` if predicate returns true (where val is the wrapped value), and
//   - `None` if predicate returns false.
func (o Option[T]) Filter(predicate func(T) bool) Option[T] {
	if o.IsNone() {
		return o
	}

	if !predicate(o.val) {
		return None[T]()
	}

	return o
}

// Returns the option if it contains a value, otherwise returns `optb`.
//
// Arguments passed to or are eagerly evaluated; if you are passing the result of a function call, it is recommended to use `OrElse`, which is lazily evaluated.
func (o Option[T]) Or(optb Option[T]) Option[T] {
	if o.IsSome() {
		return o
	}

	return optb
}

// Returns the option if it contains a value, otherwise calls `f` and returns the result.
func (o Option[T]) OrElse(f func() Option[T]) Option[T] {
	if o.IsSome() {
		return o
	}

	return f()
}

// Inserts `value` into the option, then returns a reference to it.
//
// If the option already contains a value, the old value is dropped.
//
// See also `GetOrInsert`, which doesn’t update the value if the option already contains `Some`.
func (o *Option[T]) Insert(value T) *T {
	o.val = value
	o.isSome = true

	return &o.val
}

// Inserts value into the option if it is `None`, then returns a reference to the contained value.
//
// See also Insert, which updates the value even if the option already contains `Some`.
func (o *Option[T]) GetOrInsert(value T) *T {
	if o.IsNone() {
		o.val = value
		o.isSome = true
	}

	return &o.val
}

// Inserts the default value into the option if it is `None`, then returns a mutable reference to the contained value.
func (o *Option[T]) GetOrInsertDefault() *T {
	if o.IsNone() {
		var defaultValue T
		*o = Some(defaultValue)
	}

	return &o.val
}

// Takes the value out of the option, leaving a `None` in its place.
func (o *Option[T]) Take() Option[T] {
	if o.IsNone() {
		return None[T]()
	}

	// store the old value
	old := Some(o.val)

	// turn o into a `None`
	*o = None[T]()
	return old
}

// Replaces the actual value in the option by the value given in parameter, returning the old value if present, leaving a `Some` in its place without deinitializing either one.
func (o *Option[T]) Replace(value T) Option[T] {
	if o.IsNone() {
		*o = Some(value)
		return None[T]()
	}

	// store the old value
	old := *o

	// replace the value
	*o = Some(value)
	return old
}

func (o *Option[T]) Contains(value T) bool {
	if o.IsNone() {
		return false
	}

	return o.val == value
}

func (o Option[T]) MarshalJSON() ([]byte, error) {
	if o.IsSome() {
		return json.Marshal(o.val)
	}
	return json.Marshal(nil)
}

func (o *Option[T]) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		*o = None[T]()
		return nil
	}

	var value T
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}

	*o = Some(value)
	return nil
}
