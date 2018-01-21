// Package etl transforms legacy data about scrable scores into a new given format
package etl

import "strings"

// Transform takes a map with ints as k and slice of strings as values, returns map of ints w/ strings as k
func Transform(in map[int][]string) map[string]int {
	var out = make(map[string]int)

	for key, value := range in {
		for _, v := range value {
			out[strings.ToLower(v)] = key
		}

	}

	return out
}
