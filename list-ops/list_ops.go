// Package listops implements basic
// list operations for type int.
package listops

// IntList is a slice of type int.
type IntList []int

// Append returns an updated list with
// elements added to the end.
func (l IntList) Append(elements IntList) IntList {
	return append(l, elements...)
}

// Concat flattens a combined series of
// lists and returns the updated list.
func (l IntList) Concat(lists ...IntList) IntList {
	for _, list := range lists {
		l = l.Append(list)
	}
	return l
}

// Filter returns the list of all elements
// for which predicate(element) is true.
func (l IntList) Filter(predicate func(int) bool) IntList {
	il := make(IntList, 0, l.Length())
	for _, n := range l {
		if predicate(n) {
			il = append(il, n)
		}
	}
	return il
}

// Length returns the total number
// of elements in the list.
func (l IntList) Length() (size int) {
	for range l {
		size++
	}
	return
}

// Map returns a list of the results of
// applying fn(element) on all elements.
func (l IntList) Map(fn func(int) int) IntList {
	for i, elem := range l {
		l[i] = fn(elem)
	}
	return l
}

// Foldl reduces each element into the accumulator
// from the left using fn(accumulator, element).
func (l IntList) Foldl(fn func(int, int) int, acc int) int {
	for _, elem := range l {
		acc = fn(acc, elem)
	}
	return acc
}

// Foldr reduces each element into the accumulator
// from the right using fn(element, accumulator).
func (l IntList) Foldr(fn func(int, int) int, acc int) int {
	for _, elem := range l.Reverse() {
		acc = fn(elem, acc)
	}
	return acc
}

// Reverse returns a list with all the
// original elements, but in reversed order.
func (l IntList) Reverse() IntList {
	size := l.Length()
	for i := 0; i < size/2; i++ {
		l[i], l[size-1-i] = l[size-1-i], l[i]
	}
	return l
}
