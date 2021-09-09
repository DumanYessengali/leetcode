package main

func isPowerOfTwo(n int) bool {
	if n == 0 {
		return false
	} else if n == 1 {
		return true
	} else if n%2 != 0 {
		return false
	}
	for i := 2; i <= n; i *= 2 {
		if i == n {
			return true
		}
	}
	return false
}
