package main

import "os"
import "github.com/01-edu/z01"

func main() {

	arguments := os.Args

	for index, k := range arguments {
		if index != 0 {
			str := k
			for _, i := range str {
				z01.PrintRune(i)
			}
			z01.PrintRune(10)
		}
	}
}
