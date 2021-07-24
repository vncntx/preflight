package preflight

import (
	"os"
	"testing"

	"preflight/expect"
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

// ExpectExitCode overrides the scaffolding osExit function
// func (t *Test) ExpectExitCode(act Action) expect.Expectation {
// 	var exitCode int

// 	// capture exit code through scaffold
// 	Scaffold.OSExit = func(code int) {
// 		exitCode = code
// 	}
// 	defer Scaffold.Restore()

// 	// invoke the action
// 	act()

// 	return t.Expect(exitCode)
// }
