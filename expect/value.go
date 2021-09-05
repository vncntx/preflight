package expect

import (
	"fmt"
	"reflect"
	"regexp"
	"testing"

	"vincent.click/pkg/preflight/expect/kind"
)

// ExpectedValue is an expectation on a realized value
type ExpectedValue struct {
	*testing.T

	actual interface{}
}

// Value returns a new value-based Expectation
func Value(t *testing.T, actual interface{}) Expectation {
	return &ExpectedValue{
		T:      t,
		actual: actual,
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
		inverse: e,
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

// At returns an expectation about the element at the given index
func (e *ExpectedValue) At(index interface{}) Expectation {
	k := kind.Of(e.actual)

	switch k {
	case kind.List:
		if i, ok := index.(int); ok {
			element := reflect.ValueOf(e.actual).Index(i).Interface()

			return Value(e.T, element)
		}

		return Faulty(e.T, fmt.Errorf("%s: slice or array index must be an int", e.Name()))

	case kind.String:
		if i, ok := index.(int); ok {
			element := []rune(e.actual.(string))[i]

			return Value(e.T, element)
		}

		return Faulty(e.T, fmt.Errorf("%s: string index must be an int", e.Name()))

	case kind.Map:

		key := reflect.ValueOf(index)
		element := reflect.ValueOf(e.actual).MapIndex(key).Interface()

		return Value(e.T, element)

	default:
		return Faulty(e.T, fmt.Errorf("%s: %v: not a slice, array, string, or map", e.Name(), e.actual))
	}
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
	if reflect.ValueOf(e.actual).Len() != expected {
		e.Errorf("%s: len(%#v) != %d", e.Name(), e.actual, expected)
	}
}

// HaveLength is equivalent to HasLength()
func (e *ExpectedValue) HaveLength(expected int) {
	e.HasLength(expected)
}

// Equals asserts equality to an expected value
func (e *ExpectedValue) Equals(expected interface{}) {
	if !equal(e.actual, expected) {
		e.Errorf("%s: %#v != %#v", e.Name(), e.actual, expected)
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
	str := fmt.Sprint(e.actual)

	match, err := regexp.MatchString(pattern, str)
	if err != nil {
		e.Errorf("%s: %s", e.Name(), err)
	} else if !match {
		e.Errorf("%s: '%s' does not match /%s/", e.Name(), str, pattern)
	}
}

// Match is equivalent to Matches()
func (e *ExpectedValue) Match(pattern string) {
	e.Matches(pattern)
}
