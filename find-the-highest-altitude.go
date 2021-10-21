package main

func largestAltitude(gain []int) int {
	nums := make([]int, len(gain)+1)
	var max int
	nums[0] = 0
	for i := range gain {
		nums[i+1] = gain[i] + nums[i]
		if max < nums[i] {
			max = nums[i]
		}
	}
	if max < nums[len(nums)-1] {
		return nums[len(nums)-1]
	}
	return max
}
