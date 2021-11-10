package main

func rearrangeArray(nums []int) []int {
	nums = mergeS(nums)
	for i := 1; i < len(nums); i += 2 {
		nums[i], nums[i-1] = nums[i-1], nums[i]
	}
	return nums
}
