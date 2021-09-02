package stream

import (
	"os"
	"testing"

	"vincent.click/pkg/preflight/expect"
)

// Writable is a set of expectations about a writable stream
type Writable struct {
	*testing.T

	r reader
	b []byte
}

// FromWritten returns a new Writable
func FromWritten(t *testing.T, consumer Consumer) Stream {
	w, err := os.CreateTemp(os.TempDir(), "preflight-")
	if err != nil {
		return Faulty(t, err)
	}

	consumer(w)

	// open for reading
	r, err := os.OpenFile(w.Name(), os.O_RDONLY, 0)
	if err != nil {
		return Faulty(t, err)
	}

	return &Writable{
		T: t,
		r: r,
	}
}

// Close the stream and remove the temporary file
func (w *Writable) Close() error {
	if err := w.r.Close(); err != nil {
		return err
	}

	return os.Remove(w.r.Name())
}

// Size returns an Expectation about the number of bytes written
func (w *Writable) Size() expect.Expectation {
	stat, err := w.r.Stat()
	if err != nil {
		return expect.Faulty(w.T, err)
	}

	return expect.Value(w.T, stat.Size())
}

// Text returns an Expectation about all text written to the stream
func (w *Writable) Text() expect.Expectation {
	txt, err := readAll(w.r)
	if err != nil {
		return expect.Faulty(w.T, err)
	}

	return expect.Value(w.T, string(txt))
}

// NextText returns an Expectation about the next chunk of text written to the stream
func (w *Writable) NextText(bytes int) expect.Expectation {
	data, err := read(w.r, bytes)
	if err != nil {
		return expect.Faulty(w.T, err)
	}

	return expect.Value(w.T, string(data))
}

// TextAt returns an Expectation about the text written at a specific position
func (w *Writable) TextAt(pos int64, bytes int) expect.Expectation {
	data, err := readAt(w.r, pos, bytes)
	if err != nil {
		return expect.Faulty(w.T, err)
	}

	return expect.Value(w.T, string(data))
}

// Bytes returns an Expectation about all bytes written to the stream
func (w *Writable) Bytes() expect.Expectation {
	bytes, err := readAll(w.r)
	if err != nil {
		return expect.Faulty(w.T, err)
	}

	return expect.Value(w.T, bytes)
}

// NextBytes returns an Expectation about the next chunk of bytes written to the stream
func (w *Writable) NextBytes(bytes int) expect.Expectation {
	data, err := read(w.r, bytes)
	if err != nil {
		return expect.Faulty(w.T, err)
	}

	return expect.Value(w.T, data)
}

// BytesAt returns an Expectation about the bytes written at a specific position
func (w *Writable) BytesAt(pos int64, bytes int) expect.Expectation {
	data, err := readAt(w.r, pos, bytes)
	if err != nil {
		return expect.Faulty(w.T, err)
	}

	return expect.Value(w.T, data)
}

// Lines returns an Expectation about all lines of text written to the stream
func (w *Writable) Lines() expect.Expectation {
	lines := []string{}

	for {
		line, bytes, err := readLine(w.r, w.b)
		if err != nil {
			return expect.Faulty(w.T, err)
		}

		w.b = bytes

		if len(line) < 1 {
			break
		}

		lines = append(lines, string(line))
	}

	return expect.Value(w.T, lines)
}

// NextLine returns an Expectation about the next line of text
func (w *Writable) NextLine() expect.Expectation {
	line, bytes, err := readLine(w.r, w.b)
	if err != nil {
		return expect.Faulty(w.T, err)
	}

	w.b = bytes

	return expect.Value(w.T, string(line))
}

// ContentType returns an Expectation about content type written to the stream
func (w *Writable) ContentType() expect.Expectation {
	if typ, err := detectContentType(w.r); err != nil {
		return expect.Faulty(w.T, err)
	} else {
		return expect.Value(w.T, typ)
	}
}
