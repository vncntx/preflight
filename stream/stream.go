package stream

import (
	"vincent.click/pkg/preflight/expect"
)

// Stream is a set of expectations about a data stream
type Stream interface {
	Close() error
	Size() expect.Expectation
	Text() expect.Expectation
	Bytes() expect.Expectation
	ContentType() expect.Expectation
}
