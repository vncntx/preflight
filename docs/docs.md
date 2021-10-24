# Documentation

## Expectations

An [**Expectation**](https://pkg.go.dev/vincent.click/pkg/preflight/expect#Expectation) provides a common interface for making assertions about values and behaviors.

### Equality

Assert equality using the `Eq`, `Equal`, `Equals`, and `EqualTo` methods. Arrays, slices, and maps are considered equal if all of their elements are equal. Otherwise, the normal rules of equality in Go apply.

```go
func TestEquals(test *testing.T) {
    t := preflight.Unit(test)

    t.Expect(2 + 5).Equals(7)
}
```

There are also shortcut methods for checking if a value is `Nil`, `True`, and `False`.

### Length

Assert the length of an array, slice, map, channel, or string using the `HasLength` and `HaveLength` methods.

```go
func TestLength(test *testing.T) {
    t := preflight.Unit(test)

    list := []int{2, 5, 7}
    t.Expect(list).HasLength(3)
}
```

The `Empty` method asserts a length of zero.

### Pattern Matching

Assert if a value matches a given regular expression with the `Match` and `Matches` methods. If available, these methods use the representation provided by types that can [describe themselves as a string](https://pkg.go.dev/fmt#Stringer).

```go
func TestPattern(test *testing.T) {
    t := preflight.Unit(test)

    t.Expect(5).Matches("[0-9]+")
}
```

### Elements

Given an array, slice, or map, the `At` method returns a new expectation about the element at the given index or key. Given a string, it returns a new expectation about the [rune](https://golang.org/ref/spec#Rune_literals) at the given index.

```go
func TestElement(test *testing.T) {
    t := preflight.Unit(test)

    list := []int{2, 5, 7}
    t.Expect(list).At(2).Equals(7)
}
```

### Negation

The methods `Not`, `IsNot`, and `DoesNot` negate the current expectation.

```go
func TestNegation(test *testing.T) {
    t := preflight.Unit(test)

    t.Expect(2 * 5).IsNot().EqualTo(10)
}
```

### Sugar

Some methods don't do anything and simply return the current expectation. These include `To`, `Be`, `Is`, and `Should` and can be used to make your tests easier to read.

## Files and Data Streams

Sets of expectations can be created either directly from a file descriptor or from a function that writes to a data stream.

<details>
<summary>File Descriptor</summary>


```go
func TestFile(test *testing.T) {
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
func TestWritable(test *testing.T) {
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

`Size` returns an expectation about the size of the contents in bytes.

`ContentType` returns an expectation about the media type, inferred either from the file extension or from part of the contents.

`Text` returns an expectation about the entire contents, interpreted as UTF-8 encoded text.

`NextText` returns an expectation about the next chunk of text with the given bytelength.

`TextAt` returns an expectation about the chunk of text starting at the given byte offset.

`Bytes` returns an expectation about the entire contents as a byte array.

`NextBytes` returns an expectation about the next chunk of bytes with the given bytelength.

`BytesAt` returns an expectation about the data at the given byte offset.

`Lines` returns an expectation about the entire contents as an array of lines, excluding the newline characters.

`NextLine` returns an expectation about the next line of text.

## Captor

Capture calls to builtin functions and the standard library by using **Captor**. This provides a collection of aliases that can be used to intercept arguments during tests.

`Exit` is an alias for `os.Exit`. Capture the exit code using `ExpectExitCode`.

`Panic` is an alias for the builtin `panic` function. Capture the cause of a panicking goroutine using `ExpectPanic`
