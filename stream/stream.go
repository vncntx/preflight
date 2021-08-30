package stream

import "vincent.click/pkg/preflight/expect"

// Stream is a set of expectations about a data stream
type Stream interface {
	// Close releases the underlying resources
	Close() error
	// Size returns an Expectation about the size of data in bytes
	Size() expect.Expectation
	// Text returns an Expectation about the data as text
	Text() expect.Expectation
	// NextText returns an Expectation about the next chunk of text
	NextText(bytes int) expect.Expectation
	// TextAt returns an Expectation about text at a specific position
	TextAt(pos int64, bytes int) expect.Expectation
	// Bytes returns an Expectation about the data as a byte array
	Bytes() expect.Expectation
	// NextBytes returns an Expectation about the next chunk of bytes
	NextBytes(bytes int) expect.Expectation
	// BytesAt returns an Expectation about the bytes at a specific position
	BytesAt(pos int64, bytes int) expect.Expectation
	// NextLine returns an Expectation about the next line of text
	NextLine() expect.Expectation
	// ContentType returns an Expectation about the media type.
	// See https://en.wikipedia.org/wiki/Media_type
	ContentType() expect.Expectation
}
