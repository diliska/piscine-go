package piscine

func IsPrime1(nb int) bool {

	if (nb == 0) || (nb == 1) {
		return false
	} else if nb == 2 || nb == 3 {
		return true
	} else {
		flag := true

		for i := 2; i <= nb/2; i++ {
			if nb%i == 0 {
				flag = false
				break
			}
		}
		return flag
	}
}

func FindNextPrime(nb int) int {
	res := 0
	if nb <= 1 {
		return 2
	} else if nb <= 1000000088 {

		for i := nb; i <= 2*nb; i++ {

			if IsPrime1(i) {
				res = i
				break
			}
		}
	}

	return res

}
