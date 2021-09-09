package main

func isPowerOfThree(n int) bool {
	for n != 1 {
		if n%3 != 0 || n <= 0 {
			break
		}
		n /= 3
	}
	return n == 1
}
