package piscine

func Join(strs []string, sep string) string {
	res := ""
	count := 0
	for range strs {
		count++
	}

	for index, k := range strs {
		if index != count-1 {
			res = res + k + sep
		} else {
			res = res + k
		}

	}
	return res
}
