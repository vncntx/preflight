package stream

import (
	"errors"
	"fmt"
	"io"
	"io/fs"
	"net/http"
	"os"
	"testing"

	"vincent.click/pkg/preflight/expect"
)

// File is a set of expectations about a file stream
type File struct {
	*testing.T

	d   *os.File
	mod fs.FileMode
	pos int64
	buf []byte
}

// FromFile returns expectations based on a file descriptor
func FromFile(t *testing.T, file *os.File) Stream {
	info, err := file.Stat()
	if err != nil {
		return Faulty(t, err)
	}

	return &File{
		T:   t,
		d:   file,
		mod: info.Mode(),
		buf: make([]byte, 0, bufferSize),
	}
}

// Close the underlying file descriptor
func (f *File) Close() error {
	return f.d.Close()
}

// Size returns an Expectation about the file size in bytes
func (f *File) Size() expect.Expectation {
	info, err := f.d.Stat()
	if err != nil {
		return expect.Faulty(f.T, err)
	}

	return expect.Value(f.T, info.Size())
}

// Text returns an Expectation about the entire file contents as text
func (f *File) Text() expect.Expectation {
	if err := f.readAll(); err != nil {
		return expect.Faulty(f.T, err)
	}

	return expect.Value(f.T, string(f.buf))
}

// Bytes returns an Expectation about the entire file contents
func (f *File) Bytes() expect.Expectation {
	if err := f.readAll(); err != nil {
		return expect.Faulty(f.T, err)
	}

	return expect.Value(f.T, f.buf)
}

// ContentType returns an Expectation about the content type
func (f *File) ContentType() expect.Expectation {
	if f.pos < 1 {
		err := f.read(true)
		if err != nil && !errors.Is(err, io.EOF) {
			return expect.Faulty(f.T, err)
		}
	}
	var contentType = http.DetectContentType(f.buf)

	fmt.Println(contentType)

	return expect.Value(f.T, contentType)
}

func (f *File) seek(pos int64) error {
	// cannot seek in append mode or named pipes
	if isAppend(f.mod) || isPipe(f.mod) {
		return nil
	}

	ret, err := f.d.Seek(pos, io.SeekStart)
	f.pos = ret

	return err
}

func (f *File) read(overwrite bool) error {
	start := 0

	if !overwrite {
		// next read will append to end of buffer
		start = len(f.buf)
		// add more capacity as needed
		if len(f.buf) == cap(f.buf) {
			f.buf = append(f.buf, 0)[:len(f.buf)]
		}
	}

	n, err := f.d.Read(f.buf[start:cap(f.buf)])
	f.buf = f.buf[:start+n]
	f.pos += int64(n)

	return err
}

func (f *File) readAll() error {
	err := f.seek(0)
	if err != nil {
		return err
	}

	for {
		err := f.read(false)
		if errors.Is(io.EOF, err) {
			break
		} else if err != nil {
			return err
		}
	}

	return nil
}
