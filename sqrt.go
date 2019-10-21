package piscine

func Sqrt(nb int) int {

	if nb <= 0 {
		return 0
	} else {

		res := 0

		for i := 1; i <= nb; i++ {

			if i*i == nb {
				res = i
				break
			}
		}

		return res
	}
}
