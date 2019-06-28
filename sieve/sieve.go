// Package sieve contains the functions to calculate sieves: eg Sieve of Erostane.
package sieve

// Sieve calculates the prime numbers up to n included
// using the sieve of Erostane.
func Sieve(n int) (s []int) {
	arrayS := make([]bool, n+1)
	s = make([]int, 0, n/2)
	for i := 2; i <= n; i++ {
		if arrayS[i] {
			continue
		}
		for j := i * 2; j <= n; j += i {
			arrayS[j] = true
		}
	}
	for i := 2; i <= n; i++ {
		if !arrayS[i] {
			s = append(s, i)
		}
	}
	return
}
