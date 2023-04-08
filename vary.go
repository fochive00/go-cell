package evil

import (
	"encoding/json"
	"fmt"
)

type Vary[T comparable] struct {
	value T
}

func defaultval[T any]() T {
	var defaultval T
	return defaultval
}

func (v Vary[T]) IsNone() bool {
	return v.value == defaultval[T]()
}

func (v Vary[T]) IsSome() bool {
	return v.value != defaultval[T]()
}

func (v *Vary[T]) Set(value T) {
	v.value = value
}

func (v Vary[T]) Unwrap() (value T, ok bool) {
	if v.IsNone() {
		return value, false
	}

	return v.value, true
}

func (v *Vary[T]) Take() (value T, ok bool) {
	if v.IsNone() {
		return value, false
	}

	value = v.value
	v.value = defaultval[T]()
	return value, true
}

func (v Vary[T]) String() string {
	if v.IsNone() {
		return "None"
	}

	return fmt.Sprintf("Some(%v)", v.value)
}

func (v Vary[T]) MarshalJSON() ([]byte, error) {
	if v.IsSome() {
		return json.Marshal(v.value)
	}
	return json.Marshal(nil)
}

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
