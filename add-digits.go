package main

func addDigits(num int) int {
	var result int

	for num > 9 {
		for i := num; i > 0; i /= 10 {
			result += i % 10
		}
		num = result
		result = 0
	}
	return num
}
