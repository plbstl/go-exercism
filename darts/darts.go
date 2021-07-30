// Package darts provide functions
// for implementing a Darts game.
package darts

const (
	outer  = 10 * 10 // squared radius of the outer circle.
	middle = 5 * 5   // squared radius of the middle circle.
	inner  = 1 * 1   // squared radius of the inner circle.
)

// Score returns the earned points in
// a single toss of a Darts game.
func Score(x, y float64) int {
	// position of the dart.
	//
	// 	r² = (x-h)² + (y-h)²; // but h = 0, k = 0.
	var position = x*x + y*y

	switch {
	case position <= inner:
		return 10
	case position <= middle:
		return 5
	case position <= outer:
		return 1
	default:
		return 0
	}
}
