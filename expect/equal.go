package expect

import (
	"reflect"

	"vincent.click/pkg/preflight/expect/kind"
)

func equal(a, b interface{}) bool {
	x := reflect.ValueOf(a)
	y := reflect.ValueOf(b)

	aK := kind.Of(a)
	bK := kind.Of(b)

	if aK != bK {
		return false
	}

	switch aK {
	case kind.Bool:

		return x.Bool() == y.Bool()

	case kind.Int:

		return x.Int() == y.Int()

	case kind.Uint:

		return x.Uint() == y.Uint()

	case kind.Real:

		return x.Float() == y.Float()

	case kind.Complex:

		return x.Complex() == y.Complex()

	case kind.List:

		return listsEqual(x, y)

	case kind.Map:

		return mapsEqual(x, y)

	default:
		return a == b
	}
}

func listsEqual(x, y reflect.Value) bool {
	if x.Len() != y.Len() {
		return false
	}

	for i := 0; i < x.Len(); i++ {
		xitem := x.Index(i).Interface()
		yitem := y.Index(i).Interface()

		if !equal(xitem, yitem) {
			return false
		}
	}

	return true
}

func mapsEqual(x, y reflect.Value) bool {
	if x.Len() != y.Len() {
		return false
	}

	for _, key := range x.MapKeys() {
		xitem := x.MapIndex(key).Interface()
		yitem := y.MapIndex(key).Interface()

		if !equal(xitem, yitem) {
			return false
		}
	}

	return true
}
