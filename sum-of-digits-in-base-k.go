package main

func sumBase(n int, k int) int {
	var sum int
	for n > 0 {
		sum += n % k
		n /= k
	}
	return sum
}
