// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package acronym should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package acronym

import (
	"strings"
)

// Abbreviate returns the acronym of each sentence.
func Abbreviate(s string) string {
	replacer := strings.NewReplacer("'", "", "-", " ","_","")
	s = replacer.Replace(s)
	res := ""
	for _,c := range strings.Fields(s) {
		res +=string(c[0])
	}
	return  strings.ToUpper(res)
}

