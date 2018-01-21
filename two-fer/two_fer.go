// Package twofer gives one for you/X and one for me
package twofer

// ShareWith returns a string that's different given the input string.
func ShareWith(s string) string {
	if s == "" {
		return "One for you, one for me."
	}
	return "One for " + s + ", one for me."
}
