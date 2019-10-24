package piscine

import "github.com/01-edu/z01"

func PrintWordsTables(table []string) {

	for _, k := range table {
		for _, sym := range k {
			z01.PrintRune(sym)
		}
		z01.PrintRune(10)
	}
}
