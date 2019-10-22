package piscine

func BasicJoin(strs []string) string {
	res := ""
	for _, k := range strs {
		res = res + k
	}
	return res
}
