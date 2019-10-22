package piscine

func LastRune(s string) rune {

	counter := 0
	counter1 := 0
	var res rune

	for range s {
		counter++
	}

	for _, k := range s {
		counter1++
		if counter1 == counter {
			res = k
		}
	}
	return res

}
