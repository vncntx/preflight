package stream_test

import (
	"os"
	"testing"

	"vincent.click/pkg/preflight"
	"vincent.click/pkg/preflight/stream"
)

func TestWritableClose(test *testing.T) {
	t := preflight.Unit(test)

	var name string
	w := stream.FromWritten(t.T, func(w *os.File) {
		defer w.Close()

		name = w.Name()
		w.WriteString(content)
	})

	t.Expect(w.Close()).Is().Nil()

	// temp file should no longer exist
	_, err := os.Stat(name)
	t.Expect(os.IsNotExist(err)).Is().True()
}

func TestWritableSize(test *testing.T) {
	t := preflight.Unit(test)

	w := stream.FromWritten(t.T, func(w *os.File) {
		defer w.Close()

		w.WriteString(content)
	})

	w.Size().Eq(len(content))
}

func TestWritableText(test *testing.T) {
	t := preflight.Unit(test)

	w := stream.FromWritten(t.T, func(w *os.File) {
		defer w.Close()

		w.WriteString(content)
	})

	w.Text().Eq(content)
}

func TestWritableBytes(test *testing.T) {
	t := preflight.Unit(test)

	w := stream.FromWritten(t.T, func(w *os.File) {
		defer w.Close()

		w.WriteString(content)
	})

	bytes := []byte(content)
	w.Bytes().Eq(bytes)
}

func TestWritableContentType(test *testing.T) {
	t := preflight.Unit(test)

	w := stream.FromWritten(t.T, func(w *os.File) {
		defer w.Close()

		w.WriteString(content)
	})

	w.ContentType().Matches("text/plain")
}
