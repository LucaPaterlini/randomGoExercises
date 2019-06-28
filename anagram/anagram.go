// Package anagram contain a function that filters an array of strings
// checking those that are anagrams
// NOTE: this solution is optimized for english letters only
package anagram

import "strings"

const NENGLISHLETTERS = 26

// countLetters update cArray with the number of the occurrences of each letter in word
func countLetters(word string, cArray *[]int, step int) {
	for _, c := range word {
		q := c - 'a'
		if q < 0 {
			q += 32
		}
		if 0 <= q && q < NENGLISHLETTERS {
			(*cArray)[q] += step
		}
	}
}

// none check if the array of int passaed as parameters have only 0 as values
func noneInt(a []int) bool {
	for _, x := range a {
		if x != 0 {
			return false
		}
	}
	return true
}

// Detect check if contain a function that filters
// an array of strings checking those that are anagrams
func Detect(aS string, tS []string) []string {
	// preparing the array with the frequencies each letter in the test word
	alpha := make([]int, NENGLISHLETTERS)
	countLetters(aS, &alpha, 1)
	tmp := make([]int, NENGLISHLETTERS)
	resS := []string{}
	for _, word := range tS {
		if strings.ToLower(word) == strings.ToLower(aS) {
			continue
		}
		copy(tmp, alpha)
		countLetters(word, &tmp, -1)
		if noneInt(tmp) {
			resS = append(resS, word)
		}
	}
	return resS
}
