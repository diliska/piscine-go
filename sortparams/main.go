package main

import "os"
import "github.com/01-edu/z01"

func sort(a []string, n int) []string {
	res := a
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			if a[i] > a[j] {
				temp := res[i]
				res[i] = res[j]
				res[j] = temp

			}
		}
	}
	return res
}

func printbyrune(s string) {

	for _, k := range s {
		z01.PrintRune(k)
	}
}

func main() {

	arguments := os.Args[1:]
	count := 0

	for range arguments {
		count++
	}

	str := sort(arguments, count)

	for _, j := range str {
		printbyrune(j)
		z01.PrintRune(10)
	}
}
