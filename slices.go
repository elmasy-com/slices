package slices

import (
	"fmt"
)

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
func RemoveElem[T comparable](s []T, e T) []T {

	var v []T

	for i := range s {

		if s[i] == e {
			continue
		}

		v = append(v, s[i])
	}

	return v
}

// RemoveIndex removes element with index i from s.
// If s is nil, returns nil.
func RemoveIndex[T any](s []T, i int) []T {

	var v []T

	for n := range s {

		if n == i {
			continue
		}

		v = append(v, s[n])
	}

	return v
}

// Join joins elements of s with sep and return as a string.
// If s is nil or zero-length, returns an empty string ("").
func Join[T fmt.Stringer](s []T, sep string) string {

	n := len(s)
	var v string

	for i := range s {

		v += s[i].String()

		if i != n-1 {
			v += sep
		}
	}

	return v
}

// Contain returns if s contains e.
func Contain[T comparable](s []T, e T) bool {

	for i := range s {
		if s[i] == e {
			return true
		}
	}

	return false
}

// Strings convert a slice of T to a string slice.
// If s is nil, returns nil.
func Strings[T fmt.Stringer](s []T) []string {

	var v []string

	for i := range s {
		v = append(v, s[i].String())
	}

	return v
}
