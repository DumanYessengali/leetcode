package main

func uniqueOccurrences(arr []int) bool {
	arrMap1 := map[int]int{}
	arrMap2 := map[int]int{}

	for i := range arr {
		arrMap1[arr[i]]++
	}
	for _, v1 := range arrMap1 {
		arrMap2[v1]++
		if arrMap2[v1] > 1 {
			return false
		}
	}

	return true
}
