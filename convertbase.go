package piscine

func power(a, b int) int {
	res := 1
	if b == 0 {
		res = 1
	} else if a == 0 {
		res = 0
	} else {
		for i := 1; i <= b; i++ {
			res = res * a
		}
	}
	return res
}

func todecimal(s, base string) int {
	lens := 0
	res := 0
	baseint := 0
	var temp int

	for range base {
		baseint++
	}

	for range s {
		lens++
	}

	for i := 0; i < lens; i++ {
		if '0' <= s[lens-i-1] && s[lens-i-1] <= '9' {
			temp = int(s[lens-1-i] - '0')
		} else {
			temp = int(s[lens-i-1] - 'A')
		}
		temp = temp * power(baseint, i)
		res = res + temp
	}
	return res
}

func inttostr(n int) string {
	res := ""

	f := n

	for f/10 > 0 {
		res = res + string(f%10+'0')
		f = f / 10
	}

	res = res + string(f+'0')
	res = reverse(res)

	return res
}

//----

func nbrtobase(n int, s string) string {
	lens := 0
	res1 := ""
	res := ""

	for range s {
		lens++
	}

	if lens == 10 {
		res = inttostr(n)
	} else {
		f := n
		for f/lens > 0 {
			if 0 <= f%lens && f%lens <= 9 {
				res1 = res1 + string(f%lens+'0')
			} else {
				res1 = res1 + string(f%lens+'A'-10)
			}
			f = f / lens
		}
		if 0 <= f && f <= 9 {
			res1 = res1 + string(f+'0')
		} else {
			res1 = res1 + string(f+'A'-10)
		}
		res = reverse(res1)
	}
	return res
}

//---

func reverse(s string) string {
	res := ""
	lens := 0

	for range s {
		lens++
	}

	for i := lens - 1; i >= 0; i-- {
		res = res + string(s[i])
	}
	return res
}

func ConvertBase(nbr, baseFrom, baseTo string) string {
	nbrdecimal := todecimal(nbr, baseFrom)
	res := nbrtobase(nbrdecimal, baseTo)
	return res
}
