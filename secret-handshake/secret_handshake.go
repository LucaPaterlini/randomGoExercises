// Package secret contains a suite of functions that calculates secrets.
package secret

var commands = []string{
	"wink",
	"double blink",
	"close your eyes",
	"jump",
}

func reverse(s []string) []string {
	ls := len(s)
	for i := 0; i < ls/2; i++ {
		s[ls-i-1], s[i] = s[i], s[ls-i-1]
	}
	return s
}

// Handshake returns an array of strings base on the uint in input.
func Handshake(x uint) []string {
	var tmpRes []string
	for i := 0; i < len(commands) && x != 0; i++ {
		if x%2 == 1 {
			tmpRes = append(tmpRes, commands[i])
		}
		x /= 2
	}
	if x > 0 {
		tmpRes = reverse(tmpRes)
	}
	return tmpRes
}
