package piscine

func IsPrime(nb int) bool {

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
