// Package cryptosquare takes strings and outputs a crypted version of the string
package cryptosquare

import (
	"math"
	"strings"
	"unicode"
)

const testVersion = 2

// Encode normalizes a string (only letters and numbers)
// then creates chunks of it using letters at regular positions
func Encode(s string) string {

	var ns string

	for _, runeValue := range strings.ToLower(s) {
		if unicode.IsDigit(runeValue) || unicode.IsLetter(runeValue) {
			ns += string(runeValue)
		}
	}

	l := int(math.Ceil(math.Sqrt(float64(len(ns)))))
	cs := make([]string, l)

	for i := 0; i < l; i++ {
		for j := i; j < len(ns); j += l {
			cs[i] += string(ns[j])
		}
	}

	return strings.Join(cs, " ")

}
