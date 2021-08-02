// Package diffsquares implement routines
// for finding the difference of squares.
package diffsquares

// Difference returns the difference between
// the square of the sum of the first N
// natural numbers and the sum of the squares
// of the first N natural numbers.
//
// 	if n = 10
// 	(1 + 2 + ... + 10)² - (1² + 2² + ... + 10²)
func Difference(n float64) float64 {
	sqOfSum := n * (n + 1) / 2
	sumOfSq := n * (2*n + 1) * (n + 1) / 6
	return sqOfSum*sqOfSum - sumOfSq
}
