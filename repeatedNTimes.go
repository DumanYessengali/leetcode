package main

func repeatedNTimes(nums []int) int {
	arrMap := map[int]int{}

	for _, v1 := range nums {
		if arrMap[v1] == 0 {
			arrMap[v1] = 1
		} else {
			arrMap[v1] += 1
		}
	}
	for i, v1 := range arrMap {
		if len(nums)/2 == v1 {
			return i
		}
	}
	return 0
}
