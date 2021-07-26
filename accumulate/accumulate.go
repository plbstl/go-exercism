// Package accumulate provides functionality
// for accumulation with different data types.
package accumulate

// Accumulate executes the given function
// on every item in a collection of strings.
func Accumulate(words []string, fn func(string) string) []string {
	for i, v := range words {
		words[i] = fn(v)
	}
	return words
}
