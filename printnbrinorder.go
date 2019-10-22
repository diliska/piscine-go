package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {

	digits := [10]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	if n/10 == 0 {
		z01.PrintRune('0')
	} else {
		f := n
		for f/10 > 0 {

			k := f % 10
			digits[k]++
			f = f / 10
		}
		digits[f]++

		for i := 0; i <= 9; i++ {
			for j := 1; j <= digits[i]; j++ {
				z01.PrintRune(rune('0' + i))
			}
		}
	}
}
