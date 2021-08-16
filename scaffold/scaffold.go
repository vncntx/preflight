package scaffold

import "os"

// Scaffold is a set of functions that may be replaced in tests
type Scaffold struct {
	OSExit func(int)
	Panic  func(interface{})
}

// New returns a new Scaffold
func New() *Scaffold {
	scaffold := &Scaffold{}
	scaffold.Reset()

	return scaffold
}

// Reset returns all functions to their initial state
func (s *Scaffold) Reset() {
	s.OSExit = os.Exit
	s.Panic = panik
}

func panik(cause interface{}) {
	panic(cause)
}

// CaptureExitCode overrides OSExit and returns the captured exit code
func (s *Scaffold) CaptureExitCode(act Action) int {
	var code int

	s.OSExit = func(c int) {
		code = c
	}
	defer s.Reset()

	act()

	return code
}

// CapturePanic overrides Panic and returns the captured cause
func (s *Scaffold) CapturePanic(act Action) interface{} {
	var cause interface{}

	s.Panic = func(c interface{}) {
		cause = c
	}
	defer s.Reset()

	act()

	return cause
}
