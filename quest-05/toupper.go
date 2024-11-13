package piscine

func ToUpper(s string) string {
	// substract a from A to get the diff to add and convert
	rune_a := rune('a')
	rune_A := rune('A')
	diff := rune_a - rune_A
	// s_to_upper := make([]rune, len(s)) loooool this is hilarious
	s_to_upper := []rune(s)
	for i := 0; i < len(s); i++ {
		if s_to_upper[i] >= 'a' && s_to_upper[i] <= 'z' {
			s_to_upper[i] -= diff
			// fmt.Print(s_to_upper[i])
		}

		// forgot about special characters ...
		// but if already upper nothing should happen
	}
	return string(s_to_upper)
}
