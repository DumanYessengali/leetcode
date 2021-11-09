package main

func findKthLargest(nums []int, k int) int {
	sort(nums)
	return nums[len(nums)-k]
}
