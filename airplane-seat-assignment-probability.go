package main

func nthPersonGetsNthSeat(n int) float64 {
	if n == 1 {
		return 1
	}

	prev := 1.0 / float64(n)

	for i := 1; i < n-1; i++ {
		prev = prev * (1.0 + 1.0/float64(n-i))
	}
	return 1 - prev
}
