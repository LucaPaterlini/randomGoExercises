// Package etl provide a suite to improve performance on data structures of the scrabble game
package etl

import "strings"

// Transform convert a val:[]letter data structure to a letter:value map
func Transform(oldMap map[int][]string) map[string]int {
	newMap := map[string]int{}
	for i, arrayLetters := range oldMap {
		for _, letter := range arrayLetters {
			newMap[strings.ToLower(letter)] = i
		}
	}
	return newMap
}
