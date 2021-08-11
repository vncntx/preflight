package stream

import "vincent.click/pkg/preflight/expect"

// Stream is a set of expectations about a data stream
type Stream interface {
	Close() error
	Size() expect.Expectation
	Text() expect.Expectation
	TextAt(pos int64, length int) expect.Expectation
	Bytes() expect.Expectation
	NextBytes(length int) expect.Expectation
	BytesAt(pos int64, length int) expect.Expectation
	ContentType() expect.Expectation
}
