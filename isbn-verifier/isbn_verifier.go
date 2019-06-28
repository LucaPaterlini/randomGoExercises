// Package isbn contains a suite of functions about ISBN.
package isbn
// IsValidISBN checks if an isbn 10 is valid.
func IsValidISBN(s string) bool {
	// check len s
	lastPos := len(s) - 1
	if lastPos < 0 {
		return false
	}
	c, sum := 10, 0
	for _, x := range s[:lastPos] {
		if '0' <= x && x <= '9' {
			sum += c * (int(x) - 48)
			c--
		}
	}
	// check the last character
	if c > 1 {
		return false
	}
	if s[lastPos] == 'X' {
		sum += 10
	} else {
		sum += int(s[lastPos]) - 48
	}

	return sum%11 == 0
}
