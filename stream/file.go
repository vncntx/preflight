package stream

import (
	"mime"
	"os"
	"path/filepath"
	"testing"

	"vincent.click/pkg/preflight/expect"
)

// File is a set of expectations about a file stream
type File struct {
	*testing.T

	d *os.File
	b []byte
}

// FromFile returns expectations based on a file descriptor
func FromFile(t *testing.T, file *os.File) Stream {
	return &File{
		T: t,
		d: file,
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
	txt, err := readAll(f.d)
	if err != nil {
		return expect.Faulty(f.T, err)
	}

	return expect.Value(f.T, string(txt))
}

// NextText returns an Expectation about the next chunk of text
func (f *File) NextText(bytes int) expect.Expectation {
	data, err := read(f.d, bytes)
	if err != nil {
		return expect.Faulty(f.T, err)
	}

	return expect.Value(f.T, string(data))
}

// TextAt returns an Expectation about the text contents at a specific position
func (f *File) TextAt(pos int64, bytes int) expect.Expectation {
	data, err := readAt(f.d, pos, bytes)
	if err != nil {
		return expect.Faulty(f.T, err)
	}

	return expect.Value(f.T, string(data))
}

// Bytes returns an Expectation about the entire file contents
func (f *File) Bytes() expect.Expectation {
	bytes, err := readAll(f.d)
	if err != nil {
		return expect.Faulty(f.T, err)
	}

	return expect.Value(f.T, bytes)
}

// NextBytes returns an Expectation about the next chunk of bytes
func (f *File) NextBytes(bytes int) expect.Expectation {
	data, err := read(f.d, bytes)
	if err != nil {
		return expect.Faulty(f.T, err)
	}

	return expect.Value(f.T, data)
}

// BytesAt returns an Expectation about the file contents at a specific position
func (f *File) BytesAt(pos int64, bytes int) expect.Expectation {
	data, err := readAt(f.d, pos, bytes)
	if err != nil {
		return expect.Faulty(f.T, err)
	}

	return expect.Value(f.T, data)
}

// Lines returns an Expectation about the entire file contents as lines of text
func (f *File) Lines() expect.Expectation {
	lines := []string{}

	for {
		line, bytes, err := readLine(f.d, f.b)
		if err != nil {
			return expect.Faulty(f.T, err)
		}

		f.b = bytes

		if len(line) < 1 {
			break
		}

		lines = append(lines, string(line))
	}

	return expect.Value(f.T, lines)
}

// NextLine returns an Expectation about the next line of text
func (f *File) NextLine() expect.Expectation {
	line, bytes, err := readLine(f.d, f.b)
	if err != nil {
		return expect.Faulty(f.T, err)
	}

	f.b = bytes

	return expect.Value(f.T, string(line))
}

// ContentType returns an Expectation about the content type
func (f *File) ContentType() expect.Expectation {
	ext := filepath.Ext(f.d.Name())
	if len(ext) > 0 {
		typ := mime.TypeByExtension(ext)

		return expect.Value(f.T, typ)
	}

	if typ, err := detectContentType(f.d); err != nil {
		return expect.Faulty(f.T, err)
	} else {
		return expect.Value(f.T, typ)
	}
}
