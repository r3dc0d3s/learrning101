package piscine

func Concat(str1 string, str2 string) string {
	str1_to_rune := []rune(str1)
	str2_to_rune := []rune(str2)
	// total_l :=
	concat := make([]rune, len(str1)+len(str2))
	for i := 0; i < len(str1); i++ {
		concat[i] = str1_to_rune[i]
		// stuck on which str len to use
	}

	// well couldnt include both in one loop
	for i := 0; i < len(str2); i++ {
		concat[i+len(str1)] = str2_to_rune[i]
		// stuck on which str len to use
	}

	return string(concat)
}
