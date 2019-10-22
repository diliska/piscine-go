package piscine

func ToLower(s string) string {

	res := []rune(s)

	for index, k := range s {
		if 'A' <= k && k <= 'Z' {
			res[index] = rune(k + 32)
		}

	}
	str := string(res)
	return str
}
