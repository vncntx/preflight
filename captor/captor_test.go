package captor_test

import (
	"testing"

	"vincent.click/pkg/preflight"
	"vincent.click/pkg/preflight/captor"
)

func TestCaptureExitCode(test *testing.T) {
	t := preflight.Unit(test)

	s := captor.New()

	exit := 1
	action := func() {
		s.Exit(exit)
	}

	t.Expect(s.CaptureExitCode(action)).Eq(exit)
}

func TestCapturePanic(test *testing.T) {
	t := preflight.Unit(test)

	s := captor.New()

	cause := "at the disco"
	action := func() {
		s.Panic(cause)
	}

	t.Expect(s.CapturePanic(action)).Eq(cause)
}
