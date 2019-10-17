package piscine

import "github.com/01-edu/z01"

func PrintStr(str string) {
	strchanged := []rune(str)

	for _, s := range strchanged {
		z01.PrintRune(s)
	}
}
