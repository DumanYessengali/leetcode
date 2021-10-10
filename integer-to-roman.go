package main

func intToRoman(num int) string {
	k := 1
	str := ""
	d := 0

	for num > 0 {
		d = num % 10
		num /= 10

		if k == 1 {
			str = changeToRoman(d, []string{"I", "V", "X"}) + str
		} else if k == 10 {
			str = changeToRoman(d, []string{"X", "L", "C"}) + str
		} else if k == 100 {
			str = changeToRoman(d, []string{"C", "D", "M"}) + str
		} else if k == 1000 {
			str = changeToRoman(d, []string{"M", "", ""}) + str
		}
		k *= 10
	}
	return str
}

func changeToRoman(d int, str []string) string {
	s := ""
	if d < 5 {
		if d == 4 {
			s += str[1]
			return repeat(s, str[0], 1)
		} else {
			return repeat(s, str[0], d)
		}
	} else {
		if d == 5 {
			return str[1]
		} else if d < 9 {
			s += str[1]
			return s + repeat("", str[0], d-5)
		} else {
			return str[0] + str[2]
		}
	}
}

func repeat(s, str string, n int) string {
	for i := 0; i < n; i++ {
		s = str + s
	}
	return s
}
