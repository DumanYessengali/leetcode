package main

func searchInsert(nums []int, target int) int {
	for i, v1 := range nums {
		if v1 == target {
			return i
		} else if target < v1 {
			return i
		}
	}
	return len(nums)
}
