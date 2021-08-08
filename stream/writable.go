package stream

import (
	"errors"
	"io"
	"net/http"
	"os"
	"testing"

	"vincent.click/pkg/preflight/expect"
)

// Writable is a set of expectations about a writable stream
type Writable struct {
	*testing.T

	r   *os.File
	buf []byte
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
		T:   t,
		r:   r,
		buf: make([]byte, 0, bufferSize),
	}

}

// Close and the stream
func (s *Writable) Close() error {
	if err := s.r.Close(); err != nil {
		return err
	}

	return os.Remove(s.r.Name())
}

// Size returns an Expectation about the number of bytes written
func (s *Writable) Size() expect.Expectation {
	info, err := s.r.Stat()
	if err != nil {
		return expect.Faulty(s.T, err)
	}

	return expect.Value(s.T, info.Size())
}

// Text returns an Expectation about all text written to the stream
func (s *Writable) Text() expect.Expectation {
	if err := s.readAll(); err != nil {
		return expect.Faulty(s.T, err)
	}

	return expect.Value(s.T, string(s.buf))
}

// Bytes returns an Expectation about all bytes written to the stream
func (s *Writable) Bytes() expect.Expectation {
	if err := s.readAll(); err != nil {
		return expect.Faulty(s.T, err)
	}

	return expect.Value(s.T, s.buf)
}

// ContentType returns an Expectation about content type written to the stream
func (s *Writable) ContentType() expect.Expectation {
	if len(s.buf) < 1 {
		err := s.read(true)
		if err != nil {
			return expect.Faulty(s.T, err)
		}
	}
	var contentType = http.DetectContentType(s.buf)

	return expect.Value(s.T, contentType)
}

func (s *Writable) read(overwrite bool) error {
	start := 0

	if !overwrite {
		// next read will append to end of buffer
		start = len(s.buf)
		// add more capacity as needed
		if len(s.buf) == cap(s.buf) {
			s.buf = append(s.buf, 0)[:len(s.buf)]
		}
	}

	n, err := s.r.Read(s.buf[start:cap(s.buf)])
	s.buf = s.buf[:start+n]

	return err
}

func (s *Writable) readAll() error {
	for {
		err := s.read(false)
		if errors.Is(io.EOF, err) {
			break
		} else if err != nil {
			return err
		}
	}

	return nil
}
