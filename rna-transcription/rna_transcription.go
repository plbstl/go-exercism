// Package strand provides simple functions
// for DNA and RNA transcriptions.
package strand

import (
	"errors"
	"strconv"
	"strings"
)

// ToRNA returns the RNA complement of a DNA sequence.
//
// The four nucleotides found in DNA are adenine (A), cytosine (C), guanine (G) and thymine (T).
//
// The four nucleotides found in RNA are adenine (A), cytosine (C), guanine (G) and uracil (U).
func ToRNA(dna string) (string, error) {
	var s strings.Builder
	s.Grow(len(dna))
	for i, v := range dna {
		switch string(v) {
		case "G":
			s.WriteString("C")
		case "C":
			s.WriteString("G")
		case "T":
			s.WriteString("A")
		case "A":
			s.WriteString("U")
		default:
			return "", errors.New("strand: invalid nucleotide at index " + strconv.Itoa(i))
		}
	}
	return s.String(), nil
}
