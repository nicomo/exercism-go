// Package space Given an age in seconds, calculate how old someone would be on other planets
package space

// Planet is the name of the given planet
type Planet string

const earthYear = 31557600

var multipliers = map[Planet]float64{
	"Earth":   1.0,
	"Mercury": 0.2408467,
	"Venus":   0.61519726,
	"Mars":    1.8808158,
	"Jupiter": 11.862615,
	"Saturn":  29.447498,
	"Uranus":  84.016846,
	"Neptune": 164.79132,
}

// Age returns the name on any planet in our system
func Age(s float64, p Planet) float64 {
	return (s / earthYear) / multipliers[p]
}
