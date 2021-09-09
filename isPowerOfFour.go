package main

func isPowerOfFour(n int) bool {
	if n == 0 {
		return false
	} else if n == 1 {
		return true
	} else if n%4 != 0 {
		return false
	}
	for i := 4; i <= n; i *= 4 {
		if i == n {
			return true
		}
	}
	return false
}
