package main

func majorityElement(nums []int) int {
	k := 0
	arrMap := map[int]int{}
	max := 0
	for i := range nums {
		arrMap[nums[i]]++
		if arrMap[nums[i]] > k {
			k = arrMap[nums[i]]
			max = nums[i]
		}
	}
	return max
}
