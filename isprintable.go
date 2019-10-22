package piscine

func IsPrintable(str string) bool {

	s := []rune(str)

	res := true

	for _, k := range s {

		if !(32 <= k && k <= 126) {
			res = false
			break
		}
	}
	return res
}
