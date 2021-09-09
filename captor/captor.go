package captor

import "os"

// Captor is a set of aliases to builtin functions
type Captor struct {
	// Exit is an alias for os.Exit
	Exit func(int)
	// Panic is an alias for panic
	Panic func(interface{})
}

// New returns a new Captor
func New() *Captor {
	captor := &Captor{}
	captor.Reset()

	return captor
}

// Reset returns all functions to their initial state
func (s *Captor) Reset() {
	s.Exit = os.Exit
	s.Panic = panik
}

func panik(cause interface{}) {
	panic(cause)
}

// CaptureExitCode captures the argument to Exit
func (s *Captor) CaptureExitCode(act Action) int {
	var code int

	s.Exit = func(c int) {
		code = c
	}
	defer s.Reset()

	act()

	return code
}

// CapturePanic captures the argument to Panic
func (s *Captor) CapturePanic(act Action) interface{} {
	var cause interface{}

	s.Panic = func(c interface{}) {
		cause = c
	}
	defer s.Reset()

	act()

	return cause
}
