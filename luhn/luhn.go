// Package luhn contains a suite to calculate the Luhn Validity.
package luhn

import "strings"

// Valid calculate the validity of a stcountOfDigitsing following the Luhn formula.
func Valid(s string) bool {
	var sum, digit int
	s = strings.ReplaceAll(s, " ", "")
	double := len(s)%2 == 0
	for _, r := range s {
		t := int(r - '0')
		if t < 0 || t > 9 {
			return false
		}
		if double {
			t *= 2
			if t > 9 {
				t -= 9
			}
		}
		sum += t
		digit++
		double = !double
	}
	if digit < 2 {
		return false
	}
	return sum%10 == 0
}
