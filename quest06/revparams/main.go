package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	for i := len(os.Args) - 1; i >= 1; i-- {
		for j := 0; j < len(os.Args[i]); j++ {
			z01.PrintRune(rune(os.Args[i][j]))
		}
		z01.PrintRune('\n')
	}
}
