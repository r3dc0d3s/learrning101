package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	// Copy os.Args (excluding the program name at index 0) into a new slice
	args := os.Args[1:]

	// Simple bubble sort to sort args in ascending order
	for i := 0; i < len(args)-1; i++ {
		for j := 0; j < len(args)-i-1; j++ {
			if args[j] > args[j+1] {
				// Swap args[j] and args[j+1]
				args[j], args[j+1] = args[j+1], args[j]
			}
		}
	}

	// Iterate over the sorted arguments and print each character
	for _, arg := range args {
		for _, char := range arg {
			z01.PrintRune(char)
		}
		z01.PrintRune('\n')
	}
}

// package main

// import (
// 	"os"

// 	"github.com/01-edu/z01"
// )

// func main() {
// 	for i := len(os.Args) - 1; i >= 1; i-- {
// 		for j := 0; j < len(os.Args[i]); j++ {
// 			z01.PrintRune(rune(os.Args[i][j]))
// 		}
// 		z01.PrintRune('\n')
// 	}
// }

// // sorting pseudocode

// // S = [a , b , c]

// // 1 - Calculate the max value in an array

// // max = S[0]

// // loop over S ---> Give Max the next bigger value and so on (not as straightforward)

// // true max is found , youpiee

// // build a clone of S named O (as in ordered)

// // give O[0] the value of Max

// // i need to remove the true max from the list , swap it with index 0 yes i need to swap

// // repeat the process for the res starting at S[]

// // key insight : boolean indices list BRILLIANT BRILLAINT BRILLIANT

// func helperSort([]){

// }
