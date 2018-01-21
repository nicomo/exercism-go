// Package bob can't be bothered
package bob

import (
	"strings"
	"unicode"
)

const testVersion = 2

// Hey manages questions and provides the right answers
func Hey(s string) string {
	s = strings.TrimSpace(s)

	switch {
	case s == "":
		return "Fine. Be that way!"
	case upper(s, unicode.IsUpper) && !upper(s, unicode.IsLower): // all upper, no lower
		return "Whoa, chill out!"
	case s[len(s)-1] == '?':
		return "Sure."
	default:
		return "Whatever."
	}
}

// upper determines if any items in a string are upper / lower case
func upper(items string, b func(rune) bool) bool {
	for _, item := range items {
		if b(item) {
			return true
		}
	}
	return false
}
