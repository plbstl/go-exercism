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
func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}

// SquareOfSum returns the square of the
// sum of the first N natural numbers.
//
// 	if n = 10
// 	(1 + 2 + ... + 10)² = 55² = 3025
func SquareOfSum(n int) int {
	var sqOfSum int
	for i := 1; i <= n; i++ {
		sqOfSum += i
	}
	return sqOfSum * sqOfSum
}

// SumOfSquares returns the sum of the
// squares of the first N natural numbers.
//
// 	if n = 10
// 	1² + 2² + ... + 10² = 385
func SumOfSquares(n int) int {
	var sumOfSqu int
	for i := 1; i <= n; i++ {
		sumOfSqu += i * i
	}
	return sumOfSqu
}
