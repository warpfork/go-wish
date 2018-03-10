wish: a test assertion library for Go
=====================================

`wish` is a test assertion library for Go(lang), designed to gracefully enhance
the Go standard library `testing` package and behaviors of the `go test` command.


Why another testing library?
----------------------------

It's true -- there are *many* testing libraries in the Go ecosystem already.

Wish is not *that* different.  Nonetheless, there are some reasons to make a clean start:

- `wish` is first established as of go1.10.  There have been several significant
  improvements in the golang `testing` package since the earliest versions of Go
  (`t.Run` trees, and `t.Helper` indicators, to name a few) which move the goalposts.
  Some existing libraries have features that, while they were excellent at the time
  of their creation, are now at odds with the standard ways of doing things.
- `wish` is explicitly aiming to be a *library*, not a *framework*.
  This sets it apart from many other go testing frameworks.
  Concretely, this means `wish` *always* explicitly passes on `*testing.T` reference.
  You'll never have to construct a wrapper object and carry it around.
  This makes `wish` easier to use incrementally, or with other test libraries as desired.
- `wish` wants to focus heavily on ergonomics of output messages on failures.
  For example, when comparing multi-line strings, `wish` will emit a "diff" format.
  When comparing large objects, [go-cmp](https://github.com/google/go-cmp) will be used.
- `wish` assertions (corresponding with the previous point) are generally designed to
  take large objects.  For example, if you have a large slice of strings, with `wish`
  it's preferred to simply make your equality assertions on the entire slice.
  This is more anti-fragile and typing-efficient than making one test-aborting
  assertion on the slice length, then individual asserts on the entries.
- `wish` supports custom comparator functions.  Furthermore, the interface is
  explicitly designed so that *no concrete types from `wish` are required to implement it*.
  This means other projects can provide helpful `wish`-compatable testutil functions,
  and they can do so without incurring an dependency on `wish` -- so you're free to
  provide testutils like this even in library code which must remaining unopinionated.
- `wish` aims to keep a minimal API surface area.  Just asserts.  Most methods take
  wildcard `interface{}` -- because while compile time type checking is nice,
  *this is test code*; if you're about to have a type error, you're already on track
  to see it *very* shortly, and the error messages will point straight to the spot.
- `wish` also aims to provide batteries-included solutions for things like fixture
  data files and temporary filesystem helpers... *in sub-packages*.  The main
  `wish` package shall remain small.  Import the more opinionated helpers only when
  (and if) you need them.
