package wishfix

import (
	"io"
	"strings"

	"github.com/warpfork/go-wish"
)

// marshalHunks writes out wishfix.Hunks in a deterministic way.
func marshalHunks(w io.Writer, h Hunks) error {
	// Write file header.
	w.Write(wordPoundSpace)
	w.Write([]byte(h.title))
	w.Write(wordLF)
	w.Write(wordLF)
	w.Write(wordSectionBreak)
	w.Write(wordLF)

	// Write each section.
	for _, section := range h.sections {
		// Title
		w.Write(wordPoundSpace)
		w.Write([]byte(section.title))
		w.Write(wordLF)
		// Comments (optionally)
		if section.comment != "" {
			lines := strings.Split(section.comment, "\n")
			for _, line := range lines {
				w.Write(wordPoundPoundSpace)
				w.Write([]byte(line))
				w.Write(wordLF)
			}
		}
		// Gap before body
		w.Write(wordLF)

		// Body
		w.Write(wish.IndentBytes(section.body))
		w.Write(wordLF)

		// Always a trailing section break.
		//  (though note the parser is forgiving about this.)
		w.Write(wordSectionBreak)
		w.Write(wordLF)
	}

	return nil
}

var (
	wordLF              = []byte{'\n'}
	wordPoundSpace      = []byte("# ")  // section headers
	wordPoundPoundSpace = []byte("## ") // comments
	wordSectionBreak    = []byte("---")
)
