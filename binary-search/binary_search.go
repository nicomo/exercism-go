// Package binarysearch
package binarysearch

import (
	"fmt"
)

// SearchInts does a binary search on a sorted slice of ints
// to find the position of a given int
func SearchInts(A []int, k int) int {

	L := 0
	R := len(A) - 1

	for L <= R {
		m := (L + R) / 2
		if k <= A[m] {
			R = m - 1
		} else {
			L = m + 1
		}
	}
	return L
}

// Message prints message corresponding to various results of SearchInts
func Message(A []int, k int) string {
	i := SearchInts(A, k)
	upperOut := i == len(A)
	found := !upperOut && A[i] == k // watch out for order so as not to have an Index out of bound panic

	// order of cases matters
	switch {
	case len(A) == 0:
		return fmt.Sprintf("slice has no values")
	case i == 0 && found:
		return fmt.Sprintf("%d found at beginning of slice", k)
	case i == len(A)-1 && found:
		return fmt.Sprintf("%d found at end of slice", k)
	case found:
		return fmt.Sprintf("%d found at index %d", k, i)
	case i == 0 && !found:
		return fmt.Sprintf("%d < all values", k)
	case upperOut && !found:
		return fmt.Sprintf("%d > all %d values", k, len(A))
	case !found:
		return fmt.Sprintf("%d > %d at index %d, < %d at index %d", k, A[i-1], i-1, A[i], i)

	default:
		return ""

	}

}
