package main

func nthPersonGetsNthSeat(n int) float64 {
	return 1/float64(n) + (float64(n)-2)/float64(n)*nthPersonGetsNthSeat(n-1)
}
