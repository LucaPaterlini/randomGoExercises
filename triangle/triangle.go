// Package triangle provide a suite that calculate functions related to triangles
package triangle

import "math"

type Kind int

const (
	NaT = iota // not a triangle
	Equ        // equilateral
	Iso        // isosceles
	Sca        // scalene
)

// KindFromSides returns the type for the triangle witch sides
// has been provided as input
func KindFromSides(a, b, c float64) (k Kind) {
	if a+b < c || b+c < a || c+a < b || c <= 0 || b <= 0 || a <= 0 ||
		math.IsNaN(a+b+c) || math.IsInf(a+b+c, 0) {
		return NaT
	}
	if a == b && b == c {
		return Equ
	}
	if a == b || b == c || c == a {
		return Iso
	}
	return Sca
}
