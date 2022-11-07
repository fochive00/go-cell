// implement Option type based on https://doc.rust-lang.org/std/option/enum.Option.html
package evil

import "fmt"

type Option[T any] struct {
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
func None[T any]() Option[T] {
	return Option[T]{}
}

// Make a `Some` value for the given option type
func Some[T any](val T) Option[T] {
	return Option[T]{
		val:    val,
		isSome: true,
	}
}

// Returns `true` if the option is a `Some` value.
func (o *Option[T]) IsSome() bool {
	return o.isSome
}

// Returns `true` if the option is a `Some` and the value inside of it matches a predicate.
func (o *Option[T]) IsSomeAnd(predicate func(val *T) bool) bool {
	return o.isSome && predicate(&o.val)
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
//
// Because this function may panic, its use is generally discouraged. Instead, prefer to handle the `None` case explicitly, or call UnwrapOr, UnwrapOrElse, or UnwrapOrDefault.
//
// Panics if the value is a `None`.
func (o Option[T]) Unwrap() T {
	if !o.isSome {
		panic("'called `Option[T].Unwrap()` on a `None` value'")
	}

	return o.val
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

/**************
// Couldn't implement thiese method cause method must have no type parametes
//
// May support in the future
//
// Maps an Option[T] to Option[U] by applying a function to a contained value.
// func (o Option[T]) Map[U any](f (val T) U) Option[U] {}
//
// MapOr, MapOrElse
**************/

// TODO Needs to think about this
//
// Calls the provided closure with a reference to the contained value (if `Some`).
// func (o Option[T]) Inspect(f func(*T) Option[T]) Option[T] {}

/**************
// Couldn't implement these method cause method must have no type parametes
//
// May support in the future
//
// OkOr, OkOrElse
**************/

// Returns `None` if the option is `None`, otherwise returns `optb`.
//
// Arguments passed to `And` are eagerly evaluated; if you are passing the result of a function call, it is recommended to use `AndThen`, which is lazily evaluated.
func (o Option[T]) And(optb Option[T]) Option[T] {
	if o.IsNone() {
		return o
	}

	return optb
}

/**************
// Couldn't implement these method cause method must have no type parametes
//
// May support in the future
//
// AndThen
**************/

// Returns `None` if the option is `None`, otherwise calls `predicate` with the wrapped value and returns:
//
// - `Some(val)` if predicate returns true (where val is the wrapped value), and
//
// - `None` if predicate returns false.
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

// Returns `Some` if exactly one of the option itself, `optb` is `Some`, otherwise returns `None`.
func (o Option[T]) Xor(optb Option[T]) Option[T] {
	switch {
	case o.IsSome() && optb.IsSome():
		return o
	case o.IsNone() && optb.IsNone():
		return o
	default:
		return None[T]()
	}
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
		o.val = defaultValue
		o.isSome = true
	}

	return &o.val
}

// Inserts a value computed from `f` into the option if it is `None`, then returns a mutable reference to the contained value.
func (o *Option[T]) GetOrInsertWith(f func() T) *T {
	if o.IsNone() {
		o.val = f()
		o.isSome = true
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
	var defaultValue T
	o.val = defaultValue
	o.isSome = false

	return old
}

// Replaces the actual value in the option by the value given in parameter, returning the old value if present, leaving a `Some` in its place without deinitializing either one.
func (o *Option[T]) Replace(value T) Option[T] {
	if o.IsNone() {
		o.val = value
		o.isSome = true

		return None[T]()
	}

	// store the old value
	old := Some(o.val)

	// replace the value
	o.val = value

	return old
}

// type T MUST be comparable
// func (o *Option[T]) Contains(value *T) bool {
// 	   if o.IsSome() {
// 		   return o.val == *value
// 	   }
//
//     return false
// }

/**************
// Couldn't implement these method cause method must have no type parametes
//
// May support in the future
//
// Zip, ZipWith, Unzip
**************/
