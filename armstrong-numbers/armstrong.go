package armstrong

import (
	"math"
)

// IsNumber tests if an integer
// is an Armstrong number
func IsNumber(n int) bool {
	digits := nParse(n)
	var armNumb int
	// digits are in reverse order
	for i := len(digits) - 1; i >= 0; i-- {
		armNumb += int(math.Pow(float64(digits[i]), float64(len(digits))))
	}
	if n != armNumb {
		return false
	}
	return true
}

// nParse stores digits from a number in a slice
func nParse(i int) []int {
	digits := []int{}
	for i > 0 {
		remainer := i % 10
		digits = append(digits, remainer)
		i = i / 10
	}
	return digits
}
