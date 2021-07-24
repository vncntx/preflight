package kind

import "reflect"

// Kind specifies how equality should be determined
type Kind uint

const (
	Std = Kind(iota)
	Nil
	Bool
	Int
	Uint
	Real
	Complex
	List
	Map
)

// Of returns a value's equality Kind
func Of(x interface{}) Kind {
	if x == nil {
		return Nil
	}

	k := reflect.TypeOf(x).Kind()

	switch k {
	case reflect.Bool:

		return Bool

	case reflect.Int,
		reflect.Int8,
		reflect.Int16,
		reflect.Int32,
		reflect.Int64:

		return Int

	case reflect.Uint,
		reflect.Uint8,
		reflect.Uint16,
		reflect.Uint32,
		reflect.Uint64:

		return Uint

	case reflect.Float32,
		reflect.Float64:

		return Real

	case reflect.Complex64,
		reflect.Complex128:

		return Complex

	case reflect.Array,
		reflect.Slice:

		return List

	case reflect.Map:

		return Map

	default:
		return Std
	}
}
