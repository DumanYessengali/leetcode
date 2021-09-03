package main

func canBeIncreasing(nums []int) bool {
	for i := range nums {
		if i > 0 {
			arr := make([]int, 0, len(nums)-1)
			arr = append(arr, nums[:i-1]...)
			arr = append(arr, nums[i:]...)
			if isSorted(arr) {
				return true
			}
		}
	}
	arr := make([]int, 0, len(nums)-1)
	arr = append(arr, nums[:len(nums)-1]...)
	if isSorted(arr) {
		return true
	}
	return false
}

func isSorted(arr []int) bool {
	k := 0
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] < arr[i+1] {
			k++
		}
	}
	if k+1 == len(arr) {
		return true
	}
	return false
}
