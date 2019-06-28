// Package clock provide  struct clock with few primitives.
package clock

import "fmt"

// Clock is a struct that contains the hours and minutes of the clock
type Clock struct {
	minutes int
}

// New returns a new clock
func New(h, m int) Clock {
	minutes := h*60 + m
	minutes %= 1440
	if minutes < 0 {
		minutes += 1440
	}

	return Clock{minutes}
}

// String returns the stringed version of the class clock
func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.minutes/60, c.minutes%60)
}

// Add generates a new clock with a positive time difference of addMinutes
func (c Clock) Add(addMinutes int) Clock {
	newTimeMinutes := c.minutes + addMinutes
	return New(0, newTimeMinutes)
}

// Subtract generates a new clock with a positive time difference of subMinutes
func (c Clock) Subtract(subMinutes int) Clock {
	newTimeMinutes := c.minutes - subMinutes
	return New(0, newTimeMinutes)
}
