//Package isogram checks if a given word contains more than one character.
package isogram

import (
	"unicode"
)

//IsIsogram loops trough a word and check if contains more than one character
//hypes and spaces don't match the count.
func IsIsogram(word string) bool {
	var counter = map[rune]bool{}
	for _, letter := range word {
		letter = unicode.ToLower(letter)
		if unicode.IsLetter(letter) && counter[letter] == true {
			return false
		}
		counter[letter] = true
	}
	return true
}