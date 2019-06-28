// Package atbash provide the function Atbash
package atbash

import (
	"bytes"
)

// Atbash code a string in Atbash encoding format
func Atbash(s string) string {
	lS := len(s)
	tmpCh := make([]byte, 0, lS+lS/5)
	for _, x := range []byte(s) {
		switch {
		case x >= 'a' && x <= 'z':
			{
				tmpCh = append(tmpCh, 'z'-x+'a')
			}
		case x >= 'A' && x <= 'Z':
			{
				tmpCh = append(tmpCh, 'z'-x+'A')
			}
		case x >= '0' && x <= '9':
			{
				tmpCh = append(tmpCh, x)
			}
		default:
			continue
		}
		if len(tmpCh)%6 == 5 {
			tmpCh = append(tmpCh, ' ')
		}
	}
	return string(bytes.TrimSuffix(tmpCh, []byte{' '}))
}