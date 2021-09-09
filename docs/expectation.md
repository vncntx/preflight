# Expectations

An [**Expectation**](https://pkg.go.dev/vincent.click/pkg/preflight/expect#Expectation) provides a common interface for making assertions about values and behaviors.

## Equality

Assert equality using the `Eq`, `Equal`, `Equals`, and `EqualTo` methods. Arrays, slices, and maps are considered equal if all of their elements are equal. Otherwise, the normal rules of equality in Go apply.

```go
func TestEquals(test *testing.T) {
    t := preflight.Unit(test)

    t.Expect(2 * 5).Equals(25)
}
```

There are also shortcut methods for checking if a value is `Nil`, `True`, and `False`.

## Length

Assert the length of an array, slice, map, channel, or string using the `HasLength` and `HaveLength` methods.

```go
func TestLength(test *testing.T) {
    t := preflight.Unit(test)

    list := []int{2, 5, 7}
    t.Expect(list).HasLength(3)
}
```

The `Empty` method asserts a length of zero.

## Pattern Matching

Assert if a value matches a given regular expression with the `Match` and `Matches` methods. If available, these methods use the representation provided by types that can [describe themselves as a string](https://tour.golang.org/methods/17).

```go
func TestPattern(test *testing.T) {
    t := preflight.Unit(test)

    t.Expect(5).Matches("[0-9]+")
}
```

## Elements

Given an array, slice, or map, the `At` method returns a new expectation about the element at the given index or key. Given a string, it returns a new expectation about the [rune](https://golangdocs.com/rune-in-golang) at the given index.

```go
func TestElement(test *testing.T) {
    t := preflight.Unit(test)

    list := []int{2, 5, 7}
    t.Expect(list).At(2).Equals(7)
}
```

## Negation

The methods `Not`, `IsNot`, and `DoesNot` negate the current expectation.

```go
func TestNegation(test *testing.T) {
    t := preflight.Unit(test)

    t.Expect(2 * 5).IsNot().EqualTo(10)
}
```

## Sugar

Some methods don't do anything and simply return the current expectation. These include `To`, `Be`, `Is`, and `Should` and can be used to make your tests easier to read.