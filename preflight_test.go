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

		if _, err := w.WriteString("xyz"); err != nil {
			t.Error(err)
		}
	})
	defer written.Close()

	written.Text().Eq("xyz")
	written.Bytes().Eq([]byte("xyz"))
	written.Size().Eq(3)
}
