package starter

func FirstRune(text string) rune {
	for _, symbol := range text {
		return symbol
	}
	return 0
}

func AppendMarker(values []int, marker int) []int {
	return append(values, marker)
}
