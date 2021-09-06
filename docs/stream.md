# Streams

A [**Stream**](https://pkg.go.dev/vincent.click/pkg/preflight/stream#Stream) is a set of [expectations](./expectation.md) about a file or data stream. The test is responsible for closing the stream by calling the `Close` method. It can be created either directly from a file descriptor or from a function that writes to a data stream.

<details>
<summary>File Descriptor</summary>


```go
func TestStreams(test *testing.T) {
    t := preflight.Unit(test)

    file, err := os.Open("file.txt")
    if err != nil {
        panic(err)
    }

    f := t.ExpectFile(file)
    defer f.Close()

    f.ContentType().Equals("text/plain")
}
```

</details>

<details>
<summary>Function</summary>


```go
func TestStreams(test *testing.T) {
    t := preflight.Unit(test)

    f := t.ExpectWritten(func (w *os.File) {
        if _, err := w.Write("text"); err != nil {
            panic(err)
        }
    })
    defer f.Close()

    f.ContentType().Equals("text/plain")
}
```

</details>

## Contents

`Size` returns an expectation about the size of the contents in bytes.

`ContentType` returns an expectation about the media type, inferred either from the file extension or from part of the contents.

### Text

These methods treat the data as UTF-8 encoded text.

`Text` returns an expectation about the entire contents.

`NextText` returns an expectation about the next chunk of text with the given bytelength.

`TextAt` returns an expectation about the chunk of text starting at the given byte offset.

### Bytes

These methods treat the data as an array of bytes.

`Bytes` returns an expectation about the entire contents as a byte array.

`NextBytes` returns an expectation about the next chunk of bytes with the given bytelength.

`BytesAt` returns an expectation about the data at the given byte offset.

### Lines

These methods interpret the data as lines of text, excluding the newline characters.

`Lines` returns an expectation about the entire contents as an array of lines.

`NextLine` returns an expectation about the next line of text.