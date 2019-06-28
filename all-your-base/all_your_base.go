// Package allyourbase provide a function for base conversion.
package allyourbase

import (
	"errors"
)

// ConvertToBase gets in input a base, an array of int and returns an array of
// ints int in a new base bOut.
func ConvertToBase(bIn int, dig []int, bOut int) ([]int, error) {
	var b10 int
	var res []int
	// check wrong inputs
	if bIn < 2 {
		return []int{}, errors.New("input base must be >= 2")
	}
	if bOut < 2 {
		return []int{}, errors.New("output base must be >= 2")
	}
	// convert to 10
	for _, d := range dig {
		if d < 0 || d >= bIn {
			return []int{}, errors.New("all digits must satisfy 0 <= d < input base")
		}
		b10 *= bIn
		b10 += d
	}
	for b10 != 0 {
		res = append([]int{b10 % bOut}, res...)
		b10 /= bOut
	}
	if len(res) == 0 {
		res = []int{0}
	}
	return res, nil
}
