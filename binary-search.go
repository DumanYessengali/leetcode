package main

func search(nums []int, target int) int {
	i := 0
	j := len(nums) - 1
	for i <= j {
		if nums[i] == target {
			return i
		}
		if nums[j] == target {
			return j
		}
		i++
		j--
	}
	return -1
}
