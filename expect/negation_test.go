package expect_test

import (
	"testing"

	"vincent.click/pkg/preflight"
	"vincent.click/pkg/preflight/expect"
)

func TestNegation(test *testing.T) {
	t := preflight.Unit(test)

	not := expect.Value(t.T, 1).Not()

	t.Expect(not.To()).Eq(not)
	t.Expect(not.Be()).Eq(not)
	t.Expect(not.Is()).Eq(not)
	t.Expect(not.Should()).Eq(not)
}

func TestNegationIsNot(test *testing.T) {
	t := preflight.Unit(test)

	v := expect.Value(t.T, 1)
	not := v.Not()

	t.Expect(not.Not()).Eq(v)
	t.Expect(not.IsNot()).Eq(v)
	t.Expect(not.DoesNot()).Eq(v)
}

func TestNegationAt(test *testing.T) {
	t := preflight.Unit(test)

	v1 := expect.Value(t.T, "xyz")
	v1.Not().At(2).Eq('x')

	v2 := expect.Value(t.T, []int{5, 10, 15})
	v2.Not().At(2).Eq(10)

	v3 := expect.Value(t.T, [3]int{5, 10, 15})
	v3.Not().At(2).Eq(10)

	v4 := expect.Value(t.T, map[int]int{
		1: 5,
		2: 10,
		3: 15,
	})
	v4.Not().At(3).Eq(10)
}

func TestNegationIsNil(test *testing.T) {
	t := preflight.Unit(test)

	v := expect.Value(t.T, 1)

	v.IsNot().Nil()
}

func TestNegationIsTrue(test *testing.T) {
	t := preflight.Unit(test)

	v := expect.Value(t.T, false)

	v.IsNot().True()
}

func TestNegationIsFalse(test *testing.T) {
	t := preflight.Unit(test)

	v := expect.Value(t.T, true)

	v.IsNot().False()
}

func TestNegationIsEmpty(test *testing.T) {
	t := preflight.Unit(test)

	v1 := expect.Value(t.T, []int{1})
	v2 := expect.Value(t.T, map[int]int{2: 3})

	v1.IsNot().Empty()
	v2.IsNot().Empty()
}

func TestNegationArrayHasLength(test *testing.T) {
	t := preflight.Unit(test)

	v := expect.Value(t.T, []int{3, 4, 3})

	v.Not().HasLength(2)
	v.DoesNot().HaveLength(2)
}

func TestNegationMapHasLength(test *testing.T) {
	t := preflight.Unit(test)

	v := expect.Value(t.T, map[int]int{
		1: 3,
		2: 4,
		3: 3,
	})

	v.Not().HasLength(2)
	v.DoesNot().HaveLength(2)
}

func TestNegationEquals(test *testing.T) {
	t := preflight.Unit(test)

	v := expect.Value(t.T, 343)

	v.Not().Eq(23)
	v.Not().Equal(23)
	v.Not().Equals(23)
	v.Not().EqualTo(23)
}

func TestNegationEqualsArray(test *testing.T) {
	t := preflight.Unit(test)

	list1 := []int{2, 3}
	list2 := []int{3, 4, 3}
	v := expect.Value(t.T, list1)

	v.Not().Eq(list2)
	v.Not().Equal(list2)
	v.Not().Equals(list2)
	v.Not().EqualTo(list2)
}

func TestNegationEqualsMap(test *testing.T) {
	t := preflight.Unit(test)

	map1 := map[int]int{
		1: 2,
		2: 3,
	}
	map2 := map[int]int{
		1: 3,
		2: 4,
		3: 3,
	}
	v := expect.Value(t.T, map1)

	v.Not().Eq(map2)
	v.Not().Equal(map2)
	v.Not().Equals(map2)
	v.Not().EqualTo(map2)
}

func TestNegationMatches(test *testing.T) {
	t := preflight.Unit(test)

	v := expect.Value(t.T, "343")

	v.DoesNot().Match("[a-z]+")
	v.Not().Matches("[a-z]+")
}
