package main

func jump(nums []int) int {
	var i, maxReach int

	for i < len(nums) && i <= maxReach {
		maxReach = max(i+nums[i], maxReach)
		i++

	}
	return i
}
