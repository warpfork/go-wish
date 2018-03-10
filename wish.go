package wish

import (
	"fmt"
)

// T is an interface alternative to `*testing.T` -- wherever you see this used,
// use your `*testing.T` object.
type T interface {
	Helper()
	Fail()
	FailNow()
	SkipNow()
	Log(...interface{})
	Name() string

	// Note the lack of `t.Run` in this interface.  Two reasons:
	//  - wish never launches sub-tests; that's for *you* to do (per "library, not framework");
	//  - we... can't really make it useful.  Stdlib's `Run` takes a *concrete* `*testing.T`.
}

// Checker functions compare two objects, report if they "match" (the semantics of
// which should be documented by the Checker function), and if the objects do not
// "match", should provide a descriptive message of how the objects mismatch.
type Checker func(actual interface{}, desire interface{}) (problem string, passed bool)

// Wish makes an assertion that two objects match, using criteria defined by a
// Checker function, and will log information about this to the given T.
//
// Failure to match will log to T, fail the test, and return false (so you
// may take alternative debugging paths, or handle halting on your own).
// Failure to match will *not* cause FailNow; execution will continue.
func Wish(t T, actual interface{}, check Checker, desired interface{}, opts ...options) bool {
	t.Helper()
	problemMsg, passed := check(actual, desired)
	if !passed {
		t.Log(fmt.Sprintf("%s check rejected:\n%s", getCheckerShortName(check), Indent(problemMsg)))
		t.Fail()
	}
	return passed
}

type options interface {
	_options()
}
