// Package diffsquares contains a suite to of functions to calculate the sum of squares up to n
// the square of the sum of the number up to n and the function
// that calculate the difference in between the two
package diffsquares

// SquareOfSum calculate the square of the sum of the number 1 to n
func SquareOfSum(n int) int {
	if n < 1 {
		return 0
	}
	t := n * (n + 1) / 2
	return t * t
}

// SumOfSquares calculate the sum of the square of the number 1 to n
func SumOfSquares(n int) int {
	if n < 1 {
		return 0
	}
	return (n * (n + 1) * (2*n + 1)) / 6
}

// Difference calculate the difference between SquareOfSum and SumOfSquares
// on the same input n
func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
