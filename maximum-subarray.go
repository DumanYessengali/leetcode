package main

func maxSubArray(nums []int) int {
	sum := nums[0]
	max := nums[0]
	for i := 1; i < len(nums); i++ {
		if sum+nums[i] > nums[i] {
			sum = sum + nums[i]
		} else {
			sum = nums[i]
		}

		if sum > max {
			max = sum
		}
		if sum < 0 {
			sum = 0
		}
	}
	return max
}
