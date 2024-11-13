package main

import (
	"fmt"

	"piscine"
)

func main() {
	fmt.Println(piscine.Index("Hello!", "l"))
	fmt.Println(piscine.Index("Salut!", "alu"))
	fmt.Println(piscine.Index("Ola!", "hOl"))
}

// func main() {
// 	fmt.Println(piscine.ToUpper("Hello! How are you?"))
// }

// func main() {
// 	fmt.Println(piscine.IsPrintable("Hello"))
// 	fmt.Println(piscine.IsPrintable("Hello\n"))
// }

// func main() {
// 	fmt.Println(piscine.IsNumeric("010203"))
// 	fmt.Println(piscine.IsNumeric("01,02,03"))
// }

// func main() {
// 	fmt.Println(piscine.IsAlpha("Hello! How are you?"))
// 	fmt.Println(piscine.IsAlpha("HelloHowareyou"))
// 	fmt.Println(piscine.IsAlpha("What's this 4?"))
// 	fmt.Println(piscine.IsAlpha("Whatsthis4"))
// }

// func main() {
// 	fmt.Println(piscine.IsUpper("HELLO"))
// 	fmt.Println(piscine.IsUpper("HELLO!"))
// }

// func main() {
// 	fmt.Println(piscine.Concat("Hello!", " How are you?"))
// }

// func main() {
// 	s := "Hello 78 World!    4455 /"
// 	nb := piscine.AlphaCount(s)
// 	fmt.Println(nb)
// }

// func main() {
// 	fmt.Println(piscine.Compare("Hello!", "Hello!"))
// 	fmt.Println(piscine.Compare("Salut!", "lut!"))
// 	fmt.Println(piscine.Compare("Ola!", "Ol"))
// }

// func main() {
// 	z01.PrintRune(piscine.FirstRune("Hello!"))
// 	z01.PrintRune(piscine.FirstRune("Salut!"))
// 	z01.PrintRune(piscine.FirstRune("Ola!"))
// 	z01.PrintRune('\n')
// }

// func main() {
// 	z01.PrintRune(piscine.LastRune("Hello!"))
// 	z01.PrintRune(piscine.LastRune("Salut!"))
// 	z01.PrintRune(piscine.LastRune("Ola!"))
// 	z01.PrintRune('\n')
// }

// func main() {
// 	z01.PrintRune(piscine.NRune("Hello!", 3))
// 	z01.PrintRune(piscine.NRune("Salut!", 2))
// 	z01.PrintRune(piscine.NRune("Bye!", -1)) // out of range ,code compare.go

// 	z01.PrintRune(piscine.NRune("Bye!", 5))
// 	z01.PrintRune(piscine.NRune("Ola!", 4))
// 	z01.PrintRune('\n')
// }
