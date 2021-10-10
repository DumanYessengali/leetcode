package main

func findDuplicates(nums []int) []int {
	maparr := map[int]int{}
	var arr []int

	for _, v1 := range nums {
		maparr[v1]++
		if maparr[v1] >= 2 {
			arr = append(arr, v1)
		}
	}
	return arr
}
