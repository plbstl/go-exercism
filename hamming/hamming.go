// Package hamming provides functions
// for error estimation.
package hamming

import "errors"

// Distance receives case-sensitive strands
// of DNA and calculates the hamming distance.
func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, errors.New("strands must have equal length")
	}

	n := 0
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			n++
		}
	}

	return n, nil
}

// @todo:
// 	ln 16: --- if a[i] != b[i] { ---
// byte comparisons are safe when dealing with ASCII characters.
// As soon as utf-8 characters are involved, it complicates matters.
// To explore, add an extra test with special characters to see what happens.
