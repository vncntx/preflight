package stream

import (
	"testing"

	"vincent.click/pkg/preflight/expect"
)

// Fault is a set of stream expectations that always result in an error
type Fault struct {
	*testing.T

	err error
}

// Faulty returns a new Fault
func Faulty(t *testing.T, err error) Expectations {
	return &Fault{t, err}
}

// Close returns no error
func (f *Fault) Close() error {
	return nil
}

// Size returns a faulty Expectation
func (f *Fault) Size() expect.Expectation {
	return f.toExpectation()
}

// Text returns a faulty Expectation
func (f *Fault) Text() expect.Expectation {
	return f.toExpectation()
}

// NextText returns a faulty Expectation
func (f *Fault) NextText(bytes int) expect.Expectation {
	return f.toExpectation()
}

// TextAt returns a faulty Expectation
func (f *Fault) TextAt(pos int64, bytes int) expect.Expectation {
	return f.toExpectation()
}

// Bytes returns a faulty Expectation
func (f *Fault) Bytes() expect.Expectation {
	return f.toExpectation()
}

// NextBytes returns a faulty Expectation
func (f *Fault) NextBytes(bytes int) expect.Expectation {
	return f.toExpectation()
}

// BytesAt returns a faulty Expectation
func (f *Fault) BytesAt(pos int64, bytes int) expect.Expectation {
	return f.toExpectation()
}

// Lines returns a faulty Expectation
func (f *Fault) Lines() expect.Expectation {
	return f.toExpectation()
}

// NextLine returns a faulty Expectation
func (f *Fault) NextLine() expect.Expectation {
	return f.toExpectation()
}

// ContentType returns a faulty Expectation
func (f *Fault) ContentType() expect.Expectation {
	return f.toExpectation()
}

// toExpectation returns a faulty Expectation
func (f *Fault) toExpectation() expect.Expectation {
	return expect.Faulty(f.T, f.err)
}
