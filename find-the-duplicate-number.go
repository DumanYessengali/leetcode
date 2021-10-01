package main

func findDuplicate(nums []int) int {
	arrMap := map[int]int{}

	for _, v1 := range nums {
		arrMap[v1]++
		if arrMap[v1] == 2 {
			return v1
		}
	}
	return nums[0]
}
