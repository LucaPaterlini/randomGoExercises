// Package perfect provide a function to check if a positive number is perfect.
package perfect

import (
	"errors"
	"math"
)

// Classification is the type of classification that can be attributed to a number.
type Classification int

// ClassificationDeficient means the sum of divisor is smaller the the number itself.
const ClassificationDeficient = -1

// ClassificationPerfect means the sum of divisor is the same the the number itself and the number is perfect.
const ClassificationPerfect = 0

// ClassificationAbundant means the sum of divisor is greater the the number itself.
const ClassificationAbundant = 1

// ErrOnlyPositive is an error that signal that only poisitive numbers are allowed.
var ErrOnlyPositive = errors.New("only positive numbers are allowed")

// Classify returns the classification of the number and in case its negative an error.
func Classify(x int64) (Classification, error) {
	if x < 1 {
		return 0, ErrOnlyPositive
	}
	var sumFactors int64 = 1
	for i := int64(2); i <= int64(math.Sqrt(float64(x))); i++ {
		if x%i == 0 {
			sumFactors += i + x/i
		}
	}
	// special case perfect squares
	if math.Sqrt(float64(x)) == math.Round(math.Sqrt(float64(x))) {
		sumFactors -= int64(math.Sqrt(float64(x)))
	}
	switch {
	case sumFactors > x:
		return ClassificationAbundant, nil
	case sumFactors < x:
		return ClassificationDeficient, nil
	}
	return ClassificationPerfect, nil
}
