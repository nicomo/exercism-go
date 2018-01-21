// Package gigasecond tells you the date you'll be or have been a gigasecond old - coool
package gigasecond

import "time"

const testVersion = 4

// AddGigasecond uses a type from the Go standard library.
func AddGigasecond(t time.Time) time.Time {
	return t.Add(1e9 * time.Second)
}
