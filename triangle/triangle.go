// Package triangle provides basic constants and
// functions for working with 3 sided polygons.
package triangle

import "math"

// Kind represents the type of triangle.
type Kind int

const (
	NaT = iota // not a triangle
	Equ        // equilateral
	Iso        // isosceles
	Sca        // scalene
	Deg        // degenerate
)

// KindFromSides determines if a triangle is
// equilateral, isosceles, scalene, or degenerate.
func KindFromSides(a, b, c float64) Kind {
	// all sides have to be of length > 0.
	if a <= 0 || b <= 0 || c <= 0 {
		return NaT
	}
	if math.IsNaN(a+b+c) || math.IsInf(a+b+c, 0) {
		return NaT
	}
	// the sum of the lengths of any two sides
	// must be greater than or equal to the
	// length of the third side.
	if a+b < c || a+c < b || b+c < a {
		return NaT
	}

	// the sum of the lengths of two sides equals
	// that of the third [degenerate triangle].
	if a+b == c || a+c == b || b+c == a {
		return Deg
	}

	switch {
	case a == b && b == c:
		return Equ
	case a == b || b == c || a == c:
		return Iso
	default:
		return Sca
	}
}
