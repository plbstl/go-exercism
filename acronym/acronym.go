// Package acronym provides basic functions for
// combining the initial letters of a multipart name.
package acronym

import "strings"

// Abbreviate returns an acronym.
func Abbreviate(s string) string {
	s = strings.Replace(s, string("-"), " ", -1)
	s = strings.Replace(s, string("_"), " ", -1)
	var out []byte
	for _, word := range strings.Split(s, " ") {
		if len(word) > 0 {
			out = append(out, word[0])
		}
	}
	return strings.ToUpper(string(out))
}
