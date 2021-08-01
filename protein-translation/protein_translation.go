// Package protein provide routines for operations
// carried out during protein synthesis.
package protein

// Think of this package as a ribosome.

import "errors"

// ErrStop occurs when a 'STOP' codon is reached.
var ErrStop = errors.New("stop sequence")

// ErrInvalidBase occurs when the passed in codon
// has invalid nucleotides or nucleotide sequence.
var ErrInvalidBase = errors.New("invalid base")

// ErrInvalidLength occurs when the passed in codon has
// a length not equal to three (3).
var ErrInvalidLength = errors.New("invalid codon length")

// FromRNA translates RNA sequences into proteins.
//
// RNA can be broken into three nucleotide sequences called codons, and then translated to a polypeptide like so:
//
// RNA: `"AUGUUUUCU"` => translates to
//
// Codons: `"AUG", "UUU", "UCU"`
// => which become a polypeptide with the following sequence =>
//
// Protein: `"Methionine", "Phenylalanine", "Serine"`
// .
func FromRNA(rna string) (polypeptide []string, err error) {
	var codon string
	var protein string
	for i := 0; i < len(rna); i += 3 {
		codon = rna[i : i+3]
		protein, err = FromCodon(codon)
		if err != nil {
			if err == ErrStop {
				return polypeptide, nil
			}
			return
		}
		polypeptide = append(polypeptide, protein)
	}
	return
}

// Translations contain a map of codons
// and its resulting Amino Acids.
var Translations map[string]string

var translations = map[string]string{
	"AUG": "Methionine",
	"UUU": "Phenylalanine",
	"UUC": "Phenylalanine",
	"UUA": "Leucine",
	"UUG": "Leucine",
	"UCU": "Serine",
	"UCC": "Serine",
	"UCA": "Serine",
	"UCG": "Serine",
	"UAU": "Tyrosine",
	"UAC": "Tyrosine",
	"UGU": "Cysteine",
	"UGC": "Cysteine",
	"UGG": "Tryptophan",
}

// FromCodon returns the relevant protein associated with
// the codon.
//
// A codon is a sequence of three adjacent nucleotides which
// encode for a specific amino acid during protein synthesis.
func FromCodon(codon string) (protein string, err error) {
	if len(codon) != 3 {
		return "", ErrInvalidLength
	}
	if codon == "UAA" || codon == "UAG" || codon == "UGA" {
		return "", ErrStop
	}
	protein, ok := translations[codon]
	if !ok {
		return "", ErrInvalidBase
	}
	return
}
