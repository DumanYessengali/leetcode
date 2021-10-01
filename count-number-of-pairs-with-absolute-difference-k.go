package main

func countKDifference(nums []int, k int) int {
	l := 0
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]-nums[j] == k || (nums[i]-nums[j])*-1 == k {
				l++
			}
		}
	}
	return l
}
