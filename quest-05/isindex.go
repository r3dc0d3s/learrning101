package piscine

func Index(s string, toFind string) int {
	s_to_rune := []rune(s)      // ['H' (0) , 'e' (1), 'l' (2), 'l' (3), 'o' (4), ' ' (6), 'W' (7), ... ]
	f_to_rune := []rune(toFind) // ['H' (0) , 'e' (1), 'l' (2), 'l' (3), 'o' (4)]

	for i := 0; i < len(s); i++ {
		count := 0
		for j := 0; j < len(toFind); j++ {
			if s_to_rune[i+j] != f_to_rune[j] { // shit this where i+j is needed
				break
			}
			count++
		}
		if count == len(toFind) {
			return i
		}
	}

	return -1
}

//     s = "Hello World"
//     toFind = "Hello"
