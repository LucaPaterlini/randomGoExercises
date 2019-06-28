// Package encode contains functions to change a string encoding
// NOTE: this works only for ascii letters
package encode

import (
	"strconv"
	"strings"
)

// RunLengthEncode returns a RLE conversion of the string in input
func RunLengthEncode(s string) string {
	if len(s) == 0 {
		return s
	}
	t := int32(s[0])
	c := 1
	res := ""
	for _, ch := range s[1:] {
		if ch != t {
			if c > 1 {
				res += strconv.Itoa(c)
			}
			res += string(t)
			t = ch
			c = 1
		} else {
			c += 1
		}
	}
	if c > 1 {
		res += strconv.Itoa(c)
	}
	res += string(t)
	return res
}

// RunLengthDecode returns a decoded string of the RLE in input
func RunLengthDecode(s string) string {
	res := ""
	t := 0
	for _, c := range s {
		if '0' <= c && c <= '9' {
			t = t*10 + (int(c) - 48)
		} else {
			if t == 0 {
				t++
			}
			res += strings.Repeat(string(c), t)
			t = 0
		}
	}
	return res
}
