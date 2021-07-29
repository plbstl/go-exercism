// Package etl provide routines for
// Extract-Transform-Load data integration process.
package etl

import "strings"

// Transform converts a legacy format of storing
// scrabble scores to a modern format that can be
// analyzed by the Load process of ETL.
func Transform(legacy map[int][]string) map[string]int {
	modern := make(map[string]int, len(legacy))
	for point, letters := range legacy {
		for _, l := range letters {
			modern[strings.ToLower(l)] = point
		}
	}
	return modern
}
