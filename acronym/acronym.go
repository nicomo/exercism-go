// Package acronym returns acronyms for strings
package acronym

import (
	"strings"
	"unicode"
)

const testVersion = 3

// Abbreviate returns the acronym for a string
func Abbreviate(s string) string {

	f := func(c rune) bool {
		return !unicode.IsLetter(c)
	}

	tmp := strings.FieldsFunc(s, f)

	var a string
	for _, v := range tmp {
		a += strings.ToUpper(string(v[0]))
	}

	return a
}
