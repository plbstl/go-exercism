// Package twofer implements a popular
// idiom — get 2 for the price of 1.
package twofer

// ShareWith shares an item with someone.
func ShareWith(name string) string {
	if name == "" {
		return "One for you, one for me."
	}
	return "One for " + name + ", one for me."
}
