package wish

import (
	"fmt"
	"reflect"

	"github.com/warpfork/go-wish/cmp"
)

// T is an interface alternative to `*testing.T` -- wherever you see this used,
// use your `*testing.T` object.
type T interface {
	Helper()
	FailNow()
	SkipNow()
	Name() string

	// Note the lack of `t.Run` in this interface.  Two reasons:
	//  - wish never launches sub-tests; that's for *you* to do (per "library, not framework");
	//  - we... can't really make it useful.  Stdlib's `Run` takes a *concrete* `*testing.T`.
}

// Wish makes an assertion that two objects match, using criteria defined by a
// Checker function, and will log information about this to the given T.
//
// Failure to match will not halt the test; it will only log to T, and return
// false (so you may take alternative debugging paths, or handle halting on your
// own).
func Wish(t T, actual interface{}, check Checker, desired interface{}, opts ...options) bool {
	t.Helper()
	problemMsg, passed := check(actual, desired)
	if !passed {
		fmt.Printf("%s check rejected:\n%s\n", getCheckerShortName(check), problemMsg)
	}
	return passed
}

// Checker functions compare two objects, report if they "match" (the semantics of
// which should be documented by the Checker function), and if the objects do not
// "match", should provide a descriptive message of how the objects mismatch.
type Checker func(actual interface{}, desire interface{}) (problem string, passed bool)

var (
	_ Checker = ShouldBe
	_ Checker = ShouldEqual
)

type options interface {
	_options()
}

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
func ShouldEqual(actual interface{}, desire interface{}) (string, bool) {
	s1, ok1 := actual.(string)
	s2, ok2 := desire.(string)
	if ok1 && ok2 {
		diff := strdiff(s1, s2)
		return diff, diff == ""
	}
	diff := cmp.Diff(actual, desire)
	return diff, diff == ""
}
