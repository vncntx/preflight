package preflight

import (
	"os"
	"testing"

	"preflight/expect"
	"preflight/scaffold"
	"preflight/stream"
)

// Test provides utilities for testing
type Test struct {
	*testing.T
}

// Expect returns a new value-based expectation
func (t *Test) Expect(actual interface{}) expect.Expectation {
	return expect.Value(t.T, actual)
}

// ExpectFile returns a set of expectations about a file
func (t *Test) ExpectFile(f *os.File) stream.Stream {
	return stream.FromFile(t.T, f)
}

// ExpectWritten returns a set of expectations about data written to a stream
func (t *Test) ExpectWritten(consumer stream.Consumer) stream.Stream {
	return stream.FromWritable(t.T, consumer)
}

// ExpectExitCode returns an expectation about a captured exit code
func (t *Test) ExpectExitCode(act scaffold.Action) expect.Expectation {
	code := Scaffold.CaptureExitCode(act)

	return expect.Value(t.T, code)
}
