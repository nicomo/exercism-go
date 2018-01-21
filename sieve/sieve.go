// Package sieve uses the Sieve of Eratosthenes
// to find all the primes from 2 up to a given number
package sieve

const testVersion = 1

// Sieve implement the Sieve of Eratosthenes
func Sieve(limit int) []int {

	var r []int
	sieve := make([]bool, limit)
	for i := 2; i < limit; i++ {
		if !sieve[i] {
			r = append(r, i)
			for j := 2 * i; j < limit; j += i {
				sieve[j] = true
			}
		}
	}

	return r

}
