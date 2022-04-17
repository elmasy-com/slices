package slices

import (
	"fmt"
	"testing"
)

type Test struct {
	i int
	f float64
}

// Implement Comparable interface on Test
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

func TestCopy(t *testing.T) {

	a := []Test{{0, 0}, {1, 1}, {2, 2}, {3, 3}, {4, 4}}

	b := Copy(a)

	if b == nil {
		t.Fatalf("Copy() failed!")
	}
}

func TestRemoveElem(t *testing.T) {

	a := []Test{{0, 0}, {1, 1}, {2, 2}, {3, 3}, {4, 4}}

	b := RemoveElem(a, Test{0, 0})

	c := Test{1, 1}
	if b[0] != c {
		t.Fatalf("RemoveElem() failed!")
	}
}

func TestRemoveIndex(t *testing.T) {

	a := []Test{{0, 0}, {1, 1}, {2, 2}, {3, 3}, {4, 4}}

	b := RemoveIndex(a, 0)

	c := Test{1, 1}
	if b[0] != c {
		t.Fatalf("RemoveIndex() failed!")
	}
}

func TestJoin(t *testing.T) {

	a := []Test{{0, 0}, {1, 1}, {2, 2}, {3, 3}, {4, 4}}

	b := Join(a, " ")

	if b != "0|0 1|1 2|2 3|3 4|4" {
		t.Fatalf("Join() failed!")
	}
}

func TestSortInc(t *testing.T) {

	a := []Test{{0, 0}, {2, 2}, {1, 1}, {4, 4}, {3, 3}}

	SortInc(a)

	if a[1].Compare(Test{1, 1}) != 0 {
		t.Fatalf("Join() failed!")
	}
}

func TestSortDec(t *testing.T) {

	a := []Test{{0, 0}, {2, 2}, {1, 1}, {4, 4}, {3, 3}}

	SortDec(a)

	if a[1].Compare(Test{3, 3}) != 0 {
		t.Fatalf("Join() failed!")
	}
}
