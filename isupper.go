package piscine

func IsUpper(str string) bool {

	s := []rune(str)

	res := true

	for _, k := range s {

		if !('A' <= k && k <= 'Z') {
			res = false
			break
		}
	}
	return res
}
