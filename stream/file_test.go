package stream_test

import (
	"io"
	"io/ioutil"
	"os"
	"testing"

	"vincent.click/pkg/preflight"
	"vincent.click/pkg/preflight/stream"
)

var content = inflate("Ad astra per aspera.\n")

func TestFileClose(test *testing.T) {
	t := preflight.Unit(test)

	temp := createTemp(t, content)
	file := stream.FromFile(t.T, temp)
	defer cleanup(temp)

	t.Expect(file.Close()).Is().Nil()
}

func TestFileSize(test *testing.T) {
	t := preflight.Unit(test)

	temp := createTemp(t, content)
	file := stream.FromFile(t.T, temp)
	defer cleanup(temp)

	size := len(content)
	file.Size().Eq(size)
}

func TestFileText(test *testing.T) {
	t := preflight.Unit(test)

	temp := createTemp(t, content)
	file := stream.FromFile(t.T, temp)
	defer cleanup(temp)

	file.Text().Eq(content)
}

func TestFileNextText(test *testing.T) {
	t := preflight.Unit(test)

	temp := createTemp(t, content)
	file := stream.FromFile(t.T, temp)
	defer cleanup(temp)

	file.NextText(3).Eq("Ad ")
	file.NextText(5).Eq("astra")
}

func TestFileTextAt(test *testing.T) {
	t := preflight.Unit(test)

	temp := createTemp(t, content)
	file := stream.FromFile(t.T, temp)
	defer cleanup(temp)

	file.TextAt(3, 5).Eq("astra")
}

func TestFileBytes(test *testing.T) {
	t := preflight.Unit(test)

	temp := createTemp(t, content)
	file := stream.FromFile(t.T, temp)
	defer cleanup(temp)

	bytes := []byte(content)
	file.Bytes().Eq(bytes)
}

func TestFileNextBytes(test *testing.T) {
	t := preflight.Unit(test)

	temp := createTemp(t, content)
	file := stream.FromFile(t.T, temp)
	defer cleanup(temp)

	file.NextBytes(3).Eq([]byte("Ad "))
	file.NextBytes(5).Eq([]byte("astra"))
}

func TestFileBytesAt(test *testing.T) {
	t := preflight.Unit(test)

	temp := createTemp(t, content)
	file := stream.FromFile(t.T, temp)
	defer cleanup(temp)

	bytes := []byte("astra")
	file.BytesAt(3, 5).Eq(bytes)
}

func TestFileNextLine(test *testing.T) {
	t := preflight.Unit(test)

	temp := createTemp(t, content)
	file := stream.FromFile(t.T, temp)
	defer cleanup(temp)

	file.NextLine().Eq("Ad astra per aspera.")
	file.NextLine().Eq("Ad astra per aspera.")
}

func TestFileContentType(test *testing.T) {
	t := preflight.Unit(test)

	temp := createTemp(t, content)
	file := stream.FromFile(t.T, temp)
	defer cleanup(temp)

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
	// go back to the beginning to prepare for reads
	if _, err := file.Seek(0, io.SeekStart); err != nil {
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

func cleanup(f *os.File) {
	_ = f.Close()

	if err := os.Remove(f.Name()); err != nil {
		panic(err)
	}
}
