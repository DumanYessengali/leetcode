package main

func findErrorNums(nums []int) []int {
	var arr []int
	for i := 1; i <= len(nums); i++ {
		if nums[i-1] != i {
			arr = append(arr, nums[i-1])
			arr = append(arr, i)
			break
		}
	}
	return arr
}
