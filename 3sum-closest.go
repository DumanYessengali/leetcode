package main

func threeSumClosest(nums []int, target int) int {
	sort(nums)
	n := nums[0] + nums[1] + nums[2]

	for i := 0; i < len(nums)-2; i++ {
		j, k := i+1, len(nums)-1
		for j < k {
			num := nums[i] + nums[j] + nums[k]
			if abs(num-target) < abs(n-target) {
				n = num
			}
			if num < target {
				j++
			} else {
				k--
			}
		}
	}
	return n
}
