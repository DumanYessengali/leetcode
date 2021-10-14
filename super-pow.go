package main

func superPow(a int, b []int) int {
	if a == 1 {
		return 1
	}
	a %= 1337
	n := a
	for {
		if b[0] == 0 {
			if len(b) > 1 {
				b = b[1:]
			} else {
				break
			}
		}
		b, _ = minusToArray(b, 1, len(b)-1)
		if len(b) == 1 && b[0] == 0 {
			break
		}
		a *= n
		a %= 1337
	}
	return a
}

func minusToArray(num []int, k, i int) ([]int, int) {
	if i >= 0 {
		if num[i] == 0 {
			num[i] = 19
		} else {
			num[i] -= k
		}
		k = num[i] / 10
		num[i] %= 10
		return minusToArray(num, k, i-1)
	}
	return num, k
}
