package piscine

func IsPrime1(nb int) bool {

	if (nb == 0) || (nb == 1) {
		return false
	} else {
		flag := false
		kol := 1
		for i := 1; i <= nb-1; i++ {
			if nb%i == 0 {
				kol = kol + 1
			}
		}
		if kol == 2 {
			flag = true
		}
		return flag
	}
}

func FindNextPrime(nb int) int {
	res := 0
	if nb <= 1 {
		return 2
	} else {

		for i := nb; i <= 2*nb; i++ {

			if IsPrime1(i) {
				res = i
				break
			}
		}

		return res
	}
}
