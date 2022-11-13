// implement Result type based on https://doc.rust-lang.org/std/result/enum.Result.html

package evil

import "fmt"

// Ok(T) or Err(error)
type Result[T any] struct {
	val T
	err error
}

func (r Result[T]) String() string {
	if r.err == nil {
		return fmt.Sprintf("Ok(%v)", r.val)
	}

	return fmt.Sprintf("Err(%v)", r.err)
}

func Ok[T any](val T) Result[T] {
	return Result[T]{
		val: val,
		err: nil,
	}
}

func Err[T any](err error) Result[T] {
	return Result[T]{
		err: err,
	}
}

func (r Result[T]) IsOk() bool {
	return r.err == nil
}

func (r *Result[T]) IsOkAnd(predicate func(val *T) bool) bool {
	return r.err == nil && predicate(&r.val)
}

func (r Result[T]) IsErr() bool {
	return r.err != nil
}

func (r Result[T]) IsErrAnd(predicate func(err error) bool) bool {
	return r.err != nil && predicate(r.err)
}

func (r Result[T]) Unwrap() T {
	if r.err != nil {
		panic(fmt.Sprintf("called `Result[T].Unwrap()` on an `Err` value: %v", r.err))
	}

	return r.val
}

func (r Result[T]) UnwrapErr() error {
	if r.err == nil {
		panic(fmt.Sprintf("called `Result[T].UnwrapErr()` on an `Ok` value: %v", r.val))
	}

	return r.err
}

// method must have no type parameters
// func (r Result[T]) AndThen[U any](op func(val T) Result[U]) Result[U] {
//
// }

func AndThen[T any, U any](r Result[T], op func(val T) Result[U]) Result[U] {
	if r.IsOk() {
		return op(r.val)
	}

	return Err[U](r.err)
}
