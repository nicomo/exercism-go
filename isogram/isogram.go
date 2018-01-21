// Package isogram determines if a word or phrase is an isogram
package isogram

import (
	"strings"
	"unicode"
)

const testVersion = 1

// IsIsogram takes only letters and checks if they appear just once in a string
func IsIsogram(s string) bool {

	tmp := ""
	for _, v := range s {
		if !unicode.IsLetter(v) {
			continue
		}
		if strings.Contains(tmp, strings.ToLower(string(v))) {
			return false
		}
		tmp += strings.ToLower(string(v))
	}

	return true
}
