package main

func maxRotateFunction(nums []int) int {
	var sum, fun int
	l := len(nums)
	for i, v1 := range nums {
		fun += i * v1
		sum += v1
	}
	maximum := fun
	for i := l - 1; i >= 1; i-- {
		fun = fun + sum - l*nums[i]
		maximum = max(fun, maximum)
	}
	return maximum
}
