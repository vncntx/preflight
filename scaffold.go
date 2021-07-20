package preflight

import "os"

// Scaffold is a set of functions that may be mocked during testing
var Scaffold *scaffold

type scaffold struct {
	OSExit func(int)
}

// Restore restores the scaffold to its default state
func (s *scaffold) Restore() {
	s.OSExit = os.Exit
}
