package preflight_test

import (
	"os"
	"testing"

	"preflight"
)

func TestExpectFile(test *testing.T) {
	t := preflight.Unit(test)

	written := t.ExpectWritten(func(w *os.File) {
		defer w.Close()

		if _, err := w.WriteString("xsyz"); err != nil {
			t.Error(err)
		}
	})
	defer written.Close()

	written.Text().Eq("xyz")
	written.Bytes().Eq([]byte("xyz"))
	written.Size().Eq(3)
}

func TestExpect(test *testing.T) {
	t := preflight.Unit(test)

	t.Expect(1 + 1).Eq(3)
}

func TestFatal(test *testing.T) {
	t := preflight.Unit(test)

	t.ExpectExitCode(func() {
		preflight.Scaffold.OSExit(1)
	}).Eq(1)
}
