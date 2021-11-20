package main

func singleNonDuplicate(nums []int) int {
	result := 0
	for i := range nums {
		result ^= nums[i]
	}
	return result
}
