package main

func findClosestElements(arr []int, k int, x int) []int {
	var nums = make([]int, k)
	n := len(arr)
	l, r := 0, n-k
	for l < r {
		mid := l + (r-l)/2
		if x-arr[mid] <= arr[mid+k]-x {
			r = mid
		} else {
			l = mid + 1
		}
	}

	for i := range nums {
		nums[i] = arr[i+l]
	}

	return nums
}
