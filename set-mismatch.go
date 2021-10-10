package main

func findErrorNums(nums []int) []int {
	dup, mis := -1, -1
	for i := 1; i <= len(nums); i++ {
		count := 0
		for j := 0; j < len(nums); j++ {
			if nums[j] == i {
				count++
			}
		}
		if count == 2 {
			dup = i
		} else if count == 0 {
			mis = i
		}
		if dup > 0 && mis > 0 {
			break
		}
	}
	return []int{dup, mis}
}
