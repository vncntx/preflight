package expect_test

import (
	"preflight"
	"preflight/expect"
	"testing"
)

func TestValue(test *testing.T) {
	t := preflight.Unit(test)

	v := expect.Value(t.T, 1)

	t.Expect(v.To()).Eq(v)
	t.Expect(v.Be()).Eq(v)
	t.Expect(v.Is()).Eq(v)
	t.Expect(v.Should()).Eq(v)
}

func TestValueIsNot(test *testing.T) {
	t := preflight.Unit(test)

	v := expect.Value(t.T, 1)

	if _, ok := v.Not().(*expect.Negation); !ok {
		t.Error("Not does not return a negation")
	}
	if _, ok := v.IsNot().(*expect.Negation); !ok {
		t.Error("IsNot does not return a negation")
	}
}

func TestValueIsNil(test *testing.T) {
	t := preflight.Unit(test)

	v := expect.Value(t.T, nil)

	v.Nil()
}

func TestValueIsTrue(test *testing.T) {
	t := preflight.Unit(test)

	v := expect.Value(t.T, true)

	v.True()
}

func TestValueIsFalse(test *testing.T) {
	t := preflight.Unit(test)

	v := expect.Value(t.T, false)

	v.False()
}

func TestValueIsEmpty(test *testing.T) {
	t := preflight.Unit(test)

	v1 := expect.Value(t.T, []int{})
	v2 := expect.Value(t.T, map[int]int{})

	v1.Empty()
	v2.Empty()
}

func TestValueArrayHasLength(test *testing.T) {
	t := preflight.Unit(test)

	v := expect.Value(t.T, []int{3, 4, 3})

	v.HasLength(3)
	v.HaveLength(3)
}

func TestValueMapHasLength(test *testing.T) {
	t := preflight.Unit(test)

	v := expect.Value(t.T, map[int]int{
		1: 3,
		2: 4,
		3: 3,
	})

	v.HasLength(3)
	v.HaveLength(3)
}

func TestValueEquals(test *testing.T) {
	t := preflight.Unit(test)

	expected := 343
	v := expect.Value(t.T, expected)

	v.Eq(expected)
	v.Equal(expected)
	v.Equals(expected)
	v.EqualTo(expected)
}

func TestValueEqualsArray(test *testing.T) {
	t := preflight.Unit(test)

	expected := []int{3, 4, 3}
	v := expect.Value(t.T, expected)

	v.Eq(expected)
	v.Equal(expected)
	v.Equals(expected)
	v.EqualTo(expected)
}

func TestValueEqualsMap(test *testing.T) {
	t := preflight.Unit(test)

	expected := map[int]int{
		1: 3,
		2: 4,
		3: 3,
	}
	v := expect.Value(t.T, expected)

	v.Eq(expected)
	v.Equal(expected)
	v.Equals(expected)
	v.EqualTo(expected)
}

func TestValueMatches(test *testing.T) {
	t := preflight.Unit(test)

	v := expect.Value(t.T, "343")

	v.Match("\\d+")
	v.Matches("\\d+")
}
