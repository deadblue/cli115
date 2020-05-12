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
