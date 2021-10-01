package main

func findLucky(arr []int) int {
	nums := map[int]int{}
	for _, v1 := range arr {
		if nums[v1] == 1 {
			nums[v1]++
			continue
		}
		nums[v1]++
	}
	k := -1
	for i, v1 := range nums {
		if i == v1 && k < v1 {
			k = v1
		}
	}

	return k
}
