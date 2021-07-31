// Package bob helps to communicate with
// a lackadaisical teenager named Bob.
package bob

import (
	"strings"
	"unicode"
)

// Hey attempts to start a conversation with
// Bob and returns his response.
func Hey(remark string) string {
	remark = strings.TrimSpace(remark)
	if remark == "" {
		return "Fine. Be that way!"
	}
	question := remark[len(remark)-1] == '?'
	yelling := isYelledAt(remark)
	if yelling && question {
		return "Calm down, I know what I'm doing!"
	}
	if yelling {
		return "Whoa, chill out!"
	}
	if question {
		return "Sure."
	}
	return "Whatever."
}

func isYelledAt(remark string) bool {
	var isUppercase bool
	for _, r := range remark {
		if !unicode.IsLetter(r) {
			continue
		}
		if !unicode.IsUpper(r) {
			return false
		}
		isUppercase = true
	}
	return isUppercase
}
