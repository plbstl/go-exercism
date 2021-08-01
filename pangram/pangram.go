// Package panagram provide routines
// for panagram related operations.
package pangram

import "unicode"

// IsPangram checks if a sentence
// uses every letter of the
// alphabet at least once.
func IsPangram(s string) bool {
	var hist [26]int
	for _, r := range s {
		r = unicode.ToLower(r)
		if r >= 'a' && r <= 'z' {
			hist[r-'a']++
		}
	}
	for _, n := range hist {
		if n == 0 {
			return false
		}
	}
	return true
}
