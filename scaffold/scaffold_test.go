package scaffold_test

import (
	"testing"

	"preflight"
	"preflight/scaffold"
)

func TestCaptureExitCode(test *testing.T) {
	t := preflight.Unit(test)

	s := scaffold.New()

	exit := 1
	action := func() {
		s.OSExit(1)
	}

	t.Expect(s.CaptureExitCode(action)).Eq(exit)
}
