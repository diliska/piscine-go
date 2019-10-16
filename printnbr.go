package piscine

import "github.com/01-edu/z01"

func PrintNbr(n int) {
	zn := 1
	if n < 0 {
		z01.PrintRune('-')
		zn = -1
	}
	if n != 0 {
		a := (n / 10) * zn
		if a != 0 {
			PrintNbr(a)
		}
		ost := (n % 10 * zn) + '0'
		z01.PrintRune(rune(ost))
	} else {
		z01.PrintRune('0')
	}
}
