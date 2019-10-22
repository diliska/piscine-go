package piscine

func Compare(a, b string) int {
	var res int
	if a == b {
		res = 0
	} else if a > b {
		res = 1
	} else if a < b {
		res = -1
	}
	return res
}
