package main

func myPow(x float64, n int) float64 {
	sum := 1.00000
	if n == 0 || x == sum {
		return sum
	}
	if n < 0 {
		return 1 / myPow(x, n-1)
	} else {
		return x * myPow(x, n-1)
	}
}
