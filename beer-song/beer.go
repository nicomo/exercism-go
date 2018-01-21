package beer

import (
	"errors"
	"strconv"
)

const testVersion = 1

// Verse generates a single verse
func Verse(i int) (string, error) {
	if i < 0 || i > 99 {
		return "", errors.New("index out of bound")
	}
	var s string
	switch i {
	case 2:
		s += strconv.Itoa(i) + " bottles of beer on the wall, " +
			strconv.Itoa(i) + " bottles of beer.\nTake one down and pass it around, " +
			strconv.Itoa(i-1) + " bottle of beer on the wall.\n"
	case 1:
		s += strconv.Itoa(i) + " bottle of beer on the wall, " +
			strconv.Itoa(i) + " bottle of beer.\nTake it down and pass it around, no more bottles of beer on the wall.\n"
	case 0:
		s += "No more bottles of beer on the wall, no more bottles of beer.\nGo to the store and buy some more, 99 bottles of beer on the wall.\n"
	default:
		s += strconv.Itoa(i) + " bottles of beer on the wall, " +
			strconv.Itoa(i) + " bottles of beer.\nTake one down and pass it around, " +
			strconv.Itoa(i-1) + " bottles of beer on the wall.\n"
	}
	return s, nil
}

// Verses generates a bunch of verses
func Verses(upperBound, lowerBound int) (string, error) {
	if upperBound > 99 || lowerBound < 0 || upperBound < lowerBound {
		return "", errors.New("wrong upper or lower bound")
	}
	var s string
	for i := upperBound; i >= lowerBound; i-- {
		v, err := Verse(i)
		if err != nil {
			return "", err
		}
		s += v + "\n"
	}

	return s, nil
}

// Song generates the whole song
func Song() string {
	s, _ := Verses(99, 0)
	return s
}
