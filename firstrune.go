package piscine

func FirstRune(s string) rune {

	counter := 0
	var res rune
	for _, k := range s {
		counter++
		if counter == 1 {
			res = k
		}
	}
	return res

}
