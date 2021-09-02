package main

func createTargetArray(nums []int, index []int) []int {
	var target []int
	for i, v := range index {
		target = append(target, 0)
		copy(target[v+1:], target[v:])
		target[v] = nums[i]
	}

	return target
}
