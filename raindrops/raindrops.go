// Package raindrops provide a suite of functions to convert to Raindrops language.
package raindrops

import (
	"strconv"
)

type factorRain struct {
	divisor int
	sound   string
}

var coef = []factorRain{
	{3, "Pling"},
	{5, "Plang"},
	{7, "Plong"},
}

// Convert convert a natural number into its Raindrops conversion.
func Convert(x int) (res string) {
	for _, record := range coef {
		if x%record.divisor == 0 {
			res += record.sound
		}
	}
	if res == "" {
		res = strconv.Itoa(x)
	}
	return
}
