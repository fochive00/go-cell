// implement Option type based on https://doc.rust-lang.org/std/option/enum.Option.html
package evil

type Option[T any] struct {
	val    T
	isSome bool
}

func None[T any]() Option[T] {
	return Option[T]{}
}

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
	return o.val
}

/**************
// Couldn't implement thiese method cause method must have no type parametes
//
// May support in the future
//
// Maps an Option[T] to Option[U] by applying a function to a contained value.
// func (o Option[T]) Map[U any](f (val T) U) Option[U] {
// }
//
// MapOr(), MapOrElse()
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
// OkOr(), OkOrElse()
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
// AndThen()
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

// TODO
func (o Option[T]) Or() {}

func (o Option[T]) OrElse() {}

func (o Option[T]) Xor() {}

func (o *Option[T]) Insert() {}

func (o *Option[T]) GetOrInsert() {}

func (o *Option[T]) GetOrInsertDefault() {}

func (o *Option[T]) GetOrInsertWith() {}

func (o *Option[T]) Take() {}

func (o Option[T]) Replace() {}

func (o Option[T]) Contains() {}
