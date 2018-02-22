// Package prime Given a number n, determine what the nth prime is
package prime

// Nth returns the Nth prime number
func Nth(i int) (int, bool) {
	if i == 0 {
		return 0, false
	}
	// hardcoding a big number is bad. But it works :-)
	primes := Sieve(200000, i)
	return primes[i-1], true
}

// Sieve implement the Sieve of Eratosthenes
// to find all the primes from 2 up to a given number
func Sieve(limit, nth int) []int {

	var r []int
	sieve := make([]bool, limit)
	for i := 2; i < limit; i++ {
		if !sieve[i] {
			r = append(r, i)
			for j := 2 * i; j < limit; j += i {
				sieve[j] = true
			}
		}
		if len(r) == nth {
			break
		}
	}

	return r

}
