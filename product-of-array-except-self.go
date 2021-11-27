package main

func productExceptSelf(nums []int) []int {
	n := len(nums)
	var arr []int
	arr = append(arr, 1)
	for i := 1; i < n; i++ {
		arr = append(arr, arr[i-1]*nums[i-1])
	}
	right := 1
	for i := n - 1; i >= 0; i-- {
		arr[i] *= right
		right *= nums[i]
	}
	return arr
}
