package stream

import (
	"errors"
	"fmt"
	"io"
	"io/fs"
	"net/http"
	"os"
	"testing"

	"preflight/errs"
	"preflight/expect"
)

// Sink is a set of expectations about a writable stream
type Sink struct {
	*testing.T

	r   *os.File
	w   *os.File
	i   fs.FileInfo
	buf []byte
}

// FromWritable returns a new Sink
func FromWritable(t *testing.T, consumer Consumer) Stream {
	r, w, err := os.Pipe()
	if err != nil {
		return Faulty(t, err)
	}

	consumer(w)

	return &Sink{
		T: t,
		r: r,
		w: w,
		buf: make([]byte, 0, bufferSize),
	}

}

// Close the underlying streams
func (s *Sink) Close() error {
	fmt.Printf("closing %v, %v\n", s.r, s.w)

	re, we := s.r.Close(), s.w.Close()
	if re != nil || we != nil {
		return errs.Combine(re, we)
	}

	return nil
}

// Size returns an Expectation about the number of bytes written
func (s *Sink) Size() expect.Expectation {
	info, err := s.r.Stat()
	if err != nil {
		return expect.Faulty(s.T, err)
	}

	return expect.Value(s.T, info.Size())
}

// Text returns an Expectation about all text written to the sink
func (s *Sink) Text() expect.Expectation {
	if err := s.readAll(); err != nil {
		return expect.Faulty(s.T, err)
	}

	return expect.Value(s.T, string(s.buf))
}

// Bytes returns an Expectation about all bytes written to the sink
func (s *Sink) Bytes() expect.Expectation {
	if err := s.readAll(); err != nil {
		return expect.Faulty(s.T, err)
	}

	return expect.Value(s.T, s.buf)
}

// ContentType returns an Expectation about content type written to the sink
func (s *Sink) ContentType() expect.Expectation {
	if len(s.buf) < 1 {
		err := s.read(true)
		if err != nil {
			return expect.Faulty(s.T, err)
		}
	}
	var contentType = http.DetectContentType(s.buf)

	return expect.Value(s.T, contentType)
}

func (s *Sink) read(overwrite bool) error {
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

func (s *Sink) readAll() error {
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
