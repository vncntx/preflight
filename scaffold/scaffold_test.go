package scaffold_test

import (
	"testing"

	"vincent.click/pkg/preflight"
	"vincent.click/pkg/preflight/scaffold"
)

func TestCaptureExitCode(test *testing.T) {
	t := preflight.Unit(test)

	s := scaffold.New()

	exit := 1
	action := func() {
		s.OSExit(exit)
	}

	t.Expect(s.CaptureExitCode(action)).Eq(exit)
}

func TestCapturePanic(test *testing.T) {
	t := preflight.Unit(test)

	s := scaffold.New()

	cause := "at the disco"
	action := func() {
		s.Panic(cause)
	}

	t.Expect(s.CapturePanic(action)).Eq(cause)
}
