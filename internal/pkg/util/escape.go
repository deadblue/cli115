package util

import "strings"

const (
	StdEscaper = '\\'
)

func Escape(text string, escaper rune, chars string) string {
	charSet := make(map[rune]bool)
	if chars != "" {
		for _, ch := range chars {
			charSet[ch] = true
		}
	}
	sb := strings.Builder{}
	sb.Grow(len(text) * 2)
	for _, r := range text {
		if r == escaper || charSet[r] {
			sb.WriteRune(escaper)
		}
		sb.WriteRune(r)
	}

	return sb.String()
}

func Unescape(text string, escaper rune) string {
	sb, escape := strings.Builder{}, false
	sb.Grow(len(text))
	for _, r := range text {
		escape = !escape && r == escaper
		if !escape {
			sb.WriteRune(r)
		}
	}
	return sb.String()
}

func StdEscape(text string, chars string) string {
	return Escape(text, StdEscaper, chars)
}

func StdUnescape(text string) string {
	return Unescape(text, StdEscaper)
}
