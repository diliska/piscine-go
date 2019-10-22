package piscine

func AlphaCount(str string) int {
	res := 0
	for _, k := range str {
		if (k >= 'a' && k <= 'z') || (k >= 'A' && k <= 'Z') {
			res++
		}
	}
	return res
}
