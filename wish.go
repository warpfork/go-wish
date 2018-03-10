package wish

import (
	"fmt"

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
// For pointers, ShouldBe checks pointer equality (*not* a recursion into the values;
// use ShouldEqual to recursively inspect the values).
// For primitives, ShouldBe is the same as ShouldEqual.
// For structs, ShouldBe rejects all values; use ShouldEqual for recursive compares.
//
// Typically, ShouldEqual is more appropriate to use.  ShouldBe is only useful
// if you explicitly wish to check pointer equality.
func ShouldBe(actual interface{}, desire interface{}) (problem string, passed bool) {
	//	actual_rv := reflect.ValueOf(actual)
	//	desire_rv := reflect.ValueOf(desire)
	return "", true
}

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
