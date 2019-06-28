// Package scale contains as suite targeted to scales composition
package scale

import (
	"strings"
)

var (
	sharpScale   = []string{"A", "A#", "B", "C", "C#", "D", "D#", "E", "F", "F#", "G", "G#"}
	sharpStarter = []string{"G", "D", "A", "E", "B", "F#", "e", "b", "f#", "c#", "g#", "d#", "C", "a"}
	flatScale    = []string{"Ab", "A", "Bb", "B", "C", "Db", "D", "Eb", "E", "F", "Gb", "G"}
	markerValue  = map[rune]int{'m': 1, 'M': 2, 'A': 3}
)

// scaleDetect detects witch scale is in use evaluating the first tonic
func scaleDetect(tonic string) *[]string {
	for _, note := range sharpStarter {
		if tonic == note {
			return &sharpScale
		}
	}
	return &flatScale
}

// findStringIndex linear search on an unsorted string array for the first
// occurrence of the string x
func findStringIndex(arr []string, x string) int {
	for i, item := range arr {
		if item == strings.ToUpper(x[:1])+x[1:] {
			return i
		}
	}
	return -1
}

// Scale returns a scale appropriate for the tonic and interval
func Scale(tonic, interval string) []string {
	scale := *scaleDetect(tonic)
	lScale := len(scale)
	index := findStringIndex(scale, tonic)
	shiftedScale := append(scale[index:], scale[:index]...)
	if interval != "" {
		var tmp []string
		i := 0
		for _, c := range interval {
			tmp = append(tmp, shiftedScale[i])
			i = (i + markerValue[c]) % lScale
		}
		shiftedScale = tmp
	}
	return shiftedScale
}
