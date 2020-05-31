package paginator

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/dhowden/numerus"
)

// NumeralCursor represents a numeral (of any system) that can be incremented
type NumeralCursor interface {
	PeekNext() string
	Current() string
	Increment()
	Set(string) error
	IsNextValue(string) (bool, bool)
}

// NewNumeralCursorFromString returns a NumeralCursor appropriate for the string passed in.
// For example: if "ii" is passed in, a RomanNumeralCursor is returned.
// If "5" is passed in, an ArabicNumeralCursor is returned.
// Currently lowercase roman numerals and arabic numerals are supported.
// If any other type of numerals are passed in, an error will be returned
func NewNumeralCursorFromString(str string) (numeral NumeralCursor, err error) {
	numeral = NewArabicNumeralCursor()
	err = numeral.Set(str)
	if err == nil {
		return
	}
	numeral = NewRomanNumeralCursor()
	err = numeral.Set(str)
	if err == nil {
		return
	}
	return nil, fmt.Errorf("invalid numeral %#v", str)
}

// ArabicNumeralCursor can increment arabic numerals (1, 2, 3...)
type ArabicNumeralCursor struct {
	numeral string
}

// NewArabicNumeralCursor returns an ArabicNumeralCursor starting at "1"
func NewArabicNumeralCursor() *ArabicNumeralCursor {
	return &ArabicNumeralCursor{
		numeral: "1",
	}
}

// PeekNext returns the next numeral without incrementing the cursor.
func (cursor ArabicNumeralCursor) PeekNext() string {
	pageInt, err := strconv.Atoi(cursor.numeral)
	if err != nil {
		panic(err)
	}
	return strconv.Itoa(pageInt + 1)
}

// Current returns the current numeral the cursor is on
func (cursor ArabicNumeralCursor) Current() string {
	return cursor.numeral
}

// Increment advances the numeral to the next value
func (cursor *ArabicNumeralCursor) Increment() {
	pageInt, err := strconv.Atoi(cursor.numeral)
	if err != nil {
		panic(err)
	}
	cursor.numeral = strconv.Itoa(pageInt + 1)
}

// Set parses a string into an arabic numeral.
// Errors if value is not an arabic numeral
func (cursor *ArabicNumeralCursor) Set(value string) error {
	_, err := strconv.Atoi(value)
	if err != nil {
		return fmt.Errorf("invalid arabic numeral %#v", value)
	}
	cursor.numeral = value
	return nil
}

//IsNextValue checks if candidate the next larger arabic numeral, and if its even a valid arabic numeral
func (cursor ArabicNumeralCursor) IsNextValue(candidate string) (isNextValue bool, isSameType bool) {
	_, err := strconv.Atoi(candidate)
	isSameType = err == nil
	isNextValue = cursor.PeekNext() == candidate
	return
}

// RomanNumeralCursor can increment lower case roman numerals (i, ii, iii, iv, v...)
type RomanNumeralCursor struct {
	numeral numerus.Numeral
}

// NewRomanNumeralCursor returns an RomanNumeralCursor starting at "i"
func NewRomanNumeralCursor() *RomanNumeralCursor {
	return &RomanNumeralCursor{numeral: 1}
}

// PeekNext returns the next numeral without incrementing the cursor.
func (cursor RomanNumeralCursor) PeekNext() string {
	return strings.ToLower((cursor.numeral + 1).String())
}

// Current returns the current numeral the cursor is on
func (cursor RomanNumeralCursor) Current() string {
	return strings.ToLower(cursor.numeral.String())
}

// Increment advances the numeral to the next value
func (cursor *RomanNumeralCursor) Increment() {
	cursor.numeral++
}

// Set parses a string into an arabic numeral.
// Errors if value is not a roman numeral
func (cursor *RomanNumeralCursor) Set(value string) error {
	newNumeral, err := numerus.Parse(strings.ToUpper(value))
	if err != nil {
		return fmt.Errorf("invalid roman numeral %#v", value)
	}
	cursor.numeral = newNumeral
	return nil
}

//IsNextValue checks if candidate the next larger roman numeral, and if its even a valid roman numeral
func (cursor RomanNumeralCursor) IsNextValue(candidate string) (isNextValue bool, isSameType bool) {
	candidateNumeral, err := numerus.Parse(strings.ToUpper(candidate))
	if err != nil {
		return false, false
	}
	nextNumeral := cursor.numeral + 1
	return nextNumeral == candidateNumeral, true
}
