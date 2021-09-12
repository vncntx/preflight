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
	w := stream.ExpectWritten(t.T, func(w *os.File) {
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

	w := stream.ExpectWritten(t.T, writeContent)
	defer w.Close()

	w.Size().Eq(len(contents))
}

func TestWritableText(test *testing.T) {
	t := preflight.Unit(test)

	w := stream.ExpectWritten(t.T, writeContent)
	defer w.Close()

	text := string(contents)
	w.Text().Eq(text)
}

func TestWritableNextText(test *testing.T) {
	t := preflight.Unit(test)

	w := stream.ExpectWritten(t.T, writeContent)
	defer w.Close()

	w.NextText(3).Eq("Ad ")
	w.NextText(5).Eq("astra")
}

func TestWritableTextAt(test *testing.T) {
	t := preflight.Unit(test)

	w := stream.ExpectWritten(t.T, writeContent)
	defer w.Close()

	w.TextAt(3, 5).Eq("astra")
}

func TestWritableBytes(test *testing.T) {
	t := preflight.Unit(test)

	w := stream.ExpectWritten(t.T, writeContent)
	defer w.Close()

	w.Bytes().Eq(contents)
}

func TestWritableNextBytes(test *testing.T) {
	t := preflight.Unit(test)

	w := stream.ExpectWritten(t.T, writeContent)
	defer w.Close()

	w.NextBytes(3).Eq([]byte("Ad "))
	w.NextBytes(5).Eq([]byte("astra"))
}

func TestWritableBytesAt(test *testing.T) {
	t := preflight.Unit(test)

	w := stream.ExpectWritten(t.T, writeContent)
	defer w.Close()

	bytes := []byte("astra")
	w.BytesAt(3, 5).Eq(bytes)
}

func TestWritableLines(test *testing.T) {
	t := preflight.Unit(test)

	w := stream.ExpectWritten(t.T, writeContent)
	defer w.Close()

	w.Lines().At(1).Eq("Sic itur ad astra.")
}

func TestWritableNextLine(test *testing.T) {
	t := preflight.Unit(test)

	w := stream.ExpectWritten(t.T, writeContent)
	defer w.Close()

	w.NextLine().Eq("Ad astra per aspera.")
	w.NextLine().Eq("Sic itur ad astra.")
	w.NextLine().Is().Empty()
}

func TestWritableContentType(test *testing.T) {
	t := preflight.Unit(test)

	w := stream.ExpectWritten(t.T, writeContent)
	defer w.Close()

	w.ContentType().Matches("text/plain")
}

func writeContent(w *os.File) {
	defer w.Close()

	if _, err := w.Write(contents); err != nil {
		panic(err)
	}
}
