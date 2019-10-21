package piscine

func RecursivePower(nb int, power int) int {
	if power == 0 {
		return 1
	} else {
		res := nb * RecursivePower(nb, power-1)
		return res
	}
}
