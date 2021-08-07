package expect

import (
	"fmt"
	"reflect"
	"regexp"
	"testing"
)

// ExpectedValue is an expectation on a realized value
type ExpectedValue struct {
	*testing.T

	Actual interface{}
}

// Value returns a new value-based Expectation
func Value(t *testing.T, actual interface{}) Expectation {
	return &ExpectedValue{
		T:      t,
		Actual: actual,
	}
}

// To returns the current expectation
func (e *ExpectedValue) To() Expectation {
	return e
}

// Be returns the current expectation
func (e *ExpectedValue) Be() Expectation {
	return e
}

// Is returns the current expectation
func (e *ExpectedValue) Is() Expectation {
	return e
}

// Should returns the current expectation
func (e *ExpectedValue) Should() Expectation {
	return e
}

// Not negates the current expectation
func (e *ExpectedValue) Not() Expectation {
	return &Negation{
		T:       e.T,
		Actual:  e.Actual,
		Inverse: e,
	}
}

// IsNot is equivalent to Not()
func (e *ExpectedValue) IsNot() Expectation {
	return e.Not()
}

// DoesNot is equivalent to Not()
func (e *ExpectedValue) DoesNot() Expectation {
	return e.Not()
}

// Nil asserts the value is nil
func (e *ExpectedValue) Nil() {
	e.Equals(nil)
}

// True asserts the value is true
func (e *ExpectedValue) True() {
	e.Equals(true)
}

// False asserts the value is false
func (e *ExpectedValue) False() {
	e.Equals(false)
}

// Empty asserts the value has length 0
func (e *ExpectedValue) Empty() {
	e.HasLength(0)
}

// HasLength asserts the value is an array with a given length
func (e *ExpectedValue) HasLength(expected int) {
	if reflect.ValueOf(e.Actual).Len() != expected {
		e.Errorf("%s: len(%#v) != %d", e.Name(), e.Actual, expected)
	}
}

// HaveLength is equivalent to HasLength()
func (e *ExpectedValue) HaveLength(expected int) {
	e.HasLength(expected)
}

// Equals asserts equality to an expected value
func (e *ExpectedValue) Equals(expected interface{}) {
	if !equal(e.Actual, expected) {
		e.Errorf("%s: %#v != %#v", e.Name(), e.Actual, expected)
	}
}

// Eq is equivalent to Equals()
func (e *ExpectedValue) Eq(expected interface{}) {
	e.Equals(expected)
}

// Equal is equivalent to Equals()
func (e *ExpectedValue) Equal(expected interface{}) {
	e.Equals(expected)
}

// EqualTo is equivalent to Equals()
func (e *ExpectedValue) EqualTo(expected interface{}) {
	e.Equals(expected)
}

// Matches asserts that the value matches a given pattern
func (e *ExpectedValue) Matches(pattern string) {
	actual := fmt.Sprint(e.Actual) // convert to string

	match, err := regexp.MatchString(pattern, actual)
	if err != nil {
		e.Error(err)
	} else if !match {
		e.Errorf("'%#v' does not match /%s/", e.Actual, pattern)
	}
}

// Match is equivalent to Matches()
func (e *ExpectedValue) Match(pattern string) {
	e.Matches(pattern)
}
