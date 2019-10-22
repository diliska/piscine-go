package piscine

func IsLower(str string) bool {

	s := []rune(str)

	res := true

	for _, k := range s {

		if !('a' <= k && k <= 'z') {
			res = false
			break
		}
	}
	return res
}
