package errs_test

import (
	"errors"
	"testing"

	"preflight"
	"preflight/errs"
)

func TestCombine(test *testing.T) {
	t := preflight.Unit(test)

	err1 := errors.New("error 1")
	err2 := errors.New("error 2")
	err3 := errors.New("error 3")

	err := errs.Combine(err1, err2, err3)

	t.Expect(err.Error()).Eq("error 1; error 2; error 3")
}
