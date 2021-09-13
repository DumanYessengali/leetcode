package main

import "fmt"

func findDisappearedNumbers(nums []int) []int {
	var arr []int
	arrMap := map[int]int{}

	for _, v1 := range nums {
		if arrMap[v1] == 0 {
			arrMap[v1] = 1
		} else {
			arrMap[v1] += 1
		}
	}
	for i := range nums {
		if arrMap[i+1] == 0 {
			fmt.Println(1)
			arr = append(arr, i+1)
		}
	}
	return arr
}
