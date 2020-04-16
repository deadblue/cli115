package util

import (
	"strings"
)

const (
	charSpace  = ' '
	charEscape = '\\'
)

func InputSplit(input string) (fields []string) {
	fields = make([]string, 0)
	buf := &strings.Builder{}
	inEscape := false

	for _, ch := range input {
		// When in escaping, append the char directly, and reset escape flag.
		if inEscape {
			buf.WriteRune(ch)
			inEscape = false
			continue
		}
		if ch == charSpace {
			if buf.Len() > 0 {
				fields = append(fields, buf.String())
				buf.Reset()
			}
		} else if ch == charEscape {
			inEscape = true
		} else {
			buf.WriteRune(ch)
		}
	}
	fields = append(fields, buf.String())
	return
}

func InputFieldEscape(field string) string {
	buf := strings.Builder{}
	for _, ch := range field {
		if ch == charEscape || ch == charSpace {
			buf.WriteRune(charEscape)
		}
		buf.WriteRune(ch)
	}
	return buf.String()
}
