package main

func search(nums []int, target int) int {
	i := 0
	j := len(nums) - 1

	for i <= j {
		mid := i + (j-i)/2

		if nums[mid] == target {
			return mid
		}
		if nums[mid] < target {
			i = mid + 1
		}
		if nums[mid] > target {
			j = mid - 1
		}
	}
	return -1
}

func binSearch(num []int, target, i, j int) int {
	if i <= j {
		mid := i + (j-i)/2
		if num[mid] == target {
			return mid
		}
		if num[mid] < target {
			return binSearch(num, target, mid+1, j)
		} else {
			return binSearch(num, target, i, mid-1)
		}
	}
	return -1
}
