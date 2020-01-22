package wish

import (
	"reflect"

	"github.com/warpfork/go-wish/cmp"
)

var (
	_ Checker = ShouldBe
	_ Checker = ShouldEqual
)

// ShouldBe asserts that two values are *exactly* the same.
//
// In almost all cases, prefer ShouldEqual.
// ShouldBe differs from ShouldEqual in that it does *not* recurse, and
// thus can be used to explicitly check pointer equality.
//
// For pointers, ShouldBe checks pointer equality.
// Using ShouldBe on any kind of values which require recursion to meaningfully
// compare (e.g., structs, maps, arrays) will be rejected, as will using
// ShouldBe on any kind of value which is already never recursive (e.g. any
// primitives) since you can already use ShouldEqual to compare these.
func ShouldBe(actual interface{}, desire interface{}) (problem string, passed bool) {
	switch reflect.TypeOf(desire).Kind() {
	case // primitives
		reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr,
		reflect.Bool, reflect.Float32, reflect.Float64, reflect.Complex64, reflect.Complex128, reflect.UnsafePointer:
		panic("use ShouldEqual instead of ShouldBe when comparing primitives")
	case // recursives
		reflect.Interface, reflect.Array, reflect.Map, reflect.Slice, reflect.Struct:
		panic("use ShouldEqual instead of ShouldBe when comparing recursive values")
	case reflect.Ptr:
		panic("TODO")
	case reflect.Func:
		panic("TODO")
	case reflect.Chan:
		panic("TODO")
	default:
		panic("unknown kind")
	}
}

// ShouldEqual asserts that two values are the same, examining the values
// recursively as necessary.  Maps, slices, and structs are all
// valid to compare with ShouldEqual.  Pointers will be traversed, and
// comparison continues with the values referenced by the pointer.
func ShouldEqual(actual interface{}, desire interface{}) (diff string, eq bool) {
	s1, ok1 := actual.(string)
	s2, ok2 := desire.(string)
	if ok1 && ok2 {
		diff = strdiff(s1, s2)
	} else {
		diff = cmp.Diff(actual, desire)
	}
	return diff, diff == ""
}

// ShouldBeSameTypeAs asserts that two values have the same concrete type,
// while completely ignoring the contents of the values.
// A nil value of 'actual' with no type is also correctly handled,
// and reports as an error.  (A nil value of actual *with* a type is not given
// special treatment, and may pass the check.)
//
// ShouldBeSameTypeAs is often particularly useful as part of testing errors,
// if your package uses strongly typed errors.  (A combination of a Wish of
// ShouldBeSameTypeAs followed by a Wish of ShouldEqual on the err.Error string
// is often a good combination of clear and terse and good coverage -- and
// nicely handles and reports an unexpectedly nil value of the error, as well.)
func ShouldBeSameTypeAs(actual interface{}, desire interface{}) (diff string, eq bool) {
	rt_desire := reflect.ValueOf(desire).Type()
	if actual == nil {
		return "got untyped nil; wanted a value of type " + rt_desire.String(), false
	}
	rt_actual := reflect.ValueOf(actual).Type()
	if rt_actual == rt_desire {
		return "", true
	}
	if rt_desire.String() == rt_actual.String() { // these names aren't always unique; use longer ones in that case.
		desire_fullname := rt_desire.PkgPath()
		if desire_fullname != "" {
			desire_fullname += "."
		}
		desire_fullname += rt_desire.Name()
		actual_fullname := rt_actual.PkgPath()
		if actual_fullname != "" {
			actual_fullname += "."
		}
		actual_fullname += rt_actual.Name()
		return "got value of type " + actual_fullname + "; wanted a value of type " + desire_fullname, false
	}
	return "got value of type " + rt_actual.String() + "; wanted a value of type " + rt_desire.String(), false
}
