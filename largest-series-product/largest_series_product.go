// Package lsproduct implements the function that calculate the largest product
// of the digits of substring of length span
package lsproduct

import (
	"github.com/pkg/errors"
)

// LargestSeriesProduct calculate the largest product for a contiguous substring of digits of length n.
func LargestSeriesProduct(digits string, span int) (maxProduct int, err error) {
	lDigits := len(digits)
	// check the boundaries
	if span == 0 {
		maxProduct = 1
		return
	}
	if span < 0 || span > lDigits {
		return -1, errors.New("span value out of range")
	}
	//// checking for empty string
	product, lenSub, maxProduct := 1, 0, 0
	for i := 0; i < len(digits); i++ {
		// new Digit
		newDigit := int(digits[i]) - 48
		if newDigit == 0 {
			lenSub, product = 0, 1
			continue
		}
		if lenSub < span {
			product *= newDigit
			lenSub++
		}
		if lenSub == span {
			if maxProduct < product {
				maxProduct = product
			}
			product /= int(digits[i-span+1]) - 48
			lenSub--
		}
	}
	return
}
