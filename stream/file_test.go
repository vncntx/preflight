package stream_test

import (
	"io"
	"io/ioutil"
	"os"
	"testing"

	"vincent.click/pkg/preflight"
	"vincent.click/pkg/preflight/stream"
)

var contents = []byte("Ad astra per aspera.\nSic itur ad astra.\n")

func TestFileClose(test *testing.T) {
	t := preflight.Unit(test)

	temp := createTemp(t, contents)
	file := stream.FromFile(t.T, temp)
	defer rm(temp)

	t.Expect(file.Close()).Is().Nil()
}

func TestFileSize(test *testing.T) {
	t := preflight.Unit(test)

	temp := createTemp(t, contents)
	file := stream.FromFile(t.T, temp)
	defer cleanup(temp)

	size := len(contents)
	file.Size().Eq(size)
}

func TestFileText(test *testing.T) {
	t := preflight.Unit(test)

	temp := createTemp(t, contents)
	file := stream.FromFile(t.T, temp)
	defer cleanup(temp)

	text := string(contents)
	file.Text().Eq(text)
}

func TestFileNextText(test *testing.T) {
	t := preflight.Unit(test)

	temp := createTemp(t, contents)
	file := stream.FromFile(t.T, temp)
	defer cleanup(temp)

	file.NextText(3).Eq("Ad ")
	file.NextText(5).Eq("astra")
}

func TestFileTextAt(test *testing.T) {
	t := preflight.Unit(test)

	temp := createTemp(t, contents)
	file := stream.FromFile(t.T, temp)
	defer cleanup(temp)

	file.TextAt(3, 5).Eq("astra")
}

func TestFileBytes(test *testing.T) {
	t := preflight.Unit(test)

	temp := createTemp(t, contents)
	file := stream.FromFile(t.T, temp)
	defer cleanup(temp)

	file.Bytes().Eq(contents)
}

func TestFileNextBytes(test *testing.T) {
	t := preflight.Unit(test)

	temp := createTemp(t, contents)
	file := stream.FromFile(t.T, temp)
	defer cleanup(temp)

	file.NextBytes(3).Eq([]byte("Ad "))
	file.NextBytes(5).Eq([]byte("astra"))
}

func TestFileBytesAt(test *testing.T) {
	t := preflight.Unit(test)

	temp := createTemp(t, contents)
	file := stream.FromFile(t.T, temp)
	defer cleanup(temp)

	bytes := []byte("astra")
	file.BytesAt(3, 5).Eq(bytes)
}

func TestFileLines(test *testing.T) {
	t := preflight.Unit(test)

	temp := createTemp(t, contents)
	file := stream.FromFile(t.T, temp)
	defer cleanup(temp)

	file.Lines().HasLength(2)
}

func TestFileNextLine(test *testing.T) {
	t := preflight.Unit(test)

	temp := createTemp(t, contents)
	file := stream.FromFile(t.T, temp)
	defer cleanup(temp)

	file.NextLine().Eq("Ad astra per aspera.")
	file.NextLine().Eq("Sic itur ad astra.")
	file.NextLine().Is().Empty()
}

func TestFileContentType(test *testing.T) {
	t := preflight.Unit(test)

	temp := createTemp(t, contents)
	file := stream.FromFile(t.T, temp)
	defer cleanup(temp)

	file.ContentType().Matches("text/plain")
}

func createTemp(t *preflight.Test, content []byte) *os.File {
	file, err := ioutil.TempFile(os.TempDir(), "preflight-")
	if err != nil {
		t.Error(err)

		return nil
	}

	if _, err := file.Write(content); err != nil {
		t.Error(err)

		return nil
	}

	if _, err := file.Seek(0, io.SeekStart); err != nil {
		t.Error(err)

		return nil
	}

	return file
}

func cleanup(f *os.File) {
	if err := f.Close(); err != nil {
		panic(err)
	}

	rm(f)
}

func rm(f *os.File) {
	if err := os.Remove(f.Name()); err != nil {
		panic(err)
	}
}
