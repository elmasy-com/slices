package slices

import (
	"sort"
)

type Comparable[T any] interface {
	// Compare should return:
	// int < 0  if a < b
	// int > 0  if a > b
	// int == 0 if a == b
	Compare(T) int
}

type Stringer interface {
	String() string
}

// Copy create a copy of s and returns it.
// If copy fail, returns nil.
func Copy[T any](s []T) []T {

	n := len(s)

	v := make([]T, n)

	if copy(v, s) != n {
		return nil
	}

	return v
}

// RemoveElem removes element e from s.
func RemoveElem[T Comparable[T]](s []T, e T) []T {

	v := make([]T, 0)

	for i := range s {

		if s[i].Compare(e) == 0 {
			continue
		}

		v = append(v, s[i])
	}

	return v
}

// RemoveIndex removes element with index i from s.
func RemoveIndex[T any](s []T, i int) []T {

	v := make([]T, 0)

	for n := range s {

		if n == i {
			continue
		}

		v = append(v, s[n])
	}

	return v
}

// Join joins elements of s with sep and return as a string.
func Join[T Stringer](s []T, sep string) string {

	n := len(s)
	v := ""

	for i := range s {

		v += s[i].String()

		if i != n-1 {
			v += sep
		}
	}

	return v
}

// SortInc create a copy of s and do an incremental sort (low -> high).
// Return nil if copy fail.
func SortInc[T Comparable[T]](s []T) {

	sort.Slice(s, func(i, j int) bool { return s[i].Compare(s[j]) < 0 })
}

// SortDec create a copy of s and do an decremental sort (high -> low).
// Return nil if copy fail.
func SortDec[T Comparable[T]](s []T) {

	sort.Slice(s, func(i, j int) bool { return s[i].Compare(s[j]) > 0 })
}
