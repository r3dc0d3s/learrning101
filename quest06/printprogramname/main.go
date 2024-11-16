package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arg_0_to_rune := []rune(os.Args[0])
	for i := 2; i < len(arg_0_to_rune); i++ {
		z01.PrintRune(arg_0_to_rune[i])
	}
	z01.PrintRune('\n')
}
