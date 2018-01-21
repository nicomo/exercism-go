// Package triangle loves geometry
package triangle

import (
	"math"
	"sort"
)

const (
	testVersion      = 2
	NaT         Kind = iota
	Equ
	Iso
	Sca
)

type Kind int

// KindFromSides checks if triangle is valid || equilateral || isoscele || scalene.
func KindFromSides(a, b, c float64) Kind {

	f := []float64{a, b, c}
	sort.Float64s(f)

	switch {
	case isNaT(f):
		return NaT
	case isEqu(f):
		return Equ
	case isIso(f):
		return Iso
	default:
		return Sca
	}
}

// isEqu tests equilateral triangles
func isEqu(f []float64) bool {
	return f[0] == f[1] && f[0] == f[2] && f[0] != 0
}

// isIso tests isosceles triangles
func isIso(f []float64) bool {
	return f[0] == f[1] || f[0] == f[2] || f[1] == f[2]
}

// isNaT tests various cases of not a triangle
func isNaT(f []float64) bool {
	// tests triangle inequality
	if f[0]+f[1] < f[2] {
		return true
	}
	// tests not a number and infinity number
	for i := range f {
		if math.IsNaN(f[i]) || math.IsInf(f[i], 0) || f[0] <= 0 {
			return true
		}
	}
	return false
}
