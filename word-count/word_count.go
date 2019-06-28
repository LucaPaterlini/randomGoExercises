// Package wordcount provides a function to count the occupancies of a word in a string
package wordcount

import (
	"regexp"
	"strings"
)

type Frequency map[string]int

var re = regexp.MustCompile("[a-z0-9]+('[a-z0-9])*")
// WordCount return the frequencies of a word in a string
func WordCount(s string) Frequency {
	s = strings.ToLower(s)
	res := Frequency{}
	words := re.FindAllString(s, -1)
	for _, c := range words {
		res[c]++
	}
	return res
}
