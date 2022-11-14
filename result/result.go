package result

import (
	. "github.com/fochive00/evil"
)

// method must have no type parameters
// func (r Result[T]) AndThen[U any](op func(val T) Result[U]) Result[U] {
//
// }

func AndThen[T any, U any](r Result[T], op func(val T) Result[U]) Result[U] {
	if r.IsOk() {
		return op(r.Unwrap())
	}

	return Err[U](r.UnwrapErr())
}
