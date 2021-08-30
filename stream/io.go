package stream

import (
	"errors"
	"io"
	"net/http"
	"os"
	"unicode"
	"unicode/utf8"
)

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

func read(f *os.File, max int) ([]byte, error) {
	buf := make([]byte, max)

	size, err := f.Read(buf)
	if err != nil && !errors.Is(err, io.EOF) {
		return nil, err
	}

	return buf[:size], nil
}

func readAt(f *os.File, pos int64, max int) ([]byte, error) {
	buf := make([]byte, max)

	size, err := f.ReadAt(buf, pos)
	if err != nil && !errors.Is(err, io.EOF) {
		return nil, err
	}

	return buf[:size], nil
}

func readAll(f *os.File) ([]byte, error) {
	buf := make([]byte, 0, blocksize)

	for {
		// next read will append to end of buffer
		start := len(buf)
		// add more capacity as needed
		if len(buf) == cap(buf) {
			buf = append(buf, 0)[:len(buf)]
		}

		n, err := f.Read(buf[start:cap(buf)])
		if errors.Is(err, io.EOF) {
			break
		} else if err != nil {
			return nil, err
		}

		buf = buf[:start+n]
	}

	return buf, nil
}

func readRunes(f *os.File, bytes []byte, accept rule) ([]rune, []byte, error) {
	runes := []rune{}
	dry := len(bytes) < 1 // whether the buffer needs to be replenished

	for {
		if dry {
			// read the next block into the buffer
			next, err := read(f, blocksize)
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

func detectContentType(f *os.File) (string, error) {
	content, err := readAt(f, 0, 512)
	if err != nil {
		return "", err
	}

	return http.DetectContentType(content), nil
}
