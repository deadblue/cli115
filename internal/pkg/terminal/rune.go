package terminal

func StringLeftRunes(str string, end int) string {
	runes := []rune(str)
	return string(runes[0:end])
}

func StringRightRunes(str string, start int) string {
	runes := []rune(str)
	return string(runes[start:])
}
