package piscine

func IsPrintable(s string) bool {
	s_to_rune := []rune(s)
	flag := false
	for i := 0; i < len(s); i++ {
		if s_to_rune[i] >= 32 && s_to_rune[i] <= 127 {
			flag = true
		} else {
			return false
		}
	}

	return flag
}
