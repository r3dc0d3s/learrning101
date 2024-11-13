package piscine

func IsNumeric(s string) bool {
	s_to_rune := []rune(s)
	flag := false
	for i := 0; i < len(s); i++ {
		if s_to_rune[i] >= '0' && s_to_rune[i] <= '9' {
			flag = true
		} else {
			return false
		}
	}

	return flag
}
