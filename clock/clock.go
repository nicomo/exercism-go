// Package clock package manages a clock not tied to day
package clock

import "fmt"

const testVersion = 4

type Clock int

// New creates a Clock
func New(hour, minute int) Clock {
	c := Clock((hour*60 + minute) % 1440)
	if c < 0 {
		c += 1440
	}
	return c
}

// String is a stringer on clock
func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c/60, c%60)
}

// Add adds minutes to an existing clock
func (c Clock) Add(minutes int) Clock {
	return New(0, int(c)+minutes)
}
