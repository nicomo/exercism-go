// Package anagram, given a word and a list of possible anagrams, selects the correct sublist
package anagram

import (
	"strings"
)

const testVersion = 2

// Detect checks for anagrams in a list of words
func Detect(s string, c []string) []string {
	r := []string{}

	// subjects are case insensitive
	s = strings.ToLower(s)

	for _, v := range c {

		// if lengths differ of it's the exact same word, continue
		if len(s) != len(v) || s == strings.ToLower(v) {
			continue
		}

		var count int

		for _, w := range v {
			if strings.Count(s, strings.ToLower(string(w))) != strings.Count(strings.ToLower(v), strings.ToLower(string(w))) {
				break
			}
			count++
		}
		if count == len(s) {
			r = append(r, v)
		}

	}

	return r
}
