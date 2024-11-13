package piscine

func Compare(a string, b string) int {
	// return 0 if same string
	if a == b {
		return 0
	}
	a_to_rune := []rune(a)
	b_to_rune := []rune(b)

	// compare the lengths
	var MaxLoop int = len(a)
	if len(b) < len(a) {
		MaxLoop = len(b)
	}
	// ma3reftch 3lach tan39ed loumour
	for i := 0; i < MaxLoop; i++ {
		if a_to_rune[i] > b_to_rune[i] {
			return 1
		} else if a_to_rune[i] < b_to_rune[i] {
			return -1
		}
	}

	if MaxLoop == len(b) {
		return 1
	} // else if MaxLoop == len(a) {  //missing return

	return -1

	// for i := range a_to_rune{

	// }
	// first i need to check if length is the same ?
}

// return 1 if a is ahead of b in dic

// return -1 if a is behind b in dic
