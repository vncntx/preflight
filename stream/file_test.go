package stream_test

import (
	"io/ioutil"
	"os"
	"testing"

	"vincent.click/pkg/preflight"
	"vincent.click/pkg/preflight/stream"
)

var content = inflate("Ad astra per aspera.")

func TestFileClose(test *testing.T) {
	t := preflight.Unit(test)

	temp := createTemp(t, content)
	file := stream.FromFile(t.T, temp)

	t.Expect(file.Close()).Is().Nil()
}

func TestFileSize(test *testing.T) {
	t := preflight.Unit(test)

	temp := createTemp(t, content)
	file := stream.FromFile(t.T, temp)
	defer file.Close()

	size := len(content)
	file.Size().Eq(size)
}

func TestFileText(test *testing.T) {
	t := preflight.Unit(test)

	temp := createTemp(t, content)
	file := stream.FromFile(t.T, temp)
	defer file.Close()

	file.Text().Eq(content)
}

func TestFileBytes(test *testing.T) {
	t := preflight.Unit(test)

	temp := createTemp(t, content)
	file := stream.FromFile(t.T, temp)
	defer file.Close()

	bytes := []byte(content)
	file.Bytes().Eq(bytes)
}

func TestFileContentType(test *testing.T) {
	t := preflight.Unit(test)

	temp := createTemp(t, content)
	file := stream.FromFile(t.T, temp)
	defer file.Close()

	file.ContentType().Matches("text/plain")
}

func createTemp(t *preflight.Test, content string) *os.File {
	file, err := ioutil.TempFile(os.TempDir(), "preflight-")
	if err != nil {
		t.Error(err)
	}
	if _, err := file.WriteString(content); err != nil {
		t.Error(err)
	}

	return file
}

// inflate content so that tests cover buffers and seeks
func inflate(content string) string {
	result := ""
	for i := 0; i < 10_000; i++ {
		result += content
	}

	return result
}
