// Package gigasecond implements a function to add 1 billion seconds to a date
package gigasecond

import (
	"time"
)

// AddGigasecond add 1 billion seconds to a date
func AddGigasecond(t time.Time) time.Time {
	return t.Add(1e9*time.Second)
}
