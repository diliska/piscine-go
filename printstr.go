package piscine

import "github.com/01-edu/z01"

func PrintStr(str string) {
	strchanged := []rune(str)

	for i := 0; i < len(str); i++ {
		z01.PrintRune(strchanged[i])
	}
	z01.PrintRune(10)
}
