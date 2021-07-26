// Package leap provide helpers for working
// with Gregorian calendar leap years.
package leap

// IsLeapYear checks if the given year is a leap year.
func IsLeapYear(year int) bool {
	isDivisibleBy4 := year%4 == 0
	isDivisibleBy100 := year%100 == 0
	isDivisibleBy400 := year%400 == 0

	return isDivisibleBy400 || (isDivisibleBy4 && !isDivisibleBy100)
}
