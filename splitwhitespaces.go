package piscine

func issep(n rune) bool {
	res := false

	if n == 10 || n == 9 || n == 32 {
		res = true
	}
	return res
}

func SplitWhiteSpaces(str string) []string {
	word := ""
	count := 0
	str = str + string(' ')

	for _, k := range str {
		if !issep(k) {
			word = word + string(k)
		} else {
			if word != "" {
				count++
				word = ""
			}
		}
	}

	res := make([]string, count)

	word = ""
	i := 0

	for _, k := range str {
		if !issep(k) {
			word = word + string(k)
		} else {
			if word != "" {
				i++
				res[i-1] = word
				word = ""
			}
		}
	}
	return res
}
