package stream

import (
	"errors"
	"io"
	"net/http"
	"os"
	"unicode"
	"unicode/utf8"
)

type reader interface {
	io.ReaderAt
	io.ReadCloser

	Name() string
	Stat() (os.FileInfo, error)
}

type rule func(rune) bool

// Set of line-terminating characters
// https://en.wikipedia.org/wiki/Newline#Unicode
var eol = &unicode.RangeTable{
	R32: []unicode.Range32{
		{0x000a, 0x000d, 1},
		{0x0085, 0x0085, 1},
		{0x2028, 0x2029, 1},
	},
}

func read(src io.Reader, max int) ([]byte, error) {
	buf := make([]byte, max)

	size, err := src.Read(buf)
	if err != nil && !errors.Is(err, io.EOF) {
		return nil, err
	}

	return buf[:size], nil
}

func readAt(src io.ReaderAt, pos int64, max int) ([]byte, error) {
	buf := make([]byte, max)

	size, err := src.ReadAt(buf, pos)
	if err != nil && !errors.Is(err, io.EOF) {
		return nil, err
	}

	return buf[:size], nil
}

func readAll(src io.Reader) ([]byte, error) {
	buf := make([]byte, 0, blocksize)

	for {
		// next read will append to end of buffer
		start := len(buf)
		// add more capacity as needed
		if len(buf) == cap(buf) {
			buf = append(buf, 0)[:len(buf)]
		}

		n, err := src.Read(buf[start:cap(buf)])
		if errors.Is(err, io.EOF) {
			break
		} else if err != nil {
			return nil, err
		}

		buf = buf[:start+n]
	}

	return buf, nil
}

func readRunes(src io.Reader, bytes []byte, accept rule) ([]rune, []byte, error) {
	runes := []rune{}
	dry := len(bytes) < 1 // whether the buffer needs to be replenished

	for {
		if dry {
			// read the next block into the buffer
			next, err := read(src, blocksize)
			if err != nil {
				return nil, bytes, err
			}
			bytes = append(bytes, next...)

		}

		if len(bytes) < 1 {
			break
		}

		r, size := utf8.DecodeRune(bytes)
		dry = (r == utf8.RuneError || size < 1)

		if dry {
			continue
		}

		if !accept(r) {
			break
		}

		runes = append(runes, r)
		bytes = bytes[size:]
	}

	return runes, bytes, nil
}

func readLine(src io.Reader, bytes []byte) ([]rune, []byte, error) {
	line, bytes, err := readRunes(src, bytes, func(r rune) bool {
		return !unicode.In(r, eol)
	})
	if err != nil {
		return nil, nil, err
	}

	_, bytes, err = readRunes(src, bytes, func(r rune) bool {
		return unicode.In(r, eol)
	})
	if err != nil {
		return nil, nil, err
	}

	return line, bytes, err
}

func detectContentType(src io.ReaderAt) (string, error) {
	content, err := readAt(src, 0, 512)
	if err != nil {
		return "", err
	}

	return http.DetectContentType(content), nil
}
