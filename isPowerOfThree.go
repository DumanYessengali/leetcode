package main

func isPowerOfThree(n int) bool {
	if n == 0 {
		return false
	} else if n == 1 {
		return true
	} else if n%3 != 0 {
		return false
	}
	for i := 3; i <= n; i *= 3 {
		if i == n {
			return true
		}
	}
	return false
}
