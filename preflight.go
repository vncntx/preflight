package preflight

import (
	"os"
	"testing"

	"vincent.click/pkg/preflight/captor"
	"vincent.click/pkg/preflight/expect"
	"vincent.click/pkg/preflight/stream"
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
func (t *Test) ExpectFile(f *os.File) stream.Expectations {
	return stream.ExpectFile(t.T, f)
}

// ExpectWritten returns a set of expectations about data written to a stream
func (t *Test) ExpectWritten(consumer stream.Consumer) stream.Expectations {
	return stream.ExpectWritten(t.T, consumer)
}

// ExpectExitCode returns an expectation about a captured exit code
func (t *Test) ExpectExitCode(act captor.Action) expect.Expectation {
	code := Captor.CaptureExitCode(act)

	return expect.Value(t.T, code)
}

// ExpectPanic returns an expectation about the cause of a panicking goroutine
func (t *Test) ExpectPanic(act captor.Action) expect.Expectation {
	cause := Captor.CapturePanic(act)

	return expect.Value(t.T, cause)
}
