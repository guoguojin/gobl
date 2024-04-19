package strings

import (
	"fmt"
	gs "strings"
)

func stripToRune(input string, margin rune) string {
	sb := gs.Builder{}
	nl := '\n'

	strip := false

	for _, r := range input {
		if r == margin {
			sb.WriteRune(r)
			strip = false
			continue
		}

		if r == nl {
			strip = true
			sb.WriteRune(r)
			continue
		}

		if strip {
			continue
		}

		sb.WriteRune(r)
	}

	return sb.String()
}

// StripMargin removes any leading spaces and the specified margin from multi-line strings
// allowing you to write and align multi-line strings nicely in your code similar to the Scala
// StripMargin
func StripMargin(input, margin string) string {
	mrs := []rune(margin)

	in := stripToRune(input, mrs[0])

	m := fmt.Sprintf("%c%s", '\n', margin)
	lines := gs.Split(in, m)

	sb := gs.Builder{}

	for _, l := range lines {
		sb.WriteString(l)
		sb.WriteRune('\n')
	}

	return gs.TrimSpace(gs.Trim(sb.String(), "\n"))
}

// SplitAndTrimSpace splits a string using the provided separators and removes any leading and trailing spaces around the results
func SplitAndTrimSpace(input, sep string) (output []string) {
	if input == "" {
		return output
	}

	values := gs.Split(input, sep)

	output = make([]string, len(values))

	for i, v := range values {
		output[i] = gs.TrimSpace(v)
	}

	return
}

// ToCsv takes a separator and a list of strings and constructs a char separated string
// using the given separator.
func ToCsv(sep rune, in ...string) string {
	b := gs.Builder{}

	for i, s := range in {
		if i > 0 {
			b.WriteRune(sep)
		}

		b.WriteString(s)
	}

	return b.String()
}

type QuoteMark rune

const SingleQuote QuoteMark = '\''
const DoubleQuote QuoteMark = '"'

// ToQuotedCsv takes a separator and the type of quote marks, and a list of strings
// then constructs a char separated string using the provided quote mark
func ToQuotedCsv(sep rune, quote QuoteMark, in ...string) string {
	b := gs.Builder{}

	for i, s := range in {
		if i > 0 {
			b.WriteRune(sep)
		}

		b.WriteRune(rune(quote))
		b.WriteString(s)
		b.WriteRune(rune(quote))
	}

	return b.String()
}

// MkString takes a separator and a list of strings and converts them into a single string
func MkString(sep string, input ...string) string {
	b := gs.Builder{}

	addSep := false

	for _, i := range input {
		if addSep {
			b.WriteString(sep)
		}

		b.WriteString(i)

		addSep = true
	}

	return b.String()
}
