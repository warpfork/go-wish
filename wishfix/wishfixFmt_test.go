package wishfix

import (
	"bytes"
	"testing"

	"github.com/warpfork/go-wish"
)

func TestMarshal(t *testing.T) {
	buf := bytes.Buffer{}
	err := marshalHunks(&buf, Hunks{
		title: "file header",
		sections: []section{
			{title: "section foobar", body: []byte("{\n\t\"woo\": \"zow\",\n\t\"indentation\": \"obviously preserved\",\n\t\"json\": [\"not special\"]\n}\n")},
			{title: "section baz", comment: "this will be a comment", body: []byte("it's all just\nlike, free text\nmaaaan\n")},
		},
	})
	wish.Wish(t, err, wish.ShouldEqual, nil)
	wish.Wish(t, buf.String(), wish.ShouldEqual, wish.Dedent(`
		# file header

		---
		# section foobar

			{
				"woo": "zow",
				"indentation": "obviously preserved",
				"json": ["not special"]
			}

		---
		# section baz
		## this will be a comment

			it's all just
			like, free text
			maaaan

		---
	`))
}
