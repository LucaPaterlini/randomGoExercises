// Package proverb provide a suite of functions related to proverbs
package proverb

import "fmt"

// Proverb returns the prover adapted with the new rhymes
func Proverb(rhyme []string) (res []string) {
	if len(rhyme) == 0 {
		return
	}
	for i := 0; i < len(rhyme)-1; i++ {
		res = append(res, fmt.Sprintf("For want of a %s the %s was lost.", rhyme[i], rhyme[i+1]))
	}
	res = append(res, fmt.Sprintf("And all for the want of a %s.", rhyme[0]))
	return res
}
