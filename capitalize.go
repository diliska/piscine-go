package piscine

func isletter(a rune) bool {

	if ('a' <= a && a <= 'z') || ('A' <= a && a <= 'Z') || ('0' <= a && a <= '9') {
		return true
	} else {
		return false
	}
}

func caplet(a rune) rune {

	if 'a' <= a && a <= 'z' {
		return rune(a - 32)
	} else {
		return a
	}
}

func lowlet(a rune) rune {

	if 'A' <= a && a <= 'Z' {
		return rune(a + 32)
	} else {
		return a
	}
}

func Capitalize(s string) string {

	prevrune := '!'

	res := []rune(s)

	for index, k := range s {
		if isletter(k) && !(isletter(prevrune)) {
			res[index] = rune(caplet(k))
		} else if isletter(k) && isletter(prevrune) {
			res[index] = rune(lowlet(k))
		}

		prevrune = k
	}
	str := string(res)
	return str
}
