package main

func isHappy(n int) bool {
	myMap := map[int]int{}
	for n != 1 {
		myMap[n]++
		if myMap[n] > 1 {
			return false
		}
		n = Divided(n, 0)
	}
	return true
}

func Divided(num, sum int) int {
	for num > 0 {
		sum += (num % 10) * (num % 10)
		num = num / 10
	}
	return sum
}
