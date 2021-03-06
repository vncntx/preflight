package stream

import "vincent.click/pkg/preflight/expect"

// Expectations is a set of expectations about a data stream
type Expectations interface {
	// Close frees up resources
	Close() error

	// Size returns an Expectation about the size of the data in bytes
	Size() expect.Expectation

	// Text returns an Expectation about the data as utf-8 text
	Text() expect.Expectation
	// NextText returns an Expectation about the next chunk of text
	NextText(bytes int) expect.Expectation
	// TextAt returns an Expectation about text starting at the given byte offset
	TextAt(pos int64, bytes int) expect.Expectation

	// Bytes returns an Expectation about the data as a byte array
	Bytes() expect.Expectation
	// NextBytes returns an Expectation about the next chunk of bytes
	NextBytes(bytes int) expect.Expectation
	// BytesAt returns an Expectation about the bytes starting at the given offset
	BytesAt(pos int64, bytes int) expect.Expectation

	// Lines returns an Expectation about the data as lines of text
	Lines() expect.Expectation
	// NextLine returns an Expectation about the next line of text
	NextLine() expect.Expectation

	// ContentType returns an Expectation about the media type.
	// See https://en.wikipedia.org/wiki/Media_type
	ContentType() expect.Expectation
}
