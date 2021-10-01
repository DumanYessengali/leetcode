package main

func singleNumber(nums []int) []int {
	arrMap := map[int]int{}
	for _, v1 := range nums {
		if arrMap[v1] > 0 {
			arrMap[v1] += 1
			continue
		}
		arrMap[v1] = 1
	}
	var arr []int
	for i, v1 := range arrMap {
		if v1 == 1 {
			arr = append(arr, i)
		}
	}
	return arr
}
