// Package isbn checks validity of isbn-10 string
package isbn

import (
	"errors"
	"math"
)

// IsValidISBN checks validity of isbn
func IsValidISBN(isbn string) bool {
	cleanISBN, err := cleanISBN(isbn)
	if err != nil {
		return false
	}

	return checkLastDigit(cleanISBN)
}

// cleanISBN removes hyphens
// also performs checks for length, and X only in last position
func cleanISBN(isbn string) ([]int, error) {
	res := make([]int, 0)
	for _, r := range isbn {
		if r >= 48 && r <= 57 { // ascii table 0 to 9
			res = append(res, int(r-'0'))
			continue
		}
		if r == 88 { // ascii table X
			res = append(res, 10)
			// X only allowed as a check digit, i.e. last digit
			if len(res) != 10 {
				return nil, errors.New("X only allowed as a check digit")
			}
			continue
		}
	}

	if len(res) != 10 {
		return nil, errors.New("cleaned up isbn isn't of length 10")
	}

	return res, nil
}

// checkLastDigit calculates validity of check digit at last position
func checkLastDigit(cISBN []int) bool {
	var sum int
	for i, v := range cISBN {
		sum += int(math.Abs(float64(i)-10)) * v
	}
	res := sum % 11
	return res == 0
}
