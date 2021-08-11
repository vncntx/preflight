package stream

import (
	"io/fs"
	"net/http"
	"os"
	"testing"

	"vincent.click/pkg/preflight/expect"
)

// Writable is a set of expectations about a writable stream
type Writable struct {
	*testing.T

	r   *os.File
	mod fs.FileMode
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

	info, err := r.Stat()
	if err != nil {
		return Faulty(t, err)
	}

	return &Writable{
		T:   t,
		r:   r,
		mod: info.Mode(),
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
	txt, err := readAll(w.r, w.mod)
	if err != nil {
		return expect.Faulty(w.T, err)
	}

	return expect.Value(w.T, string(txt))
}

// Bytes returns an Expectation about all bytes written to the stream
func (w *Writable) Bytes() expect.Expectation {
	bytes, err := readAll(w.r, w.mod)
	if err != nil {
		return expect.Faulty(w.T, err)
	}

	return expect.Value(w.T, bytes)
}

// BytesAt returns an Expectation about the bytes written at a specific position
func (w *Writable) BytesAt(pos int64, length int) expect.Expectation {
	if err := seek(w.r, w.mod, pos); err != nil {
		return expect.Faulty(w.T, err)
	}

	bytes, err := read(w.r, length)
	if err != nil {
		return expect.Faulty(w.T, err)
	}

	return expect.Value(w.T, bytes)
}

// ContentType returns an Expectation about content type written to the stream
func (w *Writable) ContentType() expect.Expectation {
	if err := seek(w.r, w.mod, 0); err != nil {
		return expect.Faulty(w.T, err)
	}

	content, err := read(w.r, 512)
	if err != nil {
		return expect.Faulty(w.T, err)
	}

	contentType := http.DetectContentType(content)

	return expect.Value(w.T, contentType)
}
