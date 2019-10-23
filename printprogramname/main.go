package main

import "os"
import "github.com/01-edu/z01"

func main() {

	arguments:=os.Args

	str:=arguments [0]

	for _, k:=range str {
		z01.PrintRune(k)
	}
z01.PrintRune(10)
}
