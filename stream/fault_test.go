package stream_test

import (
	"errors"
	"testing"

	"vincent.click/pkg/preflight"
	"vincent.click/pkg/preflight/expect"
	"vincent.click/pkg/preflight/stream"
)

func TestFaultyClose(test *testing.T) {
	t := preflight.Unit(test)

	f := stream.Faulty(t.T, errors.New("err"))

	err := f.Close()
	t.Expect(err).Is().Nil()
}

func TestFaultyExpectations(test *testing.T) {
	t := preflight.Unit(test)

	f := stream.Faulty(t.T, errors.New("err"))

	t.Expect(isFaulty(f.Size())).Is().True()
	t.Expect(isFaulty(f.Text())).Is().True()
	t.Expect(isFaulty(f.TextAt(3, 5))).Is().True()
	t.Expect(isFaulty(f.Bytes())).Is().True()
	t.Expect(isFaulty(f.NextBytes(5))).Is().True()
	t.Expect(isFaulty(f.BytesAt(3, 5))).Is().True()
	t.Expect(isFaulty(f.ContentType())).Is().True()
}

func isFaulty(e expect.Expectation) bool {
	_, ok := e.(*expect.Fault)

	return ok
}
