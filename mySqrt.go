package main

func mySqrt(x int) int {
	if x == 0 {
		return 0
	} else if x <= 2 {
		return 1
	}

	for i := 2; i < x; i++ {
		if i*i == x {
			return i
		}
		if i*i > x {
			return i - 1
		}
	}

	return 0
}
