package main

func findLengthOfLCIS(nums []int) int {
	k := 1
	max := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			k++
		} else {
			k = 1
		}
		if max < k {
			max = k
		}
	}
	return max
}
