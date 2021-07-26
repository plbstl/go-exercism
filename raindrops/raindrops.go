// Package raindrops implements basic routines
// for working with raindrop sounds.
package raindrops

import "strconv"

// Convert converts a number into raindrop sounds
// corresponding to certain potential factors.
func Convert(n int) string {
	drops := ""

	if n%3 == 0 {
		drops += "Pling"
	}
	if n%5 == 0 {
		drops += "Plang"
	}
	if n%7 == 0 {
		drops += "Plong"
	}

	if drops == "" {
		return strconv.Itoa(n)
	}
	return drops
}
