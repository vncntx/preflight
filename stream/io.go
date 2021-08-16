package stream

import (
	"errors"
	"io"
	"io/fs"
	"net/http"
	"os"
)

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

func readAll(f *os.File, mod fs.FileMode) ([]byte, error) {
	buf := make([]byte, 0, bufferSize)

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

func detectContentType(f *os.File) (string, error) {
	content, err := readAt(f, 0, 512)
	if err != nil {
		return "", err
	}

	return http.DetectContentType(content), nil
}
