// Exercise 03: nrune
// Instructions
// Write a function that returns the nth rune of a string. If not possible, it returns 0.

package piscine

func NRune(s string, n int) rune {
	// apparently i also need to test for negative indexes ori should implement conversion
	if len(s) < n || n < 0 {
		return 0
	}
	s_to_rune := []rune(s)
	return s_to_rune[n-1] // index starts at zero
}
