// Package collatzconjecture implements the mathematical
// conjecture introduced by Lothar Collatz.
package collatzconjecture

import "errors"

// CollatzConjecture returns the number of steps
// taken for a positive integer to reach 1 when
// implementing collatz conjecture.
func CollatzConjecture(n int) (steps int, err error) {
	if n < 1 {
		return -1, errors.New("zero and negative numbers are unsupported")
	}
	for n != 1 {
		steps++
		if n&1 == 0 {
			n /= 2
		} else {
			n = 3*n + 1
		}
	}
	return
}
