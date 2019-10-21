package piscine

func RecursiveFactorial(nb int) int {
	if (nb < 0) || (nb > 15) {
		return 0
	} else {
		if nb == 0 {
			return 1
		} else {
			fact := nb * RecursiveFactorial(nb-1)

			return fact
		}
	}
}
