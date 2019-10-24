package piscine

func ConcatParams(args []string) string {
	lent := 0
	for range args {
		lent++
	}

	res := ""
	for index, k := range args {
		res = res + k
		if index < lent-1 {
			res = res + "\n"
		}
	}
	return res
}
