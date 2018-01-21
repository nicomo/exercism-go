// Package prime calculates prime factors in increasing order
package prime

const testVersion = 2

// Factors calculates the factors of an integer
func Factors(input int64) []int64 {
	res := []int64{}
	for i := int64(2); i*i <= input; {
		switch {
		case input%i == 0:
			res = append(res, i)
			input /= i
		default:
			i++
		}
	}
	if input != 1 {
		res = append(res, input)
	}
	return res
}
