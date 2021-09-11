package main

func checkPerfectNumber(num int) bool {
	sum := 0
	for i := 1; i*i < num; i++ {
		if num%i == 0 {
			sum += i + num/i
		}
	}
	return sum-num == num
}
