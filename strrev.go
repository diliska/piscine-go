package piscine

func StrRev(s string) string {
	strreversed := ""
	strchanged := []rune(s)
	leng := 0
	for k := range strchanged {
		leng = leng + 1 + k - k

	}
	for i := leng - 1; i >= 0; i-- {
		strreversed = strreversed + string(s[i])
	}

	return strreversed
}
