package piscine

func LastRune(s string) rune {
	s_to_rune := []rune(s)
	return s_to_rune[len(s)-1]
}
