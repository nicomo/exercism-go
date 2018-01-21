// Package lsproduct: given a string of digits,
// can calculate the largest product for a contiguous substring
// of digits of length n.
package lsproduct

import (
	"errors"
	"strconv"
)

const testVersion = 3

// LargestSeriesProduct given a string of digits,
// can calculate the largest product for a contiguous substring
// of digits of length n.
func LargestSeriesProduct(digits string, span int) (int, error) {
	if span < 0 || span > len(digits) {
		err := errors.New("invalid span or empty digits string")
		return -1, err
	}

	var sum int
	for i := 0; i <= len(digits)-span; i++ {
		counter := 1
		for j := 0; j < span; j++ {
			n, err := strconv.Atoi(string(digits[j+i]))
			if err != nil {
				return 0, err
			}
			counter *= n
		}
		if counter > sum {
			sum = counter
		}

	}

	return sum, nil
}
