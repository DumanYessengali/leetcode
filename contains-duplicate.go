package main

func containsDuplicate(nums []int) bool {
	m := map[int]int{}
	for _, v1 := range nums {
		if m[v1] == 1 {
			return true
		}
		m[v1]++
	}
	return false
}
