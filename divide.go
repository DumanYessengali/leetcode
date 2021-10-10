package main

func divide(dividend int, divisor int) int {
	if -1<<31 == dividend && divisor == -1 {
		return 1<<31 - 1
	}
	return int(dividend / divisor)
}
