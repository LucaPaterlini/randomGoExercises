// Package prime contains a suite of functions related to prime number calculation
package prime

import (
	"math"
)

// is Prime check if a number is prime
func isPrime(x int) bool {
	if x%2==0 {
		return x==2
	}
	m := 1+int(math.Sqrt(float64(x)))
	for i := 3; i < m ; i+=2 {
		if x%i == 0 {
			return false
		}
	}
	return true
}

// Nth returns the n-th prime number
func Nth(n int) (int, bool) {
	if n < 1 {
		return 0, false
	}
	i := 1
	for n > 0 {
		i++
		if isPrime(i) {
			n--
		}
	}
	return i, true
}
