package piscine

func IterativeFactorial(nb int) int {

	if nb == 0 {
		return 1
	} else {

		fact := nb * IterativeFactorial(nb-1)

		return fact
	}
}
