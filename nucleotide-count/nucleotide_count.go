// Package dna provide routines for DNA specific operations.
package dna

import (
	"errors"
	"strconv"
)

// Histogram is a mapping from a nucleotide to its count.
type Histogram map[byte]int

// DNA is a case-sensitive list of nucleotides.
type DNA string

// Counts generates a histogram of valid nucleotides.
// Returns an error if DNA contains an invalid nucleotide.
func (d DNA) Counts() (Histogram, error) {
	h := Histogram{'A': 0, 'C': 0, 'G': 0, 'T': 0}
	for i := range d {
		switch d[i] {
		case 'A':
			h['A'] += 1
		case 'C':
			h['C'] += 1
		case 'G':
			h['G'] += 1
		case 'T':
			h['T'] += 1
		default:
			return nil, errors.New("dna: invalid nucleotide `" + string(d[i]) + "` at index " + strconv.Itoa(i))
		}
	}
	return h, nil
}
