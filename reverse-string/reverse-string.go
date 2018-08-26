// Package reverse reverses a string
package reverse

// String reverses a provided string
func String(s string) string {
	runes := []rune(s)
	// swapping
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
