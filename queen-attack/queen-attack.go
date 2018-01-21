//package queenattack
package queenattack

import (
	"errors"
	"math"
	"strings"
)

// CanQueenAttack takes the coordinates of 2 queens on a chess board
// and determines if they can attack one another
func CanQueenAttack(w string, b string) (bool, error) {

	// same square
	if w == b {
		err := errors.New("same square")
		return false, err
	}

	// off the board
	wb := w + b
	for _, v := range wb {
		if IsOffBoard(v) {
			err := errors.New("off board")
			return false, err
		}
	}

	for _, v := range w {
		// same file or rank: can attack
		if strings.ContainsRune(b, v) {
			return true, nil
		}
	}

	// on diagonal line?
	if IsOnDiagonal(w, b) {
		return true, nil
	}

	return false, nil
}

// IsOffBoard tests if we're still on the chess board
func IsOffBoard(v rune) bool {
	switch v {
	case
		'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', '1', '2', '3', '4', '5', '6', '7', '8':
		return false
	}
	return true
}

// IsOnDiagonal checks if 2 queens can attack each other using diagonals
// the gap between the files has to be equal to the gap between the ranks
// e.g. a -> b == 1 -> 2
// or a -> d == 1 -> 3 etc
func IsOnDiagonal(w string, b string) bool {
	wfile, wrank := int(w[0]), int(w[1])
	bfile, brank := int(b[0]), int(b[1])
	if math.Abs(float64(wfile-bfile)) == math.Abs(float64(wrank-brank)) {
		return true
	}
	return false
}
