// Package rotationalcipher provides encryption functions by rotation.
package rotationalcipher

// RotationalCipher convert a string by rotating of rot position
// on english alphabet, BE CAREFUL IT WORKS ONLY FOR ENG ALPHABET.
func RotationalCipher(s string, rot int) string {
	rotI32 := int32(rot)
	resB := []byte(s)
	for i, x := range s {
		switch {
		case x >= 'a' && x <= 'z':
			resB[i] = byte(x-'a'+rotI32)%26 + 'a'
		case x >= 'A' && x <= 'Z':
			resB[i] = byte(x-'A'+rotI32)%26 + 'A'
		}
	}
	return string(resB)
}
