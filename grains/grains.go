// Package grains counts grains of rice on the 64 squares of a chess board
package grains

import (
	"errors"
	"math"
)

// Square calculates how manu grains of rice are on a given square of the board
func Square(n int) (uint64, error) {
	if n > 64 || n <= 0 {
		err := errors.New("off board")
		return 0, err
	}
	grains := math.Pow(2, float64(n-1))
	return uint64(grains), nil
}

// Total computes the total number of rice grains on the chessboard
func Total() uint64 {
	var total uint64
	for i := 1; i <= 64; i++ {
		sq, _ := Square(i)
		total += sq
	}
	return total
}
