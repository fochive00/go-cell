// implement Result type based on https://doc.rust-lang.org/std/result/enum.Result.html

package evil

import "fmt"

type Result[T any] struct {
	val T
	err error
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
