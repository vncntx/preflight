package kind_test

import (
	"testing"

	"preflight"
	"preflight/expect/kind"
)

func TestOf(test *testing.T) {
	t := preflight.Unit(test)

	t.Expect(kind.Of(nil)).Eq(kind.Nil)

	t.Expect(kind.Of(true)).Eq(kind.Bool)
	t.Expect(kind.Of(false)).Eq(kind.Bool)

	t.Expect(kind.Of(10)).Eq(kind.Int)
	t.Expect(kind.Of(int8(8))).Eq(kind.Int)
	t.Expect(kind.Of(int16(16))).Eq(kind.Int)
	t.Expect(kind.Of(int32(32))).Eq(kind.Int)
	t.Expect(kind.Of(int64(64))).Eq(kind.Int)

	t.Expect(kind.Of(uint(10))).Eq(kind.Uint)
	t.Expect(kind.Of(uint8(8))).Eq(kind.Uint)
	t.Expect(kind.Of(uint16(16))).Eq(kind.Uint)
	t.Expect(kind.Of(uint32(32))).Eq(kind.Uint)
	t.Expect(kind.Of(uint64(64))).Eq(kind.Uint)

	t.Expect(kind.Of(float32(3.14))).Eq(kind.Real)
	t.Expect(kind.Of(float64(3.14))).Eq(kind.Real)

	t.Expect(kind.Of(complex64(3 + 43i))).Eq(kind.Complex)
	t.Expect(kind.Of(complex128(3 + 43i))).Eq(kind.Complex)

	t.Expect(kind.Of([]int{1, 2, 3})).Eq(kind.List)
	t.Expect(kind.Of(make([]int, 3))).Eq(kind.List)

	t.Expect(kind.Of(make(map[int]int))).Eq(kind.Map)

	t.Expect(kind.Of(t)).Eq(kind.Std)
	t.Expect(kind.Of("text")).Eq(kind.Std)
	t.Expect(kind.Of(struct{}{})).Eq(kind.Std)
}
