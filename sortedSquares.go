package main

func sortedSquares(nums []int) []int {
	for i, _ := range nums {
		nums[i] *= nums[i]
	}

	for i := len(nums); i > 0; i-- {
		for j := 1; j < len(nums); j++ {
			if nums[j] < nums[j-1] {
				nums[j], nums[j-1] = nums[j-1], nums[j]
			}
		}
	}

	return nums
}
