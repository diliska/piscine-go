package piscine

func BasicAtoi(s string) int {

	str := []rune(s)

	leng := 0
	for k := range str {
		leng = leng + 1 + k - k
	}

	t := 1
	i := 0

	for p := leng - 1; p >= 0; p-- {

		i += int(str[p]-'0') * t
		t = t * 10
	}
	return i

}
