// Package dna contains the Dna struct end methods
package dna

import (
	"fmt"
)

// Histogram is a mapping from nucleotide to its count in given DNA.
// Choose a suitable data type.
type Histogram map[rune]int

// DNA is a list of nucleotides. Choose a suitable data type.
type DNA string

// Counts generates a histogram of valid nucleotides in the given DNA.
// Returns an error if d contains an invalid nucleotide.
func (d DNA) Counts() (h Histogram, err error) {
	h = map[rune]int{'A': 0, 'C': 0, 'G': 0, 'T': 0}
	for _, x := range d {
		if _, ok := h[x]; !ok {
			return h, fmt.Errorf("invalid nucleotides %v", x)
		}
		h[x]++
	}
	return
}
