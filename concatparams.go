package piscine

func ConcatParams(args []string) string {
	res := ""
	for index, k := range args {
		res = res + k
		if index < len(args)-1 {
			res = res + "\n"
		}
	}
	return res
}
