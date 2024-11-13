package piscine

func AlphaCount(s string) int {
	l_alphabet := ""
	u_alphabet := ""
	accumulator := 0
	// generate the alphabet
	for i := 'a'; i <= 'z'; i++ {
		l_alphabet += string(i)
	}

	for i := 'A'; i <= 'Z'; i++ {
		u_alphabet += string(i)
	}
	s_to_rune := []rune(s)
	l_alphabet_to_rune := []rune(l_alphabet)
	u_alphabet_to_rune := []rune(u_alphabet)

	for i := 0; i < len(s); i++ {
		for j := 0; j < len(l_alphabet); j++ {
			if (s_to_rune[i] == l_alphabet_to_rune[j]) || (s_to_rune[i] == u_alphabet_to_rune[j]) {
				accumulator++
				break
			}
		}
	}
	return accumulator
}

// uppercases not considered so i should generate abother one and do a for

// chats code

// package piscine

// func AlphaCount(s string) int {
// 	accumulator := 0

// 	for _, char := range s {
// 		if (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') {
// 			accumulator++
// 		}
// 	}
// 	return accumulator
// }
