package piscine

func IsAlpha(str string) bool {

	s := []rune(str)

	res := true

	for _, k := range s {

		if !(('a' <= k && k <= 'z') || ('A' <= k && k <= 'Z') || ('0' <= k && k <= '9')) {
			res = false
			break
		}
	}
	return res
}
