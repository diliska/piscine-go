package piscine

func BasicAtoi(s string) int {

	sum := 0
	for _, i := range s {
		ost := 0
		for k := '1'; k <= i; k++ {
			ost = ost + 1
		}

		sum = sum*10 + ost

	}
	return sum

}
