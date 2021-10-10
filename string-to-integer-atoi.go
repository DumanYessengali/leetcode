package main

func myAtoi(s string) int {
	if len(s) == 0 {
		return 0
	}
	sign := 1
	index := 0
	val := 0
	for len(s) > index && s[index] == ' ' {
		index++
	}
	if index >= len(s) {
		return 0
	}
	if v := s[index]; v == '-' || v == '+' {
		if v == '-' {
			sign = -1
		}

		index++
	}
	for {
		if index == len(s) {
			break
		}
		if s[index] < 48 || s[index] > 57 {
			return sign * val
		}
		val = val*10 + toInt(string(s[index]))
		if 1<<31 <= sign*val {
			return 1<<31 - 1
		} else if (-1)*(1<<31) >= sign*val {
			return (1 << 31) * -1
		}
		index++
	}
	return sign * val
}
