// Package darts simulates a darts game
package darts

import (
	"math"
)

// Score returns points according to where
// the dart lands on the target
func Score(x, y float64) int {
	switch d := distance(x, y); {
	case d > 10:
		return 0
	case d > 5:
		return 1
	case d > 1:
		return 5
	default:
		return 10
	}
}

// distance gets the distance between 2 points.
// i.e. length of the hypotenuse is sq root of (x2 - x1)^2 + (y2 - y1)^2
// one of our points is the center (0.0,0.0)
func distance(x, y float64) float64 {
	first := math.Pow(0.0-x, 2)
	second := math.Pow(0.0-y, 2)
	return math.Sqrt(first + second)
}
