// Package atbash implements the atbash cipher
package atbash

import (
	"strings"
	"unicode"
)

const (
	testVersion = 2
	plain       = "abcdefghijklmnopqrstuvwxyz"
	cypher      = "zyxwvutsrqponmlkjihgfedcba"
)

// Atbash encrypts a string
// number goes in as is
// non letters skipped
// remaining letters cyphered
// spaces inserted every 5 characters (except at the very end of the string)
func Atbash(s string) string {

	var tmp string

	// encrypt
	for _, v := range s {

		if unicode.IsNumber(v) {
			tmp += string(v)
		}

		if !unicode.IsLetter(v) {
			continue
		}

		i := strings.Index(plain, strings.ToLower(string(v)))
		tmp += string(cypher[i])
	}

	// insert spaces every 5 characters
	var res string
	for i, c := range tmp {
		res = res + string(c)
		if i > 0 && (i+1)%5 == 0 && i+1 != len(tmp) {
			res += " "
		}
	}

	return res
}
