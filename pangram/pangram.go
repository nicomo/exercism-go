package pangram

import "strings"

const testVersion = 2

// IsPangram tests if a string uses all 26 letters of the alphabet
func IsPangram(s string) bool {
	r := make(map[rune]bool)
	for _, v := range strings.ToLower(s) {
		if v >= 'a' && v <= 'z' {
			r[v] = true
		}
	}
	return len(r) == 26

}
