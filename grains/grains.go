// Package grains provides a suite of functions to calculate
// the number of grains on a chess board of the king story.
package grains

import (
	"fmt"
)

// Square calculates the number of grains on the square x.
func Square(x int) (uint64, error) {
	if x < 1 || x > 64 {
		return 0, fmt.Errorf("bad square %d", x)
	}
	return 1 << uint(x-1), nil
}

// Total calculates the total number of grains on the 64 squares board.
func Total() uint64 {
	return 1<<64 - 1
}
