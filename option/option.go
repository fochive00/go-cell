package option

import (
	. "github.com/fochive00/evil"
)

// Returns `None` if the option is `None`, otherwise calls `f` with the wrapped value and returns the result.
//
// Some languages call this operation flatmap.
func AndThen[T any, U any](o Option[T], f func(T) Option[U]) Option[U] {
	if o.IsNone() {
		return None[U]()
	}

	return f(o.Unwrap())
}
