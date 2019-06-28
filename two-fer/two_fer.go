// Package twofer implement a suite of functions relate to
// to two for one
package twofer

import "fmt"

// ShareWith return the 2 fer sentence using as subject the name
// as parameter
func ShareWith(name string) string {
	// check if name is empty and update with you
	if name == "" {
		name = "you"
	}

	return fmt.Sprintf("One for %s, one for me.", name)
}
