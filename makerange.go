package piscine

func MakeRange(min, max int) []int {
	var res []int
	if min < max {
		size := max - min
		res = make([]int, size)
		for i := 0; i < size; i++ {
			res[i] = min + i
		}
	}
	return res
}
