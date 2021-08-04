package scaffold

import "os"

// Scaffold is a set of functions that may be replaced in tests
type Scaffold struct {
	OSExit func(int)
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
