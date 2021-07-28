// Package strain provides basic separation functions.
package strain

type Ints []int
type Lists [][]int
type Strings []string

// Keep returns a new collection containing
// elements where the predicate is true.
func (i Ints) Keep(predicate func(int) bool) Ints {
	var collection Ints
	for _, v := range i {
		if predicate(v) {
			collection = append(collection, v)
		}
	}
	return collection
}

// Discard returns a new collection containing
// elements where the predicate is false.
func (i Ints) Discard(predicate func(int) bool) Ints {
	var collection Ints
	for _, v := range i {
		if !predicate(v) {
			collection = append(collection, v)
		}
	}
	return collection
}

// Keep returns a new collection containing
// elements where the predicate is true.
func (l Lists) Keep(predicate func([]int) bool) Lists {
	var collection Lists
	for _, v := range l {
		if predicate(v) {
			collection = append(collection, v)
		}
	}
	return collection
}

// Discard returns a new collection containing
// elements where the predicate is false.
func (l Lists) Discard(predicate func([]int) bool) Lists {
	var collection Lists
	for _, v := range l {
		if !predicate(v) {
			collection = append(collection, v)
		}
	}
	return collection
}

// Keep returns a new collection containing
// elements where the predicate is true.
func (s Strings) Keep(predicate func(string) bool) Strings {
	var collection Strings
	for _, v := range s {
		if predicate(v) {
			collection = append(collection, v)
		}
	}
	return collection
}

// Discard returns a new collection containing
// elements where the predicate is false.
func (s Strings) Discard(predicate func(string) bool) Strings {
	var collection Strings
	for _, v := range s {
		if !predicate(v) {
			collection = append(collection, v)
		}
	}
	return collection
}
