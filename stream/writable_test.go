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
	})

	t.Expect(w.Close()).Is().Nil()

	// temp file should no longer exist
	_, err := os.Stat(name)
	t.Expect(os.IsNotExist(err)).Is().True()
}

func TestWritableSize(test *testing.T) {
	t := preflight.Unit(test)

	w := stream.FromWritten(t.T, writeContent)
	defer w.Close()

	w.Size().Eq(len(content))
}

func TestWritableText(test *testing.T) {
	t := preflight.Unit(test)

	w := stream.FromWritten(t.T, writeContent)
	defer w.Close()

	w.Text().Eq(content)
}

func TestWritableTextAt(test *testing.T) {
	t := preflight.Unit(test)

	w := stream.FromWritten(t.T, writeContent)
	defer w.Close()

	w.TextAt(3, 5).Eq("astra")
}

func TestWritableBytes(test *testing.T) {
	t := preflight.Unit(test)

	w := stream.FromWritten(t.T, writeContent)
	defer w.Close()

	bytes := []byte(content)
	w.Bytes().Eq(bytes)
}

func TestWritableNextBytes(test *testing.T) {
	t := preflight.Unit(test)

	w := stream.FromWritten(t.T, writeContent)
	defer w.Close()

	w.NextBytes(3).Eq([]byte("Ad "))
	w.NextBytes(5).Eq([]byte("astra"))
}

func TestWritableBytesAt(test *testing.T) {
	t := preflight.Unit(test)

	w := stream.FromWritten(t.T, writeContent)
	defer w.Close()

	bytes := []byte("astra")
	w.BytesAt(3, 5).Eq(bytes)
}

func TestWritableContentType(test *testing.T) {
	t := preflight.Unit(test)

	w := stream.FromWritten(t.T, writeContent)
	defer w.Close()

	w.ContentType().Matches("text/plain")
}

func writeContent(w *os.File) {
	defer w.Close()

	if _, err := w.WriteString(content); err != nil {
		panic(err)
	}
}
