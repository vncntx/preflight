package expect

import (
	"testing"
)

// Fault is an expectation that always results in an error
type Fault struct {
	*testing.T

	err error
}

// Faulty returns a new Fault
func Faulty(t *testing.T, err error) Expectation {
	return &Fault{t, err}
}

// To returns the current expectation
func (f *Fault) To() Expectation {
	return f
}

// Be returns the current expectation
func (f *Fault) Be() Expectation {
	return f
}

// Is returns the current expectation
func (f *Fault) Is() Expectation {
	return f
}

// Should returns the current expectation
func (f *Fault) Should() Expectation {
	return f
}

// Not returns the current expectation
func (f *Fault) Not() Expectation {
	return f
}

// IsNot returns the current expectation
func (f *Fault) IsNot() Expectation {
	return f
}

// DoesNot returns the current expectation
func (f *Fault) DoesNot() Expectation {
	return f
}

// Nil always results in an error
func (f *Fault) Nil() {
	f.Error(f.err)
}

// True always results in an error
func (f *Fault) True() {
	f.Error(f.err)
}

// False always results in an error
func (f *Fault) False() {
	f.Error(f.err)
}

// Empty always results in an error
func (f *Fault) Empty() {
	f.Error(f.err)
}

// HasLength always results in an error
func (f *Fault) HasLength(expected int) {
	f.Error(f.err)
}

// HaveLength always results in an error
func (f *Fault) HaveLength(expected int) {
	f.Error(f.err)
}

// Equals always results in an error
func (f *Fault) Equals(expected interface{}) {
	f.Error(f.err)
}

// Eq always results in an error
func (f *Fault) Eq(expected interface{}) {
	f.Error(f.err)
}

// Equal always results in an error
func (f *Fault) Equal(expected interface{}) {
	f.Error(f.err)
}

// EqualTo always results in an error
func (f *Fault) EqualTo(expected interface{}) {
	f.Error(f.err)
}

// Matches always results in an error
func (f *Fault) Matches(pattern string) {
	f.Error(f.err)
}

// Match always results in an error
func (f *Fault) Match(pattern string) {
	f.Error(f.err)
}
