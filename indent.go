package wish

import (
	"bytes"
	"strings"
)

// Indent prepends one tab character to each line of a string.
func Indent(s string) string {
	if s == "" {
		return "\t"
	}
	lines := strings.SplitAfter(s, "\n")
	buf := bytes.Buffer{}
	for _, line := range lines {
		if line != "" {
			buf.WriteByte('\t')
		}
		buf.WriteString(line)
	}
	return buf.String()
}

// Dedent strips leading tabs from every line of a string, taking a hint of
// how many tabs should be stripped from the number of consecutive tabs found
// on the first non-empty line.  Dedent also strips one leading blank
// line if it contains nothing but the linebreak.
//
// If later lines have fewer leading tab characters than the depth we intuited
// from the first line, then stripping will still only remove tab characters.
//
// Roughly, Dedent is "Do What I Mean" to normalize a heredoc string
// that contains leading indentation to make it congruent with the
// surrounding source code.
func Dedent(s string) string {
	lines := strings.SplitAfter(s, "\n")
	buf := bytes.Buffer{}
	if lines[0] == "\n" {
		lines = lines[1:]
	}
	if len(lines) == 0 {
		return ""
	}
	depth := 0
	for _, r := range lines[0] {
		depth++
		if r != '\t' {
			depth--
			break
		}
	}
	for _, line := range lines {
		for i, r := range line {
			if i < depth && r == '\t' {
				continue
			}
			buf.WriteString(line[i:])
			break
		}
	}
	return buf.String()
}
