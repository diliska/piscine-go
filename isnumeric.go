package piscine

func IsNumeric(str string) bool {

	s := []rune(str)

	res := true

	for _, k := range s {

		if !('0' <= k && k <= '9') {
			res = false
			break
		}
	}
	return res
}
