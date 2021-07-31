// Package anagram provide implementations for working with anagrams.
// An anagram is a word or phrase that is created by the rearranging
// the letters of another word or phrase.
package anagram

import "strings"

// Detect returns a list of anagrams of the subject from the given candidates.
func Detect(subject string, candidates []string) (anagrams []string) {
	for _, candidate := range candidates {
		if len(subject) != len(candidate) {
			// different lengths.
			continue
		}
		if _ASCIISum(subject) != _ASCIISum(candidate) {
			// different letters.
			continue
		}
		if bitsetOfLetters(subject) != bitsetOfLetters(candidate) {
			// different letters.
			continue
		}
		if strings.EqualFold(subject, candidate) {
			// same word.
			continue
		}
		anagrams = append(anagrams, candidate)
	}
	return
}

func _ASCIISum(s string) (sum int) {
	for _, c := range s {
		c = (c & 0xdf) - 'A'
		sum += int(c)
	}
	return
}

func bitsetOfLetters(s string) (bs rune) {
	for _, c := range s {
		c = (c & 0xdf) - 'A'
		if c > 25 || c < 0 {
			continue
		}
		bs |= 1 << c
	}
	return
}
