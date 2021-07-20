package preflight_test

import (
	"os"
	"testing"

	"preflight"
)

func TestExpectFile(test *testing.T) {
	t := preflight.Unit(test)

	expect := t.ExpectWritten(func(w *os.File) {

		w.WriteString("xyz")
		defer w.Close()

	})
	defer expect.Close()

	expect.Text().Eq("xyz")
	expect.Bytes().Eq([]byte("xyz"))
	expect.Size().Eq(3)
}
