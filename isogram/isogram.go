// Package isogram implements basic
// routines for isogram related operations.
package isogram

import (
	"strings"
	"unicode"
)

// IsIsogram tests to see if
// a string is an isogram.
func IsIsogram(s string) bool {
	s = strings.ToLower(s)
	for _, char := range s {
		if !unicode.IsLetter(char) {
			continue
		}
		if charCount := strings.Count(s, string(char)); charCount > 1 {
			return false
		}
	}
	return true
}
