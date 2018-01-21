// Package strain implements the `keep` and `discard` operation on collections
package strain

const testVersion = 1

// Ints is the slice of integers to return
type Ints []int

// Lists are the input values
type Lists [][]int

// Strings are the lists of strings to work with
type Strings []string

// Keep weeds out the false results, keeps the rest
func (list Ints) Keep(f func(int) bool) Ints {

	var r Ints
	if list == nil {
		return nil
	}
	for _, v := range list {
		if f(v) {
			r = append(r, v)
		}
	}
	return r
}

// Discard returns the discarded values in the lists
func (list Ints) Discard(f func(int) bool) Ints {
	var r Ints
	if list == nil {
		return nil
	}
	for _, v := range list {
		if !f(v) {
			r = append(r, v)
		}
	}
	return r
}

// Keep returns the true values in the lists
func (l Lists) Keep(f func([]int) bool) Lists {
	var result Lists
	for _, v := range l {
		if f(v) {
			result = append(result, v)
		}
	}
	return result
}

// Keep returns the true values in the list of strings
func (strings Strings) Keep(f func(string) bool) Strings {
	var s Strings
	for _, v := range strings {
		if f(v) {
			s = append(s, v)
		}
	}
	return s
}
