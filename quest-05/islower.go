package piscine

func IsLower(s string) bool {
	s_to_rune := []rune(s)
	for i := 0; i < len(s); i++ {
		if s_to_rune[i] < 'a' || s_to_rune[i] > 'z' {
			return false
		}
	}

	return true
}
