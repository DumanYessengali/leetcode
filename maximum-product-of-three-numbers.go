package main

func maximumProduct(nums []int) int {
	for i := len(nums); i > 0; i-- {
		for j := 1; j < i; j++ {
			if nums[j] > nums[j-1] {
				nums[j], nums[j-1] = nums[j-1], nums[j]
			}
		}
	}
	return nums[0] * nums[1] * nums[2]
}
