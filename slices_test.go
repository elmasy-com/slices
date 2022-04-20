package slices

import (
	"fmt"
	"testing"
)

type Test struct {
	i int
	f float64
}

// Implement comparable interface on Test
func (a Test) Compare(b Test) int {
	switch {
	case a.i < b.i:
		return -1
	case a.i > b.i:
		return 1
	case a.f < b.f:
		return -1
	case a.f > b.f:
		return 1
	default:
		return 0
	}
}

func (t Test) String() string {
	return fmt.Sprintf("%d|%.0f", t.i, t.f)
}

// Test with a user defined type
func TestCopyUser(t *testing.T) {

	a := []Test{{0, 0}, {1, 1}, {2, 2}, {3, 3}, {4, 4}}

	b := Copy(a)

	if b == nil {
		t.Fatalf("FAILED!")
	}
}

// Test with a builtin type
func TestCopyBuiltin(t *testing.T) {

	a := []int{0, 1, 2, 3, 4}

	b := Copy(a)

	if b == nil {
		t.Fatalf("FAILED!")
	}
}

func TestRemoveElemUser(t *testing.T) {

	a := []Test{{0, 0}, {1, 1}, {2, 2}, {3, 3}, {4, 4}}

	b := RemoveElem(a, Test{0, 0})

	c := Test{1, 1}
	if b[0] != c {
		t.Fatalf("FAILED!")
	}
}

func TestRemoveElemBuiltin(t *testing.T) {

	a := []int{0, 1, 2, 3, 4}

	b := RemoveElem(a, 0)

	if b[0] != 1 {
		t.Fatalf("FAILED!")
	}

}

func TestRemoveIndexUser(t *testing.T) {

	a := []Test{{0, 0}, {1, 1}, {2, 2}, {3, 3}, {4, 4}}

	b := RemoveIndex(a, 0)

	c := Test{1, 1}
	if b[0] != c {
		t.Fatalf("FAILED!")
	}
}

func TestRemoveIndexBuiltin(t *testing.T) {

	a := []int{0, 1, 2, 3, 4}

	b := RemoveIndex(a, 0)

	if b[0] != 1 {
		t.Fatalf("FAILED!")
	}
}

func TestJoinUser(t *testing.T) {

	a := []Test{{0, 0}, {1, 1}, {2, 2}, {3, 3}, {4, 4}}

	b := Join(a, " ")

	if b != "0|0 1|1 2|2 3|3 4|4" {
		t.Fatalf("FAILED!")
	}
}

func TestContainUser(t *testing.T) {

	a := []Test{{0, 0}, {1, 1}, {2, 2}, {3, 3}, {4, 4}}

	if !Contain(a, Test{0, 0}) {
		t.Fatalf("FAILED!")
	}
}

func TestContainBuiltin(t *testing.T) {

	a := []int{0, 1, 2, 3, 4}

	if !Contain(a, 0) {
		t.Fatalf("FAILED!")
	}
}

func TestStringsUser(t *testing.T) {

	a := []Test{{0, 0}, {1, 1}, {2, 2}, {3, 3}, {4, 4}}

	b := Strings(a)

	if b[0] != "0|0" {
		t.Fatalf("FAILED!")
	}
}
