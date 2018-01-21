// Package raindrops prints strings
// given a prime factorization of a number
package raindrops

import "strconv"

const testVersion = 2

// Convert takes an int, prime factorize it, return string
func Convert(n int) string {

	// case 1
	if n < 2 {
		return strconv.Itoa(n)
	}

	raindrops := ""

	cases := map[int]string{
		3: "Pling",
		5: "Plang",
		7: "Plong",
	}

	// ranging of map is random, we need sorted keys outside of map
	var keys = [3]int{3, 5, 7}

	// is the int divisible by 3, 5 or 7
	for _, k := range keys {
		if n%k == 0 {
			raindrops += cases[k]
		}
	}

	// it is not -> output original int as string
	if raindrops == "" {
		raindrops += strconv.Itoa(n)
	}

	return raindrops
}
