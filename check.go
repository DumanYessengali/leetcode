package main

func check(nums []int) bool {
	if isSorted(nums) {
		return true
	}

	for i := 0; i < len(nums); i++ {
		El := nums[0]
		lastEl := nums[len(nums)-1]
		nums[0] = lastEl
		for j := 1; j < len(nums); j++ {
			next := nums[j]
			nums[j] = El
			El = next
		}
		if isSorted(nums) {
			return true
		}

	}
	return false
}
