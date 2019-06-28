// Package reverse contain the function Reverse that revert a string.
package reverse

// Reverse the order of the char the string s
func Reverse(s string) string {
	sR := []rune(s)
	lsR := len(sR)
	for i := 0; i < (lsR/2 + lsR%2); i++ {
		sR[i], sR[lsR-i-1] = sR[lsR-i-1], sR[i]
	}
	return string(sR)
}