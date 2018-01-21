// Package hamming calcultes the Hamming distance between 2 strands of DNA
package hamming

import "errors"

const testVersion = 4

// Distance calculates distance between 2 DNA strands
func Distance(a, b string) (int, error) {

	err := errors.New("disallow, diff length")
	var counter int
	if len(a) != len(b) {
		return -1, err
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			counter++
		}
	}
	return counter, nil
}
