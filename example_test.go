package wish_test

import (
	"testing"

	"github.com/warpfork/go-wish"
)

func ExampleThing() {
	t := &testing.T{}
	actual := "foobar"
	objective := "bazfomp"
	wish.Wish(t, actual, wish.ShouldBe, objective) // commentary

	// Output:
}
