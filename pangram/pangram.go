// Package pangram contains the function to check
// if a string is a pangram
package pangram

// IsPangram check if a sentence is a pangram
// on the english alphabet (Ascii letters only)
func IsPangram(s string) bool {
	alpha := make([]bool, 26)
	for _, c := range s {
		q := c - 'a'
		if q < 0 {
			q += 32
		}
		if 0 <= q && q < 26 {
			alpha[q] = true
		}
	}
	for _, b := range alpha {
		if !b {
			return false
		}
	}
	return true
}
