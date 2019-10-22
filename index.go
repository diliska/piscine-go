package piscine

func Index(s string, toFind string) int {
	lens := 0
	lentf := 0
	res := -1

	for range s {
		lens++
	}

	for range toFind {
		lentf++
	}

	for i := 0; i <= lens-lentf; i++ {
		if s[i:i+lentf] == toFind {
			res = i
			break
		}
	}

	return res
}
