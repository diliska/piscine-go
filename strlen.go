package piscine

func StrLen(str string) int {

	strchanged := []rune(str)

	i := 0

	for k := range strchanged {
		i = i + 1 + k - k

	}

	return i
}
