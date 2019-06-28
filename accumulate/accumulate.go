// Package accumulate provide a function to apply map on an array of strings
package accumulate

// Accumulate apply a function on an array strings
func Accumulate(aS []string, f func(string) string) []string {
	for i, x := range aS {
		aS[i] = f(x)
	}
	return aS
}
