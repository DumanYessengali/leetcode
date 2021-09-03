package main

func findGCD(nums []int) int {
	max := nums[0]
	min := nums[0]
	for _, v1 := range nums {
		if v1 > max {
			max = v1
		}
		if v1 < min {
			min = v1
		}
	}

	div := 0

	for i := 1; i <= min; i++ {
		if max%i == 0 && min%i == 0 {
			div = i
		}
	}
	return div
}
