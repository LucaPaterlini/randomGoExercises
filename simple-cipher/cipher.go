// Package cipher provide a set of function to code and decode plain english text.
package cipher

// Cipher is the interface of the ciphers minimum requirements Encode and Decode.
type Cipher interface {
	Encode(string) string
	Decode(string) string
}

// Caesar is the implementation of the Cipher interface for the Caesar encryption.
type Caesar struct {
	shift byte
}

// Encode encrypt a string using its Caesar cipher.
func (c Caesar) Encode(s string) string {
	res := make([]byte, 0, len(s))
	for _, x := range []byte(s) {
		switch {
		case x >= 'a' && x <= 'z':
			res = append(res, (x-'a'+c.shift)%26+'a')
		case x >= 'A' && x <= 'Z':
			res = append(res, (x-'A'+c.shift)%26+'a')
		}
	}
	return string(res)
}

// Decode decrypt a string using its Caesar cipher.
func (c Caesar) Decode(s string) string {
	c.shift = 26 - c.shift
	return c.Encode(s)
}

// NewCaesar returns a new struct of Caesar with the shift of 3 letters.
func NewCaesar() Cipher {
	return Caesar{3}
}

// NewShift returns a Cipher for shift encryption with x shifts.
func NewShift(x int) Cipher {
	if x < (-25) || (x > 25) || x == 0 {
		return nil
	}
	return Caesar{byte((26 + x) % 26)}
}

// Vigenere is the implementation of the Cipher interface for the Vigenere encryption.
type Vigenere struct {
	key  []byte
	lKey int
}

// Encode encrypt a string using its Vigenere cipher.
func (v Vigenere) Encode(s string) string {
	res := make([]byte, len(s))
	i := 0
	for _, x := range []byte(s) {
		switch {
		case x >= 'a' && x <= 'z':
			res[i] = (x-'a'+v.key[i%v.lKey])%26 + 'a'
		case x >= 'A' && x <= 'Z':
			res[i] =
				(x-'A'+v.key[i%v.lKey])%26 + 'a'
		default:
			continue
		}
		i++
	}
	return string(res[:i])
}

// Decode decrypt a string using its Vigenere cipher.
func (v Vigenere) Decode(s string) string {
	e := Vigenere{make([]byte, len(v.key)), v.lKey}
	for i, x := range []byte(v.key) {
		e.key[i] = 26 - x
	}
	return e.Encode(s)
}

// NewVigenere returns a Cipher for shift encryption applying the string s.
func NewVigenere(s string) Cipher {
	if len(s) < 3 {
		return nil
	}
	for _, x := range s {
		if x < 'a' || x > 'z' {
			return nil
		}
	}
	sPos := make([]byte, 0, len(s))
	for _, x := range []byte(s) {
		sPos = append(sPos, x-'a')
	}
	return Vigenere{sPos, len(s)}
}
