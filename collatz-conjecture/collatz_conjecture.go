// Package collatzconjecture containst the function
// to calculate the number of step to reach 1 starting
// from n
package collatzconjecture

import (
	"github.com/pkg/errors"
)

// CollatzConjecture calculate the number of steps
// required to reach 1 starting from n
func CollatzConjecture(n int) (int, error) {
	if n < 1 {
		return -1, errors.New("Invalid Input")
	}
	c := 0
	for n != 1 {
		c += 1
		if (n & 1) == 0 {
			n >>= 1
		} else {
			n = 3*n + 1
		}
	}
	return c, nil
}
