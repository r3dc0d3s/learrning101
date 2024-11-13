package piscine

func IsUpper(s string) bool {
	s_to_rune := []rune(s)
	for i := 0; i < len(s); i++ {
		if s_to_rune[i] < 'A' || s_to_rune[i] > 'Z' {
			return false
		}
	}

	return true
}
