package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	var insertStr string
	order := false
	help := false
	args := []string{}

	// Parse arguments
	for _, arg := range os.Args[1:] {
		switch {
		case arg == "--help" || arg == "-h":
			help = true
		case len(arg) > 9 && arg[:9] == "--insert=":
			insertStr = arg[9:]
		case len(arg) > 3 && arg[:3] == "-i=":
			insertStr = arg[3:]
		case arg == "--order" || arg == "-o":
			order = true
		default:
			args = append(args, arg)
		}
	}

	// Display help information if no arguments are given or if help is requested
	if help || len(os.Args[1:]) == 0 {
		printHelp()
		return
	}

	// Handle argument string (if any)
	var result string
	if len(args) > 0 {
		result = args[0]
	}

	// Insert the specified string if provided
	if insertStr != "" {
		result += insertStr
	}

	// Sort the string in ASCII order if the order flag is set
	if order {
		result = sortString(result)
	}

	// Print the final result
	for _, r := range result {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}

// sortString sorts the string characters in ASCII order without using additional packages
func sortString(s string) string {
	runes := []rune(s)
	for i := 0; i < len(runes)-1; i++ {
		for j := i + 1; j < len(runes); j++ {
			if runes[i] > runes[j] {
				runes[i], runes[j] = runes[j], runes[i]
			}
		}
	}
	return string(runes)
}

// printHelp prints the help information
// printHelp prints the help information with the correct spacing and indentation.
func printHelp() {
	helpText := `--insert
  -i
	 This flag inserts the string into the string passed as argument.
--order
  -o
	 This flag will behave like a boolean, if it is called it will order the argument.`
	for _, r := range helpText {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}
