package main

func sumOfUnique(nums []int) int {
	sum := 0
	arrMap := map[int]int{}

	for _, v1 := range nums {
		if arrMap[v1] == 0 {
			arrMap[v1] = 1
		} else {
			arrMap[v1] += 1
		}
	}
	for i, v1 := range arrMap {
		if v1 == 1 {
			sum += i
		}
	}
	return sum
}
