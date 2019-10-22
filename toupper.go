package piscine

func ToUpper(s string) string {

	res := []rune(s)

	for index, k := range s {
		if 'a' <= k && k <= 'z' {
			res[index] = rune(k - 32)
		}

	}
	str := string(res)
	return str
}
