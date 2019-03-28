package wish

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/warpfork/go-wish/difflib"
)

func strdiff(a, b string) string {
	result, err := difflib.GetUnifiedDiffString(difflib.UnifiedDiff{
		A:       escapishSlice(strings.SplitAfter(a, "\n")),
		B:       escapishSlice(strings.SplitAfter(b, "\n")),
		Context: 3,
	})
	if err != nil {
		panic(fmt.Errorf("diffing failed: %s", err))
	}
	return result
}

func escapish(s string) string {
	q := strconv.Quote(s)
	return q[1:len(q)-1] + "\n"
}

func escapishSlice(ss []string) []string {
	for i, s := range ss {
		ss[i] = escapish(s)
	}
	return ss
}
