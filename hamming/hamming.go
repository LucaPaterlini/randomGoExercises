// Package hamming implements the Distance function between 2 strings.
package hamming

import (
	"errors"
)

// Distance calculate the distance between two string, return error
// in case of the lengths of the 2 strings are different.
func Distance(a, b string) (int, error) {
	bR := []rune(b)
	aR := []rune(a)
	if len(aR) != len(bR) {
		return 0, errors.New("input strings were different lengths, need to be the same")
	}
	c := 0
	for i, x := range aR {
		if bR[i] != x {
			c++
		}
	}
	return c, nil
}
