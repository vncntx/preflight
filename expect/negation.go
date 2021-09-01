package expect

import (
	"fmt"
	"reflect"
	"regexp"
	"testing"
)

// Negation is a negated expectation
type Negation struct {
	*testing.T

	inverse *ExpectedValue
}

// To returns the current Expectation
func (not *Negation) To() Expectation {
	return not
}

// Be returns the current Expectation
func (not *Negation) Be() Expectation {
	return not
}

// Is returns the current Expectation
func (not *Negation) Is() Expectation {
	return not
}

// Should returns the current Expectation
func (not *Negation) Should() Expectation {
	return not
}

// Not negates the current Expectation
func (not *Negation) Not() Expectation {
	return not.inverse
}

// IsNot is equivalent to Not()
func (not *Negation) IsNot() Expectation {
	return not.Not()
}

// DoesNot is equivalent to Not()
func (not *Negation) DoesNot() Expectation {
	return not.Not()
}

// At returns an negated expectation about the element at the given index
func (not *Negation) At(index interface{}) Expectation {
	return not.inverse.At(index).Not()
}

// Nil asserts the value is not nil
func (not *Negation) Nil() {
	not.Equals(nil)
}

// True asserts the value is not true
func (not *Negation) True() {
	not.Equals(true)
}

// False asserts the value is not false
func (not *Negation) False() {
	not.Equals(false)
}

// Empty asserts the value is a non-empty array
func (not *Negation) Empty() {
	not.HasLength(0)
}

// HasLength asserts the value is an array with length != given
func (not *Negation) HasLength(given int) {
	if reflect.ValueOf(not.inverse.actual).Len() == given {
		not.Errorf("%s: len(%#v) == %d", not.Name(), not.inverse.actual, given)
	}
}

// HaveLength is equivalent to HasLength()
func (not *Negation) HaveLength(given int) {
	not.HasLength(given)
}

// Equals asserts inequality to a given value
func (not *Negation) Equals(given interface{}) {
	if equal(not.inverse.actual, given) {
		not.Errorf("%s: %#v == %#v", not.Name(), not.inverse.actual, given)
	}
}

// Eq is equivalent to Equals()
func (not *Negation) Eq(given interface{}) {
	not.Equals(given)
}

// Equal is equivalent to Equals()
func (not *Negation) Equal(given interface{}) {
	not.Equals(given)
}

// EqualTo is equivalent to Equals()
func (not *Negation) EqualTo(given interface{}) {
	not.Equals(given)
}

// Matches asserts that the value does not match a given pattern
func (not *Negation) Matches(pattern string) {
	actual := fmt.Sprint(not.inverse.actual) // convert to string

	match, err := regexp.MatchString(pattern, actual)
	if err != nil {
		not.Error(err)
	} else if match {
		not.Errorf("%s: '%v' matches /%s/", not.Name(), not.inverse.actual, pattern)
	}
}

// Match is equivalent to Matches()
func (not *Negation) Match(pattern string) {
	not.Matches(pattern)
}
