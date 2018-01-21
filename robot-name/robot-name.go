// Package robotname generates unique random names for robots
package robotname

import (
	"math/rand"
	"time"
)

const (
	letters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	numbers = "0123456789"
)

type Robot struct {
	name string
}

// Name returns a robot's name if it's already set
// and creates a random one otherwise
func (r *Robot) Name() string {

	if r.name != "" {
		return r.name
	} else {
		rand.Seed(time.Now().UTC().UnixNano())

		n := make([]byte, 5)

		for i := 0; i < 2; i++ {
			n[i] = letters[rand.Intn(len(letters))]
		}

		for i := 2; i < 5; i++ {
			n[i] = numbers[rand.Intn(len(numbers))]
		}

		r.name = string(n)

		return string(n)
	}

}

// Reset resets robot's name to nil value
func (r *Robot) Reset() {
	r.name = ""
}
