package stream

import (
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
	txt, err := readAll(f.d, f.mod)
	if err != nil {
		return expect.Faulty(f.T, err)
	}

	return expect.Value(f.T, string(txt))
}

// NextText returns an Expectation about the next chunk of text
func (f *File) NextText(length int) expect.Expectation {
	bytes, err := read(f.d, length)
	if err != nil {
		return expect.Faulty(f.T, err)
	}

	return expect.Value(f.T, string(bytes))
}

// TextAt returns an Expectation about the text contents at a specific position
func (f *File) TextAt(pos int64, length int) expect.Expectation {
	bytes, err := readAt(f.d, pos, length)
	if err != nil {
		return expect.Faulty(f.T, err)
	}

	return expect.Value(f.T, string(bytes))
}

// Bytes returns an Expectation about the entire file contents
func (f *File) Bytes() expect.Expectation {
	bytes, err := readAll(f.d, f.mod)
	if err != nil {
		return expect.Faulty(f.T, err)
	}

	return expect.Value(f.T, bytes)
}

// NextBytes returns an Expectation about the next chunk of bytes
func (f *File) NextBytes(length int) expect.Expectation {
	bytes, err := read(f.d, length)
	if err != nil {
		return expect.Faulty(f.T, err)
	}

	return expect.Value(f.T, bytes)
}

// BytesAt returns an Expectation about the file contents at a specific position
func (f *File) BytesAt(pos int64, length int) expect.Expectation {
	bytes, err := readAt(f.d, pos, length)
	if err != nil {
		return expect.Faulty(f.T, err)
	}

	return expect.Value(f.T, bytes)
}

// ContentType returns an Expectation about the content type
func (f *File) ContentType() expect.Expectation {
	content, err := readAt(f.d, 0, 512)
	if err != nil {
		return expect.Faulty(f.T, err)
	}

	contentType := http.DetectContentType(content)

	return expect.Value(f.T, contentType)
}
