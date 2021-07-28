// Package proverb provides functionality for
// creating brief popular epigrams or maxims.
package proverb

// Proverb generates a relevant proverb
// based on the given rhyme.
func Proverb(rhyme []string) []string {
	size := len(rhyme)
	if size == 0 {
		return []string{}
	}
	proverb := make([]string, size)
	for i := 0; i < size-1; i++ {
		proverb[i] = "For want of a " + rhyme[i] + " the " + rhyme[i+1] + " was lost."
	}
	proverb[size-1] = "And all for the want of a " + rhyme[0] + "."
	return proverb
}
