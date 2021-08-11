package stream

import (
	"errors"
	"io"
	"io/fs"
	"os"
)

func seek(f *os.File, mod fs.FileMode, pos int64) error {
	// cannot seek in append mode or named pipes
	if isAppend(mod) || isPipe(mod) {
		return nil
	}

	_, err := f.Seek(pos, io.SeekStart)

	return err
}

func read(f *os.File, n int) ([]byte, error) {
	buf := make([]byte, n)

	_, err := f.Read(buf)
	if err != nil && !errors.Is(err, io.EOF) {
		return nil, err
	}

	return buf, nil
}

func readAll(f *os.File, mod fs.FileMode) ([]byte, error) {
	buf := make([]byte, 0, bufferSize)

	err := seek(f, mod, 0)
	if err != nil {
		return nil, err
	}

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
