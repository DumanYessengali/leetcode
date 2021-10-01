package main

func missingNumber(nums []int) int {
	for i := len(nums); i > 0; i-- {
		for j := 1; j < i; j++ {
			if nums[j] < nums[j-1] {
				nums[j], nums[j-1] = nums[j-1], nums[j]
			}
		}
	}
	j := 0
	for i := 0; i <= nums[len(nums)-1]; i++ {
		if i != nums[j] {
			return i
		}
		j++
	}
	return nums[len(nums)-1] + 1
}
