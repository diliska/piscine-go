package piscine

func NRune(s string, n int) rune {

	counter := 0
	var res rune
	for _, k := range s {
		counter++
		if counter == n {
			res = k
		}
	}
	return res

}
