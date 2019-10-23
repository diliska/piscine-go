package main

import "os"
import "github.com/01-edu/z01"

func main() {

	arguments := os.Args
	count := 0

	for range arguments {
		count++
	}

	for i := 1; i < count; i++ {

		str := arguments[count-i]
		for _, j := range str {
			z01.PrintRune(j)
		}
		z01.PrintRune(10)
	}
}
