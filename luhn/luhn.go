// Package Luhn calculates the check-digit of a number
package luhn

import (
	"strconv"
	"strings"
)

// Valide tests if the given string is a valid luhn string
// with a correct check digit at the end
func Valid(in string) bool {
	in = strings.Replace(in, " ", "", -1)
	if len(in) == 0 {
		return false
	}
	inLast, _ := strconv.Atoi(in[len(in)-1:])
	inPartialReversed := Reverse(in[:len(in)-1])
	luhnSum := getCheckDigit(inPartialReversed)

	if ((luhnSum * 9) % 10) == inLast {
		return true
	}
	return false
}

// AddCheck: we are given a partial number and add the check digit at the end
func AddCheck(in string) (out string) {
	inReversed := Reverse(strings.Replace(in, " ", "", -1))
	cd := getCheckDigit(inReversed)
	cd = (cd * 9) % 10
	out = in + strconv.Itoa(cd)
	return out
}

// getCheckDigit: given a number, what is the correct luhn check digit
func getCheckDigit(in string) int {
	sum := 0
	for i := 0; i < len(in); i++ {
		n, _ := strconv.Atoi(string(in[i]))
		if i%2 == 0 {
			if n*2 > 9 {
				sum += 2*n - 9
			} else {
				sum += 2 * n
			}
		} else {
			sum += n
		}
	}
	return sum
}

// Reverse reverses the string
func Reverse(s string) (result string) {
	for _, v := range s {
		result = string(v) + result
	}
	return
}
