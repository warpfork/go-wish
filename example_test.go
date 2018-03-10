package wish_test

import (
	"fmt"
	"testing"

	"github.com/warpfork/go-wish"
)

func ExampleThing() {
	t := &testing.T{}
	actual := "foobar"
	objective := "bazfomp"
	fmt.Printf("%v\n", wish.Wish(t, actual, wish.ShouldEqual, objective))

	// Output:
	// ShouldEqual check rejected:
	// @@ -1 +1 @@
	// -foobar
	// +bazfomp
	//
	// false
}

func ExampleMultilineString() {
	t := &testing.T{}
	actual := "foobar\nwoop\nwow"
	objective := "bazfomp\nwoop\nwowdiff"
	fmt.Printf("%v\n", wish.Wish(t, actual, wish.ShouldEqual, objective))

	// Output:
	// ShouldEqual check rejected:
	// @@ -1,3 +1,3 @@
	// -foobar\n
	// +bazfomp\n
	//  woop\n
	// -wow
	// +wowdiff
	//
	// false
}

func ExampleWish_ShouldEqual_CompareStructs() {
	t := &testing.T{}
	actual := struct{ Baz string }{"asdf"}
	objective := struct{ Baz string }{"asdf"}
	fmt.Printf("%v\n", wish.Wish(t, actual, wish.ShouldEqual, objective))

	// Output:
	// true
}

func ExampleWish_ShouldEqual_CompareStructsReject() {
	t := &testing.T{}
	actual := struct{ Bar string }{"asdf"}
	objective := struct{ Baz string }{"qwer"}
	fmt.Printf("%v\n", wish.Wish(t, actual, wish.ShouldEqual, objective))

	// Output:
	// ShouldEqual check rejected:
	// :
	//	-: struct { Bar string }{Bar: "asdf"}
	//	+: struct { Baz string }{Baz: "qwer"}
	//
	// false
}

func ExampleWish_ShouldEqual_TypeMismatch() {
	t := &testing.T{}
	actual := "foobar"
	objective := struct{}{}
	fmt.Printf("%v\n", wish.Wish(t, actual, wish.ShouldEqual, objective))

	// Output:
	// ShouldEqual check rejected:
	// :
	// 	-: "foobar"
	// 	+: struct {}{}
	//
	// false
}
