package evil_test

import (
	"testing"

	. "github.com/fochive00/evil"
)

type A struct {
	Str    string
	IntPtr *int
}

// TODO
func TestOption(t *testing.T) {
	var i Option[int]
	var astruct Option[A]
	var ptr Option[*A]
	var anya Option[any]

	// None if not initialized
	if i != None[int]() {
		t.Fatalf("")
	}

	if astruct != None[A]() {
		t.Fatalf("")
	}

	if ptr != None[*A]() {
		t.Fatalf("")
	}

	if anya != None[any]() {
		t.Fatalf("")
	}

	//
	switch {
	case i.IsSome():
		a := i.Unwrap()
		if a == 5 {

		}

	case i.IsNone():
	}

	// Example

}
