package terminal

import "strings"

const (
	charSpace  = ' '
	charEscape = '\\'
)

/*
Split an input line by space, unless the space is escaping.
The heading spaces will be ignored, while the tailing spaces will be remained.

In the splited parts, escaped-string won't be unescaped.

For example:
	Input "a b c" will be split into ["a", "b", "c"]
	Input "a\ b c" will be split into ["a\ b", "c"]
	Input "a b " will be split into ["a", "b", " "]
	Input " a b" will be split into ["a", "b"]
*/
func split(line string) (fields []string) {
	fields = make([]string, 0)
	sb, escape := strings.Builder{}, false
	for _, ch := range line {
		if !escape && ch == charSpace {
			if sb.Len() > 0 {
				fields = append(fields, sb.String())
				sb.Reset()
			}
		} else {
			sb.WriteRune(ch)
			escape = !escape && ch == charEscape
		}
	}
	fields = append(fields, sb.String())
	return
}

// Escape space in str.
// A string looks like "a b" will be "a\ b".
func Escape(str string) string {
	sb := strings.Builder{}
	sb.Grow(len(str) * 2)
	for _, ch := range str {
		if ch == charEscape || ch == charSpace {
			sb.WriteRune(charEscape)
		}
		sb.WriteRune(ch)
	}
	return sb.String()
}

// Unescape space in str.
// A string looks like "a\ b" will get "a b".
func Unescape(str string) string {
	sb, escape := strings.Builder{}, false
	sb.Grow(len(str))
	for _, ch := range str {
		if escape {
			sb.WriteRune(ch)
			escape = false
			continue
		}
		if ch == charEscape {
			escape = true
		} else {
			sb.WriteRune(ch)
		}
	}
	return sb.String()
}
