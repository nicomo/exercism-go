// scrabble computes the scrabble score for a given word (string)
package scrabble

import "strings"

const (
	testVersion = 4
)

// Score takes a word/string and computes points
// based on the value of the letters
func Score(s string) int {
	points := map[string]int{
		"a": 1,
		"e": 1,
		"i": 1,
		"o": 1,
		"u": 1,
		"l": 1,
		"n": 1,
		"r": 1,
		"s": 1,
		"t": 1,
		"d": 2,
		"g": 2,
		"b": 3,
		"c": 3,
		"m": 3,
		"p": 3,
		"f": 4,
		"h": 4,
		"v": 4,
		"w": 4,
		"y": 4,
		"k": 5,
		"j": 8,
		"x": 8,
		"q": 10,
		"z": 10,
	}
	score := 0

	// range over string, no need to split
	for _, letter := range s {
		// range sends back a rune, need to cast as string to lower
		score += points[strings.ToLower(string(letter))]
	}
	return score
}
