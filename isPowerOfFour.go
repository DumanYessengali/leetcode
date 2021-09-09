package main

func isPowerOfFour(n int) bool {
	for n != 1 {
		if n%4 != 0 || n <= 0 {
			break
		}
		n /= 4
	}
	return n == 1
}
