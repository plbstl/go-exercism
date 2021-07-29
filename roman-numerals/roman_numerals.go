// Package romannumerals provides functionality
// for roman numeral conversions.
package romannumerals

import (
	"errors"
	"strconv"
)

const (
	one = iota + 1
	two
	three
	four
	five
	six
	seven
	eight
	nine
)

const (
	ten      = 10
	hundred  = 100
	thousand = 1000
)

// ToRomanNumeral returns the equilivant roman
// number for positive integers up to 3000.
func ToRomanNumeral(n int) (string, error) {
	var (
		romanNumber string
		err         error
	)

	switch {
	case n < 1:
		err = errors.New("number `" + strconv.Itoa(n) + "` is less than 1")
	case n < 10:
		romanNumber = units(n)
	case n < 100:
		digit1, digit2 := twoDigits(n)
		romanNumber = tens(digit1) + units(digit2)
	case n < 1000:
		digit1, digit2, digit3 := threeDigits(n)
		romanNumber = hundreds(digit1) + tens(digit2) + units(digit3)
	case n < 3001:
		digit1, digit2, digit3, digit4 := fourDigits(n)
		romanNumber = thousands(digit1) + hundreds(digit2) + tens(digit3) + units(digit4)
	default:
		err = errors.New("number `" + strconv.Itoa(n) + "` is greater than 3000")
	}

	return romanNumber, err
}

func units(n int) string {
	switch n {
	case one:
		return "I"
	case two:
		return "II"
	case three:
		return "III"
	case four:
		return "IV"
	case five:
		return "V"
	case six:
		return "VI"
	case seven:
		return "VII"
	case eight:
		return "VIII"
	case nine:
		return "IX"
	default:
		return ""
	}
}

func tens(n int) string {
	switch n {
	case one:
		return "X"
	case two:
		return "XX"
	case three:
		return "XXX"
	case four:
		return "XL"
	case five:
		return "L"
	case six:
		return "LX"
	case seven:
		return "LXX"
	case eight:
		return "LXXX"
	case nine:
		return "XC"
	default:
		return ""
	}
}

func hundreds(n int) string {
	switch n {
	case one:
		return "C"
	case two:
		return "CC"
	case three:
		return "CCC"
	case four:
		return "CD"
	case five:
		return "D"
	case six:
		return "DC"
	case seven:
		return "DCC"
	case eight:
		return "DCCC"
	case nine:
		return "CM"
	default:
		return ""
	}
}

func thousands(n int) string {
	switch n {
	case one:
		return "M"
	case two:
		return "MM"
	case three:
		return "MMM"
	default:
		return ""
	}
}

// twoDigits separates a two digit number.
// Expects numbers 10⁠–⁠99.
func twoDigits(n int) (int, int) {
	a, b := extractFirstDigit(n, ten)
	return a, b
}

// threeDigits separates a three digit number.
// Expects numbers 100⁠–⁠999.
func threeDigits(n int) (int, int, int) {
	a, b := extractFirstDigit(n, hundred)
	b, c := twoDigits(b)
	return a, b, c
}

// fourDigits separates a four digit number.
// Expects numbers 1000⁠–⁠9999.
func fourDigits(n int) (int, int, int, int) {
	a, b := extractFirstDigit(n, thousand)
	b, c, d := threeDigits(b)
	return a, b, c, d
}

// extractFirstDigit returns the first digit of number n based on the
// number value — tens, hundreds, thousands — and the remaining digit(s).
//
// For example:
//
// var a, b = extractFirstDigit(1234, thousand)
//
// // a = 1, b = 234
func extractFirstDigit(n, numValue int) (int, int) {
	var a, b = 9, -1
	for ; a > 0; a-- {
		b = n - (a * numValue)
		if b >= 0 {
			break
		}
	}
	if a == 0 {
		b = n
	}
	return a, b
}
